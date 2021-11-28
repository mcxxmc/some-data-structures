package common

// Template the interface to be implemented
//
// can be used as:
//
// 1. Val in trees
type Template interface {
	String() string     // stringify
	Copy() interface{}  // makes a deep copy
}
