package camionlogistica

import(
	context "context"
	colas "../../colas"
	registroseguimiento "../../registroseguimiento"
	financieroLogistica "../../financierologistica"
	//"fmt"
	"time"
	//"golang.org/x/net/context"
	//"golang.org/x/net/context"
	//"log"
)

/*
Ojo con camionlogistica.pb.go, ahí se definen:
- structs: "Paquete", "ParPaquetes", "Ok" y "Camion"
- La interfaz CamionLogisticaClient que se debe llamar desde, obviamente, el cliente
- La interfaz CamionLogisticaServer, cuyos métodos se deben definir en este archivo:
*/

type Camion_Logistica_Server struct{
	CamionCount int
	ColasPaquetes *colas.Colas
	RegistrosSeguimientos *[]registroseguimiento.RegistroSeguimiento
	RabbitSender *financieroLogistica.RabbitSender
}

func (cls *Camion_Logistica_Server) ReportarIntento(ctx context.Context,paquete *Paquete) (*Ok, error){

	for i,reg:=range(*(cls.RegistrosSeguimientos)){
		if paquete.GetIDPaquete()==reg.IDPaquete{
			(*cls.RegistrosSeguimientos)[i].Estado=paquete.GetEstado()
			(*cls.RegistrosSeguimientos)[i].CantidadIntentos=int(paquete.GetIntentos())
			//TODO: COMUNICAR A FINANCIERO POR RABBITMQ!!!
			//fmt.Println("registros actualizados: ",(*cls.RegistrosSeguimientos))
			break

		}
	}
	return &Ok{Ok:0}, nil
}

func (cls *Camion_Logistica_Server) RegistrarEntrega(idCamion string, paquete *Paquete){

	if paquete.GetIDPaquete()==""{
		return
	}

	for i,reg:=range(*(cls.RegistrosSeguimientos)){
		if paquete.GetIDPaquete()==reg.IDPaquete{
			(*cls.RegistrosSeguimientos)[i].IDCamion=idCamion
			break

		}
	}
	return
}



/*
	Mira, basicamente creé una estructura "colas.Paquete" independiente de la estructura 
	"Paquete" ya existente y definida en este package con ayuda del protoc (Paquete!=package), cuyos atributos son casi los mismos.
	La razón de porqué decidí hacer esto es porque de caso contrario, también hubiese tenido que usar la estructura Paquete en los packages
	clientelogistica y logistica, pero logistica también importa los packages clientelogistica y camionlogistica.
	Básicamente me hubiese dado error por loop de importaciones, por lo que decidí que se trabajara desde una estructura de paquetes
	"independiente", y paf! nació el package colas colas.Paquete 
	

	De todos modos, esta función es una conversión entre colas.Paquete a (camionlogistica.)Paquete
*/
func ColaspaqToPaq(colasPaq colas.Paquete)(Paquete){
	
	return Paquete{
		IDPaquete:colasPaq.IDPedido,
		CodigoSeguimiento:string(colasPaq.CodigoSeguimiento),
		Tipo:colasPaq.Tipo,
		ValorProducto: colasPaq.ValorProducto,
		Intentos:0,//inicialmente...
		Estado:"En Camino",
		Origen: colasPaq.Origen,
		Destino: colasPaq.Destino,
	}
}

func (cls *Camion_Logistica_Server) EnviarInfoFinanciero(tipoCamion string, paquete *Paquete){
	//fmt.Println("enviando paquete: ",paquete, "a registro")

	msg:=financieroLogistica.Message{
		"Tipo Camion":tipoCamion,
		"ID Paquete":paquete.GetIDPaquete(),
		"Tipo Paquete":paquete.GetTipo(),
		"Estado Final":paquete.GetEstado(),//entregado vs no entregado
		"Intentos Totales":paquete.GetIntentos(),
		"Valor producto":paquete.GetValorProducto(),
	}
	cls.RabbitSender.Publish(msg)

}

func (cls *Camion_Logistica_Server) AsignarPaquetes(ctx context.Context,parpaquetes *ParPaquetes) (*ParPaquetes, error){
	if(parpaquetes.GetPaquete1().GetIDPaquete()!=""){
		cls.EnviarInfoFinanciero(parpaquetes.Camion.GetTipoCamion(),parpaquetes.Paquete1)
	}
	if(parpaquetes.GetPaquete2().GetIDPaquete()!=""){
		cls.EnviarInfoFinanciero(parpaquetes.Camion.GetTipoCamion(),parpaquetes.Paquete2)
	}


	
	var PaqRes1,PaqRes2 Paquete

	switch parpaquetes.Camion.TipoCamion{
	case "Normal":
		//1er paquete
		for{
			//cola prioritaria tiene prioridad
			if len(*((*(cls.ColasPaquetes)).ColaPrioritaria))>=1{

				PaqRes1,*((*(cls.ColasPaquetes)).ColaPrioritaria)=ColaspaqToPaq((*((*(cls.ColasPaquetes)).ColaPrioritaria))[0]), (*((*(cls.ColasPaquetes)).ColaPrioritaria))[1:]
				break
			}
			//cola normal:
			if len(*((*(cls.ColasPaquetes)).ColaNormal))>=1{
				PaqRes1,*((*(cls.ColasPaquetes)).ColaNormal)=ColaspaqToPaq((*((*(cls.ColasPaquetes)).ColaNormal))[0]), (*((*(cls.ColasPaquetes)).ColaNormal))[1:]
				break
			}
		}
		//2do paquete
		for start := time.Now(); time.Since(start) < time.Duration(parpaquetes.Camion.TiempoEspera)*time.Second;{
			//cola prioritaria tiene prioridad
			if len(*((*(cls.ColasPaquetes)).ColaPrioritaria))>=1{
				PaqRes2,*((*(cls.ColasPaquetes)).ColaPrioritaria)=ColaspaqToPaq((*((*(cls.ColasPaquetes)).ColaPrioritaria))[0]), (*((*(cls.ColasPaquetes)).ColaPrioritaria))[1:]
				break
			}
			//cola normal:
			if len(*((*(cls.ColasPaquetes)).ColaNormal))>=1{
				PaqRes2,*((*(cls.ColasPaquetes)).ColaNormal)=ColaspaqToPaq((*((*(cls.ColasPaquetes)).ColaNormal))[0]), (*((*(cls.ColasPaquetes)).ColaNormal))[1:]
				break
			}
		}
	case "Retail":
		if(parpaquetes.GetPaquete1().GetTipo()=="Prioritario"||parpaquetes.GetPaquete2().GetTipo()=="Prioritario"){
			//1er paquete
			for{
				//cola retail tiene prioridad
				if len(*((*(cls.ColasPaquetes)).ColaRetail))>=1{
					PaqRes1,*((*(cls.ColasPaquetes)).ColaRetail)=ColaspaqToPaq((*((*(cls.ColasPaquetes)).ColaRetail))[0]), (*((*(cls.ColasPaquetes)).ColaRetail))[1:]
					break
				}
	
				//cola prioridad
				if len(*((*(cls.ColasPaquetes)).ColaPrioritaria))>=1{
					PaqRes1,*((*(cls.ColasPaquetes)).ColaPrioritaria)=ColaspaqToPaq((*((*(cls.ColasPaquetes)).ColaPrioritaria))[0]), (*((*(cls.ColasPaquetes)).ColaPrioritaria))[1:]
					break
				}
			}
			//2do paquete
			for start := time.Now(); time.Since(start) < time.Duration(parpaquetes.Camion.TiempoEspera)*time.Second;{
				//cola retail tiene prioridad
				if len(*((*(cls.ColasPaquetes)).ColaRetail))>=1{
					PaqRes2,*((*(cls.ColasPaquetes)).ColaRetail)=ColaspaqToPaq((*((*(cls.ColasPaquetes)).ColaRetail))[0]), (*((*(cls.ColasPaquetes)).ColaRetail))[1:]
					break
				}
	
				//cola prioridad
				if len(*((*(cls.ColasPaquetes)).ColaPrioritaria))>=1{
					PaqRes2,*((*(cls.ColasPaquetes)).ColaPrioritaria)=ColaspaqToPaq((*((*(cls.ColasPaquetes)).ColaPrioritaria))[0]), (*((*(cls.ColasPaquetes)).ColaPrioritaria))[1:]
					break
				}
			}	
		} else {
			//paq1. solo retail
			for{
				if len(*((*(cls.ColasPaquetes)).ColaRetail))>=1{
					PaqRes1,*((*(cls.ColasPaquetes)).ColaRetail)=ColaspaqToPaq((*((*(cls.ColasPaquetes)).ColaRetail))[0]), (*((*(cls.ColasPaquetes)).ColaRetail))[1:]
					break
				}
			}
			//paq2. solo retail
			for start := time.Now(); time.Since(start) < time.Duration(parpaquetes.Camion.TiempoEspera)*time.Second;{
				if len(*((*(cls.ColasPaquetes)).ColaRetail))>=1{
					PaqRes2,*((*(cls.ColasPaquetes)).ColaRetail)=ColaspaqToPaq((*((*(cls.ColasPaquetes)).ColaRetail))[0]), (*((*(cls.ColasPaquetes)).ColaRetail))[1:]
					break
				}
			}
		}
	}

	PaqRes1.Estado="En Camino"
	if PaqRes2.IDPaquete!=""{
		PaqRes2.Estado="En Camino"
	}

	//actualiza los estados en el registro de memoria, añadiendo el id del camion al que se le entregó y cambiando el estado a: "En Camino"
	cls.RegistrarEntrega(parpaquetes.Camion.IDCamion,&PaqRes1)
	cls.ReportarIntento(context.Background(),&PaqRes1)

	cls.RegistrarEntrega(parpaquetes.Camion.IDCamion,&PaqRes2)
	cls.ReportarIntento(context.Background(),&PaqRes2)

	//fmt.Println("asignados paquetes a camion: ",parpaquetes.Camion.IDCamion)

	return &ParPaquetes{Paquete1:&PaqRes1, Paquete2:&PaqRes2, Camion:parpaquetes.GetCamion()}, nil
}

func (cls *Camion_Logistica_Server) RegistrarCamion(ctx context.Context, ok *Ok)  (*Camion, error){
	
	var tipoCamion, idCamion string
	switch cls.CamionCount{
	case 0:
		tipoCamion="Normal"
		idCamion="Camion A"
	case 1:
		tipoCamion="Retail"
		idCamion="Camion B"
	case 2:
		tipoCamion="Retail"
		idCamion="Camion C"
	}
	return &Camion{IDCamion:idCamion, TipoCamion: tipoCamion}, nil
}