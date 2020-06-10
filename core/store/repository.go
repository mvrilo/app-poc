package store

import (
	"time"

	"github.com/fatih/structs"
	"github.com/jinzhu/copier"
	"github.com/mvrilo/app-poc/pkg/database"
	"github.com/mvrilo/app-poc/proto/v1"
)

type Repository struct {
	db *database.Database
}

func (r *Repository) Find(params *proto.FindRequest) (*proto.Store, error) {
	var store proto.Store
	err := r.db.Where(structs.Map(params)).Find(&store).Error
	return &store, err
}

func (r *Repository) Create(params *proto.CreateRequest) (*proto.Store, error) {
	var store proto.Store
	if err := copier.Copy(&store, params); err != nil {
		return nil, err
	}
	err := r.db.Create(&store).Error
	return &store, err
}

func (r *Repository) List(params *proto.ListRequest) (*proto.Stores, error) {
	var stores []*proto.Store
	storeParams := structs.Map(params)
	if storeParams["Name"] == "" {
		delete(storeParams, "Name")
	}
	err := r.db.Where(storeParams).Find(&stores).Error
	return &proto.Stores{Stores: stores}, err
}

func (r *Repository) ChangeStatus(params *proto.ChangeStatusRequest) (*proto.Store, error) {
	var store proto.Store
	storeParams := structs.Map(params)
	storeParams["updated_at"] = time.Now()
	err := r.db.Model(&store).Updates(storeParams).Error
	if err != nil {
		return nil, err
	}
	return &store, nil
}
