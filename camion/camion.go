package main

import(
	"log"
	//firstGrpc "../interfaces"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	camionLogistica "../camionlogistica/camionlogistica"
	"fmt"
	//"strconv"
	//"io"
)

var tipoCamion string

func main(){
	//aca tengo que ingresar el numero de puerto que se ingresó a logistica (para camiones)
	entry:
		fmt.Println("ingrese dirección IP del servidor (en el formato: 255.255.255.255)")
		var IPaddr string
		fmt.Scanln(&IPaddr)
		fmt.Println("ingrese el numero de puerto en el que el logística está escuchando un camión")
		var PortNum string
		fmt.Scanln(&PortNum)

		CompleteAddr:=IPaddr+":"+PortNum
		fmt.Println(CompleteAddr)
		conn, err:=grpc.Dial(CompleteAddr,grpc.WithInsecure(),grpc.WithBlock())
		defer conn.Close()

	if err!=nil{
		goto entry
	}
	
	
	c:=camionLogistica.NewCamionLogisticaClient(conn)

	ok:=camionLogistica.Ok{Ok:int32(0)}

	//lo primero que hay que hacer obligatoriamente. Se nos asigna el tipo de camión segun el puerto
	CamionRes, err:=c.RegistrarCamion(context.Background(),&ok)

	if err!=nil{
		log.Println("error: ",err)
		return
	} 
	tipoCamion=CamionRes.GetTipoCamion()
	fmt.Println("este camion es de tipo ",tipoCamion)


		


}

