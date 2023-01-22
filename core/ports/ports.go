package ports

import (
	"P2/RESTAPI/api/request"
	"P2/RESTAPI/core/entities"
	"context"
)

type PortService interface {
	CreateHuman(context.Context,*entities.Humans)(*entities.Humans,error)
	FindAll(context.Context,[]entities.Humans)([]entities.Humans,error)
	FindById(context.Context,[]entities.Humans,string)([]entities.Humans,error)
	UpdateById(context.Context,entities.Humans,string)error
	DeleteById(context.Context,entities.Humans,string)error

	FindForm(context.Context,request.Req)([]entities.Humans,error)
	Login(context.Context,entities.Humans)(string,error)

}