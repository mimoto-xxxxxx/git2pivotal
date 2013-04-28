package pivotal

import (
	"testing"
)

func TestFinderStandardUrl(t *testing.T) {
	ret := FindStoryFromString(`https://www.pivotaltracker.com/story/show/123456`)
	if len(ret) != 1 {
		t.Fail()
		return
	}
	if ret[0].Id != "123456" {
		t.Fail()
		return
	}
}

func TestFinderShortForm(t *testing.T) {
	ret := FindStoryFromString(`pivotal:123456`)
	if len(ret) != 1 {
		t.Fail()
		return
	}
	if ret[0].Id != "123456" {
		t.Fail()
		return
	}
}

func TestFinderMultiple(t *testing.T) {
	ret := FindStoryFromString(`https://www.pivotaltracker.com/story/show/012345
pivotal:123456pivotal:234567
pivotal:345678`)
	if len(ret) != 4 {
		t.Fail()
		return
	}
	correct := []string{"012345", "123456", "234567", "345678"}
	for i, v := range ret {
		if correct[i] != v.Id {
			t.Fail()
			return
		}
	}
}

func TestRemove(t *testing.T) {
	ret := RemoveStoryFromString(`https://www.pivotaltracker.com/story/show/012345
pivotal:123456pivotal:234567 hello
pivotal:345678`)
	if ret != "hello" {
		t.Fail()
		return
	}
}
