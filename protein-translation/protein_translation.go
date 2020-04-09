package protein

import (
	"errors"
)

var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// ErrStop represents a STOP codon.
var ErrStop = errors.New("recieved stop codon")

// ErrInvalidBase repesents a erronious codon base input.
var ErrInvalidBase = errors.New("recieved bad codon base")

// FromCodon returns the protein or stop signal as error for a given codon.
func FromCodon(codon string) (string, error) {
	protein, ok := codons[codon]
	if !ok {
		return "", ErrInvalidBase
	} else if protein == "STOP" {
		return "", ErrStop
	}
	return protein, nil
}

// FromRNA returns a slice of protiens for a given RNA strand.
func FromRNA(rna string) ([]string, error) {
	var proteins []string
	for i := 0; i < len(rna); i += 3 {
		protein, err := FromCodon(rna[i : i+3])
		if err == ErrStop {
			break
		} else if err == ErrInvalidBase {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}
