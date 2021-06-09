package models

type RepoServiceIf interface {
	GetById(id uint) (item interface{}, err error)
	GetAll() (items interface{}, err error)
	GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error)
	GetBySort(sort string) (items interface{}, err error)
	GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error)
	Exists() (bool, error)
	Create() (err error)
	Update() (err error)
	Remove() (err error)
}

type RepoFactoryIf interface {
	NewFactory(_type string) RepoServiceIf
}

