package colas

// casi lo mismo que un "Pedido", pero por problemas de importaciones era mas facil hacer otra struct
type Paquete struct{
	IDPedido string
	NombreProducto string
	ValorProducto int32
	Tipo string
	Destino string
	Origen string
}

type Colas struct{
	ColaNormal *[]Paquete
	ColaPrioritaria *[]Paquete
	ColaRetail *[]Paquete
}
