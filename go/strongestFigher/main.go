package main

import "fmt"

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func main() {
	fight(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Lew")
}

func (f *Fighter) takeDamage(i int) {
	f.Health = f.Health - i
}

func (f Fighter) checkIfAlive() bool {
	return f.Health > 0
}

func fight(f1 Fighter, f2 Fighter, firstAttacker string) string {
	// Create a conditional to run while loop
	continue_fight := true
	// Define attacker a and defender d to allow subsequent loop to run regardless of order
	var a, d Fighter
	if f1.Name == firstAttacker {
		a, d = f1, f2
	} else {
		a, d = f2, f1
	}
	var winner string
	for continue_fight {
		d.takeDamage(a.DamagePerAttack)
		fmt.Println(d)
		if !d.checkIfAlive() {
			winner = a.Name
			break
		}
		a.takeDamage(d.DamagePerAttack)
		fmt.Println(a)
		if !a.checkIfAlive() {
			winner = d.Name
			break
		}
	}
	fmt.Println("The winner is:", winner)
	return winner
}
