package particles

import (
	"math/rand"
	"project-particles/config"
	"project-particles/particle"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
var compte float64 = 0

func (s *System) Update() {
	for e := s.Content.Front(); e != nil; {
		next := e.Next()

		p := e.Value.(*particle.Particle)

		p.PositionX += p.SpeedX
		if config.General.Toggle_gravity {
			p.PositionY += p.SpeedY + p.Gravity
		} else {
			p.PositionY += p.SpeedY
		}

		if config.General.ToggleLifetime {
			if !Lifetime(p) {
				s.Content.Remove(e)
				e = next
				continue
			}
		}
		e = next
	}

	compte += config.General.SpawnRate

	for compte >= 1 {
		switch config.General.SpawnMode {
		case "rain":
			new_p := particle.Raining()
			s.Content.PushFront(&new_p)
		case "snow":
			new_p := particle.Snowing()
			s.Content.PushFront(&new_p)
		case "circle":
			new_p := particle.Circle()
			s.Content.PushFront(&new_p)
		case "spiral":
			new_p := particle.Spiral()
			s.Content.PushFront(&new_p)
		default:
			if config.General.RandomSpawn {
				new_p := particle.NewParticle(rand.Float64()*float64(config.General.WindowSizeX)-1, rand.Float64()*float64(config.General.WindowSizeY)-1, 0, 1, 1, 1, 1, 1, 1, (rand.Float64()*2)-1, (rand.Float64()*2)-1, config.General.Gravity)
				s.Content.PushFront(&new_p)
			} else {
				new_p := particle.NewParticle(float64(config.General.WindowSizeX)/2, float64(config.General.WindowSizeY)/2, 0, 1, 1, 1, 1, 1, 1, (rand.Float64()*2)-1, (rand.Float64()*2)-1, config.General.Gravity)
				s.Content.PushFront(&new_p)
			}
		}
		compte -= 1
	}
}
