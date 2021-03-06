package board

import "fmt"

// Board provides access to the board that the game is played on.
// Put makes marks at the requested position.
// IsWritable finds out if a position is stil free.
// Get returns a copy of the data.
type Board interface {
	Put(p Position, player Player) (ok bool, reason string)
	IsWritable(p Position) (ok bool, reason string)
	Get() Data
}

// Data is the board the game is played on.
// The size is customizable but the board is always quadratic.
// Marks stores the moves the players have made.
// 0 means the board is empty, all other values represent the player numbers.
type Data struct {
	Marks Marks
	Size  int
}

// Marks stores the players moves on the board.
// The (sequential) array slice is interpreted as a line-wise traversal of the
// board.
type Marks []Player

// Player is the ID of a player or the number of players.
type Player int

// Position is a position on the board.
// The two values are the x- and y-coordinate.
// The top-left corner is (0, 0).
type Position [2]int

// NewPosition creates a Position from an index in a Marks array and the size
// (edge length) of the board.
func NewPosition(index int, size int) Position {
	return Position{index % size, index / size}
}

// ToIndex returns the array index of a Position.
func (p Position) ToIndex(s int) int {
	return p[1]*s + p[0]
}

// NewData constructs Data with size s*s
func NewData(size int) Data {
	return Data{
		make(Marks, size*size),
		size,
	}
}

// Put makes a mark for a player on the board.
// If the position is not on the board or a mark has already been made at the
// specified position, Put returns an error.
func (bo Data) Put(p Position, player Player) (ok bool, reason string) {
	if ok, reason := bo.IsWritable(p); !ok {
		return ok, reason
	}

	bo.Marks[p.ToIndex(bo.Size)] = player
	return true, ""
}

// IsWritable checks is a position can be written in.
// It has to be within the limits of the board and empty.
// If the position is not writable a reason is given.
func (bo Data) IsWritable(p Position) (ok bool, reason string) {
	// TODO try to get rid of this function
	if p[0] < 0 || p[0] >= bo.Size || p[1] < 0 || p[1] >= bo.Size {
		return false, fmt.Sprintf("position out of range, board has size %vx%v",
			bo.Size, bo.Size)
	}

	if bo.Marks[p.ToIndex(bo.Size)] != 0 {
		return false, fmt.Sprintf("position is not empty")
	}

	return true, ""
}

// Get returns a copy of the board.
func (bo Data) Get() Data {
	// TODO test
	result := Data{Marks: make(Marks, bo.Size*bo.Size), Size: bo.Size}

	for i, v := range bo.Marks {
		result.Marks[i] = v
	}

	return result
}

// Equal returns true when the Marks in the parameter board are equal to the
// one of the receiver.
func (marks Marks) Equal(other Marks) bool {
	// TODO this function is kinda ugly, since it is only used in tests
	if len(marks) != len(other) {
		return false
	}
	for i := 0; i < len(marks); i++ {
		if marks[i] != other[i] {
			return false
		}
	}
	return true
}
