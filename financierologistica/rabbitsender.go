package financierologistica


import (
  "log"
  //"fmt"
  "net"
  "github.com/streadway/amqp"
  "bytes"
  "encoding/json"
  //"strconv"
)

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



type Message map[string]interface{}

type RabbitSender struct{
	queue *amqp.Queue
	channel *amqp.Channel
}

func (rs *RabbitSender) failOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
}

func(rs *RabbitSender) DeclareChannel(ch *amqp.Channel){
	rs.channel=ch
}

func (rs *RabbitSender) Close(){
	rs.channel.Close()
}

func(rs *RabbitSender) DeclareQueue(){
//declare queue
	q, err := rs.channel.QueueDeclare(
		"finanzas", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	
	rs.failOnError(err, "Failed to declare a queue")
	rs.queue=&q

}

//serializa mensaje como un json
func(rs *RabbitSender) Serialize(msg Message)([]byte, error){
	var b bytes.Buffer
    encoder := json.NewEncoder(&b)
    err := encoder.Encode(msg)
    return b.Bytes(), err
}

func(rs *RabbitSender) Publish(msg Message){

	body,err:=rs.Serialize(msg)

	  
	err = rs.channel.Publish(
		"",     // exchange
		rs.queue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "text/plain",
		  Body:        body,
		})
	rs.failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
	rs.failOnError(err, "Failed to publish a message")
}

/*
func main(){
	IPAddr:=getIPAddr()
	//log.Println(IpAddr)
	fmt.Println("dirección IP de logística: ",IPAddr)

	conn, err := amqp.Dial("amqp://user:pass@"+IPAddr+":5672/")
	if err != nil {
		log.Fatalf("failed to connect to rabbitMq: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel: %s", err)
	}
	defer ch.Close()

	var rabbitS RabbitSender
	rabbitS.DeclareChannel(ch)
	rabbitS.DeclareQueue()


	

	msg:=Message{
		"name":"raul",
		"age":19,
	}

	rabbitS.Publish(msg)
	
}
*/

/*





func main(){

	IPAddr:=getIPAddr()
	//log.Println(IpAddr)
	fmt.Println("dirección IP de logística: ",IPAddr)
	
	/*
	portnum:
		fmt.Println("ingrese puerto para escuchar a a serv. financiero")
		var port string
		fmt.Scanln(&port)
		if p,err:=strconv.Atoi(port);err!=nil{
			goto portnum
		} else {
			_=p
		}
	*/
/*
	sdasdasdasdsad

	
}

*/