package main

import "log"

type road func()

type car struct {
	road road
}

func (c *car) Drive() {
	c.road()
}

func main() {
	car := car{}
	leftRoad := func() {
		log.Println("turning left")
	}
	rightRoad := func() {
		log.Println("turning right")
	}
	aheadRoad := func() {
		log.Println("go ahead")
	}
	backRoad := func() {
		log.Println("turning back")
	}
	car.road = leftRoad
	car.Drive()
	car.road = rightRoad
	car.Drive()
	car.road = aheadRoad
	car.Drive()
	car.road = backRoad
	car.Drive()

}
