package utils

import (
	"strconv"
	"time"
)

type TimeUtils struct {
	TimeStr 		string
	FormatLayout	string
	Result 			string
	Time 			time.Time
	// Parts of Time
	Year			string
	Month			string
	Day				string
	Hour			string
	Minute			string
	Second			string
}

var instance *TimeUtils

func GetInstance() *TimeUtils  {
	if instance == nil {
		instance = &TimeUtils{}
	}

	return instance
}

func (timeUtils *TimeUtils) Parse()  {
	_time := timeUtils.Time
	timeUtils.Year 		= strconv.Itoa(_time.Year())
	timeUtils.Month 	= strconv.Itoa(int(_time.Month()))
	timeUtils.Day		= strconv.Itoa(_time.Day())
	timeUtils.Hour		= strconv.Itoa(_time.Hour())
	timeUtils.Minute	= strconv.Itoa(_time.Minute())
	timeUtils.Second	= strconv.Itoa(_time.Second())
}

func (timeUtils *TimeUtils) GetNowWithFormatOfAutoticket() string  {
	timeUtils.Time = time.Now()
	timeUtils.Parse()
	str := timeUtils.Year +
		"-" + timeUtils.Month +
		"-" + timeUtils.Day +
		"T" + timeUtils.Hour +
		":" + timeUtils.Minute +
		":" + timeUtils.Second
	return str
}