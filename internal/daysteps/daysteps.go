package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	s := strings.Split(datastring, ",")
	if len(s) != 2 {
		return errors.New("error in function daysteps.Parse. len(s) != 2")
	}
	steps, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	ds.Steps = steps
	if ds.Steps <= 0 {
		return errors.New("error in function daysteps.Parse. Steps <= 0")
	}
	duration, err := time.ParseDuration(s[1])
	if err != nil {
		return err
	}
	ds.Duration = duration
	if ds.Duration <= 0 {
		return errors.New("error in function daysteps.Parse. Duration <= 0")
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	distance := spentenergy.Distance(ds.Steps, ds.Height)
	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return " ", err
	}
	result := fmt.Sprintf("Количество шагов: %d \nДистанция составила %.2f км.\nВы сожгли %.2f \n", ds.Steps, distance, spentCalories)
	return result, nil
}
