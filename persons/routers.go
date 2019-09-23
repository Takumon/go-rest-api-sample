package persons

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {
	var person Person

	if err := c.BindJSON(&person); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPerson, dbErr := CreatePersonModel(person)

	if dbErr != nil {
		c.IndentedJSON(dbErr.Code, gin.H{"error": dbErr.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"person": createdPerson})
}

func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	person, dbErr := GetPersonModel(id)

	if dbErr != nil {
		c.IndentedJSON(dbErr.Code, gin.H{"error": dbErr.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"person": person})
}

func GetPersons(c *gin.Context) {
	persons, dbErr := GetPersonsModel()
	if dbErr != nil {
		c.IndentedJSON(dbErr.Code, gin.H{"error": dbErr.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"persons": persons})
}

func UpdatePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person

	if err := c.BindJSON(&person); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPerson, dbErr := UpdatePersonModel(id, person)
	if dbErr != nil {
		c.IndentedJSON(dbErr.Code, gin.H{"error": dbErr.Message})
	}

	c.IndentedJSON(http.StatusOK, updatedPerson)
}

func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	dbErr := DeletePersonModel(id)

	if dbErr != nil {
		c.IndentedJSON(dbErr.Code, gin.H{"error": dbErr.Message})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Delete success"})
}
