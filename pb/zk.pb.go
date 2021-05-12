// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: zk.proto

package pb

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

type ZKMod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	W *Int   `protobuf:"bytes,1,opt,name=w,proto3" json:"w,omitempty"`
	X []*Int `protobuf:"bytes,2,rep,name=x,proto3" json:"x,omitempty"`
	A []bool `protobuf:"varint,3,rep,packed,name=a,proto3" json:"a,omitempty"`
	B []bool `protobuf:"varint,4,rep,packed,name=b,proto3" json:"b,omitempty"`
	Z []*Int `protobuf:"bytes,5,rep,name=z,proto3" json:"z,omitempty"`
}

func (x *ZKMod) Reset() {
	*x = ZKMod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZKMod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZKMod) ProtoMessage() {}

func (x *ZKMod) ProtoReflect() protoreflect.Message {
	mi := &file_zk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZKMod.ProtoReflect.Descriptor instead.
func (*ZKMod) Descriptor() ([]byte, []int) {
	return file_zk_proto_rawDescGZIP(), []int{0}
}

func (x *ZKMod) GetW() *Int {
	if x != nil {
		return x.W
	}
	return nil
}

func (x *ZKMod) GetX() []*Int {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *ZKMod) GetA() []bool {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *ZKMod) GetB() []bool {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *ZKMod) GetZ() []*Int {
	if x != nil {
		return x.Z
	}
	return nil
}

type ZKPrm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A []*Int `protobuf:"bytes,1,rep,name=a,proto3" json:"a,omitempty"`
	Z []*Int `protobuf:"bytes,2,rep,name=z,proto3" json:"z,omitempty"`
}

func (x *ZKPrm) Reset() {
	*x = ZKPrm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZKPrm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZKPrm) ProtoMessage() {}

func (x *ZKPrm) ProtoReflect() protoreflect.Message {
	mi := &file_zk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZKPrm.ProtoReflect.Descriptor instead.
func (*ZKPrm) Descriptor() ([]byte, []int) {
	return file_zk_proto_rawDescGZIP(), []int{1}
}

func (x *ZKPrm) GetA() []*Int {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *ZKPrm) GetZ() []*Int {
	if x != nil {
		return x.Z
	}
	return nil
}

type ZKEnc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S  *Int        `protobuf:"bytes,1,opt,name=S,proto3" json:"S,omitempty"`
	A  *Ciphertext `protobuf:"bytes,2,opt,name=A,proto3" json:"A,omitempty"`
	C  *Int        `protobuf:"bytes,3,opt,name=C,proto3" json:"C,omitempty"`
	Z1 *Int        `protobuf:"bytes,4,opt,name=Z1,proto3" json:"Z1,omitempty"`
	Z2 *Int        `protobuf:"bytes,5,opt,name=Z2,proto3" json:"Z2,omitempty"`
	Z3 *Int        `protobuf:"bytes,6,opt,name=Z3,proto3" json:"Z3,omitempty"`
}

func (x *ZKEnc) Reset() {
	*x = ZKEnc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZKEnc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZKEnc) ProtoMessage() {}

func (x *ZKEnc) ProtoReflect() protoreflect.Message {
	mi := &file_zk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZKEnc.ProtoReflect.Descriptor instead.
func (*ZKEnc) Descriptor() ([]byte, []int) {
	return file_zk_proto_rawDescGZIP(), []int{2}
}

func (x *ZKEnc) GetS() *Int {
	if x != nil {
		return x.S
	}
	return nil
}

func (x *ZKEnc) GetA() *Ciphertext {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *ZKEnc) GetC() *Int {
	if x != nil {
		return x.C
	}
	return nil
}

func (x *ZKEnc) GetZ1() *Int {
	if x != nil {
		return x.Z1
	}
	return nil
}

func (x *ZKEnc) GetZ2() *Int {
	if x != nil {
		return x.Z2
	}
	return nil
}

func (x *ZKEnc) GetZ3() *Int {
	if x != nil {
		return x.Z3
	}
	return nil
}

type ZKLogStar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S  *Int        `protobuf:"bytes,1,opt,name=S,proto3" json:"S,omitempty"`
	A  *Ciphertext `protobuf:"bytes,2,opt,name=A,proto3" json:"A,omitempty"`
	Y  *Point      `protobuf:"bytes,3,opt,name=Y,proto3" json:"Y,omitempty"`
	D  *Int        `protobuf:"bytes,4,opt,name=D,proto3" json:"D,omitempty"`
	Z1 *Int        `protobuf:"bytes,5,opt,name=Z1,proto3" json:"Z1,omitempty"`
	Z2 *Int        `protobuf:"bytes,6,opt,name=Z2,proto3" json:"Z2,omitempty"`
	Z3 *Int        `protobuf:"bytes,7,opt,name=Z3,proto3" json:"Z3,omitempty"`
}

func (x *ZKLogStar) Reset() {
	*x = ZKLogStar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZKLogStar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZKLogStar) ProtoMessage() {}

func (x *ZKLogStar) ProtoReflect() protoreflect.Message {
	mi := &file_zk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZKLogStar.ProtoReflect.Descriptor instead.
func (*ZKLogStar) Descriptor() ([]byte, []int) {
	return file_zk_proto_rawDescGZIP(), []int{3}
}

func (x *ZKLogStar) GetS() *Int {
	if x != nil {
		return x.S
	}
	return nil
}

func (x *ZKLogStar) GetA() *Ciphertext {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *ZKLogStar) GetY() *Point {
	if x != nil {
		return x.Y
	}
	return nil
}

func (x *ZKLogStar) GetD() *Int {
	if x != nil {
		return x.D
	}
	return nil
}

func (x *ZKLogStar) GetZ1() *Int {
	if x != nil {
		return x.Z1
	}
	return nil
}

func (x *ZKLogStar) GetZ2() *Int {
	if x != nil {
		return x.Z2
	}
	return nil
}

func (x *ZKLogStar) GetZ3() *Int {
	if x != nil {
		return x.Z3
	}
	return nil
}

type ZKAffG struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A  *Ciphertext `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	Bx *Point      `protobuf:"bytes,2,opt,name=Bx,proto3" json:"Bx,omitempty"`
	By *Ciphertext `protobuf:"bytes,3,opt,name=By,proto3" json:"By,omitempty"`
	E  *Int        `protobuf:"bytes,4,opt,name=E,proto3" json:"E,omitempty"`
	S  *Int        `protobuf:"bytes,5,opt,name=S,proto3" json:"S,omitempty"`
	F  *Int        `protobuf:"bytes,6,opt,name=F,proto3" json:"F,omitempty"`
	T  *Int        `protobuf:"bytes,7,opt,name=T,proto3" json:"T,omitempty"`
	Z1 *Int        `protobuf:"bytes,8,opt,name=Z1,proto3" json:"Z1,omitempty"`
	Z2 *Int        `protobuf:"bytes,9,opt,name=Z2,proto3" json:"Z2,omitempty"`
	Z3 *Int        `protobuf:"bytes,10,opt,name=Z3,proto3" json:"Z3,omitempty"`
	Z4 *Int        `protobuf:"bytes,11,opt,name=Z4,proto3" json:"Z4,omitempty"`
	W  *Int        `protobuf:"bytes,12,opt,name=W,proto3" json:"W,omitempty"`
	Wy *Int        `protobuf:"bytes,13,opt,name=Wy,proto3" json:"Wy,omitempty"`
}

func (x *ZKAffG) Reset() {
	*x = ZKAffG{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZKAffG) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZKAffG) ProtoMessage() {}

func (x *ZKAffG) ProtoReflect() protoreflect.Message {
	mi := &file_zk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZKAffG.ProtoReflect.Descriptor instead.
func (*ZKAffG) Descriptor() ([]byte, []int) {
	return file_zk_proto_rawDescGZIP(), []int{4}
}

func (x *ZKAffG) GetA() *Ciphertext {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *ZKAffG) GetBx() *Point {
	if x != nil {
		return x.Bx
	}
	return nil
}

func (x *ZKAffG) GetBy() *Ciphertext {
	if x != nil {
		return x.By
	}
	return nil
}

func (x *ZKAffG) GetE() *Int {
	if x != nil {
		return x.E
	}
	return nil
}

func (x *ZKAffG) GetS() *Int {
	if x != nil {
		return x.S
	}
	return nil
}

func (x *ZKAffG) GetF() *Int {
	if x != nil {
		return x.F
	}
	return nil
}

func (x *ZKAffG) GetT() *Int {
	if x != nil {
		return x.T
	}
	return nil
}

func (x *ZKAffG) GetZ1() *Int {
	if x != nil {
		return x.Z1
	}
	return nil
}

func (x *ZKAffG) GetZ2() *Int {
	if x != nil {
		return x.Z2
	}
	return nil
}

func (x *ZKAffG) GetZ3() *Int {
	if x != nil {
		return x.Z3
	}
	return nil
}

func (x *ZKAffG) GetZ4() *Int {
	if x != nil {
		return x.Z4
	}
	return nil
}

func (x *ZKAffG) GetW() *Int {
	if x != nil {
		return x.W
	}
	return nil
}

func (x *ZKAffG) GetWy() *Int {
	if x != nil {
		return x.Wy
	}
	return nil
}

type ZKMulStar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A  *Ciphertext `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	Bx *Point      `protobuf:"bytes,2,opt,name=Bx,proto3" json:"Bx,omitempty"`
	E  *Int        `protobuf:"bytes,4,opt,name=E,proto3" json:"E,omitempty"`
	S  *Int        `protobuf:"bytes,5,opt,name=S,proto3" json:"S,omitempty"`
	Z1 *Int        `protobuf:"bytes,8,opt,name=Z1,proto3" json:"Z1,omitempty"`
	Z2 *Int        `protobuf:"bytes,9,opt,name=Z2,proto3" json:"Z2,omitempty"`
	W  *Int        `protobuf:"bytes,12,opt,name=W,proto3" json:"W,omitempty"`
}

func (x *ZKMulStar) Reset() {
	*x = ZKMulStar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zk_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZKMulStar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZKMulStar) ProtoMessage() {}

func (x *ZKMulStar) ProtoReflect() protoreflect.Message {
	mi := &file_zk_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZKMulStar.ProtoReflect.Descriptor instead.
func (*ZKMulStar) Descriptor() ([]byte, []int) {
	return file_zk_proto_rawDescGZIP(), []int{5}
}

func (x *ZKMulStar) GetA() *Ciphertext {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *ZKMulStar) GetBx() *Point {
	if x != nil {
		return x.Bx
	}
	return nil
}

func (x *ZKMulStar) GetE() *Int {
	if x != nil {
		return x.E
	}
	return nil
}

func (x *ZKMulStar) GetS() *Int {
	if x != nil {
		return x.S
	}
	return nil
}

func (x *ZKMulStar) GetZ1() *Int {
	if x != nil {
		return x.Z1
	}
	return nil
}

func (x *ZKMulStar) GetZ2() *Int {
	if x != nil {
		return x.Z2
	}
	return nil
}

func (x *ZKMulStar) GetW() *Int {
	if x != nil {
		return x.W
	}
	return nil
}

type ZKMul struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A *Ciphertext `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	B *Ciphertext `protobuf:"bytes,2,opt,name=B,proto3" json:"B,omitempty"`
	Z *Int        `protobuf:"bytes,3,opt,name=Z,proto3" json:"Z,omitempty"`
	U *Int        `protobuf:"bytes,4,opt,name=U,proto3" json:"U,omitempty"`
	V *Int        `protobuf:"bytes,5,opt,name=V,proto3" json:"V,omitempty"`
}

func (x *ZKMul) Reset() {
	*x = ZKMul{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zk_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZKMul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZKMul) ProtoMessage() {}

func (x *ZKMul) ProtoReflect() protoreflect.Message {
	mi := &file_zk_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZKMul.ProtoReflect.Descriptor instead.
func (*ZKMul) Descriptor() ([]byte, []int) {
	return file_zk_proto_rawDescGZIP(), []int{6}
}

func (x *ZKMul) GetA() *Ciphertext {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *ZKMul) GetB() *Ciphertext {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *ZKMul) GetZ() *Int {
	if x != nil {
		return x.Z
	}
	return nil
}

func (x *ZKMul) GetU() *Int {
	if x != nil {
		return x.U
	}
	return nil
}

func (x *ZKMul) GetV() *Int {
	if x != nil {
		return x.V
	}
	return nil
}

type ZKDec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	S     *Int        `protobuf:"bytes,1,opt,name=S,proto3" json:"S,omitempty"`
	T     *Int        `protobuf:"bytes,2,opt,name=T,proto3" json:"T,omitempty"`
	A     *Ciphertext `protobuf:"bytes,3,opt,name=A,proto3" json:"A,omitempty"`
	Gamma *Scalar     `protobuf:"bytes,4,opt,name=Gamma,proto3" json:"Gamma,omitempty"`
	Z1    *Int        `protobuf:"bytes,5,opt,name=Z1,proto3" json:"Z1,omitempty"`
	Z2    *Int        `protobuf:"bytes,6,opt,name=Z2,proto3" json:"Z2,omitempty"`
	W     *Int        `protobuf:"bytes,7,opt,name=W,proto3" json:"W,omitempty"`
}

func (x *ZKDec) Reset() {
	*x = ZKDec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zk_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZKDec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZKDec) ProtoMessage() {}

func (x *ZKDec) ProtoReflect() protoreflect.Message {
	mi := &file_zk_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZKDec.ProtoReflect.Descriptor instead.
func (*ZKDec) Descriptor() ([]byte, []int) {
	return file_zk_proto_rawDescGZIP(), []int{7}
}

func (x *ZKDec) GetS() *Int {
	if x != nil {
		return x.S
	}
	return nil
}

func (x *ZKDec) GetT() *Int {
	if x != nil {
		return x.T
	}
	return nil
}

func (x *ZKDec) GetA() *Ciphertext {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *ZKDec) GetGamma() *Scalar {
	if x != nil {
		return x.Gamma
	}
	return nil
}

func (x *ZKDec) GetZ1() *Int {
	if x != nil {
		return x.Z1
	}
	return nil
}

func (x *ZKDec) GetZ2() *Int {
	if x != nil {
		return x.Z2
	}
	return nil
}

func (x *ZKDec) GetW() *Int {
	if x != nil {
		return x.W
	}
	return nil
}

var File_zk_proto protoreflect.FileDescriptor

var file_zk_proto_rawDesc = []byte{
	0x0a, 0x08, 0x7a, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0b,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x05, 0x5a,
	0x4b, 0x4d, 0x6f, 0x64, 0x12, 0x15, 0x0a, 0x01, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x77, 0x12, 0x15, 0x0a, 0x01, 0x78,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52,
	0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x08, 0x52, 0x01, 0x61,
	0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x04, 0x20, 0x03, 0x28, 0x08, 0x52, 0x01, 0x62, 0x12, 0x15,
	0x0a, 0x01, 0x7a, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49,
	0x6e, 0x74, 0x52, 0x01, 0x7a, 0x22, 0x35, 0x0a, 0x05, 0x5a, 0x4b, 0x50, 0x72, 0x6d, 0x12, 0x15,
	0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49,
	0x6e, 0x74, 0x52, 0x01, 0x61, 0x12, 0x15, 0x0a, 0x01, 0x7a, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x7a, 0x22, 0x9e, 0x01, 0x0a,
	0x05, 0x5a, 0x4b, 0x45, 0x6e, 0x63, 0x12, 0x15, 0x0a, 0x01, 0x53, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x53, 0x12, 0x1c, 0x0a,
	0x01, 0x41, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x69,
	0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x52, 0x01, 0x41, 0x12, 0x15, 0x0a, 0x01, 0x43,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52,
	0x01, 0x43, 0x12, 0x17, 0x0a, 0x02, 0x5a, 0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x02, 0x5a, 0x31, 0x12, 0x17, 0x0a, 0x02, 0x5a,
	0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74,
	0x52, 0x02, 0x5a, 0x32, 0x12, 0x17, 0x0a, 0x02, 0x5a, 0x33, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x02, 0x5a, 0x33, 0x22, 0xbb, 0x01,
	0x0a, 0x09, 0x5a, 0x4b, 0x4c, 0x6f, 0x67, 0x53, 0x74, 0x61, 0x72, 0x12, 0x15, 0x0a, 0x01, 0x53,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52,
	0x01, 0x53, 0x12, 0x1c, 0x0a, 0x01, 0x41, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x52, 0x01, 0x41,
	0x12, 0x17, 0x0a, 0x01, 0x59, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x01, 0x59, 0x12, 0x15, 0x0a, 0x01, 0x44, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x44,
	0x12, 0x17, 0x0a, 0x02, 0x5a, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70,
	0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x02, 0x5a, 0x31, 0x12, 0x17, 0x0a, 0x02, 0x5a, 0x32, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x02,
	0x5a, 0x32, 0x12, 0x17, 0x0a, 0x02, 0x5a, 0x33, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x02, 0x5a, 0x33, 0x22, 0xd1, 0x02, 0x0a, 0x06,
	0x5a, 0x4b, 0x41, 0x66, 0x66, 0x47, 0x12, 0x1c, 0x0a, 0x01, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78,
	0x74, 0x52, 0x01, 0x41, 0x12, 0x19, 0x0a, 0x02, 0x42, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02, 0x42, 0x78, 0x12,
	0x1e, 0x0a, 0x02, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x52, 0x02, 0x42, 0x79, 0x12,
	0x15, 0x0a, 0x01, 0x45, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e,
	0x49, 0x6e, 0x74, 0x52, 0x01, 0x45, 0x12, 0x15, 0x0a, 0x01, 0x53, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x53, 0x12, 0x15, 0x0a,
	0x01, 0x46, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e,
	0x74, 0x52, 0x01, 0x46, 0x12, 0x15, 0x0a, 0x01, 0x54, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x54, 0x12, 0x17, 0x0a, 0x02, 0x5a,
	0x31, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74,
	0x52, 0x02, 0x5a, 0x31, 0x12, 0x17, 0x0a, 0x02, 0x5a, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x02, 0x5a, 0x32, 0x12, 0x17, 0x0a,
	0x02, 0x5a, 0x33, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49,
	0x6e, 0x74, 0x52, 0x02, 0x5a, 0x33, 0x12, 0x17, 0x0a, 0x02, 0x5a, 0x34, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x02, 0x5a, 0x34, 0x12,
	0x15, 0x0a, 0x01, 0x57, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e,
	0x49, 0x6e, 0x74, 0x52, 0x01, 0x57, 0x12, 0x17, 0x0a, 0x02, 0x57, 0x79, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x02, 0x57, 0x79, 0x22,
	0xbb, 0x01, 0x0a, 0x09, 0x5a, 0x4b, 0x4d, 0x75, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x12, 0x1c, 0x0a,
	0x01, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x69,
	0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x52, 0x01, 0x41, 0x12, 0x19, 0x0a, 0x02, 0x42,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x02, 0x42, 0x78, 0x12, 0x15, 0x0a, 0x01, 0x45, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x45, 0x12, 0x15, 0x0a,
	0x01, 0x53, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e,
	0x74, 0x52, 0x01, 0x53, 0x12, 0x17, 0x0a, 0x02, 0x5a, 0x31, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x02, 0x5a, 0x31, 0x12, 0x17, 0x0a,
	0x02, 0x5a, 0x32, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49,
	0x6e, 0x74, 0x52, 0x02, 0x5a, 0x32, 0x12, 0x15, 0x0a, 0x01, 0x57, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x57, 0x22, 0x88, 0x01,
	0x0a, 0x05, 0x5a, 0x4b, 0x4d, 0x75, 0x6c, 0x12, 0x1c, 0x0a, 0x01, 0x41, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65,
	0x78, 0x74, 0x52, 0x01, 0x41, 0x12, 0x1c, 0x0a, 0x01, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74,
	0x52, 0x01, 0x42, 0x12, 0x15, 0x0a, 0x01, 0x5a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x5a, 0x12, 0x15, 0x0a, 0x01, 0x55, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01,
	0x55, 0x12, 0x15, 0x0a, 0x01, 0x56, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70,
	0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x56, 0x22, 0xbe, 0x01, 0x0a, 0x05, 0x5a, 0x4b, 0x44,
	0x65, 0x63, 0x12, 0x15, 0x0a, 0x01, 0x53, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x53, 0x12, 0x15, 0x0a, 0x01, 0x54, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x54,
	0x12, 0x1c, 0x0a, 0x01, 0x41, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x52, 0x01, 0x41, 0x12, 0x20,
	0x0a, 0x05, 0x47, 0x61, 0x6d, 0x6d, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x05, 0x47, 0x61, 0x6d, 0x6d, 0x61,
	0x12, 0x17, 0x0a, 0x02, 0x5a, 0x31, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70,
	0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x02, 0x5a, 0x31, 0x12, 0x17, 0x0a, 0x02, 0x5a, 0x32, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x02,
	0x5a, 0x32, 0x12, 0x15, 0x0a, 0x01, 0x57, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x57, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x75, 0x72, 0x75, 0x73, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2f, 0x63, 0x6d, 0x70, 0x2d, 0x65, 0x63, 0x64, 0x73, 0x61, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zk_proto_rawDescOnce sync.Once
	file_zk_proto_rawDescData = file_zk_proto_rawDesc
)

func file_zk_proto_rawDescGZIP() []byte {
	file_zk_proto_rawDescOnce.Do(func() {
		file_zk_proto_rawDescData = protoimpl.X.CompressGZIP(file_zk_proto_rawDescData)
	})
	return file_zk_proto_rawDescData
}

var file_zk_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_zk_proto_goTypes = []interface{}{
	(*ZKMod)(nil),      // 0: pb.ZKMod
	(*ZKPrm)(nil),      // 1: pb.ZKPrm
	(*ZKEnc)(nil),      // 2: pb.ZKEnc
	(*ZKLogStar)(nil),  // 3: pb.ZKLogStar
	(*ZKAffG)(nil),     // 4: pb.ZKAffG
	(*ZKMulStar)(nil),  // 5: pb.ZKMulStar
	(*ZKMul)(nil),      // 6: pb.ZKMul
	(*ZKDec)(nil),      // 7: pb.ZKDec
	(*Int)(nil),        // 8: pb.Int
	(*Ciphertext)(nil), // 9: pb.Ciphertext
	(*Point)(nil),      // 10: pb.Point
	(*Scalar)(nil),     // 11: pb.Scalar
}
var file_zk_proto_depIdxs = []int32{
	8,  // 0: pb.ZKMod.w:type_name -> pb.Int
	8,  // 1: pb.ZKMod.x:type_name -> pb.Int
	8,  // 2: pb.ZKMod.z:type_name -> pb.Int
	8,  // 3: pb.ZKPrm.a:type_name -> pb.Int
	8,  // 4: pb.ZKPrm.z:type_name -> pb.Int
	8,  // 5: pb.ZKEnc.S:type_name -> pb.Int
	9,  // 6: pb.ZKEnc.A:type_name -> pb.Ciphertext
	8,  // 7: pb.ZKEnc.C:type_name -> pb.Int
	8,  // 8: pb.ZKEnc.Z1:type_name -> pb.Int
	8,  // 9: pb.ZKEnc.Z2:type_name -> pb.Int
	8,  // 10: pb.ZKEnc.Z3:type_name -> pb.Int
	8,  // 11: pb.ZKLogStar.S:type_name -> pb.Int
	9,  // 12: pb.ZKLogStar.A:type_name -> pb.Ciphertext
	10, // 13: pb.ZKLogStar.Y:type_name -> pb.Point
	8,  // 14: pb.ZKLogStar.D:type_name -> pb.Int
	8,  // 15: pb.ZKLogStar.Z1:type_name -> pb.Int
	8,  // 16: pb.ZKLogStar.Z2:type_name -> pb.Int
	8,  // 17: pb.ZKLogStar.Z3:type_name -> pb.Int
	9,  // 18: pb.ZKAffG.A:type_name -> pb.Ciphertext
	10, // 19: pb.ZKAffG.Bx:type_name -> pb.Point
	9,  // 20: pb.ZKAffG.By:type_name -> pb.Ciphertext
	8,  // 21: pb.ZKAffG.E:type_name -> pb.Int
	8,  // 22: pb.ZKAffG.S:type_name -> pb.Int
	8,  // 23: pb.ZKAffG.F:type_name -> pb.Int
	8,  // 24: pb.ZKAffG.T:type_name -> pb.Int
	8,  // 25: pb.ZKAffG.Z1:type_name -> pb.Int
	8,  // 26: pb.ZKAffG.Z2:type_name -> pb.Int
	8,  // 27: pb.ZKAffG.Z3:type_name -> pb.Int
	8,  // 28: pb.ZKAffG.Z4:type_name -> pb.Int
	8,  // 29: pb.ZKAffG.W:type_name -> pb.Int
	8,  // 30: pb.ZKAffG.Wy:type_name -> pb.Int
	9,  // 31: pb.ZKMulStar.A:type_name -> pb.Ciphertext
	10, // 32: pb.ZKMulStar.Bx:type_name -> pb.Point
	8,  // 33: pb.ZKMulStar.E:type_name -> pb.Int
	8,  // 34: pb.ZKMulStar.S:type_name -> pb.Int
	8,  // 35: pb.ZKMulStar.Z1:type_name -> pb.Int
	8,  // 36: pb.ZKMulStar.Z2:type_name -> pb.Int
	8,  // 37: pb.ZKMulStar.W:type_name -> pb.Int
	9,  // 38: pb.ZKMul.A:type_name -> pb.Ciphertext
	9,  // 39: pb.ZKMul.B:type_name -> pb.Ciphertext
	8,  // 40: pb.ZKMul.Z:type_name -> pb.Int
	8,  // 41: pb.ZKMul.U:type_name -> pb.Int
	8,  // 42: pb.ZKMul.V:type_name -> pb.Int
	8,  // 43: pb.ZKDec.S:type_name -> pb.Int
	8,  // 44: pb.ZKDec.T:type_name -> pb.Int
	9,  // 45: pb.ZKDec.A:type_name -> pb.Ciphertext
	11, // 46: pb.ZKDec.Gamma:type_name -> pb.Scalar
	8,  // 47: pb.ZKDec.Z1:type_name -> pb.Int
	8,  // 48: pb.ZKDec.Z2:type_name -> pb.Int
	8,  // 49: pb.ZKDec.W:type_name -> pb.Int
	50, // [50:50] is the sub-list for method output_type
	50, // [50:50] is the sub-list for method input_type
	50, // [50:50] is the sub-list for extension type_name
	50, // [50:50] is the sub-list for extension extendee
	0,  // [0:50] is the sub-list for field type_name
}

func init() { file_zk_proto_init() }
func file_zk_proto_init() {
	if File_zk_proto != nil {
		return
	}
	file_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_zk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZKMod); i {
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
		file_zk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZKPrm); i {
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
		file_zk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZKEnc); i {
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
		file_zk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZKLogStar); i {
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
		file_zk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZKAffG); i {
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
		file_zk_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZKMulStar); i {
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
		file_zk_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZKMul); i {
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
		file_zk_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZKDec); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zk_proto_goTypes,
		DependencyIndexes: file_zk_proto_depIdxs,
		MessageInfos:      file_zk_proto_msgTypes,
	}.Build()
	File_zk_proto = out.File
	file_zk_proto_rawDesc = nil
	file_zk_proto_goTypes = nil
	file_zk_proto_depIdxs = nil
}
