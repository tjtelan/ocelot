// Code generated by protoc-gen-go. DO NOT EDIT.
// source: respositories.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PaginatedRepository struct {
	Pagelen float64                                 `protobuf:"fixed64,1,opt,name=pagelen" json:"pagelen,omitempty"`
	Size    float64                                 `protobuf:"fixed64,2,opt,name=size" json:"size,omitempty"`
	Values  []*PaginatedRepository_RepositoryValues `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
	Page    float64                                 `protobuf:"fixed64,4,opt,name=page" json:"page,omitempty"`
	Next    string                                  `protobuf:"bytes,5,opt,name=next" json:"next,omitempty"`
}

func (m *PaginatedRepository) Reset()                    { *m = PaginatedRepository{} }
func (m *PaginatedRepository) String() string            { return proto.CompactTextString(m) }
func (*PaginatedRepository) ProtoMessage()               {}
func (*PaginatedRepository) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *PaginatedRepository) GetPagelen() float64 {
	if m != nil {
		return m.Pagelen
	}
	return 0
}

func (m *PaginatedRepository) GetSize() float64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PaginatedRepository) GetValues() []*PaginatedRepository_RepositoryValues {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *PaginatedRepository) GetPage() float64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *PaginatedRepository) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

type PaginatedRepository_RepositoryValues struct {
	Name        string                                                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Links       *PaginatedRepository_RepositoryValues_RepositoryLinks `protobuf:"bytes,2,opt,name=links" json:"links,omitempty"`
	Project     *PaginatedRepository_RepositoryValues_Project         `protobuf:"bytes,3,opt,name=project" json:"project,omitempty"`
	CreatedOn   *google_protobuf.Timestamp                            `protobuf:"bytes,4,opt,name=created_on,json=createdOn" json:"created_on,omitempty"`
	Mainbranch  *PaginatedRepository_RepositoryValues_MainBranch      `protobuf:"bytes,5,opt,name=mainbranch" json:"mainbranch,omitempty"`
	FullName    string                                                `protobuf:"bytes,6,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	UpdatedOn   *google_protobuf.Timestamp                            `protobuf:"bytes,7,opt,name=updated_on,json=updatedOn" json:"updated_on,omitempty"`
	Size        float64                                               `protobuf:"fixed64,8,opt,name=size" json:"size,omitempty"`
	Type        string                                                `protobuf:"bytes,9,opt,name=type" json:"type,omitempty"`
	Slug        string                                                `protobuf:"bytes,10,opt,name=slug" json:"slug,omitempty"`
	IsPrivate   bool                                                  `protobuf:"varint,11,opt,name=is_private,json=isPrivate" json:"is_private,omitempty"`
	Description string                                                `protobuf:"bytes,12,opt,name=description" json:"description,omitempty"`
}

func (m *PaginatedRepository_RepositoryValues) Reset()         { *m = PaginatedRepository_RepositoryValues{} }
func (m *PaginatedRepository_RepositoryValues) String() string { return proto.CompactTextString(m) }
func (*PaginatedRepository_RepositoryValues) ProtoMessage()    {}
func (*PaginatedRepository_RepositoryValues) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0}
}

func (m *PaginatedRepository_RepositoryValues) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues) GetLinks() *PaginatedRepository_RepositoryValues_RepositoryLinks {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues) GetProject() *PaginatedRepository_RepositoryValues_Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues) GetCreatedOn() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedOn
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues) GetMainbranch() *PaginatedRepository_RepositoryValues_MainBranch {
	if m != nil {
		return m.Mainbranch
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues) GetUpdatedOn() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedOn
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues) GetSize() float64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PaginatedRepository_RepositoryValues) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues) GetIsPrivate() bool {
	if m != nil {
		return m.IsPrivate
	}
	return false
}

func (m *PaginatedRepository_RepositoryValues) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type PaginatedRepository_RepositoryValues_RepositoryLinks struct {
	Watchers     *LinkUrl       `protobuf:"bytes,1,opt,name=watchers" json:"watchers,omitempty"`
	Branches     *LinkUrl       `protobuf:"bytes,2,opt,name=branches" json:"branches,omitempty"`
	Tags         *LinkUrl       `protobuf:"bytes,3,opt,name=tags" json:"tags,omitempty"`
	Commits      *LinkUrl       `protobuf:"bytes,4,opt,name=commits" json:"commits,omitempty"`
	Clone        []*LinkAndName `protobuf:"bytes,5,rep,name=clone" json:"clone,omitempty"`
	Self         *LinkUrl       `protobuf:"bytes,6,opt,name=self" json:"self,omitempty"`
	Source       *LinkUrl       `protobuf:"bytes,7,opt,name=source" json:"source,omitempty"`
	Hooks        *LinkUrl       `protobuf:"bytes,8,opt,name=hooks" json:"hooks,omitempty"`
	Forks        *LinkUrl       `protobuf:"bytes,9,opt,name=forks" json:"forks,omitempty"`
	Downloads    *LinkUrl       `protobuf:"bytes,10,opt,name=downloads" json:"downloads,omitempty"`
	Pullrequests *LinkUrl       `protobuf:"bytes,11,opt,name=pullrequests" json:"pullrequests,omitempty"`
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) Reset() {
	*m = PaginatedRepository_RepositoryValues_RepositoryLinks{}
}
func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) String() string {
	return proto.CompactTextString(m)
}
func (*PaginatedRepository_RepositoryValues_RepositoryLinks) ProtoMessage() {}
func (*PaginatedRepository_RepositoryValues_RepositoryLinks) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0, 0}
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetWatchers() *LinkUrl {
	if m != nil {
		return m.Watchers
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetBranches() *LinkUrl {
	if m != nil {
		return m.Branches
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetTags() *LinkUrl {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetCommits() *LinkUrl {
	if m != nil {
		return m.Commits
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetClone() []*LinkAndName {
	if m != nil {
		return m.Clone
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetSelf() *LinkUrl {
	if m != nil {
		return m.Self
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetSource() *LinkUrl {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetHooks() *LinkUrl {
	if m != nil {
		return m.Hooks
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetForks() *LinkUrl {
	if m != nil {
		return m.Forks
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetDownloads() *LinkUrl {
	if m != nil {
		return m.Downloads
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_RepositoryLinks) GetPullrequests() *LinkUrl {
	if m != nil {
		return m.Pullrequests
	}
	return nil
}

type PaginatedRepository_RepositoryValues_Project struct {
	Key   string                                                `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Type  string                                                `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Uuid  string                                                `protobuf:"bytes,3,opt,name=uuid" json:"uuid,omitempty"`
	Links *PaginatedRepository_RepositoryValues_RepositoryLinks `protobuf:"bytes,4,opt,name=links" json:"links,omitempty"`
	Name  string                                                `protobuf:"bytes,5,opt,name=name" json:"name,omitempty"`
}

func (m *PaginatedRepository_RepositoryValues_Project) Reset() {
	*m = PaginatedRepository_RepositoryValues_Project{}
}
func (m *PaginatedRepository_RepositoryValues_Project) String() string {
	return proto.CompactTextString(m)
}
func (*PaginatedRepository_RepositoryValues_Project) ProtoMessage() {}
func (*PaginatedRepository_RepositoryValues_Project) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0, 1}
}

func (m *PaginatedRepository_RepositoryValues_Project) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues_Project) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues_Project) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues_Project) GetLinks() *PaginatedRepository_RepositoryValues_RepositoryLinks {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *PaginatedRepository_RepositoryValues_Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PaginatedRepository_RepositoryValues_MainBranch struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PaginatedRepository_RepositoryValues_MainBranch) Reset() {
	*m = PaginatedRepository_RepositoryValues_MainBranch{}
}
func (m *PaginatedRepository_RepositoryValues_MainBranch) String() string {
	return proto.CompactTextString(m)
}
func (*PaginatedRepository_RepositoryValues_MainBranch) ProtoMessage() {}
func (*PaginatedRepository_RepositoryValues_MainBranch) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 0, 2}
}

func (m *PaginatedRepository_RepositoryValues_MainBranch) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PaginatedRepository_RepositoryValues_MainBranch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// some fields are omitted from the response, but if you need anything else,
// see https://confluence.atlassian.com/bitbucket/changesets-resource-296095208.html#changesetsResource-GETanindividualchangeset
// for a list of avilable data on changesets
type ChangeSetV1 struct {
	Node      string `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	RawAuthor string `protobuf:"bytes,2,opt,name=raw_author,json=rawAuthor" json:"raw_author,omitempty"`
	Author    string `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	RawNode   string `protobuf:"bytes,4,opt,name=raw_node,json=rawNode" json:"raw_node,omitempty"`
	Branch    string `protobuf:"bytes,5,opt,name=branch" json:"branch,omitempty"`
}

func (m *ChangeSetV1) Reset()                    { *m = ChangeSetV1{} }
func (m *ChangeSetV1) String() string            { return proto.CompactTextString(m) }
func (*ChangeSetV1) ProtoMessage()               {}
func (*ChangeSetV1) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ChangeSetV1) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *ChangeSetV1) GetRawAuthor() string {
	if m != nil {
		return m.RawAuthor
	}
	return ""
}

func (m *ChangeSetV1) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *ChangeSetV1) GetRawNode() string {
	if m != nil {
		return m.RawNode
	}
	return ""
}

func (m *ChangeSetV1) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

type PaginatedRepoBranches struct {
	Pagelen float64   `protobuf:"fixed64,1,opt,name=pagelen" json:"pagelen,omitempty"`
	Page    float64   `protobuf:"fixed64,2,opt,name=page" json:"page,omitempty"`
	Next    string    `protobuf:"bytes,3,opt,name=next" json:"next,omitempty"`
	Values  []*Branch `protobuf:"bytes,4,rep,name=values" json:"values,omitempty"`
}

func (m *PaginatedRepoBranches) Reset()                    { *m = PaginatedRepoBranches{} }
func (m *PaginatedRepoBranches) String() string            { return proto.CompactTextString(m) }
func (*PaginatedRepoBranches) ProtoMessage()               {}
func (*PaginatedRepoBranches) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *PaginatedRepoBranches) GetPagelen() float64 {
	if m != nil {
		return m.Pagelen
	}
	return 0
}

func (m *PaginatedRepoBranches) GetPage() float64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *PaginatedRepoBranches) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func (m *PaginatedRepoBranches) GetValues() []*Branch {
	if m != nil {
		return m.Values
	}
	return nil
}

type Branch struct {
	Name   string         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Target *Branch_Target `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
}

func (m *Branch) Reset()                    { *m = Branch{} }
func (m *Branch) String() string            { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()               {}
func (*Branch) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *Branch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Branch) GetTarget() *Branch_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

type Branch_Target struct {
	Hash    string                     `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Author  *Author                    `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	Date    *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=date" json:"date,omitempty"`
	Message string                     `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	Type    string                     `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
}

func (m *Branch_Target) Reset()                    { *m = Branch_Target{} }
func (m *Branch_Target) String() string            { return proto.CompactTextString(m) }
func (*Branch_Target) ProtoMessage()               {}
func (*Branch_Target) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3, 0} }

func (m *Branch_Target) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Branch_Target) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Branch_Target) GetDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Branch_Target) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Branch_Target) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*PaginatedRepository)(nil), "protos.PaginatedRepository")
	proto.RegisterType((*PaginatedRepository_RepositoryValues)(nil), "protos.PaginatedRepository.RepositoryValues")
	proto.RegisterType((*PaginatedRepository_RepositoryValues_RepositoryLinks)(nil), "protos.PaginatedRepository.RepositoryValues.RepositoryLinks")
	proto.RegisterType((*PaginatedRepository_RepositoryValues_Project)(nil), "protos.PaginatedRepository.RepositoryValues.Project")
	proto.RegisterType((*PaginatedRepository_RepositoryValues_MainBranch)(nil), "protos.PaginatedRepository.RepositoryValues.MainBranch")
	proto.RegisterType((*ChangeSetV1)(nil), "protos.ChangeSetV1")
	proto.RegisterType((*PaginatedRepoBranches)(nil), "protos.PaginatedRepoBranches")
	proto.RegisterType((*Branch)(nil), "protos.Branch")
	proto.RegisterType((*Branch_Target)(nil), "protos.Branch.Target")
}

func init() { proto.RegisterFile("respositories.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 848 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x51, 0x6f, 0x1b, 0x45,
	0x10, 0xd6, 0xc5, 0xf6, 0xd9, 0x37, 0x17, 0xd1, 0x6a, 0xa3, 0x56, 0x57, 0x23, 0x84, 0x15, 0x04,
	0xa4, 0x82, 0xba, 0xc2, 0xad, 0x84, 0x2a, 0xf1, 0xd2, 0xc2, 0x23, 0xa4, 0xd1, 0x51, 0xca, 0x63,
	0xb4, 0xf1, 0x4d, 0xec, 0x25, 0xe7, 0xdd, 0x63, 0x77, 0x2f, 0x26, 0x3c, 0xc1, 0x33, 0xff, 0x82,
	0x57, 0x5e, 0xf9, 0x4d, 0x3c, 0xf3, 0x13, 0xd0, 0xce, 0xee, 0xf9, 0xce, 0xb5, 0x55, 0x14, 0xa9,
	0x4f, 0x99, 0x9d, 0xf9, 0xe6, 0xbb, 0xf1, 0x7c, 0xb3, 0xb3, 0x81, 0x23, 0x8d, 0xa6, 0x52, 0x46,
	0x58, 0xa5, 0x05, 0x9a, 0x69, 0xa5, 0x95, 0x55, 0x2c, 0xa6, 0x3f, 0x66, 0xfc, 0xe1, 0x42, 0xa9,
	0x45, 0x89, 0x8f, 0xe9, 0x78, 0x51, 0x5f, 0x3e, 0xb6, 0x62, 0x85, 0xc6, 0xf2, 0x55, 0xe5, 0x81,
	0xe3, 0xc3, 0xb9, 0x5a, 0xad, 0x94, 0x0c, 0xa7, 0x07, 0xfe, 0x84, 0xd7, 0x28, 0x2d, 0x4a, 0x2b,
	0xec, 0x86, 0xf1, 0xf8, 0xef, 0x14, 0x8e, 0xce, 0xf8, 0x42, 0x48, 0x6e, 0xb1, 0xc8, 0x31, 0x7c,
	0xf1, 0x86, 0x65, 0x30, 0xac, 0xf8, 0x02, 0x4b, 0x94, 0x59, 0x34, 0x89, 0x4e, 0xa2, 0xbc, 0x39,
	0x32, 0x06, 0x7d, 0x23, 0x7e, 0xc5, 0xec, 0x80, 0xdc, 0x64, 0xb3, 0x6f, 0x20, 0xbe, 0xe6, 0x65,
	0x8d, 0x26, 0xeb, 0x4d, 0x7a, 0x27, 0xe9, 0xec, 0x73, 0xcf, 0x6e, 0xa6, 0x7b, 0xa8, 0xa7, 0xad,
	0xf9, 0x9a, 0x72, 0xf2, 0x90, 0xeb, 0x98, 0xdd, 0x47, 0xb2, 0xbe, 0x67, 0x76, 0xb6, 0xf3, 0x49,
	0xfc, 0xc5, 0x66, 0x83, 0x49, 0x74, 0x92, 0xe4, 0x64, 0x8f, 0x7f, 0x03, 0xb8, 0xfb, 0x26, 0x09,
	0x01, 0xf9, 0x0a, 0xa9, 0x5a, 0x07, 0xe4, 0x2b, 0x64, 0x39, 0x0c, 0x4a, 0x21, 0xaf, 0x0c, 0xd5,
	0x9a, 0xce, 0xbe, 0xba, 0x4d, 0x55, 0x1d, 0xc7, 0xb7, 0x8e, 0x23, 0xf7, 0x54, 0xec, 0x14, 0x86,
	0x95, 0x56, 0x3f, 0xe1, 0xdc, 0x66, 0x3d, 0x62, 0x7d, 0x7a, 0x2b, 0xd6, 0x33, 0x9f, 0x9b, 0x37,
	0x24, 0xec, 0x19, 0xc0, 0x5c, 0xa3, 0x4b, 0x3b, 0x57, 0x92, 0x7e, 0x7a, 0x3a, 0x1b, 0x4f, 0xbd,
	0xbe, 0xd3, 0x46, 0xdf, 0xe9, 0xab, 0x46, 0xdf, 0x3c, 0x09, 0xe8, 0x97, 0x92, 0xfd, 0x08, 0xb0,
	0xe2, 0x42, 0x5e, 0x68, 0x2e, 0xe7, 0x4b, 0xea, 0x50, 0x3a, 0xfb, 0xf2, 0x56, 0xd5, 0x7c, 0xc7,
	0x85, 0x7c, 0x41, 0xe9, 0x79, 0x87, 0x8a, 0xbd, 0x0f, 0xc9, 0x65, 0x5d, 0x96, 0xe7, 0xd4, 0xd0,
	0x98, 0x1a, 0x3a, 0x72, 0x8e, 0x53, 0xd7, 0xd4, 0x67, 0x00, 0x75, 0x55, 0x34, 0x05, 0x0f, 0xff,
	0xbf, 0xe0, 0x80, 0x7e, 0xd9, 0x8e, 0xce, 0xa8, 0x33, 0x3a, 0x0c, 0xfa, 0xf6, 0xa6, 0xc2, 0x2c,
	0xf1, 0xba, 0x39, 0x9b, 0x70, 0x65, 0xbd, 0xc8, 0xc0, 0xfb, 0x9c, 0xcd, 0x3e, 0x00, 0x10, 0xe6,
	0xbc, 0xd2, 0xe2, 0x9a, 0x5b, 0xcc, 0xd2, 0x49, 0x74, 0x32, 0xca, 0x13, 0x61, 0xce, 0xbc, 0x83,
	0x4d, 0x20, 0x2d, 0xd0, 0xcc, 0xb5, 0xa8, 0xac, 0x50, 0x32, 0x3b, 0xa4, 0xcc, 0xae, 0x6b, 0xfc,
	0x4f, 0x0f, 0xee, 0xbc, 0xa1, 0x29, 0xfb, 0x0c, 0x46, 0x6b, 0x6e, 0xe7, 0x4b, 0xd4, 0x86, 0x06,
	0x27, 0x9d, 0xdd, 0x69, 0xfa, 0xe7, 0x00, 0x3f, 0xe8, 0x32, 0xdf, 0x00, 0x1c, 0xd8, 0xf7, 0x07,
	0x9b, 0x81, 0xda, 0x05, 0x37, 0x00, 0xf6, 0x11, 0xf4, 0x2d, 0x5f, 0x98, 0x30, 0x23, 0x3b, 0x40,
	0x0a, 0xb2, 0x87, 0x30, 0x74, 0x37, 0x53, 0x58, 0x13, 0x84, 0xdf, 0xc1, 0x35, 0x71, 0xf6, 0x10,
	0x06, 0xf3, 0x52, 0x49, 0xcc, 0x06, 0x74, 0xc1, 0x8e, 0xba, 0xc0, 0xe7, 0xb2, 0x70, 0xca, 0xe4,
	0x1e, 0xe1, 0x3e, 0x6d, 0xb0, 0xbc, 0x24, 0xe1, 0xf6, 0x7d, 0xda, 0x05, 0xd9, 0xa7, 0x10, 0x1b,
	0x55, 0xeb, 0x39, 0x06, 0x05, 0x77, 0x60, 0x21, 0xcc, 0x3e, 0x86, 0xc1, 0x52, 0xa9, 0x2b, 0x43,
	0xa2, 0xed, 0xc1, 0xf9, 0xa8, 0x83, 0x5d, 0x2a, 0x7d, 0x65, 0x48, 0xc7, 0x7d, 0x30, 0x8a, 0xb2,
	0x47, 0x90, 0x14, 0x6a, 0x2d, 0x4b, 0xc5, 0x0b, 0x43, 0xf2, 0xee, 0x81, 0xb6, 0x08, 0xf6, 0x04,
	0x0e, 0xab, 0xba, 0x2c, 0x35, 0xfe, 0x5c, 0xa3, 0xb1, 0x86, 0x64, 0xdf, 0x93, 0xb1, 0x05, 0x1a,
	0xff, 0x15, 0xc1, 0x30, 0x5c, 0x33, 0x76, 0x17, 0x7a, 0x57, 0x78, 0x13, 0x96, 0x82, 0x33, 0x37,
	0xf3, 0x76, 0xb0, 0x3d, 0x6f, 0x75, 0x2d, 0x0a, 0x12, 0x2b, 0xc9, 0xc9, 0x6e, 0x77, 0x47, 0xff,
	0xdd, 0xed, 0x8e, 0x66, 0x47, 0x0d, 0xda, 0x1d, 0x35, 0x7e, 0x0a, 0xd0, 0xde, 0xc2, 0x4d, 0x75,
	0xd1, 0x76, 0x75, 0x94, 0x75, 0xd0, 0x66, 0x1d, 0xff, 0x11, 0x41, 0xfa, 0xf5, 0x92, 0xcb, 0x05,
	0x7e, 0x8f, 0xf6, 0xf5, 0x17, 0x84, 0x51, 0x45, 0xbb, 0xfd, 0x54, 0x81, 0xee, 0xc6, 0x68, 0xbe,
	0x3e, 0xe7, 0xb5, 0x5d, 0x2a, 0x1d, 0xb2, 0x13, 0xcd, 0xd7, 0xcf, 0xc9, 0xc1, 0xee, 0x43, 0x1c,
	0x42, 0xfe, 0x67, 0x87, 0x13, 0x7b, 0x00, 0x23, 0x97, 0x46, 0x74, 0x7d, 0x8a, 0x0c, 0x35, 0x5f,
	0x9f, 0x3a, 0xc6, 0xfb, 0x10, 0x77, 0x96, 0x4d, 0x92, 0x87, 0xd3, 0xf1, 0xef, 0x11, 0xdc, 0xdb,
	0xea, 0xcb, 0x8b, 0xe6, 0x1a, 0xbc, 0xf5, 0x19, 0xa1, 0x65, 0x7f, 0xb0, 0x67, 0xd9, 0xf7, 0xda,
	0x65, 0xcf, 0x3e, 0xd9, 0x3c, 0x2d, 0x7d, 0x9a, 0xfc, 0xf7, 0x1a, 0x21, 0xc2, 0xde, 0x0a, 0xd1,
	0xe3, 0x7f, 0x23, 0x88, 0xdb, 0x26, 0xee, 0x3c, 0x05, 0x8f, 0x20, 0xb6, 0x5c, 0x2f, 0xd0, 0x86,
	0xab, 0x7b, 0x6f, 0x9b, 0x66, 0xfa, 0x8a, 0x82, 0x79, 0x00, 0x8d, 0xff, 0x8c, 0x20, 0xf6, 0x2e,
	0xc7, 0xb6, 0xe4, 0x66, 0xd9, 0xb0, 0x39, 0xdb, 0x15, 0xd5, 0x69, 0x6b, 0xa7, 0x28, 0xdf, 0xdb,
	0x4d, 0x2f, 0xa7, 0xd0, 0x77, 0xbb, 0x2f, 0x6c, 0x81, 0xb7, 0x6d, 0x49, 0xc2, 0xb9, 0x76, 0xad,
	0xd0, 0x98, 0xe6, 0x11, 0x4c, 0xf2, 0xe6, 0xb8, 0x19, 0x8c, 0x41, 0x3b, 0x18, 0x17, 0xfe, 0xbf,
	0x81, 0x27, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x90, 0x05, 0x3e, 0x98, 0x2b, 0x08, 0x00, 0x00,
}
