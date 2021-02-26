package useragent

import (
	"regexp"
	"sync"
	"testing"
)

func TestDesktop(t *testing.T) {
	ua, err := Desktop()
	if err != nil {
		t.Error(err)
	}
	t.Error(ua)

	match, _ := regexp.MatchString("Mozilla/5.0", ua)
	if match == false {
		t.Errorf("Proxy should return a User Agent String. Instead of: %v", ua)
	}
}

func TestDesktopConcurrent(t *testing.T) {
	n := 100
	var uas []string
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			ua, err := Desktop()
			if err != nil {
				t.Error(err)
				return
			}
			uas = append(uas, ua)
		}()
	}
	wg.Wait()

	for i, ua := range uas {
		match, _ := regexp.MatchString("Mozilla/5.0", ua)
		if match == false {
			t.Errorf("Proxy should return a User Agent String on record %v. Instead of: %v", i, ua)
		}
	}
}
