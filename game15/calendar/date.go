package calendar

import (
	"fmt"
	"errors"
)

// Date type defined a date in year
type Date struct {
	year int
	month int
	day int
}
// SetYear : set year of date
func (d *Date) SetYear(year int) error {
	if year < 1 { return errors.New("in valid year") }
	d.year = year
	return nil
}
// SetMonth : set month of date
func (d *Date) SetMonth(month int) error{
	if month < 1 || month > 12 {
		return fmt.Errorf("Month must be between %v and %v", 1, 12)
	}
	d.month = month
	return nil
}
// SetDay : set day of date
func (d *Date) SetDay(day int) error{
	daysInMonth := []int{ 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	maxDayInMonth := daysInMonth[d.month - 1]
	if day < 1 || day > maxDayInMonth {
		return fmt.Errorf("Month must be between %v and %v", 1, maxDayInMonth)
	}	
	d.day = day
	return nil
}

// Year : get year of date
func (d *Date) Year() int {
	return d.year
}

// Month : get month of date
func (d *Date) Month() int {
	return d.month
}

// Day : get year of date
func (d *Date) Day() int {
	return d.day
}