package actions

import (
	"bytes"
	"fmt"
	"github.com/coopstools/minibeast/app/beast"
	"github.com/coopstools/minibeast/app/scene/util"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTailWhip(t *testing.T) {
	buffout := bytes.Buffer{}
	ctx := util.GameCtx{Writer: &buffout}

	fmt.Println(os.Getwd())
	bf := beast.CreateFactory("../../beasts.yaml")
	pet := bf.BuildPet("p", "Tiny Leviathan")
	opp := bf.BuildPet("o", "Tiny Leviathan")

	tw := BuildActionSet().Op("Tail Whip")

	pet, opp = tw(ctx, pet, opp)
	pet, opp = tw(ctx, pet, opp)
	pet, opp = tw(ctx, pet, opp)

	assert.Equal(t, 65, pet.Def())
	assert.Equal(t, 26, opp.Def())
}
