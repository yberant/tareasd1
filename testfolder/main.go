package main

import(
	"fmt"
	//csvo "./csvordenes"
	"time"
)

var(
	secs int
)

type Paq struct{
	Nombre string
	Tipo string
}

type Colas struct{
	ColaN *[]Paq
	ColaR *[]Paq
	ColaP *[]Paq
}

func (cls *Colas) ImprimirColas(){
	fmt.Println("colas:")
	fmt.Println("N: ",*(cls.ColaN))
	fmt.Println("P: ",*(cls.ColaP))
	fmt.Println("R: ",*(cls.ColaR))
	fmt.Println("")
}

func AgregarColas(cls *Colas, paq Paq){
	switch paq.Tipo{
	case "Normal":
		*(cls.ColaN)=append(*(cls.ColaN),paq)
	case "Prioritario":
		*(cls.ColaP)=append(*(cls.ColaP),paq)
	case "Retail":
		*(cls.ColaR)=append(*(cls.ColaR),paq)
	
	}
}

func (cls *Colas) RecibirPaquetes(tipoCamion string,paqEnt1 Paq, paqEnt2 Paq) (Paq, Paq){

	var paqRes1,paqRes2 Paq

	switch tipoCamion{
	case "Normal":
		//primer paquete
		fmt.Println("camion ",tipoCamion," esperando a 1er paquete")
		for{
			//priorizacion de cola prioritaria
			if len(*(cls.ColaP))>=1{
				paqRes1, (*cls.ColaP) = (*cls.ColaP)[0], (*cls.ColaP)[1:]
				break
			}
			//cola normal
			if len(*(cls.ColaN))>=1{
				paqRes1, (*cls.ColaN) = (*cls.ColaN)[0], (*cls.ColaN)[1:]
				break
			}
		}

		fmt.Println("camion ",tipoCamion," esperando a 2do paquete")		
		//segundo paquete (opcional) Espera secs segundos
		for start := time.Now(); time.Since(start) < time.Duration(secs)*time.Second; {
			//priorizacion de cola prioritaria
			if len(*(cls.ColaP))>=1{
				paqRes2, (*cls.ColaP) = (*cls.ColaP)[0], (*cls.ColaP)[1:]
				break
			}
			//cola normal
			if len(*(cls.ColaN))>=1{
				paqRes2, (*cls.ColaN) = (*cls.ColaN)[0], (*cls.ColaN)[1:]
				break
			}
		}

	case "Retail":
		//primer paquete, cito:
		//"Se puede asignar un paquete prioritario a los camiones de retail tras volver de una entrega con paquetes de retail."
		if paqEnt1.Tipo=="Prioritario"||paqEnt2.Tipo=="Prioritario"{
			fmt.Println("camion ",tipoCamion," esperando a 1er paquete")
			for{
				//priorizacion de cola retail
				if len(*(cls.ColaR))>=1{
					paqRes1, (*cls.ColaR) = (*cls.ColaR)[0], (*cls.ColaR)[1:]
					break
				}
					//cola prioritaria
				if len(*(cls.ColaP))>=1{
					paqRes1, (*cls.ColaP) = (*cls.ColaP)[0], (*cls.ColaP)[1:]
					break
				}
			}
			
			fmt.Println("camion ",tipoCamion," esperando a 2do paquete")
			//segundo paquete (opcional) Espera secs segundos
			for start := time.Now(); time.Since(start) < time.Duration(secs)*time.Second; {		
				//priorizacion de cola prioritaria
				if len(*(cls.ColaR))>=1{
					paqRes2, (*cls.ColaR) = (*cls.ColaR)[0], (*cls.ColaR)[1:]
					break
				}
				//cola normal
				if len(*(cls.ColaP))>=1{
					paqRes2, (*cls.ColaP) = (*cls.ColaP)[0], (*cls.ColaP)[1:]
					break
				}
			}

		} else {
			fmt.Println("camion ",tipoCamion," esperando a 1er paquete")
			for{
				//solo cola retail
				if len(*(cls.ColaR))>=1{
					paqRes1, (*cls.ColaR) = (*cls.ColaR)[0], (*cls.ColaR)[1:]
					break
				}
			}

			fmt.Println("camion ",tipoCamion," esperando a 2do paquete")
			//segundo paquete (opcional) Espera secs segundos
			for start := time.Now(); time.Since(start) < time.Duration(secs)*time.Second; {		
				
				//solo retail
				if len(*(cls.ColaR))>=1{
					paqRes2, (*cls.ColaR) = (*cls.ColaR)[0], (*cls.ColaR)[1:]
					break
				}
			}

		}

		//segundo paquete
	}
	fmt.Println("agregados ",paqRes1," y ",paqRes2," a camion ",tipoCamion)
	return paqRes1,paqRes2
}

func Agregador(cls *Colas){
	for{
		fmt.Println("ingrese nombre de paquete a agregar")
		var nombre string
		fmt.Scanln(&nombre)
		var tipo string
		ingtipo:
			fmt.Println("ingrese tipo de paquete a agregar (R,N o P)")
			fmt.Scanln(&tipo)
		if tipo!="R"&&tipo!="N"&&tipo!="P"{
			goto ingtipo
		}
		switch tipo{
		case "R":
			AgregarColas(cls,Paq{nombre,"Retail"})
		case "N":
			AgregarColas(cls,Paq{nombre,"Normal"})
		case "P":
			AgregarColas(cls,Paq{nombre,"Prioritario"})
		}
		fmt.Println("paquete agregado!")
		cls.ImprimirColas()

	}
}

func main() {
	secs=10
	fmt.Println(secs)
	colas:=Colas{
		ColaN:&[]Paq{},
		ColaR:&[]Paq{},
		ColaP:&[]Paq{},
	}
	go Agregador(&colas)

	var pp1,pp2 Paq
	colas.RecibirPaquetes("Normal",pp1,pp2)

	p1,p2:=colas.RecibirPaquetes("Retail",pp1,pp2)

	colas.RecibirPaquetes("Retail",p1,p2)

	colas.RecibirPaquetes("Normal",p1,p2)

	colas.RecibirPaquetes("Normal",pp1,pp2)

	colas.RecibirPaquetes("Retail",pp1,pp2)

	/*
	AgregarColas(&colas,Paq{"A","Prioritario"})
	AgregarColas(&colas,Paq{"B","Normal"})
	AgregarColas(&colas,Paq{"C","Retail"})
	AgregarColas(&colas,Paq{"D","Retail"})
	AgregarColas(&colas,Paq{"E","Retail"})
	AgregarColas(&colas,Paq{"F","Normal"})
	AgregarColas(&colas,Paq{"H","Normal"})
	AgregarColas(&colas,Paq{"I","Prioritario"})
	AgregarColas(&colas,Paq{"J","Normal"})
	AgregarColas(&colas,Paq{"X","Normal"})
	AgregarColas(&colas,Paq{"Y","Prioritario"})
	AgregarColas(&colas,Paq{"Z","Retail"})
	colas.ImprimirColas()
	fmt.Println("\nNormal:")

	var Paq1,Paq2 Paq
	PaqRes1,PaqRes2:=colas.RecibirPaquetes("Normal",Paq1,Paq2)
	
	fmt.Println("paquetes: ",PaqRes1,PaqRes2)
	fmt.Println("colas actuales")
	colas.ImprimirColas()

	fmt.Println("\nRetail:")
	PaqRes1,PaqRes2=colas.RecibirPaquetes("Retail",PaqRes1,PaqRes2)

	fmt.Println("paquetes: ",PaqRes1,PaqRes2)
	fmt.Println("colas actuales")
	colas.ImprimirColas()

	fmt.Println("\nRetail:")
	PaqRes1,PaqRes2=colas.RecibirPaquetes("Retail",PaqRes1,PaqRes2)

	fmt.Println("paquetes: ",PaqRes1,PaqRes2)
	fmt.Println("colas actuales")
	colas.ImprimirColas()

	fmt.Println("\nNormal:")
	PaqRes1,PaqRes2=colas.RecibirPaquetes("Normal",PaqRes1,PaqRes2)

	fmt.Println("paquetes: ",PaqRes1,PaqRes2)
	fmt.Println("colas actuales")
	colas.ImprimirColas()

	fmt.Println("Retail:")
	PaqRes1,PaqRes2=colas.RecibirPaquetes("Retail",PaqRes1,PaqRes2)

	fmt.Println("paquetes: ",PaqRes1,PaqRes2)
	fmt.Println("colas actuales")
	colas.ImprimirColas()*/




}