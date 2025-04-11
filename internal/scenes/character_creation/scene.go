package character_creation

import (
	"fmt"
	"image/color"
	"time"
	"unicode"

	"github.com/coopstools/minibeast/internal/models"
	"github.com/coopstools/minibeast/internal/scenes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Scene struct {
	gameState    *models.GameState
	sceneManager scenes.SceneManager
	selected     int
	keyDelay     time.Duration
	nameEditing  bool
	nameBuffer   []rune
}

func NewScene(state *models.GameState, manager scenes.SceneManager) *Scene {
	return &Scene{
		gameState:    state,
		sceneManager: manager,
		selected:     0,
		keyDelay:     time.Millisecond * 200,
		nameBuffer:   []rune(state.Character.Name),
	}
}

func (s *Scene) Update() error {
	if s.nameEditing {
		return s.handleNameInput()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		s.selected--
		if s.selected < 0 {
			s.selected = 4
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		s.selected++
		if s.selected > 4 {
			s.selected = 0
		}
	}

	// Handle name field selection
	if s.selected == 4 && inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.nameEditing = true
		return nil
	}

	// Handle stat adjustments
	if s.selected < 4 {
		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			if s.gameState.Character.StatPool > 0 {
				switch s.selected {
				case 0:
					s.gameState.Character.Strength++
					s.gameState.Character.StatPool--
				case 1:
					s.gameState.Character.Dexterity++
					s.gameState.Character.StatPool--
				case 2:
					s.gameState.Character.Vitality++
					s.gameState.Character.StatPool--
				case 3:
					s.gameState.Character.Intelligence++
					s.gameState.Character.StatPool--
				}
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			switch s.selected {
			case 0:
				if s.gameState.Character.Strength > models.MinStatValue {
					s.gameState.Character.Strength--
					s.gameState.Character.StatPool++
				}
			case 1:
				if s.gameState.Character.Dexterity > models.MinStatValue {
					s.gameState.Character.Dexterity--
					s.gameState.Character.StatPool++
				}
			case 2:
				if s.gameState.Character.Vitality > models.MinStatValue {
					s.gameState.Character.Vitality--
					s.gameState.Character.StatPool++
				}
			case 3:
				if s.gameState.Character.Intelligence > models.MinStatValue {
					s.gameState.Character.Intelligence--
					s.gameState.Character.StatPool++
				}
			}
		}
	}

	// Scene transition
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && !s.nameEditing {
		s.sceneManager.SwitchTo("profession")
	}

	return nil
}

func (s *Scene) handleNameInput() error {
	// Handle name editing mode
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.nameEditing = false
		s.gameState.Character.Name = string(s.nameBuffer)
		return nil
	}

	// Handle backspace
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(s.nameBuffer) > 0 {
		s.nameBuffer = s.nameBuffer[:len(s.nameBuffer)-1]
		return nil
	}

	// Handle character input
	for _, char := range ebiten.InputChars() {
		if isValidNameChar(char) && len(s.nameBuffer) < 20 { // 20 char limit
			s.nameBuffer = append(s.nameBuffer, char)
		}
	}

	return nil
}

func isValidNameChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r) || r == ' ' || r == '-' || r == '\''
}

func (s *Scene) Draw(screen *ebiten.Image, assets *scenes.SceneAssets) {
	screen.Fill(color.RGBA{40, 40, 40, 255})
	font := text.NewGoXFace(assets.Font)

	// Draw stats and pool
	stats := []struct {
		n string
		v int
	}{
		{"Strength:", s.gameState.Character.Strength},
		{"Dexterity:", s.gameState.Character.Dexterity},
		{"Vitality:", s.gameState.Character.Vitality},
		{"Intelligence:", s.gameState.Character.Intelligence},
	}

	// Draw stat pool
	poolOpts := &text.DrawOptions{}
	poolOpts.GeoM.Translate(100, 50)
	text.Draw(screen, fmt.Sprintf("Stat Points Available: %d", s.gameState.Character.StatPool), font, poolOpts)

	// Draw stats
	for i, stat := range stats {
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(100, float64(100+i*30))
		prefix := "  "
		if i == s.selected && !s.nameEditing {
			prefix = "> "
		}
		text.Draw(screen, prefix+stat.n, font, opts)
		opts.GeoM.Translate(200, 0)
		text.Draw(screen, fmt.Sprintf("%d", stat.v), font, opts)
	}

	// Draw name field
	nameOpts := &text.DrawOptions{}
	nameOpts.GeoM.Translate(100, float64(100+len(stats)*30))
	namePrefix := "  "
	if s.selected == len(stats) && !s.nameEditing {
		namePrefix = "> "
	}
	if s.nameEditing {
		namePrefix = "* "
	}
	text.Draw(screen, namePrefix+"Name: "+string(s.nameBuffer), font, nameOpts)

	// Draw instructions
	instructOpts := &text.DrawOptions{}
	instructOpts.GeoM.Translate(100, 250)
	if s.nameEditing {
		text.Draw(screen, "Press ENTER to confirm name", font, instructOpts)
	} else {
		text.Draw(screen, "Use arrow keys to navigate and modify stats", font, instructOpts)
		instructOpts.GeoM.Translate(0, 30)
		text.Draw(screen, "Press ENTER to continue", font, instructOpts)
	}
}
