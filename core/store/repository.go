package store

import (
	"time"

	"github.com/fatih/structs"
	"github.com/jinzhu/copier"
	"github.com/mvrilo/storepoc/pkg/database"
	"github.com/mvrilo/storepoc/proto"
)

type Repository struct {
	db *database.Database
}

func (r *Repository) Find(params *proto.FindRequest) (store *proto.Store, err error) {
	err = r.db.Find(structs.Map(params), &store).Error
	return
}

func (r *Repository) Create(params *proto.CreateRequest) (*proto.Store, error) {
	var store proto.Store
	err := copier.Copy(&store, params)
	if err != nil {
		return nil, err
	}

	err = r.db.Create(&store).Error
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

func (r *Repository) ChangeStatus(params *proto.ChangeStatusRequest) (store *proto.Store, err error) {
	storeParams := structs.Map(params)
	storeParams["updated_at"] = time.Now()
	err = r.db.Model(&store).Updates(storeParams).Error
	return
}
