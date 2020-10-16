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
	Valor int32
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
	
	
	//fmt.Println("nombres cols:", nombresCols)
	csvWriter := csv.NewWriter(file)
    csvWriter.Write(nombresCols)
    csvWriter.Flush()
}

func(csvo *CSVOrdenes) AñadirOrden(orden Orden) {
	//filas:=csvo.LeerTexto()
	filaOrden:=[]string{orden.TimeStamp,orden.IDPaquete,orden.Tipo,orden.Nombre,strconv.Itoa(int(orden.Valor)),orden.Origen,orden.Destino,strconv.Itoa(orden.Seguimiento)}
	
	/*dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}*/
	//fmt.Println("escribiendo desde: ",dir)

	//fmt.Println("archivo: ",csvo.FileName)
    
	file, err:=os.OpenFile(csvo.FileName,os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Println("error: ",err)
        os.Exit(1)
	}
	csvWriter:=csv.NewWriter(file)
	csvWriter.Write(filaOrden)
	
	csvWriter.Flush()
	
	//fmt.Println("escrito")
	



}