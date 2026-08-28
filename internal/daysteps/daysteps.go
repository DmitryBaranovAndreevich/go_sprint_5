package daysteps

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 2 {
		return fmt.Errorf("invalid string: %s", datastring)
	}

	stringSteps, durationString := parts[0], parts[1]

	parseSteps, err := strconv.Atoi(stringSteps)

	if err != nil || parseSteps <= 0 {
		return fmt.Errorf("invalid steps: %s", stringSteps)
	}

	ds.Steps = parseSteps

	parseDuration, err := time.ParseDuration(durationString)

	if err != nil || parseDuration <= 0 {
		return fmt.Errorf("invalid duration: %s", durationString)
	}

	ds.Duration = parseDuration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	countCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, countCalories), nil

}
