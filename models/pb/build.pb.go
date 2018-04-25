// Code generated by protoc-gen-go. DO NOT EDIT.
// source: build.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	build.proto
	guideocelot.proto
	werkerserver.proto

It has these top-level messages:
	BuildConfig
	Stage
	Result
	Triggers
	WerkerTask
	BuildReq
	AllCredsWrapper
	CredWrapper
	SSHKeyWrapper
	SSHWrap
	VCSCreds
	RepoCredWrapper
	RepoCreds
	K8SCreds
	K8SCredsWrapper
	StatusQuery
	BuildQuery
	Builds
	BuildRuntimeInfo
	LineResponse
	RepoAccount
	Status
	StageStatus
	BuildSummary
	Summaries
	PollRequest
	Polls
	Exists
	AcctRepo
	AcctRepos
	Request
	Response
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StageResultVal int32

const (
	StageResultVal_PASS StageResultVal = 0
	StageResultVal_FAIL StageResultVal = 1
)

var StageResultVal_name = map[int32]string{
	0: "PASS",
	1: "FAIL",
}
var StageResultVal_value = map[string]int32{
	"PASS": 0,
	"FAIL": 1,
}

func (x StageResultVal) String() string {
	return proto.EnumName(StageResultVal_name, int32(x))
}
func (StageResultVal) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// this is a direct translation of the ocelot.yaml file
type BuildConfig struct {
	Image string `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	// @inject_tag: yaml:"buildTool"
	BuildTool string   `protobuf:"bytes,2,opt,name=buildTool" json:"buildTool,omitempty" yaml:"buildTool"`
	Packages  []string `protobuf:"bytes,3,rep,name=packages" json:"packages,omitempty"`
	Branches  []string `protobuf:"bytes,4,rep,name=branches" json:"branches,omitempty"`
	Env       []string `protobuf:"bytes,5,rep,name=env" json:"env,omitempty"`
	Stages    []*Stage `protobuf:"bytes,7,rep,name=stages" json:"stages,omitempty"`
}

func (m *BuildConfig) Reset()                    { *m = BuildConfig{} }
func (m *BuildConfig) String() string            { return proto.CompactTextString(m) }
func (*BuildConfig) ProtoMessage()               {}
func (*BuildConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BuildConfig) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *BuildConfig) GetBuildTool() string {
	if m != nil {
		return m.BuildTool
	}
	return ""
}

func (m *BuildConfig) GetPackages() []string {
	if m != nil {
		return m.Packages
	}
	return nil
}

func (m *BuildConfig) GetBranches() []string {
	if m != nil {
		return m.Branches
	}
	return nil
}

func (m *BuildConfig) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *BuildConfig) GetStages() []*Stage {
	if m != nil {
		return m.Stages
	}
	return nil
}

type Stage struct {
	Env    []string `protobuf:"bytes,1,rep,name=env" json:"env,omitempty"`
	Script []string `protobuf:"bytes,2,rep,name=script" json:"script,omitempty"`
	Name   string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// @inject_tag: yaml:"trigger"
	Trigger *Triggers `protobuf:"bytes,4,opt,name=trigger" json:"trigger,omitempty" yaml:"trigger"`
}

func (m *Stage) Reset()                    { *m = Stage{} }
func (m *Stage) String() string            { return proto.CompactTextString(m) }
func (*Stage) ProtoMessage()               {}
func (*Stage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Stage) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Stage) GetScript() []string {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *Stage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stage) GetTrigger() *Triggers {
	if m != nil {
		return m.Trigger
	}
	return nil
}

type Result struct {
	Stage    string         `protobuf:"bytes,1,opt,name=stage" json:"stage,omitempty"`
	Status   StageResultVal `protobuf:"varint,2,opt,name=status,enum=models.StageResultVal" json:"status,omitempty"`
	Error    string         `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	Messages []string       `protobuf:"bytes,4,rep,name=messages" json:"messages,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Result) GetStage() string {
	if m != nil {
		return m.Stage
	}
	return ""
}

func (m *Result) GetStatus() StageResultVal {
	if m != nil {
		return m.Status
	}
	return StageResultVal_PASS
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Result) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

type Triggers struct {
	Branches []string `protobuf:"bytes,1,rep,name=branches" json:"branches,omitempty"`
}

func (m *Triggers) Reset()                    { *m = Triggers{} }
func (m *Triggers) String() string            { return proto.CompactTextString(m) }
func (*Triggers) ProtoMessage()               {}
func (*Triggers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Triggers) GetBranches() []string {
	if m != nil {
		return m.Branches
	}
	return nil
}

type WerkerTask struct {
	VaultToken   string       `protobuf:"bytes,1,opt,name=vaultToken" json:"vaultToken,omitempty"`
	CheckoutHash string       `protobuf:"bytes,2,opt,name=checkoutHash" json:"checkoutHash,omitempty"`
	BuildConf    *BuildConfig `protobuf:"bytes,3,opt,name=buildConf" json:"buildConf,omitempty"`
	VcsToken     string       `protobuf:"bytes,4,opt,name=vcsToken" json:"vcsToken,omitempty"`
	VcsType      SubCredType  `protobuf:"varint,9,opt,name=vcsType,enum=models.SubCredType" json:"vcsType,omitempty"`
	FullName     string       `protobuf:"bytes,6,opt,name=fullName" json:"fullName,omitempty"`
	Id           int64        `protobuf:"varint,7,opt,name=id" json:"id,omitempty"`
	Branch       string       `protobuf:"bytes,8,opt,name=branch" json:"branch,omitempty"`
}

func (m *WerkerTask) Reset()                    { *m = WerkerTask{} }
func (m *WerkerTask) String() string            { return proto.CompactTextString(m) }
func (*WerkerTask) ProtoMessage()               {}
func (*WerkerTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WerkerTask) GetVaultToken() string {
	if m != nil {
		return m.VaultToken
	}
	return ""
}

func (m *WerkerTask) GetCheckoutHash() string {
	if m != nil {
		return m.CheckoutHash
	}
	return ""
}

func (m *WerkerTask) GetBuildConf() *BuildConfig {
	if m != nil {
		return m.BuildConf
	}
	return nil
}

func (m *WerkerTask) GetVcsToken() string {
	if m != nil {
		return m.VcsToken
	}
	return ""
}

func (m *WerkerTask) GetVcsType() SubCredType {
	if m != nil {
		return m.VcsType
	}
	return SubCredType_NIL_SCT
}

func (m *WerkerTask) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *WerkerTask) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *WerkerTask) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func init() {
	proto.RegisterType((*BuildConfig)(nil), "models.BuildConfig")
	proto.RegisterType((*Stage)(nil), "models.Stage")
	proto.RegisterType((*Result)(nil), "models.Result")
	proto.RegisterType((*Triggers)(nil), "models.Triggers")
	proto.RegisterType((*WerkerTask)(nil), "models.WerkerTask")
	proto.RegisterEnum("models.StageResultVal", StageResultVal_name, StageResultVal_value)
}

func init() { proto.RegisterFile("build.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x53, 0xef, 0x8a, 0xd3, 0x4e,
	0x14, 0xfd, 0xe5, 0x4f, 0xd3, 0xf6, 0xe6, 0x67, 0xa9, 0xa3, 0x2c, 0xc3, 0x22, 0x12, 0x82, 0x4a,
	0x58, 0xb0, 0x60, 0x7d, 0x82, 0xdd, 0x05, 0x51, 0x10, 0x91, 0xb4, 0x28, 0xf8, 0x2d, 0x7f, 0xee,
	0xa6, 0x21, 0x69, 0x26, 0xce, 0x24, 0x05, 0xbf, 0xf9, 0x0c, 0x3e, 0x8a, 0x4f, 0x28, 0x73, 0x93,
	0x69, 0x37, 0xdf, 0xee, 0x39, 0xf7, 0xe6, 0xcc, 0xc9, 0xb9, 0x33, 0xe0, 0xa7, 0x7d, 0x59, 0xe7,
	0x9b, 0x56, 0x8a, 0x4e, 0x30, 0xef, 0x28, 0x72, 0xac, 0xd5, 0xf5, 0xd3, 0xa2, 0x2f, 0x73, 0x14,
	0x19, 0xd6, 0xa2, 0x1b, 0x5a, 0xe1, 0x5f, 0x0b, 0xfc, 0x3b, 0x3d, 0x7a, 0x2f, 0x9a, 0x87, 0xb2,
	0x60, 0xcf, 0x61, 0x56, 0x1e, 0x93, 0x02, 0xb9, 0x15, 0x58, 0xd1, 0x32, 0x1e, 0x00, 0x7b, 0x01,
	0x4b, 0xd2, 0xdb, 0x0b, 0x51, 0x73, 0x9b, 0x3a, 0x17, 0x82, 0x5d, 0xc3, 0xa2, 0x4d, 0xb2, 0x2a,
	0x29, 0x50, 0x71, 0x27, 0x70, 0xa2, 0x65, 0x7c, 0xc6, 0xba, 0x97, 0xca, 0xa4, 0xc9, 0x0e, 0xa8,
	0xb8, 0x3b, 0xf4, 0x0c, 0x66, 0x6b, 0x70, 0xb0, 0x39, 0xf1, 0x19, 0xd1, 0xba, 0x64, 0xaf, 0xc1,
	0x53, 0x1d, 0xe9, 0xcc, 0x03, 0x27, 0xf2, 0xb7, 0x4f, 0x36, 0x83, 0xf3, 0xcd, 0x4e, 0xb3, 0xf1,
	0xd8, 0x0c, 0x7f, 0xc2, 0x8c, 0x08, 0xa3, 0x60, 0x5d, 0x14, 0xae, 0xc0, 0x53, 0x99, 0x2c, 0xdb,
	0x8e, 0xdb, 0x44, 0x8e, 0x88, 0x31, 0x70, 0x9b, 0xe4, 0x88, 0xdc, 0x21, 0xf3, 0x54, 0xb3, 0x1b,
	0x98, 0x77, 0xb2, 0x2c, 0x0a, 0x94, 0xdc, 0x0d, 0xac, 0xc8, 0xdf, 0xae, 0xcd, 0x71, 0xfb, 0x81,
	0x56, 0xb1, 0x19, 0x08, 0x7f, 0x5b, 0xe0, 0xc5, 0xa8, 0xfa, 0xba, 0xd3, 0x11, 0x91, 0x0f, 0x13,
	0x11, 0x01, 0xb6, 0x21, 0xeb, 0x5d, 0xaf, 0x28, 0x9f, 0xd5, 0xf6, 0x6a, 0x6a, 0x9d, 0x3e, 0xfd,
	0x96, 0xd4, 0xf1, 0x38, 0xa5, 0x55, 0x50, 0x4a, 0x21, 0x47, 0x47, 0x03, 0xd0, 0x71, 0x1d, 0x51,
	0x29, 0x8a, 0x60, 0x8c, 0xcb, 0xe0, 0xf0, 0x0d, 0x2c, 0x8c, 0xaf, 0x49, 0xac, 0xd6, 0x34, 0xd6,
	0xf0, 0x8f, 0x0d, 0xf0, 0x1d, 0x65, 0x85, 0x72, 0x9f, 0xa8, 0x8a, 0xbd, 0x04, 0x38, 0x25, 0x7d,
	0xdd, 0xed, 0x45, 0x85, 0xcd, 0xe8, 0xf9, 0x11, 0xc3, 0x42, 0xf8, 0x3f, 0x3b, 0x60, 0x56, 0x89,
	0xbe, 0xfb, 0x98, 0xa8, 0xc3, 0xb8, 0xde, 0x09, 0xc7, 0xde, 0x8d, 0xfb, 0xd7, 0x97, 0x84, 0x0c,
	0xfb, 0xdb, 0x67, 0xe6, 0xff, 0x1e, 0xdd, 0x9e, 0xf8, 0x32, 0xa5, 0x1d, 0x9e, 0x32, 0x35, 0x1c,
	0xea, 0x92, 0xe4, 0x19, 0xb3, 0xb7, 0x30, 0xd7, 0xf5, 0xaf, 0x16, 0xf9, 0x92, 0xc2, 0x3a, 0x8b,
	0xed, 0xfa, 0xf4, 0x5e, 0x62, 0xae, 0x5b, 0xb1, 0x99, 0xd1, 0x52, 0x0f, 0x7d, 0x5d, 0x7f, 0xd1,
	0xfb, 0xf3, 0x06, 0x29, 0x83, 0xd9, 0x0a, 0xec, 0x32, 0xe7, 0xf3, 0xc0, 0x8a, 0x9c, 0xd8, 0x2e,
	0x73, 0xbd, 0xff, 0x21, 0x08, 0xbe, 0xa0, 0xc9, 0x11, 0xdd, 0xbc, 0x82, 0xd5, 0x74, 0x11, 0x6c,
	0x01, 0xee, 0xd7, 0xdb, 0xdd, 0x6e, 0xfd, 0x9f, 0xae, 0x3e, 0xdc, 0x7e, 0xfa, 0xbc, 0xb6, 0xee,
	0xdc, 0x1f, 0x76, 0x9b, 0xa6, 0x1e, 0x3d, 0x8d, 0xf7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8e,
	0x4a, 0xd4, 0x2c, 0x44, 0x03, 0x00, 0x00,
}
