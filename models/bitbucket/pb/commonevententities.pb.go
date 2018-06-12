// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commonevententities.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// https://confluence.atlassian.com/bitbucket/event-payloads-740262817.html#EventPayloads-entity_userOwner
type Owner struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Links                *Links   `protobuf:"bytes,3,opt,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Owner) Reset()         { *m = Owner{} }
func (m *Owner) String() string { return proto.CompactTextString(m) }
func (*Owner) ProtoMessage()    {}
func (*Owner) Descriptor() ([]byte, []int) {
	return fileDescriptor_commonevententities_aa030c92829d4266, []int{0}
}
func (m *Owner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Owner.Unmarshal(m, b)
}
func (m *Owner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Owner.Marshal(b, m, deterministic)
}
func (dst *Owner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Owner.Merge(dst, src)
}
func (m *Owner) XXX_Size() int {
	return xxx_messageInfo_Owner.Size(m)
}
func (m *Owner) XXX_DiscardUnknown() {
	xxx_messageInfo_Owner.DiscardUnknown(m)
}

var xxx_messageInfo_Owner proto.InternalMessageInfo

func (m *Owner) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Owner) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Owner) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

// https://confluence.atlassian.com/bitbucket/event-payloads-740262817.html#EventPayloads-entity_repositoryRepository
type Repository struct {
	Links                *Links   `protobuf:"bytes,1,opt,name=links,proto3" json:"links,omitempty"`
	Project              *Project `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	FullName             string   `protobuf:"bytes,3,opt,name=fullName,json=full_name,proto3" json:"fullName,omitempty"`
	Website              string   `protobuf:"bytes,4,opt,name=website,proto3" json:"website,omitempty"`
	Owner                *Owner   `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}
func (*Repository) Descriptor() ([]byte, []int) {
	return fileDescriptor_commonevententities_aa030c92829d4266, []int{1}
}
func (m *Repository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repository.Unmarshal(m, b)
}
func (m *Repository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repository.Marshal(b, m, deterministic)
}
func (dst *Repository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repository.Merge(dst, src)
}
func (m *Repository) XXX_Size() int {
	return xxx_messageInfo_Repository.Size(m)
}
func (m *Repository) XXX_DiscardUnknown() {
	xxx_messageInfo_Repository.DiscardUnknown(m)
}

var xxx_messageInfo_Repository proto.InternalMessageInfo

func (m *Repository) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Repository) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *Repository) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *Repository) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Repository) GetOwner() *Owner {
	if m != nil {
		return m.Owner
	}
	return nil
}

// https://confluence.atlassian.com/bitbucket/event-payloads-740262817.html#EventPayloads-Pullrequestevents
type PullRequestEntity struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	State                string               `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	Author               *Owner               `protobuf:"bytes,5,opt,name=author,proto3" json:"author,omitempty"`
	Source               *PRInfo              `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	Destination          *PRInfo              `protobuf:"bytes,7,opt,name=destination,proto3" json:"destination,omitempty"`
	MergeCommit          *Commit              `protobuf:"bytes,8,opt,name=mergeCommit,json=merge_commit,proto3" json:"mergeCommit,omitempty"`
	Participants         []*Owner             `protobuf:"bytes,9,rep,name=participants,proto3" json:"participants,omitempty"`
	Reviewers            []*Owner             `protobuf:"bytes,10,rep,name=reviewers,proto3" json:"reviewers,omitempty"`
	CloseSourceBranch    bool                 `protobuf:"varint,11,opt,name=closeSourceBranch,json=close_source_branch,proto3" json:"closeSourceBranch,omitempty"`
	ClosedBy             *Owner               `protobuf:"bytes,12,opt,name=closedBy,json=closed_by,proto3" json:"closedBy,omitempty"`
	Reason               string               `protobuf:"bytes,13,opt,name=reason,proto3" json:"reason,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,14,opt,name=createdAt,json=created_at,proto3" json:"createdAt,omitempty"`
	UpdatedOn            *timestamp.Timestamp `protobuf:"bytes,15,opt,name=updatedOn,json=updated_on,proto3" json:"updatedOn,omitempty"`
	Links                *Links               `protobuf:"bytes,16,opt,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PullRequestEntity) Reset()         { *m = PullRequestEntity{} }
func (m *PullRequestEntity) String() string { return proto.CompactTextString(m) }
func (*PullRequestEntity) ProtoMessage()    {}
func (*PullRequestEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_commonevententities_aa030c92829d4266, []int{2}
}
func (m *PullRequestEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullRequestEntity.Unmarshal(m, b)
}
func (m *PullRequestEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullRequestEntity.Marshal(b, m, deterministic)
}
func (dst *PullRequestEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullRequestEntity.Merge(dst, src)
}
func (m *PullRequestEntity) XXX_Size() int {
	return xxx_messageInfo_PullRequestEntity.Size(m)
}
func (m *PullRequestEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_PullRequestEntity.DiscardUnknown(m)
}

var xxx_messageInfo_PullRequestEntity proto.InternalMessageInfo

func (m *PullRequestEntity) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PullRequestEntity) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PullRequestEntity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PullRequestEntity) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *PullRequestEntity) GetAuthor() *Owner {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *PullRequestEntity) GetSource() *PRInfo {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *PullRequestEntity) GetDestination() *PRInfo {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *PullRequestEntity) GetMergeCommit() *Commit {
	if m != nil {
		return m.MergeCommit
	}
	return nil
}

func (m *PullRequestEntity) GetParticipants() []*Owner {
	if m != nil {
		return m.Participants
	}
	return nil
}

func (m *PullRequestEntity) GetReviewers() []*Owner {
	if m != nil {
		return m.Reviewers
	}
	return nil
}

func (m *PullRequestEntity) GetCloseSourceBranch() bool {
	if m != nil {
		return m.CloseSourceBranch
	}
	return false
}

func (m *PullRequestEntity) GetClosedBy() *Owner {
	if m != nil {
		return m.ClosedBy
	}
	return nil
}

func (m *PullRequestEntity) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *PullRequestEntity) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *PullRequestEntity) GetUpdatedOn() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedOn
	}
	return nil
}

func (m *PullRequestEntity) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

type PRInfo struct {
	Branch               *PRInfoBr   `protobuf:"bytes,1,opt,name=branch,proto3" json:"branch,omitempty"`
	Commit               *Commit     `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	Repository           *Repository `protobuf:"bytes,3,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PRInfo) Reset()         { *m = PRInfo{} }
func (m *PRInfo) String() string { return proto.CompactTextString(m) }
func (*PRInfo) ProtoMessage()    {}
func (*PRInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_commonevententities_aa030c92829d4266, []int{3}
}
func (m *PRInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PRInfo.Unmarshal(m, b)
}
func (m *PRInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PRInfo.Marshal(b, m, deterministic)
}
func (dst *PRInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PRInfo.Merge(dst, src)
}
func (m *PRInfo) XXX_Size() int {
	return xxx_messageInfo_PRInfo.Size(m)
}
func (m *PRInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PRInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PRInfo proto.InternalMessageInfo

func (m *PRInfo) GetBranch() *PRInfoBr {
	if m != nil {
		return m.Branch
	}
	return nil
}

func (m *PRInfo) GetCommit() *Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *PRInfo) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type PRInfoBr struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PRInfoBr) Reset()         { *m = PRInfoBr{} }
func (m *PRInfoBr) String() string { return proto.CompactTextString(m) }
func (*PRInfoBr) ProtoMessage()    {}
func (*PRInfoBr) Descriptor() ([]byte, []int) {
	return fileDescriptor_commonevententities_aa030c92829d4266, []int{3, 0}
}
func (m *PRInfoBr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PRInfoBr.Unmarshal(m, b)
}
func (m *PRInfoBr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PRInfoBr.Marshal(b, m, deterministic)
}
func (dst *PRInfoBr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PRInfoBr.Merge(dst, src)
}
func (m *PRInfoBr) XXX_Size() int {
	return xxx_messageInfo_PRInfoBr.Size(m)
}
func (m *PRInfoBr) XXX_DiscardUnknown() {
	xxx_messageInfo_PRInfoBr.DiscardUnknown(m)
}

var xxx_messageInfo_PRInfoBr proto.InternalMessageInfo

func (m *PRInfoBr) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// https://confluence.atlassian.com/bitbucket/event-payloads-740262817.html#EventPayloads-entity_projectProject
type Project struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Links                *Links   `protobuf:"bytes,3,opt,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_commonevententities_aa030c92829d4266, []int{4}
}
func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (dst *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(dst, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Project) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

type Changeset struct {
	New                  *Changeset_Head `protobuf:"bytes,1,opt,name=new,proto3" json:"new,omitempty"`
	Old                  *Changeset_Head `protobuf:"bytes,2,opt,name=old,proto3" json:"old,omitempty"`
	Links                *Links          `protobuf:"bytes,3,opt,name=links,proto3" json:"links,omitempty"`
	Closed               bool            `protobuf:"varint,4,opt,name=closed,proto3" json:"closed,omitempty"`
	Created              bool            `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	Forced               bool            `protobuf:"varint,6,opt,name=forced,proto3" json:"forced,omitempty"`
	Commits              []*Commit       `protobuf:"bytes,7,rep,name=commits,proto3" json:"commits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Changeset) Reset()         { *m = Changeset{} }
func (m *Changeset) String() string { return proto.CompactTextString(m) }
func (*Changeset) ProtoMessage()    {}
func (*Changeset) Descriptor() ([]byte, []int) {
	return fileDescriptor_commonevententities_aa030c92829d4266, []int{5}
}
func (m *Changeset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Changeset.Unmarshal(m, b)
}
func (m *Changeset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Changeset.Marshal(b, m, deterministic)
}
func (dst *Changeset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Changeset.Merge(dst, src)
}
func (m *Changeset) XXX_Size() int {
	return xxx_messageInfo_Changeset.Size(m)
}
func (m *Changeset) XXX_DiscardUnknown() {
	xxx_messageInfo_Changeset.DiscardUnknown(m)
}

var xxx_messageInfo_Changeset proto.InternalMessageInfo

func (m *Changeset) GetNew() *Changeset_Head {
	if m != nil {
		return m.New
	}
	return nil
}

func (m *Changeset) GetOld() *Changeset_Head {
	if m != nil {
		return m.Old
	}
	return nil
}

func (m *Changeset) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Changeset) GetClosed() bool {
	if m != nil {
		return m.Closed
	}
	return false
}

func (m *Changeset) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Changeset) GetForced() bool {
	if m != nil {
		return m.Forced
	}
	return false
}

func (m *Changeset) GetCommits() []*Commit {
	if m != nil {
		return m.Commits
	}
	return nil
}

type Changeset_Head struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Target               *Commit  `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Changeset_Head) Reset()         { *m = Changeset_Head{} }
func (m *Changeset_Head) String() string { return proto.CompactTextString(m) }
func (*Changeset_Head) ProtoMessage()    {}
func (*Changeset_Head) Descriptor() ([]byte, []int) {
	return fileDescriptor_commonevententities_aa030c92829d4266, []int{5, 0}
}
func (m *Changeset_Head) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Changeset_Head.Unmarshal(m, b)
}
func (m *Changeset_Head) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Changeset_Head.Marshal(b, m, deterministic)
}
func (dst *Changeset_Head) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Changeset_Head.Merge(dst, src)
}
func (m *Changeset_Head) XXX_Size() int {
	return xxx_messageInfo_Changeset_Head.Size(m)
}
func (m *Changeset_Head) XXX_DiscardUnknown() {
	xxx_messageInfo_Changeset_Head.DiscardUnknown(m)
}

var xxx_messageInfo_Changeset_Head proto.InternalMessageInfo

func (m *Changeset_Head) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Changeset_Head) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Changeset_Head) GetTarget() *Commit {
	if m != nil {
		return m.Target
	}
	return nil
}

type Commit struct {
	Hash    string               `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Author  *Owner               `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Message string               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Date    *timestamp.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	// ignoring the "parents" field
	Links                *Links      `protobuf:"bytes,5,opt,name=links,proto3" json:"links,omitempty"`
	Repository           *Repository `protobuf:"bytes,6,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}
func (*Commit) Descriptor() ([]byte, []int) {
	return fileDescriptor_commonevententities_aa030c92829d4266, []int{6}
}
func (m *Commit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commit.Unmarshal(m, b)
}
func (m *Commit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commit.Marshal(b, m, deterministic)
}
func (dst *Commit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit.Merge(dst, src)
}
func (m *Commit) XXX_Size() int {
	return xxx_messageInfo_Commit.Size(m)
}
func (m *Commit) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit.DiscardUnknown(m)
}

var xxx_messageInfo_Commit proto.InternalMessageInfo

func (m *Commit) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Commit) GetAuthor() *Owner {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Commit) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Commit) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Commit) GetLinks() *Links {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Commit) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

// /2.0/repositories/{username}/{repo_slug}/commits
type Commits struct {
	Values               []*Commit `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Commits) Reset()         { *m = Commits{} }
func (m *Commits) String() string { return proto.CompactTextString(m) }
func (*Commits) ProtoMessage()    {}
func (*Commits) Descriptor() ([]byte, []int) {
	return fileDescriptor_commonevententities_aa030c92829d4266, []int{7}
}
func (m *Commits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commits.Unmarshal(m, b)
}
func (m *Commits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commits.Marshal(b, m, deterministic)
}
func (dst *Commits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commits.Merge(dst, src)
}
func (m *Commits) XXX_Size() int {
	return xxx_messageInfo_Commits.Size(m)
}
func (m *Commits) XXX_DiscardUnknown() {
	xxx_messageInfo_Commits.DiscardUnknown(m)
}

var xxx_messageInfo_Commits proto.InternalMessageInfo

func (m *Commits) GetValues() []*Commit {
	if m != nil {
		return m.Values
	}
	return nil
}

// bitbucket api 1.0
// https://confluence.atlassian.com/bitbucket/src-resources-296095214.html
type RepoSourceFile struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RepoSourceFile) Reset()         { *m = RepoSourceFile{} }
func (m *RepoSourceFile) String() string { return proto.CompactTextString(m) }
func (*RepoSourceFile) ProtoMessage()    {}
func (*RepoSourceFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_commonevententities_aa030c92829d4266, []int{8}
}
func (m *RepoSourceFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepoSourceFile.Unmarshal(m, b)
}
func (m *RepoSourceFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepoSourceFile.Marshal(b, m, deterministic)
}
func (dst *RepoSourceFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepoSourceFile.Merge(dst, src)
}
func (m *RepoSourceFile) XXX_Size() int {
	return xxx_messageInfo_RepoSourceFile.Size(m)
}
func (m *RepoSourceFile) XXX_DiscardUnknown() {
	xxx_messageInfo_RepoSourceFile.DiscardUnknown(m)
}

var xxx_messageInfo_RepoSourceFile proto.InternalMessageInfo

func (m *RepoSourceFile) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *RepoSourceFile) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *RepoSourceFile) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Author struct {
	Raw                  string   `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	User                 *Owner   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Author) Reset()         { *m = Author{} }
func (m *Author) String() string { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()    {}
func (*Author) Descriptor() ([]byte, []int) {
	return fileDescriptor_commonevententities_aa030c92829d4266, []int{9}
}
func (m *Author) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Author.Unmarshal(m, b)
}
func (m *Author) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Author.Marshal(b, m, deterministic)
}
func (dst *Author) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Author.Merge(dst, src)
}
func (m *Author) XXX_Size() int {
	return xxx_messageInfo_Author.Size(m)
}
func (m *Author) XXX_DiscardUnknown() {
	xxx_messageInfo_Author.DiscardUnknown(m)
}

var xxx_messageInfo_Author proto.InternalMessageInfo

func (m *Author) GetRaw() string {
	if m != nil {
		return m.Raw
	}
	return ""
}

func (m *Author) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Author) GetUser() *Owner {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*Owner)(nil), "protos.Owner")
	proto.RegisterType((*Repository)(nil), "protos.Repository")
	proto.RegisterType((*PullRequestEntity)(nil), "protos.PullRequestEntity")
	proto.RegisterType((*PRInfo)(nil), "protos.PRInfo")
	proto.RegisterType((*PRInfoBr)(nil), "protos.PRInfo.br")
	proto.RegisterType((*Project)(nil), "protos.Project")
	proto.RegisterType((*Changeset)(nil), "protos.Changeset")
	proto.RegisterType((*Changeset_Head)(nil), "protos.Changeset.Head")
	proto.RegisterType((*Commit)(nil), "protos.Commit")
	proto.RegisterType((*Commits)(nil), "protos.Commits")
	proto.RegisterType((*RepoSourceFile)(nil), "protos.RepoSourceFile")
	proto.RegisterType((*Author)(nil), "protos.Author")
}

func init() {
	proto.RegisterFile("commonevententities.proto", fileDescriptor_commonevententities_aa030c92829d4266)
}

var fileDescriptor_commonevententities_aa030c92829d4266 = []byte{
	// 856 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x8e, 0x1b, 0x35,
	0x14, 0x56, 0xfe, 0x26, 0xc9, 0xc9, 0x76, 0xdb, 0x35, 0xa8, 0x32, 0xe1, 0x82, 0x10, 0x44, 0x95,
	0x82, 0x34, 0x65, 0x97, 0x2b, 0x2e, 0xdb, 0x0a, 0x04, 0x52, 0x45, 0x8b, 0x41, 0xbd, 0x42, 0x8a,
	0x9c, 0x99, 0xb3, 0x89, 0x61, 0x66, 0x3c, 0xd8, 0x9e, 0x8d, 0xf2, 0x1e, 0xbc, 0x04, 0x6f, 0xc0,
	0x63, 0xf0, 0x14, 0x3c, 0x07, 0xf2, 0x5f, 0x7e, 0x76, 0xb3, 0x61, 0xaf, 0xe2, 0x73, 0xbe, 0xcf,
	0xce, 0xf1, 0x77, 0xfc, 0x9d, 0x81, 0x8f, 0x32, 0x59, 0x96, 0xb2, 0xc2, 0x1b, 0xac, 0x0c, 0x56,
	0x46, 0x18, 0x81, 0x3a, 0xad, 0x95, 0x34, 0x92, 0x24, 0xee, 0x47, 0x8f, 0xcf, 0x3c, 0xc5, 0x67,
	0xc7, 0x9f, 0x2c, 0xa5, 0x5c, 0x16, 0xf8, 0xc2, 0x45, 0x8b, 0xe6, 0xfa, 0x85, 0x11, 0x25, 0x6a,
	0xc3, 0xcb, 0xda, 0x13, 0xa6, 0xbf, 0x42, 0xef, 0xed, 0xba, 0x42, 0x45, 0x08, 0x74, 0xcd, 0xa6,
	0x46, 0xda, 0x9a, 0xb4, 0x66, 0x43, 0xe6, 0xd6, 0x64, 0x0c, 0x83, 0x46, 0xa3, 0xaa, 0x78, 0x89,
	0xb4, 0xed, 0xf2, 0xdb, 0x98, 0x7c, 0x06, 0xbd, 0x42, 0x54, 0xbf, 0x6b, 0xda, 0x99, 0xb4, 0x66,
	0xa3, 0xab, 0x47, 0xfe, 0x3c, 0x9d, 0xbe, 0xb1, 0x49, 0xe6, 0xb1, 0xe9, 0xdf, 0x2d, 0x00, 0x86,
	0xb5, 0xd4, 0xc2, 0x48, 0xb5, 0xd9, 0xed, 0x69, 0xdd, 0xbf, 0x87, 0x3c, 0x87, 0x7e, 0xad, 0xe4,
	0x6f, 0x98, 0x19, 0xf7, 0x9f, 0xa3, 0xab, 0xc7, 0x91, 0xf6, 0xce, 0xa7, 0x59, 0xc4, 0xc9, 0xc7,
	0x30, 0xb8, 0x6e, 0x8a, 0xe2, 0x47, 0x5b, 0x5f, 0xc7, 0xd5, 0x37, 0xb4, 0xf1, 0xdc, 0x15, 0x48,
	0xa1, 0xbf, 0xc6, 0x85, 0x16, 0x06, 0x69, 0xd7, 0x61, 0x31, 0xb4, 0x65, 0x48, 0x7b, 0x67, 0xda,
	0x3b, 0x2c, 0xc3, 0x09, 0xc1, 0x3c, 0x36, 0xfd, 0xb3, 0x07, 0x17, 0xef, 0x9a, 0xa2, 0x60, 0xf8,
	0x47, 0x83, 0xda, 0x7c, 0x6b, 0xd5, 0xde, 0x90, 0x73, 0x68, 0x8b, 0xdc, 0x95, 0xdf, 0x61, 0x6d,
	0x91, 0x93, 0x0f, 0xa1, 0x67, 0x84, 0x29, 0xa2, 0x3c, 0x3e, 0x20, 0x13, 0x18, 0xe5, 0xa8, 0x33,
	0x25, 0x6a, 0x23, 0x64, 0x15, 0x4a, 0xdb, 0x4f, 0xd9, 0x7d, 0xda, 0xf0, 0x6d, 0x69, 0x3e, 0x20,
	0x9f, 0x43, 0xc2, 0x1b, 0xb3, 0x92, 0xf7, 0x54, 0x16, 0x40, 0xf2, 0x0c, 0x12, 0x2d, 0x1b, 0x95,
	0x21, 0x4d, 0x1c, 0xed, 0x7c, 0x2b, 0x10, 0xfb, 0xa1, 0xba, 0x96, 0x2c, 0xa0, 0xe4, 0x2b, 0x57,
	0x86, 0x11, 0x15, 0x77, 0x65, 0xf4, 0x8f, 0x92, 0xf7, 0x29, 0xe4, 0x12, 0x46, 0x25, 0xaa, 0x25,
	0xbe, 0x96, 0x65, 0x29, 0x0c, 0x1d, 0x1c, 0xee, 0xf0, 0x59, 0x76, 0xe6, 0x28, 0xf3, 0xcc, 0x45,
	0xe4, 0x12, 0xce, 0x6a, 0xae, 0x8c, 0xc8, 0x44, 0xcd, 0x2b, 0xa3, 0xe9, 0x70, 0xd2, 0xb9, 0x5b,
	0xf9, 0x01, 0x85, 0x7c, 0x09, 0x43, 0x85, 0x37, 0x02, 0xd7, 0xa8, 0x34, 0x85, 0x63, 0xfc, 0x1d,
	0x4e, 0x52, 0xb8, 0xc8, 0x0a, 0xa9, 0xf1, 0x67, 0x77, 0xa7, 0x57, 0x8a, 0x57, 0xd9, 0x8a, 0x8e,
	0x26, 0xad, 0xd9, 0x80, 0x7d, 0xe0, 0x80, 0xb9, 0xbf, 0xed, 0x7c, 0xe1, 0x20, 0xf2, 0x05, 0x0c,
	0x5c, 0x3a, 0x7f, 0xb5, 0xa1, 0x67, 0xc7, 0x54, 0x1c, 0x7a, 0x78, 0xbe, 0xd8, 0x90, 0xa7, 0x90,
	0x28, 0xe4, 0x5a, 0x56, 0xf4, 0x91, 0x6b, 0x43, 0x88, 0xc8, 0x37, 0x30, 0xcc, 0x14, 0x72, 0x83,
	0xf9, 0x4b, 0x43, 0xcf, 0xdd, 0x21, 0xe3, 0xd4, 0x3b, 0x29, 0x8d, 0x4e, 0x4a, 0x7f, 0x89, 0x4e,
	0x62, 0x10, 0xc8, 0x73, 0x6e, 0xec, 0xd6, 0xa6, 0xce, 0x6d, 0xf4, 0xb6, 0xa2, 0x8f, 0xff, 0x7f,
	0x6b, 0x20, 0xcf, 0x65, 0xb5, 0x73, 0xc7, 0x93, 0x13, 0x8e, 0xfa, 0xab, 0x05, 0x89, 0xef, 0x1c,
	0x79, 0x0e, 0x89, 0xbf, 0x73, 0xb0, 0xd3, 0xc5, 0x61, 0x67, 0xd3, 0x85, 0x62, 0x81, 0x60, 0x5f,
	0x8c, 0x6f, 0x57, 0xb0, 0xd4, 0xed, 0x96, 0x06, 0x94, 0x5c, 0x01, 0xa8, 0xad, 0x5d, 0x83, 0xb3,
	0x49, 0xe4, 0xee, 0x8c, 0xcc, 0xf6, 0x58, 0x63, 0x0a, 0xed, 0x85, 0x1b, 0x1f, 0x6e, 0x4c, 0x84,
	0xf1, 0x61, 0xd7, 0xd3, 0xf7, 0xd0, 0x0f, 0x96, 0x3d, 0x06, 0xdb, 0x5c, 0xd3, 0x88, 0x3c, 0x58,
	0xc7, 0xad, 0x1f, 0x36, 0x55, 0xfe, 0x69, 0xc3, 0xf0, 0xf5, 0x8a, 0x57, 0x4b, 0xd4, 0x68, 0xc8,
	0x0c, 0x3a, 0x15, 0xae, 0x83, 0x06, 0x4f, 0xb7, 0x17, 0x8b, 0x78, 0xfa, 0x3d, 0xf2, 0x9c, 0x59,
	0x8a, 0x65, 0xca, 0x22, 0x0f, 0x12, 0xdc, 0xcb, 0x94, 0xc5, 0xc3, 0xca, 0xb0, 0xaf, 0xc7, 0x3f,
	0x25, 0x67, 0xe2, 0x01, 0x0b, 0x91, 0x1d, 0x3c, 0xe1, 0x41, 0x38, 0x1b, 0x0f, 0x58, 0x0c, 0xed,
	0x8e, 0x6b, 0xa9, 0x32, 0xcc, 0x9d, 0x71, 0x07, 0x2c, 0x44, 0x64, 0x06, 0x7d, 0xdf, 0x00, 0x4d,
	0xfb, 0xce, 0x0e, 0xb7, 0xfb, 0x13, 0xe1, 0xf1, 0x7b, 0xe8, 0xda, 0x2a, 0x8f, 0x4e, 0xeb, 0xa8,
	0x71, 0x7b, 0x4f, 0xe3, 0x67, 0x90, 0x18, 0xae, 0x96, 0x68, 0xc2, 0x4d, 0xee, 0x34, 0xde, 0xa3,
	0xd3, 0x7f, 0x5b, 0x90, 0xf8, 0x94, 0x3d, 0x66, 0xc5, 0xf5, 0x2a, 0x1e, 0x6d, 0xd7, 0x7b, 0x83,
	0xa9, 0x7d, 0x6a, 0x30, 0x51, 0xe8, 0x97, 0xa8, 0x35, 0x5f, 0xc6, 0x71, 0x1c, 0x43, 0x92, 0x42,
	0x37, 0x8f, 0xe3, 0xee, 0xb4, 0x23, 0x1c, 0x6f, 0xd7, 0x80, 0xde, 0x89, 0x06, 0x1c, 0xbe, 0xd6,
	0xe4, 0x21, 0xaf, 0x75, 0x7a, 0x09, 0x7d, 0x7f, 0x4f, 0x6d, 0xb5, 0xb9, 0xe1, 0x45, 0x83, 0xf6,
	0x73, 0x74, 0x4c, 0xf4, 0x80, 0x4e, 0xdf, 0xc0, 0xb9, 0x3d, 0xcc, 0x0f, 0xa0, 0xef, 0x44, 0xe1,
	0x95, 0x96, 0xf9, 0xee, 0x35, 0xcb, 0xdc, 0xe5, 0x6a, 0x6e, 0x56, 0x51, 0x7d, 0xbb, 0xb6, 0xb9,
	0x9c, 0x1b, 0x1e, 0xc4, 0x70, 0xeb, 0xe9, 0x4f, 0x90, 0xbc, 0xf4, 0x6a, 0x3d, 0x81, 0x8e, 0xe2,
	0xeb, 0x70, 0x88, 0x5d, 0x6e, 0xbb, 0xda, 0xde, 0xeb, 0xea, 0xa7, 0xd0, 0xb5, 0xdf, 0xdc, 0xdb,
	0x2f, 0xd1, 0x0b, 0xef, 0xa0, 0x85, 0xff, 0xf4, 0x7f, 0xfd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x9f, 0xe8, 0xde, 0x9a, 0x1e, 0x08, 0x00, 0x00,
}
