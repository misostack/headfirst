package main

import (
	"fmt"
)
// MyType is a type of string
type MyType string

// CMTOPIXEL : 1cm = 37.7952755906 px
const CMTOPIXEL = 37.7952755906

// PTTOPIXEL : 1pt = 37.7952755906 px
const PTTOPIXEL = 1.333333

// Centimeter : cm
type Centimeter float64
// Pixel : px
type Pixel float64
// Point : pt
type Point float64

func (p Pixel) String() string {
	return fmt.Sprintf("%0.2f px", p)
}

// ToPixel : convert centimeter to pixel
func (c Centimeter) ToPixel() Pixel {
	return Pixel( c * CMTOPIXEL)
}

func (c Centimeter) String() string {
	return fmt.Sprintf("%0.2f cm", c)
}

// ToPixel : convert PT to Pixel
func (p Point) ToPixel() Pixel {
	return Pixel( p * PTTOPIXEL)
}

func (p Point) String() string {
	return fmt.Sprintf("%0.2f pt", p)
}


func (m MyType) firstLetter() string{
	return string(m[0])
}

func (m *MyType) capitalize() {
	cm := map[rune]rune{}
	uc := 'A'
	for lc := 'a'; lc<= 'z'; lc++{		
		cm[lc] = uc
		uc++
	}

	// uc := map[byte]byte{
	ns := []rune{}
	// }
	for i,v := range *m {
		if i == 0 {
			c, ok := cm[v]
			if ok {
				ns = append(ns, c)
			}		
		} else {
			ns = append(ns, v)
		}
	}
	*m = ""
	for _, v := range ns {
		*m += MyType(v)
	}
}

func main() {
	fmt.Println("[Headfirst Series]: Game 14 , the secret of Defined Types.")
	value := MyType("the first letter should be capitalized")
	fmt.Println(value)
	fmt.Printf("The first letter of %v is '%v'\n", value, value.firstLetter())
	value.capitalize()
	fmt.Println(value)
	cx, cy := Centimeter(2.5), Centimeter(3)
	fmt.Printf("cx=%v and cy=%v ==> x=%v and y=%v\n", cx.String(), cy.String(), cx.ToPixel(), cy.ToPixel())
	px, py := Point(16), Point(20)
	fmt.Printf("px=%v and py=%v ==> x=%v and y=%v\n", px.String(), py.String(), px.ToPixel(), py.ToPixel())
}