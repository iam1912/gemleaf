package models

type LeafSegment struct {
	Cursor uint64
	Max    uint64
	Min    uint64
	InitOk bool
}

func NewLeafSegment(leaf *Leaf) *LeafSegment {
	return &LeafSegment{
		Cursor: leaf.MaxID,
		Max:    leaf.MaxID + uint64(leaf.Step) - 1,
		Min:    leaf.MaxID,
		InitOk: true,
	}
}

func (s *LeafSegment) HasID() bool {
	if s.InitOk && s.Cursor <= s.Max {
		return true
	}
	return false
}
