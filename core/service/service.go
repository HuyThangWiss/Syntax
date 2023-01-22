package service

import (
	"P2/RESTAPI/api/middleware/HashPassword"
	"P2/RESTAPI/api/request"
	"P2/RESTAPI/core/entities"
	"P2/RESTAPI/core/ports"
	"context"
	"log"
)

type HumansService struct {
	HumanService ports.PortService
}

func NewHumansService(HumansPort ports.PortService)*HumansService  {
	return &HumansService{
		HumanService: HumansPort,
	}
}

func (u *HumansService)InsertIntoHumans(ctx context.Context,req *request.Humans)(*entities.Humans,error)  {
	post,err := u.HumanService.CreateHuman(ctx,&entities.Humans{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	})
	if err != nil{
		return nil,err
	}
	return post,nil
}

func (u *HumansService)FindAll(ctx context.Context,req []entities.Humans)([]entities.Humans,error)  {
	arr,err := u.HumanService.FindAll(ctx,req)
	if err != nil{
		return nil,err
	}
	return arr,nil
}
func (u *HumansService)FindById(ctx context.Context,req []entities.Humans,Id string)([]entities.Humans,error)  {
	arr,err := u.HumanService.FindById(ctx,req,Id)
	if err != nil{
		return nil,err
	}
	return arr,nil

}
func (u *HumansService)UpdateById(ctx context.Context,req entities.Humans,Id string)  error{
	err := u.HumanService.UpdateById(ctx,req,Id)
	if err != nil{
		return err
	}
	return nil
}

func (u *HumansService)DeleteById(ctx context.Context,req entities.Humans,Id string)error  {
	err:= u.HumanService.DeleteById(ctx,req,Id)
	if err != nil{
		return err
	}
	return nil
}
func (u *HumansService)FindForm(ctx context.Context,req request.Req)([]entities.Humans,error)  {
	post,err := u.HumanService.FindForm(ctx,req)
	if err != nil{
		log.Fatalf("err",err)
		return nil,err
	}
	return post,nil
}
func (u *HumansService)Login(ctx context.Context,req entities.Humans)(string,error) {
	str,err := u.HumanService.Login(ctx,req)
	if err != nil{
		return "", err
	}
	return str,nil
}
func (u *HumansService)RegisterUser(ctx context.Context,req *entities.Humans)(*entities.Humans,error) {
	EmailHash,_ := HashPassword.HashPassword(req.Email)
	req.Email=EmailHash

	post,err := u.HumanService.CreateHuman(ctx,&entities.Humans{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	})
	if err != nil{
		return nil,err
	}
	return post,nil
}



















