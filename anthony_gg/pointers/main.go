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

func takeDamagePointer(player *Player) {
	fmt.Println("Player took some damage.")
	explosionDmg := 10
	player.health -= explosionDmg
}

// Syntactic sugar
// Method not a function
func (player Player) takeDamageReceiver(dmg int) {
	fmt.Println("Player took some damage.")
	player.health -= dmg
}

// Same as `takeDamageReceiver`
func takeDamageNoReceiver(player Player, dmg int) {
	fmt.Println("Player took some damage.")
	player.health -= dmg
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

	player2 := &Player{
		health: 100,
	}

	// pointer approach
	// 8 byte long integer in 64 bits
	fmt.Printf("\nBefore damage %+v\n", player)
	takeDamagePointer(player2)
	fmt.Printf("After damage %+v\n", player)

	// player3 := &Player{
	// 	health: 100,
	// }

	// pointer issue
	// fmt.Printf("\nBefore damage %+v\n", player)
	// somehow deleted player3
	// invalid memory address or nil pointer dereference
	// player3 = nil
	// takeDamagePointer(player3)
	// fmt.Printf("After damage %+v\n", player)

	// receivers
	fmt.Printf("\nBefore damage %+v\n", player)
	player.takeDamageReceiver(20) // Copy
	fmt.Printf("After damage %+v\n", player)

	fmt.Printf("\nBefore damage %+v\n", player)
	takeDamageNoReceiver(player, 30) // Copy
	fmt.Printf("After damage %+v\n", player)
}
