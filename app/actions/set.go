package actions

import (
	"fmt"
	"github.com/coopstools/minibeast/app/beast"
	"math/rand"
)

type Action func(active, passive beast.Pet) (beast.Pet, beast.Pet)

type Actions struct {
	byName  map[string]Action
	byIndex []string
}

func (a Actions) Op(i int) Action {
	return a.byName[a.byIndex[i]]
}

func (a Actions) GetOps() []string {
	return a.byIndex
}

func BuildActionSet() Actions {
	l := map[string]Action{
		"Attack": Attack,
		"Growl":  Growl,
	}
	il := make([]string, 0, len(l))
	for k, _ := range l {
		il = append(il, k)
	}
	return Actions{byName: l, byIndex: il}
}

func Attack(attacker, defender beast.Pet) (beast.Pet, beast.Pet) {
	fmt.Println(attacker.Name + " attacks!")
	A, D := attacker.Att(), defender.Def()
	if A > 255 || D > 255 {
		A, D = A/4, D/4
	}
	if A == 0 {
		A = 1
	}
	L := 1  //attacker's level
	P := 40 //power of attack; using scratch

	dmg := L * 2 / 5
	dmg += 2
	dmg *= P * A
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

func Growl(attacker, defender beast.Pet) (beast.Pet, beast.Pet) {
	fmt.Println(attacker.Name + " growls!")
	fmt.Println(defender.Name + "'s defense drops.")
	return attacker, defender
}
