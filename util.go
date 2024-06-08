package util

import (
	"cmp"
	"encoding/json"
)

/*
Zero returns unassigned value of Type T
*/
func Zero[T any]() T {
	var t T
	return t
}

/*
UnmarshalReturn json.Unmarshal's to type T and returns it
as a pointer.
*/
func UnmarshalReturn[T any](data []byte) (*T, error) {
	var t T
	var err error

	if err = json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

type Converter[T any, V any] func(T) V
type Comparer[T, V any] func(T, V) int

/*
IdentityFunc returns identity function of type T
*/
func IdentityFunc[T any]() func(T) T {
	return func(t T) T {
		return t
	}
}

/*
SliceConvert converts slices with given converter function.
*/
func SliceConvert[T any, V any](ts []T, converter Converter[T, V]) (vs []V) {
	vs = make([]V, len(ts))
	for i := range len(ts) {
		vs[i] = converter(ts[i])
	}
	return
}

/*
MergeOrdered concatinates and sorts input1 and input2 and returns it.
Parameters input1 and input2 must be sorted slices.
*/
func MergeOrdered[C cmp.Ordered, IN1, IN2, OUT []C](
	input1 IN1,
	input2 IN2,
) (out OUT) {
	out = make(OUT, len(input1)+len(input2))
	var i, j int = 0, 0
	for i < len(input1) && j < len(input2) {
		if input2[j] > input1[i] {
			out[i+j] = input1[i]
			i++
		} else {
			out[i+j] = input2[j]
			j++
		}
	}
	if j == len(input2) {
		for ; i < len(input1); i++ {
			out[i+j] = input1[i]
		}
	} else {
		for ; j < len(input2); j++ {
			out[i+j] = input2[j]
		}
	}
	return
}

/*
MergeWithFunc merges input1 and input2 by comparing them with cmp function
and constructing and returning a merged slice by using conv1, conv2 Converter functions.
*/
func MergeWithFunc[IN1, IN2, OUT any](
	input1 []IN1,
	input2 []IN2,
	cmp Comparer[IN1, IN2],
	conv1 Converter[IN1, OUT],
	conv2 Converter[IN2, OUT],
) (out []OUT) {
	out = make([]OUT, len(input1)+len(input2))
	var i, j int = 0, 0
	for i < len(input1) && j < len(input2) {
		if cmp(input1[i], input2[j]) < 0 {
			out[i+j] = conv1(input1[i])
			i++
		} else {
			out[i+j] = conv2(input2[j])
			j++
		}
	}
	if j == len(input2) {
		for ; i < len(input1); i++ {
			out[i+j] = conv1(input1[i])
		}
	} else {
		for ; j < len(input2); j++ {
			out[i+j] = conv2(input2[j])
		}
	}
	return
}

/*
MergeWithFuncSimple merges input1 and input2 by comparing them with cmp function
and constructing and returning a merged slice.
*/
func MergeWithFuncSimple[T any](in1 []T, in2 []T, cmp Comparer[T, T]) []T {
	return MergeWithFunc(in1, in2, cmp, IdentityFunc[T](), IdentityFunc[T]())
}
