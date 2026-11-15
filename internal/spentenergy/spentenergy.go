package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	cal, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	return cal * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("invalid number of steps: %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("invalid weight: %.2f", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("invalid height: %.2f", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid duration: %v", duration)
	}

	mSpeed := MeanSpeed(steps, height, duration)

	return (weight * mSpeed * duration.Minutes()) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	dist := Distance(steps, height)

	return dist / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 || height <= 0 {
		return 0
	}

	stepLength := height * stepLengthCoefficient

	return stepLength * float64(steps) / mInKm
}
