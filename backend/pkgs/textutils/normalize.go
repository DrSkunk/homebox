package textutils

import (
	"strings"
	"unicode"
	"golang.org/x/text/unicode/norm"
)

// RemoveAccents removes accents from text by normalizing Unicode characters
// and removing diacritical marks. This allows for accent-insensitive search.
//
// Example:
// - "electrónica" becomes "electronica"
// - "café" becomes "cafe"
// - "père" becomes "pere"
func RemoveAccents(text string) string {
       // Alleen diacritics van Latijnse tekens verwijderen, niet van Cyrillisch
       var b strings.Builder
       runes := []rune(norm.NFD.String(text))
       for i := 0; i < len(runes); i++ {
	       r := runes[i]
	       if unicode.In(r, unicode.Mn) {
		       // Check basisletter (vorige rune)
		       if i > 0 {
			       base := runes[i-1]
			       if unicode.In(base, unicode.Latin) {
				       // diacritic op Latijns: verwijderen
				       continue
			       }
		       }
		       // diacritic op niet-Latijns: behouden
	       }
	       b.WriteRune(r)
       }
       return norm.NFC.String(b.String())
}

// NormalizeSearchQuery normalizes a search query for accent-insensitive matching.
// This function removes accents and converts to lowercase for consistent search behavior.
func NormalizeSearchQuery(query string) string {
	normalized := RemoveAccents(query)
	return strings.ToLower(normalized)
}
