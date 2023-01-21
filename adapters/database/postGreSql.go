package database

import (
	"P2/RESTAPI/api/request"
	"P2/RESTAPI/core/entities"
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
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
func (s *CollectionImpl)CreateHuman(ctx context.Context,req *entities.Humans)error{
	id := uuid.NewV4()
	human:= &entities.Humans{
		Id: id.String(),
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	}
	if err := s.postGre.Create(human);err.Error != nil{
		return err.Error
	}
	return nil
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














