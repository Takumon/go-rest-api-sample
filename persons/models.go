package persons

import (
	"net/http"

	"./../common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Unit uint8

type Person struct {
	gorm.Model
	CityName  string `form:"city_name" json:"city_name" gorm:"column:city_name;type:varchar(200);"`
	FirstName string `form:"first_name" json:"first_name" binding:"required" gorm:"column:first_name;not null;type:varchar(100);"`
	LastName  string `form:"last_name" json:"last_name" binding:"required" gorm:"column:last_name;not null;type:varchar(100);"`
}

func CreatePersonModel(person Person) (Person, *common.CommonError) {
	tx := common.GetDB().Begin()

	if tx.Error != nil {
		return person, &common.CommonError{
			Error:   tx.Error,
			Message: tx.Error.Error(),
			Code:    http.StatusOK,
		}
	}

	if err := tx.Create(&person).Error; err != nil {
		tx.Rollback()
		return person, &common.CommonError{
			Error:   err,
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return person, &common.CommonError{
			Error:   err,
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	return person, nil
}

func GetPersonModel(id string) (Person, *common.CommonError) {
	var person Person
	if err := common.GetDB().Where("id = ?", id).First(&person).Error; err != nil {
		return person, &common.CommonError{
			Error:   err,
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
	}

	return person, nil
}

func GetPersonsModel() ([]Person, *common.CommonError) {
	var persons []Person
	db := common.GetDB()
	if err := db.Find(&persons).Error; err != nil {
		return persons, &common.CommonError{
			Error:   err,
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	return persons, nil
}

func UpdatePersonModel(id string, person Person) (Person, *common.CommonError) {
	tx := common.GetDB().Begin()

	var currentPerson Person

	if tx.Error != nil {
		return currentPerson, &common.CommonError{
			Error:   tx.Error,
			Message: tx.Error.Error(),
			Code:    http.StatusOK,
		}
	}

	if err := tx.Where("id = ?", id).First(&currentPerson).Error; err != nil {
		return currentPerson, &common.CommonError{
			Error:   err,
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
	}

	if err := tx.Model(&currentPerson).Updates(&person).Error; err != nil {
		tx.Rollback()
		return currentPerson, &common.CommonError{
			Error:   err,
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return currentPerson, &common.CommonError{
			Error:   err,
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	return currentPerson, nil
}

func DeletePersonModel(id string) *common.CommonError {
	tx := common.GetDB().Begin()

	if tx.Error != nil {
		return &common.CommonError{
			Error:   tx.Error,
			Message: tx.Error.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	var person Person
	if err := tx.Where("id = ?", id).Delete(&person).Error; err != nil {
		tx.Rollback()
		return &common.CommonError{
			Error:   err,
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &common.CommonError{
			Error:   err,
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}
