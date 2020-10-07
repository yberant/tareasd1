package csvordenes

import(
	"encoding/csv"
	"os"
	"log"
	"fmt"
	"strconv"
)

type Orden struct {
	TimeStamp string
	IDPaquete string
	Tipo string
	Nombre string
	Valor int
	Origen string
	Destino string
	Seguimiento int

}

type CSVOrdenes struct {
	FileName string
}

//retorna las filas
func(csvo *CSVOrdenes) LeerTexto() [][] string {

	path, err := os.Getwd()
	if err != nil {
    	log.Println(err)
	}
	fmt.Println(path) 

	f, err := os.Open(csvo.FileName)
	if err !=nil {
		log.Fatalf("Cannot open '%s': %s\n", csvo.FileName, err.Error())
	}

	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = ';'
	filas, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}
	return filas
}

func(csvo *CSVOrdenes) CrearArchivo(){
	file, err := os.Create(csvo.FileName)
    defer file.Close()
 
    if err != nil {
        os.Exit(1)
    }
    nombresCols := []string{"Timestamp", "IDPaquete", "Tipo", "Nombre", "Valor", "Origen", "Destino","Seguimiento"}
	
	
	fmt.Println("nombres cols:", nombresCols)
	csvWriter := csv.NewWriter(file)
    csvWriter.Write(nombresCols)
    csvWriter.Flush()
}

func(csvo *CSVOrdenes) AñadirOrden(orden Orden) {
	//filas:=csvo.LeerTexto()
	filaOrden:=[]string{orden.TimeStamp,orden.IDPaquete,orden.Tipo,orden.Nombre,strconv.Itoa(orden.Valor),orden.Origen,orden.Destino,strconv.Itoa(orden.Seguimiento)}
	/*filasOrden=append(filasOrden,orden.TimeStamp)
	filasOrden=append(filasOrden,orden.IDPaquete)
	filasOrden=append(filasOrden,orden.Tipo)
	filasOrden=append(filasOrden,orden.Nombre)
	filasOrden=append(filasOrden,strconv.Itoa(orden.Valor))
	filasOrden=append(filasOrden,orden.Origen)
	filasOrden=append(filasOrden,orden.Destino)
	filasOrden=append(filasOrden,strconv.Itoa(orden.Seguimiento))*/

	//filas=append(filas,filaOrden)


	/*file, err := os.Create(csvo.FileName)
    defer file.Close()
 
    */
	file, err:=os.OpenFile(csvo.FileName,os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Println("error: ",err)
        os.Exit(1)
	}
	csvWriter:=csv.NewWriter(file)
	csvWriter.Write(filaOrden)
	
    csvWriter.Flush()
	



}