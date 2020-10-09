package entities

import (
	"fmt"
	"strings"
	"time"
)

type CustomTime time.Time
type CustomDate time.Time

const customTimeLayout = "2006-01-02 15:04:05"
const customDateLayout = "2006-01-02"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(customTimeLayout, s)
	*ct = CustomTime(nt)
	return
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

func (ct *CustomTime) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(customTimeLayout))
}

func (ct *CustomDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(customDateLayout, s)
	*ct = CustomDate(nt)
	return
}

func (ct CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

// String returns the time in the custom format
func (ct *CustomDate) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(customDateLayout))
}

// PagerMovie ...
type PagerMovie struct {
	Items        *[]Movie `json:"items"`
	ActualPage   int      `json:"actualPage"`
	TotalRecords int32    `json:"totalRecords"`

	PageSize int    `json:"pageSize"`
	OrderBy  string `json:"orderBy"`
	Order    string `json:"order"`
}

// Movie ...
type Movie struct {
	ID          int        `db:"id" json:"id"`
	Title       string     `db:"name" json:"title"`
	ReleaseDate CustomDate `db:"release_date" json:"releaseDate"`
	Genre       *Genre     `db:"genre_id" json:"genre"`
}
