// Copyright (c) 2019 Snowflake Inc. All rights reserved.
//
//Licensed under the Apache License, Version 2.0 (the
//"License"); you may not use this file except in compliance
//with the License.  You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing,
//software distributed under the License is distributed on an
//"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
//KIND, either express or implied.  See the License for the
//specific language governing permissions and limitations
//under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: ansible.proto

package ansible

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Var struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Var) Reset() {
	*x = Var{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ansible_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Var) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Var) ProtoMessage() {}

func (x *Var) ProtoReflect() protoreflect.Message {
	mi := &file_ansible_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Var.ProtoReflect.Descriptor instead.
func (*Var) Descriptor() ([]byte, []int) {
	return file_ansible_proto_rawDescGZIP(), []int{0}
}

func (x *Var) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Var) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type RunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The playbook to execute. Needs to be a fully qualified path.
	Playbook string `protobuf:"bytes,1,opt,name=playbook,proto3" json:"playbook,omitempty"`
	// Will become N -e options to ansible-playbook
	Vars []*Var `protobuf:"bytes,2,rep,name=vars,proto3" json:"vars,omitempty"`
	// The user to use for exection.
	User string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	// If set --check is passed to ansible. Depending on playbook settings
	// this may still execute tasks.
	Check bool `protobuf:"varint,4,opt,name=check,proto3" json:"check,omitempty"`
	// If set --diff is passed to ansible
	Diff bool `protobuf:"varint,5,opt,name=diff,proto3" json:"diff,omitempty"`
	// If true, execute ansible with verbose output enabled (equivilant to -vvv)
	Verbose bool `protobuf:"varint,6,opt,name=verbose,proto3" json:"verbose,omitempty"`
}

func (x *RunRequest) Reset() {
	*x = RunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ansible_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunRequest) ProtoMessage() {}

func (x *RunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ansible_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunRequest.ProtoReflect.Descriptor instead.
func (*RunRequest) Descriptor() ([]byte, []int) {
	return file_ansible_proto_rawDescGZIP(), []int{1}
}

func (x *RunRequest) GetPlaybook() string {
	if x != nil {
		return x.Playbook
	}
	return ""
}

func (x *RunRequest) GetVars() []*Var {
	if x != nil {
		return x.Vars
	}
	return nil
}

func (x *RunRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RunRequest) GetCheck() bool {
	if x != nil {
		return x.Check
	}
	return false
}

func (x *RunRequest) GetDiff() bool {
	if x != nil {
		return x.Diff
	}
	return false
}

func (x *RunRequest) GetVerbose() bool {
	if x != nil {
		return x.Verbose
	}
	return false
}

type RunReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All output sent to stdout.
	Stdout string `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	// All output sent to stderr.
	Stderr string `protobuf:"bytes,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	// The return code from the ansible command. Ansible returning non-zero will
	// not be an RPC failure as exec'ing is all we guarentee and some playbooks
	// are designed to return non-zero.
	ReturnCode int32 `protobuf:"varint,3,opt,name=return_code,json=returnCode,proto3" json:"return_code,omitempty"`
}

func (x *RunReply) Reset() {
	*x = RunReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ansible_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunReply) ProtoMessage() {}

func (x *RunReply) ProtoReflect() protoreflect.Message {
	mi := &file_ansible_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunReply.ProtoReflect.Descriptor instead.
func (*RunReply) Descriptor() ([]byte, []int) {
	return file_ansible_proto_rawDescGZIP(), []int{2}
}

func (x *RunReply) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *RunReply) GetStderr() string {
	if x != nil {
		return x.Stderr
	}
	return ""
}

func (x *RunReply) GetReturnCode() int32 {
	if x != nil {
		return x.ReturnCode
	}
	return 0
}

var File_ansible_proto protoreflect.FileDescriptor

var file_ansible_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x41, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x22, 0x2d, 0x0a, 0x03, 0x56, 0x61, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x6f,
	0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x6f,
	0x6f, 0x6b, 0x12, 0x20, 0x0a, 0x04, 0x76, 0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x41, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x2e, 0x56, 0x61, 0x72, 0x52, 0x04,
	0x76, 0x61, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x69, 0x66, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x69,
	0x66, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x22, 0x5b, 0x0a, 0x08,
	0x52, 0x75, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x6f,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x3b, 0x0a, 0x08, 0x50, 0x6c, 0x61,
	0x79, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x2f, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x13, 0x2e, 0x41,
	0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x41, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x6e, 0x6f, 0x77, 0x66, 0x6c, 0x61, 0x6b, 0x65, 0x2d, 0x4c,
	0x61, 0x62, 0x73, 0x2f, 0x73, 0x61, 0x6e, 0x73, 0x73, 0x68, 0x65, 0x6c, 0x6c, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x61, 0x6e, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ansible_proto_rawDescOnce sync.Once
	file_ansible_proto_rawDescData = file_ansible_proto_rawDesc
)

func file_ansible_proto_rawDescGZIP() []byte {
	file_ansible_proto_rawDescOnce.Do(func() {
		file_ansible_proto_rawDescData = protoimpl.X.CompressGZIP(file_ansible_proto_rawDescData)
	})
	return file_ansible_proto_rawDescData
}

var file_ansible_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ansible_proto_goTypes = []interface{}{
	(*Var)(nil),        // 0: Ansible.Var
	(*RunRequest)(nil), // 1: Ansible.RunRequest
	(*RunReply)(nil),   // 2: Ansible.RunReply
}
var file_ansible_proto_depIdxs = []int32{
	0, // 0: Ansible.RunRequest.vars:type_name -> Ansible.Var
	1, // 1: Ansible.Playbook.Run:input_type -> Ansible.RunRequest
	2, // 2: Ansible.Playbook.Run:output_type -> Ansible.RunReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ansible_proto_init() }
func file_ansible_proto_init() {
	if File_ansible_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ansible_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Var); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ansible_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ansible_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ansible_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ansible_proto_goTypes,
		DependencyIndexes: file_ansible_proto_depIdxs,
		MessageInfos:      file_ansible_proto_msgTypes,
	}.Build()
	File_ansible_proto = out.File
	file_ansible_proto_rawDesc = nil
	file_ansible_proto_goTypes = nil
	file_ansible_proto_depIdxs = nil
}
