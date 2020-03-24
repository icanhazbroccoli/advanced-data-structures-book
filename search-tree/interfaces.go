package search_tree

type Comparable interface {
	LessThan(Comparable) bool
	EqualsTo(Comparable) bool
	LessThanOrEqualsTo(Comparable) bool
}

type SearchKey Comparable

type StoredObject interface{}

type FindStatus bool
type InsertStatus bool
type DeleteStatus bool

const (
	FindOk   FindStatus = true
	FindNone            = false
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
	Find(SearchKey) (StoredObject, FindStatus)
	Insert(SearchKey, StoredObject) InsertStatus
	Delete(SearchKey) (StoredObject, DeleteStatus)
}

type IntervalSearchTree interface {
	SearchTree
	FindInterval(SearchKey, SearchKey) []StoredObject
}
