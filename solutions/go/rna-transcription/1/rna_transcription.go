package strand

import "strings"

func ToRNA(dna string) string {
	if len(dna) < 1 {
        return ""
    }

    var rna strings.Builder

    rna.Grow(len(dna))
    
    for _, char := range dna{
        switch char {
            case 'G':
            	 rna.WriteRune('C')
            case 'T':
            	 rna.WriteRune('A')
            case 'A':
            	 rna.WriteRune('U')
            case 'C':
            	 rna.WriteRune('G')
        }
    }
    
    return rna.String()
}
