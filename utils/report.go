package utils

import (
	"strconv"
	"time"
	"xpress/models"
)

func MonthlyReport(m int) models.Chart {
	now := time.Now()
	if m != 0 {
		now = time.Date(now.Year(), time.Month(m), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())

	}

	first_day := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	last_day := first_day.AddDate(0, 1, -1)
	var labels []string
	var data []int

	total_days := time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
	for i := 0; i < total_days; i++ {
		labels = append(labels, strconv.Itoa(i+1))
		data = append(data, 0)
	}
	var packages []models.Package
	if DB.Select("created_at").Find(&packages).Error != nil {
		panic("error")
	}
	for _, p := range packages {
		if p.Created_at > first_day.Unix() && p.Created_at < last_day.Unix() {
			t := time.Unix(p.Created_at, 0).Day() - 1
			data[t]++
		}
	}
	return models.Chart{
		Labels: labels,
		Data:   data,
	}
}

func YearlyReport(y int) models.Chart {
	now := time.Now()
	if y != 0 {
		now = time.Date(y, now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), now.Location())
	}
	var labels []string
	var data []int

	first_day := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	last_day := first_day.AddDate(1, 0, -1)

	for i := 0; i < 12; i++ {
		labels = append(labels, strconv.Itoa(i+1))
		data = append(data, 0)
	}
	var packages []models.Package

	if DB.Select("created_at").Find(&packages).Error != nil {
		panic("error")
	}
	for _, p := range packages {
		if p.Created_at > first_day.Unix() && p.Created_at < last_day.Unix() {
			t := time.Unix(p.Created_at, 0).Month() - 1
			data[t]++
		}
	}
	return models.Chart{
		Labels: labels,
		Data:   data,
	}
}
