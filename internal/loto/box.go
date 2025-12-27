package loto

import (
	"slices"

	"github.com/kawana77b/loto/internal/util"
)

type Box struct {
	items []int
}

// NewBox creates a new Box containing integers from min to max (inclusive).
func NewBox(min, max int) *Box {
	items := make([]int, 0, max-min+1)
	for i := min; i <= max; i++ {
		items = append(items, i)
	}
	return &Box{
		items: items,
	}
}

// Length returns the number of items in the box.
func (b *Box) Length() int {
	return len(b.items)
}

// Append adds one or more items to the box.
func (b *Box) Append(item ...int) {
	b.items = append(b.items, item...)
}

// Clear removes all items from the box.
func (b *Box) Clear() {
	b.items = b.items[:0]
}

// Clone creates and returns a deep copy of the box.
func (b *Box) Clone() *Box {
	clonedItems := make([]int, len(b.items))
	copy(clonedItems, b.items)
	return &Box{
		items: clonedItems,
	}
}

// Sort sorts the items in the box in ascending order.
func (b *Box) Sort() {
	slices.Sort(b.items)
}

// Shuffle randomly shuffles the items in the box.
func (b *Box) Shuffle() {
	b.items = util.Shuffle(b.items)
}

// Contains checks if the box contains the specified element.
func (b *Box) Contains(item int) bool {
	return slices.Contains(b.items, item)
}

// PickN randomly selects and returns n unique items from the box.
func (b *Box) PickN(n int) []int {
	if n <= 0 {
		return []int{}
	}
	shuffled := util.Shuffle(b.items)
	return shuffled[:n]
}

// PickDupN randomly selects and returns n items from the box, allowing for duplicates.
func (b *Box) PickDupN(n int) []int {
	if n <= 0 {
		return []int{}
	}
	result := make([]int, n)
	for i := range result {
		result[i], _ = util.RandomPick(b.items)
	}
	return result
}
