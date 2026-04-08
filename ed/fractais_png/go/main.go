package main

import (
	"fmt"
	"math/rand"
)

func ri(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

func embua(pen *Pen, dist float64) {
	pen.Walk(dist)
	pen.Right(90)
	pen.Walk(-dist)
	pen.Left(-90)
	pen.Walk(dist)
	pen.Right(90)
	pen.Walk(-dist)
}

func main() {
	pen := NewPen(500, 500)
	pen.SetHeading(90)
	pen.SetPosition(350, 350)
	embua(pen, 250)
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
