package person

import (
	"github.com/gin-gonic/gin"
	"github.com/rhinoman/couchdb-go"
	"net/http"
)

type Person struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Children []string `json:"children"`
}

type PersonRepository interface {
	GetPerson(id string) (Person, error)
	SavePerson(id string, person Person) error
}

type CouchDBPersonRepository struct {
	DB *couchdb.Database
}

func (this CouchDBPersonRepository) GetPerson(id string) (Person, error) {
	var doc Person
	if _, err := this.DB.Read(id, &doc, nil); err != nil {
		return doc, err
	}
	return doc, nil
}

func (this CouchDBPersonRepository) SavePerson(id string, person Person) error {
	rev, err := this.DB.Read(id, &person, nil)
	if err != nil {
		return err
	}
	//this is actually racy...
	if _, err := this.DB.Save(person, id, rev); err != nil {
		return err
	}
	return nil
}

func CreatePersonService(router *gin.Engine, db PersonRepository) {
	person := router.Group("api/v1/person")
	{
		person.GET(":id", func(c *gin.Context) {
			id := c.Param("id")
			doc, err := db.GetPerson(id)
			if err != nil {

				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
			c.JSON(http.StatusOK, doc)
		})

		person.PUT(":id", func(c *gin.Context) {
			id := c.Param("id")
			var doc Person
			if err := c.BindJSON(&doc); err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}

			if err := db.SavePerson(id, doc); err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
			c.JSON(http.StatusOK, doc)
		})
	}
}
