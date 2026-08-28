package trainings

import (
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("invalid input string: %s", datastring)
	}

	stringSteps, typeStr, durationString := parts[0], parts[1], parts[2]

	parseSteps, err := strconv.Atoi(stringSteps)

	if err != nil || parseSteps <= 0 {
		return fmt.Errorf("invalid steps: %s", stringSteps)
	}

	t.Steps = parseSteps
	t.TrainingType = typeStr

	parseDuration, err := time.ParseDuration(durationString)

	if err != nil || parseDuration <= 0 {
		return fmt.Errorf("invalid duration: %s", durationString)
	}

	t.Duration = parseDuration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64

	switch t.TrainingType {
	case "Бег":
		{
			countCalories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

			if err != nil {
				return "", err
			}

			calories = countCalories
		}

	case "Ходьба":
		{
			countCalories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

			if err != nil {
				return "", err
			}

			calories = countCalories
		}

	default:
		{
			return "", fmt.Errorf("неизвестный тип тренировки")
		}
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
}
