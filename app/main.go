package main

import (
	"fmt"
	"github.com/coopstools/minibeast/app/actions"
	"github.com/coopstools/minibeast/app/beast"
	"strconv"
	"time"
)

func main() {
	beastFactory := beast.CreateFactory()

	listOptions(beastFactory.Starters)

	usersPet := askForPet(beastFactory)
	wild := beastFactory.BuildPet("Wild Brown Robin", "Brown Robin")
	usersPet, wild = randomEncounter(usersPet, wild)
}

func askForPet(f beast.Factory) beast.Pet {
	var choice string
	for {
		fmt.Print("Choose your beast: ")
		_, _ = fmt.Scanln(&choice)
		if n, err := strconv.Atoi(choice); err == nil && n < len(f.Starters) {
			choice = f.Starters[n]
			break
		}
		fmt.Println("Invalid choice. Please choose a numeric value the from options above.")
	}

	var name string
	for {
		fmt.Printf("Choose a name for your %s: ", choice)
		n, err := fmt.Scanln(&name)
		if n == 0 || err != nil {
			fmt.Println("Sorry, that name doesn't work.")
			continue
		}
		break
	}
	return f.BuildPet(name, choice)
}

func randomEncounter(pet, opp beast.Pet) (beast.Pet, beast.Pet) {
	act := actions.BuildActionSet()
	fmt.Println(pet.Name + " encountered a " + opp.Name)

	var sel string
	for {
		fmt.Printf("\n%s\n%s\n", pet, opp)
		listOptions(act.GetOps())
		fmt.Print("Choose your action: ")
		_, _ = fmt.Scanln(&sel)
		seli, _ := strconv.Atoi(sel)
		pet, opp = act.Op(seli)(pet, opp)

		if pet.Hp() <= 0 {
			fmt.Println(pet.Name + " has fainted. You loose.")
			break
		}
		if opp.Hp() <= 0 {
			fmt.Println(opp.Name + " has fainted. You win.")
			break
		}

		time.Sleep(1 * time.Second)
		fmt.Printf("\n%s\n%s\n", opp, pet)
		fmt.Println("Opponents turn")
		pet, opp = act.Op(0)(opp, pet)

		if pet.Hp() <= 0 {
			fmt.Println(pet.Name + " has fainted. You loose.")
			break
		}
		if opp.Hp() <= 0 {
			fmt.Println(opp.Name + " has fainted. You win.")
			break
		}
	}
	return pet, opp
}

func listOptions(opts []string) {
	for i, name := range opts {
		fmt.Printf("%d: %s\n", i, name)
	}
}
