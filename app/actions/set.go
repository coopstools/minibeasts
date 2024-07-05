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
	attStat, defStat := attacker.Att(), defender.Def()
	dmg := attStat * attStat / defStat
	fmt.Printf("%s takes %d damage.\n", defender.Name, dmg)
	defender = defender.OffsetHP(-1 * dmg)
	return attacker, defender
}

func Growl(attacker, defender beast.Pet) (beast.Pet, beast.Pet) {
	fmt.Println(attacker.Name + " attacks!")
	fmt.Println(defender.Name + "'s defense drops.")
	return attacker, defender
}
