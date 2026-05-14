package bt

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// func intCompare(lhs *LHS, rhs *RHS) Status {
// 	ll := lhs.Value.(int)
// 	rr := rhs.Value.(int)

// 	if ll <= rr {
// 		return StatusFailure
// 	} else {
// 		return StatusSuccess
// 	}
// }

// func TestSequenceExecution(t *testing.T) {
// 	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn})))

// 	tests := []struct {
// 		name     string
// 		nodes    []*Node
// 		kind     SequenceKind
// 		expected Status
// 	}{
// 		{
// 			name: "FirstSuccess with one successful node",
// 			nodes: []*Node{
// 				{
// 					Name: "Node2",
// 					Condition: &Condition{
// 						LHS:        LHS{Value: 5},
// 						RHS:        RHS{Value: 4},
// 						Comparator: intCompare,
// 					},
// 				},
// 			},
// 			kind:     SequenceKindFirstSuccess,
// 			expected: StatusSuccess,
// 		},
// 		{
// 			name: "FirstSuccess with second node successful",
// 			nodes: []*Node{
// 				{
// 					Name: "Node1",
// 					Condition: &Condition{
// 						LHS:        LHS{Value: 2},
// 						RHS:        RHS{Value: 3},
// 						Comparator: intCompare,
// 					},
// 				},
// 				{
// 					Name: "Node2",
// 					Condition: &Condition{
// 						LHS:        LHS{Value: 5},
// 						RHS:        RHS{Value: 4},
// 						Comparator: intCompare,
// 					},
// 				},
// 			},
// 			kind:     SequenceKindFirstSuccess,
// 			expected: StatusSuccess,
// 		},
// 		{
// 			name: "FirstSuccess with one failing node",
// 			nodes: []*Node{
// 				{
// 					Name: "Node2",
// 					Condition: &Condition{
// 						LHS:        LHS{Value: 5},
// 						RHS:        RHS{Value: 4},
// 						Comparator: intCompare,
// 					},
// 				},
// 			},
// 			kind:     SequenceKindFirstSuccess,
// 			expected: StatusSuccess,
// 		},
// 		{
// 			name: "FirstSuccess with all nodes failing",
// 			nodes: []*Node{
// 				{
// 					Name: "Node1",
// 					Condition: &Condition{
// 						LHS:        LHS{Value: 2},
// 						RHS:        RHS{Value: 3},
// 						Comparator: intCompare,
// 					},
// 				},
// 				{
// 					Name: "Node2",
// 					Condition: &Condition{
// 						LHS:        LHS{Value: 5},
// 						RHS:        RHS{Value: 6},
// 						Comparator: intCompare,
// 					},
// 				},
// 				{
// 					Name: "Node3",
// 					Condition: &Condition{
// 						LHS:        LHS{Value: 1},
// 						RHS:        RHS{Value: 7},
// 						Comparator: intCompare,
// 					},
// 				},
// 			},
// 			kind:     SequenceKindFirstSuccess,
// 			expected: StatusFailure,
// 		},
// 		{
// 			name: "AllSuccess with all nodes successful",
// 			nodes: []*Node{
// 				{
// 					Name: "Node1",
// 					Condition: &Condition{
// 						LHS:        LHS{Value: 4},
// 						RHS:        RHS{Value: 3},
// 						Comparator: intCompare,
// 					},
// 				},
// 				{
// 					Name: "Node2",
// 					Condition: &Condition{
// 						LHS:        LHS{Value: 7},
// 						RHS:        RHS{Value: 6},
// 						Comparator: intCompare,
// 					},
// 				},
// 				{
// 					Name: "Node3",
// 					Condition: &Condition{
// 						LHS:        LHS{Value: 8},
// 						RHS:        RHS{Value: 7},
// 						Comparator: intCompare,
// 					},
// 				},
// 			},
// 			kind:     SequenceKindAllSuccess,
// 			expected: StatusSuccess,
// 		},
// 		{
// 			name: "AllSuccess with a node failing",
// 			nodes: []*Node{
// 				{
// 					Name: "Node1",
// 					Condition: &Condition{
// 						LHS:        LHS{Value: 4},
// 						RHS:        RHS{Value: 3},
// 						Comparator: intCompare,
// 					},
// 				},
// 				{
// 					Name: "Node2",
// 					Condition: &Condition{
// 						LHS:        LHS{Value: 7},
// 						RHS:        RHS{Value: 6},
// 						Comparator: intCompare,
// 					},
// 				},
// 				{
// 					Name: "Node3",
// 					Condition: &Condition{
// 						LHS:        LHS{Value: 8},
// 						RHS:        RHS{Value: 9},
// 						Comparator: intCompare,
// 					},
// 				},
// 			},
// 			kind:     SequenceKindAllSuccess,
// 			expected: StatusFailure,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			seq := &Sequence{
// 				Nodes: tt.nodes,
// 				Kind:  tt.kind,
// 			}
// 			result := seq.Execute()
// 			require.Equal(t, tt.expected, result)
// 		})
// 	}
// }

// func TestExecuteNextSequence(t *testing.T) {
// 	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn})))

// 	s2 := &Sequence{
// 		Name: "Sequence2",
// 		Nodes: []*Node{
// 			{
// 				Name: "Node2",
// 				Condition: &Condition{
// 					LHS:        LHS{Value: 5},
// 					RHS:        RHS{Value: 4},
// 					Comparator: intCompare,
// 				},
// 			},
// 		},
// 		Kind: SequenceKindFirstSuccess,
// 	}

// 	s1 := &Sequence{
// 		Name: "Sequence1",
// 		Nodes: []*Node{
// 			{
// 				Name: "Node1",
// 				Condition: &Condition{
// 					LHS:        LHS{Value: 4},
// 					RHS:        RHS{Value: 3},
// 					Comparator: intCompare,
// 				},
// 			},
// 		},
// 		Kind:      SequenceKindFirstSuccess,
// 		OnSuccess: s2,
// 	}

// 	result := s1.Execute()
// 	require.Equal(t, StatusSuccess, result)
// 	require.Equal(t, StatusSuccess, s1.LastStatus)
// 	require.Equal(t, StatusSuccess, s2.LastStatus)

// }

type Character struct {
	Name   string
	Health int
	Energy int
	Gold   int
}

func healthConditionFactory(character *Character) *Condition {
	return &Condition{
		LHS: LHS{Value: character},
		Comparator: func(lhs *LHS, rhs *RHS) Status {
			character := lhs.Value.(*Character)
			if character.Health <= 50 {
				return StatusSuccess
			}

			return StatusFailure
		},
	}
}

func energyConditionFactory(character *Character) *Condition {
	return &Condition{
		LHS: LHS{Value: character},
		Comparator: func(lhs *LHS, rhs *RHS) Status {
			character := lhs.Value.(*Character)
			if character.Energy <= 50 {
				return StatusSuccess
			}

			return StatusFailure
		},
	}
}

func goldConditionFactory(character *Character) *Condition {
	return &Condition{
		LHS: LHS{Value: character},
		Comparator: func(lhs *LHS, rhs *RHS) Status {
			character := lhs.Value.(*Character)
			if character.Gold > 10 {
				return StatusSuccess
			}

			return StatusFailure
		},
	}
}

type SucceedNode struct {
	name      string
	condition *Condition
	onSuccess *Sequence
}

func (n *SucceedNode) Name() string          { return n.name }
func (n *SucceedNode) Condition() *Condition { return n.condition }
func (n *SucceedNode) OnSuccess() *Sequence  { return n.onSuccess }
func (n *SucceedNode) Execute() Status       { return StatusSuccess }

type HealNode struct {
	character *Character
}

func (n *HealNode) Name() string { return "heal" }
func (n *HealNode) Condition() *Condition {
	return &Condition{Comparator: func(lhs *LHS, rhs *RHS) Status {
		return StatusSuccess
	}}
}
func (n *HealNode) OnSuccess() *Sequence { return nil }
func (n *HealNode) Execute() Status {
	n.character.Health = 100
	slog.Info("healed")
	return StatusSuccess
}

type RestNode struct {
	character *Character
}

func (n *RestNode) Name() string { return "rest" }
func (n *RestNode) Condition() *Condition {
	return &Condition{Comparator: func(lhs *LHS, rhs *RHS) Status {
		return StatusSuccess
	}}
}
func (n *RestNode) OnSuccess() *Sequence { return nil }
func (n *RestNode) Execute() Status {
	n.character.Energy = 100
	slog.Info("rested")
	return StatusSuccess
}

type TavernNode struct {
	character *Character
}

func (n *TavernNode) Name() string { return "tavern" }
func (n *TavernNode) Condition() *Condition {
	return &Condition{Comparator: func(lhs *LHS, rhs *RHS) Status {
		return StatusSuccess
	}}
}
func (n *TavernNode) OnSuccess() *Sequence { return nil }
func (n *TavernNode) Execute() Status {
	newGold := n.character.Gold - 10
	if newGold < 0 {
		newGold = 0
	}
	n.character.Gold = newGold
	slog.Info("tavern visited")
	return StatusSuccess
}

func executeScenario(ch *Character) *Character {
	healSequence := &Sequence{
		Name: "HealSequence",
		Nodes: []N{
			&HealNode{character: ch},
		},
		Kind: SequenceKindAllSuccess,
	}

	restSequence := &Sequence{
		Name: "RestSequence",
		Nodes: []N{
			&RestNode{character: ch},
		},
		Kind: SequenceKindAllSuccess,
	}

	tavernSequence := &Sequence{
		Name: "TavernSequence",
		Nodes: []N{
			&TavernNode{character: ch},
		},
		Kind: SequenceKindAllSuccess,
	}

	catchCondition := &Condition{
		Comparator: func(lhs *LHS, rhs *RHS) Status {
			return StatusSuccess
		},
	}

	healthNode := &SucceedNode{name: "HealthCheck", condition: healthConditionFactory(ch), onSuccess: healSequence}
	energyNode := &SucceedNode{name: "EnergyCheck", condition: energyConditionFactory(ch), onSuccess: restSequence}
	goldNode := &SucceedNode{name: "GoldCheck", condition: goldConditionFactory(ch), onSuccess: tavernSequence}
	catchNode := &SucceedNode{name: "Catch", condition: catchCondition}

	sequence := &Sequence{
		Name: "MainSequence",
		Nodes: []N{
			healthNode,
			energyNode,
			goldNode,
			catchNode,
		},
		Kind: SequenceKindFirstSuccess,
	}

	sequence.Execute()
	return ch
}

func TestBasicScenario(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn})))

	tests := []struct {
		label          string
		ch             *Character
		expectedHealth int
		expectedEnergy int
		expectedGold   int
	}{
		{
			label: "healthy and rested, go to tavern",
			ch: &Character{
				Name:   "Hero",
				Health: 100,
				Energy: 100,
				Gold:   20,
			},
			expectedHealth: 100,
			expectedEnergy: 100,
			expectedGold:   10,
		},
		{
			label: "healthy and no energy, rest",
			ch: &Character{
				Name:   "Hero",
				Health: 100,
				Energy: 30,
				Gold:   20,
			},
			expectedHealth: 100,
			expectedEnergy: 100,
			expectedGold:   20,
		},
		{
			label: "unhealthy, go heal",
			ch: &Character{
				Name:   "Hero",
				Health: 30,
				Energy: 30,
				Gold:   20,
			},
			expectedHealth: 100,
			expectedEnergy: 30,
			expectedGold:   20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			ch := executeScenario(tt.ch)
			require.Equal(t, tt.expectedHealth, ch.Health)
			require.Equal(t, tt.expectedEnergy, ch.Energy)
			require.Equal(t, tt.expectedGold, ch.Gold)
		})
	}

}
