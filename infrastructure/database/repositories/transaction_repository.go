package repositories

import (
	"transaction/infrastructure/database/entities"

	"github.com/jinzhu/gorm"
)

type TransactionRepository interface {
	Find(id string) (*entities.TransactionEntity, error)
	Insert(transaction *entities.TransactionEntity) error
	Update(transaction *entities.TransactionEntity) error
	Delete(id string) error
}

type TransactionRepositoryImpl struct {
	Db *gorm.DB
}

func (t *TransactionRepositoryImpl) Find(id string) (*entities.TransactionEntity, error) {
	var transaction entities.TransactionEntity
	err := t.Db.Where("transaction_id = ?", id).Find(&transaction).Error
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (t *TransactionRepositoryImpl) Insert(transaction *entities.TransactionEntity) error {
	err := t.Db.Create(transaction).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryImpl) Update(transaction *entities.TransactionEntity) error {
	err := t.Db.Save(transaction).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepositoryImpl) Delete(id string) error {
	err := t.Db.Where("transaction_id = ?", id).Delete(entities.TransactionEntity{}).Error
	if err != nil {
		return err
	}

	return nil
}