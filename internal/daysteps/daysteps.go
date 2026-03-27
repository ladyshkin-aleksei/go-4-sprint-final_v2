package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	slice := strings.Split(data, ",")
	if len(slice) != 2 {
		return 0, 0, fmt.Errorf("Количество переданных параметров о количестве шагов и времени прогулки меньше необходимого\n")
	}
	steps, err := strconv.Atoi(slice[0])
	if err != nil {
		return 0, 0, fmt.Errorf("Некорректный формат количества шагов: %v", err)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("Количество шагов должно быть больше 0\n")
	}
	duration, err := time.ParseDuration(slice[1])
	if err != nil {
		return 0, 0, fmt.Errorf("Некорректный формат длительности: %v", err)
	}

	if duration == 0 {
		return 0, 0, fmt.Errorf("Продолжительность не может быть нулевой\n")
	}

	if duration < 0 {
		return 0, 0, fmt.Errorf("Продолжительность не может быть отрицательной\n")
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Printf("Ошибка парсинга данных: %v", err)
		//fmt.Println("Ошибка при парсинге данных:", err)
		return ""
	}

	if steps <= 0 {
		fmt.Println("Для определения дистанции и количества сожженых квлорий количество шагов должно быть больше 0")
		return ""
	}

	if duration < 0 {
		fmt.Println("Продолжительность не может быть отрицательной")
		return ""
	}

	length := float64(steps) * stepLength
	distance := length / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		log.Printf("Ошибка расчета калорий: %v", err)
		return ""
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)

	return result
}
