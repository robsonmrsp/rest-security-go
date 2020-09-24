package parser

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Parameters struct {
	m map[string]string
}

func NewParameters() *Parameters {
	m := new(Parameters)
	m.m = make(map[string]string)
	return m
}

func (m *Parameters) Get(key string) (val string) {
	return m.m[key]
}

func (m *Parameters) Exists(key string) bool {
	return m.m[key] != ""
}

func (m *Parameters) Set(key string, val string) {
	m.m[key] = val
}

type PageParameters struct {
	Page       int
	PageSize   int
	TotalPages int
	OrderBy    string
	Order      string

	FilterParameters *Parameters
}

// Create ...
func Create(c *gin.Context) *PageParameters {
	page := new(PageParameters)
	query := c.Request.URL.Query()
	page.Page, _ = strconv.Atoi(query.Get("page"))
	page.PageSize, _ = strconv.Atoi(query.Get("pageSize"))
	page.TotalPages, _ = strconv.Atoi(query.Get("totalPages"))
	page.Order = query.Get("direction")
	page.OrderBy = query.Get("orderBy")
	page.FilterParameters = NewParameters()

	// TODO must be a better way to do this
	for key := range query {
		if key != "page" && key != "pageSize" && key != "totalPages" && key != "direction" && key != "orderBy" {
			page.FilterParameters.Set(key, query.Get(key))
		}
	}
	return page
}
