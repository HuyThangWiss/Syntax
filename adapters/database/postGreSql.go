package database

import (
	"P2/RESTAPI/api/middleware"
	"P2/RESTAPI/api/request"
	"P2/RESTAPI/core/entities"
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type CollectionImpl struct {
	postGre *gorm.DB
}

//func NewPostGreSql(impl *gorm.DB) *CollectionImpl {
//	return &CollectionImpl{
//		postGre: impl,
//	}
//}
func NewPostGresql(impl *gorm.DB)*CollectionImpl  {
	return &CollectionImpl{
		postGre: impl,
	}
}
func (s *CollectionImpl)CreateHuman(ctx context.Context,req *entities.Humans)(*entities.Humans,error){
	id := uuid.NewV4()
	human:= &entities.Humans{
		Id: id.String(),
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	}
	if err := s.postGre.Create(human);err.Error != nil{
		return nil,err.Error
	}
	return human,nil
}
func (s *CollectionImpl)FindAll(ctx context.Context,req []entities.Humans)([]entities.Humans,error){
	if result := s.postGre.Find(&req);result.Error != nil{
		return nil,result.Error
	}
	return req,nil
}
func (s *CollectionImpl)FindById(ctx context.Context,req []entities.Humans,Id string)([]entities.Humans,error){
	res := s.postGre.Find(&req,"id = ?",Id)
	if res.RowsAffected == 0 {
		return nil, errors.New("movie not found")
	}
	return req,nil
}
func (s *CollectionImpl)UpdateById(ctx context.Context,req entities.Humans,Id string)error{
	update := request.Humans{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	}
	res := s.postGre.Model(&req).Where("Id = ?",Id).Updates(update)
	if res.RowsAffected == 0 {
		return errors.New("movies not found")
	}
	return nil
}
func (s *CollectionImpl)DeleteById(ctx context.Context,req entities.Humans,Id string)error{
	res := s.postGre.Where("Id = ?",Id).Delete(&req)
	if res.RowsAffected == 0 {
		return errors.New("movie not found")
	}
	return nil
}


func (s *CollectionImpl)FindForm(ctx context.Context,req request.Req)([]entities.Humans,error){
	var arr []entities.Humans
	result:= s.postGre.Where(entities.Humans{
		Id:      req.Id,
		Name:    req.Name,
		Address: req.Address,
	}).Find(&arr)
	if result.Error!= nil {
		fmt.Println(result.Error.Error())
		return nil,errors.New("Err ")
	}
	return arr,nil
}
func (s *CollectionImpl)Login(ctx context.Context,req entities.Humans)(string,error){
	var admin entities.Humans
	err := s.postGre.Where("name = ?",req.Name).Find(&admin)
	if err.Error != nil{
		log.Println("User not found:", err)
		return "",errors.New("err ada 1")
	}
	record := bcrypt.CompareHashAndPassword([]byte(admin.Email),[]byte(req.Email))
	if record != nil{
		return "",errors.New("err ada 2")
		log.Println("Invalid password:", err)
	}
	token,err2 := middleware.GenerateJWT(req.Name,req.Email)
	if err2 != nil{
		return "",errors.New("err ada 3")
	}
	fmt.Println("Login successful!")
	return token,nil
}














