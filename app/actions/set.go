package actions

import (
	"fmt"
	"github.com/coopstools/minibeast/app/beast"
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

	//baseDmg := (((L * 2) / 5) * P * A) / (50 * D)
	baseDmg := (2 * L * P * A) / (250 * D)
	if baseDmg >= 997 {
		baseDmg = 997
	}
	baseDmg += 2

	//r := rand.Intn(39) + 217
	//dmg := baseDmg * r / 255
	dmg := baseDmg
	fmt.Printf("%s takes %d damage.\n", defender.Name, dmg)
	defender = defender.OffsetHP(-1 * dmg)
	return attacker, defender
}

func Growl(attacker, defender beast.Pet) (beast.Pet, beast.Pet) {
	fmt.Println(attacker.Name + " growls!")
	fmt.Println(defender.Name + "'s defense drops.")
	return attacker, defender
}
