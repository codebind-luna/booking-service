package models

import (
	"fmt"

	bookingv1 "github.com/codebind-luna/booking-service/gen/go/booking/v1"
)

var (
	statusMap = map[Status]bookingv1.Status{
		Booked:    bookingv1.Status_STATUS_BOOKED,
		Available: bookingv1.Status_STATUS_AVAILABLE,
	}

	secMap = map[Section]bookingv1.Section{
		A: bookingv1.Section_SECTION_A,
		B: bookingv1.Section_SECTION_B,
	}
)

type Status string

func (s Status) String() string {
	return string(s)
}

func (s Status) Status() bookingv1.Status {
	return statusMap[s]
}

const (
	Available Status = "available"
	Booked    Status = "booked"
)

// Seat entity represents a seat on the map.
type Seat struct {
	r       int
	c       int
	status  Status
	seatNo  int
	user    *User
	section Section
}

func NewSeat(r, c int) *Seat {
	sec := Section(r)
	sNo := c + 1
	return &Seat{r: r, c: c, status: Available, section: sec, seatNo: sNo}
}

func (s *Seat) IsAvailable() bool {
	return s.status == Available
}

func (s *Seat) Status() Status {
	return s.status
}

func (s *Seat) Section() Section {
	return s.section
}

func (s *Seat) SeatNo() int {
	return s.seatNo
}

func (s *Seat) User() *User {
	return s.user
}

func (s *Seat) SetUser(u *User) {
	s.user = u
}

func (s *Seat) SetStatus(st Status) {
	s.status = st
}

type Section int

const (
	A Section = iota
	B
)

func (s Section) String() string {
	return [...]string{"SECTION_A", "SECTION_B"}[s]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (s Section) EnumIndex() int {
	return int(s)
}

func (s Section) Section() bookingv1.Section {
	return secMap[s]
}

var (
	validSections = []Section{
		A,
		B,
	}
	ErrInvalidSection = fmt.Errorf("invalid section")
	sectionMap        = map[string]Section{
		A.String(): A,
		B.String(): B,
	}
)

func isValidSection(section Section) bool {
	for _, valid := range validSections {
		if valid == section {
			return true
		}
	}
	return false
}

func ParseSection(section string) (Section, error) {
	r, ok := sectionMap[section]
	if !ok {
		return 0, ErrInvalidSection
	}
	if !isValidSection(r) {
		return 0, ErrInvalidSection
	}
	return r, nil
}

// SeatMap holds the seat arrangement and provides methods to manipulate the seats.
type SeatMap struct {
	cols  int
	seats [][]*Seat
}

func NewSeatMap(cols int) SeatMap {
	seats := make([][]*Seat, 2)

	for i := 0; i < 2; i++ {
		seats[i] = make([]*Seat, cols)
		for j := 0; j < cols; j++ {
			seats[i][j] = NewSeat(i, j)
		}
	}
	return SeatMap{seats: seats, cols: cols}
}

func (sm *SeatMap) Cols() int {
	return sm.cols
}

func (sm *SeatMap) Seats() [][]*Seat {
	return sm.seats
}
