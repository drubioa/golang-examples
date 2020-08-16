package main

import (
	"fmt"
)

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

func toArrayRec(t *Tree,  n int )  []int{
	var a []int
	if t.Left != nil  {
		l := toArrayRec(t.Left, n)
		for i:= 0 ; i < len(l) ; i++ {
			a = append(a, l[i])
		}
	}
	a = append(a, t.Value)
	if t.Right != nil  {
		r := toArrayRec(t.Right, n)
		for i:= 0 ; i < len(r) ; i++ {
			a = append(a, r[i])
		}
	}
	return a
}

func toArray(t *Tree,  ch chan []int, n int ) {
	ch <- toArrayRec(t, n)

}

func equivalent(t1, t2 *Tree) bool{
	ch := make(chan []int, 2)
	go toArray(t1, ch, 1)
	go toArray(t2, ch, 2)
	a:= <-ch
	b := <-ch
	close(ch)

	if len(a) != len(b) {
		return false
	}
	for i := 1; i < len(a) ; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func generateTreeExample1() Tree {
	t := Tree {Value: 3}
	t.Left = &Tree {Value: 1, Left: &Tree{Value: 1}, Right: &Tree{Value: 2}}
	t.Right = &Tree {Value: 8, Left: &Tree{Value: 5}, Right: &Tree{Value: 13}}
	return t
}

func generateTreeExample2() Tree {
	t := Tree {Value: 8}
	t.Left = &Tree {Value: 3, Left: &Tree{Value: 1, Left: &Tree{ Value: 1}, Right: &Tree{Value: 2}}, Right: &Tree{Value: 5}}
	t.Right = &Tree {Value: 13}
	return t
}

func main() {
	t1 := generateTreeExample1();
	t2 := generateTreeExample2();
	fmt.Println(equivalent(&t1, &t2))
}