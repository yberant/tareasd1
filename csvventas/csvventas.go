package csvventas

import(
	"encoding/csv"
	"os"
	"log"
	//"fmt"
	"strings"
	//"time"
	//"strconv"
	//clientelogistica "../clientelogistica/clientelogistica"
)

type CSVVentas struct {
	NombreArchivo string
}

func (csvv *CSVVentas) LeerPedidos() [][]string {

	//var filas [][]string 

	f, err := os.Open(csvv.NombreArchivo)
	if err !=nil {
		log.Fatalf("Cannot open '%s': %s\n", csvv.NombreArchivo, err.Error())
	}

	defer f.Close()
	r := csv.NewReader(f)

	r.Comma = ';'
	filas, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}
	//fmt.Println("filas:",filas)

	var filasSep [][]string
	
	for _,fila:=range(filas){
		filasSep=append(filasSep,strings.Split(fila[0],","))
	}
	

	return filasSep
}

