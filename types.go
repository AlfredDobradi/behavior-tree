package bt

import "log/slog"

type Status uint8

const (
	StatusFailure Status = iota
	StatusSuccess
)

type SequenceKind uint8

const (
	SequenceKindFirstSuccess SequenceKind = iota
	SequenceKindAllSuccess
)

type RHS struct {
	Value any
}

type LHS struct {
	Value any
}

type Condition struct {
	LHS        LHS
	RHS        RHS
	Comparator func(*LHS, *RHS) Status
}

func (c *Condition) Evaluate() Status {
	return c.Comparator(&c.LHS, &c.RHS)
}

type N interface {
	Name() string
	Condition() *Condition
	OnSuccess() *Sequence
	Execute() Status
}

type Node struct {
	Name      string
	Condition *Condition
	OnSuccess *Sequence
	// OnFailure *Sequence
}

func (n *Node) Execute() Status {
	// Placeholder for actual execution logic
	return StatusSuccess
}

type Sequence struct {
	Name       string
	Nodes      []N
	LastStatus Status
	Kind       SequenceKind
	OnSuccess  *Sequence
	OnFailure  *Sequence
}

func (s *Sequence) Execute() Status {
	var nextSequence *Sequence
	for _, node := range s.Nodes {
		if conditionMet := node.Condition().Evaluate(); conditionMet == StatusFailure {
			slog.Info("Condition failed for node", "name", node.Name())
			s.LastStatus = StatusFailure
		} else {
			slog.Info("Condition met for node", "name", node.Name())
			s.LastStatus = node.Execute()
		}

		if s.LastStatus == StatusFailure && s.Kind == SequenceKindAllSuccess {
			if s.OnFailure != nil {
				nextSequence = s.OnFailure
			}
			break
		} else if s.Kind == SequenceKindFirstSuccess && s.LastStatus == StatusSuccess {
			if node.OnSuccess() != nil {
				nextSequence = node.OnSuccess()
			}
			break
		}
	}

	if nextSequence != nil {
		slog.Info("Transitioning to next sequence", "name", nextSequence.Name)
		return nextSequence.Execute()
	}

	if s.LastStatus == StatusSuccess && s.OnSuccess != nil {
		slog.Info("Transitioning to onSuccess sequence", "name", s.OnSuccess.Name)
		return s.OnSuccess.Execute()
	}

	slog.Info("No next sequence found, returning status", "sequence", s.Name, "status", s.LastStatus)
	return s.LastStatus
}
