//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Kruise Authors.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationType) DeepCopyInto(out *AllocationType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationType.
func (in *AllocationType) DeepCopy() *AllocationType {
	if in == nil {
		return nil
	}
	out := new(AllocationType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Entry) DeepCopyInto(out *Entry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Entry.
func (in *Entry) DeepCopy() *Entry {
	if in == nil {
		return nil
	}
	out := new(Entry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDNAT) DeepCopyInto(out *PodDNAT) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDNAT.
func (in *PodDNAT) DeepCopy() *PodDNAT {
	if in == nil {
		return nil
	}
	out := new(PodDNAT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodDNAT) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDNATList) DeepCopyInto(out *PodDNATList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodDNAT, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDNATList.
func (in *PodDNATList) DeepCopy() *PodDNATList {
	if in == nil {
		return nil
	}
	out := new(PodDNATList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodDNATList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDNATSpec) DeepCopyInto(out *PodDNATSpec) {
	*out = *in
	if in.VSwitch != nil {
		in, out := &in.VSwitch, &out.VSwitch
		*out = new(string)
		**out = **in
	}
	if in.ENI != nil {
		in, out := &in.ENI, &out.ENI
		*out = new(string)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	if in.ExternalIP != nil {
		in, out := &in.ExternalIP, &out.ExternalIP
		*out = new(string)
		**out = **in
	}
	if in.ExternalPort != nil {
		in, out := &in.ExternalPort, &out.ExternalPort
		*out = new(string)
		**out = **in
	}
	if in.InternalIP != nil {
		in, out := &in.InternalIP, &out.InternalIP
		*out = new(string)
		**out = **in
	}
	if in.InternalPort != nil {
		in, out := &in.InternalPort, &out.InternalPort
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.TableId != nil {
		in, out := &in.TableId, &out.TableId
		*out = new(string)
		**out = **in
	}
	if in.EntryId != nil {
		in, out := &in.EntryId, &out.EntryId
		*out = new(string)
		**out = **in
	}
	if in.PortMapping != nil {
		in, out := &in.PortMapping, &out.PortMapping
		*out = make([]PortMapping, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDNATSpec.
func (in *PodDNATSpec) DeepCopy() *PodDNATSpec {
	if in == nil {
		return nil
	}
	out := new(PodDNATSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDNATStatus) DeepCopyInto(out *PodDNATStatus) {
	*out = *in
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.Entries != nil {
		in, out := &in.Entries, &out.Entries
		*out = make([]Entry, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDNATStatus.
func (in *PodDNATStatus) DeepCopy() *PodDNATStatus {
	if in == nil {
		return nil
	}
	out := new(PodDNATStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodEIP) DeepCopyInto(out *PodEIP) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodEIP.
func (in *PodEIP) DeepCopy() *PodEIP {
	if in == nil {
		return nil
	}
	out := new(PodEIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodEIP) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodEIPList) DeepCopyInto(out *PodEIPList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodEIP, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodEIPList.
func (in *PodEIPList) DeepCopy() *PodEIPList {
	if in == nil {
		return nil
	}
	out := new(PodEIPList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodEIPList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodEIPSpec) DeepCopyInto(out *PodEIPSpec) {
	*out = *in
	out.AllocationType = in.AllocationType
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodEIPSpec.
func (in *PodEIPSpec) DeepCopy() *PodEIPSpec {
	if in == nil {
		return nil
	}
	out := new(PodEIPSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodEIPStatus) DeepCopyInto(out *PodEIPStatus) {
	*out = *in
	in.PodLastSeen.DeepCopyInto(&out.PodLastSeen)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodEIPStatus.
func (in *PodEIPStatus) DeepCopy() *PodEIPStatus {
	if in == nil {
		return nil
	}
	out := new(PodEIPStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortMapping) DeepCopyInto(out *PortMapping) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortMapping.
func (in *PortMapping) DeepCopy() *PortMapping {
	if in == nil {
		return nil
	}
	out := new(PortMapping)
	in.DeepCopyInto(out)
	return out
}
