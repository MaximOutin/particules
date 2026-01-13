package particles

import (
	"project-particles/config"
	"project-particles/particle"
)

func Lifetime(p *particle.Particle) bool {
	if p.Opacity <= 0 {
		return false
	}
	p.Opacity = p.Opacity - 0.008*(1/config.General.DurationTime)
	return true
}
