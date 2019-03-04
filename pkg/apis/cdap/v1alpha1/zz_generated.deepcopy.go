// +build !ignore_autogenerated

/*
Copyright 2019 The CDAP Authors.

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
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDAPMaster) DeepCopyInto(out *CDAPMaster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDAPMaster.
func (in *CDAPMaster) DeepCopy() *CDAPMaster {
	if in == nil {
		return nil
	}
	out := new(CDAPMaster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CDAPMaster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDAPMasterList) DeepCopyInto(out *CDAPMasterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CDAPMaster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDAPMasterList.
func (in *CDAPMasterList) DeepCopy() *CDAPMasterList {
	if in == nil {
		return nil
	}
	out := new(CDAPMasterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CDAPMasterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDAPMasterServiceSpec) DeepCopyInto(out *CDAPMasterServiceSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServicePort != nil {
		in, out := &in.ServicePort, &out.ServicePort
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDAPMasterServiceSpec.
func (in *CDAPMasterServiceSpec) DeepCopy() *CDAPMasterServiceSpec {
	if in == nil {
		return nil
	}
	out := new(CDAPMasterServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDAPMasterSpec) DeepCopyInto(out *CDAPMasterSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LogLevels != nil {
		in, out := &in.LogLevels, &out.LogLevels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.AppFabric.DeepCopyInto(&out.AppFabric)
	in.Logs.DeepCopyInto(&out.Logs)
	in.Messaging.DeepCopyInto(&out.Messaging)
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Metrics.DeepCopyInto(&out.Metrics)
	in.Preview.DeepCopyInto(&out.Preview)
	in.Router.DeepCopyInto(&out.Router)
	in.UserInterface.DeepCopyInto(&out.UserInterface)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDAPMasterSpec.
func (in *CDAPMasterSpec) DeepCopy() *CDAPMasterSpec {
	if in == nil {
		return nil
	}
	out := new(CDAPMasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CDAPMasterStatus) DeepCopyInto(out *CDAPMasterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CDAPMasterStatus.
func (in *CDAPMasterStatus) DeepCopy() *CDAPMasterStatus {
	if in == nil {
		return nil
	}
	out := new(CDAPMasterStatus)
	in.DeepCopyInto(out)
	return out
}
