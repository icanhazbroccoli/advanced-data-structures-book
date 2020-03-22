package search_tree

type Comparable interface {
	LessThan(Comparable) bool
	EqualsTo(Comparable) bool
}

type SearchKey Comparable

type StoredObject interface{}

type FoundStatus bool
type InsertStatus bool
type DeleteStatus bool

const (
	FoundOk   FoundStatus = true
	FoundNone             = false
)

const (
	InsertOk   InsertStatus = true
	InsertNone              = false
)

const (
	DeleteOk   DeleteStatus = true
	DeleteNone              = false
)

type SearchTree interface {
	RotateLeft()
	RotateRight()

	Find(SearchKey) (StoredObject, FoundStatus)
	Insert(SearchKey, StoredObject) InsertStatus
	Delete(SearchKey) (StoredObject, DeleteStatus)
}
