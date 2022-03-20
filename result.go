package abstraction

// Result ...
type Result struct {
	Error        error
	LastInsertId int64
	RowsAffected int64
}
