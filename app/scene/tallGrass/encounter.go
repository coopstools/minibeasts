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
	openingDialogue            = "==========Enter Wilderness==========\nYou wonder out of town into the dangers of the tall grass."
	creatureType, creatureName = "Wild Brown Robin", "Brown Robin"
	encounterDialogue          = "You hear the rustling of leaves and spot a %s. %s prepare to battle.\n"
	combatantDisplay           = "\n%s\n%s\n"
	selectActionReq            = "What will %s do: "
)

func RandomEncounter(ctx util.GameCtx, f beast.Factory, pet beast.Pet) beast.Pet {
	_, _ = fmt.Fprintln(ctx, openingDialogue)

	opp := f.BuildPet(creatureType, creatureName)
	_, _ = fmt.Fprintf(ctx, encounterDialogue, opp.Name, pet.Name)

	act := actions.BuildActionSet()
	var sel string
	for {
		_, _ = fmt.Fprintf(ctx, combatantDisplay, pet, opp)
		util.ListOptions(ctx, act.GetOps())
		_, _ = fmt.Fprintf(ctx, selectActionReq, pet.Name)
		_, _ = fmt.Fscanln(ctx, &sel)
		seli, _ := strconv.Atoi(sel)
		pet, opp = act.Op(seli)(pet, opp)
		if isFightOver(ctx, pet, opp) {
			break
		}

		time.Sleep(1 * time.Second)

		oppAction := 0
		pet, opp = act.Op(oppAction)(opp, pet)
		if isFightOver(ctx, pet, opp) {
			break
		}
	}
	return pet
}

func isFightOver(ctx util.GameCtx, pet, opp beast.Pet) bool {
	if pet.Hp() <= 0 {
		_, _ = fmt.Fprintln(ctx, pet.Name+" has fainted. You loose.")
		return true
	}
	if opp.Hp() <= 0 {
		_, _ = fmt.Fprintln(ctx, opp.Name+" has fainted. You win.")
		return true
	}
	return false
}
