package main

/*
	@description: check given element is at slice
*/
func SliceContains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

/*
	@description: UNUSED allows unused variables to be included in Go programs
*/
func UNUSED(x ...interface{}) {}
