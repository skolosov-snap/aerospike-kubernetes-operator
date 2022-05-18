// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeAccessControlSpec) DeepCopyInto(out *AerospikeAccessControlSpec) {
	*out = *in
	if in.AdminPolicy != nil {
		in, out := &in.AdminPolicy, &out.AdminPolicy
		*out = new(AerospikeClientAdminPolicy)
		**out = **in
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]AerospikeRoleSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]AerospikeUserSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeAccessControlSpec.
func (in *AerospikeAccessControlSpec) DeepCopy() *AerospikeAccessControlSpec {
	if in == nil {
		return nil
	}
	out := new(AerospikeAccessControlSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeCertPathInOperatorSource) DeepCopyInto(out *AerospikeCertPathInOperatorSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeCertPathInOperatorSource.
func (in *AerospikeCertPathInOperatorSource) DeepCopy() *AerospikeCertPathInOperatorSource {
	if in == nil {
		return nil
	}
	out := new(AerospikeCertPathInOperatorSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeClientAdminPolicy) DeepCopyInto(out *AerospikeClientAdminPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeClientAdminPolicy.
func (in *AerospikeClientAdminPolicy) DeepCopy() *AerospikeClientAdminPolicy {
	if in == nil {
		return nil
	}
	out := new(AerospikeClientAdminPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeCluster) DeepCopyInto(out *AerospikeCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeCluster.
func (in *AerospikeCluster) DeepCopy() *AerospikeCluster {
	if in == nil {
		return nil
	}
	out := new(AerospikeCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AerospikeCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeClusterList) DeepCopyInto(out *AerospikeClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AerospikeCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeClusterList.
func (in *AerospikeClusterList) DeepCopy() *AerospikeClusterList {
	if in == nil {
		return nil
	}
	out := new(AerospikeClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AerospikeClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeClusterSpec) DeepCopyInto(out *AerospikeClusterSpec) {
	*out = *in
	in.Storage.DeepCopyInto(&out.Storage)
	if in.AerospikeAccessControl != nil {
		in, out := &in.AerospikeAccessControl, &out.AerospikeAccessControl
		*out = new(AerospikeAccessControlSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AerospikeConfig != nil {
		in, out := &in.AerospikeConfig, &out.AerospikeConfig
		*out = (*in).DeepCopy()
	}
	if in.ValidationPolicy != nil {
		in, out := &in.ValidationPolicy, &out.ValidationPolicy
		*out = new(ValidationPolicySpec)
		**out = **in
	}
	in.RackConfig.DeepCopyInto(&out.RackConfig)
	out.AerospikeNetworkPolicy = in.AerospikeNetworkPolicy
	if in.OperatorClientCertSpec != nil {
		in, out := &in.OperatorClientCertSpec, &out.OperatorClientCertSpec
		*out = new(AerospikeOperatorClientCertSpec)
		(*in).DeepCopyInto(*out)
	}
	in.PodSpec.DeepCopyInto(&out.PodSpec)
	in.SeedsFinderServices.DeepCopyInto(&out.SeedsFinderServices)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeClusterSpec.
func (in *AerospikeClusterSpec) DeepCopy() *AerospikeClusterSpec {
	if in == nil {
		return nil
	}
	out := new(AerospikeClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeClusterStatus) DeepCopyInto(out *AerospikeClusterStatus) {
	*out = *in
	in.AerospikeClusterStatusSpec.DeepCopyInto(&out.AerospikeClusterStatusSpec)
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make(map[string]AerospikePodStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeClusterStatus.
func (in *AerospikeClusterStatus) DeepCopy() *AerospikeClusterStatus {
	if in == nil {
		return nil
	}
	out := new(AerospikeClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeClusterStatusSpec) DeepCopyInto(out *AerospikeClusterStatusSpec) {
	*out = *in
	in.Storage.DeepCopyInto(&out.Storage)
	if in.AerospikeAccessControl != nil {
		in, out := &in.AerospikeAccessControl, &out.AerospikeAccessControl
		*out = new(AerospikeAccessControlSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AerospikeConfig != nil {
		in, out := &in.AerospikeConfig, &out.AerospikeConfig
		*out = (*in).DeepCopy()
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.ValidationPolicy != nil {
		in, out := &in.ValidationPolicy, &out.ValidationPolicy
		*out = new(ValidationPolicySpec)
		**out = **in
	}
	in.RackConfig.DeepCopyInto(&out.RackConfig)
	out.AerospikeNetworkPolicy = in.AerospikeNetworkPolicy
	if in.OperatorClientCertSpec != nil {
		in, out := &in.OperatorClientCertSpec, &out.OperatorClientCertSpec
		*out = new(AerospikeOperatorClientCertSpec)
		(*in).DeepCopyInto(*out)
	}
	in.PodSpec.DeepCopyInto(&out.PodSpec)
	in.SeedsFinderServices.DeepCopyInto(&out.SeedsFinderServices)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeClusterStatusSpec.
func (in *AerospikeClusterStatusSpec) DeepCopy() *AerospikeClusterStatusSpec {
	if in == nil {
		return nil
	}
	out := new(AerospikeClusterStatusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeConfigSpec) DeepCopyInto(out *AerospikeConfigSpec) {
	clone := in.DeepCopy()
	*out = *clone
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeContainerSpec) DeepCopyInto(out *AerospikeContainerSpec) {
	*out = *in
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeContainerSpec.
func (in *AerospikeContainerSpec) DeepCopy() *AerospikeContainerSpec {
	if in == nil {
		return nil
	}
	out := new(AerospikeContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeInstanceSummary) DeepCopyInto(out *AerospikeInstanceSummary) {
	*out = *in
	if in.AccessEndpoints != nil {
		in, out := &in.AccessEndpoints, &out.AccessEndpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AlternateAccessEndpoints != nil {
		in, out := &in.AlternateAccessEndpoints, &out.AlternateAccessEndpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLSAccessEndpoints != nil {
		in, out := &in.TLSAccessEndpoints, &out.TLSAccessEndpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TLSAlternateAccessEndpoints != nil {
		in, out := &in.TLSAlternateAccessEndpoints, &out.TLSAlternateAccessEndpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeInstanceSummary.
func (in *AerospikeInstanceSummary) DeepCopy() *AerospikeInstanceSummary {
	if in == nil {
		return nil
	}
	out := new(AerospikeInstanceSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeNetworkPolicy) DeepCopyInto(out *AerospikeNetworkPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeNetworkPolicy.
func (in *AerospikeNetworkPolicy) DeepCopy() *AerospikeNetworkPolicy {
	if in == nil {
		return nil
	}
	out := new(AerospikeNetworkPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeObjectMeta) DeepCopyInto(out *AerospikeObjectMeta) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeObjectMeta.
func (in *AerospikeObjectMeta) DeepCopy() *AerospikeObjectMeta {
	if in == nil {
		return nil
	}
	out := new(AerospikeObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeOperatorCertSource) DeepCopyInto(out *AerospikeOperatorCertSource) {
	*out = *in
	if in.SecretCertSource != nil {
		in, out := &in.SecretCertSource, &out.SecretCertSource
		*out = new(AerospikeSecretCertSource)
		**out = **in
	}
	if in.CertPathInOperator != nil {
		in, out := &in.CertPathInOperator, &out.CertPathInOperator
		*out = new(AerospikeCertPathInOperatorSource)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeOperatorCertSource.
func (in *AerospikeOperatorCertSource) DeepCopy() *AerospikeOperatorCertSource {
	if in == nil {
		return nil
	}
	out := new(AerospikeOperatorCertSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeOperatorClientCertSpec) DeepCopyInto(out *AerospikeOperatorClientCertSpec) {
	*out = *in
	in.AerospikeOperatorCertSource.DeepCopyInto(&out.AerospikeOperatorCertSource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeOperatorClientCertSpec.
func (in *AerospikeOperatorClientCertSpec) DeepCopy() *AerospikeOperatorClientCertSpec {
	if in == nil {
		return nil
	}
	out := new(AerospikeOperatorClientCertSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikePersistentVolumePolicySpec) DeepCopyInto(out *AerospikePersistentVolumePolicySpec) {
	*out = *in
	if in.InputInitMethod != nil {
		in, out := &in.InputInitMethod, &out.InputInitMethod
		*out = new(AerospikeVolumeInitMethod)
		**out = **in
	}
	if in.InputCascadeDelete != nil {
		in, out := &in.InputCascadeDelete, &out.InputCascadeDelete
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikePersistentVolumePolicySpec.
func (in *AerospikePersistentVolumePolicySpec) DeepCopy() *AerospikePersistentVolumePolicySpec {
	if in == nil {
		return nil
	}
	out := new(AerospikePersistentVolumePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikePodSpec) DeepCopyInto(out *AerospikePodSpec) {
	*out = *in
	in.AerospikeContainerSpec.DeepCopyInto(&out.AerospikeContainerSpec)
	in.AerospikeObjectMeta.DeepCopyInto(&out.AerospikeObjectMeta)
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.SchedulingPolicy.DeepCopyInto(&out.SchedulingPolicy)
	if in.InputDNSPolicy != nil {
		in, out := &in.InputDNSPolicy, &out.InputDNSPolicy
		*out = new(v1.DNSPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikePodSpec.
func (in *AerospikePodSpec) DeepCopy() *AerospikePodSpec {
	if in == nil {
		return nil
	}
	out := new(AerospikePodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikePodStatus) DeepCopyInto(out *AerospikePodStatus) {
	*out = *in
	in.Aerospike.DeepCopyInto(&out.Aerospike)
	if in.InitializedVolumes != nil {
		in, out := &in.InitializedVolumes, &out.InitializedVolumes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InitializedVolumePaths != nil {
		in, out := &in.InitializedVolumePaths, &out.InitializedVolumePaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikePodStatus.
func (in *AerospikePodStatus) DeepCopy() *AerospikePodStatus {
	if in == nil {
		return nil
	}
	out := new(AerospikePodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeRoleSpec) DeepCopyInto(out *AerospikeRoleSpec) {
	*out = *in
	if in.Privileges != nil {
		in, out := &in.Privileges, &out.Privileges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Whitelist != nil {
		in, out := &in.Whitelist, &out.Whitelist
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeRoleSpec.
func (in *AerospikeRoleSpec) DeepCopy() *AerospikeRoleSpec {
	if in == nil {
		return nil
	}
	out := new(AerospikeRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeSecretCertSource) DeepCopyInto(out *AerospikeSecretCertSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeSecretCertSource.
func (in *AerospikeSecretCertSource) DeepCopy() *AerospikeSecretCertSource {
	if in == nil {
		return nil
	}
	out := new(AerospikeSecretCertSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeServerVolumeAttachment) DeepCopyInto(out *AerospikeServerVolumeAttachment) {
	*out = *in
	in.AttachmentOptions.DeepCopyInto(&out.AttachmentOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeServerVolumeAttachment.
func (in *AerospikeServerVolumeAttachment) DeepCopy() *AerospikeServerVolumeAttachment {
	if in == nil {
		return nil
	}
	out := new(AerospikeServerVolumeAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeStorageSpec) DeepCopyInto(out *AerospikeStorageSpec) {
	*out = *in
	in.FileSystemVolumePolicy.DeepCopyInto(&out.FileSystemVolumePolicy)
	in.BlockVolumePolicy.DeepCopyInto(&out.BlockVolumePolicy)
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]VolumeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeStorageSpec.
func (in *AerospikeStorageSpec) DeepCopy() *AerospikeStorageSpec {
	if in == nil {
		return nil
	}
	out := new(AerospikeStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AerospikeUserSpec) DeepCopyInto(out *AerospikeUserSpec) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AerospikeUserSpec.
func (in *AerospikeUserSpec) DeepCopy() *AerospikeUserSpec {
	if in == nil {
		return nil
	}
	out := new(AerospikeUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachmentOptions) DeepCopyInto(out *AttachmentOptions) {
	*out = *in
	in.MountOptions.DeepCopyInto(&out.MountOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachmentOptions.
func (in *AttachmentOptions) DeepCopy() *AttachmentOptions {
	if in == nil {
		return nil
	}
	out := new(AttachmentOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerSpec) DeepCopyInto(out *LoadBalancerSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LoadBalancerSourceRanges != nil {
		in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerSpec.
func (in *LoadBalancerSpec) DeepCopy() *LoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountOptions) DeepCopyInto(out *MountOptions) {
	*out = *in
	if in.MountPropagation != nil {
		in, out := &in.MountPropagation, &out.MountPropagation
		*out = new(v1.MountPropagationMode)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountOptions.
func (in *MountOptions) DeepCopy() *MountOptions {
	if in == nil {
		return nil
	}
	out := new(MountOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeSpec) DeepCopyInto(out *PersistentVolumeSpec) {
	*out = *in
	in.AerospikeObjectMeta.DeepCopyInto(&out.AerospikeObjectMeta)
	out.Size = in.Size.DeepCopy()
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeSpec.
func (in *PersistentVolumeSpec) DeepCopy() *PersistentVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rack) DeepCopyInto(out *Rack) {
	*out = *in
	if in.InputAerospikeConfig != nil {
		in, out := &in.InputAerospikeConfig, &out.InputAerospikeConfig
		*out = (*in).DeepCopy()
	}
	in.AerospikeConfig.DeepCopyInto(&out.AerospikeConfig)
	if in.InputStorage != nil {
		in, out := &in.InputStorage, &out.InputStorage
		*out = new(AerospikeStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Storage.DeepCopyInto(&out.Storage)
	if in.InputPodSpec != nil {
		in, out := &in.InputPodSpec, &out.InputPodSpec
		*out = new(RackPodSpec)
		(*in).DeepCopyInto(*out)
	}
	in.PodSpec.DeepCopyInto(&out.PodSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rack.
func (in *Rack) DeepCopy() *Rack {
	if in == nil {
		return nil
	}
	out := new(Rack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RackConfig) DeepCopyInto(out *RackConfig) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Racks != nil {
		in, out := &in.Racks, &out.Racks
		*out = make([]Rack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackConfig.
func (in *RackConfig) DeepCopy() *RackConfig {
	if in == nil {
		return nil
	}
	out := new(RackConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RackPodSpec) DeepCopyInto(out *RackPodSpec) {
	*out = *in
	in.SchedulingPolicy.DeepCopyInto(&out.SchedulingPolicy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackPodSpec.
func (in *RackPodSpec) DeepCopy() *RackPodSpec {
	if in == nil {
		return nil
	}
	out := new(RackPodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicy) DeepCopyInto(out *SchedulingPolicy) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicy.
func (in *SchedulingPolicy) DeepCopy() *SchedulingPolicy {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeedsFinderServices) DeepCopyInto(out *SeedsFinderServices) {
	*out = *in
	if in.LoadBalancer != nil {
		in, out := &in.LoadBalancer, &out.LoadBalancer
		*out = new(LoadBalancerSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeedsFinderServices.
func (in *SeedsFinderServices) DeepCopy() *SeedsFinderServices {
	if in == nil {
		return nil
	}
	out := new(SeedsFinderServices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidationPolicySpec) DeepCopyInto(out *ValidationPolicySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidationPolicySpec.
func (in *ValidationPolicySpec) DeepCopy() *ValidationPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ValidationPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeAttachment) DeepCopyInto(out *VolumeAttachment) {
	*out = *in
	in.AttachmentOptions.DeepCopyInto(&out.AttachmentOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeAttachment.
func (in *VolumeAttachment) DeepCopy() *VolumeAttachment {
	if in == nil {
		return nil
	}
	out := new(VolumeAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSource) DeepCopyInto(out *VolumeSource) {
	*out = *in
	if in.EmptyDir != nil {
		in, out := &in.EmptyDir, &out.EmptyDir
		*out = new(v1.EmptyDirVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.SecretVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(v1.ConfigMapVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.HostPath != nil {
		in, out := &in.HostPath, &out.HostPath
		*out = new(v1.HostPathVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.PersistentVolume != nil {
		in, out := &in.PersistentVolume, &out.PersistentVolume
		*out = new(PersistentVolumeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSource.
func (in *VolumeSource) DeepCopy() *VolumeSource {
	if in == nil {
		return nil
	}
	out := new(VolumeSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	in.AerospikePersistentVolumePolicySpec.DeepCopyInto(&out.AerospikePersistentVolumePolicySpec)
	in.Source.DeepCopyInto(&out.Source)
	if in.Aerospike != nil {
		in, out := &in.Aerospike, &out.Aerospike
		*out = new(AerospikeServerVolumeAttachment)
		(*in).DeepCopyInto(*out)
	}
	if in.Sidecars != nil {
		in, out := &in.Sidecars, &out.Sidecars
		*out = make([]VolumeAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]VolumeAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}
