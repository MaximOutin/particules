package main

import (
	"fmt"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particle"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.
func (g *game) Draw(screen *ebiten.Image) {

	for e := g.system.Content.Front(); e != nil; e = e.Next() {

		p, ok := e.Value.(*particle.Particle)
		if ok {
			if config.General.ToggleLifetime && config.General.Toggle_InScreen {
				if particles.Lifetime(p) && particles.InScreen(p) {
					options := ebiten.DrawImageOptions{}
					options.GeoM.Rotate(p.Rotation)
					options.GeoM.Scale(p.ScaleX, p.ScaleY)
					options.GeoM.Translate(p.PositionX, p.PositionY)
					options.ColorScale.Scale(float32(p.ColorRed), float32(p.ColorGreen), float32(p.ColorBlue), float32(p.Opacity))
					screen.DrawImage(assets.ParticleImage, &options)
				}
			} else if config.General.ToggleLifetime {
				if particles.Lifetime(p) {
					options := ebiten.DrawImageOptions{}
					options.GeoM.Rotate(p.Rotation)
					options.GeoM.Scale(p.ScaleX, p.ScaleY)
					options.GeoM.Translate(p.PositionX, p.PositionY)
					options.ColorScale.Scale(float32(p.ColorRed), float32(p.ColorGreen), float32(p.ColorBlue), float32(p.Opacity))
					screen.DrawImage(assets.ParticleImage, &options)
				}
			} else if config.General.Toggle_InScreen {
				if particles.InScreen(p) {
					options := ebiten.DrawImageOptions{}
					options.GeoM.Rotate(p.Rotation)
					options.GeoM.Scale(p.ScaleX, p.ScaleY)
					options.GeoM.Translate(p.PositionX, p.PositionY)
					options.ColorScale.Scale(float32(p.ColorRed), float32(p.ColorGreen), float32(p.ColorBlue), float32(p.Opacity))
					screen.DrawImage(assets.ParticleImage, &options)
				}
			} else {
				options := ebiten.DrawImageOptions{}
				options.GeoM.Rotate(p.Rotation)
				options.GeoM.Scale(p.ScaleX, p.ScaleY)
				options.GeoM.Translate(p.PositionX, p.PositionY)
				options.ColorScale.Scale(float32(p.ColorRed), float32(p.ColorGreen), float32(p.ColorBlue), float32(p.Opacity))
				screen.DrawImage(assets.ParticleImage, &options)

			}
		}

		if config.General.Debug {
			ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS()))
		}

		if MenuOpen {
			msg := "\nMenu de sélection"
			msg += "\nUtilisez les fleches haut/bas pour naviguer et gauche/droite pour modifier.\n"

			if MenuSelection == 0 {
				msg += "> "
			} else {
				msg += "  "
			}
			msg += fmt.Sprintf("Spawn Rate : %.2f\n", config.General.SpawnRate)

			if MenuSelection == 1 {
				msg += "> "
			} else {
				msg += "  "
			}
			msg += fmt.Sprintf("Gravity : %.2f\n", config.General.Gravity)

			if MenuSelection == 2 {
				msg += "> "
			} else {
				msg += "  "
			}

			msg += fmt.Sprintf("Mode : %s\n", config.General.SpawnMode)

			if MenuSelection == 3 {
				msg += "> "
			} else {
				msg += "  "
			}
			msg += fmt.Sprintf("Random : %t\n", config.General.RandomSpawn)

			if MenuSelection == 4 {
				msg += "> "
			} else {
				msg += "  "
			}
			msg += fmt.Sprintf("Degrade : %t\n", config.General.ToggleDegrade)

			ebitenutil.DebugPrint(screen, msg)
		} else {
			ebitenutil.DebugPrintAt(screen, "I: Menu", 0, 20)
		}

	}
}
