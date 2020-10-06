package camionlogistica

import(
	context "context"
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
	
}

func (cls *Camion_Logistica_Server) ReportarIntento(ctx context.Context,paquete *Paquete) (*Ok, error){
	return &Ok{Ok:0}, nil
}

func (cls *Camion_Logistica_Server) AsignarPaquetes(ctx context.Context,camion *Camion) (*ParPaquetes, error){
	return &ParPaquetes{Paquete1:nil, Paquete2:nil}, nil
}

func (cls *Camion_Logistica_Server) RegistrarCamion(ctx context.Context, ok *Ok)  (*Camion, error){
	return &Camion{IDCamion:"0", TipoCamion: "todavia no se implementa esto xd"}, nil
}