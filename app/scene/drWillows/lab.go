package drWillows

import (
	"fmt"
	"github.com/coopstools/minibeast/app/beast"
	"github.com/coopstools/minibeast/app/scene/util"
	"strconv"
)

const (
	openingDialogue  = "==========Enter Lab==========\nYou walk into a large lab filled with beakers, bunsen burners, and endless shelves of books. Upon the center bench are three, well behaved mini beasts."
	selectCritterReq = "Professor Willow stands beside the table. \"Well,\" he begins. \"Choose one and start your adventure.\"\nChoose: "
	invalidChoice    = "Invalid choice. Please choose a numeric value the from options above."
	selectNameReq    = "Choose a name for your %s: "
	invalidName      = "Sorry, that name doesn't work."
	departDialogue   = "And so, with %s the %s in tow, you set off into the world to make your mark. An 18 year old (because any younger would be dumb) with their companion, destined for greatness.\n==========Leave Lab==========\n"
)

func AskForPet(ctx util.GameCtx, f beast.Factory) beast.Pet {
	_, _ = fmt.Fprintln(ctx, openingDialogue)
	util.ListOptions(ctx, f.Starters)

	var choice string
	for {
		_, _ = fmt.Fprint(ctx, selectCritterReq)
		_, _ = fmt.Fscanln(ctx, &choice)
		if n, err := strconv.Atoi(choice); err == nil && n < len(f.Starters) {
			choice = f.Starters[n]
			break
		}
		_, _ = fmt.Fprintln(ctx, invalidChoice)
	}

	var name string
	for {
		_, _ = fmt.Fprintf(ctx, selectNameReq, choice)
		n, err := fmt.Fscanln(ctx, &name)
		if n == 0 || err != nil {
			_, _ = fmt.Fprintln(ctx, invalidName)
			continue
		}
		break
	}

	pet := f.BuildPet(name, choice)
	_, _ = fmt.Fprintf(ctx, departDialogue, pet.Name, choice)
	return pet
}
