// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: camionlogistica/camionlogistica.proto

package camionlogistica

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Paquete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDPaquete         string `protobuf:"bytes,1,opt,name=IDPaquete,proto3" json:"IDPaquete,omitempty"`
	CodigoSeguimiento string `protobuf:"bytes,2,opt,name=CodigoSeguimiento,proto3" json:"CodigoSeguimiento,omitempty"`
	Tipo              string `protobuf:"bytes,3,opt,name=Tipo,proto3" json:"Tipo,omitempty"` //retail, normal (pyme), prioritario (pyme)
	ValorProducto     int32  `protobuf:"varint,4,opt,name=ValorProducto,proto3" json:"ValorProducto,omitempty"`
	Intentos          int32  `protobuf:"varint,5,opt,name=Intentos,proto3" json:"Intentos,omitempty"`
	Estado            string `protobuf:"bytes,6,opt,name=Estado,proto3" json:"Estado,omitempty"` //"En bodega", "En Camino", "Entregado", "No entregado"
	Origen            string `protobuf:"bytes,7,opt,name=Origen,proto3" json:"Origen,omitempty"`
	Destino           string `protobuf:"bytes,8,opt,name=Destino,proto3" json:"Destino,omitempty"`
}

func (x *Paquete) Reset() {
	*x = Paquete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_camionlogistica_camionlogistica_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paquete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paquete) ProtoMessage() {}

func (x *Paquete) ProtoReflect() protoreflect.Message {
	mi := &file_camionlogistica_camionlogistica_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paquete.ProtoReflect.Descriptor instead.
func (*Paquete) Descriptor() ([]byte, []int) {
	return file_camionlogistica_camionlogistica_proto_rawDescGZIP(), []int{0}
}

func (x *Paquete) GetIDPaquete() string {
	if x != nil {
		return x.IDPaquete
	}
	return ""
}

func (x *Paquete) GetCodigoSeguimiento() string {
	if x != nil {
		return x.CodigoSeguimiento
	}
	return ""
}

func (x *Paquete) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *Paquete) GetValorProducto() int32 {
	if x != nil {
		return x.ValorProducto
	}
	return 0
}

func (x *Paquete) GetIntentos() int32 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

func (x *Paquete) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *Paquete) GetOrigen() string {
	if x != nil {
		return x.Origen
	}
	return ""
}

func (x *Paquete) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

//cambio de planes, en lugar de retornar un stream de 2 paquetes, se retornan los
//2 paquetes al mismo tiempo en esta estructura.
//Son sólo 2 paquetes después de todo.
type ParPaquetes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paquete1     *Paquete `protobuf:"bytes,1,opt,name=Paquete1,proto3" json:"Paquete1,omitempty"`
	Paquete2     *Paquete `protobuf:"bytes,2,opt,name=Paquete2,proto3" json:"Paquete2,omitempty"`
	TipoCamion   string   `protobuf:"bytes,3,opt,name=TipoCamion,proto3" json:"TipoCamion,omitempty"`
	TiempoEspera int32    `protobuf:"varint,4,opt,name=TiempoEspera,proto3" json:"TiempoEspera,omitempty"`
}

func (x *ParPaquetes) Reset() {
	*x = ParPaquetes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_camionlogistica_camionlogistica_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParPaquetes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParPaquetes) ProtoMessage() {}

func (x *ParPaquetes) ProtoReflect() protoreflect.Message {
	mi := &file_camionlogistica_camionlogistica_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParPaquetes.ProtoReflect.Descriptor instead.
func (*ParPaquetes) Descriptor() ([]byte, []int) {
	return file_camionlogistica_camionlogistica_proto_rawDescGZIP(), []int{1}
}

func (x *ParPaquetes) GetPaquete1() *Paquete {
	if x != nil {
		return x.Paquete1
	}
	return nil
}

func (x *ParPaquetes) GetPaquete2() *Paquete {
	if x != nil {
		return x.Paquete2
	}
	return nil
}

func (x *ParPaquetes) GetTipoCamion() string {
	if x != nil {
		return x.TipoCamion
	}
	return ""
}

func (x *ParPaquetes) GetTiempoEspera() int32 {
	if x != nil {
		return x.TiempoEspera
	}
	return 0
}

//tipo de mensaje "dummy"
type Ok struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok int32 `protobuf:"varint,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
}

func (x *Ok) Reset() {
	*x = Ok{}
	if protoimpl.UnsafeEnabled {
		mi := &file_camionlogistica_camionlogistica_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ok) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ok) ProtoMessage() {}

func (x *Ok) ProtoReflect() protoreflect.Message {
	mi := &file_camionlogistica_camionlogistica_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ok.ProtoReflect.Descriptor instead.
func (*Ok) Descriptor() ([]byte, []int) {
	return file_camionlogistica_camionlogistica_proto_rawDescGZIP(), []int{2}
}

func (x *Ok) GetOk() int32 {
	if x != nil {
		return x.Ok
	}
	return 0
}

type Camion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDCamion   string `protobuf:"bytes,1,opt,name=IDCamion,proto3" json:"IDCamion,omitempty"`
	TipoCamion string `protobuf:"bytes,2,opt,name=TipoCamion,proto3" json:"TipoCamion,omitempty"`
}

func (x *Camion) Reset() {
	*x = Camion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_camionlogistica_camionlogistica_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Camion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Camion) ProtoMessage() {}

func (x *Camion) ProtoReflect() protoreflect.Message {
	mi := &file_camionlogistica_camionlogistica_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Camion.ProtoReflect.Descriptor instead.
func (*Camion) Descriptor() ([]byte, []int) {
	return file_camionlogistica_camionlogistica_proto_rawDescGZIP(), []int{3}
}

func (x *Camion) GetIDCamion() string {
	if x != nil {
		return x.IDCamion
	}
	return ""
}

func (x *Camion) GetTipoCamion() string {
	if x != nil {
		return x.TipoCamion
	}
	return ""
}

var File_camionlogistica_camionlogistica_proto protoreflect.FileDescriptor

var file_camionlogistica_camionlogistica_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x61, 0x2f, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x6c,
	0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x22, 0xf5, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x71,
	0x75, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x44, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x44, 0x50, 0x61, 0x71, 0x75, 0x65,
	0x74, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x43, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x53, 0x65, 0x67, 0x75,
	0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x43,
	0x6f, 0x64, 0x69, 0x67, 0x6f, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x70, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x54, 0x69, 0x70, 0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x56, 0x61, 0x6c,
	0x6f, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x49, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x4f, 0x72, 0x69, 0x67, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f,
	0x22, 0xbd, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x73,
	0x12, 0x34, 0x0a, 0x08, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x31, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x61, 0x2e, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x52, 0x08, 0x50, 0x61,
	0x71, 0x75, 0x65, 0x74, 0x65, 0x31, 0x12, 0x34, 0x0a, 0x08, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74,
	0x65, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x61, 0x6d, 0x69, 0x6f,
	0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x50, 0x61, 0x71, 0x75, 0x65,
	0x74, 0x65, 0x52, 0x08, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x32, 0x12, 0x1e, 0x0a, 0x0a,
	0x54, 0x69, 0x70, 0x6f, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x54, 0x69, 0x70, 0x6f, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x54, 0x69, 0x65, 0x6d, 0x70, 0x6f, 0x45, 0x73, 0x70, 0x65, 0x72, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x54, 0x69, 0x65, 0x6d, 0x70, 0x6f, 0x45, 0x73, 0x70, 0x65, 0x72, 0x61,
	0x22, 0x14, 0x0a, 0x02, 0x4f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x4f, 0x6b, 0x22, 0x44, 0x0a, 0x06, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x49, 0x44, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x49, 0x44, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x54, 0x69, 0x70, 0x6f, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x54, 0x69, 0x70, 0x6f, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x32, 0xe9, 0x01, 0x0a,
	0x0f, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61,
	0x12, 0x42, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x6f, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x67, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x1a, 0x13, 0x2e,
	0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e,
	0x4f, 0x6b, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0f, 0x41, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x72, 0x50,
	0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e,
	0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x50, 0x61, 0x72, 0x50, 0x61, 0x71,
	0x75, 0x65, 0x74, 0x65, 0x73, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x6c, 0x6f,
	0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x50, 0x61, 0x72, 0x50, 0x61, 0x71, 0x75, 0x65,
	0x74, 0x65, 0x73, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x72, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x63, 0x61, 0x6d, 0x69, 0x6f,
	0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e, 0x4f, 0x6b, 0x1a, 0x17, 0x2e,
	0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x2e,
	0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_camionlogistica_camionlogistica_proto_rawDescOnce sync.Once
	file_camionlogistica_camionlogistica_proto_rawDescData = file_camionlogistica_camionlogistica_proto_rawDesc
)

func file_camionlogistica_camionlogistica_proto_rawDescGZIP() []byte {
	file_camionlogistica_camionlogistica_proto_rawDescOnce.Do(func() {
		file_camionlogistica_camionlogistica_proto_rawDescData = protoimpl.X.CompressGZIP(file_camionlogistica_camionlogistica_proto_rawDescData)
	})
	return file_camionlogistica_camionlogistica_proto_rawDescData
}

var file_camionlogistica_camionlogistica_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_camionlogistica_camionlogistica_proto_goTypes = []interface{}{
	(*Paquete)(nil),     // 0: camionlogistica.Paquete
	(*ParPaquetes)(nil), // 1: camionlogistica.ParPaquetes
	(*Ok)(nil),          // 2: camionlogistica.Ok
	(*Camion)(nil),      // 3: camionlogistica.Camion
}
var file_camionlogistica_camionlogistica_proto_depIdxs = []int32{
	0, // 0: camionlogistica.ParPaquetes.Paquete1:type_name -> camionlogistica.Paquete
	0, // 1: camionlogistica.ParPaquetes.Paquete2:type_name -> camionlogistica.Paquete
	0, // 2: camionlogistica.CamionLogistica.ReportarIntento:input_type -> camionlogistica.Paquete
	1, // 3: camionlogistica.CamionLogistica.AsignarPaquetes:input_type -> camionlogistica.ParPaquetes
	2, // 4: camionlogistica.CamionLogistica.RegistrarCamion:input_type -> camionlogistica.Ok
	2, // 5: camionlogistica.CamionLogistica.ReportarIntento:output_type -> camionlogistica.Ok
	1, // 6: camionlogistica.CamionLogistica.AsignarPaquetes:output_type -> camionlogistica.ParPaquetes
	3, // 7: camionlogistica.CamionLogistica.RegistrarCamion:output_type -> camionlogistica.Camion
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_camionlogistica_camionlogistica_proto_init() }
func file_camionlogistica_camionlogistica_proto_init() {
	if File_camionlogistica_camionlogistica_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_camionlogistica_camionlogistica_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paquete); i {
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
		file_camionlogistica_camionlogistica_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParPaquetes); i {
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
		file_camionlogistica_camionlogistica_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ok); i {
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
		file_camionlogistica_camionlogistica_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Camion); i {
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
			RawDescriptor: file_camionlogistica_camionlogistica_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_camionlogistica_camionlogistica_proto_goTypes,
		DependencyIndexes: file_camionlogistica_camionlogistica_proto_depIdxs,
		MessageInfos:      file_camionlogistica_camionlogistica_proto_msgTypes,
	}.Build()
	File_camionlogistica_camionlogistica_proto = out.File
	file_camionlogistica_camionlogistica_proto_rawDesc = nil
	file_camionlogistica_camionlogistica_proto_goTypes = nil
	file_camionlogistica_camionlogistica_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CamionLogisticaClient is the client API for CamionLogistica service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CamionLogisticaClient interface {
	//En caso de que el intento haya sido exitoso, Estado="Entregado".
	//En el caso de que los 3 intentos hayan fallado, el paquete se manda
	//con estado="No entregado".
	//retorna un mensaje dummy "Ok"
	ReportarIntento(ctx context.Context, in *Paquete, opts ...grpc.CallOption) (*Ok, error)
	//Logística asigna paquetes al camión, tras recibir la info de los paquetes
	//entregados (o no entregados) y reportar al servicio financiero.
	AsignarPaquetes(ctx context.Context, in *ParPaquetes, opts ...grpc.CallOption) (*ParPaquetes, error)
	//una vez se inicializa un camion, se tiene que ejecutar esta funcion una y solo una vez
	//logistica entregará al camion informacion: ID del camion y el tipo (según el puerto al que se conecte)
	RegistrarCamion(ctx context.Context, in *Ok, opts ...grpc.CallOption) (*Camion, error)
}

type camionLogisticaClient struct {
	cc grpc.ClientConnInterface
}

func NewCamionLogisticaClient(cc grpc.ClientConnInterface) CamionLogisticaClient {
	return &camionLogisticaClient{cc}
}

func (c *camionLogisticaClient) ReportarIntento(ctx context.Context, in *Paquete, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/camionlogistica.CamionLogistica/ReportarIntento", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *camionLogisticaClient) AsignarPaquetes(ctx context.Context, in *ParPaquetes, opts ...grpc.CallOption) (*ParPaquetes, error) {
	out := new(ParPaquetes)
	err := c.cc.Invoke(ctx, "/camionlogistica.CamionLogistica/AsignarPaquetes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *camionLogisticaClient) RegistrarCamion(ctx context.Context, in *Ok, opts ...grpc.CallOption) (*Camion, error) {
	out := new(Camion)
	err := c.cc.Invoke(ctx, "/camionlogistica.CamionLogistica/RegistrarCamion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CamionLogisticaServer is the server API for CamionLogistica service.
type CamionLogisticaServer interface {
	//En caso de que el intento haya sido exitoso, Estado="Entregado".
	//En el caso de que los 3 intentos hayan fallado, el paquete se manda
	//con estado="No entregado".
	//retorna un mensaje dummy "Ok"
	ReportarIntento(context.Context, *Paquete) (*Ok, error)
	//Logística asigna paquetes al camión, tras recibir la info de los paquetes
	//entregados (o no entregados) y reportar al servicio financiero.
	AsignarPaquetes(context.Context, *ParPaquetes) (*ParPaquetes, error)
	//una vez se inicializa un camion, se tiene que ejecutar esta funcion una y solo una vez
	//logistica entregará al camion informacion: ID del camion y el tipo (según el puerto al que se conecte)
	RegistrarCamion(context.Context, *Ok) (*Camion, error)
}

// UnimplementedCamionLogisticaServer can be embedded to have forward compatible implementations.
type UnimplementedCamionLogisticaServer struct {
}

func (*UnimplementedCamionLogisticaServer) ReportarIntento(context.Context, *Paquete) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportarIntento not implemented")
}
func (*UnimplementedCamionLogisticaServer) AsignarPaquetes(context.Context, *ParPaquetes) (*ParPaquetes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AsignarPaquetes not implemented")
}
func (*UnimplementedCamionLogisticaServer) RegistrarCamion(context.Context, *Ok) (*Camion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistrarCamion not implemented")
}

func RegisterCamionLogisticaServer(s *grpc.Server, srv CamionLogisticaServer) {
	s.RegisterService(&_CamionLogistica_serviceDesc, srv)
}

func _CamionLogistica_ReportarIntento_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Paquete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CamionLogisticaServer).ReportarIntento(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/camionlogistica.CamionLogistica/ReportarIntento",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CamionLogisticaServer).ReportarIntento(ctx, req.(*Paquete))
	}
	return interceptor(ctx, in, info, handler)
}

func _CamionLogistica_AsignarPaquetes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParPaquetes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CamionLogisticaServer).AsignarPaquetes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/camionlogistica.CamionLogistica/AsignarPaquetes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CamionLogisticaServer).AsignarPaquetes(ctx, req.(*ParPaquetes))
	}
	return interceptor(ctx, in, info, handler)
}

func _CamionLogistica_RegistrarCamion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ok)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CamionLogisticaServer).RegistrarCamion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/camionlogistica.CamionLogistica/RegistrarCamion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CamionLogisticaServer).RegistrarCamion(ctx, req.(*Ok))
	}
	return interceptor(ctx, in, info, handler)
}

var _CamionLogistica_serviceDesc = grpc.ServiceDesc{
	ServiceName: "camionlogistica.CamionLogistica",
	HandlerType: (*CamionLogisticaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportarIntento",
			Handler:    _CamionLogistica_ReportarIntento_Handler,
		},
		{
			MethodName: "AsignarPaquetes",
			Handler:    _CamionLogistica_AsignarPaquetes_Handler,
		},
		{
			MethodName: "RegistrarCamion",
			Handler:    _CamionLogistica_RegistrarCamion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "camionlogistica/camionlogistica.proto",
}
