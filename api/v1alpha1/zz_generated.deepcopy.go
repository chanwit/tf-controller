//go:build !ignore_autogenerated

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

package v1alpha1

import (
	"github.com/fluxcd/pkg/apis/meta"
	corev1 "k8s.io/api/core/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendConfigSpec) DeepCopyInto(out *BackendConfigSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendConfigSpec.
func (in *BackendConfigSpec) DeepCopy() *BackendConfigSpec {
	if in == nil {
		return nil
	}
	out := new(BackendConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendConfigsReference) DeepCopyInto(out *BackendConfigsReference) {
	*out = *in
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendConfigsReference.
func (in *BackendConfigsReference) DeepCopy() *BackendConfigsReference {
	if in == nil {
		return nil
	}
	out := new(BackendConfigsReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudSpec) DeepCopyInto(out *CloudSpec) {
	*out = *in
	if in.Workspaces != nil {
		in, out := &in.Workspaces, &out.Workspaces
		*out = new(CloudWorkspacesSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudSpec.
func (in *CloudSpec) DeepCopy() *CloudSpec {
	if in == nil {
		return nil
	}
	out := new(CloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudWorkspacesSpec) DeepCopyInto(out *CloudWorkspacesSpec) {
	*out = *in
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudWorkspacesSpec.
func (in *CloudWorkspacesSpec) DeepCopy() *CloudWorkspacesSpec {
	if in == nil {
		return nil
	}
	out := new(CloudWorkspacesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossNamespaceSourceReference) DeepCopyInto(out *CrossNamespaceSourceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossNamespaceSourceReference.
func (in *CrossNamespaceSourceReference) DeepCopy() *CrossNamespaceSourceReference {
	if in == nil {
		return nil
	}
	out := new(CrossNamespaceSourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileMapping) DeepCopyInto(out *FileMapping) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileMapping.
func (in *FileMapping) DeepCopy() *FileMapping {
	if in == nil {
		return nil
	}
	out := new(FileMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheck) DeepCopyInto(out *HealthCheck) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheck.
func (in *HealthCheck) DeepCopy() *HealthCheck {
	if in == nil {
		return nil
	}
	out := new(HealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LockStatus) DeepCopyInto(out *LockStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LockStatus.
func (in *LockStatus) DeepCopy() *LockStatus {
	if in == nil {
		return nil
	}
	out := new(LockStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanStatus) DeepCopyInto(out *PlanStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanStatus.
func (in *PlanStatus) DeepCopy() *PlanStatus {
	if in == nil {
		return nil
	}
	out := new(PlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadInputsFromSecretSpec) DeepCopyInto(out *ReadInputsFromSecretSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadInputsFromSecretSpec.
func (in *ReadInputsFromSecretSpec) DeepCopy() *ReadInputsFromSecretSpec {
	if in == nil {
		return nil
	}
	out := new(ReadInputsFromSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInventory) DeepCopyInto(out *ResourceInventory) {
	*out = *in
	if in.Entries != nil {
		in, out := &in.Entries, &out.Entries
		*out = make([]ResourceRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInventory.
func (in *ResourceInventory) DeepCopy() *ResourceInventory {
	if in == nil {
		return nil
	}
	out := new(ResourceInventory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRef) DeepCopyInto(out *ResourceRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRef.
func (in *ResourceRef) DeepCopy() *ResourceRef {
	if in == nil {
		return nil
	}
	out := new(ResourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerPodMetadata) DeepCopyInto(out *RunnerPodMetadata) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerPodMetadata.
func (in *RunnerPodMetadata) DeepCopy() *RunnerPodMetadata {
	if in == nil {
		return nil
	}
	out := new(RunnerPodMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerPodSpec) DeepCopyInto(out *RunnerPodSpec) {
	*out = *in
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]corev1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
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
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerPodSpec.
func (in *RunnerPodSpec) DeepCopy() *RunnerPodSpec {
	if in == nil {
		return nil
	}
	out := new(RunnerPodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerPodTemplate) DeepCopyInto(out *RunnerPodTemplate) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerPodTemplate.
func (in *RunnerPodTemplate) DeepCopy() *RunnerPodTemplate {
	if in == nil {
		return nil
	}
	out := new(RunnerPodTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TFStateSpec) DeepCopyInto(out *TFStateSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TFStateSpec.
func (in *TFStateSpec) DeepCopy() *TFStateSpec {
	if in == nil {
		return nil
	}
	out := new(TFStateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Terraform) DeepCopyInto(out *Terraform) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Terraform.
func (in *Terraform) DeepCopy() *Terraform {
	if in == nil {
		return nil
	}
	out := new(Terraform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Terraform) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformList) DeepCopyInto(out *TerraformList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Terraform, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformList.
func (in *TerraformList) DeepCopy() *TerraformList {
	if in == nil {
		return nil
	}
	out := new(TerraformList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TerraformList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformSpec) DeepCopyInto(out *TerraformSpec) {
	*out = *in
	if in.BackendConfig != nil {
		in, out := &in.BackendConfig, &out.BackendConfig
		*out = new(BackendConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.BackendConfigsFrom != nil {
		in, out := &in.BackendConfigsFrom, &out.BackendConfigsFrom
		*out = make([]BackendConfigsReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Cloud != nil {
		in, out := &in.Cloud, &out.Cloud
		*out = new(CloudSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Vars != nil {
		in, out := &in.Vars, &out.Vars
		*out = make([]Variable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VarsFrom != nil {
		in, out := &in.VarsFrom, &out.VarsFrom
		*out = make([]VarsReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.FileMappings != nil {
		in, out := &in.FileMappings, &out.FileMappings
		*out = make([]FileMapping, len(*in))
		copy(*out, *in)
	}
	out.Interval = in.Interval
	if in.RetryInterval != nil {
		in, out := &in.RetryInterval, &out.RetryInterval
		*out = new(v1.Duration)
		**out = **in
	}
	out.SourceRef = in.SourceRef
	if in.ReadInputsFromSecrets != nil {
		in, out := &in.ReadInputsFromSecrets, &out.ReadInputsFromSecrets
		*out = make([]ReadInputsFromSecretSpec, len(*in))
		copy(*out, *in)
	}
	if in.WriteOutputsToSecret != nil {
		in, out := &in.WriteOutputsToSecret, &out.WriteOutputsToSecret
		*out = new(WriteOutputsToSecretSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CliConfigSecretRef != nil {
		in, out := &in.CliConfigSecretRef, &out.CliConfigSecretRef
		*out = new(corev1.SecretReference)
		**out = **in
	}
	if in.HealthChecks != nil {
		in, out := &in.HealthChecks, &out.HealthChecks
		*out = make([]HealthCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AlwaysCleanupRunnerPod != nil {
		in, out := &in.AlwaysCleanupRunnerPod, &out.AlwaysCleanupRunnerPod
		*out = new(bool)
		**out = **in
	}
	if in.RunnerTerminationGracePeriodSeconds != nil {
		in, out := &in.RunnerTerminationGracePeriodSeconds, &out.RunnerTerminationGracePeriodSeconds
		*out = new(int64)
		**out = **in
	}
	in.RunnerPodTemplate.DeepCopyInto(&out.RunnerPodTemplate)
	if in.TFState != nil {
		in, out := &in.TFState, &out.TFState
		*out = new(TFStateSpec)
		**out = **in
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]Webhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DependsOn != nil {
		in, out := &in.DependsOn, &out.DependsOn
		*out = make([]meta.NamespacedObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Enterprise != nil {
		in, out := &in.Enterprise, &out.Enterprise
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformSpec.
func (in *TerraformSpec) DeepCopy() *TerraformSpec {
	if in == nil {
		return nil
	}
	out := new(TerraformSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerraformStatus) DeepCopyInto(out *TerraformStatus) {
	*out = *in
	out.ReconcileRequestStatus = in.ReconcileRequestStatus
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LastDriftDetectedAt != nil {
		in, out := &in.LastDriftDetectedAt, &out.LastDriftDetectedAt
		*out = (*in).DeepCopy()
	}
	if in.LastAppliedByDriftDetectionAt != nil {
		in, out := &in.LastAppliedByDriftDetectionAt, &out.LastAppliedByDriftDetectionAt
		*out = (*in).DeepCopy()
	}
	if in.AvailableOutputs != nil {
		in, out := &in.AvailableOutputs, &out.AvailableOutputs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Plan = in.Plan
	if in.Inventory != nil {
		in, out := &in.Inventory, &out.Inventory
		*out = new(ResourceInventory)
		(*in).DeepCopyInto(*out)
	}
	out.Lock = in.Lock
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerraformStatus.
func (in *TerraformStatus) DeepCopy() *TerraformStatus {
	if in == nil {
		return nil
	}
	out := new(TerraformStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Variable) DeepCopyInto(out *Variable) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(corev1.EnvVarSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Variable.
func (in *Variable) DeepCopy() *Variable {
	if in == nil {
		return nil
	}
	out := new(Variable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarsReference) DeepCopyInto(out *VarsReference) {
	*out = *in
	if in.VarsKeys != nil {
		in, out := &in.VarsKeys, &out.VarsKeys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarsReference.
func (in *VarsReference) DeepCopy() *VarsReference {
	if in == nil {
		return nil
	}
	out := new(VarsReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webhook) DeepCopyInto(out *Webhook) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webhook.
func (in *Webhook) DeepCopy() *Webhook {
	if in == nil {
		return nil
	}
	out := new(Webhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WriteOutputsToSecretSpec) DeepCopyInto(out *WriteOutputsToSecretSpec) {
	*out = *in
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WriteOutputsToSecretSpec.
func (in *WriteOutputsToSecretSpec) DeepCopy() *WriteOutputsToSecretSpec {
	if in == nil {
		return nil
	}
	out := new(WriteOutputsToSecretSpec)
	in.DeepCopyInto(out)
	return out
}
