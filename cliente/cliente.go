package main

import(
	"log"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	clienteLogistica "../clientelogistica/clientelogistica"
	"fmt"
	"strconv"
	csvventas "../csvventas"
	"strings"
	"time"
	//"io"
)

func StringToPedido(fila []string)(clienteLogistica.Pedido){
	var pedido clienteLogistica.Pedido

	if TipoCliente=="Retail"{

		val,err:=strconv.Atoi(fila[2])
		if err!=nil{
			log.Fatalf("Error en columna de valor: %s\n", err.Error())
		}
		pedido=clienteLogistica.Pedido{
			IDPedido:fila[0],
			NombreProducto:fila[1],
			ValorProducto:int32(val),
			Tipo: "Retail",
			Origen: fila[3],
			Destino: fila[4],
		}
	} else {//tipo Pyme

		var tipoproducto string
		if fila[5]=="0"{
			tipoproducto="Normal"
		} else {
			tipoproducto="Prioritario"
		}
		val,err:=strconv.Atoi(fila[2])
		if err!=nil{
			log.Fatalf("Error en columna de valor: %s\n", err.Error())
		}

		var origen string
		if f:=strings.Split(fila[4],"-");len(f)==2{//ej: "casa-A"
			origen="tienda-"+f[1]//ej: "tienda-A"
		} else {
			origen="tienda-?"
		}
		pedido=clienteLogistica.Pedido{
			IDPedido:fila[0],
			NombreProducto:fila[1],
			ValorProducto:int32(val),
			Tipo: tipoproducto,
			Origen: origen,
			Destino: fila[4],
		}
	}

	return pedido

}

func ModoSeguimiento(){
	fmt.Println("modo seguimiento!")
	var codS string
	var codigoSeguimiento int64
	var seguimiento clienteLogistica.Seguimiento
	for{
		codigo:
			fmt.Println("ingrese nuémro de seguimiento para solicitar")
			fmt.Scanln(&codS)
		if cs,err:=strconv.Atoi(codS);err!=nil{
			fmt.Println("ingrese número valido")
			goto codigo
		} else {
			codigoSeguimiento=int64(cs)
		}
		seguimiento=clienteLogistica.Seguimiento{
			CodigoSeguimiento:codigoSeguimiento,
		}
		est,err:=c.SolicitarEstado(context.Background(),&seguimiento)

		if err!=nil{
			log.Fatalf("error en seguimiento recibido: %s",err)
		}

		estado:=est.Estado
		if(estado=="No disponible"){
			fmt.Println("codigo de seguimiento inválido")
		} else if(estado=="No encontrado"){
			fmt.Println("producto no encontrado, intente con otro código")
		} else{
			fmt.Println("estado del producto: ",estado)
		}



	}

}

var TipoCliente string
var c clienteLogistica.ClienteLogisticaClient

func main(){

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

	c=clienteLogistica.NewClienteLogisticaClient(conn)

	var Mode string
	var filas [][]string
	var CsvVentas csvventas.CSVVentas
	mode:
		fmt.Println("¿Que tipo de cliente es usted? (ingrese 0, 1 o 2)")
		fmt.Println("0: Pyme")
		fmt.Println("1: Retail")
		fmt.Println("2: Seguidor")
		fmt.Scanln(&Mode)
	if(Mode=="0"){
		TipoCliente="Pyme"
		CsvVentas=csvventas.CSVVentas{NombreArchivo:"cliente/pymes.csv"}
		} else if (Mode=="1") {
		TipoCliente="Retail"
		CsvVentas=csvventas.CSVVentas{NombreArchivo:"cliente/retail.csv"}
		} else if (Mode=="2") {
			ModoSeguimiento()
	} else {
		fmt.Println("error, ingrese de nuevo")
		goto mode
	}
	if(Mode!="2"){
		filas=CsvVentas.LeerPedidos()
	}

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


	//esto debería ir después
	for i,fila:=range(filas){
		if i==0{
			continue//headers del csv no se cuentan
		}
		//fmt.Println("fila de largo ",len(fila),": ",fila," ",i)
		pedido:=StringToPedido(fila)
		fmt.Println("\nrealizando pedido: ",pedido)
		seg, err:=c.HacerPedido(context.Background(),&pedido)
		if err!=nil{
			log.Fatalf("Reuqest error: %s",err)
		}
		fmt.Println("codigo de seguimiento recibido por logistica: ",seg.CodigoSeguimiento)
		time.Sleep(time.Second*time.Duration(TiemposPedidos))
	}


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