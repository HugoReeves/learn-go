package main

import "fmt"

//The enemy struct
type enemy struct {
	name      string
	maxhealth int
	health    int
}

//In both of these methods we recieve a pointer to an enemy because we want to be able to update a value instead of returning something
//Revive - Revive is a method that recieves an enemy struct pointer but takes no arguments
func (enemy *enemy) revive() {
	if enemy.health <= 0 {
		//if the enemy has 0 or less health
		enemy.health = enemy.maxhealth
		fmt.Println("You revived", enemy.name, "to max health!")
	} else {
		//If the enemy has more than 0 health, then the enemy is still alive and cannot be revived
		fmt.Println(enemy.name, "could not be revived as they are not dead!")
	}
}

//Attack - Attack is a method that recieves an enemy pointer and takes a damage int argument
func (enemy *enemy) attack(damage int) {
	if enemy.health-damage <= 0 {
		//If you deal enouh damage to kill the enemy
		enemy.health = 0
		fmt.Println("You killed", enemy.name, "by dealing", damage, "to them.")
	} else {
		//If you deal damage but not enough to kill
		enemy.health = enemy.health - damage
		fmt.Println("You delt", damage, "damage to", enemy.name, "now they have", enemy.health, "health remaining.")
	}
}

func main() {
	//defining enemy1 as an enemy struct
	enemy1 := enemy{"Bill", 20, 10}
	fmt.Println(enemy1) //{Bill 20 10}

	//Using the revive method on an enemy who is not dead (Greater than 0 health)
	enemy1.revive() //Bill could not be revived as they are not dead!

	//Attacking the enemy
	enemy1.attack(5) //You delt 5 to Bill now they have 5 remaining.
	enemy1.attack(5) //You killed Bill by dealing 5 to them.

	//reviving a dead enemy
	enemy1.revive() //You revived Bill to max health!
}
