package main

import (
	"fmt"
	"github.com/coopstools/minibeast/app/beast"
	"github.com/coopstools/minibeast/app/scene/drWillows"
	RatallGrass "github.com/coopstools/minibeast/app/scene/tallGrass"
	"github.com/coopstools/minibeast/app/scene/util"
	"os"
	"time"
)

func main() {
	beastFactory := beast.CreateFactory("beasts.yaml")
	ctx := util.GameCtx{Writer: os.Stdout, Reader: os.Stdin}

	_, _ = fmt.Fprintln(ctx, "~~~WELCOME TO THE WORLD OF MINI BEASTS~~~")

	usersPet := drWillows.AskForPet(ctx, beastFactory)
	time.Sleep(1 * time.Second)
	usersPet = RatallGrass.RandomEncounter(ctx, beastFactory, usersPet)
}
