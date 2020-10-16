package clientelogistica

import(
	context "context"
	csvo "../../csvordenes"
	colas "../../colas"
	registroseguimiento "../../registroseguimiento"
	"time"
	"fmt"
)


type Cliente_Logistica_Server struct{
	CsvOrdenes *csvo.CSVOrdenes
	ColasPedidos *colas.Colas
	SeguimientoActual *int
	RegistrosSeguimientos *[]registroseguimiento.RegistroSeguimiento

}

func(cls *Cliente_Logistica_Server)HacerPedido(ctx context.Context, pedido *Pedido) (*Seguimiento, error){
	
	//fmt.Println("/npedido recibido a logistica: ",pedido)

	//crear codigo de seguimiento
	var codigoSeguimiento int
	if pedido.GetTipo()!="Retail"{
		codigoSeguimiento=*(cls.SeguimientoActual)
		//fmt.Println("codigo de seguimiento creado: ",codigoSeguimiento)
		*(cls.SeguimientoActual)=*(cls.SeguimientoActual)+1//habria que poner un semaforo para asegurar que no hayan errores, pero filo
	} else {
		codigoSeguimiento=0
		//fmt.Println("sin codigo de seguimiento creado: ",codigoSeguimiento)
	}

	//crear paquete
	//fmt.Println("creando paquete")
	paquete:=colas.Paquete{
		IDPedido: pedido.GetIDPedido(),
		NombreProducto: pedido.GetNombreProducto(),
		ValorProducto: pedido.GetValorProducto(),
		Tipo: pedido.GetTipo(),
		Destino: pedido.GetDestino(),
		Origen: pedido.GetOrigen(),
		CodigoSeguimiento: codigoSeguimiento,
	}
	//fmt.Println("paquete creado:",paquete)
	//fmt.Println("creando registro seguimiento ")

	
	
	//ingresar al registro de seguimientos
	registroSeguimiento:=registroseguimiento.RegistroSeguimiento{
		IDPaquete: pedido.GetIDPedido(),
		Estado: "En Bodega",
		IDCamion: "Ninguno Aún",
		IDSeguimiento: codigoSeguimiento,
		CantidadIntentos: 0,
	}


	fmt.Println("agregando a registros")
	*(cls.RegistrosSeguimientos)=append(*(cls.RegistrosSeguimientos), registroSeguimiento)
	//fmt.Println("registro de seguimientos:",*(cls.RegistrosSeguimientos))

	fmt.Println("agregando a colas")
	//fmt.Println((*cls.ColasPedidos).ColaNormal)
	//fmt.Println(cls.ColasPedidos)


	//ingresar a la cola
	switch paquete.Tipo{
	case "Normal":
		//fmt.Println("(cli)agregando a cola normal")
		(*(*(*cls).ColasPedidos).ColaNormal)=append((*(*(*cls).ColasPedidos).ColaNormal),paquete)
	case "Prioritario":
		//fmt.Println("(cli)agregando a cola prioritaria")
		(*(*(*cls).ColasPedidos).ColaPrioritaria)=append((*(*(*cls).ColasPedidos).ColaPrioritaria),paquete)
	case "Retail":
		//fmt.Println("(cli)agregando a cola retail")
		(*(*(*cls).ColasPedidos).ColaRetail)=append((*(*(*cls).ColasPedidos).ColaRetail),paquete)
	}
	
	//fmt.Println("cola retail: ",*(*(*cls).ColasPedidos).ColaRetail)



	//ingresar al csv, una orden nueva:
	//fmt.Println("actualizando csv de registros de pedidos:")
	orden:=csvo.Orden{
		TimeStamp: time.Now().Format("2-Jan-2006 3:04PM"),
		IDPaquete: pedido.GetIDPedido(),
		Tipo: pedido.GetTipo(),
		Nombre: pedido.GetNombreProducto(),
		Valor: pedido.GetValorProducto(),
		Origen: pedido.GetOrigen(),
		Destino: pedido.GetDestino(),
		Seguimiento: codigoSeguimiento,
	}
	//fmt.Println("añadiendo orden")
	cls.CsvOrdenes.AñadirOrden(orden)
	//fmt.Println("registro csv actualizado")

	//retornar codigo de seguimiento
	//fmt.Println("retornando codigo de seguimiento: ",codigoSeguimiento)
	//fmt.Println("retornando codigo de seguimiento: ",int64(codigoSeguimiento))
	return &Seguimiento{CodigoSeguimiento:int64(codigoSeguimiento)}, nil
}


func(cls *Cliente_Logistica_Server)ImprimirColas(){
	(*cls).ColasPedidos.ImprimirColas()
	
}

func(cls *Cliente_Logistica_Server)SolicitarEstado(ctx context.Context, seguimiento *Seguimiento) (*Estado, error){
	
	codigoseg:=seguimiento.GetCodigoSeguimiento()

	if codigoseg==0{
		return &Estado{Estado:"No disponible"}, nil
	}

	for _,seg:=range(*(cls.RegistrosSeguimientos)){
		if int64(seg.IDSeguimiento)==codigoseg{
			//fmt.Println("producto encontrado: ",seg.IDPaquete)
			return &Estado{Estado: seg.Estado},nil
		}
	}
	
	return &Estado{Estado:"No encontrado"}, nil
}