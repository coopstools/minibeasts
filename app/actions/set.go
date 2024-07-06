package actions

import (
	"fmt"
	"github.com/coopstools/minibeast/app/beast"
	"github.com/coopstools/minibeast/app/scene/util"
	"math/rand"
)

type Action func(ctx util.GameCtx, active, passive beast.Pet) (beast.Pet, beast.Pet)

type Actions struct {
	byName  map[string]Action
	byIndex []string
}

func (a Actions) Op(oppName string) Action {
	return a.byName[oppName]
}

func (a Actions) GetOps() []string {
	return a.byIndex
}

func BuildActionSet() Actions {
	l := map[string]Action{
		"Gust":      BuildAttack("Gust", 40),
		"Growl":     Growl,
		"Tackle":    BuildAttack("Tackle", 35),
		"Scratch":   BuildAttack("Scratch", 40),
		"Tail Whip": TailWhip,
	}
	il := make([]string, 0, len(l))
	for k, _ := range l {
		il = append(il, k)
	}
	return Actions{byName: l, byIndex: il}
}

func BuildAttack(attName string, powerOfAtt int) Action {
	return func(ctx util.GameCtx, attacker, defender beast.Pet) (beast.Pet, beast.Pet) {
		ctx.Printf("%s uses %s!\n", attacker.Name, attName)
		A, D := attacker.Att(), defender.Def()
		if A > 255 || D > 255 {
			A, D = A/4, D/4
		}
		if A == 0 {
			A = 1
		}
		L := 1 //attacker's level

		dmg := L * 2 / 5
		dmg += 2
		dmg *= powerOfAtt * A
		dmg /= D
		dmg /= 50
		dmg += 2
		if dmg >= 999 {
			dmg = 999
		}

		r := rand.Intn(39) + 236 //should be in the range 217 to 255, but I changed it
		dmg = dmg * r / 255
		fmt.Printf("%s takes %d damage.\n", defender.Name, dmg)
		defender = defender.OffsetHP(-1 * dmg)
		return attacker, defender
	}
}

func TailWhip(ctx util.GameCtx, attacker, defender beast.Pet) (beast.Pet, beast.Pet) {
	ctx.Print(attacker.Name + " uses Tail Whip!\n")
	ctx.Print(defender.Name + "'s defense drops.\n")
	defender = defender.ModDef(-1)
	return attacker, defender
}

func Growl(ctx util.GameCtx, attacker, defender beast.Pet) (beast.Pet, beast.Pet) {
	ctx.Print(attacker.Name + " uses growl!\n")
	ctx.Print(defender.Name + "'s attack drops.\n")
	defender = defender.ModAtt(-1)
	return attacker, defender
}
