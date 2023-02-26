package linkedlist

// type ImmutableLinkedList[T comparable] struct {
// 	list forwardlist.ImmutableForwadList[T]
// }

// // Of creates a list with the given elements.
// func ImmutableOf[T comparable](elements ...T) ImmutableLinkedList[T] {
// 	list := forwardlist.ImmutableOf[T](elements...)
// 	return ImmutableLinkedList[T]{list: list}
// }

// // AddSlice unsupported operation.
// func (list ImmutableLinkedList[T]) AddSlice(s []T) bool {
// 	panic(errors.UnsupportedOperation("AddSlice", "ImmutableLinkedList"))
// }

// // AddAll unsupported operation.
// func (list ImmutableLinkedList[T]) AddAll(iterable collections.Iterable[T]) bool {
// 	panic(errors.UnsupportedOperation("AddAll", "ImmutableLinkedList"))
// }

// // Empty returns true if the list contains no elements.
// func (list ImmutableLinkedList[T]) Empty() bool {
// 	return list.list.Empty()
// }

// // Add unsupported operation.
// func (list ImmutableLinkedList[T]) Add(e T) bool {
// 	panic(errors.UnsupportedOperation("Add", "ImmutableLinkedList"))
// }

// // AddAt unsupported operation.
// func (list ImmutableLinkedList[T]) AddAt(i int, e T) {
// 	panic(errors.UnsupportedOperation("AddAt", "ImmutableLinkedList"))
// }

// // At returns the element at the specified index in the list.
// func (list ImmutableLinkedList[T]) At(i int) T {
// 	return list.list.At(i)
// }

// // Set unsupported operation.
// func (list ImmutableLinkedList[T]) Set(i int, e T) T {
// 	panic(errors.UnsupportedOperation("Set", "ImmutableLinkedList"))
// }

// // Len returns the number of elements in the list.
// func (list ImmutableLinkedList[T]) Len() int {
// 	return list.list.Len()
// }

// // Clear unsupported operation.
// func (list ImmutableLinkedList[T]) Clear() {
// 	panic(errors.UnsupportedOperation("Clear", "ImmutableLinkedList"))
// }

// // Contains returns true if the list contains the specified element.
// func (list ImmutableLinkedList[T]) Contains(e T) bool {
// 	return list.list.Contains(e)
// }

// // Remove unsupported operation.
// func (list ImmutableLinkedList[T]) Remove(e T) bool {
// 	panic(errors.UnsupportedOperation("Remove", "ImmutableLinkedList"))
// }

// // RemoveAt unsupported operation.
// func (list ImmutableLinkedList[T]) RemoveAt(i int) T {
// 	panic(errors.UnsupportedOperation("RemoveAt", "ImmutableLinkedList"))
// }

// // RemoveIf unsupported operation.
// func (list ImmutableLinkedList[T]) RemoveIf(f func(T) bool) bool {
// 	panic(errors.UnsupportedOperation("RemoveIf", "ImmutableLinkedList"))
// }

// // RemoveAll unsupported operation
// func (list ImmutableLinkedList[T]) RemoveAll(iterable collections.Iterable[T]) bool {
// 	panic(errors.UnsupportedOperation("RemoveAll", "ImmutableLinkedList"))
// }

// // RetainAll unsupported operation.
// func (list ImmutableLinkedList[T]) RetainAll(c collections.Collection[T]) bool {
// 	panic(errors.UnsupportedOperation("RetainAll", "ImmutableLinkedList"))
// }

// // RemoveSlice unsupported operation.
// func (list ImmutableLinkedList[T]) RemoveSlice(s []T) bool {
// 	panic(errors.UnsupportedOperation("RemoveSlice", "ImmutableLinkedList"))
// }

// // ToSlice returns a slice containing the elements of the list.
// func (list ImmutableLinkedList[T]) ToSlice() []T {
// 	return list.list.ToSlice()
// }

// // ForEach performs the given action for each element of the list.
// func (list ImmutableLinkedList[T]) ForEach(f func(T)) {
// 	list.list.ForEach(f)
// }

// // SubList returns a view of the portion of the list between the specified start and end indices (exclusive).
// func (list ImmutableLinkedList[T]) SubList(start int, end int) ImmutableLinkedList[T] {
// 	if start < 0 || start >= list.Len() {
// 		panic(errors.IndexOutOfBounds(start, list.Len()))
// 	} else if end < 0 || end > list.Len() {
// 		panic(errors.IndexOutOfBounds(end, list.Len()))
// 	} else if start > end {
// 		panic(errors.IndexBoundsOutOfRange(start, end))
// 	} else if start == end {
// 		return ImmutableOf[T]()
// 	}
// 	return ImmutableLinkedList[T]{list: list.list.SubList(start, end)}
// }

// // Equals returns true if the list is equivalent to the given list. Two lists are equal if they have the same size
// // and contain the same elements in the same order.
// func (list ImmutableLinkedList[T]) Equals(other ImmutableLinkedList[T]) bool {
// 	return list.list.Equals(other.list)
// }

// // Iterator returns an iterator over the elements in the list.
// func (list ImmutableLinkedList[T]) Iterator() collections.Iterator[T] {
// 	return list.list.Iterator()
// }

// // String returns the string representation of the list.
// func (list ImmutableLinkedList[T]) String() string {
// 	return fmt.Sprint(list.ToSlice())
// }
