package interfaces

// Equitable specifies method that a type must implement in order to be able to check two members of the
// type for equality. The method equals must be an equivalence relation.
type Equitable[T any] interface {
	Equals(other T) bool
}

// Hashable specifies methods that a type must implement in order to be able to generate a hash code for a member of the type.
// If x and y are of type T and satisy x.Equals(y) then they must also sastify x.HashCode() == y.HashCode().
type Hashable[T any] interface {
	HashCode() int
	Equitable[T] // see .
}

// Comparable specifies methods that a type must implement in order to allow ordering between 2 members of that type.
// The Less method must be transitive.
type Comparable[T any] interface {
	Equitable[T]
	Less(other T) bool
}

// Collection a blue print for which methods a collection must implement. Collection generally referes to linear
// data structures i.e linked lists, queue and so on.
type Collection[T Equitable[T]] interface {
	Iterable[T]              // Returns an iterator for iterating through the collection.
	Add(e T) bool            // Adds element e to the collection.
	AddAll(c Iterable[T])    // Adds all elements from another collection into the collection.
	Len() int                // Returns the size (number of items) stored in the collection.
	Contains(e T) bool       // Checks if the element e is a member of the collection.
	Remove(e T) bool         // Tries to remove a specified element in the collection. It removes the first occurence of the element.
	RemoveAll(c Iterable[T]) // // Removes all elements from another collections that appear in the collection. This removes first occurence!
	Empty() bool             // Checks if the collection contains any elements.
	Clear()                  // Removes all elements in the collection.
}
