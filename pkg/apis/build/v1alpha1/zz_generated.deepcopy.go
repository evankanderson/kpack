// +build !ignore_autogenerated

/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNBBuild) DeepCopyInto(out *CNBBuild) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNBBuild.
func (in *CNBBuild) DeepCopy() *CNBBuild {
	if in == nil {
		return nil
	}
	out := new(CNBBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CNBBuild) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNBBuildList) DeepCopyInto(out *CNBBuildList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CNBBuild, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNBBuildList.
func (in *CNBBuildList) DeepCopy() *CNBBuildList {
	if in == nil {
		return nil
	}
	out := new(CNBBuildList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CNBBuildList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNBBuildSpec) DeepCopyInto(out *CNBBuildSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNBBuildSpec.
func (in *CNBBuildSpec) DeepCopy() *CNBBuildSpec {
	if in == nil {
		return nil
	}
	out := new(CNBBuildSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNBBuildStatus) DeepCopyInto(out *CNBBuildStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.BuildMetadata != nil {
		in, out := &in.BuildMetadata, &out.BuildMetadata
		*out = make([]CNBBuildpackMetadata, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNBBuildStatus.
func (in *CNBBuildStatus) DeepCopy() *CNBBuildStatus {
	if in == nil {
		return nil
	}
	out := new(CNBBuildStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNBBuildpackMetadata) DeepCopyInto(out *CNBBuildpackMetadata) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNBBuildpackMetadata.
func (in *CNBBuildpackMetadata) DeepCopy() *CNBBuildpackMetadata {
	if in == nil {
		return nil
	}
	out := new(CNBBuildpackMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (im *CNBImage) DeepCopyInto(out *CNBImage) {
	*out = *im
	out.TypeMeta = im.TypeMeta
	im.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = im.Spec
	im.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNBImage.
func (im *CNBImage) DeepCopy() *CNBImage {
	if im == nil {
		return nil
	}
	out := new(CNBImage)
	im.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (im *CNBImage) DeepCopyObject() runtime.Object {
	if c := im.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNBImageList) DeepCopyInto(out *CNBImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CNBImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNBImageList.
func (in *CNBImageList) DeepCopy() *CNBImageList {
	if in == nil {
		return nil
	}
	out := new(CNBImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CNBImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNBImageSpec) DeepCopyInto(out *CNBImageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNBImageSpec.
func (in *CNBImageSpec) DeepCopy() *CNBImageSpec {
	if in == nil {
		return nil
	}
	out := new(CNBImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CNBImageStatus) DeepCopyInto(out *CNBImageStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CNBImageStatus.
func (in *CNBImageStatus) DeepCopy() *CNBImageStatus {
	if in == nil {
		return nil
	}
	out := new(CNBImageStatus)
	in.DeepCopyInto(out)
	return out
}