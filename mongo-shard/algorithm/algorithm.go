package algorithm

// This interface will support different types of hash alorithms
type hashAlgorithm interface {
	hash(data []byte) int
}
