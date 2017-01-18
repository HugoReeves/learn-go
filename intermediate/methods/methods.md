# Methods

[Code](methods.go)

Methods are simmilar to functions and are written in a simmilar way, they can take arguments but they must always have a reciever. For the following explanations we will be using the following struct,
```go
type enemy struct {
	name      string
	maxhealth int
	health    int
}
```
And we will create one variable called enemy1 with the enemy struct, `enemy1 := enemy{"Bill", 20, 10}`.

### Defining Methods

Defining Methods is very simmilar to defining a function, we just have to include a statement that declares the reciever. We define a function with the following syntax `func (recievername type) methodname(args) optionalreturntype { CODE }`, where recievername is treated just like an argname and the type can be a standard type or a struct.
In our code, we defined two methods, one method, revive, takes a pointer to an enemy type as the reciever and no args, the other method, attack, takes a pointer to an enemy type as the reciever and a single int damage arg.
The code for the revive method:
```go
func (enemy *enemy) revive() {
	if enemy.health <= 0 {
		enemy.health = enemy.maxhealth
		fmt.Println("You revived", enemy.name, "to max health!")
	} else {
		fmt.Println(enemy.name, "could not be revived as they are not dead!")
	}
}
```
This code, checks if the health of the enemy is less than or equal to 0, if so, it sets the enemies health to max and tells the user who they have revived. Otherwise, it tells the user that the enemy coul not be revived as they are not dead.

The code for the attack method:
```go
func (enemy *enemy) attack(damage int) {
	if enemy.health-damage <= 0 {
		enemy.health = 0
		fmt.Println("You killed", enemy.name, "by dealing", damage, "to them.")
	} else {
		enemy.health = enemy.health - damage
		fmt.Println("You delt", damage, "damage to", enemy.name, "now they have", enemy.health, "health remaining.")
	}
}
```
This code, takes an enemy reciever and an arg specifing the ammount of damage to apply. If the damage would kill the enemy, it sets the enemies health to 0 and tells the user that the enemy was killed. If the damage would not kill the enemy, the health of the enemy is reduced accordingly and the user is told how much damage was delt.

### Using Methods

You can apply methods to objects using the following syntax, `object.METHOD(arguments if needed)`. Using our previously defined methods, we can apply them in the following ways to the enemy1 enemy object, `enemy1.revive()` to revive the enemy or `enemy1.attack(5)` to deal 5 damage to the enemy.

### Next

* [Home](../../README.md)
* [Intermediate Overview](../intermediate.md)
* [Next Lesson: Methods](../methods/methods.md)