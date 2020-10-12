package main

import(
	"log"
	//firstGrpc "../interfaces"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	camionLogistica "../camionlogistica/camionlogistica"
	"fmt"
	//"io"
	"encoding/csv"
	"math/rand"
	"os"
	"strconv"
	"time"

)

var tipoCamion string
var csvPath string
var c camionLogistica.CamionLogisticaClient



//esta funcion actualiza el registro del camión
func actualizarRegistro(par *camionLogistica.ParPaquetes, data [2][7]string, cantidad int, path string) {
	data[0] = [7]string{par.Paquete1.IDPaquete, par.Paquete1.Tipo, strconv.FormatInt(int64(par.Paquete1.ValorProducto), 10), par.Paquete1.Origen, par.Paquete1.Destino, strconv.FormatInt(int64(par.Paquete1.Intentos), 10), strconv.Itoa(0)}
	//se debe escribir la data
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	//se chequea por errores al escribir
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Se agrego correctamente la información en el registro csv")
	if cantidad == 2 {
		//hay dos paquetes por lo que hay que agregar otro a data
		data[1] = [7]string{par.Paquete2.IDPaquete, par.Paquete2.Tipo, strconv.FormatInt(int64(par.Paquete2.ValorProducto), 10), par.Paquete2.Origen, par.Paquete2.Destino, strconv.FormatInt(int64(par.Paquete2.Intentos), 10), strconv.Itoa(0)}
		writer.Write(data[0][:])
		writer.Write(data[1][:])
		writer.Flush()
		if err := writer.Error(); err != nil {
			// an error occurred during the flush
			fmt.Println("Ocurrio un error al hacer el flush")
		}
	} else {
		writer.Write(data[0][:])
		writer.Flush()
		if err := writer.Error(); err != nil {
			// an error occurred during the flush
			fmt.Println("Ocurrio un error al hacer el flush")
		}
	}
}

//esta función simulará la entrega de los paquetes, 0 significa que lo entrega, 1 que fallo la entrega
func repartir(paquete *camionLogistica.Paquete) (respuesta int) {
	probability := []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1}
	//se suma 1 al intento, partiendo de 0
	paquete.Intentos = paquete.Intentos + 1
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	if probability[r.Intn(len(probability))] == 0 {
		fmt.Println("intento exitoso")
		return 0
	} else {
		fmt.Println("intento fallido")
		return 1
	}
}



//esto debería funcionar para dos camiones
func entregaPaquete(paquete *camionLogistica.Paquete, csvPath string, tiempo int){

	//en este caso, o ya fue entregado o se agotaron los intentos
	if paquete.Estado!="En Camino" {
		return
	}

	if paquete.Tipo=="Normal"{
		//tope de 2 REintentos (3 intentos maximos en total)
		if paquete.Intentos < 3 && float32((paquete.Intentos)*10) < float32(paquete.ValorProducto) {//YOEL: y los intentos??????
			fmt.Println("intentando entregar paquete "+paquete.IDPaquete+" de tipo normal")
			time.Sleep(time.Duration(tiempo) * time.Second)

			respuesta := repartir(paquete)
			cambiarIntentos(csvPath, paquete)

			if respuesta == 0 {
				recibir(paquete)
				cambiarFechaEntrega(csvPath, paquete)
				return
			} 

			if paquete.Intentos>=3 || float32((paquete.Intentos)*10) >= float32(paquete.ValorProducto){
				noRecibir(paquete)
				cambiarFechaEntrega(csvPath, paquete)
			}

		} else {
			noRecibir(paquete)
			cambiarFechaEntrega(csvPath, paquete)
		}
		return

	} else if paquete.Tipo=="Prioritario" {

		if paquete.Intentos < 3 && float32((paquete.Intentos)*10) < 1.3*float32(paquete.ValorProducto) {
			fmt.Println("intentando entregar paquete "+paquete.IDPaquete+" de tipo prioritario")
			time.Sleep(time.Duration(tiempo) * time.Second)
			
			respuesta := repartir(paquete)
			cambiarIntentos(csvPath, paquete)

			if respuesta == 0 {
				recibir(paquete)
				cambiarFechaEntrega(csvPath, paquete)
				return
			}

			if paquete.Intentos>=3 || float32((paquete.Intentos)*10) >= 1.3*float32(paquete.ValorProducto){
				noRecibir(paquete)
				cambiarFechaEntrega(csvPath, paquete)
			}

		} else {
			noRecibir(paquete)
			cambiarFechaEntrega(csvPath, paquete)
		}
		return

	} else {//"Retail"

		if paquete.Intentos < 3 {
			
			fmt.Println("intentando entregar paquete "+paquete.IDPaquete+" de tipo retail")
			time.Sleep(time.Duration(tiempo) * time.Second)

			respuesta := repartir(paquete)
			cambiarIntentos(csvPath, paquete)

			if respuesta == 0 {
				recibir(paquete)
				cambiarFechaEntrega(csvPath, paquete)
				return
			}

			if paquete.Intentos>=3{
				noRecibir(paquete)
				cambiarFechaEntrega(csvPath, paquete)
			}	
			c.ReportarIntento(context.Background(), paquete)

		} else {
			noRecibir(paquete)
			cambiarFechaEntrega(csvPath, paquete)
			//return
		}
		return

	}
}

//cambia estado a recibido
func recibir(paquete *camionLogistica.Paquete) {
	fmt.Println("paquete "+paquete.IDPaquete+" recibido!")
	paquete.Estado = "Recibido"
}

//cambia estado a no recibido
func noRecibir(paquete *camionLogistica.Paquete) {
	fmt.Println("paquete "+paquete.IDPaquete+" no pudo ser recibido!")
	paquete.Estado = "No Recibido"
}

//escribe en el registro csv que se hizo un intento
func cambiarIntentos(path string, paquete *camionLogistica.Paquete) {
	csvFile, err1 := os.Open(path)
	if err1 != nil {
		panic(err1)
	}

	csvReader := csv.NewReader(csvFile)
	rows, err := csvReader.ReadAll() // `rows` is of type [][]string
	if err != nil {
		panic(err)
	}
	//fmt.Println(rows)
	//fmt.Println("\n")
	var id = paquete.IDPaquete
	for i, row := range rows {
		// process the `row` here
		if row[0] == id {
			//se encuentra el registro que se quiere cambiar.
			//se actualiza la cantidad de intentos.
			rows[i][5] = strconv.FormatInt(int64(paquete.Intentos), 10)
			csvFile.Close()
			break
		}
	}
	//se debe escribir nuevamente el archivo
	csvFile, err2 := os.Create(path)
	if err2 != nil {
		panic(err2)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	if err := csvWriter.WriteAll(rows); err != nil {
		log.Fatal(err)
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		// an error occurred during the flush
	}
}

//una vez que se entrego el paquete, esta funcion cambia a la fecha en la que se entrego
func cambiarFechaEntrega(path string, paquete *camionLogistica.Paquete) {
	csvFile, err1 := os.Open(path)
	if err1 != nil {
		panic(err1)
	}

	csvReader := csv.NewReader(csvFile)
	rows, err := csvReader.ReadAll() // `rows` is of type [][]string
	if err != nil {
		panic(err)
	}
	var id = paquete.IDPaquete
	t := time.Now()
	loc, err3 := time.LoadLocation("America/Buenos_Aires")

	if err3 != nil {
		fmt.Println("Opps: ", err3)
		return
	}
	fecha := t.In(loc).Format("02-01-2006 03:04")
	fechaEntrega := fecha + "\n"
	for i, row := range rows {
		// process the `row` here
		if row[0] == id {
			//se encuentra el registro que se quiere cambiar.
			//se actualiza la cantidad de intentos.
			rows[i][6] = fechaEntrega
			csvFile.Close()
			break
		}
	}
	//se debe escribir nuevamente el archivo
	csvFile, err2 := os.Create(path)
	if err2 != nil {
		panic(err2)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	if err := csvWriter.WriteAll(rows); err != nil {
		log.Fatal(err)
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		// an error occurred during the flush
	}
}


func main(){
	//data v a contener la información de los paquetes a escribir en el registro
	var data [2][7]string
	// cantidad indicará si se recibieron dos o un paquete
	var cantidad int
	//numero del paquete a entregar, se inicializa en cero para mostrar que es la primera iteración
	
	//aca tengo que ingresar el numero de IP y puerto que se ingresó a logistica (para camiones)
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
	
	//cliente que puede utilizar la interfaz camionlogistica, llamando funciones de forma remota
	c=camionLogistica.NewCamionLogisticaClient(conn)

	ok:=camionLogistica.Ok{Ok:int32(0)}

	//lo primero que hay que hacer obligatoriamente. Se nos asigna el tipo de camión segun el puerto
	CamionRes, err:=c.RegistrarCamion(context.Background(),&ok)
	
	fmt.Println("este camion es de tipo: ",CamionRes.GetTipoCamion())


	//crear csv que almacene los campos de los paquetes
	csvPath = "./camion/registro"+CamionRes.GetIDCamion()+".csv"
	csvFile, er := os.Create(csvPath)
	if er != nil {
		panic(er)
	}
	csvWriter := csv.NewWriter(csvFile)
	err0 := csvWriter.Write([]string{"ID-Paquete", "Tipo de Paquete", "Valor Paquete", "Origen", "Destino", "Numero de Intentos", "Fecha de Entrega"})
	if err0 != nil {
		fmt.Println("ocurrió el medio error al escribir", err0)
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		// an error occurred during the flush
		fmt.Println("Ocurrio un error al hacer el flush")
	}
	//se cierra el archivo ya que despues este debe ser abierto desde funciones
	csvFile.Close()




	var tiempoespera int
	fmt.Println("ingresar tiempo espera...")
	fmt.Scanln(&tiempoespera)

	//Se crea una instancia de camion de camion.Logistica
	//Esta tendrá toda la info del camión de nuestro programa
	MyCamion:=camionLogistica.Camion{
		//Se copian los datos del camion asignado
		IDCamion: CamionRes.GetIDCamion(),
		TipoCamion: CamionRes.GetTipoCamion(),
		TiempoEspera: int32(tiempoespera),
	}
	tiempoEspera := int(MyCamion.TiempoEspera)
	var par *(camionLogistica.ParPaquetes)
	
	par=&(camionLogistica.ParPaquetes{
		Camion: &MyCamion,
	})

	for{
		fmt.Println("esperando paquetes en logistica...")
		par,err=c.AsignarPaquetes(context.Background(),par)
		//debo cerciorarme de cuanto paquetes tengo en el camion
		//cuántos paquetes hay?
		if err!=nil{
			log.Fatalf("error en asignacion de paquetes: %s",err)
		}

		if par.Paquete2.IDPaquete == "" {
			//hay 1 paquete
			cantidad = 1
			par.Paquete1.Estado="En Camino"
			fmt.Println("Recibido paquete con id: ",par.Paquete1.IDPaquete)
		} else {
			//hay 2 paquetes
			cantidad = 2

			if par.Paquete2.ValorProducto> par.Paquete1.ValorProducto{
				aux:=par.Paquete1
				par.Paquete1=par.Paquete2
				par.Paquete2=aux
			}
			par.Paquete1.Estado="En Camino"
			par.Paquete2.Estado="En Camino"
			fmt.Println("Recibido paquete con id: ",par.Paquete1.IDPaquete)
			fmt.Println("Recibido paquete con id: ",par.Paquete2.IDPaquete)
		}

		//actualizamos el registro agregando los nuevos paquetes recibidos
		actualizarRegistro(par, data, cantidad, csvPath)

		//simular repartición:
		if cantidad == 1 {
			for par.Paquete1.Estado == "En Camino" {
				//puede ir con && nose aun
				entregaPaquete(par.Paquete1, csvPath, tiempoEspera)
			}
		} else {
			for par.Paquete1.Estado == "En Camino" || par.Paquete2.Estado == "En Camino" {
				//puede ir con && nose aun
				entregaPaquete(par.Paquete1, csvPath, tiempoEspera)
				c.ReportarIntento(context.Background(), par.Paquete1)
				entregaPaquete(par.Paquete2, csvPath, tiempoEspera)
				c.ReportarIntento(context.Background(), par.Paquete2)
			}

		}
	}
}

