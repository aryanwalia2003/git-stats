package app

// isEmoji checks if a rune is an emoji character.
// Covers common emoji ranges in Unicode.
func isEmoji(r rune) bool {
	return (r >= 0x1F600 && r <= 0x1F64F) || // Emoticons
		(r >= 0x1F300 && r <= 0x1F5FF) || // Misc Symbols
		(r >= 0x1F680 && r <= 0x1F6FF) || // Transport
		(r >= 0x1F900 && r <= 0x1F9FF) || // Supplemental
		(r >= 0x2600 && r <= 0x26FF) || // Misc symbols
		(r >= 0x2700 && r <= 0x27BF) || // Dingbats
		(r >= 0xFE00 && r <= 0xFE0F) || // Variation selectors
		(r >= 0x1FA00 && r <= 0x1FA6F) // Chess, extended-A
}
