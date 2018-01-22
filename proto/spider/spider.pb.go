// Code generated by protoc-gen-go.
// source: github.com/roscopecoltran/gcse/proto/spider/spider.proto
// DO NOT EDIT!

/*
Package sppb is a generated protocol buffer package.

It is generated from these files:
	github.com/roscopecoltran/gcse/proto/spider/spider.proto

It has these top-level messages:
	GoFileInfo
	RepoInfo
	FolderInfo
	CrawlingInfo
	HistoryEvent
	HistoryInfo
	Package
*/
package sppb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type GoFileInfo_Status int32

const (
	GoFileInfo_Unknown      GoFileInfo_Status = 0
	GoFileInfo_ParseSuccess GoFileInfo_Status = 1
	GoFileInfo_ParseFailed  GoFileInfo_Status = 2
	GoFileInfo_ShouldIgnore GoFileInfo_Status = 3
)

var GoFileInfo_Status_name = map[int32]string{
	0: "Unknown",
	1: "ParseSuccess",
	2: "ParseFailed",
	3: "ShouldIgnore",
}
var GoFileInfo_Status_value = map[string]int32{
	"Unknown":      0,
	"ParseSuccess": 1,
	"ParseFailed":  2,
	"ShouldIgnore": 3,
}

func (x GoFileInfo_Status) String() string {
	return proto.EnumName(GoFileInfo_Status_name, int32(x))
}
func (GoFileInfo_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type HistoryEvent_Action_Enum int32

const (
	HistoryEvent_Action_None    HistoryEvent_Action_Enum = 0
	HistoryEvent_Action_Success HistoryEvent_Action_Enum = 1
	HistoryEvent_Action_Failed  HistoryEvent_Action_Enum = 2
	HistoryEvent_Action_Invalid HistoryEvent_Action_Enum = 3
)

var HistoryEvent_Action_Enum_name = map[int32]string{
	0: "None",
	1: "Success",
	2: "Failed",
	3: "Invalid",
}
var HistoryEvent_Action_Enum_value = map[string]int32{
	"None":    0,
	"Success": 1,
	"Failed":  2,
	"Invalid": 3,
}

func (x HistoryEvent_Action_Enum) String() string {
	return proto.EnumName(HistoryEvent_Action_Enum_name, int32(x))
}
func (HistoryEvent_Action_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0, 0}
}

type GoFileInfo struct {
	Status      GoFileInfo_Status `protobuf:"varint,1,opt,name=status,enum=gcse.spider.GoFileInfo_Status" json:"status,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	IsTest      bool              `protobuf:"varint,4,opt,name=is_test,json=isTest" json:"is_test,omitempty"`
	Imports     []string          `protobuf:"bytes,5,rep,name=imports" json:"imports,omitempty"`
}

func (m *GoFileInfo) Reset()                    { *m = GoFileInfo{} }
func (m *GoFileInfo) String() string            { return proto.CompactTextString(m) }
func (*GoFileInfo) ProtoMessage()               {}
func (*GoFileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RepoInfo struct {
	// The timestamp this repo-info is crawled
	CrawlingTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=crawling_time,json=crawlingTime" json:"crawling_time,omitempty"`
	Stars        int32                      `protobuf:"varint,2,opt,name=stars" json:"stars,omitempty"`
	Description  string                     `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// Where this project was forked from, full path
	Source string `protobuf:"bytes,5,opt,name=source" json:"source,omitempty"`
	// As far as we know, when this repo was updated
	LastUpdated *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
}

func (m *RepoInfo) Reset()                    { *m = RepoInfo{} }
func (m *RepoInfo) String() string            { return proto.CompactTextString(m) }
func (*RepoInfo) ProtoMessage()               {}
func (*RepoInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RepoInfo) GetCrawlingTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CrawlingTime
	}
	return nil
}

func (m *RepoInfo) GetLastUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Information for a non-repository folder.
type FolderInfo struct {
	// E.g. "sub"
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// E.g. "spider/sub"
	Path    string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Sha     string `protobuf:"bytes,3,opt,name=sha" json:"sha,omitempty"`
	HtmlUrl string `protobuf:"bytes,4,opt,name=html_url,json=htmlUrl" json:"html_url,omitempty"`
	// The timestamp this folder-info is crawled
	CrawlingTime *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=crawling_time,json=crawlingTime" json:"crawling_time,omitempty"`
}

func (m *FolderInfo) Reset()                    { *m = FolderInfo{} }
func (m *FolderInfo) String() string            { return proto.CompactTextString(m) }
func (*FolderInfo) ProtoMessage()               {}
func (*FolderInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FolderInfo) GetCrawlingTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CrawlingTime
	}
	return nil
}

type CrawlingInfo struct {
	// The timestamp the related entry was crawled
	CrawlingTime *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=crawling_time,json=crawlingTime" json:"crawling_time,omitempty"`
	Etag         string                     `protobuf:"bytes,2,opt,name=etag" json:"etag,omitempty"`
}

func (m *CrawlingInfo) Reset()                    { *m = CrawlingInfo{} }
func (m *CrawlingInfo) String() string            { return proto.CompactTextString(m) }
func (*CrawlingInfo) ProtoMessage()               {}
func (*CrawlingInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CrawlingInfo) GetCrawlingTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CrawlingTime
	}
	return nil
}

type HistoryEvent struct {
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Action    HistoryEvent_Action_Enum   `protobuf:"varint,2,opt,name=action,enum=gcse.spider.HistoryEvent_Action_Enum" json:"action,omitempty"`
}

func (m *HistoryEvent) Reset()                    { *m = HistoryEvent{} }
func (m *HistoryEvent) String() string            { return proto.CompactTextString(m) }
func (*HistoryEvent) ProtoMessage()               {}
func (*HistoryEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *HistoryEvent) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type HistoryEvent_Action struct {
}

func (m *HistoryEvent_Action) Reset()                    { *m = HistoryEvent_Action{} }
func (m *HistoryEvent_Action) String() string            { return proto.CompactTextString(m) }
func (*HistoryEvent_Action) ProtoMessage()               {}
func (*HistoryEvent_Action) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type HistoryInfo struct {
	Events    []*HistoryEvent            `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
	FoundTime *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=found_time,json=foundTime" json:"found_time,omitempty"`
	// Possible value:
	//   web                  added from web
	//   user:<user>          found from user crawling
	//   parent               found by crawling his parent
	//   imported:<pkg>       imported by a <pkg>
	//   testimported:<pkg>   test imported by a <pkg>
	//   package:<pkg>
	//   reference:<pkg>      referenced in the readme file of <pkg>
	//   godoc                found by godoc.org/api
	FoundWay      string                     `protobuf:"bytes,3,opt,name=found_way,json=foundWay" json:"found_way,omitempty"`
	LatestSuccess *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=latest_success,json=latestSuccess" json:"latest_success,omitempty"`
	LatestFailed  *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=latest_failed,json=latestFailed" json:"latest_failed,omitempty"`
}

func (m *HistoryInfo) Reset()                    { *m = HistoryInfo{} }
func (m *HistoryInfo) String() string            { return proto.CompactTextString(m) }
func (*HistoryInfo) ProtoMessage()               {}
func (*HistoryInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *HistoryInfo) GetEvents() []*HistoryEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *HistoryInfo) GetFoundTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.FoundTime
	}
	return nil
}

func (m *HistoryInfo) GetLatestSuccess() *google_protobuf.Timestamp {
	if m != nil {
		return m.LatestSuccess
	}
	return nil
}

func (m *HistoryInfo) GetLatestFailed() *google_protobuf.Timestamp {
	if m != nil {
		return m.LatestFailed
	}
	return nil
}

type Package struct {
	// package "name"
	Name string `protobuf:"bytes,1,opt,name=Name,json=name" json:"Name,omitempty"`
	// Relative path to the repository, "" for root repository, "/sub" for a sub package.
	// Full path: site + "/" + user + "/" + repo + path
	Path        string `protobuf:"bytes,2,opt,name=Path,json=path" json:"Path,omitempty"`
	Synopsis    string `protobuf:"bytes,9,opt,name=Synopsis,json=synopsis" json:"Synopsis,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,json=description" json:"Description,omitempty"`
	// No directory info
	ReadmeFn string `protobuf:"bytes,4,opt,name=ReadmeFn,json=readmeFn" json:"ReadmeFn,omitempty"`
	// Raw content, cound be md, txt, etc.
	ReadmeData  string   `protobuf:"bytes,5,opt,name=ReadmeData,json=readmeData" json:"ReadmeData,omitempty"`
	Imports     []string `protobuf:"bytes,6,rep,name=Imports,json=imports" json:"Imports,omitempty"`
	TestImports []string `protobuf:"bytes,7,rep,name=TestImports,json=testImports" json:"TestImports,omitempty"`
	// URL to the package source code.
	Url string `protobuf:"bytes,8,opt,name=url" json:"url,omitempty"`
}

func (m *Package) Reset()                    { *m = Package{} }
func (m *Package) String() string            { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()               {}
func (*Package) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*GoFileInfo)(nil), "gcse.spider.GoFileInfo")
	proto.RegisterType((*RepoInfo)(nil), "gcse.spider.RepoInfo")
	proto.RegisterType((*FolderInfo)(nil), "gcse.spider.FolderInfo")
	proto.RegisterType((*CrawlingInfo)(nil), "gcse.spider.CrawlingInfo")
	proto.RegisterType((*HistoryEvent)(nil), "gcse.spider.HistoryEvent")
	proto.RegisterType((*HistoryEvent_Action)(nil), "gcse.spider.HistoryEvent.Action")
	proto.RegisterType((*HistoryInfo)(nil), "gcse.spider.HistoryInfo")
	proto.RegisterType((*Package)(nil), "gcse.spider.Package")
	proto.RegisterEnum("gcse.spider.GoFileInfo_Status", GoFileInfo_Status_name, GoFileInfo_Status_value)
	proto.RegisterEnum("gcse.spider.HistoryEvent_Action_Enum", HistoryEvent_Action_Enum_name, HistoryEvent_Action_Enum_value)
}

var fileDescriptor0 = []byte{
	// 738 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x3e, 0xf9, 0x73, 0xdc, 0x71, 0xda, 0x13, 0xad, 0x8e, 0xce, 0x49, 0x73, 0xa4, 0xaa, 0x8a,
	0x84, 0xc4, 0x95, 0x23, 0x8a, 0xa8, 0x40, 0x08, 0x41, 0xa1, 0x0d, 0x84, 0x8b, 0xaa, 0x72, 0x5a,
	0x21, 0x71, 0x13, 0x6d, 0xec, 0x8d, 0x63, 0xd5, 0xde, 0xb5, 0xbc, 0xeb, 0x56, 0x79, 0x17, 0xae,
	0x78, 0x1b, 0x5e, 0x81, 0x37, 0xe1, 0x8e, 0xfd, 0x73, 0x93, 0xf2, 0x53, 0x82, 0xc4, 0x95, 0x77,
	0xfe, 0x3c, 0x33, 0xdf, 0xcc, 0x37, 0xf0, 0x28, 0x4e, 0xc4, 0xa2, 0x9c, 0xf9, 0x21, 0xcb, 0x86,
	0x11, 0xbe, 0x4a, 0xa2, 0x88, 0xd0, 0x38, 0xa4, 0xc3, 0x38, 0xe4, 0x64, 0x98, 0x17, 0x4c, 0xb0,
	0x21, 0xcf, 0x93, 0x88, 0x14, 0xf6, 0xe3, 0x6b, 0x1d, 0xf2, 0x94, 0xdd, 0x37, 0xaa, 0xfe, 0xd3,
	0xb5, 0x7f, 0xc4, 0x2c, 0xc5, 0x34, 0x36, 0x91, 0xb3, 0x72, 0x3e, 0xcc, 0xc5, 0x32, 0x27, 0x7c,
	0x28, 0x92, 0x8c, 0x70, 0x81, 0xb3, 0x7c, 0xf5, 0x32, 0x7f, 0x1a, 0x7c, 0xa9, 0x01, 0xbc, 0x66,
	0xa3, 0x24, 0x25, 0x63, 0x3a, 0x67, 0xe8, 0x10, 0x1c, 0x69, 0x15, 0x25, 0xef, 0xd5, 0xf6, 0x6b,
	0xf7, 0x77, 0x0e, 0xf6, 0xfc, 0xb5, 0x4c, 0xfe, 0xca, 0xd1, 0x9f, 0x68, 0xaf, 0xc0, 0x7a, 0x23,
	0x04, 0x4d, 0x8a, 0x33, 0xd2, 0xab, 0xcb, 0xa8, 0xad, 0x40, 0xbf, 0xd1, 0x3e, 0x78, 0x11, 0xe1,
	0x61, 0x91, 0xe4, 0x22, 0x61, 0xb4, 0xd7, 0xd0, 0xa6, 0x75, 0x15, 0xfa, 0x0f, 0xda, 0x09, 0x9f,
	0x0a, 0x59, 0x50, 0xaf, 0x29, 0xad, 0x6e, 0xe0, 0x24, 0xfc, 0x5c, 0x4a, 0xa8, 0x27, 0x0d, 0x59,
	0xce, 0x0a, 0xc1, 0x7b, 0xad, 0xfd, 0x86, 0x0c, 0xab, 0xc4, 0xc1, 0x5b, 0x70, 0x4c, 0x6a, 0xe4,
	0x41, 0xfb, 0x82, 0x5e, 0x52, 0x76, 0x4d, 0xbb, 0x7f, 0xa1, 0x2e, 0x74, 0xce, 0x70, 0xc1, 0xc9,
	0xa4, 0x0c, 0x43, 0xc2, 0x79, 0xb7, 0x86, 0xfe, 0x06, 0x4f, 0x6b, 0x46, 0x58, 0x96, 0x1c, 0x75,
	0xeb, 0xca, 0x65, 0xb2, 0x60, 0x65, 0x1a, 0x8d, 0x63, 0xca, 0x0a, 0xd2, 0x6d, 0x0c, 0x3e, 0xd7,
	0xc0, 0x0d, 0x48, 0xce, 0x74, 0xe7, 0xcf, 0x61, 0x3b, 0x2c, 0xf0, 0x75, 0x9a, 0xd0, 0x78, 0xaa,
	0x40, 0xd2, 0x00, 0x78, 0x07, 0x7d, 0x3f, 0x66, 0x2c, 0x4e, 0x89, 0x5f, 0x41, 0xea, 0x9f, 0x57,
	0x08, 0x06, 0x9d, 0x2a, 0x40, 0xa9, 0xd0, 0x3f, 0xd0, 0x92, 0xea, 0x82, 0x6b, 0x0c, 0x5a, 0x81,
	0x11, 0x36, 0x00, 0xe1, 0x5f, 0x09, 0x39, 0x2b, 0x8b, 0x90, 0xc8, 0x56, 0x95, 0xd1, 0x4a, 0xe8,
	0x19, 0x74, 0x52, 0xcc, 0xc5, 0xb4, 0xcc, 0x23, 0x2c, 0x48, 0xa4, 0x11, 0xba, 0xbb, 0x1e, 0x4f,
	0xf9, 0x5f, 0x18, 0xf7, 0xc1, 0x47, 0x39, 0xd8, 0x11, 0x4b, 0xe5, 0xd8, 0x74, 0x7b, 0xd5, 0x80,
	0x6a, 0x6b, 0x03, 0x92, 0xba, 0x1c, 0x8b, 0x45, 0x35, 0x34, 0xf5, 0x96, 0x28, 0x35, 0xf8, 0x02,
	0xdb, 0x3a, 0xd5, 0x13, 0xed, 0x82, 0xbb, 0x10, 0x59, 0x3a, 0x2d, 0x8b, 0x54, 0xd7, 0x20, 0x87,
	0xa1, 0xe4, 0x8b, 0x22, 0xfd, 0x1e, 0xb3, 0xd6, 0xef, 0x61, 0x36, 0x08, 0xa1, 0xf3, 0xca, 0xca,
	0x7f, 0x66, 0x08, 0xb2, 0x25, 0x22, 0x70, 0x5c, 0xb5, 0xa4, 0xde, 0x83, 0x4f, 0x35, 0xe8, 0xbc,
	0x49, 0xb8, 0x60, 0xc5, 0xf2, 0xe4, 0x8a, 0x50, 0x81, 0x1e, 0xc3, 0xd6, 0x0d, 0x0d, 0x36, 0xc8,
	0xb0, 0x72, 0x96, 0x33, 0x71, 0x70, 0xa8, 0x07, 0x59, 0xd7, 0xf4, 0xb8, 0x77, 0x8b, 0x1e, 0xeb,
	0x49, 0xfc, 0x23, 0xed, 0xe7, 0x9f, 0xd0, 0x32, 0x0b, 0x6c, 0x50, 0xff, 0x05, 0x38, 0x46, 0x3d,
	0x38, 0x84, 0xa6, 0xb2, 0x20, 0x17, 0x9a, 0xa7, 0x8c, 0x12, 0xb9, 0xc1, 0x72, 0x9d, 0x57, 0xcb,
	0x0b, 0xe0, 0xdc, 0xec, 0xad, 0x34, 0x8c, 0xe9, 0x15, 0x4e, 0x93, 0x48, 0xae, 0xec, 0x87, 0x3a,
	0x78, 0x36, 0x8d, 0x06, 0xec, 0x01, 0x38, 0x44, 0xa5, 0x53, 0x7c, 0x6d, 0xc8, 0x3e, 0x76, 0x7f,
	0x5a, 0x50, 0x60, 0x1d, 0xd1, 0x13, 0x80, 0x39, 0x2b, 0x69, 0x64, 0x00, 0xae, 0xff, 0xba, 0x7d,
	0xed, 0xad, 0xd1, 0xfd, 0x1f, 0x8c, 0x30, 0xbd, 0xc6, 0x4b, 0xbb, 0x22, 0xae, 0x56, 0xbc, 0xc3,
	0x4b, 0x74, 0x04, 0x3b, 0x29, 0x56, 0x5c, 0x9e, 0x72, 0xd3, 0xc7, 0x06, 0x1b, 0xbb, 0x6d, 0x22,
	0x6c, 0xe3, 0x6a, 0xfc, 0xf6, 0x17, 0x73, 0xdd, 0xfd, 0x26, 0xfb, 0x64, 0x02, 0x0c, 0x5a, 0xea,
	0x9a, 0xb5, 0xcf, 0x70, 0x78, 0x89, 0x63, 0xbd, 0x0a, 0xa7, 0x3f, 0xd8, 0xf8, 0xb3, 0x6f, 0x37,
	0xbe, 0x0f, 0xee, 0x64, 0x49, 0x59, 0xce, 0x13, 0xde, 0xdb, 0x32, 0x3d, 0x71, 0x2b, 0x2b, 0xf6,
	0x1e, 0xdf, 0xcd, 0xde, 0xbe, 0x3a, 0x21, 0x38, 0xca, 0xc8, 0x88, 0x5a, 0x76, 0xb8, 0x85, 0x95,
	0xd1, 0x1e, 0x80, 0xb1, 0x1d, 0x63, 0x81, 0x2d, 0xbb, 0xa1, 0xb8, 0xd1, 0xa8, 0x2b, 0x37, 0xb6,
	0x57, 0xce, 0xb9, 0x75, 0xe5, 0x54, 0x5e, 0x75, 0x07, 0x2b, 0x6b, 0x5b, 0x5b, 0x3d, 0xb1, 0x52,
	0x29, 0x9e, 0x2a, 0x42, 0xba, 0x86, 0xa7, 0xf2, 0xf9, 0xd2, 0x79, 0xdf, 0xe4, 0x79, 0x3e, 0x9b,
	0x39, 0x1a, 0xa5, 0x87, 0x5f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xee, 0xec, 0x45, 0x2d, 0x5b, 0x06,
	0x00, 0x00,
}
