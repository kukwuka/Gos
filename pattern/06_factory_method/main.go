package main

import (
	"errors"
	"fmt"
)

type Car interface {
	GetFuel() int
	AddFuel(int)
	Drive(int) error
}

type BMW struct {
	fuel int
}

type VAZ struct {
	fuel int
}

func (v *VAZ) GetFuel() int {
	return v.fuel
}

func (v *VAZ) Drive(distance int) error {

	if distance*8 > v.fuel {
		return errors.New("not enough Fuel")
	}
	return nil
}
func (v *VAZ) AddFuel(value int) {
	v.fuel += value
}

func (b *BMW) GetFuel() int {
	return b.fuel
}

func (b *BMW) Drive(distance int) error {

	if distance*15 > b.fuel {
		return errors.New("not enough Fuel")
	}
	return nil
}
func (b *BMW) AddFuel(value int) {
	b.fuel += value
}

func BuyCar(money int) (car Car, err error) {

	if 1000 < money && money < 10000 {
		return &VAZ{
			fuel: 0,
		}, nil
	}
	if money > 10000 {
		return &BMW{
			fuel: 0,
		}, nil
	}
	return nil, errors.New("not enough money")
}

func main() {

	bmw, err := BuyCar(100000)
	vaz, err := BuyCar(5000)
	_, err = BuyCar(0)

	fmt.Println(err)
	err = nil

	bmw.AddFuel(40)
	vaz.AddFuel(40)

	err = bmw.Drive(3)
	fmt.Println(err)
	err = nil

	err = vaz.Drive(3)
	fmt.Println(err)

}
