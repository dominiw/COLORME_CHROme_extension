
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/container-storage-interface/spec/csi.proto

package csi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type PluginCapability_Service_Type int32

const (
	PluginCapability_Service_UNKNOWN PluginCapability_Service_Type = 0
	// CONTROLLER_SERVICE indicates that the Plugin provides RPCs for
	// the ControllerService. Plugins SHOULD provide this capability.
	// In rare cases certain plugins MAY wish to omit the
	// ControllerService entirely from their implementation, but such
	// SHOULD NOT be the common case.
	// The presence of this capability determines whether the CO will
	// attempt to invoke the REQUIRED ControllerService RPCs, as well
	// as specific RPCs as indicated by ControllerGetCapabilities.
	PluginCapability_Service_CONTROLLER_SERVICE PluginCapability_Service_Type = 1
	// VOLUME_ACCESSIBILITY_CONSTRAINTS indicates that the volumes for
	// this plugin MAY NOT be equally accessible by all nodes in the
	// cluster. The CO MUST use the topology information returned by
	// CreateVolumeRequest along with the topology information
	// returned by NodeGetInfo to ensure that a given volume is
	// accessible from a given node when scheduling workloads.
	PluginCapability_Service_VOLUME_ACCESSIBILITY_CONSTRAINTS PluginCapability_Service_Type = 2
	// GROUP_CONTROLLER_SERVICE indicates that the Plugin provides
	// RPCs for operating on groups of volumes. Plugins MAY provide
	// this capability.
	// The presence of this capability determines whether the CO will
	// attempt to invoke the REQUIRED GroupController service RPCs, as
	// well as specific RPCs as indicated by
	// GroupControllerGetCapabilities.
	PluginCapability_Service_GROUP_CONTROLLER_SERVICE PluginCapability_Service_Type = 3
)

var PluginCapability_Service_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "CONTROLLER_SERVICE",
	2: "VOLUME_ACCESSIBILITY_CONSTRAINTS",
	3: "GROUP_CONTROLLER_SERVICE",
}

var PluginCapability_Service_Type_value = map[string]int32{
	"UNKNOWN":                          0,
	"CONTROLLER_SERVICE":               1,
	"VOLUME_ACCESSIBILITY_CONSTRAINTS": 2,
	"GROUP_CONTROLLER_SERVICE":         3,
}

func (x PluginCapability_Service_Type) String() string {
	return proto.EnumName(PluginCapability_Service_Type_name, int32(x))
}

func (PluginCapability_Service_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9cdb00adce470e01, []int{4, 0, 0}
}

type PluginCapability_VolumeExpansion_Type int32

const (
	PluginCapability_VolumeExpansion_UNKNOWN PluginCapability_VolumeExpansion_Type = 0
	// ONLINE indicates that volumes may be expanded when published to
	// a node. When a Plugin implements this capability it MUST
	// implement either the EXPAND_VOLUME controller capability or the
	// EXPAND_VOLUME node capability or both. When a plugin supports
	// ONLINE volume expansion and also has the EXPAND_VOLUME
	// controller capability then the plugin MUST support expansion of
	// volumes currently published and available on a node. When a
	// plugin supports ONLINE volume expansion and also has the
	// EXPAND_VOLUME node capability then the plugin MAY support
	// expansion of node-published volume via NodeExpandVolume.
	//
	// Example 1: Given a shared filesystem volume (e.g. GlusterFs),
	//
	//	the Plugin may set the ONLINE volume expansion capability and
	//	implement ControllerExpandVolume but not NodeExpandVolume.
	//
	// Example 2: Given a block storage volume type (e.g. EBS), the
	//
	//	Plugin may set the ONLINE volume expansion capability and
	//	implement both ControllerExpandVolume and NodeExpandVolume.
	//
	// Example 3: Given a Plugin that supports volume expansion only
	//
	//	upon a node, the Plugin may set the ONLINE volume
	//	expansion capability and implement NodeExpandVolume but not
	//	ControllerExpandVolume.
	PluginCapability_VolumeExpansion_ONLINE PluginCapability_VolumeExpansion_Type = 1
	// OFFLINE indicates that volumes currently published and
	// available on a node SHALL NOT be expanded via
	// ControllerExpandVolume. When a plugin supports OFFLINE volume
	// expansion it MUST implement either the EXPAND_VOLUME controller
	// capability or both the EXPAND_VOLUME controller capability and
	// the EXPAND_VOLUME node capability.
	//
	// Example 1: Given a block storage volume type (e.g. Azure Disk)
	//
	//	that does not support expansion of "node-attached" (i.e.
	//	controller-published) volumes, the Plugin may indicate
	//	OFFLINE volume expansion support and implement both
	//	ControllerExpandVolume and NodeExpandVolume.
	PluginCapability_VolumeExpansion_OFFLINE PluginCapability_VolumeExpansion_Type = 2
)

var PluginCapability_VolumeExpansion_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "ONLINE",
	2: "OFFLINE",
}

var PluginCapability_VolumeExpansion_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"ONLINE":  1,
	"OFFLINE": 2,
}

func (x PluginCapability_VolumeExpansion_Type) String() string {
	return proto.EnumName(PluginCapability_VolumeExpansion_Type_name, int32(x))
}

func (PluginCapability_VolumeExpansion_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9cdb00adce470e01, []int{4, 1, 0}
}

type VolumeCapability_AccessMode_Mode int32

const (
	VolumeCapability_AccessMode_UNKNOWN VolumeCapability_AccessMode_Mode = 0
	// Can only be published once as read/write on a single node, at
	// any given time.
	VolumeCapability_AccessMode_SINGLE_NODE_WRITER VolumeCapability_AccessMode_Mode = 1
	// Can only be published once as readonly on a single node, at
	// any given time.
	VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY VolumeCapability_AccessMode_Mode = 2
	// Can be published as readonly at multiple nodes simultaneously.
	VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY VolumeCapability_AccessMode_Mode = 3
	// Can be published at multiple nodes simultaneously. Only one of
	// the node can be used as read/write. The rest will be readonly.
	VolumeCapability_AccessMode_MULTI_NODE_SINGLE_WRITER VolumeCapability_AccessMode_Mode = 4
	// Can be published as read/write at multiple nodes
	// simultaneously.
	VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER VolumeCapability_AccessMode_Mode = 5
	// Can only be published once as read/write at a single workload
	// on a single node, at any given time. SHOULD be used instead of
	// SINGLE_NODE_WRITER for COs using the experimental
	// SINGLE_NODE_MULTI_WRITER capability.
	VolumeCapability_AccessMode_SINGLE_NODE_SINGLE_WRITER VolumeCapability_AccessMode_Mode = 6
	// Can be published as read/write at multiple workloads on a
	// single node simultaneously. SHOULD be used instead of
	// SINGLE_NODE_WRITER for COs using the experimental
	// SINGLE_NODE_MULTI_WRITER capability.
	VolumeCapability_AccessMode_SINGLE_NODE_MULTI_WRITER VolumeCapability_AccessMode_Mode = 7
)

var VolumeCapability_AccessMode_Mode_name = map[int32]string{
	0: "UNKNOWN",
	1: "SINGLE_NODE_WRITER",
	2: "SINGLE_NODE_READER_ONLY",
	3: "MULTI_NODE_READER_ONLY",
	4: "MULTI_NODE_SINGLE_WRITER",
	5: "MULTI_NODE_MULTI_WRITER",
	6: "SINGLE_NODE_SINGLE_WRITER",
	7: "SINGLE_NODE_MULTI_WRITER",
}

var VolumeCapability_AccessMode_Mode_value = map[string]int32{
	"UNKNOWN":                   0,
	"SINGLE_NODE_WRITER":        1,
	"SINGLE_NODE_READER_ONLY":   2,
	"MULTI_NODE_READER_ONLY":    3,
	"MULTI_NODE_SINGLE_WRITER":  4,
	"MULTI_NODE_MULTI_WRITER":   5,
	"SINGLE_NODE_SINGLE_WRITER": 6,
	"SINGLE_NODE_MULTI_WRITER":  7,
}

func (x VolumeCapability_AccessMode_Mode) String() string {
	return proto.EnumName(VolumeCapability_AccessMode_Mode_name, int32(x))
}

func (VolumeCapability_AccessMode_Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9cdb00adce470e01, []int{10, 2, 0}
}

type ControllerServiceCapability_RPC_Type int32

const (
	ControllerServiceCapability_RPC_UNKNOWN                  ControllerServiceCapability_RPC_Type = 0
	ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME     ControllerServiceCapability_RPC_Type = 1
	ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME ControllerServiceCapability_RPC_Type = 2
	ControllerServiceCapability_RPC_LIST_VOLUMES             ControllerServiceCapability_RPC_Type = 3
	ControllerServiceCapability_RPC_GET_CAPACITY             ControllerServiceCapability_RPC_Type = 4
	// Currently the only way to consume a snapshot is to create
	// a volume from it. Therefore plugins supporting
	// CREATE_DELETE_SNAPSHOT MUST support creating volume from
	// snapshot.
	ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT ControllerServiceCapability_RPC_Type = 5
	ControllerServiceCapability_RPC_LIST_SNAPSHOTS         ControllerServiceCapability_RPC_Type = 6
	// Plugins supporting volume cloning at the storage level MAY
	// report this capability. The source volume MUST be managed by
	// the same plugin. Not all volume sources and parameters
	// combinations MAY work.
	ControllerServiceCapability_RPC_CLONE_VOLUME ControllerServiceCapability_RPC_Type = 7
	// Indicates the SP supports ControllerPublishVolume.readonly
	// field.
	ControllerServiceCapability_RPC_PUBLISH_READONLY ControllerServiceCapability_RPC_Type = 8
	// See VolumeExpansion for details.
	ControllerServiceCapability_RPC_EXPAND_VOLUME ControllerServiceCapability_RPC_Type = 9
	// Indicates the SP supports the
	// ListVolumesResponse.entry.published_node_ids field and the
	// ControllerGetVolumeResponse.published_node_ids field.
	// The SP MUST also support PUBLISH_UNPUBLISH_VOLUME.
	ControllerServiceCapability_RPC_LIST_VOLUMES_PUBLISHED_NODES ControllerServiceCapability_RPC_Type = 10
	// Indicates that the Controller service can report volume
	// conditions.
	// An SP MAY implement `VolumeCondition` in only the Controller
	// Plugin, only the Node Plugin, or both.
	// If `VolumeCondition` is implemented in both the Controller and
	// Node Plugins, it SHALL report from different perspectives.
	// If for some reason Controller and Node Plugins report
	// misaligned volume conditions, CO SHALL assume the worst case
	// is the truth.
	// Note that, for alpha, `VolumeCondition` is intended be
	// informative for humans only, not for automation.
	ControllerServiceCapability_RPC_VOLUME_CONDITION ControllerServiceCapability_RPC_Type = 11
	// Indicates the SP supports the ControllerGetVolume RPC.
	// This enables COs to, for example, fetch per volume
	// condition after a volume is provisioned.
	ControllerServiceCapability_RPC_GET_VOLUME ControllerServiceCapability_RPC_Type = 12
	// Indicates the SP supports the SINGLE_NODE_SINGLE_WRITER and/or
	// SINGLE_NODE_MULTI_WRITER access modes.
	// These access modes are intended to replace the
	// SINGLE_NODE_WRITER access mode to clarify the number of writers
	// for a volume on a single node. Plugins MUST accept and allow
	// use of the SINGLE_NODE_WRITER access mode when either
	// SINGLE_NODE_SINGLE_WRITER and/or SINGLE_NODE_MULTI_WRITER are
	// supported, in order to permit older COs to continue working.
	ControllerServiceCapability_RPC_SINGLE_NODE_MULTI_WRITER ControllerServiceCapability_RPC_Type = 13
)

var ControllerServiceCapability_RPC_Type_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CREATE_DELETE_VOLUME",
	2:  "PUBLISH_UNPUBLISH_VOLUME",
	3:  "LIST_VOLUMES",
	4:  "GET_CAPACITY",
	5:  "CREATE_DELETE_SNAPSHOT",
	6:  "LIST_SNAPSHOTS",
	7:  "CLONE_VOLUME",
	8:  "PUBLISH_READONLY",
	9:  "EXPAND_VOLUME",
	10: "LIST_VOLUMES_PUBLISHED_NODES",
	11: "VOLUME_CONDITION",
	12: "GET_VOLUME",
	13: "SINGLE_NODE_MULTI_WRITER",
}

var ControllerServiceCapability_RPC_Type_value = map[string]int32{
	"UNKNOWN":                      0,
	"CREATE_DELETE_VOLUME":         1,
	"PUBLISH_UNPUBLISH_VOLUME":     2,
	"LIST_VOLUMES":                 3,
	"GET_CAPACITY":                 4,
	"CREATE_DELETE_SNAPSHOT":       5,
	"LIST_SNAPSHOTS":               6,
	"CLONE_VOLUME":                 7,
	"PUBLISH_READONLY":             8,
	"EXPAND_VOLUME":                9,
	"LIST_VOLUMES_PUBLISHED_NODES": 10,
	"VOLUME_CONDITION":             11,
	"GET_VOLUME":                   12,
	"SINGLE_NODE_MULTI_WRITER":     13,
}

func (x ControllerServiceCapability_RPC_Type) String() string {
	return proto.EnumName(ControllerServiceCapability_RPC_Type_name, int32(x))
}

func (ControllerServiceCapability_RPC_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9cdb00adce470e01, []int{31, 0, 0}
}

type VolumeUsage_Unit int32

const (
	VolumeUsage_UNKNOWN VolumeUsage_Unit = 0
	VolumeUsage_BYTES   VolumeUsage_Unit = 1
	VolumeUsage_INODES  VolumeUsage_Unit = 2
)

var VolumeUsage_Unit_name = map[int32]string{
	0: "UNKNOWN",
	1: "BYTES",
	2: "INODES",
}

var VolumeUsage_Unit_value = map[string]int32{
	"UNKNOWN": 0,
	"BYTES":   1,
	"INODES":  2,
}

func (x VolumeUsage_Unit) String() string {
	return proto.EnumName(VolumeUsage_Unit_name, int32(x))
}

func (VolumeUsage_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9cdb00adce470e01, []int{51, 0}
}

type NodeServiceCapability_RPC_Type int32

const (
	NodeServiceCapability_RPC_UNKNOWN              NodeServiceCapability_RPC_Type = 0
	NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME NodeServiceCapability_RPC_Type = 1
	// If Plugin implements GET_VOLUME_STATS capability
	// then it MUST implement NodeGetVolumeStats RPC
	// call for fetching volume statistics.
	NodeServiceCapability_RPC_GET_VOLUME_STATS NodeServiceCapability_RPC_Type = 2
	// See VolumeExpansion for details.
	NodeServiceCapability_RPC_EXPAND_VOLUME NodeServiceCapability_RPC_Type = 3
	// Indicates that the Node service can report volume conditions.
	// An SP MAY implement `VolumeCondition` in only the Node
	// Plugin, only the Controller Plugin, or both.
	// If `VolumeCondition` is implemented in both the Node and
	// Controller Plugins, it SHALL report from different
	// perspectives.
	// If for some reason Node and Controller Plugins report
	// misaligned volume conditions, CO SHALL assume the worst case
	// is the truth.
	// Note that, for alpha, `VolumeCondition` is intended to be
	// informative for humans only, not for automation.
	NodeServiceCapability_RPC_VOLUME_CONDITION NodeServiceCapability_RPC_Type = 4
	// Indicates the SP supports the SINGLE_NODE_SINGLE_WRITER and/or
	// SINGLE_NODE_MULTI_WRITER access modes.
	// These access modes are intended to replace the
	// SINGLE_NODE_WRITER access mode to clarify the number of writers
	// for a volume on a single node. Plugins MUST accept and allow
	// use of the SINGLE_NODE_WRITER access mode (subject to the
	// processing rules for NodePublishVolume), when either
	// SINGLE_NODE_SINGLE_WRITER and/or SINGLE_NODE_MULTI_WRITER are
	// supported, in order to permit older COs to continue working.
	NodeServiceCapability_RPC_SINGLE_NODE_MULTI_WRITER NodeServiceCapability_RPC_Type = 5
	// Indicates that Node service supports mounting volumes
	// with provided volume group identifier during node stage
	// or node publish RPC calls.
	NodeServiceCapability_RPC_VOLUME_MOUNT_GROUP NodeServiceCapability_RPC_Type = 6
)

var NodeServiceCapability_RPC_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "STAGE_UNSTAGE_VOLUME",
	2: "GET_VOLUME_STATS",
	3: "EXPAND_VOLUME",
	4: "VOLUME_CONDITION",
	5: "SINGLE_NODE_MULTI_WRITER",
	6: "VOLUME_MOUNT_GROUP",
}

var NodeServiceCapability_RPC_Type_value = map[string]int32{
	"UNKNOWN":                  0,
	"STAGE_UNSTAGE_VOLUME":     1,
	"GET_VOLUME_STATS":         2,
	"EXPAND_VOLUME":            3,
	"VOLUME_CONDITION":         4,
	"SINGLE_NODE_MULTI_WRITER": 5,
	"VOLUME_MOUNT_GROUP":       6,
}

func (x NodeServiceCapability_RPC_Type) String() string {
	return proto.EnumName(NodeServiceCapability_RPC_Type_name, int32(x))
}

func (NodeServiceCapability_RPC_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9cdb00adce470e01, []int{55, 0, 0}
}

type GroupControllerServiceCapability_RPC_Type int32

const (
	GroupControllerServiceCapability_RPC_UNKNOWN GroupControllerServiceCapability_RPC_Type = 0
	// Indicates that the group controller plugin supports
	// creating, deleting, and getting details of a volume
	// group snapshot.
	GroupControllerServiceCapability_RPC_CREATE_DELETE_GET_VOLUME_GROUP_SNAPSHOT GroupControllerServiceCapability_RPC_Type = 1
)

var GroupControllerServiceCapability_RPC_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "CREATE_DELETE_GET_VOLUME_GROUP_SNAPSHOT",
}

var GroupControllerServiceCapability_RPC_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"CREATE_DELETE_GET_VOLUME_GROUP_SNAPSHOT": 1,
}

func (x GroupControllerServiceCapability_RPC_Type) String() string {
	return proto.EnumName(GroupControllerServiceCapability_RPC_Type_name, int32(x))
}

func (GroupControllerServiceCapability_RPC_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9cdb00adce470e01, []int{62, 0, 0}
}

type GetPluginInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPluginInfoRequest) Reset()         { *m = GetPluginInfoRequest{} }
func (m *GetPluginInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetPluginInfoRequest) ProtoMessage()    {}
func (*GetPluginInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdb00adce470e01, []int{0}
}

func (m *GetPluginInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginInfoRequest.Unmarshal(m, b)
}
func (m *GetPluginInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetPluginInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginInfoRequest.Merge(m, src)
}
func (m *GetPluginInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetPluginInfoRequest.Size(m)
}
func (m *GetPluginInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginInfoRequest proto.InternalMessageInfo

type GetPluginInfoResponse struct {
	// The name MUST follow domain name notation format
	// (https://tools.ietf.org/html/rfc1035#section-2.3.1). It SHOULD
	// include the plugin's host company name and the plugin name,
	// to minimize the possibility of collisions. It MUST be 63
	// characters or less, beginning and ending with an alphanumeric
	// character ([a-z0-9A-Z]) with dashes (-), dots (.), and
	// alphanumerics between. This field is REQUIRED.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// This field is REQUIRED. Value of this field is opaque to the CO.
	VendorVersion string `protobuf:"bytes,2,opt,name=vendor_version,json=vendorVersion,proto3" json:"vendor_version,omitempty"`
	// This field is OPTIONAL. Values are opaque to the CO.
	Manifest             map[string]string `protobuf:"bytes,3,rep,name=manifest,proto3" json:"manifest,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetPluginInfoResponse) Reset()         { *m = GetPluginInfoResponse{} }
func (m *GetPluginInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetPluginInfoResponse) ProtoMessage()    {}
func (*GetPluginInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdb00adce470e01, []int{1}
}

func (m *GetPluginInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginInfoResponse.Unmarshal(m, b)
}
func (m *GetPluginInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetPluginInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginInfoResponse.Merge(m, src)
}
func (m *GetPluginInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetPluginInfoResponse.Size(m)
}
func (m *GetPluginInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginInfoResponse proto.InternalMessageInfo

func (m *GetPluginInfoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetPluginInfoResponse) GetVendorVersion() string {
	if m != nil {
		return m.VendorVersion
	}
	return ""
}

func (m *GetPluginInfoResponse) GetManifest() map[string]string {
	if m != nil {
		return m.Manifest
	}
	return nil
}

type GetPluginCapabilitiesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPluginCapabilitiesRequest) Reset()         { *m = GetPluginCapabilitiesRequest{} }
func (m *GetPluginCapabilitiesRequest) String() string { return proto.CompactTextString(m) }
func (*GetPluginCapabilitiesRequest) ProtoMessage()    {}
func (*GetPluginCapabilitiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdb00adce470e01, []int{2}
}

func (m *GetPluginCapabilitiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginCapabilitiesRequest.Unmarshal(m, b)
}
func (m *GetPluginCapabilitiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginCapabilitiesRequest.Marshal(b, m, deterministic)
}
func (m *GetPluginCapabilitiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginCapabilitiesRequest.Merge(m, src)
}
func (m *GetPluginCapabilitiesRequest) XXX_Size() int {
	return xxx_messageInfo_GetPluginCapabilitiesRequest.Size(m)
}
func (m *GetPluginCapabilitiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginCapabilitiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginCapabilitiesRequest proto.InternalMessageInfo

type GetPluginCapabilitiesResponse struct {
	// All the capabilities that the controller service supports. This
	// field is OPTIONAL.
	Capabilities         []*PluginCapability `protobuf:"bytes,1,rep,name=capabilities,proto3" json:"capabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetPluginCapabilitiesResponse) Reset()         { *m = GetPluginCapabilitiesResponse{} }
func (m *GetPluginCapabilitiesResponse) String() string { return proto.CompactTextString(m) }
func (*GetPluginCapabilitiesResponse) ProtoMessage()    {}
func (*GetPluginCapabilitiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdb00adce470e01, []int{3}
}

func (m *GetPluginCapabilitiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPluginCapabilitiesResponse.Unmarshal(m, b)
}
func (m *GetPluginCapabilitiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPluginCapabilitiesResponse.Marshal(b, m, deterministic)
}
func (m *GetPluginCapabilitiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPluginCapabilitiesResponse.Merge(m, src)
}
func (m *GetPluginCapabilitiesResponse) XXX_Size() int {
	return xxx_messageInfo_GetPluginCapabilitiesResponse.Size(m)
}
func (m *GetPluginCapabilitiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPluginCapabilitiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPluginCapabilitiesResponse proto.InternalMessageInfo

func (m *GetPluginCapabilitiesResponse) GetCapabilities() []*PluginCapability {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

// Specifies a capability of the plugin.
type PluginCapability struct {
	// Types that are valid to be assigned to Type:
	//
	//	*PluginCapability_Service_
	//	*PluginCapability_VolumeExpansion_
	Type                 isPluginCapability_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PluginCapability) Reset()         { *m = PluginCapability{} }
func (m *PluginCapability) String() string { return proto.CompactTextString(m) }
func (*PluginCapability) ProtoMessage()    {}
func (*PluginCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdb00adce470e01, []int{4}
}

func (m *PluginCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginCapability.Unmarshal(m, b)
}
func (m *PluginCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginCapability.Marshal(b, m, deterministic)
}
func (m *PluginCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginCapability.Merge(m, src)
}
func (m *PluginCapability) XXX_Size() int {
	return xxx_messageInfo_PluginCapability.Size(m)
}
func (m *PluginCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginCapability.DiscardUnknown(m)
}

var xxx_messageInfo_PluginCapability proto.InternalMessageInfo

type isPluginCapability_Type interface {
	isPluginCapability_Type()
}

type PluginCapability_Service_ struct {
	Service *PluginCapability_Service `protobuf:"bytes,1,opt,name=service,proto3,oneof"`
}

type PluginCapability_VolumeExpansion_ struct {
	VolumeExpansion *PluginCapability_VolumeExpansion `protobuf:"bytes,2,opt,name=volume_expansion,json=volumeExpansion,proto3,oneof"`
}

func (*PluginCapability_Service_) isPluginCapability_Type() {}

func (*PluginCapability_VolumeExpansion_) isPluginCapability_Type() {}

func (m *PluginCapability) GetType() isPluginCapability_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *PluginCapability) GetService() *PluginCapability_Service {
	if x, ok := m.GetType().(*PluginCapability_Service_); ok {
		return x.Service
	}
	return nil
}

func (m *PluginCapability) GetVolumeExpansion() *PluginCapability_VolumeExpansion {
	if x, ok := m.GetType().(*PluginCapability_VolumeExpansion_); ok {
		return x.VolumeExpansion
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PluginCapability) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PluginCapability_Service_)(nil),