package particle

import (
	"math"
	"math/rand"
	"project-particles/config"
)

var angle float64 = 0

// Fonction de génération pour le mode : 'snow'
func Snowing() Particle {
	//définie les valeurs par défaut des paramètres
	params := []float64{
		rand.Float64()*float64(config.General.WindowSizeX) - 1, 0, // PositionX, PositionY / Génération des particules en haut de l'écran
		0,          // Rotation / Pas de rotation pour la pluie
		0.50, 0.50, // ScaleX, ScaleY / Taille réduite pour faire les flocons
		1, 1, 1, // ColorRed, ColorGreen, ColorBlue / Couleur blanche pour le mode "snow"
		1,                                        // Opacity
		(rand.Float64() * 2) - 1, rand.Float64(), // SpeedX, SpeedY / On veux que la neige tombe
		config.General.Gravity, //Gravity
	}
	return Particle{
		PositionX:  params[0],
		PositionY:  params[1],
		Rotation:   params[2],
		ScaleX:     params[3],
		ScaleY:     params[4],
		ColorRed:   params[5],
		ColorGreen: params[6],
		ColorBlue:  params[7],
		Opacity:    params[8],
		SpeedX:     params[9],
		SpeedY:     params[10],
		Gravity:    params[11],
	}
}

// Fonction de génération pour le mode : 'rain'
func Raining() Particle {
	//définie les valeurs par défaut des paramètres
	params := []float64{
		rand.Float64()*float64(config.General.WindowSizeX) - 1, 0, // PositionX, PositionY / Génération des particules en haut de l'écran
		0,       // Rotation / Pas de rotation pour la pluie
		0.10, 1, // ScaleX, ScaleY / Taille réduite sur la largeur pour faire la pluie
		0, 0, 1, // ColorRed, ColorGreen, ColorBlue / Couleur bleu pour le mode "rain"
		1,                       // Opacity
		0, rand.Float64() + 1*5, // SpeedX, SpeedY / On veux que la pluie tombe
		config.General.Gravity} // Gravity
	return Particle{
		PositionX:  params[0],
		PositionY:  params[1],
		Rotation:   params[2],
		ScaleX:     params[3],
		ScaleY:     params[4],
		ColorRed:   params[5],
		ColorGreen: params[6],
		ColorBlue:  params[7],
		Opacity:    params[8],
		SpeedX:     params[9],
		SpeedY:     params[10],
		Gravity:    params[11],
	}
}

// Fonction de génération pour le mode : 'circle'
func Circle() Particle {

	// Détermine la position du centre et le rayon du cercle
	centerX := float64(config.General.WindowSizeX) / 2
	centerY := float64(config.General.WindowSizeY) / 2
	radius := 250.0

	//Détermine les angles sur le cercle selon le nombre de particule à générer
	angle += (2 * math.Pi) / (float64(config.General.SpawnRate) * 25)

	//Redémarage de l'angle à 0
	if angle >= 2*math.Pi {
		angle = 0
	}

	//Détermine la position d'apparition de la particule sur le cercle
	posX := centerX + radius*math.Cos(angle)
	posY := centerY + radius*math.Sin(angle)

	//La particule ne peut pas entrer dans le rayon du cercle
	speedX := rand.Float64() * math.Cos(angle)
	speedY := rand.Float64() * math.Sin(angle)

	params := []float64{
		posX, posY, // PositionX, PositionY
		0,        // Rotation
		0.5, 0.5, // ScaleX, ScaleY
		rand.Float64(), rand.Float64(), rand.Float64(), // ColorRed, ColorGreen, ColorBlue / Multicolor
		1,              // Opacity
		speedX, speedY, // SpeedX, SpeedY
		0, // Gravity
	}

	return Particle{
		PositionX:  params[0],
		PositionY:  params[1],
		Rotation:   params[2],
		ScaleX:     params[3],
		ScaleY:     params[4],
		ColorRed:   params[5],
		ColorGreen: params[6],
		ColorBlue:  params[7],
		Opacity:    params[8],
		SpeedX:     params[9],
		SpeedY:     params[10],
		Gravity:    params[11],
	}
}

func Spiral() Particle {

	spawnX := float64(config.General.WindowSizeX) / 2
	spawnY := float64(config.General.WindowSizeY) / 2
	radius := angle * 10
	angleIncrement := (2 * math.Pi) / (float64(config.General.SpawnRate) * 100)
	angle += angleIncrement

	if angle >= 10*math.Pi {
		angle = 0
	}
	posX := spawnX + radius*math.Cos(angle)
	posY := spawnY + radius*math.Sin(angle)

	speedX := 0.2 * math.Cos(angle)
	speedY := 0.2 * math.Sin(angle)

	params := []float64{
		posX, posY, // Position
		0,        // Rotation
		0.1, 0.1, // ScaleX, ScaleY (petites particules)
		1, 1, rand.Float64(), // RGB
		1,              // Opacity
		speedX, speedY, // SpeedX, SpeedY
		0, // Gravity (souvent désactivée pour une spirale centrée)
	}

	return Particle{
		PositionX:  params[0],
		PositionY:  params[1],
		Rotation:   params[2],
		ScaleX:     params[3],
		ScaleY:     params[4],
		ColorRed:   params[5],
		ColorGreen: params[6],
		ColorBlue:  params[7],
		Opacity:    params[8],
		SpeedX:     params[9],
		SpeedY:     params[10],
		Gravity:    params[11],
	}
}
