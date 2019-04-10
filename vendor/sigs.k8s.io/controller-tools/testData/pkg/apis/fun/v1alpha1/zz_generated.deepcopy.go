// +build !ignore_autogenerated

/*

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Toy) DeepCopyInto(out *Toy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Toy.
func (in *Toy) DeepCopy() *Toy {
	if in == nil {
		return nil
	}
	out := new(Toy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Toy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToyList) DeepCopyInto(out *ToyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Toy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToyList.
func (in *ToyList) DeepCopy() *ToyList {
	if in == nil {
		return nil
	}
	out := new(ToyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ToyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToySpec) DeepCopyInto(out *ToySpec) {
	*out = *in
	if in.Knights != nil {
		in, out := &in.Knights, &out.Knights
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	in.Template.DeepCopyInto(&out.Template)
	in.Claim.DeepCopyInto(&out.Claim)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToySpec.
func (in *ToySpec) DeepCopy() *ToySpec {
	if in == nil {
		return nil
	}
	out := new(ToySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToyStatus) DeepCopyInto(out *ToyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToyStatus.
func (in *ToyStatus) DeepCopy() *ToyStatus {
	if in == nil {
		return nil
	}
	out := new(ToyStatus)
	in.DeepCopyInto(out)
	return out
}