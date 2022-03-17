package slices

import "testing"

func TestDelete(t *testing.T) {
	t.Run("Delete from slice", func(t *testing.T) {
		s := []int{0, 1, 2, 3}
		s = Delete(s, 0)
		if len(s) != 3 {
			t.Fatalf("could not delete data")
		}
		if s[0] != 1 {
			t.Fatalf("could not delete data")
		}
	})

	t.Run("Delete out of index", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("should've not panic")
			}
		}()

		s := []int{0, 1, 2, 3}
		s = Delete(s, 4)
		if len(s) != 4 {
			t.Fatalf("could not delete data")
		}
		if s[0] != 0 {
			t.Fatalf("could not delete data")
		}
	})

	t.Run("On empty slice", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("should've not panic")
			}
		}()

		var s []struct{}
		_ = Delete(s, 0)
	})
}

func TestInsert(t *testing.T) {
	t.Run("Insert another slice from beginning", func(t *testing.T) {
		s1 := []int{0, 1, 2, 3}
		s2 := []int{4, 5}
		s1 = Insert(s1, 0, s2...)
		if len(s1) != 6 {
			t.Fatalf("should've insert slice")
		}
		if s1[1] != 5 {
			t.Fatalf("should've insert slice correctly")
		}
	})

	t.Run("Insert another slice from out of index", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("should've not panic")
			}
		}()

		s1 := []int{0, 1, 2, 3}
		s2 := []int{4, 5}
		s1 = Insert(s1, 5, s2...)
		if len(s1) != 6 {
			t.Fatalf("should've insert slice")
		}
		if s1[3] != 4 {
			t.Fatalf("should've insert slice correctly")
		}
	})

	t.Run("On empty slice", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("should've not panic")
			}
		}()

		var s1 []int
		s2 := []int{4, 5}
		s1 = Insert(s1, 5, s2...)
		if len(s1) != 2 {
			t.Fatalf("should've insert slice")
		}
		if s1[1] != 5 {
			t.Fatalf("should've insert slice correctly")
		}
	})
}

func TestDuplicate(t *testing.T) {
	t.Run("Duplicate a slice", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("should've not panic")
			}
		}()

		s1 := []int{0, 1, 2, 3}
		s2 := Duplicate(s1)
		if len(s1) != len(s2) {
			t.Fatalf("should've duplicate slice")
		}
		for i, e := range s1 {
			if s2[i] != e {
				t.Fatalf("should've duplicated in correct order")
			}
		}
	})

	t.Run("Duplicate an empty slice", func(t *testing.T) {
		var s1 []int
		s2 := Duplicate(s1)
		if len(s1) != len(s2) {
			t.Fatalf("should've duplicate slice")
		}
	})
}
