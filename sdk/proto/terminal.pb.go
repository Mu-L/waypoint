// Code generated by protoc-gen-go. DO NOT EDIT.
// source: terminal.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type TerminalUI struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminalUI) Reset()         { *m = TerminalUI{} }
func (m *TerminalUI) String() string { return proto.CompactTextString(m) }
func (*TerminalUI) ProtoMessage()    {}
func (*TerminalUI) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{0}
}

func (m *TerminalUI) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminalUI.Unmarshal(m, b)
}
func (m *TerminalUI) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminalUI.Marshal(b, m, deterministic)
}
func (m *TerminalUI) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminalUI.Merge(m, src)
}
func (m *TerminalUI) XXX_Size() int {
	return xxx_messageInfo_TerminalUI.Size(m)
}
func (m *TerminalUI) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminalUI.DiscardUnknown(m)
}

var xxx_messageInfo_TerminalUI proto.InternalMessageInfo

type TerminalUI_OutputRequest struct {
	Lines                []string `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminalUI_OutputRequest) Reset()         { *m = TerminalUI_OutputRequest{} }
func (m *TerminalUI_OutputRequest) String() string { return proto.CompactTextString(m) }
func (*TerminalUI_OutputRequest) ProtoMessage()    {}
func (*TerminalUI_OutputRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{0, 0}
}

func (m *TerminalUI_OutputRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminalUI_OutputRequest.Unmarshal(m, b)
}
func (m *TerminalUI_OutputRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminalUI_OutputRequest.Marshal(b, m, deterministic)
}
func (m *TerminalUI_OutputRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminalUI_OutputRequest.Merge(m, src)
}
func (m *TerminalUI_OutputRequest) XXX_Size() int {
	return xxx_messageInfo_TerminalUI_OutputRequest.Size(m)
}
func (m *TerminalUI_OutputRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminalUI_OutputRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TerminalUI_OutputRequest proto.InternalMessageInfo

func (m *TerminalUI_OutputRequest) GetLines() []string {
	if m != nil {
		return m.Lines
	}
	return nil
}

type TerminalUI_Event struct {
	// Types that are valid to be assigned to Event:
	//	*TerminalUI_Event_Line_
	//	*TerminalUI_Event_Status_
	//	*TerminalUI_Event_NamedValues_
	//	*TerminalUI_Event_Raw_
	//	*TerminalUI_Event_Table_
	Event                isTerminalUI_Event_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TerminalUI_Event) Reset()         { *m = TerminalUI_Event{} }
func (m *TerminalUI_Event) String() string { return proto.CompactTextString(m) }
func (*TerminalUI_Event) ProtoMessage()    {}
func (*TerminalUI_Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{0, 1}
}

func (m *TerminalUI_Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminalUI_Event.Unmarshal(m, b)
}
func (m *TerminalUI_Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminalUI_Event.Marshal(b, m, deterministic)
}
func (m *TerminalUI_Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminalUI_Event.Merge(m, src)
}
func (m *TerminalUI_Event) XXX_Size() int {
	return xxx_messageInfo_TerminalUI_Event.Size(m)
}
func (m *TerminalUI_Event) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminalUI_Event.DiscardUnknown(m)
}

var xxx_messageInfo_TerminalUI_Event proto.InternalMessageInfo

type isTerminalUI_Event_Event interface {
	isTerminalUI_Event_Event()
}

type TerminalUI_Event_Line_ struct {
	Line *TerminalUI_Event_Line `protobuf:"bytes,1,opt,name=line,proto3,oneof"`
}

type TerminalUI_Event_Status_ struct {
	Status *TerminalUI_Event_Status `protobuf:"bytes,2,opt,name=status,proto3,oneof"`
}

type TerminalUI_Event_NamedValues_ struct {
	NamedValues *TerminalUI_Event_NamedValues `protobuf:"bytes,3,opt,name=named_values,json=namedValues,proto3,oneof"`
}

type TerminalUI_Event_Raw_ struct {
	Raw *TerminalUI_Event_Raw `protobuf:"bytes,4,opt,name=raw,proto3,oneof"`
}

type TerminalUI_Event_Table_ struct {
	Table *TerminalUI_Event_Table `protobuf:"bytes,5,opt,name=table,proto3,oneof"`
}

func (*TerminalUI_Event_Line_) isTerminalUI_Event_Event() {}

func (*TerminalUI_Event_Status_) isTerminalUI_Event_Event() {}

func (*TerminalUI_Event_NamedValues_) isTerminalUI_Event_Event() {}

func (*TerminalUI_Event_Raw_) isTerminalUI_Event_Event() {}

func (*TerminalUI_Event_Table_) isTerminalUI_Event_Event() {}

func (m *TerminalUI_Event) GetEvent() isTerminalUI_Event_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *TerminalUI_Event) GetLine() *TerminalUI_Event_Line {
	if x, ok := m.GetEvent().(*TerminalUI_Event_Line_); ok {
		return x.Line
	}
	return nil
}

func (m *TerminalUI_Event) GetStatus() *TerminalUI_Event_Status {
	if x, ok := m.GetEvent().(*TerminalUI_Event_Status_); ok {
		return x.Status
	}
	return nil
}

func (m *TerminalUI_Event) GetNamedValues() *TerminalUI_Event_NamedValues {
	if x, ok := m.GetEvent().(*TerminalUI_Event_NamedValues_); ok {
		return x.NamedValues
	}
	return nil
}

func (m *TerminalUI_Event) GetRaw() *TerminalUI_Event_Raw {
	if x, ok := m.GetEvent().(*TerminalUI_Event_Raw_); ok {
		return x.Raw
	}
	return nil
}

func (m *TerminalUI_Event) GetTable() *TerminalUI_Event_Table {
	if x, ok := m.GetEvent().(*TerminalUI_Event_Table_); ok {
		return x.Table
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TerminalUI_Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TerminalUI_Event_Line_)(nil),
		(*TerminalUI_Event_Status_)(nil),
		(*TerminalUI_Event_NamedValues_)(nil),
		(*TerminalUI_Event_Raw_)(nil),
		(*TerminalUI_Event_Table_)(nil),
	}
}

type TerminalUI_Event_Status struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Step                 bool     `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminalUI_Event_Status) Reset()         { *m = TerminalUI_Event_Status{} }
func (m *TerminalUI_Event_Status) String() string { return proto.CompactTextString(m) }
func (*TerminalUI_Event_Status) ProtoMessage()    {}
func (*TerminalUI_Event_Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{0, 1, 0}
}

func (m *TerminalUI_Event_Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminalUI_Event_Status.Unmarshal(m, b)
}
func (m *TerminalUI_Event_Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminalUI_Event_Status.Marshal(b, m, deterministic)
}
func (m *TerminalUI_Event_Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminalUI_Event_Status.Merge(m, src)
}
func (m *TerminalUI_Event_Status) XXX_Size() int {
	return xxx_messageInfo_TerminalUI_Event_Status.Size(m)
}
func (m *TerminalUI_Event_Status) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminalUI_Event_Status.DiscardUnknown(m)
}

var xxx_messageInfo_TerminalUI_Event_Status proto.InternalMessageInfo

func (m *TerminalUI_Event_Status) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TerminalUI_Event_Status) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *TerminalUI_Event_Status) GetStep() bool {
	if m != nil {
		return m.Step
	}
	return false
}

type TerminalUI_Event_Line struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Style                string   `protobuf:"bytes,2,opt,name=style,proto3" json:"style,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminalUI_Event_Line) Reset()         { *m = TerminalUI_Event_Line{} }
func (m *TerminalUI_Event_Line) String() string { return proto.CompactTextString(m) }
func (*TerminalUI_Event_Line) ProtoMessage()    {}
func (*TerminalUI_Event_Line) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{0, 1, 1}
}

func (m *TerminalUI_Event_Line) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminalUI_Event_Line.Unmarshal(m, b)
}
func (m *TerminalUI_Event_Line) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminalUI_Event_Line.Marshal(b, m, deterministic)
}
func (m *TerminalUI_Event_Line) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminalUI_Event_Line.Merge(m, src)
}
func (m *TerminalUI_Event_Line) XXX_Size() int {
	return xxx_messageInfo_TerminalUI_Event_Line.Size(m)
}
func (m *TerminalUI_Event_Line) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminalUI_Event_Line.DiscardUnknown(m)
}

var xxx_messageInfo_TerminalUI_Event_Line proto.InternalMessageInfo

func (m *TerminalUI_Event_Line) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *TerminalUI_Event_Line) GetStyle() string {
	if m != nil {
		return m.Style
	}
	return ""
}

type TerminalUI_Event_Raw struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Stderr               bool     `protobuf:"varint,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminalUI_Event_Raw) Reset()         { *m = TerminalUI_Event_Raw{} }
func (m *TerminalUI_Event_Raw) String() string { return proto.CompactTextString(m) }
func (*TerminalUI_Event_Raw) ProtoMessage()    {}
func (*TerminalUI_Event_Raw) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{0, 1, 2}
}

func (m *TerminalUI_Event_Raw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminalUI_Event_Raw.Unmarshal(m, b)
}
func (m *TerminalUI_Event_Raw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminalUI_Event_Raw.Marshal(b, m, deterministic)
}
func (m *TerminalUI_Event_Raw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminalUI_Event_Raw.Merge(m, src)
}
func (m *TerminalUI_Event_Raw) XXX_Size() int {
	return xxx_messageInfo_TerminalUI_Event_Raw.Size(m)
}
func (m *TerminalUI_Event_Raw) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminalUI_Event_Raw.DiscardUnknown(m)
}

var xxx_messageInfo_TerminalUI_Event_Raw proto.InternalMessageInfo

func (m *TerminalUI_Event_Raw) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TerminalUI_Event_Raw) GetStderr() bool {
	if m != nil {
		return m.Stderr
	}
	return false
}

type TerminalUI_Event_NamedValue struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminalUI_Event_NamedValue) Reset()         { *m = TerminalUI_Event_NamedValue{} }
func (m *TerminalUI_Event_NamedValue) String() string { return proto.CompactTextString(m) }
func (*TerminalUI_Event_NamedValue) ProtoMessage()    {}
func (*TerminalUI_Event_NamedValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{0, 1, 3}
}

func (m *TerminalUI_Event_NamedValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminalUI_Event_NamedValue.Unmarshal(m, b)
}
func (m *TerminalUI_Event_NamedValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminalUI_Event_NamedValue.Marshal(b, m, deterministic)
}
func (m *TerminalUI_Event_NamedValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminalUI_Event_NamedValue.Merge(m, src)
}
func (m *TerminalUI_Event_NamedValue) XXX_Size() int {
	return xxx_messageInfo_TerminalUI_Event_NamedValue.Size(m)
}
func (m *TerminalUI_Event_NamedValue) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminalUI_Event_NamedValue.DiscardUnknown(m)
}

var xxx_messageInfo_TerminalUI_Event_NamedValue proto.InternalMessageInfo

func (m *TerminalUI_Event_NamedValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TerminalUI_Event_NamedValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type TerminalUI_Event_NamedValues struct {
	Values               []*TerminalUI_Event_NamedValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TerminalUI_Event_NamedValues) Reset()         { *m = TerminalUI_Event_NamedValues{} }
func (m *TerminalUI_Event_NamedValues) String() string { return proto.CompactTextString(m) }
func (*TerminalUI_Event_NamedValues) ProtoMessage()    {}
func (*TerminalUI_Event_NamedValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{0, 1, 4}
}

func (m *TerminalUI_Event_NamedValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminalUI_Event_NamedValues.Unmarshal(m, b)
}
func (m *TerminalUI_Event_NamedValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminalUI_Event_NamedValues.Marshal(b, m, deterministic)
}
func (m *TerminalUI_Event_NamedValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminalUI_Event_NamedValues.Merge(m, src)
}
func (m *TerminalUI_Event_NamedValues) XXX_Size() int {
	return xxx_messageInfo_TerminalUI_Event_NamedValues.Size(m)
}
func (m *TerminalUI_Event_NamedValues) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminalUI_Event_NamedValues.DiscardUnknown(m)
}

var xxx_messageInfo_TerminalUI_Event_NamedValues proto.InternalMessageInfo

func (m *TerminalUI_Event_NamedValues) GetValues() []*TerminalUI_Event_NamedValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type TerminalUI_Event_TableEntry struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Color                string   `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TerminalUI_Event_TableEntry) Reset()         { *m = TerminalUI_Event_TableEntry{} }
func (m *TerminalUI_Event_TableEntry) String() string { return proto.CompactTextString(m) }
func (*TerminalUI_Event_TableEntry) ProtoMessage()    {}
func (*TerminalUI_Event_TableEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{0, 1, 5}
}

func (m *TerminalUI_Event_TableEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminalUI_Event_TableEntry.Unmarshal(m, b)
}
func (m *TerminalUI_Event_TableEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminalUI_Event_TableEntry.Marshal(b, m, deterministic)
}
func (m *TerminalUI_Event_TableEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminalUI_Event_TableEntry.Merge(m, src)
}
func (m *TerminalUI_Event_TableEntry) XXX_Size() int {
	return xxx_messageInfo_TerminalUI_Event_TableEntry.Size(m)
}
func (m *TerminalUI_Event_TableEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminalUI_Event_TableEntry.DiscardUnknown(m)
}

var xxx_messageInfo_TerminalUI_Event_TableEntry proto.InternalMessageInfo

func (m *TerminalUI_Event_TableEntry) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *TerminalUI_Event_TableEntry) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

type TerminalUI_Event_TableRow struct {
	Entries              []*TerminalUI_Event_TableEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TerminalUI_Event_TableRow) Reset()         { *m = TerminalUI_Event_TableRow{} }
func (m *TerminalUI_Event_TableRow) String() string { return proto.CompactTextString(m) }
func (*TerminalUI_Event_TableRow) ProtoMessage()    {}
func (*TerminalUI_Event_TableRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{0, 1, 6}
}

func (m *TerminalUI_Event_TableRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminalUI_Event_TableRow.Unmarshal(m, b)
}
func (m *TerminalUI_Event_TableRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminalUI_Event_TableRow.Marshal(b, m, deterministic)
}
func (m *TerminalUI_Event_TableRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminalUI_Event_TableRow.Merge(m, src)
}
func (m *TerminalUI_Event_TableRow) XXX_Size() int {
	return xxx_messageInfo_TerminalUI_Event_TableRow.Size(m)
}
func (m *TerminalUI_Event_TableRow) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminalUI_Event_TableRow.DiscardUnknown(m)
}

var xxx_messageInfo_TerminalUI_Event_TableRow proto.InternalMessageInfo

func (m *TerminalUI_Event_TableRow) GetEntries() []*TerminalUI_Event_TableEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type TerminalUI_Event_Table struct {
	Headers              []string                     `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	Rows                 []*TerminalUI_Event_TableRow `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *TerminalUI_Event_Table) Reset()         { *m = TerminalUI_Event_Table{} }
func (m *TerminalUI_Event_Table) String() string { return proto.CompactTextString(m) }
func (*TerminalUI_Event_Table) ProtoMessage()    {}
func (*TerminalUI_Event_Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8b8260c8ef16ad, []int{0, 1, 7}
}

func (m *TerminalUI_Event_Table) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TerminalUI_Event_Table.Unmarshal(m, b)
}
func (m *TerminalUI_Event_Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TerminalUI_Event_Table.Marshal(b, m, deterministic)
}
func (m *TerminalUI_Event_Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TerminalUI_Event_Table.Merge(m, src)
}
func (m *TerminalUI_Event_Table) XXX_Size() int {
	return xxx_messageInfo_TerminalUI_Event_Table.Size(m)
}
func (m *TerminalUI_Event_Table) XXX_DiscardUnknown() {
	xxx_messageInfo_TerminalUI_Event_Table.DiscardUnknown(m)
}

var xxx_messageInfo_TerminalUI_Event_Table proto.InternalMessageInfo

func (m *TerminalUI_Event_Table) GetHeaders() []string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *TerminalUI_Event_Table) GetRows() []*TerminalUI_Event_TableRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

func init() {
	proto.RegisterType((*TerminalUI)(nil), "proto.TerminalUI")
	proto.RegisterType((*TerminalUI_OutputRequest)(nil), "proto.TerminalUI.OutputRequest")
	proto.RegisterType((*TerminalUI_Event)(nil), "proto.TerminalUI.Event")
	proto.RegisterType((*TerminalUI_Event_Status)(nil), "proto.TerminalUI.Event.Status")
	proto.RegisterType((*TerminalUI_Event_Line)(nil), "proto.TerminalUI.Event.Line")
	proto.RegisterType((*TerminalUI_Event_Raw)(nil), "proto.TerminalUI.Event.Raw")
	proto.RegisterType((*TerminalUI_Event_NamedValue)(nil), "proto.TerminalUI.Event.NamedValue")
	proto.RegisterType((*TerminalUI_Event_NamedValues)(nil), "proto.TerminalUI.Event.NamedValues")
	proto.RegisterType((*TerminalUI_Event_TableEntry)(nil), "proto.TerminalUI.Event.TableEntry")
	proto.RegisterType((*TerminalUI_Event_TableRow)(nil), "proto.TerminalUI.Event.TableRow")
	proto.RegisterType((*TerminalUI_Event_Table)(nil), "proto.TerminalUI.Event.Table")
}

func init() {
	proto.RegisterFile("terminal.proto", fileDescriptor_ff8b8260c8ef16ad)
}

var fileDescriptor_ff8b8260c8ef16ad = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xb5, 0xf1, 0xa5, 0xc9, 0xa4, 0x20, 0x58, 0x55, 0xc5, 0xda, 0x72, 0x89, 0x8a, 0x90, 0xf2,
	0xe4, 0x88, 0x70, 0x51, 0x05, 0xbc, 0x80, 0x14, 0xe4, 0x4a, 0x08, 0xa4, 0x6d, 0x81, 0x47, 0xb4,
	0x69, 0x86, 0x60, 0xc9, 0xb1, 0xc3, 0x7a, 0x1d, 0x2b, 0xdf, 0xc1, 0xe7, 0xf0, 0x55, 0xfc, 0x01,
	0xda, 0x59, 0xbb, 0x2e, 0x42, 0xa6, 0x4f, 0xde, 0x33, 0x3a, 0xe7, 0xcc, 0xce, 0xf1, 0x0e, 0xdc,
	0xd2, 0xa8, 0xd6, 0x69, 0x2e, 0xb3, 0x78, 0xa3, 0x0a, 0x5d, 0xb0, 0x80, 0x3e, 0xfc, 0x68, 0x55,
	0x14, 0xab, 0x0c, 0xa7, 0x84, 0x16, 0xd5, 0xb7, 0x29, 0xae, 0x37, 0x7a, 0x67, 0x39, 0xc7, 0xbf,
	0x43, 0x80, 0xf3, 0x46, 0xf6, 0xe9, 0x94, 0x3f, 0x86, 0x9b, 0x1f, 0x2b, 0xbd, 0xa9, 0xb4, 0xc0,
	0x1f, 0x15, 0x96, 0x9a, 0x1d, 0x40, 0x90, 0xa5, 0x39, 0x96, 0x91, 0x3b, 0xf6, 0x26, 0x43, 0x61,
	0x01, 0xff, 0x15, 0x42, 0x30, 0xdf, 0x62, 0xae, 0xd9, 0x0c, 0x7c, 0x53, 0x8a, 0xdc, 0xb1, 0x3b,
	0x19, 0xcd, 0xee, 0x59, 0xd7, 0xb8, 0x73, 0x8c, 0x89, 0x16, 0xbf, 0x4f, 0x73, 0x4c, 0x1c, 0x41,
	0x5c, 0x76, 0x02, 0x61, 0xa9, 0xa5, 0xae, 0xca, 0xe8, 0x06, 0xa9, 0x1e, 0xf4, 0xa9, 0xce, 0x88,
	0x95, 0x38, 0xa2, 0xe1, 0xb3, 0x04, 0xf6, 0x73, 0xb9, 0xc6, 0xe5, 0xd7, 0xad, 0xcc, 0x2a, 0x2c,
	0x23, 0x8f, 0xf4, 0x8f, 0xfa, 0xf4, 0x1f, 0x0c, 0xf7, 0x33, 0x51, 0x13, 0x47, 0x8c, 0xf2, 0x0e,
	0xb2, 0x29, 0x78, 0x4a, 0xd6, 0x91, 0x4f, 0x06, 0x47, 0x7d, 0x06, 0x42, 0xd6, 0x89, 0x23, 0x0c,
	0x93, 0x3d, 0x87, 0x40, 0xcb, 0x45, 0x86, 0x51, 0x40, 0x92, 0xfb, 0x7d, 0x92, 0x73, 0x43, 0x4a,
	0x1c, 0x61, 0xd9, 0xfc, 0x1d, 0x84, 0x76, 0x0a, 0x76, 0x78, 0x39, 0xb5, 0xc9, 0x6a, 0x78, 0x39,
	0xd3, 0x6d, 0xf0, 0xd6, 0xe5, 0x8a, 0xa2, 0x18, 0x0a, 0x73, 0x64, 0x0c, 0xfc, 0x52, 0xe3, 0x86,
	0xa6, 0x1b, 0x08, 0x3a, 0xf3, 0x18, 0x7c, 0x93, 0x61, 0xcb, 0x76, 0x3b, 0xf6, 0x01, 0x04, 0xa5,
	0xde, 0x65, 0xd8, 0x38, 0x58, 0xc0, 0x9f, 0x80, 0x27, 0x64, 0x6d, 0xac, 0x96, 0x52, 0x4b, 0xe2,
	0xef, 0x0b, 0x3a, 0xdb, 0x8b, 0x2c, 0x51, 0x29, 0x52, 0x0c, 0x44, 0x83, 0xf8, 0x0b, 0x80, 0x2e,
	0x30, 0xa3, 0x34, 0x79, 0x35, 0x9d, 0xe8, 0x6c, 0x5a, 0x51, 0xf0, 0x6d, 0x2b, 0x02, 0xfc, 0x14,
	0x46, 0x57, 0x82, 0x66, 0x2f, 0x21, 0x6c, 0xfe, 0x8e, 0x79, 0x32, 0xa3, 0xd9, 0xf1, 0xf5, 0x7f,
	0x47, 0x34, 0x0a, 0x7e, 0x02, 0x40, 0xf9, 0xcd, 0x73, 0xad, 0x76, 0x5d, 0x3b, 0xf7, 0x4a, 0x3b,
	0x53, 0xbd, 0x28, 0xb2, 0x42, 0xb5, 0x97, 0x20, 0xc0, 0x13, 0x18, 0x90, 0x52, 0x14, 0x35, 0x7b,
	0x0d, 0x7b, 0x98, 0x6b, 0x95, 0x5e, 0x7f, 0x85, 0xae, 0x99, 0x68, 0x25, 0xfc, 0x0b, 0x04, 0x54,
	0x66, 0x11, 0xec, 0x7d, 0x47, 0xb9, 0x44, 0xd5, 0x3e, 0xfe, 0x16, 0xb2, 0x67, 0xe0, 0xab, 0xa2,
	0x36, 0xcf, 0xd7, 0xb8, 0x8f, 0xff, 0xeb, 0x2e, 0x8a, 0x5a, 0x10, 0xfb, 0xed, 0x1e, 0x04, 0x68,
	0xea, 0xb3, 0x9f, 0x2e, 0xdc, 0xe9, 0xc8, 0x67, 0xa8, 0xb6, 0xe9, 0x05, 0xb2, 0x37, 0x10, 0xda,
	0xd5, 0x63, 0x0f, 0xff, 0x35, 0xfc, 0x6b, 0x29, 0xf9, 0x61, 0x6c, 0x57, 0x3a, 0x6e, 0x57, 0x3a,
	0x9e, 0x9b, 0x95, 0x66, 0xaf, 0x20, 0xa4, 0xce, 0x25, 0xbb, 0xdb, 0x73, 0xa7, 0x3e, 0xe9, 0xc4,
	0x5d, 0x84, 0x54, 0x79, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0x83, 0x38, 0xad, 0x15, 0x46, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TerminalUIServiceClient is the client API for TerminalUIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TerminalUIServiceClient interface {
	Output(ctx context.Context, in *TerminalUI_OutputRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Events(ctx context.Context, opts ...grpc.CallOption) (TerminalUIService_EventsClient, error)
}

type terminalUIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTerminalUIServiceClient(cc grpc.ClientConnInterface) TerminalUIServiceClient {
	return &terminalUIServiceClient{cc}
}

func (c *terminalUIServiceClient) Output(ctx context.Context, in *TerminalUI_OutputRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.TerminalUIService/Output", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terminalUIServiceClient) Events(ctx context.Context, opts ...grpc.CallOption) (TerminalUIService_EventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TerminalUIService_serviceDesc.Streams[0], "/proto.TerminalUIService/Events", opts...)
	if err != nil {
		return nil, err
	}
	x := &terminalUIServiceEventsClient{stream}
	return x, nil
}

type TerminalUIService_EventsClient interface {
	Send(*TerminalUI_Event) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type terminalUIServiceEventsClient struct {
	grpc.ClientStream
}

func (x *terminalUIServiceEventsClient) Send(m *TerminalUI_Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *terminalUIServiceEventsClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TerminalUIServiceServer is the server API for TerminalUIService service.
type TerminalUIServiceServer interface {
	Output(context.Context, *TerminalUI_OutputRequest) (*empty.Empty, error)
	Events(TerminalUIService_EventsServer) error
}

// UnimplementedTerminalUIServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTerminalUIServiceServer struct {
}

func (*UnimplementedTerminalUIServiceServer) Output(ctx context.Context, req *TerminalUI_OutputRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Output not implemented")
}
func (*UnimplementedTerminalUIServiceServer) Events(srv TerminalUIService_EventsServer) error {
	return status.Errorf(codes.Unimplemented, "method Events not implemented")
}

func RegisterTerminalUIServiceServer(s *grpc.Server, srv TerminalUIServiceServer) {
	s.RegisterService(&_TerminalUIService_serviceDesc, srv)
}

func _TerminalUIService_Output_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminalUI_OutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalUIServiceServer).Output(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TerminalUIService/Output",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalUIServiceServer).Output(ctx, req.(*TerminalUI_OutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerminalUIService_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TerminalUIServiceServer).Events(&terminalUIServiceEventsServer{stream})
}

type TerminalUIService_EventsServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*TerminalUI_Event, error)
	grpc.ServerStream
}

type terminalUIServiceEventsServer struct {
	grpc.ServerStream
}

func (x *terminalUIServiceEventsServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *terminalUIServiceEventsServer) Recv() (*TerminalUI_Event, error) {
	m := new(TerminalUI_Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TerminalUIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TerminalUIService",
	HandlerType: (*TerminalUIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Output",
			Handler:    _TerminalUIService_Output_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Events",
			Handler:       _TerminalUIService_Events_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "terminal.proto",
}
