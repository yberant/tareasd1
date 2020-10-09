package colas

import(
	"fmt"
)
// existe también una estructura de mismo nombre y con (casi) los mismos atributos definida en el proto de camionlogistica
//, pero por problemas de importaciones era mas facil hacer otra struct
type Paquete struct{
	IDPedido string
	NombreProducto string
	ValorProducto int32
	Tipo string
	Destino string
	Origen string
	CodigoSeguimiento int
}

/*
IDPaquete         string `protobuf:"bytes,1,opt,name=IDPaquete,proto3" json:"IDPaquete,omitempty"`
	CodigoSeguimiento string `protobuf:"bytes,2,opt,name=CodigoSeguimiento,proto3" json:"CodigoSeguimiento,omitempty"`
	Tipo              string `protobuf:"bytes,3,opt,name=Tipo,proto3" json:"Tipo,omitempty"` //retail, normal (pyme), prioritario (pyme)
	ValorProducto     int32  `protobuf:"varint,4,opt,name=ValorProducto,proto3" json:"ValorProducto,omitempty"`
	Intentos          int32  `protobuf:"varint,5,opt,name=Intentos,proto3" json:"Intentos,omitempty"`
	Estado            string `protobuf:"bytes,6,opt,name=Estado,proto3" json:"Estado,omitempty"` //"En bodega", "En Camino", "Entregado", "No entregado"
	Origen            string `protobuf:"bytes,7,opt,name=Origen,proto3" json:"Origen,omitempty"`
	Destino 
*/

type Colas struct{
	ColaNormal *[]Paquete
	ColaPrioritaria *[]Paquete
	ColaRetail *[]Paquete
}

func (cls *Colas) ImprimirColas(){
	fmt.Println("colas")
	fmt.Println("normal: ",len(*(*cls).ColaNormal))
	fmt.Println("priotritaria: ",len(*(*cls).ColaPrioritaria))
	fmt.Println("retail: ",len(*(*cls).ColaRetail))
}
