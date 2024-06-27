package main

import "fmt"

type Player struct {
	Points int
}

func main() {
	
	player1 := Player{
		Points:0,
	}

	fmt.Println("Player Points:", player1.Points)

	// first game add 2 points

}

func (r Player) AddPoints(points int) {
	
	r.Points += points
} 