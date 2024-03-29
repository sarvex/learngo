package main

import (
	"sort"
	"strings"
)

// product is now *product because sorting will unnecessarily copy the product values

type list []*product

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery 🚚.\n"
	}

	var str strings.Builder
	for _, p := range l {
		// fmt.Printf("* %s\n", p)
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}
	return str.String()
}

func (l list) discount(ratio float64) {
	for _, p := range l {
		p.discount(ratio)
	}
}

// implementation of the sort.Interface:

// by default `list` sorts by `title`.
func (l list) Len() int           { return len(l) }
func (l list) Less(i, j int) bool { return l[i].title < l[j].title }
func (l list) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

// byRelease sorts by product release dates.
type byRelease struct {
	// byRelease embeds `list` and reuses list's Len and Swap methods.
	list
}

// Less takes priority over the Less method of the `list`.
// `sort.Sort` will first call this method instead of the `list`'s Less method.
func (bp byRelease) Less(i, j int) bool {
	// `Before()` accepts a `time.Time` but `released` is not `time.Time`.
	return bp.list[i].released.Before(bp.list[j].released.Time)
}

/*
Anonymous embedding means auto-forwarding method calls to the embedded value:

func (bp byRelease) Len() int      { return bp.list.Len() }
func (bp byRelease) Swap(i, j int) { bp.list.Swap(i, j)   }
*/

func byReleaseDate(l list) sort.Interface {
	return byRelease{l}
}
