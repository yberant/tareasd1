package camionlogistica

import(
	context "context"
	colas "../../colas"
	registroseguimiento "../../registroseguimiento"
	"fmt"
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
}

func (cls *Camion_Logistica_Server) ReportarIntento(ctx context.Context,paquete *Paquete) (*Ok, error){

	for i,reg:=range(*(cls.RegistrosSeguimientos)){
		if paquete.GetIDPaquete()==reg.IDPaquete{
			(*cls.RegistrosSeguimientos)[i].Estado=paquete.GetEstado()
			(*cls.RegistrosSeguimientos)[i].CantidadIntentos=int(paquete.GetIntentos())
			//TODO: COMUNICAR A FINANCIERO POR RABBITMQ!!!
			fmt.Println("registros actualizados: ",(*cls.RegistrosSeguimientos))
			break

		}
	}
	

	return &Ok{Ok:0}, nil
}

/*
	Mira, basicamente creé una estructura "colas.Paquete" independiente de la estructura 
	"Paquete" ya existente y definida en este package con ayuda del protoc (Paquete!=package), cuyos atributos son casi los mismos.
	La razón de porqué decidí hacer esto es porque de caso contrario, también hubiese tenido que usar la estructura Paquete en los packages
	clientelogistica y logistica, pero logistica también importa los packages clientelogistica y camionlogistica.
	Básicamente me hubiese dado error por loop de importaciones, por lo que decidí que se trabajara desde una estructura de paquetes
	"independiente", y paf! nació el package colas colas.Paquete 
	
	Hay una razón más que justificada de porqué hice esto, pero no se explicarla bien. Tiene que ver con los loops de importación.

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

func (cls *Camion_Logistica_Server) AsignarPaquetes(ctx context.Context,parpaquetes *ParPaquetes) (*ParPaquetes, error){
	
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

	//actualiza los estados en el registro de memoria a "En Camino"
	cls.ReportarIntento(context.Background(),&PaqRes1)
	cls.ReportarIntento(context.Background(),&PaqRes2)
	return &ParPaquetes{Paquete1:&PaqRes1, Paquete2:&PaqRes2}, nil
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