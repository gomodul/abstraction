package abstraction

import (
	"fmt"
	"reflect"
)

// Pagination ...
type Pagination struct {
	Page     *int    `json:"page"`
	PageSize *int    `json:"page_size"`
	OrderBy  *string `json:"order_by"`
	Order    *string `json:"order"`
}

// NewPagination ...
func NewPagination() *Pagination {
	page := 1
	pageSize := 10
	order := "desc"
	orderBy := "id"

	return &Pagination{
		Page:     &page,
		PageSize: &pageSize,
		Order:    &order,
		OrderBy:  &orderBy,
	}
}

// Init ...
func (p *Pagination) Init() {
	page := 1
	if p.Page != nil {
		page = *p.Page
	}

	pageSize := 0
	if p.PageSize != nil {
		pageSize = *p.PageSize
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	p.Page = &page
	p.PageSize = &pageSize
}

// GetOffset ...
func (p Pagination) GetOffset() int {
	return (*p.Page - 1) * *p.PageSize
}

// GetLimit ...
func (p Pagination) GetLimit() int {
	return *p.PageSize + 1
}

// GetOrderBy ...
func (p Pagination) GetOrderBy() string {
	orderBy := "id"
	if p.OrderBy != nil {
		orderBy = *p.OrderBy
	}

	order := "desc"
	if p.Order != nil {
		order = *p.Order
	}

	return fmt.Sprintf("%s %s", orderBy, order)
}

// NewPaginationInfo ...
func (p Pagination) NewPaginationInfo(data interface{}) (interface{}, *PaginationInfo) {
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Slice {
		return data, nil
	}

	info := &PaginationInfo{
		Pagination: &p,
		More:       false,
	}

	if p.PageSize != nil && v.Len() > *p.PageSize {
		data = v.Slice(0, *p.PageSize).Interface()
		info.More = true
	}

	return data, info
}
