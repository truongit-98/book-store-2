package models

import "gorm.io/gorm"

func Transaction(fc func(tx *gorm.DB) error) error {
	return db.Transaction(fc)
}
func Commit() error {
	return db.Commit().Error
}

func RollBack() error {
	return db.Rollback().Error
}

func


GetById(sv RepoServiceIf, id uint) (item interface{}, err error) {
	item, err = sv.GetById(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func GetAll(sv RepoServiceIf) (items interface{}, err error) {
	items, err = sv.GetAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func GetPaginate(sv RepoServiceIf, pos, count int32) (items interface{}, totalCount int32, err error) {
	items, totalCount, err = sv.GetPaginate(pos, count)
	if err != nil {
		return nil,0, err
	}
	return items, totalCount,nil
}


func GetBySort(sv RepoServiceIf, sort string) (items interface{}, err error) {
	items, err = sv.GetBySort(sort)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func GetWithCondition(sv RepoServiceIf, condition interface{}, params ...interface{}) (items interface{}, err error) {
	items, err = sv.GetWithConditions(condition, params)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func Exists(sv RepoServiceIf) (exist bool, err error) {
	exist, err = sv.Exists()
	if err != nil {
		return false, err
	}
	return exist, nil
}

func Create(sv RepoServiceIf) (err error) {
	err = sv.Create()
	if err != nil {
		return err
	}
	return nil
}

func Update(sv RepoServiceIf) (err error) {
	err = sv.Update()
	if err != nil {
		return err
	}
	return nil
}

func Remove(sv RepoServiceIf) (err error) {
	err = sv.Remove()
	if err != nil {
		return err
	}
	return nil
}
