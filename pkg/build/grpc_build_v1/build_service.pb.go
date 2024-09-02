// Copyright 2024 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.27.3
// source: pkg/build/grpc_build_v1/build_service.proto

package grpc_build_v1

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

type BuildKind int32

const (
	BuildKind_BUILD_KIND_UNSPECIFIED                     BuildKind = 0
	BuildKind_BUILD_KIND_APP_BUILD_WITH_SOURCE_UPLOAD    BuildKind = 1 // tsuru app build ... /path/to/my/files.sh
	BuildKind_BUILD_KIND_APP_DEPLOY_WITH_SOURCE_UPLOAD   BuildKind = 1 // tsuru app deploy ... /path/to/my/files.sh
	BuildKind_BUILD_KIND_APP_BUILD_WITH_CONTAINER_IMAGE  BuildKind = 2 // tsuru app build ... -i registry.example.com/tsuru/my-app:staging
	BuildKind_BUILD_KIND_APP_DEPLOY_WITH_CONTAINER_IMAGE BuildKind = 2 // tsuru app deploy ... -i registry.example.com/tsuru/my-app:staging
	BuildKind_BUILD_KIND_APP_BUILD_WITH_CONTAINER_FILE   BuildKind = 3 // tsuru app build ... --dockerfile Dockerfile --dockerfile-context ./
	BuildKind_BUILD_KIND_APP_DEPLOY_WITH_CONTAINER_FILE  BuildKind = 3 // tsuru app deploy ... --dockerfile Dockerfile --dockerfile-context ./
	BuildKind_BUILD_KIND_PLATFORM_WITH_CONTAINER_IMAGE   BuildKind = 5 // tsuru platform add/update ... -i registry.example.com/tsuru/python:latest
	BuildKind_BUILD_KIND_PLATFORM_WITH_CONTAINER_FILE    BuildKind = 6 // tsuru platform add/update ... --dockerfile Dockerfile
	BuildKind_BUILD_KIND_JOB_CREATE_WITH_CONTAINER_IMAGE BuildKind = 7 // tsuru job create ... -i registry.example.com/tsuru/my-job:latest
)

// Enum value maps for BuildKind.
var (
	BuildKind_name = map[int32]string{
		0: "BUILD_KIND_UNSPECIFIED",
		1: "BUILD_KIND_APP_BUILD_WITH_SOURCE_UPLOAD",
		// Duplicate value: 1: "BUILD_KIND_APP_DEPLOY_WITH_SOURCE_UPLOAD",
		2: "BUILD_KIND_APP_BUILD_WITH_CONTAINER_IMAGE",
		// Duplicate value: 2: "BUILD_KIND_APP_DEPLOY_WITH_CONTAINER_IMAGE",
		3: "BUILD_KIND_APP_BUILD_WITH_CONTAINER_FILE",
		// Duplicate value: 3: "BUILD_KIND_APP_DEPLOY_WITH_CONTAINER_FILE",
		5: "BUILD_KIND_PLATFORM_WITH_CONTAINER_IMAGE",
		6: "BUILD_KIND_PLATFORM_WITH_CONTAINER_FILE",
		7: "BUILD_KIND_JOB_CREATE_WITH_CONTAINER_IMAGE",
	}
	BuildKind_value = map[string]int32{
		"BUILD_KIND_UNSPECIFIED":                     0,
		"BUILD_KIND_APP_BUILD_WITH_SOURCE_UPLOAD":    1,
		"BUILD_KIND_APP_DEPLOY_WITH_SOURCE_UPLOAD":   1,
		"BUILD_KIND_APP_BUILD_WITH_CONTAINER_IMAGE":  2,
		"BUILD_KIND_APP_DEPLOY_WITH_CONTAINER_IMAGE": 2,
		"BUILD_KIND_APP_BUILD_WITH_CONTAINER_FILE":   3,
		"BUILD_KIND_APP_DEPLOY_WITH_CONTAINER_FILE":  3,
		"BUILD_KIND_PLATFORM_WITH_CONTAINER_IMAGE":   5,
		"BUILD_KIND_PLATFORM_WITH_CONTAINER_FILE":    6,
		"BUILD_KIND_JOB_CREATE_WITH_CONTAINER_IMAGE": 7,
	}
)

func (x BuildKind) Enum() *BuildKind {
	p := new(BuildKind)
	*p = x
	return p
}

func (x BuildKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BuildKind) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_build_grpc_build_v1_build_service_proto_enumTypes[0].Descriptor()
}

func (BuildKind) Type() protoreflect.EnumType {
	return &file_pkg_build_grpc_build_v1_build_service_proto_enumTypes[0]
}

func (x BuildKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BuildKind.Descriptor instead.
func (BuildKind) EnumDescriptor() ([]byte, []int) {
	return file_pkg_build_grpc_build_v1_build_service_proto_rawDescGZIP(), []int{0}
}

type BuildRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// BuildKind indicates what kind of process started the build on the caller side.
	Kind BuildKind `protobuf:"varint,1,opt,name=kind,proto3,enum=grpc_build_v1.BuildKind" json:"kind,omitempty"`
	// App is the Tsuru app which is being deployed, if any.
	//
	// NOTE: mandatory field when build kind starts with BUILD_KIND_APP_.
	App *TsuruApp `protobuf:"bytes,2,opt,name=app,proto3" json:"app,omitempty"`
	// Platform is the Tsuru platform which is being builded, if any.
	//
	// NOTE: mandatory field when build kind starts with BUILD_KIND_PLATFORM_.
	Platform *TsuruPlatform `protobuf:"bytes,3,opt,name=platform,proto3" json:"platform,omitempty"`
	// SourceImage is the source container image name.
	//
	// When deploy is from app's source code (BUILD_KIND_APP_DEPLOY_WITH_SOURCE_UPLOAD), it holds
	// the plataform's container image (e.g. docker.io/tsuru/scratch:latest).
	// When deploy is from container image (BUILD_KIND_APP_DEPLOY_WITH_CONTAINER_IMAGE), it holds
	// the app's container image (e.g. registry.example.com/company/app:v100).
	// Otherwise it's empty.
	SourceImage string `protobuf:"bytes,4,opt,name=source_image,json=sourceImage,proto3" json:"source_image,omitempty"`
	// DestinationImages are the tags of the container image after build.
	DestinationImages []string `protobuf:"bytes,5,rep,name=destination_images,json=destinationImages,proto3" json:"destination_images,omitempty"`
	// Data is the app's source data (or container context).
	// Cannot exceed 2^32 of size.
	//
	// See more: https://developers.google.com/protocol-buffers/docs/proto3#scalar
	Data []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// Containerfile is the container file definition.
	Containerfile string `protobuf:"bytes,7,opt,name=Containerfile,proto3" json:"Containerfile,omitempty"`
	// PushOptions contains the options push the generated images.
	PushOptions *PushOptions `protobuf:"bytes,10,opt,name=push_options,json=pushOptions,proto3" json:"push_options,omitempty"`
	// Job is the Tsuru job which is being deployed, if any.
	//
	// NOTE: mandatory field when build kind starts with BUILD_KIND_JOB_.
	Job *TsuruJob `protobuf:"bytes,11,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *BuildRequest) Reset() {
	*x = BuildRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildRequest) ProtoMessage() {}

func (x *BuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildRequest.ProtoReflect.Descriptor instead.
func (*BuildRequest) Descriptor() ([]byte, []int) {
	return file_pkg_build_grpc_build_v1_build_service_proto_rawDescGZIP(), []int{0}
}

func (x *BuildRequest) GetKind() BuildKind {
	if x != nil {
		return x.Kind
	}
	return BuildKind_BUILD_KIND_UNSPECIFIED
}

func (x *BuildRequest) GetApp() *TsuruApp {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *BuildRequest) GetPlatform() *TsuruPlatform {
	if x != nil {
		return x.Platform
	}
	return nil
}

func (x *BuildRequest) GetSourceImage() string {
	if x != nil {
		return x.SourceImage
	}
	return ""
}

func (x *BuildRequest) GetDestinationImages() []string {
	if x != nil {
		return x.DestinationImages
	}
	return nil
}

func (x *BuildRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *BuildRequest) GetContainerfile() string {
	if x != nil {
		return x.Containerfile
	}
	return ""
}

func (x *BuildRequest) GetPushOptions() *PushOptions {
	if x != nil {
		return x.PushOptions
	}
	return nil
}

func (x *BuildRequest) GetJob() *TsuruJob {
	if x != nil {
		return x.Job
	}
	return nil
}

type BuildResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*BuildResponse_Output
	//	*BuildResponse_TsuruConfig
	Data isBuildResponse_Data `protobuf_oneof:"data"`
}

func (x *BuildResponse) Reset() {
	*x = BuildResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildResponse) ProtoMessage() {}

func (x *BuildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildResponse.ProtoReflect.Descriptor instead.
func (*BuildResponse) Descriptor() ([]byte, []int) {
	return file_pkg_build_grpc_build_v1_build_service_proto_rawDescGZIP(), []int{1}
}

func (m *BuildResponse) GetData() isBuildResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *BuildResponse) GetOutput() string {
	if x, ok := x.GetData().(*BuildResponse_Output); ok {
		return x.Output
	}
	return ""
}

func (x *BuildResponse) GetTsuruConfig() *TsuruConfig {
	if x, ok := x.GetData().(*BuildResponse_TsuruConfig); ok {
		return x.TsuruConfig
	}
	return nil
}

type isBuildResponse_Data interface {
	isBuildResponse_Data()
}

type BuildResponse_Output struct {
	// Output is the progress messages during the build and push phase.
	Output string `protobuf:"bytes,1,opt,name=output,proto3,oneof"`
}

type BuildResponse_TsuruConfig struct {
	// TsuruConfig is the configuration of the application.
	TsuruConfig *TsuruConfig `protobuf:"bytes,2,opt,name=tsuru_config,json=tsuruConfig,proto3,oneof"`
}

func (*BuildResponse_Output) isBuildResponse_Data() {}

func (*BuildResponse_TsuruConfig) isBuildResponse_Data() {}

type TsuruApp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the Tsuru app name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// EnvVars are the enviroment variables set on app.
	EnvVars map[string]string `protobuf:"bytes,3,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TsuruApp) Reset() {
	*x = TsuruApp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TsuruApp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TsuruApp) ProtoMessage() {}

func (x *TsuruApp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TsuruApp.ProtoReflect.Descriptor instead.
func (*TsuruApp) Descriptor() ([]byte, []int) {
	return file_pkg_build_grpc_build_v1_build_service_proto_rawDescGZIP(), []int{2}
}

func (x *TsuruApp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TsuruApp) GetEnvVars() map[string]string {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

type TsuruJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the Tsuru job name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// EnvVars are the enviroment variables set on the job.
	EnvVars map[string]string `protobuf:"bytes,3,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TsuruJob) Reset() {
	*x = TsuruJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TsuruJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TsuruJob) ProtoMessage() {}

func (x *TsuruJob) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TsuruJob.ProtoReflect.Descriptor instead.
func (*TsuruJob) Descriptor() ([]byte, []int) {
	return file_pkg_build_grpc_build_v1_build_service_proto_rawDescGZIP(), []int{3}
}

func (x *TsuruJob) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TsuruJob) GetEnvVars() map[string]string {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

type TsuruPlatform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the Tsuru platform name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TsuruPlatform) Reset() {
	*x = TsuruPlatform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TsuruPlatform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TsuruPlatform) ProtoMessage() {}

func (x *TsuruPlatform) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TsuruPlatform.ProtoReflect.Descriptor instead.
func (*TsuruPlatform) Descriptor() ([]byte, []int) {
	return file_pkg_build_grpc_build_v1_build_service_proto_rawDescGZIP(), []int{4}
}

func (x *TsuruPlatform) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PushOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Disable turns off the push for container registry.
	Disable bool `protobuf:"varint,1,opt,name=disable,proto3" json:"disable,omitempty"`
	// InsecureRegistry allows sending an image to registry running in plain HTTP.
	InsecureRegistry bool `protobuf:"varint,2,opt,name=insecure_registry,json=insecureRegistry,proto3" json:"insecure_registry,omitempty"`
}

func (x *PushOptions) Reset() {
	*x = PushOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushOptions) ProtoMessage() {}

func (x *PushOptions) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushOptions.ProtoReflect.Descriptor instead.
func (*PushOptions) Descriptor() ([]byte, []int) {
	return file_pkg_build_grpc_build_v1_build_service_proto_rawDescGZIP(), []int{5}
}

func (x *PushOptions) GetDisable() bool {
	if x != nil {
		return x.Disable
	}
	return false
}

func (x *PushOptions) GetInsecureRegistry() bool {
	if x != nil {
		return x.InsecureRegistry
	}
	return false
}

type ContainerImageConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entrypoint   []string `protobuf:"bytes,1,rep,name=entrypoint,proto3" json:"entrypoint,omitempty"`
	Cmd          []string `protobuf:"bytes,2,rep,name=cmd,proto3" json:"cmd,omitempty"`
	ExposedPorts []string `protobuf:"bytes,3,rep,name=exposed_ports,json=exposedPorts,proto3" json:"exposed_ports,omitempty"`
	WorkingDir   string   `protobuf:"bytes,4,opt,name=working_dir,json=workingDir,proto3" json:"working_dir,omitempty"`
}

func (x *ContainerImageConfig) Reset() {
	*x = ContainerImageConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerImageConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerImageConfig) ProtoMessage() {}

func (x *ContainerImageConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerImageConfig.ProtoReflect.Descriptor instead.
func (*ContainerImageConfig) Descriptor() ([]byte, []int) {
	return file_pkg_build_grpc_build_v1_build_service_proto_rawDescGZIP(), []int{6}
}

func (x *ContainerImageConfig) GetEntrypoint() []string {
	if x != nil {
		return x.Entrypoint
	}
	return nil
}

func (x *ContainerImageConfig) GetCmd() []string {
	if x != nil {
		return x.Cmd
	}
	return nil
}

func (x *ContainerImageConfig) GetExposedPorts() []string {
	if x != nil {
		return x.ExposedPorts
	}
	return nil
}

func (x *ContainerImageConfig) GetWorkingDir() string {
	if x != nil {
		return x.WorkingDir
	}
	return ""
}

type TsuruConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Procfile definition found during the build.
	Procfile string `protobuf:"bytes,1,opt,name=procfile,proto3" json:"procfile,omitempty"`
	// TsuruYAML definition found during the build.
	TsuruYaml string `protobuf:"bytes,2,opt,name=tsuru_yaml,json=tsuruYaml,proto3" json:"tsuru_yaml,omitempty"`
	// ContainerImageConfig found in the container image registry.
	ImageConfig *ContainerImageConfig `protobuf:"bytes,3,opt,name=image_config,json=imageConfig,proto3" json:"image_config,omitempty"`
}

func (x *TsuruConfig) Reset() {
	*x = TsuruConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TsuruConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TsuruConfig) ProtoMessage() {}

func (x *TsuruConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TsuruConfig.ProtoReflect.Descriptor instead.
func (*TsuruConfig) Descriptor() ([]byte, []int) {
	return file_pkg_build_grpc_build_v1_build_service_proto_rawDescGZIP(), []int{7}
}

func (x *TsuruConfig) GetProcfile() string {
	if x != nil {
		return x.Procfile
	}
	return ""
}

func (x *TsuruConfig) GetTsuruYaml() string {
	if x != nil {
		return x.TsuruYaml
	}
	return ""
}

func (x *TsuruConfig) GetImageConfig() *ContainerImageConfig {
	if x != nil {
		return x.ImageConfig
	}
	return nil
}

var File_pkg_build_grpc_build_v1_build_service_proto protoreflect.FileDescriptor

var file_pkg_build_grpc_build_v1_build_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x76, 0x31, 0x22, 0x97, 0x03, 0x0a,
	0x0c, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x29, 0x0a, 0x03, 0x61,
	0x70, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x75, 0x72, 0x75, 0x41, 0x70,
	0x70, 0x52, 0x03, 0x61, 0x70, 0x70, 0x12, 0x38, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x75, 0x72, 0x75, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x11, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x3d, 0x0a, 0x0c,
	0x70, 0x75, 0x73, 0x68, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f,
	0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b,
	0x70, 0x75, 0x73, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x29, 0x0a, 0x03, 0x6a,
	0x6f, 0x62, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x75, 0x72, 0x75, 0x4a, 0x6f,
	0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x72, 0x0a, 0x0d, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x3f, 0x0a, 0x0c, 0x74, 0x73, 0x75, 0x72, 0x75, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x5f, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x75, 0x72, 0x75, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x73, 0x75, 0x72, 0x75, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9b, 0x01, 0x0a, 0x08, 0x54,
	0x73, 0x75, 0x72, 0x75, 0x41, 0x70, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x65,
	0x6e, 0x76, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x76, 0x31, 0x2e, 0x54, 0x73,
	0x75, 0x72, 0x75, 0x41, 0x70, 0x70, 0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c,
	0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9b, 0x01, 0x0a, 0x08, 0x54, 0x73, 0x75,
	0x72, 0x75, 0x4a, 0x6f, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x65, 0x6e, 0x76,
	0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x76, 0x31, 0x2e, 0x54, 0x73, 0x75, 0x72,
	0x75, 0x4a, 0x6f, 0x62, 0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x65, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x45, 0x6e,
	0x76, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x23, 0x0a, 0x0d, 0x54, 0x73, 0x75, 0x72, 0x75, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x54, 0x0a, 0x0b, 0x50,
	0x75, 0x73, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x10, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x22, 0x8e, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d,
	0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x65, 0x64, 0x50, 0x6f, 0x72, 0x74,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44,
	0x69, 0x72, 0x22, 0x90, 0x01, 0x0a, 0x0b, 0x54, 0x73, 0x75, 0x72, 0x75, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x63, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x63, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x73, 0x75, 0x72, 0x75, 0x5f, 0x79, 0x61, 0x6d, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x73, 0x75, 0x72, 0x75, 0x59, 0x61, 0x6d, 0x6c, 0x12, 0x46, 0x0a,
	0x0c, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2a, 0xcd, 0x03, 0x0a, 0x09, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4b,
	0x69, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x16, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x4b, 0x49, 0x4e,
	0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x2b, 0x0a, 0x27, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x50,
	0x50, 0x5f, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x53, 0x4f, 0x55,
	0x52, 0x43, 0x45, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x01, 0x12, 0x2c, 0x0a, 0x28,
	0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x44,
	0x45, 0x50, 0x4c, 0x4f, 0x59, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x10, 0x01, 0x12, 0x2d, 0x0a, 0x29, 0x42, 0x55,
	0x49, 0x4c, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x42, 0x55, 0x49,
	0x4c, 0x44, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45,
	0x52, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x2e, 0x0a, 0x2a, 0x42, 0x55, 0x49,
	0x4c, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x44, 0x45, 0x50, 0x4c,
	0x4f, 0x59, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45,
	0x52, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x2c, 0x0a, 0x28, 0x42, 0x55, 0x49,
	0x4c, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x42, 0x55, 0x49, 0x4c,
	0x44, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52,
	0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x2d, 0x0a, 0x29, 0x42, 0x55, 0x49, 0x4c, 0x44,
	0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59,
	0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x5f,
	0x46, 0x49, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x2c, 0x0a, 0x28, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f,
	0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x57, 0x49,
	0x54, 0x48, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x5f, 0x49, 0x4d, 0x41,
	0x47, 0x45, 0x10, 0x05, 0x12, 0x2b, 0x0a, 0x27, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x4b, 0x49,
	0x4e, 0x44, 0x5f, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x57, 0x49, 0x54, 0x48,
	0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x10,
	0x06, 0x12, 0x2e, 0x0a, 0x2a, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f,
	0x4a, 0x4f, 0x42, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f,
	0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x45, 0x52, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10,
	0x07, 0x1a, 0x02, 0x10, 0x01, 0x32, 0x4f, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x46,
	0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x5f, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x5f, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x73, 0x75, 0x72, 0x75, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_build_grpc_build_v1_build_service_proto_rawDescOnce sync.Once
	file_pkg_build_grpc_build_v1_build_service_proto_rawDescData = file_pkg_build_grpc_build_v1_build_service_proto_rawDesc
)

func file_pkg_build_grpc_build_v1_build_service_proto_rawDescGZIP() []byte {
	file_pkg_build_grpc_build_v1_build_service_proto_rawDescOnce.Do(func() {
		file_pkg_build_grpc_build_v1_build_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_build_grpc_build_v1_build_service_proto_rawDescData)
	})
	return file_pkg_build_grpc_build_v1_build_service_proto_rawDescData
}

var file_pkg_build_grpc_build_v1_build_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_build_grpc_build_v1_build_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_build_grpc_build_v1_build_service_proto_goTypes = []interface{}{
	(BuildKind)(0),               // 0: grpc_build_v1.BuildKind
	(*BuildRequest)(nil),         // 1: grpc_build_v1.BuildRequest
	(*BuildResponse)(nil),        // 2: grpc_build_v1.BuildResponse
	(*TsuruApp)(nil),             // 3: grpc_build_v1.TsuruApp
	(*TsuruJob)(nil),             // 4: grpc_build_v1.TsuruJob
	(*TsuruPlatform)(nil),        // 5: grpc_build_v1.TsuruPlatform
	(*PushOptions)(nil),          // 6: grpc_build_v1.PushOptions
	(*ContainerImageConfig)(nil), // 7: grpc_build_v1.ContainerImageConfig
	(*TsuruConfig)(nil),          // 8: grpc_build_v1.TsuruConfig
	nil,                          // 9: grpc_build_v1.TsuruApp.EnvVarsEntry
	nil,                          // 10: grpc_build_v1.TsuruJob.EnvVarsEntry
}
var file_pkg_build_grpc_build_v1_build_service_proto_depIdxs = []int32{
	0,  // 0: grpc_build_v1.BuildRequest.kind:type_name -> grpc_build_v1.BuildKind
	3,  // 1: grpc_build_v1.BuildRequest.app:type_name -> grpc_build_v1.TsuruApp
	5,  // 2: grpc_build_v1.BuildRequest.platform:type_name -> grpc_build_v1.TsuruPlatform
	6,  // 3: grpc_build_v1.BuildRequest.push_options:type_name -> grpc_build_v1.PushOptions
	4,  // 4: grpc_build_v1.BuildRequest.job:type_name -> grpc_build_v1.TsuruJob
	8,  // 5: grpc_build_v1.BuildResponse.tsuru_config:type_name -> grpc_build_v1.TsuruConfig
	9,  // 6: grpc_build_v1.TsuruApp.env_vars:type_name -> grpc_build_v1.TsuruApp.EnvVarsEntry
	10, // 7: grpc_build_v1.TsuruJob.env_vars:type_name -> grpc_build_v1.TsuruJob.EnvVarsEntry
	7,  // 8: grpc_build_v1.TsuruConfig.image_config:type_name -> grpc_build_v1.ContainerImageConfig
	1,  // 9: grpc_build_v1.Build.Build:input_type -> grpc_build_v1.BuildRequest
	2,  // 10: grpc_build_v1.Build.Build:output_type -> grpc_build_v1.BuildResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_pkg_build_grpc_build_v1_build_service_proto_init() }
func file_pkg_build_grpc_build_v1_build_service_proto_init() {
	if File_pkg_build_grpc_build_v1_build_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildRequest); i {
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
		file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildResponse); i {
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
		file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TsuruApp); i {
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
		file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TsuruJob); i {
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
		file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TsuruPlatform); i {
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
		file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushOptions); i {
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
		file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerImageConfig); i {
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
		file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TsuruConfig); i {
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
	file_pkg_build_grpc_build_v1_build_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*BuildResponse_Output)(nil),
		(*BuildResponse_TsuruConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_build_grpc_build_v1_build_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_build_grpc_build_v1_build_service_proto_goTypes,
		DependencyIndexes: file_pkg_build_grpc_build_v1_build_service_proto_depIdxs,
		EnumInfos:         file_pkg_build_grpc_build_v1_build_service_proto_enumTypes,
		MessageInfos:      file_pkg_build_grpc_build_v1_build_service_proto_msgTypes,
	}.Build()
	File_pkg_build_grpc_build_v1_build_service_proto = out.File
	file_pkg_build_grpc_build_v1_build_service_proto_rawDesc = nil
	file_pkg_build_grpc_build_v1_build_service_proto_goTypes = nil
	file_pkg_build_grpc_build_v1_build_service_proto_depIdxs = nil
}
