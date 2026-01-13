package particle

import (
	"math"
	"project-particles/config"
	"reflect"
	"testing"
)

func TestNoArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		1,
		config.General.Gravity}
	result := NewParticle()
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestOneArg(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1,
		1,
		1,
		1,
		2,
		1,
		1,
		config.General.Gravity}
	result := NewParticle(2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestTwoArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1,
		1,
		1,
		0.5,
		2,
		1,
		1,
		config.General.Gravity}
	result := NewParticle(0.5, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestThreeArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1,
		1,
		0.4,
		0.5,
		2,
		1,
		1,
		config.General.Gravity}
	result := NewParticle(0.4, 0.5, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestFourArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1,
		0.3,
		0.4,
		0.5,
		2,
		1,
		1,
		config.General.Gravity}
	result := NewParticle(0.3, 0.4, 0.5, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestFiveArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1,
		1.1,
		0.3,
		0.4,
		0.5,
		2,
		1,
		1,
		config.General.Gravity}
	result := NewParticle(1.1, 0.3, 0.4, 0.5, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestSixArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		0,
		1.2,
		1.1,
		0.3,
		0.4,
		0.5,
		2,
		1,
		1,
		config.General.Gravity}
	result := NewParticle(1.2, 1.1, 0.3, 0.4, 0.5, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestSevenArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY),
		math.Pi,
		1.2,
		1.1,
		0.3,
		0.4,
		0.5,
		2,
		1,
		1,
		config.General.Gravity}
	result := NewParticle(math.Pi, 1.2, 1.1, 0.3, 0.4, 0.5, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestEightArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX),
		float64(config.General.WindowSizeY) / 2,
		math.Pi,
		1.2,
		1.1,
		0.3,
		0.4,
		0.5,
		2,
		1,
		1,
		config.General.Gravity}
	result := NewParticle(float64(config.General.WindowSizeY), math.Pi, 1.2, 1.1, 0.3, 0.4, 0.5, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}
func TestTenArgs(t *testing.T) {
	expected := Particle{
		float64(config.General.WindowSizeX) / 3,
		float64(config.General.WindowSizeY) / 2,
		math.Pi,
		1.2,
		1.1,
		0.3,
		0.4,
		0.5,
		2,
		1,
		1,
		config.General.Gravity}
	result := NewParticle(float64(config.General.WindowSizeY)/3, float64(config.General.WindowSizeY), math.Pi, 1.2, 1.1, 0.3, 0.4, 0.5, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}
