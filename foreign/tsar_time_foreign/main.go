package tsar_time_foreign

import (
	. "github.com/adam-mcdaniel/xgopher"
	"time"
)


func Xasm_time_month(m *Machine) {
	m.Push(NewNumber(float64(time.Now().Month())))
}

func Xasm_time_day(m *Machine) {
	m.Push(NewNumber(float64(time.Now().Day())))
}

func Xasm_time_year(m *Machine) {
	m.Push(NewNumber(float64(time.Now().Year())))
}

func Xasm_time_hour(m *Machine) {
	m.Push(NewNumber(float64(time.Now().Hour())))
}

func Xasm_time_minute(m *Machine) {
	m.Push(NewNumber(float64(time.Now().Minute())))
}

func Xasm_time_second(m *Machine) {
	m.Push(NewNumber(float64(time.Now().Second())))
}

func Xasm_time_nanosecond(m *Machine) {
	m.Push(NewNumber(float64(time.Now().Nanosecond())))
}
