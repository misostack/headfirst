package calendar

import(
	"fmt"
	"unicode/utf8"
)

const titlelength = 30

// Event : defined an event such as birthday, anniversary
type Event struct {
	title string
	Date
}

// SetTitle : Set the event title
func (e *Event) SetTitle(title string) error{
	l := utf8.RuneCountInString(title)
	if l > titlelength { 
		return fmt.Errorf("The title length should be less than %v characters", titlelength)
	}
	e.title = title
	return nil
}

// Title : get the event  title
func (e Event) Title() string {
	return e.title
}

