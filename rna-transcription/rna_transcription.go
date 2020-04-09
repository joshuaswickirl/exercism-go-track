package strand

import "strings"

var r = strings.NewReplacer("G", "C", "C", "G", "T", "A", "A", "U")

// ToRNA returns the RNA complement to a given DNA strand.
func ToRNA(dna string) string {
	return r.Replace(dna)
}
