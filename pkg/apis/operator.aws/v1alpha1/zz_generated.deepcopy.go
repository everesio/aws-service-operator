// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudFormationTemplate) DeepCopyInto(out *CloudFormationTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Data = in.Data
	out.Status = in.Status
	out.Output = in.Output
	out.AdditionalResources = in.AdditionalResources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudFormationTemplate.
func (in *CloudFormationTemplate) DeepCopy() *CloudFormationTemplate {
	if in == nil {
		return nil
	}
	out := new(CloudFormationTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudFormationTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudFormationTemplateAdditionalResources) DeepCopyInto(out *CloudFormationTemplateAdditionalResources) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudFormationTemplateAdditionalResources.
func (in *CloudFormationTemplateAdditionalResources) DeepCopy() *CloudFormationTemplateAdditionalResources {
	if in == nil {
		return nil
	}
	out := new(CloudFormationTemplateAdditionalResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudFormationTemplateData) DeepCopyInto(out *CloudFormationTemplateData) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudFormationTemplateData.
func (in *CloudFormationTemplateData) DeepCopy() *CloudFormationTemplateData {
	if in == nil {
		return nil
	}
	out := new(CloudFormationTemplateData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudFormationTemplateList) DeepCopyInto(out *CloudFormationTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudFormationTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudFormationTemplateList.
func (in *CloudFormationTemplateList) DeepCopy() *CloudFormationTemplateList {
	if in == nil {
		return nil
	}
	out := new(CloudFormationTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudFormationTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudFormationTemplateOutput) DeepCopyInto(out *CloudFormationTemplateOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudFormationTemplateOutput.
func (in *CloudFormationTemplateOutput) DeepCopy() *CloudFormationTemplateOutput {
	if in == nil {
		return nil
	}
	out := new(CloudFormationTemplateOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudFormationTemplateStatus) DeepCopyInto(out *CloudFormationTemplateStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudFormationTemplateStatus.
func (in *CloudFormationTemplateStatus) DeepCopy() *CloudFormationTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(CloudFormationTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Bucket) DeepCopyInto(out *S3Bucket) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	out.Output = in.Output
	out.AdditionalResources = in.AdditionalResources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Bucket.
func (in *S3Bucket) DeepCopy() *S3Bucket {
	if in == nil {
		return nil
	}
	out := new(S3Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3Bucket) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketAdditionalResources) DeepCopyInto(out *S3BucketAdditionalResources) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketAdditionalResources.
func (in *S3BucketAdditionalResources) DeepCopy() *S3BucketAdditionalResources {
	if in == nil {
		return nil
	}
	out := new(S3BucketAdditionalResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketList) DeepCopyInto(out *S3BucketList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]S3Bucket, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketList.
func (in *S3BucketList) DeepCopy() *S3BucketList {
	if in == nil {
		return nil
	}
	out := new(S3BucketList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *S3BucketList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketLogging) DeepCopyInto(out *S3BucketLogging) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketLogging.
func (in *S3BucketLogging) DeepCopy() *S3BucketLogging {
	if in == nil {
		return nil
	}
	out := new(S3BucketLogging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketOutput) DeepCopyInto(out *S3BucketOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketOutput.
func (in *S3BucketOutput) DeepCopy() *S3BucketOutput {
	if in == nil {
		return nil
	}
	out := new(S3BucketOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketSpec) DeepCopyInto(out *S3BucketSpec) {
	*out = *in
	out.Logging = in.Logging
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketSpec.
func (in *S3BucketSpec) DeepCopy() *S3BucketSpec {
	if in == nil {
		return nil
	}
	out := new(S3BucketSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketStatus) DeepCopyInto(out *S3BucketStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketStatus.
func (in *S3BucketStatus) DeepCopy() *S3BucketStatus {
	if in == nil {
		return nil
	}
	out := new(S3BucketStatus)
	in.DeepCopyInto(out)
	return out
}
