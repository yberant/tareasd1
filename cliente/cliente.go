package main

import(
	"log"
	//firstGrpc "../interfaces"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	clienteLogistica "../clientelogistica/clientelogistica"
	"fmt"
	//"strconv"
	//"io"
)

func main(){
	//var conn *grpc.ClientConn
	//192.168.1.17:9000
	entry:
		fmt.Println("ingrese dirección IP del servidor (en el formato: 255.255.255.255)")
		var IPaddr string
		fmt.Scanln(&IPaddr)
		fmt.Println("ingrese el numero de puerto en el que el servidor está escuchando")
		var PortNum string
		fmt.Scanln(&PortNum)

		CompleteAddr:=IPaddr+":"+PortNum
		fmt.Println(CompleteAddr)
		conn, err:=grpc.Dial(CompleteAddr,grpc.WithInsecure(),grpc.WithBlock())
		defer conn.Close()


	if err!=nil{
		goto entry
	}


	c:=clienteLogistica.NewClienteLogisticaClient(conn)



	pedido:=clienteLogistica.Pedido{
		IDPedido: "aaa",
		NombreProducto: "bbb",
		ValorProducto: 32,
		Tipo: "Retail",
		Destino: "Ciudad A",
		Origen: "Tienda A",
	}

	seguimiento, err:=c.HacerPedido(context.Background(), &pedido)

	fmt.Println("OOOO")


	if err!=nil{
		log.Fatalf("something went wrong with the response 1")
	} else {
		fmt.Println("respuesta recibida")
		fmt.Println(seguimiento.CodigoSeguimiento)
	}

}