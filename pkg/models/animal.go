package models

import (
	"time"
)

type Animal struct {
	Id                 int         `json:"id"`
	AnimalTypes        []int       `json:"animalTypes"`
	Weight             float64     `json:"weight"`
	Length             float64     `json:"length"`
	Height             float64     `json:"height"`
	Gender             string      `json:"gender"`
	LifeStatus         string      `json:"lifeStatus"`
	ChippingDateTime   time.Time   `json:"chippingDateTime"`
	ChipperId          int         `json:"chipperId"`
	ChippingLocationId int         `json:"chippingLocationId"`
	VisitedLocations   []int       `json:"visitedLocations"`
	DeathDateTime      interface{} `json:"deathDateTime"`
}

type AnimalSearch struct {
	StartDateTime      string
	EndDateTime        string
	ChipperId          int
	ChippingLocationId int
	LifeStatus         string
	Gender             string
	From               int
	Size               int
}
