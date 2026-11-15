package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	personaldata.Personal

	Steps    int
	Duration time.Duration
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf(`invalid number of values: %d (%q), must be 2 ("steps,duration")`, len(parts), datastring)
	}

	stepsRaw, durationRaw := parts[0], parts[1]

	steps, err := strconv.Atoi(stepsRaw)
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("nonpositive number of steps: %d", steps)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(durationRaw)
	if err != nil {
		return fmt.Errorf("invalid duration: %s", durationRaw)
	}
	if duration <= 0 {
		return fmt.Errorf("nonpositive duration: %v", duration)
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps, ds.Height)
	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	formatStr := `Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`
	return fmt.Sprintf(formatStr, ds.Steps, dist, cal), nil
}
