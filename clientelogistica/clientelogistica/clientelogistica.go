package clientelogistica

import(
	context "context"

)

type Cliente_Logistica_Server struct{

}

func(cls *Cliente_Logistica_Server)HacerPedido(ctx context.Context, pedido *Pedido) (*Seguimiento, error){
	return &Seguimiento{CodigoSeguimiento:0}, nil
}

func(cls *Cliente_Logistica_Server)SolicitarEstado(ctx context.Context, seguimiento *Seguimiento) (*Estado, error){
	return &Estado{Estado:"no implementado"}, nil
}