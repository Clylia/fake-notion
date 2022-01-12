package id

// AccountID defines account id object.
type AccountID string

func (a AccountID) String() string {
	return string(a)
}

// PageID defines page id object.
type PageID string

func (p PageID) String() string {
	return string(p)
}
