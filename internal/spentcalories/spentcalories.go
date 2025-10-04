package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	slice := strings.Split(data, ",")
	if len(slice) != 3 {
		return 0, "", 0, fmt.Errorf("Количество переданных параметров (количество шагов, вид активности и продолжительность активности) меньше необходимого")
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, "", 0, err
	}
	duration, err := time.ParseDuration(slice[2])
	if err != nil {
		return 0, "", 0, err
	}
	activity := slice[1]
	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	distance := ((height * stepLengthCoefficient) * float64(steps)) / mInKm
	return distance
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	distance := distance(steps, height)
	avgspeed := distance / duration.Hours()
	return avgspeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var calories float64
	switch activity {
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("Неизвестный тип тренировки")
	}

	distance := distance(steps, height)
	meanSpeedValue := meanSpeed(steps, height, duration)

	// Формируем строку с информацией о тренировке
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, duration.Hours(), distance, meanSpeedValue, calories)

	return result, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Некорректные входные параметры")
	}
	avgspeed := meanSpeed(steps, height, duration)
	calories := (duration.Minutes() * avgspeed * weight) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Некорректные входные параметры")
	}
	avgspeed := meanSpeed(steps, height, duration)
	calories := (duration.Minutes() * avgspeed * weight) / minInH
	result := walkingCaloriesCoefficient * calories
	return result, nil
}
