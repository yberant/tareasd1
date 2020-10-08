package main

import (
    "log"
	"net"
	"fmt"
	"strconv"
	grpc "google.golang.org/grpc"
	clienteLogistica "../clientelogistica/clientelogistica"
	camionLogistica "../camionlogistica/camionlogistica"
	"time"
	csvo "../csvordenes"
	colas "../colas"
	registroseguimiento "../registroseguimiento"
    //"strings"
)

//variables
var(
	IPAddr string
	waitSeconds int
	seguiActual int
	colasPaquetes colas.Colas
	csvOrdenes csvo.CSVOrdenes
	registrosSeguimientos []registroseguimiento.RegistroSeguimiento
)



//jijiji
func getIPAddr() string{
	addrs, err := net.InterfaceAddrs()
    if err != nil {
        return ""
    }
    for _, address := range addrs {
        // check the address type and if it is not a loopback the display it
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String()
            }
        }
    }
    return ""
}

func ListenClientes(clientPort int){
	
	
	portstring:=":"+strconv.Itoa(clientPort)//por ejemplo, ":9000"
	lis, err := net.Listen("tcp", portstring)
	if err!=nil{
		log.Fatalf("Error escuchando en el puerto :%s: %v", portstring, err)
	} else{
		fmt.Println("Escuchado clientes desde: ",IPAddr+portstring)
	}

	//fmt.Println("agregando colas paquetes ",colasPaquetes)
		

	s:=clienteLogistica.Cliente_Logistica_Server{
		CsvOrdenes: &csvOrdenes,
		SeguimientoActual: &seguiActual,
		ColasPedidos: &colasPaquetes,
		RegistrosSeguimientos: &registrosSeguimientos,
	}
	


	grpcServer:=grpc.NewServer()
	clienteLogistica.RegisterClienteLogisticaServer(grpcServer,&s)



	if err:=grpcServer.Serve(lis); err!=nil{
		log.Fatalf("No se pudo servir en grpc en el puerto: %s; %v",portstring, err)
	} else {
		fmt.Println("Servidor comunicandose con cliente")
	}	
	
	
}

func EscucharCamion (camionPort int, camionCount int){
	portstring:=":"+strconv.Itoa(camionPort)//por ejemplo, ":9000"
	lis, err := net.Listen("tcp", portstring)
	if err!=nil{
		log.Fatalf("Error escuchando en el puerto :%s: %v", portstring, err)
	} else{
		fmt.Println("Escuchado a un camion desde: ",IPAddr+portstring)
	}

	s:=camionLogistica.Camion_Logistica_Server{
		CamionCount: camionCount, 
		ColasPaquetes: &colasPaquetes,
		RegistrosSeguimientos:&registrosSeguimientos,
	}


	grpcServer:=grpc.NewServer()
	camionLogistica.RegisterCamionLogisticaServer(grpcServer,&s)



	if err:=grpcServer.Serve(lis); err!=nil{
		log.Fatalf("No se pudo servir en grpc en el puerto: %s; %v",portstring, err)
	} else {
		fmt.Println("Servidor comunicandose con cliente")
	}

}

//Escucha a un camión
func AñadirCamion(camionCount int){
	enterNumber:
		switch camionCount{
		case 0:
			fmt.Println("Ingrese el numero de puerto del camion normal")
		case 1:
			fmt.Println("Ingrese el numero de puerto del primer camión de retail")
		case 2:
			fmt.Println("Ingrese el numero de puerto del segundo camión de retail")
		}
		var PortString string
		fmt.Scanln(&PortString)
	if port,err:=strconv.Atoi(PortString);err!=nil{
			fmt.Println("numero inválido")
			goto enterNumber
	} else {
		go EscucharCamion(port, camionCount)	
		time.Sleep(100*time.Millisecond)
	}	
}
//func listenCamiones(){...}



func main(){
	//creacion de csv
	csvOrdenes=csvo.CSVOrdenes{FileName: "logistica/logs.csv"}
	csvOrdenes.CrearArchivo()

	//creacion de registro de seguimiento
	registrosSeguimientos=[]registroseguimiento.RegistroSeguimiento{}

	_ = registrosSeguimientos

	//creacion de colas
	colasPaquetes=colas.Colas{
		ColaNormal: &([]colas.Paquete{}),
		ColaPrioritaria: &([]colas.Paquete{}),
		ColaRetail: &([]colas.Paquete{}),
	}

	_ = colasPaquetes
	colasPaquetes.ImprimirColas()
	//colasPaquetes.ImprimirColas()

	seguiActual=1;

	IPAddr=getIPAddr()
	//log.Println(IpAddr)
	fmt.Println("dirección IP de logística: ",IPAddr)

	//se registran los 3 camiones
	AñadirCamion(0)
	AñadirCamion(1)
	AñadirCamion(2)
	
	for {
		fmt.Println("comandos:")
		fmt.Println("\"quit\": salir")
		fmt.Println("\"listen\": agregar nuevo puerto para escuchar a cliente")

		var Command string
		fmt.Scanln(&Command)

		switch Command{
		case "quit":
			fmt.Println("adios")
			return

		case "listen":
			enterNumber:
				fmt.Println("Ingrese numero de puerto a escuchar")
				var PortString string
				fmt.Scanln(&PortString)
			if port,err:=strconv.Atoi(PortString);err!=nil{
				fmt.Println("numero inválido")
				goto enterNumber
			} else {
				go ListenClientes(port)
				time.Sleep(500 * time.Millisecond)

			}
			
		}

		


	}
}