package cmap

import (
	. "github.com/karlseguin/expect"
	"strconv"
	"testing"
)

type CMapTests struct{}

func Test_CMap(t *testing.T) {
	Expectify(new(CMapTests), t)
}

func (_ *CMapTests) GetsAndSetsAndDeletesValues() {
	c := New()
	for i := 0; i < 1000; i++ {
		c.Set(strconv.Itoa(i), i)
	}

	for i := 0; i < 1000; i++ {
		Expect(c.Get(strconv.Itoa(i))).To.Equal(i, true)
	}

	for i := 0; i < 1000; i++ {
		if 1&i == 0 {
			c.Delete(strconv.Itoa(i))
		}
	}

	for i := 0; i < 1000; i++ {
		if 1&i == 0 {
			Expect(c.Get(strconv.Itoa(i))).To.Equal(nil, false)
		} else {
			Expect(c.Get(strconv.Itoa(i))).To.Equal(i, true)
		}
	}
}

func (_ *CMapTests) SetOverwrites() {
	c := New()
	c.Set("spice", "flow")
	c.Set("spice", "worm")
	Expect(c.Get("spice")).ToEqual("worm", true)
}

func (_ *CMapTests) DeleteMissingKeyOk() {
	c := New()
	c.Set("spice", "flow")
	c.Delete("paul")
	Expect(c.Get("spice")).ToEqual("flow", true)
	Expect(c.Get("paul")).ToEqual(nil, false)
}

func (_ *CMapTests) Len() {
	c := New()
	Expect(c.Len()).To.Equal(0)

	for i := 0; i < 1000; i++ {
		c.Set(strconv.Itoa(i), i)
	}
	Expect(c.Len()).To.Equal(1000)

	for i := 0; i < 1000; i++ {
		if 1&i == 0 {
			c.Delete(strconv.Itoa(i))
		}
	}
	Expect(c.Len()).To.Equal(500)
}
