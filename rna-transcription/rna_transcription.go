//package strand
package strand

// ToRNA retunrs the corresponding RNA for a given DNA sequence(string)
func ToRNA(dna string) string {
	var nucleotide = map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}
	var rna string
	for _, v := range dna {
		rna += nucleotide[string(v)]
	}
	return rna
}
