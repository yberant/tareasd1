package main

import(
	//"fmt"
	csvo "./csvordenes"
	"time"
)

func main() {
	csvOrdenes:=&csvo.CSVOrdenes{FileName: "logistica/logs.csv"}
	csvOrdenes.CrearArchivo()

	var timeString string
	timeString=time.Now().Format("2-Jan-2006 3:04PM")
	
	ordenNueva:=csvo.Orden{
		TimeStamp: timeString,
		IDPaquete: "aaaa",
		Tipo: "retail",
		Nombre: "caja feliz",
		Valor: 30,
		Origen: "casa A",
		Destino: "casa B",
		Seguimiento: 0,
	}

	csvOrdenes.AñadirOrden(ordenNueva)
	csvOrdenes.AñadirOrden(ordenNueva)
	csvOrdenes.AñadirOrden(ordenNueva)
}