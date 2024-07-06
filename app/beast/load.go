package beast

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Factory struct {
	Beasts   map[string]Beast
	Starters []string
}

func (f Factory) BuildPet(petName, beastName string) Pet {
	b := f.Beasts[beastName]
	p := Pet{Name: petName, base: b.BaseStats, Moves: b.StarterMoves}
	p.current.HP = p.base.HP
	return p
}

func CreateFactory(filePath string) Factory {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("could not read file err   #%v ", err)
	}
	var f Factory
	err = yaml.Unmarshal(yamlFile, &f)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	f.Starters = make([]string, 0, 3)
	for name, b := range f.Beasts {
		if b.IsStarter {
			f.Starters = append(f.Starters, name)
		}
	}

	return f
}
