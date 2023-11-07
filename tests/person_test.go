package tests

import (
	"testing"

	"github.com/Kiril-Poposki1998/GOfiberMiniAPI/model"
	"github.com/magiconair/properties/assert"
)

func TestPersonCreate(t *testing.T) {
	p := model.Person{
		Id:      1,
		Name:    "Kiril",
		Surname: "Poposki",
	}
	assert.Equal(t, p.Name, "Kiril")
	assert.Equal(t, p.Surname, "Poposki")
}
