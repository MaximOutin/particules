package main

import (
	"fmt"
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.

// 0: ConfigSpawnRate, 1: ConfigGravity, 2: ConfigMode
var MenuSelection int = 0
var MenuOpen bool = false

func (g *game) Update() error {
	//menu de selection de la config
	if inpututil.IsKeyJustPressed(ebiten.KeyI) {
		MenuOpen = true
	}
	//navigation
	if MenuOpen {

		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
			MenuSelection++
			if MenuSelection > 4 {
				MenuSelection = 0
			}
			fmt.Println(MenuSelection)
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			MenuSelection--
			if MenuSelection < 0 {
				MenuSelection = 4
			}
			fmt.Println(MenuSelection)
		}

		// 0 : spawnrate
		if MenuSelection == 0 {
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
				if config.General.SpawnRate > 100 {
					config.General.SpawnRate = 100
				} else {
					config.General.SpawnRate += 1
				}
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
				if config.General.SpawnRate < 1 {
					config.General.SpawnRate = 0
				} else {
					config.General.SpawnRate -= 1
				}
			}
		}

		// 1: Gravité
		if MenuSelection == 1 {
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
				config.General.Gravity += 0.5
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
				config.General.Gravity -= 0.5
			}
		}

		// 2 : menuselection
		if MenuSelection == 2 {
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
				if config.General.SpawnMode == "snow" {
					config.General.SpawnMode = "rain"
					fmt.Println("Mode rain active")
				} else if config.General.SpawnMode == "rain" {
					config.General.SpawnMode = "default"
					fmt.Println("Mode default active")
				} else if config.General.SpawnMode == "default" {
					config.General.SpawnMode = "spiral"
					fmt.Println("Mode spiral active")
				} else if config.General.SpawnMode == "spiral" {
					config.General.SpawnMode = "circle"
					fmt.Println("Mode circle active")
				} else if config.General.SpawnMode == "circle" {
					config.General.SpawnMode = "snow"
					fmt.Println("Mode snow active")
				}
				g.system.Content.Init()
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
				if config.General.SpawnMode == "snow" {
					config.General.SpawnMode = "circle"
					fmt.Println("Mode circle active")
				} else if config.General.SpawnMode == "circle" {
					config.General.SpawnMode = "spiral"
					fmt.Println("Mode spiral active")
				} else if config.General.SpawnMode == "spiral" {
					config.General.SpawnMode = "default"
					fmt.Println("Mode default active")
				} else if config.General.SpawnMode == "default" {
					config.General.SpawnMode = "rain"
					fmt.Println("Mode rain active")
					//config.General.ToggleLifetime = false
					//config.General.Toggle_InScreen = true
				} else if config.General.SpawnMode == "rain" {
					config.General.SpawnMode = "snow"
					fmt.Println("Mode snow active")
				}
				g.system.Content.Init()
			}
		}
		// 3: toggle random
		if MenuSelection == 3 {
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
				if config.General.RandomSpawn {
					config.General.RandomSpawn = false
				} else {
					config.General.RandomSpawn = true
				}
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
				if config.General.RandomSpawn {
					config.General.RandomSpawn = false
				} else {
					config.General.RandomSpawn = true
				}
			}
		}
		if MenuSelection == 4 {
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
				if config.General.ToggleDegrade {
					config.General.ToggleDegrade = false
				} else {
					config.General.ToggleDegrade = true
				}
			}
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
				if config.General.ToggleDegrade {
					config.General.ToggleDegrade = false
				} else {
					config.General.ToggleDegrade = true
				}
			}
		}
	}

	g.system.Update()

	return nil
}
