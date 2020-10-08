package main

import(
	//"log"
	//context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	//clienteLogistica "../clientelogistica/clientelogistica"
	"fmt"
	"strconv"
	csvventas "../csvventas"
	"time"
	//"io"
)

func main(){

	tiempopedidos:
		fmt.Println("ingrese tiempo entre pedidos en segundos")
		var TiempoPedidosString string
		fmt.Scanln(&TiempoPedidosString)
	var TiemposPedidos int
	if tt,err:=strconv.Atoi(TiempoPedidosString);err!=nil{
		goto tiempopedidos
	} else {
		TiemposPedidos=tt
	}

	var CsvVentas csvventas.CSVVentas


	var Mode string
	var filas [][]string
	mode:
		fmt.Println("¿Que tipo de cliente es usted? (ingrese 0 o 1)")
		fmt.Println("0: Pyme")
		fmt.Println("1: Retail")
		fmt.Scanln(&Mode)
	if(Mode=="0"){
		Mode="Pyme"
		CsvVentas:=csvventas.CSVVentas{NombreArchivo:"cliente/pymes.csv", TipoCliente:"Pyme"}
		CsvVentas.LeerPedidos()
		} else if (Mode=="1") {
		Mode="Retail"
		CsvVentas:=csvventas.CSVVentas{NombreArchivo:"cliente/retail.csv", TipoCliente:"Retail"}
		CsvVentas.LeerPedidos()
	} else {
		fmt.Println("error, ingrese de nuevo")
		goto mode
	}

	
	//esto debería ir después
	for _,fila:=range(filas){
		
		paquete:=CsvVentas.Pedido(fila)
		fmt.Println(paquete)
		time.Sleep(time.Second*time.Duration(TiemposPedidos))
	}

	//var conn *grpc.ClientConn
	//192.168.1.17:9000
	entry:
		fmt.Println("ingrese dirección IP del servidor (en el formato: 255.255.255.255)")
		var IPaddr string
		fmt.Scanln(&IPaddr)
		fmt.Println("ingrese el numero de puerto en el que el logística está escuchando")
		var PortNum string
		fmt.Scanln(&PortNum)

		CompleteAddr:=IPaddr+":"+PortNum
		fmt.Println(CompleteAddr)
		conn, err:=grpc.Dial(CompleteAddr,grpc.WithInsecure(),grpc.WithBlock())
		defer conn.Close()


	if err!=nil{
		goto entry
	}


	//c:=clienteLogistica.NewClienteLogisticaClient(conn)


	/*
	pedido1:=clienteLogistica.Pedido{
		IDPedido: "aaa",
		NombreProducto: "bbb",
		ValorProducto: 32,
		Tipo: "Retail",
		Destino: "Ciudad A",
		Origen: "Tienda A",
	}

	pedido2:=clienteLogistica.Pedido{
		IDPedido: "dde",
		NombreProducto: "eee",
		ValorProducto: 15,
		Tipo: "Normal",
		Destino: "Ciudad B",
		Origen: "Tienda B",
	}

	pedido3	:=clienteLogistica.Pedido{
		IDPedido: "fgf",
		NombreProducto: "ggg",
		ValorProducto: 18,
		Tipo: "Normal",
		Destino: "Ciudad C",
		Origen: "Tienda C",
	}

	pedido4:=clienteLogistica.Pedido{
		IDPedido: "efe",
		NombreProducto: "ggg",
		ValorProducto: 40,
		Tipo: "Prioritario",
		Destino: "Ciudad D",
		Origen: "Tienda D",
	}

	seguimiento, err:=c.HacerPedido(context.Background(), &pedido1)

	if err==nil{
		fmt.Println("pedido retorno seguimiento: ",seguimiento.CodigoSeguimiento)
	} else {
		log.Println("error: ",err)
	}

	seguimiento2, err:=c.HacerPedido(context.Background(), &pedido2)

	if err==nil{
		fmt.Println("pedido retorno seguimiento: ",seguimiento2.CodigoSeguimiento)
	} else {
		log.Println("error: ",err)
	}

	seguimiento3, err:=c.HacerPedido(context.Background(), &pedido3)

	if err==nil{
		fmt.Println("pedido retorno seguimiento: ",seguimiento3.CodigoSeguimiento)
	} else {
		log.Println("error: ",err)
	}

	seguimiento4, err:=c.HacerPedido(context.Background(), &pedido4)

	if err==nil{
		fmt.Println("pedido retorno seguimiento: ",seguimiento4.CodigoSeguimiento)
	} else {
		log.Println("error: ",err)
	}


	segreq:=clienteLogistica.Seguimiento{CodigoSeguimiento: 2}

	estado, err:=c.SolicitarEstado(context.Background(), &segreq)

	if err==nil{
		fmt.Println("estado:",estado.Estado)
	}
	*/



	

}