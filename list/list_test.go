package list_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"datastructure/list"
)

func TestList_Add(t *testing.T) {
	Convey("I should always should be able to add element to a list", t, func() {
		l := list.New()
		l1 := l.Prepend(2)
		l2 := l1.Prepend(4)
		So(l2, ShouldResemble, list.List{Head: 4, Tail: &list.List{Head: 2, Tail: &list.List{}}})

		Convey("and it should not affect initial list", func() {
			So(l, ShouldResemble, list.List{})
		})

		Convey("and initial list should become new list queue", func() {
			So(l1, ShouldResemble, *l2.Tail)
		})
	})
}

func TestList_Empty(t *testing.T) {
	Convey("I should be able to identify empty list", t, func() {
		So(list.New().Empty(), ShouldBeTrue)
		So(list.New().Prepend(1).Empty(), ShouldBeFalse)
		So(list.New().Prepend(1).Tail.Empty(), ShouldBeTrue)
		So(list.New().Prepend(1).Prepend("t").Prepend(true).Empty(), ShouldBeFalse)
	})
}

func TestList_Pop(t *testing.T) {
	Convey("Given", t, func() {
		Convey("an Empty list", func() {
			l := list.New()
			h, q := l.Pop()
			So(h, ShouldBeNil)
			So(q.Empty(), ShouldBeTrue)
		})

		Convey("a known list", func() {
			l := list.New().Prepend(5).Prepend(4)
			l2 := l.Prepend(3)
			h, q := l2.Pop()
			So(h, ShouldEqual, 3)
			So(q, ShouldResemble, l)
			h2, _ := q.Pop()
			So(h2, ShouldEqual, 4)
		})
	})
}

func TestList_Map(t *testing.T) {
	Convey("Given a list", t, func() {
		l := list.New().Prepend(4).Prepend(3).Prepend(6)

		expectedL := l

		times3 := func(e interface{}) interface{} { return e.(int) * 3 }
		times2 := func(e interface{}) interface{} { return e.(int) * 2 }
		times5 := func(e interface{}) interface{} { return e.(int) * 5 }

		Convey("When apply function to all element of a list", func() {
			lT3 := l.Map(times3)
			lT2 := l.Map(times2)
			lT5 := l.Map(times5)

			expectedT3 := list.New().Prepend(12).Prepend(9).Prepend(18)
			expectedT2 := list.New().Prepend(8).Prepend(6).Prepend(12)
			expectedT5 := list.New().Prepend(20).Prepend(15).Prepend(30)

			Convey("it should return a new list where function was applied", func() {
				So(l, ShouldResemble, expectedL)
				So(lT2, ShouldResemble, &expectedT2)
				So(lT3, ShouldResemble, &expectedT3)
				So(lT5, ShouldResemble, &expectedT5)
			})
		})
	})
}

func TestList_Map_String(t *testing.T) {
	Convey("Given ", t, func() {
		Convey("an Empty list", func() {
			l := list.New()
			Convey("when I apply Map function", func() {
				newL := l.Map(func(e interface{}) interface{} { return e })
				So(newL.Empty(), ShouldBeTrue)
			})
		})

		l := list.New()
		addWorld := func(e interface{}) interface{} { return e.(string) + " world!" }

		Convey("a single element list", func() {
			l = l.Prepend("hello")
			expectedL := list.New().Prepend("hello world!")
			Convey("when I apply Map function", func() {
				newL := l.Map(addWorld)
				So(*newL, ShouldResemble, expectedL)
			})
		})
		Convey("a  list", func() {
			l = l.Prepend("new").Prepend("my").Prepend("hello")
			expectedL := list.New().Prepend("new world!").Prepend("my world!").Prepend("hello world!")
			Convey("when I apply Map function", func() {
				newL := l.Map(addWorld)
				So(newL, ShouldResemble, expectedL)
			})
		})

	})
}

func TestList_AddSorted(t *testing.T) {
	Convey("Given a sorted list", t, func() {
		tmpList := list.New().Prepend(2).Prepend(4).Prepend(5).Prepend(9)
		expectedList := list.New().
			Prepend(1).Prepend(2).Prepend(3).Prepend(4).Prepend(5).
			Prepend(6).Prepend(9).Prepend(10)

		l := &tmpList
		comp := func(a, b interface{}) bool {
			return a.(int) > b.(int)
		}

		Convey("When I add element with AddSorted", func() {
			l = l.AddSorted(1, comp)
			l = l.AddSorted(10, comp)
			l = l.AddSorted(6, comp)
			l = l.AddSorted(3, comp)

			Convey("it should be added and new list should be sorted", func() {
				So(*l, ShouldResemble, expectedList)
			})
		})
	})
}

func TestList_BullSort(t *testing.T) {
	Convey("Given a list", t, func() {
		tmpList := list.New().Prepend(1).Prepend(4).Prepend(6).Prepend(9).Prepend(3).Prepend(5).Prepend(2).Prepend(10)

		expectedSorted := list.New().
			Prepend(1).Prepend(2).Prepend(3).Prepend(4).Prepend(5).
			Prepend(6).Prepend(9).Prepend(10)

		l := &tmpList
		comp := func(a, b interface{}) bool {
			return a.(int) > b.(int)
		}

		Convey("empty or single element, should stay as his", func() {
			l := list.New()
			So(l.BullSort(comp), ShouldResemble, &l)

			l = list.New().Prepend(1)
			So(l.BullSort(comp), ShouldResemble, &l)
		})

		Convey("not sorted, I should be able to sort it", func() {
			l = l.BullSort(comp)
			So(*l, ShouldResemble, expectedSorted)
		})

		Convey("sorted, I should be able to sort it", func() {
			l = l.BullSort(comp)
			l = l.BullSort(comp)
			So(*l, ShouldResemble, expectedSorted)
		})
	})
}

func TestNewList(t *testing.T) {
	Convey("When I create a list", t, func() {
		newList := list.New()
		Convey("I expect it to be empty", func() {
			So(newList.Empty(), ShouldBeTrue)
		})
	})
}
