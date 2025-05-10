package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	s := strings.Split(datastring, ",")
	if len(s) != 3 {
		return errors.New("error in function training.Parse. len(s) != 3.")
	}
	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	t.Steps = steps
	if t.Steps <= 0 {
		return errors.New("error in function training.Parse. Steps <= 0")
	}
	t.TrainingType = s[1]

	duration, err := time.ParseDuration(s[2])
	if err != nil {
		return err
	}
	t.Duration = duration
	if t.Duration <= 0 {
		return errors.New("error in function daysteps.Parse. Duration <= 0")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	var spentCalories float64
	var err error
	run := "Бег"
	walk := "Ходьба"
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	duration := t.Duration.Hours()
	training := t.TrainingType
	switch training {
	case run:
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	case walk:
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	}
	if training != run && training != walk {
		return "", errors.New("unknown training type")
	}
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", training, duration, distance, meanSpeed, spentCalories)

	return result, nil

}
