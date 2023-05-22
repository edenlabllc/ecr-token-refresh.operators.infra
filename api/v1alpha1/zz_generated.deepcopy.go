//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 @apanasiuk-el edenlabllc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECRTokenRefresh) DeepCopyInto(out *ECRTokenRefresh) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECRTokenRefresh.
func (in *ECRTokenRefresh) DeepCopy() *ECRTokenRefresh {
	if in == nil {
		return nil
	}
	out := new(ECRTokenRefresh)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ECRTokenRefresh) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECRTokenRefreshList) DeepCopyInto(out *ECRTokenRefreshList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ECRTokenRefresh, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECRTokenRefreshList.
func (in *ECRTokenRefreshList) DeepCopy() *ECRTokenRefreshList {
	if in == nil {
		return nil
	}
	out := new(ECRTokenRefreshList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ECRTokenRefreshList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECRTokenRefreshSpec) DeepCopyInto(out *ECRTokenRefreshSpec) {
	*out = *in
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(v1.Duration)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECRTokenRefreshSpec.
func (in *ECRTokenRefreshSpec) DeepCopy() *ECRTokenRefreshSpec {
	if in == nil {
		return nil
	}
	out := new(ECRTokenRefreshSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ECRTokenRefreshStatus) DeepCopyInto(out *ECRTokenRefreshStatus) {
	*out = *in
	if in.LastUpdatedTime != nil {
		in, out := &in.LastUpdatedTime, &out.LastUpdatedTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ECRTokenRefreshStatus.
func (in *ECRTokenRefreshStatus) DeepCopy() *ECRTokenRefreshStatus {
	if in == nil {
		return nil
	}
	out := new(ECRTokenRefreshStatus)
	in.DeepCopyInto(out)
	return out
}
