// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/versionpb/version.proto

package versionpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Version struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string              `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Entrypoint           *Version_Entrypoint `protobuf:"bytes,4,opt,name=entrypoint,proto3" json:"entrypoint,omitempty"`
	Config               []*Version_Config   `protobuf:"bytes,5,rep,name=config,proto3" json:"config,omitempty"`
	Workflows            []*Version_Workflow `protobuf:"bytes,6,rep,name=workflows,proto3" json:"workflows,omitempty"`
	MongoUri             string              `protobuf:"bytes,7,opt,name=mongo_uri,json=mongoUri,proto3" json:"mongo_uri,omitempty"`
	MongoDbName          string              `protobuf:"bytes,8,opt,name=mongo_db_name,json=mongoDbName,proto3" json:"mongo_db_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b189bce7aa4304, []int{0}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Version) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Version) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Version) GetEntrypoint() *Version_Entrypoint {
	if m != nil {
		return m.Entrypoint
	}
	return nil
}

func (m *Version) GetConfig() []*Version_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Version) GetWorkflows() []*Version_Workflow {
	if m != nil {
		return m.Workflows
	}
	return nil
}

func (m *Version) GetMongoUri() string {
	if m != nil {
		return m.MongoUri
	}
	return ""
}

func (m *Version) GetMongoDbName() string {
	if m != nil {
		return m.MongoDbName
	}
	return ""
}

type Version_Config struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version_Config) Reset()         { *m = Version_Config{} }
func (m *Version_Config) String() string { return proto.CompactTextString(m) }
func (*Version_Config) ProtoMessage()    {}
func (*Version_Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b189bce7aa4304, []int{0, 0}
}

func (m *Version_Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version_Config.Unmarshal(m, b)
}
func (m *Version_Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version_Config.Marshal(b, m, deterministic)
}
func (m *Version_Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version_Config.Merge(m, src)
}
func (m *Version_Config) XXX_Size() int {
	return xxx_messageInfo_Version_Config.Size(m)
}
func (m *Version_Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Version_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Version_Config proto.InternalMessageInfo

func (m *Version_Config) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Version_Config) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Version_Entrypoint struct {
	ProtoFile            string   `protobuf:"bytes,1,opt,name=proto_file,json=protoFile,proto3" json:"proto_file,omitempty"`
	Image                string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Src                  string   `protobuf:"bytes,3,opt,name=src,proto3" json:"src,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version_Entrypoint) Reset()         { *m = Version_Entrypoint{} }
func (m *Version_Entrypoint) String() string { return proto.CompactTextString(m) }
func (*Version_Entrypoint) ProtoMessage()    {}
func (*Version_Entrypoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b189bce7aa4304, []int{0, 1}
}

func (m *Version_Entrypoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version_Entrypoint.Unmarshal(m, b)
}
func (m *Version_Entrypoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version_Entrypoint.Marshal(b, m, deterministic)
}
func (m *Version_Entrypoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version_Entrypoint.Merge(m, src)
}
func (m *Version_Entrypoint) XXX_Size() int {
	return xxx_messageInfo_Version_Entrypoint.Size(m)
}
func (m *Version_Entrypoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Version_Entrypoint.DiscardUnknown(m)
}

var xxx_messageInfo_Version_Entrypoint proto.InternalMessageInfo

func (m *Version_Entrypoint) GetProtoFile() string {
	if m != nil {
		return m.ProtoFile
	}
	return ""
}

func (m *Version_Entrypoint) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Version_Entrypoint) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

type Version_Workflow struct {
	Id                   string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Entrypoint           string                   `protobuf:"bytes,3,opt,name=entrypoint,proto3" json:"entrypoint,omitempty"`
	Nodes                []*Version_Workflow_Node `protobuf:"bytes,4,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Edges                []*Version_Workflow_Edge `protobuf:"bytes,5,rep,name=edges,proto3" json:"edges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Version_Workflow) Reset()         { *m = Version_Workflow{} }
func (m *Version_Workflow) String() string { return proto.CompactTextString(m) }
func (*Version_Workflow) ProtoMessage()    {}
func (*Version_Workflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b189bce7aa4304, []int{0, 2}
}

func (m *Version_Workflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version_Workflow.Unmarshal(m, b)
}
func (m *Version_Workflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version_Workflow.Marshal(b, m, deterministic)
}
func (m *Version_Workflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version_Workflow.Merge(m, src)
}
func (m *Version_Workflow) XXX_Size() int {
	return xxx_messageInfo_Version_Workflow.Size(m)
}
func (m *Version_Workflow) XXX_DiscardUnknown() {
	xxx_messageInfo_Version_Workflow.DiscardUnknown(m)
}

var xxx_messageInfo_Version_Workflow proto.InternalMessageInfo

func (m *Version_Workflow) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Version_Workflow) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Version_Workflow) GetEntrypoint() string {
	if m != nil {
		return m.Entrypoint
	}
	return ""
}

func (m *Version_Workflow) GetNodes() []*Version_Workflow_Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *Version_Workflow) GetEdges() []*Version_Workflow_Edge {
	if m != nil {
		return m.Edges
	}
	return nil
}

type Version_Workflow_Node struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Image                string   `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Src                  string   `protobuf:"bytes,4,opt,name=src,proto3" json:"src,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version_Workflow_Node) Reset()         { *m = Version_Workflow_Node{} }
func (m *Version_Workflow_Node) String() string { return proto.CompactTextString(m) }
func (*Version_Workflow_Node) ProtoMessage()    {}
func (*Version_Workflow_Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b189bce7aa4304, []int{0, 2, 0}
}

func (m *Version_Workflow_Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version_Workflow_Node.Unmarshal(m, b)
}
func (m *Version_Workflow_Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version_Workflow_Node.Marshal(b, m, deterministic)
}
func (m *Version_Workflow_Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version_Workflow_Node.Merge(m, src)
}
func (m *Version_Workflow_Node) XXX_Size() int {
	return xxx_messageInfo_Version_Workflow_Node.Size(m)
}
func (m *Version_Workflow_Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Version_Workflow_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Version_Workflow_Node proto.InternalMessageInfo

func (m *Version_Workflow_Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Version_Workflow_Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Version_Workflow_Node) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Version_Workflow_Node) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

type Version_Workflow_Edge struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FromNode             string   `protobuf:"bytes,2,opt,name=fromNode,proto3" json:"fromNode,omitempty"`
	ToNode               string   `protobuf:"bytes,3,opt,name=toNode,proto3" json:"toNode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version_Workflow_Edge) Reset()         { *m = Version_Workflow_Edge{} }
func (m *Version_Workflow_Edge) String() string { return proto.CompactTextString(m) }
func (*Version_Workflow_Edge) ProtoMessage()    {}
func (*Version_Workflow_Edge) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b189bce7aa4304, []int{0, 2, 1}
}

func (m *Version_Workflow_Edge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version_Workflow_Edge.Unmarshal(m, b)
}
func (m *Version_Workflow_Edge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version_Workflow_Edge.Marshal(b, m, deterministic)
}
func (m *Version_Workflow_Edge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version_Workflow_Edge.Merge(m, src)
}
func (m *Version_Workflow_Edge) XXX_Size() int {
	return xxx_messageInfo_Version_Workflow_Edge.Size(m)
}
func (m *Version_Workflow_Edge) XXX_DiscardUnknown() {
	xxx_messageInfo_Version_Workflow_Edge.DiscardUnknown(m)
}

var xxx_messageInfo_Version_Workflow_Edge proto.InternalMessageInfo

func (m *Version_Workflow_Edge) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Version_Workflow_Edge) GetFromNode() string {
	if m != nil {
		return m.FromNode
	}
	return ""
}

func (m *Version_Workflow_Edge) GetToNode() string {
	if m != nil {
		return m.ToNode
	}
	return ""
}

type Request struct {
	Version              *Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b189bce7aa4304, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type Response struct {
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_60b189bce7aa4304, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Version)(nil), "version.Version")
	proto.RegisterType((*Version_Config)(nil), "version.Version.Config")
	proto.RegisterType((*Version_Entrypoint)(nil), "version.Version.Entrypoint")
	proto.RegisterType((*Version_Workflow)(nil), "version.Version.Workflow")
	proto.RegisterType((*Version_Workflow_Node)(nil), "version.Version.Workflow.Node")
	proto.RegisterType((*Version_Workflow_Edge)(nil), "version.Version.Workflow.Edge")
	proto.RegisterType((*Request)(nil), "version.Request")
	proto.RegisterType((*Response)(nil), "version.Response")
}

func init() {
	proto.RegisterFile("proto/versionpb/version.proto", fileDescriptor_60b189bce7aa4304)
}

var fileDescriptor_60b189bce7aa4304 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x55, 0x12, 0xc7, 0x8e, 0xc7, 0x50, 0x95, 0x15, 0x02, 0xe3, 0xd2, 0x2a, 0x8a, 0x38, 0x44,
	0x80, 0xdc, 0x2a, 0x05, 0x71, 0xe0, 0x06, 0x94, 0x03, 0x87, 0x0a, 0x6d, 0x14, 0x90, 0xb8, 0x44,
	0x76, 0x3c, 0x31, 0xab, 0xda, 0x5e, 0xb3, 0xeb, 0xa4, 0xea, 0x95, 0x9f, 0xe0, 0xdf, 0xf8, 0x1a,
	0xe4, 0xf5, 0xda, 0x0e, 0x04, 0x90, 0x7b, 0xca, 0xcc, 0xec, 0x7b, 0x6f, 0x76, 0xde, 0x64, 0x0d,
	0xc7, 0xb9, 0xe0, 0x05, 0x3f, 0xdd, 0xa2, 0x90, 0x8c, 0x67, 0x79, 0x58, 0x47, 0xbe, 0xaa, 0x13,
	0x4b, 0xa7, 0x93, 0x1f, 0x26, 0x58, 0x9f, 0xaa, 0x98, 0x1c, 0x40, 0x9f, 0x45, 0x6e, 0x6f, 0xdc,
	0x9b, 0xda, 0xb4, 0xcf, 0x22, 0x42, 0xc0, 0xc8, 0x82, 0x14, 0xdd, 0xbe, 0xaa, 0xa8, 0x98, 0x3c,
	0x06, 0xbb, 0xfc, 0x95, 0x79, 0xb0, 0x42, 0x77, 0xa0, 0x0e, 0xda, 0x02, 0x79, 0x0d, 0x80, 0x59,
	0x21, 0x6e, 0x72, 0xce, 0xb2, 0xc2, 0x35, 0xc6, 0xbd, 0xa9, 0x33, 0x3b, 0xf2, 0xeb, 0xd6, 0xba,
	0x8f, 0x7f, 0xd1, 0x40, 0xe8, 0x0e, 0x9c, 0x9c, 0x82, 0xb9, 0xe2, 0xd9, 0x9a, 0xc5, 0xee, 0x70,
	0x3c, 0x98, 0x3a, 0xb3, 0x87, 0x7b, 0xc4, 0xb7, 0xea, 0x98, 0x6a, 0x18, 0x79, 0x05, 0xf6, 0x35,
	0x17, 0x57, 0xeb, 0x84, 0x5f, 0x4b, 0xd7, 0x54, 0x9c, 0x47, 0x7b, 0x9c, 0xcf, 0x1a, 0x41, 0x5b,
	0x2c, 0x39, 0x02, 0x3b, 0xe5, 0x59, 0xcc, 0x97, 0x1b, 0xc1, 0x5c, 0x4b, 0x0d, 0x31, 0x52, 0x85,
	0x85, 0x60, 0x64, 0x02, 0x77, 0xab, 0xc3, 0x28, 0x5c, 0xaa, 0xf1, 0x47, 0x0a, 0xe0, 0xa8, 0xe2,
	0xbb, 0xf0, 0x32, 0x48, 0xd1, 0x3b, 0x03, 0xb3, 0xba, 0x0b, 0x39, 0x84, 0xc1, 0x15, 0xde, 0x68,
	0xd3, 0xca, 0x90, 0xdc, 0x87, 0xe1, 0x36, 0x48, 0x36, 0xb5, 0x6d, 0x55, 0xe2, 0xcd, 0x01, 0xda,
	0xb1, 0xc9, 0x31, 0x80, 0xda, 0xc3, 0x72, 0xcd, 0x12, 0xd4, 0x64, 0x5b, 0x55, 0xde, 0xb3, 0x04,
	0x4b, 0x09, 0x96, 0x06, 0x71, 0x23, 0xa1, 0x92, 0xb2, 0x95, 0x14, 0x2b, 0x6d, 0x7a, 0x19, 0x7a,
	0x3f, 0xfb, 0x30, 0xaa, 0xe7, 0xeb, 0xb4, 0xbd, 0x93, 0xdf, 0xf6, 0x53, 0x29, 0xed, 0xae, 0xe0,
	0x05, 0x0c, 0x33, 0x1e, 0xa1, 0x74, 0x0d, 0xe5, 0xe6, 0xc9, 0x3f, 0xdd, 0xf4, 0x2f, 0x79, 0x84,
	0xb4, 0x02, 0x97, 0x2c, 0x8c, 0x62, 0x94, 0x7a, 0x6f, 0xff, 0x61, 0x5d, 0x44, 0x31, 0xd2, 0x0a,
	0xec, 0x51, 0x30, 0x4a, 0x91, 0x4e, 0xf7, 0x6e, 0x0c, 0x19, 0xfc, 0xc5, 0x10, 0xa3, 0x35, 0xe4,
	0x03, 0x18, 0x65, 0x8b, 0x3d, 0x4d, 0x0f, 0x46, 0x6b, 0xc1, 0xd3, 0xb2, 0x9f, 0xd6, 0x6d, 0x72,
	0xf2, 0x00, 0xcc, 0x82, 0xab, 0x93, 0x4a, 0x5c, 0x67, 0x93, 0x97, 0x60, 0x51, 0xfc, 0xb6, 0x41,
	0x59, 0x90, 0xa7, 0x50, 0xbf, 0x17, 0xa5, 0xe9, 0xcc, 0x0e, 0xff, 0x1c, 0x91, 0x36, 0x0f, 0xea,
	0x09, 0x8c, 0x28, 0xca, 0x9c, 0x67, 0x12, 0x89, 0x0b, 0x56, 0x8a, 0x52, 0xb6, 0x9b, 0xac, 0xd3,
	0xd9, 0xf7, 0x3e, 0x1c, 0x68, 0xea, 0x1c, 0xc5, 0x96, 0xad, 0x90, 0x3c, 0x87, 0xe1, 0xbc, 0x08,
	0x44, 0x41, 0x5a, 0x71, 0xdd, 0xdf, 0xbb, 0xb7, 0x53, 0xd1, 0xd2, 0xcf, 0xc0, 0x98, 0x17, 0x3c,
	0xef, 0x06, 0xf6, 0xc1, 0xfa, 0xb8, 0x09, 0x13, 0x26, 0xbf, 0x76, 0xc3, 0x9f, 0x81, 0xbd, 0xc8,
	0xf2, 0xdb, 0x30, 0xce, 0xe1, 0xce, 0x22, 0x8f, 0x82, 0x02, 0xeb, 0x67, 0xd1, 0x81, 0xf4, 0xc6,
	0xf9, 0x62, 0x37, 0xdf, 0xa7, 0xd0, 0x54, 0x7f, 0xff, 0xf3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xc2, 0xee, 0xa1, 0x0b, 0xb9, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VersionServiceClient is the client API for VersionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VersionServiceClient interface {
	Start(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Stop(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Publish(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Unpublish(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	UpdateConfig(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type versionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVersionServiceClient(cc grpc.ClientConnInterface) VersionServiceClient {
	return &versionServiceClient{cc}
}

func (c *versionServiceClient) Start(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/version.VersionService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionServiceClient) Stop(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/version.VersionService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionServiceClient) Publish(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/version.VersionService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionServiceClient) Unpublish(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/version.VersionService/Unpublish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *versionServiceClient) UpdateConfig(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/version.VersionService/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionServiceServer is the server API for VersionService service.
type VersionServiceServer interface {
	Start(context.Context, *Request) (*Response, error)
	Stop(context.Context, *Request) (*Response, error)
	Publish(context.Context, *Request) (*Response, error)
	Unpublish(context.Context, *Request) (*Response, error)
	UpdateConfig(context.Context, *Request) (*Response, error)
}

// UnimplementedVersionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVersionServiceServer struct {
}

func (*UnimplementedVersionServiceServer) Start(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedVersionServiceServer) Stop(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (*UnimplementedVersionServiceServer) Publish(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedVersionServiceServer) Unpublish(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unpublish not implemented")
}
func (*UnimplementedVersionServiceServer) UpdateConfig(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}

func RegisterVersionServiceServer(s *grpc.Server, srv VersionServiceServer) {
	s.RegisterService(&_VersionService_serviceDesc, srv)
}

func _VersionService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.VersionService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServiceServer).Start(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _VersionService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.VersionService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServiceServer).Stop(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _VersionService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.VersionService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServiceServer).Publish(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _VersionService_Unpublish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServiceServer).Unpublish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.VersionService/Unpublish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServiceServer).Unpublish(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _VersionService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/version.VersionService/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServiceServer).UpdateConfig(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _VersionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "version.VersionService",
	HandlerType: (*VersionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _VersionService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _VersionService_Stop_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _VersionService_Publish_Handler,
		},
		{
			MethodName: "Unpublish",
			Handler:    _VersionService_Unpublish_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _VersionService_UpdateConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/versionpb/version.proto",
}