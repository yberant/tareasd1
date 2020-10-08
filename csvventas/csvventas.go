package csvventas

import(
	"encoding/csv"
	"os"
	"log"
	//"fmt"
	//"time"
	"strconv"
	clientelogistica "../clientelogistica/clientelogistica"
)

type CSVVentas struct {
	NombreArchivo string
	TipoCliente string
}

func (csvv *CSVVentas) Pedido(fila []string) (clientelogistica.Pedido){

	var pedido clientelogistica.Pedido

	if csvv.TipoCliente=="Retail"{
		pedido=clientelogistica.Pedido{
			IDPedido:fila[0],
			NombreProducto:fila[1],
			//ValorProducto:strconv.Atoi(fila[2]),
			//Tipo,

		}
	}

	return pedido
}

func (csvv *CSVVentas) LeerPedidos() [][]string {
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
	
	return filas

	
}

