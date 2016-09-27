package person

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestPersonRepository struct {
	people map[string]Person
}

func (this TestPersonRepository) GetPerson(id string) (Person, error) {
	val, ok := this.people[id]
	if ok {
		return val, nil
	}
	return val, fmt.Errorf("Person not found")
}

func (this TestPersonRepository) SavePerson(id string, person Person) error {
	this.people[id] = person
	return nil
}

func GetPerson() Person{
	return Person{
		Name: "Roger",
		Age: 40,
		Children: []string{"Alex", "Alva", "Alice", "Theo"},
	}
}

func TestGetPerson(t *testing.T) {
	//arrange
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	inputPerson := GetPerson()
	testPersonDB := TestPersonRepository{people: map[string]Person{"roger": inputPerson}}
	CreatePersonService(router, testPersonDB)

	//act
	req, _ := http.NewRequest("GET", "/api/v1/person/roger", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	var person Person
	json.Unmarshal(resp.Body.Bytes(), &person)

	//assert
	assert.Equal(t, inputPerson.Name, person.Name)
	assert.Equal(t, inputPerson.Age, person.Age)
	assert.Equal(t, inputPerson.Children, person.Children)
}

func TestPutPerson(t *testing.T) {
	//arrange
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	testPersonDB := TestPersonRepository{people: map[string]Person{}}
	CreatePersonService(router, testPersonDB)
	inputPerson := GetPerson()
	b, _ := json.Marshal(inputPerson)

	//act
	req, _ := http.NewRequest("PUT", "/api/v1/person/roger", bytes.NewBuffer(b))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	person, ok := testPersonDB.people["roger"]

	//assert
	assert.True(t, ok, "Person not found")
	assert.Equal(t, inputPerson.Name, person.Name)
	assert.Equal(t, inputPerson.Age, person.Age)
	assert.Equal(t, inputPerson.Children, person.Children)
}
