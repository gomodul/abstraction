package abstraction

// PaginationInfo ...
type PaginationInfo struct {
	*Pagination
	More bool `json:"more"`
}
