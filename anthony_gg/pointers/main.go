package main

import "fmt"

type Player struct {
	health int
}

func takeDamage(player Player) {
	fmt.Println("Player took some damage.")
	explosionDmg := 10
	player.health -= explosionDmg
}

func takeDamageFunctional(player Player) Player {
	fmt.Println("Player took some damage.")
	explosionDmg := 10
	player.health -= explosionDmg
	return player
}

func main() {
	player := Player{
		health: 100,
	}

	fmt.Printf("Before damage %+v\n", player)
	takeDamage(player) // Copy
	fmt.Printf("After damage %+v\n", player)

	// functional approach
	fmt.Printf("\nBefore damage %+v\n", player)
	player = takeDamageFunctional(player)
	fmt.Printf("After damage %+v\n", player)
}
