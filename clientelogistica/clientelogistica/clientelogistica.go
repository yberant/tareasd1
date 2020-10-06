package clientelogistica

import(
	context "context"
	csvo "../../logistica/csvordenes"
	colas "../../logistica/colas"
	"fmt"

)


type Cliente_Logistica_Server struct{
	CsvOrdenes *csvo.CSVOrdenes
	ColasPedidos *colas.Colas
	SeguimientoActual *int
}

func(cls *Cliente_Logistica_Server)HacerPedido(ctx context.Context, pedido *Pedido) (*Seguimiento, error){
	
	//crear paquete
	paquete:=colas.Paquete{
		IDPedido: pedido.GetIDPedido(),
		NombreProducto: pedido.GetNombreProducto(),
		ValorProducto: pedido.GetValorProducto(),
		Tipo: pedido.GetTipo(),
		Destino: pedido.GetDestino(),
		Origen: pedido.GetOrigen(),
	}
	//ingresar al csv
	//ingresar a la cola
	switch paquete.Tipo{
	case "Normal":
		(*(*(*cls).ColasPedidos).ColaNormal)=append((*(*(*cls).ColasPedidos).ColaNormal),paquete)
	case "Prioritaria":
		(*(*(*cls).ColasPedidos).ColaPrioritaria)=append((*(*(*cls).ColasPedidos).ColaPrioritaria),paquete)
	case "Retail":
		(*(*(*cls).ColasPedidos).ColaRetail)=append((*(*(*cls).ColasPedidos).ColaRetail),paquete)
	}
	//retornar codigo de seguimiento

	
	
	return &Seguimiento{CodigoSeguimiento:0}, nil
}

func(cls *Cliente_Logistica_Server)ImprimirColas(){
	fmt.Println("colas")
	fmt.Println("normal: ",*(*(*cls).ColasPedidos).ColaNormal)
	fmt.Println("priotritaria: ",*(*(*cls).ColasPedidos).ColaPrioritaria)
	fmt.Println("retail: ",*(*(*cls).ColasPedidos).ColaRetail)
	
}

func(cls *Cliente_Logistica_Server)SolicitarEstado(ctx context.Context, seguimiento *Seguimiento) (*Estado, error){
	return &Estado{Estado:"no implementado"}, nil
}