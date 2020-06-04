package mfstrarrayutil

import (
	"strconv"
	"strings"
	"testing"
)

func TestArrayIndexStr(t *testing.T) {

	arr := []string{"zero", "one", "two", "three", "four", "five"}

	for idx, val := range arr {

		got := StringArrayIndex(arr, val)
		if got != idx {

			t.Errorf("Expected \"%d\".  Got \"%d\".", idx, got)
		}
	}

	got := StringArrayIndex(arr, "notthere")
	if got != -1 {

		t.Errorf("Expected \"-1\".  Got \"%d\".", got)
	}
}

func TestStringArrayContains(t *testing.T) {

	arr := []string{"zero", "one", "two", "three", "four", "five"}

	for _, val := range arr {

		got := StringArrayContains(arr, val)
		if !got {

			t.Error("Expected true.  Got false.")
		}
	}

	arrfalse := []string{"orange", "apple", "pear", "banana"}

	for _, val := range arrfalse {

		got := StringArrayContains(arr, val)
		if got {

			t.Error("Expected false.  Got true.")
		}
	}
}

func TestStringArrayMap(t *testing.T) {

	arr := []string{"orange", "apple", "pear", "banana"}
	got := StringArrayMap(arr, func(s string) string {

		return "xxx" + s
	})

	if len(got) != len(arr) {

		t.Errorf("Expected input and output slice lengths to be equal.")
	}

	for idx, val := range got {

		if "xxx" + arr[idx] != val {

			t.Errorf("Expected \"xxx%s\".  Got \"%s\".", arr[idx], got)
		}
	}
}

func TestStringArrayFilter(t *testing.T) {

	arr := []string{"orange", "apple", "pear", "banana", "peach", "pomelo", "bacon"}
	expected := []string{"pear", "peach", "pomelo"}

	got := StringArrayFilter(arr, func(s string) bool {

		return strings.HasPrefix(s, "p")
	})

	if len(got) != len(expected) {

		t.Fatal("Expected got len to equal expected len.")
	}

	for idx, val := range expected {

		if got[idx] != val {

			t.Errorf("Expected \"%s\" to equal \"%s\".", val, got[idx])
		}
	}
}

func TestStringArrayCount(t *testing.T) {

	arr := []string{"orange", "apple", "pear", "banana", "peach", "pomelo", "bacon"}
	expected := 3

	got := StringArrayCount(arr, func(s string) bool {

		return strings.HasPrefix(s, "p")
	})

	if expected != got {

		t.Errorf("Expected \"%d\".  Got \"%d\".", expected, got)
	}

	got = StringArrayCount(arr, func(s string) bool {

		return s == "nothere"
	})

	if 0 != got {

		t.Errorf("Expected \"%d\".  Got \"%d\".", expected, got)
	}
}

func TestStringArrayAny(t *testing.T) {

	arr := []string{"orange", "apple", "pear", "banana", "peach", "pomelo", "bacon"}

	got := StringArrayAny(arr, func (s string) bool {

		return s == "banana"
	})

	if !got {

		t.Error("Expected true.")
	}

	got = StringArrayAny(arr, func(s string) bool {

		return s == "notthere"
	})

	if got {

		t.Error("Expected false")
	}
}

func TestStringArrayAll(t *testing.T) {

	arr := []string{"1", "2", "3", "4", "5", "6", "7"}

	got := StringArrayAll(arr, func (s string) bool {

		i, _ := strconv.Atoi(s)

		return i < 8
	})

	if !got {

		t.Error("Expected true.")
	}

	got = StringArrayAll(arr, func(s string) bool {

		i, _ := strconv.Atoi(s)

		return i < 7
	})

	if got {

		t.Error("Expected false")
	}
}