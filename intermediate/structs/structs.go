package main

import "fmt"

//Defining the enemy struct
type enemy struct {
	name   string
	id     int
	health int
	race   string
}

func main() {
	//Shorthand insertion, name, id, health, race
	ralph := enemy{"Ralph", 1, 20, "Goblin"}
	fmt.Println(ralph) //{Ralph 1 20 Goblin}

	//values not defined will be zero valued, in this case id = 0
	james := enemy{name: "James", health: 10, race: "Humanoid"}
	fmt.Println(james) //{James 0 10 Humanoid}

	//The struct is treated as a datatype so it can be used as a type in functions and slices, maps, etc.
	enemies := make([]enemy, 3)
	//Adding the data to the slice.
	//In this case, adding the variables ralph and james works because the datatype of enemy agrees with the slice's datatype of enemy
	enemies[0] = ralph
	enemies[1] = james
	//As expected you can also insert your data immediately.
	enemies[2] = enemy{"Tim", 2, 40, "Alien"}
	fmt.Println(enemies) //[{Ralph 1 20 Goblin} {James 0 10 Humanoid} {Tim 2 40 Alien}]

	//Accessing values in relation to field name
	fmt.Println(ralph.health) //20
}
