package main

// type Point2D struct {
// 	x int
// 	y int
// }

// type Option[T any] struct {
// 	data T
// 	some bool
// }

// func none[T any]() Option[T] {
// 	return Option[T]{
// 		some: false,
// 	}
// }

// func some[T any](value T) Option[T] {
// 	return Option[T]{
// 		some: true,
// 		data: value,
// 	}
// }

// func (opt Option[T]) unwrap() T {
// 	if opt.isSome() {
// 		return opt.data
// 	} else {
// 		log.Fatal("Failed to unwrap")

// 		//unreachable
// 		return opt.data
// 	}
// }

// func (opt Option[T]) isSome() bool {
// 	return opt.some
// }

// func (opt Option[T]) isNone() bool {
// 	return !opt.isSome()
// }
