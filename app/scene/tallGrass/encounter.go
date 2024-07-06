package tallGrass

import (
	"fmt"
	"github.com/coopstools/minibeast/app/actions"
	"github.com/coopstools/minibeast/app/beast"
	"github.com/coopstools/minibeast/app/scene/util"
	"strconv"
	"time"
)

const (
	openingDialogue            = "==========Enter Wilderness==========\nYou wonder out of town into the dangers of the tall grass.\n"
	creatureType, creatureName = "Wild Brown Robin", "Brown Robin"
	encounterDialogue          = "You hear the rustling of leaves and spot a %s. %s prepare to battle.\n"
	combatantDisplay           = "\n%s\n%s\n"
	selectActionReq            = "What will %s do: "
)

func RandomEncounter(ctx util.GameCtx, f beast.Factory, pet beast.Pet) beast.Pet {
	ctx.Print(openingDialogue)

	opp := f.BuildPet(creatureType, creatureName)
	ctx.Printf(encounterDialogue, opp.Name, pet.Name)

	act := actions.BuildActionSet()
	var sel, attName string
	var seli int
	for {
		ctx.Printf(combatantDisplay, pet, opp)
		util.ListOptions(ctx, pet.Moves)
		ctx.Printf(selectActionReq, pet.Name)
		_, _ = fmt.Fscanln(ctx, &sel)
		seli, _ = strconv.Atoi(sel)
		attName = pet.Moves[seli]
		pet, opp = act.Op(attName)(ctx, pet, opp)
		if isFightOver(ctx, pet, opp) {
			break
		}

		time.Sleep(1 * time.Second)

		attName = opp.Moves[0]
		opp, pet = act.Op(attName)(ctx, opp, pet)
		if isFightOver(ctx, pet, opp) {
			break
		}
	}
	return pet
}

func isFightOver(ctx util.GameCtx, pet, opp beast.Pet) bool {
	if pet.Hp() <= 0 {
		ctx.Print(pet.Name + " has fainted. You loose.\n")
		return true
	}
	if opp.Hp() <= 0 {
		ctx.Print(opp.Name + " has fainted. You win.\n")
		return true
	}
	return false
}
