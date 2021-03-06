// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1/finding.proto

package securitycenter // import "google.golang.org/genproto/googleapis/cloud/securitycenter/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The state of the finding.
type Finding_State int32

const (
	// Unspecified state.
	Finding_STATE_UNSPECIFIED Finding_State = 0
	// The finding requires attention and has not been addressed yet.
	Finding_ACTIVE Finding_State = 1
	// The finding has been fixed, triaged as a non-issue or otherwise addressed
	// and is no longer active.
	Finding_INACTIVE Finding_State = 2
)

var Finding_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "ACTIVE",
	2: "INACTIVE",
}
var Finding_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"ACTIVE":            1,
	"INACTIVE":          2,
}

func (x Finding_State) String() string {
	return proto.EnumName(Finding_State_name, int32(x))
}
func (Finding_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_finding_8dd8f7ecab8b405d, []int{0, 0}
}

// Cloud Security Command Center (Cloud SCC) finding.
//
// A finding is a record of assessment data (security, risk, health or privacy)
// ingested into Cloud SCC for presentation, notification, analysis,
// policy testing, and enforcement. For example, an XSS vulnerability in an
// App Engine application is a finding.
type Finding struct {
	// The relative resource name of this finding. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/123/sources/456/findings/789"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The relative resource name of the source the finding belongs to. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// This field is immutable after creation time.
	// For example:
	// "organizations/123/sources/456"
	Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	// The full resource name of the Google Cloud Platform (GCP) resource this
	// finding is for. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	// This field is immutable after creation time.
	ResourceName string `protobuf:"bytes,3,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The state of the finding.
	State Finding_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.cloud.securitycenter.v1.Finding_State" json:"state,omitempty"`
	// The additional taxonomy group within findings from a given source.
	// This field is immutable after creation time.
	// Example: "XSS_FLASH_INJECTION"
	Category string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	// The URI that, if available, points to a web page outside of Cloud SCC
	// where additional information about the finding can be found. This field is
	// guaranteed to be either empty or a well formed URL.
	ExternalUri string `protobuf:"bytes,6,opt,name=external_uri,json=externalUri,proto3" json:"external_uri,omitempty"`
	// Source specific properties. These properties are managed by the source
	// that writes the finding. The key names in the source_properties map must be
	// between 1 and 255 characters, and must start with a letter and contain
	// alphanumeric characters or underscores only.
	SourceProperties map[string]*_struct.Value `protobuf:"bytes,7,rep,name=source_properties,json=sourceProperties,proto3" json:"source_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. User specified security marks. These marks are entirely
	// managed by the user and come from the SecurityMarks resource that belongs
	// to the finding.
	SecurityMarks *SecurityMarks `protobuf:"bytes,8,opt,name=security_marks,json=securityMarks,proto3" json:"security_marks,omitempty"`
	// The time at which the event took place. For example, if the finding
	// represents an open firewall it would capture the time the open firewall was
	// detected.
	EventTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// The time at which the finding was created in Cloud SCC.
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Finding) Reset()         { *m = Finding{} }
func (m *Finding) String() string { return proto.CompactTextString(m) }
func (*Finding) ProtoMessage()    {}
func (*Finding) Descriptor() ([]byte, []int) {
	return fileDescriptor_finding_8dd8f7ecab8b405d, []int{0}
}
func (m *Finding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Finding.Unmarshal(m, b)
}
func (m *Finding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Finding.Marshal(b, m, deterministic)
}
func (dst *Finding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Finding.Merge(dst, src)
}
func (m *Finding) XXX_Size() int {
	return xxx_messageInfo_Finding.Size(m)
}
func (m *Finding) XXX_DiscardUnknown() {
	xxx_messageInfo_Finding.DiscardUnknown(m)
}

var xxx_messageInfo_Finding proto.InternalMessageInfo

func (m *Finding) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Finding) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Finding) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Finding) GetState() Finding_State {
	if m != nil {
		return m.State
	}
	return Finding_STATE_UNSPECIFIED
}

func (m *Finding) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Finding) GetExternalUri() string {
	if m != nil {
		return m.ExternalUri
	}
	return ""
}

func (m *Finding) GetSourceProperties() map[string]*_struct.Value {
	if m != nil {
		return m.SourceProperties
	}
	return nil
}

func (m *Finding) GetSecurityMarks() *SecurityMarks {
	if m != nil {
		return m.SecurityMarks
	}
	return nil
}

func (m *Finding) GetEventTime() *timestamp.Timestamp {
	if m != nil {
		return m.EventTime
	}
	return nil
}

func (m *Finding) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func init() {
	proto.RegisterType((*Finding)(nil), "google.cloud.securitycenter.v1.Finding")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "google.cloud.securitycenter.v1.Finding.SourcePropertiesEntry")
	proto.RegisterEnum("google.cloud.securitycenter.v1.Finding_State", Finding_State_name, Finding_State_value)
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1/finding.proto", fileDescriptor_finding_8dd8f7ecab8b405d)
}

var fileDescriptor_finding_8dd8f7ecab8b405d = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0x27, 0xeb, 0xda, 0xb5, 0xaf, 0xdd, 0xd4, 0x59, 0xda, 0x14, 0x55, 0x08, 0x4a, 0xb9, 0xf4,
	0x30, 0x12, 0xad, 0xbb, 0x0c, 0x26, 0x0e, 0xa3, 0x64, 0x52, 0x25, 0xa8, 0xaa, 0x34, 0xdb, 0x01,
	0x0e, 0x91, 0x97, 0xbd, 0x45, 0x66, 0x89, 0x1d, 0xd9, 0x4e, 0x45, 0x3f, 0x15, 0x5f, 0x11, 0xc5,
	0x49, 0x26, 0x3a, 0xfe, 0x94, 0x5b, 0x7e, 0xff, 0x9e, 0xdf, 0xf3, 0x73, 0xe0, 0x24, 0x16, 0x22,
	0x4e, 0xd0, 0x8d, 0x12, 0x91, 0xdf, 0xb9, 0x0a, 0xa3, 0x5c, 0x32, 0xbd, 0x8e, 0x90, 0x6b, 0x94,
	0xee, 0xea, 0xd4, 0xbd, 0x67, 0xfc, 0x8e, 0xf1, 0xd8, 0xc9, 0xa4, 0xd0, 0x82, 0xbc, 0x28, 0xdd,
	0x8e, 0x71, 0x3b, 0x9b, 0x6e, 0x67, 0x75, 0x3a, 0x78, 0x5e, 0x55, 0xa3, 0x19, 0x73, 0x29, 0xe7,
	0x42, 0x53, 0xcd, 0x04, 0x57, 0x65, 0x7a, 0x70, 0xb6, 0xe5, 0xac, 0x9a, 0x09, 0x53, 0x2a, 0x1f,
	0xea, 0x50, 0x5d, 0xd2, 0xa0, 0xdb, 0xfc, 0xde, 0x55, 0x5a, 0xe6, 0x91, 0xae, 0xd4, 0x97, 0x4f,
	0x55, 0xcd, 0x52, 0x54, 0x9a, 0xa6, 0x59, 0x69, 0x18, 0xfd, 0x68, 0xc2, 0xde, 0x55, 0x39, 0x03,
	0x21, 0xb0, 0xcb, 0x69, 0x8a, 0xb6, 0x35, 0xb4, 0xc6, 0x1d, 0xdf, 0x7c, 0x93, 0x63, 0x68, 0x65,
	0x54, 0x22, 0xd7, 0xf6, 0x8e, 0x61, 0x2b, 0x44, 0x5e, 0xc3, 0xbe, 0x44, 0x25, 0x72, 0x19, 0x61,
	0x68, 0x42, 0x0d, 0x23, 0xf7, 0x6a, 0x72, 0x5e, 0x84, 0xa7, 0xd0, 0x54, 0x9a, 0x6a, 0xb4, 0x77,
	0x87, 0xd6, 0xf8, 0x60, 0xf2, 0xc6, 0xf9, 0xf7, 0xf5, 0x38, 0x55, 0x23, 0xce, 0xb2, 0x08, 0xf9,
	0x65, 0x96, 0x0c, 0xa0, 0x1d, 0x51, 0x8d, 0xb1, 0x90, 0x6b, 0xbb, 0x69, 0x0e, 0x79, 0xc4, 0xe4,
	0x15, 0xf4, 0xf0, 0xbb, 0x46, 0xc9, 0x69, 0x12, 0xe6, 0x92, 0xd9, 0x2d, 0xa3, 0x77, 0x6b, 0xee,
	0x5a, 0x32, 0xf2, 0x0d, 0x0e, 0xab, 0x36, 0x33, 0x29, 0x32, 0x94, 0x9a, 0xa1, 0xb2, 0xf7, 0x86,
	0x8d, 0x71, 0x77, 0xf2, 0xfe, 0xbf, 0xfb, 0x31, 0x05, 0x16, 0x8f, 0x79, 0x8f, 0x6b, 0xb9, 0xf6,
	0xfb, 0xea, 0x09, 0x4d, 0x02, 0x38, 0xd8, 0xdc, 0x91, 0xdd, 0x1e, 0x5a, 0xe3, 0xee, 0xf6, 0xc1,
	0x97, 0x15, 0xf3, 0xb9, 0x08, 0xf9, 0xfb, 0xea, 0x57, 0x48, 0xde, 0x02, 0xe0, 0x0a, 0xb9, 0x0e,
	0x8b, 0xdd, 0xd9, 0x1d, 0x53, 0x71, 0x50, 0x57, 0xac, 0x17, 0xeb, 0x04, 0xf5, 0x62, 0xfd, 0x8e,
	0x71, 0x17, 0x98, 0x5c, 0x40, 0x37, 0x92, 0x48, 0x35, 0x96, 0x59, 0xd8, 0x9a, 0x85, 0xd2, 0x5e,
	0x10, 0x83, 0xaf, 0x70, 0xf4, 0xc7, 0xc1, 0x49, 0x1f, 0x1a, 0x0f, 0xb8, 0xae, 0x9e, 0x49, 0xf1,
	0x49, 0x4e, 0xa0, 0xb9, 0xa2, 0x49, 0x8e, 0xe6, 0x91, 0x74, 0x27, 0xc7, 0xbf, 0x9d, 0x70, 0x53,
	0xa8, 0x7e, 0x69, 0x7a, 0xb7, 0x73, 0x6e, 0x8d, 0xce, 0xa1, 0x69, 0xb6, 0x4c, 0x8e, 0xe0, 0x70,
	0x19, 0x5c, 0x06, 0x5e, 0x78, 0x3d, 0x5f, 0x2e, 0xbc, 0xe9, 0xec, 0x6a, 0xe6, 0x7d, 0xec, 0x3f,
	0x23, 0x00, 0xad, 0xcb, 0x69, 0x30, 0xbb, 0xf1, 0xfa, 0x16, 0xe9, 0x41, 0x7b, 0x36, 0xaf, 0xd0,
	0xce, 0x07, 0x0d, 0xa3, 0x48, 0xa4, 0x5b, 0x6e, 0x74, 0x61, 0x7d, 0xf9, 0x54, 0x39, 0x62, 0x91,
	0x50, 0x1e, 0x3b, 0x42, 0xc6, 0x6e, 0x8c, 0xdc, 0x74, 0xe4, 0x96, 0x12, 0xcd, 0x98, 0xfa, 0xdb,
	0xcf, 0x76, 0xb1, 0xc9, 0xdc, 0xb6, 0x4c, 0xf0, 0xec, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a,
	0x14, 0x31, 0xbc, 0x10, 0x04, 0x00, 0x00,
}
