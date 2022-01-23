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

// BlockID defines block id object.
type BlockID string

func (b BlockID) String() string {
	return string(b)
}

// BlobID defines blob id object.
type BlobID string

func (b BlobID) String() string {
	return string(b)
}
