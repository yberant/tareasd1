package camionlogistica

import(
	context "context"
	colas "../../logistica/colas"
	registroseguimiento "../../logistica/registroseguimiento"
	"fmt"
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
	WaitSeconds int
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

func (cls *Camion_Logistica_Server) AsignarPaquetes(ctx context.Context,camion *Camion) (*ParPaquetes, error){
	return &ParPaquetes{Paquete1:nil, Paquete2:nil}, nil
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