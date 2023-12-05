package team5Agent

import (
	"SOMAS2023/internal/common/utils"
	"sort"

	"github.com/google/uuid"
)

func (t5 *team5Agent) VoteForKickout() map[uuid.UUID]int {
	agentsOnBike := t5.GetFellowBikers()
	numberOfAgents := float64(len(agentsOnBike))

	internalRanking := make(map[uuid.UUID]float64)
	ranking := make(map[uuid.UUID]int)

	a := 1.0
	b := 1.0
	c := 0.6
	forceMax := utils.BikerMaxForce

	for _, agent := range agentsOnBike {

		key := agent.GetID()
		reputation := t5.QueryReputation(key)
		pedallingForce := t5.GetForces().Pedal

		utility := (a * pedallingForce) + (b * reputation) + (c * numberOfAgents)
		utilityNorm := utility / ((a * forceMax) + b + c)

		internalRanking[key] = utilityNorm
	}

	// Extras for sorting
	type kv struct {
		Key   uuid.UUID
		Value float64
	}

	var ss []kv

	for k, v := range internalRanking {
		ss = append(ss, kv{k, v})
	}

	// Sort the slice by values
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	for i, pair := range ss {
		ranking[pair.Key] = i + 1
	}

	return ranking
}

func (t5 *team5Agent) DecideJoining(pendingAgents []uuid.UUID) map[uuid.UUID]bool {

	agentMap := t5.GetGameState().GetAgents()
	result := make(map[uuid.UUID]bool)
	pendingAgentUtility := make(map[uuid.UUID]float64)
	threshold := 0.5
	maxBikers := utils.BikersOnBike
	currentBikers := maxBikers - len(t5.GetFellowBikers())

	a := 1.0
	b := 1.0
	c := 1.0
	energyMax := 1.0
	targetColor := t5.GetColour()

	for _, agentID := range pendingAgents {
		agentState := agentMap[agentID]

		key := agentState.GetID()
		reputation := agentState.QueryReputation(key)
		energyLevel := agentState.GetEnergyLevel()
		pendingAgentColor := agentState.GetColour()

		isColorSame := 0.0

		if targetColor == pendingAgentColor {
			isColorSame = 1.0
		}

		// color has to be a 0/1 and replaced with
		utility := (a * energyLevel) + (b * reputation) + (c * isColorSame)
		utilityNorm := utility / ((a * energyMax) + b + c)

		pendingAgentUtility[agentID] = utilityNorm

	}

	type kv struct {
		Key   uuid.UUID
		Value float64
	}

	var ss []kv
	for k, v := range pendingAgentUtility {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value // Sorting in descending order
	})

	for i, pair := range ss {
		result[pair.Key] = i < currentBikers && pair.Value >= threshold
	}

	return result
}
