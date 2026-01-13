package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
	"project-particles/particle"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	l := list.New()
	switch config.General.SpawnMode {
	case "rain":
		p1 := particle.Raining()
		l.PushFront(&p1)
		return System{Content: l}
	case "snow":
		p1 := particle.Snowing()
		l.PushFront(&p1)
		return System{Content: l}
	case "circle":
		return System{Content: l}
	case "spiral":
		return System{Content: l}
	default:
		if config.General.RandomSpawn {
			for i := 0; i < config.General.InitNumParticles; i++ {
				p1 := particle.NewParticle(rand.Float64()*float64(config.General.WindowSizeX)-1, rand.Float64()*float64(config.General.WindowSizeY)-1, 0, 1, 1, 1, 1, 1, 1, (rand.Float64()*2)-1, (rand.Float64()*2)-1, config.General.Gravity)
				l.PushFront(&p1)
			}
			return System{Content: l}
		} else {
			for i := 0; i < config.General.InitNumParticles; i++ {
				p1 := particle.NewParticle(float64(config.General.WindowSizeX)/2, float64(config.General.WindowSizeY)/2, 0, 1, 1, 1, 1, 1, 1, (rand.Float64()*2)-1, (rand.Float64()*2)-1, config.General.Gravity)
				l.PushFront(&p1)
			}
			return System{Content: l}
		}
	}
}
