package particles

import (
	"project-particles/config"
	"project-particles/particle"
)

func InScreen(p *particle.Particle) bool {
	maxX := float64(config.General.WindowSizeX + config.General.MarginX)
	maxY := float64(config.General.WindowSizeY + config.General.MarginY)
	minX := -float64(config.General.MarginX)
	minY := -float64(config.General.MarginY)

	return p.PositionX < maxX &&
		p.PositionX > minX &&
		p.PositionY < maxY &&
		p.PositionY > minY
}
