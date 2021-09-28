package main

import "log"

type car struct {
	Direction string
}

func (c *car) Drive() {
	switch c.Direction {
	case "left":
		log.Println("turning left")
	case "right":
		log.Println("turning left")
	case "ahead":
		log.Println("go ahead")
	case "back":
		log.Println("turning back")

	}
}

func main() {
	car := car{}

	car.Direction = "left"
	car.Drive()

	car.Direction = "right"
	car.Drive()

	car.Direction = "ahead"
	car.Drive()

	car.Direction = "back"
	car.Drive()

}
