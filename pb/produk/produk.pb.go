// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: protos/produk.proto

package produk

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_produk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_produk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_protos_produk_proto_rawDescGZIP(), []int{0}
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NamaProduk   string       `protobuf:"bytes,2,opt,name=nama_produk,json=namaProduk,proto3" json:"nama_produk,omitempty"`
	Kategori     string       `protobuf:"bytes,3,opt,name=kategori,proto3" json:"kategori,omitempty"`
	Harga        float64      `protobuf:"fixed64,4,opt,name=harga,proto3" json:"harga,omitempty"`
	Stok         uint32       `protobuf:"varint,5,opt,name=stok,proto3" json:"stok,omitempty"`
	Deskripsi    string       `protobuf:"bytes,6,opt,name=deskripsi,proto3" json:"deskripsi,omitempty"`
	Spesifikasi  *Spesifikasi `protobuf:"bytes,7,opt,name=spesifikasi,proto3" json:"spesifikasi,omitempty"`
	Merek        string       `protobuf:"bytes,8,opt,name=merek,proto3" json:"merek,omitempty"`
	TanggalRilis string       `protobuf:"bytes,9,opt,name=tanggal_rilis,json=tanggalRilis,proto3" json:"tanggal_rilis,omitempty"`
	Rating       float64      `protobuf:"fixed64,10,opt,name=rating,proto3" json:"rating,omitempty"`
	Ulasan       []*Ulasan    `protobuf:"bytes,11,rep,name=ulasan,proto3" json:"ulasan,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_produk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_protos_produk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_protos_produk_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetNamaProduk() string {
	if x != nil {
		return x.NamaProduk
	}
	return ""
}

func (x *Product) GetKategori() string {
	if x != nil {
		return x.Kategori
	}
	return ""
}

func (x *Product) GetHarga() float64 {
	if x != nil {
		return x.Harga
	}
	return 0
}

func (x *Product) GetStok() uint32 {
	if x != nil {
		return x.Stok
	}
	return 0
}

func (x *Product) GetDeskripsi() string {
	if x != nil {
		return x.Deskripsi
	}
	return ""
}

func (x *Product) GetSpesifikasi() *Spesifikasi {
	if x != nil {
		return x.Spesifikasi
	}
	return nil
}

func (x *Product) GetMerek() string {
	if x != nil {
		return x.Merek
	}
	return ""
}

func (x *Product) GetTanggalRilis() string {
	if x != nil {
		return x.TanggalRilis
	}
	return ""
}

func (x *Product) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Product) GetUlasan() []*Ulasan {
	if x != nil {
		return x.Ulasan
	}
	return nil
}

type Spesifikasi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdSpesifikasi int64  `protobuf:"varint,1,opt,name=id_spesifikasi,json=idSpesifikasi,proto3" json:"id_spesifikasi,omitempty"`
	Procesor      string `protobuf:"bytes,2,opt,name=procesor,proto3" json:"procesor,omitempty"`
	Ram           string `protobuf:"bytes,3,opt,name=ram,proto3" json:"ram,omitempty"`
	Penyimpanan   string `protobuf:"bytes,4,opt,name=penyimpanan,proto3" json:"penyimpanan,omitempty"`
	Layar         string `protobuf:"bytes,5,opt,name=layar,proto3" json:"layar,omitempty"`
	Grafis        string `protobuf:"bytes,6,opt,name=grafis,proto3" json:"grafis,omitempty"`
}

func (x *Spesifikasi) Reset() {
	*x = Spesifikasi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_produk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spesifikasi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spesifikasi) ProtoMessage() {}

func (x *Spesifikasi) ProtoReflect() protoreflect.Message {
	mi := &file_protos_produk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spesifikasi.ProtoReflect.Descriptor instead.
func (*Spesifikasi) Descriptor() ([]byte, []int) {
	return file_protos_produk_proto_rawDescGZIP(), []int{2}
}

func (x *Spesifikasi) GetIdSpesifikasi() int64 {
	if x != nil {
		return x.IdSpesifikasi
	}
	return 0
}

func (x *Spesifikasi) GetProcesor() string {
	if x != nil {
		return x.Procesor
	}
	return ""
}

func (x *Spesifikasi) GetRam() string {
	if x != nil {
		return x.Ram
	}
	return ""
}

func (x *Spesifikasi) GetPenyimpanan() string {
	if x != nil {
		return x.Penyimpanan
	}
	return ""
}

func (x *Spesifikasi) GetLayar() string {
	if x != nil {
		return x.Layar
	}
	return ""
}

func (x *Spesifikasi) GetGrafis() string {
	if x != nil {
		return x.Grafis
	}
	return ""
}

type Ulasan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdUlasan      uint64  `protobuf:"varint,1,opt,name=id_ulasan,json=idUlasan,proto3" json:"id_ulasan,omitempty"`
	NamaPelanggan string  `protobuf:"bytes,2,opt,name=nama_pelanggan,json=namaPelanggan,proto3" json:"nama_pelanggan,omitempty"`
	TanggalUlasan string  `protobuf:"bytes,3,opt,name=tanggal_ulasan,json=tanggalUlasan,proto3" json:"tanggal_ulasan,omitempty"`
	Komentar      string  `protobuf:"bytes,4,opt,name=komentar,proto3" json:"komentar,omitempty"`
	Rating        float32 `protobuf:"fixed32,5,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *Ulasan) Reset() {
	*x = Ulasan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_produk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ulasan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ulasan) ProtoMessage() {}

func (x *Ulasan) ProtoReflect() protoreflect.Message {
	mi := &file_protos_produk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ulasan.ProtoReflect.Descriptor instead.
func (*Ulasan) Descriptor() ([]byte, []int) {
	return file_protos_produk_proto_rawDescGZIP(), []int{3}
}

func (x *Ulasan) GetIdUlasan() uint64 {
	if x != nil {
		return x.IdUlasan
	}
	return 0
}

func (x *Ulasan) GetNamaPelanggan() string {
	if x != nil {
		return x.NamaPelanggan
	}
	return ""
}

func (x *Ulasan) GetTanggalUlasan() string {
	if x != nil {
		return x.TanggalUlasan
	}
	return ""
}

func (x *Ulasan) GetKomentar() string {
	if x != nil {
		return x.Komentar
	}
	return ""
}

func (x *Ulasan) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_produk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_protos_produk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_protos_produk_proto_rawDescGZIP(), []int{4}
}

func (x *Id) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_produk_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_protos_produk_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_protos_produk_proto_rawDescGZIP(), []int{5}
}

func (x *Status) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_protos_produk_proto protoreflect.FileDescriptor

var file_protos_produk_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xd2, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x61, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x12, 0x14, 0x0a, 0x05, 0x68, 0x61, 0x72, 0x67, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x68, 0x61, 0x72, 0x67, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x6f, 0x6b, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65,
	0x73, 0x6b, 0x72, 0x69, 0x70, 0x73, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x65, 0x73, 0x6b, 0x72, 0x69, 0x70, 0x73, 0x69, 0x12, 0x36, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x73,
	0x69, 0x66, 0x69, 0x6b, 0x61, 0x73, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x70, 0x65, 0x73, 0x69, 0x66, 0x69, 0x6b,
	0x61, 0x73, 0x69, 0x52, 0x0b, 0x73, 0x70, 0x65, 0x73, 0x69, 0x66, 0x69, 0x6b, 0x61, 0x73, 0x69,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x65, 0x72, 0x65, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6d, 0x65, 0x72, 0x65, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x61, 0x6e, 0x67, 0x67, 0x61,
	0x6c, 0x5f, 0x72, 0x69, 0x6c, 0x69, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74,
	0x61, 0x6e, 0x67, 0x67, 0x61, 0x6c, 0x52, 0x69, 0x6c, 0x69, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x06, 0x75, 0x6c, 0x61, 0x73, 0x61, 0x6e, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x6c,
	0x61, 0x73, 0x61, 0x6e, 0x52, 0x06, 0x75, 0x6c, 0x61, 0x73, 0x61, 0x6e, 0x22, 0xb2, 0x01, 0x0a,
	0x0b, 0x53, 0x70, 0x65, 0x73, 0x69, 0x66, 0x69, 0x6b, 0x61, 0x73, 0x69, 0x12, 0x25, 0x0a, 0x0e,
	0x69, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x73, 0x69, 0x66, 0x69, 0x6b, 0x61, 0x73, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x69, 0x64, 0x53, 0x70, 0x65, 0x73, 0x69, 0x66, 0x69, 0x6b,
	0x61, 0x73, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x6f, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x72, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x61,
	0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x6e, 0x79, 0x69, 0x6d, 0x70, 0x61, 0x6e, 0x61, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x6e, 0x79, 0x69, 0x6d, 0x70, 0x61,
	0x6e, 0x61, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x79, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x79, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x72, 0x61,
	0x66, 0x69, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x72, 0x61, 0x66, 0x69,
	0x73, 0x22, 0xa7, 0x01, 0x0a, 0x06, 0x55, 0x6c, 0x61, 0x73, 0x61, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x64, 0x5f, 0x75, 0x6c, 0x61, 0x73, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x69, 0x64, 0x55, 0x6c, 0x61, 0x73, 0x61, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x61, 0x6d,
	0x61, 0x5f, 0x70, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x67, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x61, 0x6d, 0x61, 0x50, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x67, 0x61, 0x6e,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x61, 0x6e, 0x67, 0x67, 0x61, 0x6c, 0x5f, 0x75, 0x6c, 0x61, 0x73,
	0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x61, 0x6e, 0x67, 0x67, 0x61,
	0x6c, 0x55, 0x6c, 0x61, 0x73, 0x61, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x14, 0x0a, 0x02, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0xfb, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x6b, 0x73, 0x12, 0x0e, 0x2e, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x2a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x6b, 0x12, 0x0b, 0x2e, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x1a,
	0x10, 0x2e, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x2d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x6b, 0x12, 0x10, 0x2e, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x1a, 0x0b, 0x2e, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64,
	0x12, 0x31, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x6b,
	0x12, 0x10, 0x2e, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x6b, 0x12, 0x0b, 0x2e, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64,
	0x1a, 0x0f, 0x2e, 0x67, 0x6f, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x6b, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_produk_proto_rawDescOnce sync.Once
	file_protos_produk_proto_rawDescData = file_protos_produk_proto_rawDesc
)

func file_protos_produk_proto_rawDescGZIP() []byte {
	file_protos_produk_proto_rawDescOnce.Do(func() {
		file_protos_produk_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_produk_proto_rawDescData)
	})
	return file_protos_produk_proto_rawDescData
}

var file_protos_produk_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_produk_proto_goTypes = []any{
	(*Empty)(nil),       // 0: go_grpc.Empty
	(*Product)(nil),     // 1: go_grpc.Product
	(*Spesifikasi)(nil), // 2: go_grpc.Spesifikasi
	(*Ulasan)(nil),      // 3: go_grpc.Ulasan
	(*Id)(nil),          // 4: go_grpc.Id
	(*Status)(nil),      // 5: go_grpc.Status
}
var file_protos_produk_proto_depIdxs = []int32{
	2, // 0: go_grpc.Product.spesifikasi:type_name -> go_grpc.Spesifikasi
	3, // 1: go_grpc.Product.ulasan:type_name -> go_grpc.Ulasan
	0, // 2: go_grpc.ProductClient.GetProduks:input_type -> go_grpc.Empty
	4, // 3: go_grpc.ProductClient.GetProduk:input_type -> go_grpc.Id
	1, // 4: go_grpc.ProductClient.CreateProduk:input_type -> go_grpc.Product
	1, // 5: go_grpc.ProductClient.UpdateProduk:input_type -> go_grpc.Product
	4, // 6: go_grpc.ProductClient.DeleteProduk:input_type -> go_grpc.Id
	1, // 7: go_grpc.ProductClient.GetProduks:output_type -> go_grpc.Product
	1, // 8: go_grpc.ProductClient.GetProduk:output_type -> go_grpc.Product
	4, // 9: go_grpc.ProductClient.CreateProduk:output_type -> go_grpc.Id
	5, // 10: go_grpc.ProductClient.UpdateProduk:output_type -> go_grpc.Status
	5, // 11: go_grpc.ProductClient.DeleteProduk:output_type -> go_grpc.Status
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_produk_proto_init() }
func file_protos_produk_proto_init() {
	if File_protos_produk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_produk_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_protos_produk_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Product); i {
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
		file_protos_produk_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Spesifikasi); i {
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
		file_protos_produk_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Ulasan); i {
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
		file_protos_produk_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Id); i {
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
		file_protos_produk_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_protos_produk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_produk_proto_goTypes,
		DependencyIndexes: file_protos_produk_proto_depIdxs,
		MessageInfos:      file_protos_produk_proto_msgTypes,
	}.Build()
	File_protos_produk_proto = out.File
	file_protos_produk_proto_rawDesc = nil
	file_protos_produk_proto_goTypes = nil
	file_protos_produk_proto_depIdxs = nil
}
