package main

import "fmt"

type Player struct {
	Points int
}

func main() {
	
	player1 := Player{
		Points:0,
	}

	//direction player
	fmt.Printf("Player address: %p\n", &player1)

	fmt.Println("Player Points:", player1.Points)

	// first game add 2 points
	player1.AddPoints(2);

	// can try points player 1
	fmt.Println("Player points:", player1.Points)

}

func (r *Player) AddPoints(points int) {
	
	// verify receiver direction
	fmt.Printf("Player address: %p\n", r)

	r.Points += points
} 