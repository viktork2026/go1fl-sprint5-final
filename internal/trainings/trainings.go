package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	personaldata.Personal

	Steps        int
	TrainingType string
	Duration     time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf(`invalid number of values: %d (%q), must be 3 ("steps,activity,duration")`, len(parts), datastring)
	}

	stepsRaw, trainingType, durationRaw := parts[0], parts[1], parts[2]

	steps, err := strconv.Atoi(stepsRaw)
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("nonpositive number of steps: %d", steps)
	}

	duration, err := time.ParseDuration(durationRaw)
	if err != nil {
		return fmt.Errorf("invalid duration: %s", durationRaw)
	}
	if duration <= 0 {
		return fmt.Errorf("nonpositive duration: %v", duration)
	}

	t.TrainingType = trainingType
	t.Steps = steps
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	dist := spentenergy.Distance(t.Steps, t.Height)
	mSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var cal float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		cal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		cal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		err = fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	formatStr := `Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`
	return fmt.Sprintf(formatStr, t.TrainingType, t.Duration.Hours(), dist, mSpeed, cal), nil
}
