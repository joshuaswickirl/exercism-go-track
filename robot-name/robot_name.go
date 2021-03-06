package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Robot has a name
type Robot struct {
	name string
}

const namespaceCapacity = 26 * 26 * 10 * 10 * 10

var usedNames = map[string]bool{}

// Name returns the Robot's unique name
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(usedNames) == namespaceCapacity {
		return "", errors.New("namespace is exhausted")
	}
	r.name = generateName()
	for usedNames[r.name] {
		r.name = generateName()
	}
	usedNames[r.name] = true
	return r.name, nil
}

// Reset updates a Robot's name
func (r *Robot) Reset() {
	r.name = ""
}

func generateName() string {
	l1 := rand.Intn(26) + 'A'
	l2 := rand.Intn(26) + 'A'
	n := rand.Intn(1000)
	return fmt.Sprintf("%s%s%03d", string(l1), string(l2), n)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
