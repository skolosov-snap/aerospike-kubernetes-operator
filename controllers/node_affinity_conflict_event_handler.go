package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"strings"
)

var (
	nodeAffinityConflicts = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "node_affinity_conflict",
			Help: "Number of Node Affinity Conflicts detected",
		},
	)
	deletedPvs = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "deleted_pvc",
			Help: "Number of PVCs deleted",
		},
	)
	resolveSuccess = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "node_affinity_conflict_success",
			Help: "Number of Node Affinity Conflicts resolved",
		},
	)
	resolveError = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "node_affinity_conflict_error",
			Help: "Number of Node Affinity Conflicts not resolved",
		},
	)
)

type EventWatcher struct {
	kubeClient *kubernetes.Clientset
	log        logr.Logger
}

func NewEventWatcher(factory informers.SharedInformerFactory, kubeClient *kubernetes.Clientset, log logr.Logger) *EventWatcher {
	watcher := &EventWatcher{
		kubeClient: kubeClient,
		log:        log,
	}
	factory.Core().V1().Events().Informer().AddEventHandler(watcher)
	metrics.Registry.MustRegister(nodeAffinityConflicts, resolveSuccess, resolveError, deletedPvs)
	return watcher
}

func (e *EventWatcher) OnAdd(obj interface{}) {
	event, ok := obj.(*corev1.Event)
	if !ok {
		e.log.Error(fmt.Errorf("unexpected type %T of received event", obj), "can't handle new event")
		return
	}
	e.onEvent(event)
}

func (e *EventWatcher) OnUpdate(oldObj, newObj interface{}) {
	// Ignore updates (very noisy, we get everything via adds anyway)
}

func (e *EventWatcher) OnDelete(obj interface{}) {
	// Ignore deletes
}

func (e *EventWatcher) onEvent(event *corev1.Event) {
	if isNodeAffinityConflictEvent(event) {
		nodeAffinityConflicts.Inc()
		e.log.Info("Received node affinity conflict event: ", "event", event.InvolvedObject.String())

		pod, err := e.getPodForEvent(event.InvolvedObject)
		if err != nil {
			e.log.Error(err, "Stopping processing due to issue with pod.")
			resolveError.Inc()
			return
		}

		if err = e.deleteConflictingPVCs(pod); err != nil {
			e.log.Error(err, "Stopping processing due to issue with pvcs.")
			resolveError.Inc()
			return
		}

		err = e.restartPod(pod)
		if err != nil {
			e.log.Error(err, "Error during node affinity conflict resolution. Conflict is not resolved.")
			resolveError.Inc()
			return
		}
		resolveSuccess.Inc()
	}
	return
}

func isNodeAffinityConflictEvent(event *corev1.Event) bool {
	return event.Reason == "FailedScheduling" && strings.Contains(event.Message, "volume node affinity conflict")
}

func (e *EventWatcher) getPodForEvent(obj corev1.ObjectReference) (*corev1.Pod, error) {
	pod, err := e.kubeClient.CoreV1().Pods(obj.Namespace).Get(context.TODO(), obj.Name, metav1.GetOptions{})
	if err != nil {
		e.log.Error(err, "Error retrieving pvcs")
		return pod, err
	}
	e.log.Info("Found Pod: ", "pod", pod.String())
	if pod.Status.Phase != corev1.PodPending {
		return pod, errors.New("pod is not in pending phase")
	}
	if pod.UID != obj.UID {
		return pod, fmt.Errorf("pod uid (%s) does not match event uid (%s).", pod.UID, obj.UID)
	}
	if pod.Labels["app"] != "aerospike-cluster" {
		return pod, errors.New("pod does not belong to the aerospike operator")
	}

	return pod, nil
}

func (e *EventWatcher) deleteConflictingPVCs(pod *corev1.Pod) error {
	pvcs, err := e.kubeClient.CoreV1().PersistentVolumeClaims(pod.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	// Locate pvcs for the pod
	podPvcs := extractPvcsFromPod(pod)
	numDeleted := 0
	for _, pvc := range pvcs.Items {
		if pvc.Spec.StorageClassName != nil && "local-ssd" == *pvc.Spec.StorageClassName && podPvcs[pvc.Name] {
			e.log.Info("Found conflicting local ssd pvc.", "pvc", pvc.String())
			if err := e.kubeClient.CoreV1().PersistentVolumeClaims(pvc.Namespace).Delete(context.TODO(), pvc.Name, metav1.DeleteOptions{}); err != nil {
				e.log.Error(err, "Error deleting pvc - unable to resolve node affinity conflict.", "namespace", pvc.Namespace, "name", pvc.Name)
				return err
			}
			deletedPvs.Inc()
			numDeleted++
		}
	}

	if numDeleted == 0 {
		return errors.New("did not locate any conflicting local ssd pvcs")
	}

	return nil
}

func (e *EventWatcher) restartPod(pod *corev1.Pod) error {
	if err := e.kubeClient.CoreV1().Pods(pod.Namespace).Delete(context.TODO(), pod.Name, metav1.DeleteOptions{}); err != nil {
		e.log.Error(err, "Error deleting pod - unable to resolve node affinity conflict.", "namespace", pod.Namespace, "name", pod.Name)
		return err
	}

	return nil
}

func extractPvcsFromPod(pod *corev1.Pod) map[string]bool {
	podPvcs := make(map[string]bool, len(pod.Spec.Volumes))
	for _, volume := range pod.Spec.Volumes {
		if volume.PersistentVolumeClaim != nil && len(volume.PersistentVolumeClaim.ClaimName) > 0 {
			podPvcs[volume.PersistentVolumeClaim.ClaimName] = true
		}
	}
	return podPvcs
}
