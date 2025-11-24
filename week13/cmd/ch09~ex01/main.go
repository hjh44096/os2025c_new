package main

import "fmt"

type Kilometers float64
type Meters float64
type Miles float64

//변환 메소드  (Km -> Mile, Miter -> Mile)
func (km Kilometers) ToMiles() Miles {
	return Miles(km / 1.609)
}
func (m Meters) ToMiles() Miles {
	return Miles(m / 1609)
}

func main() {
	kmph := Kilometers(151)
	fmt.Printf("%0.2f kilometer per hour equals %0.2f mile per hour\n", kmph, kmph.ToMiles())
	meter := Meters(160100)
	fmt.Printf("%0.2f meters equals %0.2f mile\n", meter, meter.ToMiles())
}
