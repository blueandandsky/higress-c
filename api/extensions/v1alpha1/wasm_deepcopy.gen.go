// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: extensions/v1alpha1/wasm.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using WasmPlugin within kubernetes types, where deepcopy-gen is used.
func (in *WasmPlugin) DeepCopyInto(out *WasmPlugin) {
	p := proto.Clone(in).(*WasmPlugin)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasmPlugin. Required by controller-gen.
func (in *WasmPlugin) DeepCopy() *WasmPlugin {
	if in == nil {
		return nil
	}
	out := new(WasmPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new WasmPlugin. Required by controller-gen.
func (in *WasmPlugin) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using MatchRule within kubernetes types, where deepcopy-gen is used.
func (in *MatchRule) DeepCopyInto(out *MatchRule) {
	p := proto.Clone(in).(*MatchRule)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchRule. Required by controller-gen.
func (in *MatchRule) DeepCopy() *MatchRule {
	if in == nil {
		return nil
	}
	out := new(MatchRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new MatchRule. Required by controller-gen.
func (in *MatchRule) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
