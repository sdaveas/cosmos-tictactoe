// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tictactoe/game.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Player int32

const (
	Player_NONE    Player = 0
	Player_CREATOR Player = 1
	Player_GUEST   Player = 2
)

var Player_name = map[int32]string{
	0: "NONE",
	1: "CREATOR",
	2: "GUEST",
}

var Player_value = map[string]int32{
	"NONE":    0,
	"CREATOR": 1,
	"GUEST":   2,
}

func (x Player) String() string {
	return proto.EnumName(Player_name, int32(x))
}

func (Player) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d1c70769b6ffdc4, []int{0}
}

type GameStatus int32

const (
	GameStatus_OPEN    GameStatus = 0
	GameStatus_RUNNING GameStatus = 1
	GameStatus_CLOSED  GameStatus = 2
)

var GameStatus_name = map[int32]string{
	0: "OPEN",
	1: "RUNNING",
	2: "CLOSED",
}

var GameStatus_value = map[string]int32{
	"OPEN":    0,
	"RUNNING": 1,
	"CLOSED":  2,
}

func (x GameStatus) String() string {
	return proto.EnumName(GameStatus_name, int32(x))
}

func (GameStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4d1c70769b6ffdc4, []int{1}
}

type Game struct {
	Creator string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Guest   string     `protobuf:"bytes,3,opt,name=guest,proto3" json:"guest,omitempty"`
	Oplayer Player     `protobuf:"varint,4,opt,name=oplayer,proto3,enum=sdaveas.tictactoe.tictactoe.Player" json:"oplayer,omitempty"`
	Xplayer Player     `protobuf:"varint,5,opt,name=xplayer,proto3,enum=sdaveas.tictactoe.tictactoe.Player" json:"xplayer,omitempty"`
	Winner  Player     `protobuf:"varint,6,opt,name=winner,proto3,enum=sdaveas.tictactoe.tictactoe.Player" json:"winner,omitempty"`
	Board   string     `protobuf:"bytes,7,opt,name=board,proto3" json:"board,omitempty"`
	Status  GameStatus `protobuf:"varint,8,opt,name=status,proto3,enum=sdaveas.tictactoe.tictactoe.GameStatus" json:"status,omitempty"`
}

func (m *Game) Reset()         { *m = Game{} }
func (m *Game) String() string { return proto.CompactTextString(m) }
func (*Game) ProtoMessage()    {}
func (*Game) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d1c70769b6ffdc4, []int{0}
}
func (m *Game) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Game) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Game.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Game) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Game.Merge(m, src)
}
func (m *Game) XXX_Size() int {
	return m.Size()
}
func (m *Game) XXX_DiscardUnknown() {
	xxx_messageInfo_Game.DiscardUnknown(m)
}

var xxx_messageInfo_Game proto.InternalMessageInfo

func (m *Game) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Game) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Game) GetGuest() string {
	if m != nil {
		return m.Guest
	}
	return ""
}

func (m *Game) GetOplayer() Player {
	if m != nil {
		return m.Oplayer
	}
	return Player_NONE
}

func (m *Game) GetXplayer() Player {
	if m != nil {
		return m.Xplayer
	}
	return Player_NONE
}

func (m *Game) GetWinner() Player {
	if m != nil {
		return m.Winner
	}
	return Player_NONE
}

func (m *Game) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

func (m *Game) GetStatus() GameStatus {
	if m != nil {
		return m.Status
	}
	return GameStatus_OPEN
}

type MsgCreateGame struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Guest   string `protobuf:"bytes,2,opt,name=guest,proto3" json:"guest,omitempty"`
}

func (m *MsgCreateGame) Reset()         { *m = MsgCreateGame{} }
func (m *MsgCreateGame) String() string { return proto.CompactTextString(m) }
func (*MsgCreateGame) ProtoMessage()    {}
func (*MsgCreateGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d1c70769b6ffdc4, []int{1}
}
func (m *MsgCreateGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateGame.Merge(m, src)
}
func (m *MsgCreateGame) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateGame) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateGame.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateGame proto.InternalMessageInfo

func (m *MsgCreateGame) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateGame) GetGuest() string {
	if m != nil {
		return m.Guest
	}
	return ""
}

type MsgAcceptGame struct {
	Guest string `protobuf:"bytes,1,opt,name=guest,proto3" json:"guest,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgAcceptGame) Reset()         { *m = MsgAcceptGame{} }
func (m *MsgAcceptGame) String() string { return proto.CompactTextString(m) }
func (*MsgAcceptGame) ProtoMessage()    {}
func (*MsgAcceptGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d1c70769b6ffdc4, []int{2}
}
func (m *MsgAcceptGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAcceptGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAcceptGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAcceptGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAcceptGame.Merge(m, src)
}
func (m *MsgAcceptGame) XXX_Size() int {
	return m.Size()
}
func (m *MsgAcceptGame) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAcceptGame.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAcceptGame proto.InternalMessageInfo

func (m *MsgAcceptGame) GetGuest() string {
	if m != nil {
		return m.Guest
	}
	return ""
}

func (m *MsgAcceptGame) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type MsgMakeMove struct {
	Caller string `protobuf:"bytes,1,opt,name=caller,proto3" json:"caller,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Cell   uint64 `protobuf:"varint,3,opt,name=cell,proto3" json:"cell,omitempty"`
}

func (m *MsgMakeMove) Reset()         { *m = MsgMakeMove{} }
func (m *MsgMakeMove) String() string { return proto.CompactTextString(m) }
func (*MsgMakeMove) ProtoMessage()    {}
func (*MsgMakeMove) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d1c70769b6ffdc4, []int{3}
}
func (m *MsgMakeMove) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMakeMove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMakeMove.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMakeMove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMakeMove.Merge(m, src)
}
func (m *MsgMakeMove) XXX_Size() int {
	return m.Size()
}
func (m *MsgMakeMove) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMakeMove.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMakeMove proto.InternalMessageInfo

func (m *MsgMakeMove) GetCaller() string {
	if m != nil {
		return m.Caller
	}
	return ""
}

func (m *MsgMakeMove) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgMakeMove) GetCell() uint64 {
	if m != nil {
		return m.Cell
	}
	return 0
}

type MsgDeleteGame struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgDeleteGame) Reset()         { *m = MsgDeleteGame{} }
func (m *MsgDeleteGame) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteGame) ProtoMessage()    {}
func (*MsgDeleteGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d1c70769b6ffdc4, []int{4}
}
func (m *MsgDeleteGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteGame.Merge(m, src)
}
func (m *MsgDeleteGame) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteGame) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteGame.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteGame proto.InternalMessageInfo

func (m *MsgDeleteGame) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeleteGame) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterEnum("sdaveas.tictactoe.tictactoe.Player", Player_name, Player_value)
	proto.RegisterEnum("sdaveas.tictactoe.tictactoe.GameStatus", GameStatus_name, GameStatus_value)
	proto.RegisterType((*Game)(nil), "sdaveas.tictactoe.tictactoe.Game")
	proto.RegisterType((*MsgCreateGame)(nil), "sdaveas.tictactoe.tictactoe.MsgCreateGame")
	proto.RegisterType((*MsgAcceptGame)(nil), "sdaveas.tictactoe.tictactoe.MsgAcceptGame")
	proto.RegisterType((*MsgMakeMove)(nil), "sdaveas.tictactoe.tictactoe.MsgMakeMove")
	proto.RegisterType((*MsgDeleteGame)(nil), "sdaveas.tictactoe.tictactoe.MsgDeleteGame")
}

func init() { proto.RegisterFile("tictactoe/game.proto", fileDescriptor_4d1c70769b6ffdc4) }

var fileDescriptor_4d1c70769b6ffdc4 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0xbd, 0xaa, 0x2c, 0x27, 0x13, 0x1a, 0xc4, 0x62, 0xca, 0xd2, 0x82, 0x08, 0xea, 0xa1,
	0xc1, 0x07, 0x09, 0x5a, 0x7a, 0x28, 0xa5, 0x84, 0xd4, 0x11, 0x6e, 0xa0, 0x96, 0x83, 0x9c, 0x5c,
	0x7a, 0x5b, 0xaf, 0x07, 0x55, 0x54, 0xce, 0x0a, 0xed, 0x3a, 0x4d, 0xde, 0xa2, 0x8f, 0xd5, 0x63,
	0x8e, 0x39, 0x16, 0xfb, 0x45, 0x8a, 0x56, 0xf2, 0x1f, 0x52, 0x48, 0x9a, 0xdb, 0xcc, 0xea, 0xfb,
	0x7d, 0xda, 0xf9, 0x86, 0x85, 0xae, 0xce, 0x84, 0xe6, 0x42, 0x4b, 0x0c, 0x53, 0x3e, 0xc3, 0xa0,
	0x28, 0xa5, 0x96, 0xf4, 0x95, 0x9a, 0xf2, 0x2b, 0xe4, 0x2a, 0x58, 0x7f, 0xdd, 0x54, 0x2f, 0xbb,
	0xa9, 0x4c, 0xa5, 0xd1, 0x85, 0x55, 0x55, 0x23, 0xfe, 0x9d, 0x05, 0xf6, 0x80, 0xcf, 0x90, 0x32,
	0xe8, 0x88, 0x12, 0xb9, 0x96, 0x25, 0x23, 0x07, 0xe4, 0x70, 0x37, 0x59, 0xb5, 0x74, 0x1f, 0xac,
	0x6c, 0xca, 0x2c, 0x73, 0x68, 0x65, 0x53, 0xda, 0x85, 0x76, 0x3a, 0x47, 0xa5, 0xd9, 0x33, 0x73,
	0x54, 0x37, 0xf4, 0x13, 0x74, 0x64, 0x91, 0xf3, 0x1b, 0x2c, 0x99, 0x7d, 0x40, 0x0e, 0xf7, 0xdf,
	0xbe, 0x0e, 0x1e, 0xb8, 0x4d, 0x70, 0x66, 0xa4, 0xc9, 0x8a, 0xa9, 0xf0, 0xeb, 0x06, 0x6f, 0x3f,
	0x01, 0x6f, 0x18, 0xfa, 0x11, 0x9c, 0x9f, 0xd9, 0xe5, 0x25, 0x96, 0xcc, 0xf9, 0x7f, 0xba, 0x41,
	0xaa, 0x81, 0x26, 0x92, 0x97, 0x53, 0xd6, 0xa9, 0x07, 0x32, 0x0d, 0x3d, 0x02, 0x47, 0x69, 0xae,
	0xe7, 0x8a, 0xed, 0x18, 0xcb, 0x37, 0x0f, 0x5a, 0x56, 0x19, 0x8e, 0x8d, 0x3c, 0x69, 0x30, 0xff,
	0x08, 0x9e, 0x0f, 0x55, 0xda, 0xaf, 0x52, 0xc4, 0x47, 0x22, 0x5e, 0x47, 0x6a, 0x6d, 0x45, 0xea,
	0xbf, 0x37, 0x06, 0xc7, 0x42, 0x60, 0xa1, 0x8d, 0xc1, 0x5a, 0x46, 0xb6, 0x93, 0xbf, 0xb7, 0x1f,
	0xff, 0x14, 0xf6, 0x86, 0x2a, 0x1d, 0xf2, 0x1f, 0x38, 0x94, 0x57, 0x48, 0x5f, 0x80, 0x23, 0x78,
	0x9e, 0xe3, 0xea, 0xa7, 0x4d, 0xf7, 0xcf, 0x5a, 0x29, 0xd8, 0x02, 0xf3, 0xdc, 0x6c, 0xd5, 0x4e,
	0x4c, 0xed, 0x7f, 0x30, 0x37, 0x38, 0xc1, 0x1c, 0x1f, 0x1d, 0xe1, 0x9e, 0x5d, 0xaf, 0x07, 0x4e,
	0x1d, 0x33, 0xdd, 0x01, 0x3b, 0x1e, 0xc5, 0x91, 0xdb, 0xa2, 0x7b, 0xd0, 0xe9, 0x27, 0xd1, 0xf1,
	0xf9, 0x28, 0x71, 0x09, 0xdd, 0x85, 0xf6, 0xe0, 0x22, 0x1a, 0x9f, 0xbb, 0x56, 0x2f, 0x04, 0xd8,
	0xe4, 0x57, 0xe9, 0x47, 0x67, 0x51, 0x5c, 0xeb, 0x93, 0x8b, 0x38, 0x3e, 0x8d, 0x07, 0x2e, 0xa1,
	0x00, 0x4e, 0xff, 0xeb, 0x68, 0x1c, 0x9d, 0xb8, 0xd6, 0xe7, 0x2f, 0xbf, 0x17, 0x1e, 0xb9, 0x5d,
	0x78, 0xe4, 0xcf, 0xc2, 0x23, 0xbf, 0x96, 0x5e, 0xeb, 0x76, 0xe9, 0xb5, 0xee, 0x96, 0x5e, 0xeb,
	0x5b, 0x90, 0x66, 0xfa, 0xfb, 0x7c, 0x12, 0x08, 0x39, 0x0b, 0x9b, 0x7d, 0x85, 0x9b, 0xb7, 0x72,
	0xbd, 0x55, 0xeb, 0x9b, 0x02, 0xd5, 0xc4, 0x31, 0xcf, 0xe0, 0xdd, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x75, 0x58, 0x71, 0x21, 0x51, 0x03, 0x00, 0x00,
}

func (m *Game) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Game) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Game) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Board) > 0 {
		i -= len(m.Board)
		copy(dAtA[i:], m.Board)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Board)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Winner != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.Winner))
		i--
		dAtA[i] = 0x30
	}
	if m.Xplayer != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.Xplayer))
		i--
		dAtA[i] = 0x28
	}
	if m.Oplayer != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.Oplayer))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Guest) > 0 {
		i -= len(m.Guest)
		copy(dAtA[i:], m.Guest)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Guest)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Guest) > 0 {
		i -= len(m.Guest)
		copy(dAtA[i:], m.Guest)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Guest)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAcceptGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAcceptGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAcceptGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Guest) > 0 {
		i -= len(m.Guest)
		copy(dAtA[i:], m.Guest)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Guest)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMakeMove) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMakeMove) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMakeMove) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Cell != 0 {
		i = encodeVarintGame(dAtA, i, uint64(m.Cell))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Caller) > 0 {
		i -= len(m.Caller)
		copy(dAtA[i:], m.Caller)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Caller)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintGame(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGame(dAtA []byte, offset int, v uint64) int {
	offset -= sovGame(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Game) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Guest)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	if m.Oplayer != 0 {
		n += 1 + sovGame(uint64(m.Oplayer))
	}
	if m.Xplayer != 0 {
		n += 1 + sovGame(uint64(m.Xplayer))
	}
	if m.Winner != 0 {
		n += 1 + sovGame(uint64(m.Winner))
	}
	l = len(m.Board)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovGame(uint64(m.Status))
	}
	return n
}

func (m *MsgCreateGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Guest)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *MsgAcceptGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Guest)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func (m *MsgMakeMove) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Caller)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	if m.Cell != 0 {
		n += 1 + sovGame(uint64(m.Cell))
	}
	return n
}

func (m *MsgDeleteGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovGame(uint64(l))
	}
	return n
}

func sovGame(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGame(x uint64) (n int) {
	return sovGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Game) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Game: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Game: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Guest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Guest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oplayer", wireType)
			}
			m.Oplayer = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Oplayer |= Player(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Xplayer", wireType)
			}
			m.Xplayer = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Xplayer |= Player(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			m.Winner = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Winner |= Player(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Board", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Board = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= GameStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGame
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Guest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Guest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGame
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAcceptGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgAcceptGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAcceptGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Guest", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Guest = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGame
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgMakeMove) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgMakeMove: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMakeMove: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Caller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Caller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cell", wireType)
			}
			m.Cell = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cell |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGame
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeleteGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGame
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeleteGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGame
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGame
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGame
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGame
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGame
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGame
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGame
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGame
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGame        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGame          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGame = fmt.Errorf("proto: unexpected end of group")
)
