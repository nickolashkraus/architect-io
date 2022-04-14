package pkg

// PROBLEM STATEMENT:
//
// Create a function. Takes in a string. Returns true if there are duplicate
// characters in the string otherwise returns false.
//
// SOLUTION:
//
// Create a hash set (hash table). Add characters to the hash set if they are
// not already in the hash set, else return true.
//
// NOTE: Hash table lookup is O(1).
//
// FOLLOW-UP:
//
// Q: What happens if the string is infinitely long?
// A: The string cannot be infinitely long, since, assuming a UTF-16 character
// set, once all the characters in the set are exhausted, you must begin to
// duplicate characters. This must occur at at least n+1, where n is the length
// of the character set.

func findDuplicate(s string) bool {
	table := make(map[rune]bool)
	// Loop over string. If rune is in hash table, return false. If not, add rune
	// to hash table.
	for _, elem := range s {
		_, ok := table[elem]
		if ok {
			return true
		} else {
			table[elem] = true
		}
	}
	return false
}
