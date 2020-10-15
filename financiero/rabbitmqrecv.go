package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	//"strconv"
	"bytes"
	"encoding/csv"
	"encoding/json"

	"github.com/streadway/amqp"
)

type Message map[string]interface{}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func deserialize(b []byte) (Message, error) {
	var msg Message
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return msg, err
}

func calcularValor(estadoFinal string, tipoCamion string, tipoPaquete string, valorProducto float64, intentosTotales float64) float64 {
	if estadoFinal == "Recibido" {
		if tipoCamion == "Retail" && tipoPaquete == "Prioritario" {
			//ganancia monto*1,3-((n#intentos-1)*10)
			return valorProducto*1.3 - (intentosTotales-1)*10

		} else {
			return valorProducto - (intentosTotales-1)*10
		}
	} else //no recibido
	{
		if tipoPaquete == "Normal" {
			return (intentosTotales-1) * -10

		} else if tipoPaquete == "Prioritario" {
			return valorProducto*0.3 - (intentosTotales-1)*10
			/*
			if tipoCamion == "Retail" {
				return (valorProducto*1.3)*0.3 - (intentosTotales-1)*10
			} else {
				return valorProducto*0.3 - (intentosTotales-1)*10
			}*/

		} else {//tipoPaquete == Retail
			return valorProducto - (intentosTotales-1)*10
		}
	}
}

func csvAppend(csvPath string, actualizacion []string) {
	file, err := os.OpenFile(csvPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	w.Write(actualizacion)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		// an error occurred during the flush
		fmt.Println("Ocurrio un error al hacer el flush")
	}
	fmt.Println("Se agrego al csv")
	return

}

func main() {
	//Variable que guardara el balance total de PrestigioExpress
	balance := 0.0
	perdidas := 0.0
	ganancias := 0.0
	costos := 0.0

	//crear csv que almacene los campos de los paquetes
	csvPath := "./financiero/finanzas.csv"
	csvFile, er := os.Create(csvPath)
	if er != nil {
		panic(er)
	}
	csvWriter := csv.NewWriter(csvFile)
	err0 := csvWriter.Write([]string{"ID-Paquete", "Intentos Entrega", "Estado", "Ganancia/Perdida"})
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

	fmt.Println("Ingrese direccion IP de logistica")
	var addrS string

	fmt.Scanln(&addrS)
	//conn, err := amqp.Dial("amqp://user:pass@"+addrS+":5672/")

	//se crea conexion rabbitmq localhost
	conn, err := amqp.Dial("amqp://user:pass@"+addrS+":5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//se crea canal rabbitmq
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//se declara una cola en el canal con nombre "hello"
	q, err := ch.QueueDeclare(
		"finanzas", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//msgs es el chan Delivery, sigue retornando mensajes hasta Connection.Close(), Channel.Cancel(), Channel.Close()
	//o cualquier AMQP exception
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for d := range msgs {
			var valor float64
			var msg Message

			msg, err := deserialize(d.Body)
			failOnError(err, "Error en deserializacion")

			
			idPaquete := fmt.Sprintf("%v", msg["ID Paquete"])
			estadoFinal := fmt.Sprintf("%v", msg["Estado Final"])
			intentosTotales := msg["Intentos Totales"].(float64)
			valorProducto := msg["Valor producto"].(float64)
			tipoCamion := fmt.Sprintf("%v", msg["Tipo Camion"])
			tipoPaquete := fmt.Sprintf("%v", msg["Tipo Paquete"])

			fmt.Println("\nha llegado un paquete: "+idPaquete)

			valor = calcularValor(estadoFinal, tipoCamion, tipoPaquete, valorProducto, intentosTotales)

			csvAppend(csvPath, []string{idPaquete, fmt.Sprintf("%v", intentosTotales), estadoFinal, fmt.Sprintf("%v", /*valorProducto-((intentosTotales-1)*10)*/valor)})
			fmt.Println("actualizado registro")
			costos = costos + (intentosTotales-1)*10
			if valor < 0 {
				fmt.Println("perdida asociada: ",valor)
				perdidas = perdidas + valor
			} else {
				fmt.Println("Beneficio asociado: ",valor)
				ganancias = ganancias + valor
			}
		}
	}()

	log.Printf(" [*]Esperando mensajes. Para salir presione CTRL+C")
	// Block until a signal is received.
	<-c
	fmt.Printf("\n Resumen:\n")
	fmt.Printf("Ganancias totales: %#v\n", ganancias)
	fmt.Printf("Perdidas totales: %#v\n", perdidas)
	fmt.Printf("Costos totales: %#v\n", costos)
	balance = ganancias + perdidas
	fmt.Printf("Balance total: %#v\n", balance)
	fmt.Println("Adios")
	os.Exit(3)
}
