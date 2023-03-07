package simulator

import (
	"sync"
	"time"

	"github.com/baribari2/pulp-calculator/common/types"
)

// The main data structure for a simulation, this is an interface so that a debate can be simulated for various
// user types and distributions.
type SimulationSet interface {
	// The type of simulation (enneagram, age, etc.)
	GetSimulationType() types.SimulationType

	// The amount of users to simulate in this simulation
	GetSimulationSize() int

	// The users that are part of this simulation
	GetUsers() []*types.User

	// The distribution of users in this simulation
	// If the simulation set is enneagram the distribution will be the distribution of enneagram types,
	// if the simulation set is age the distribution will be the distribution of ages, etc.
	//
	// The values in the array * the length of the array must equal about 1
	// For example, []float{0.2, 0.2, 0.2, 0.2, 0.2}
	GetDistribution() []float64

	// The length the simulation is to run for
	GetDuration() time.Duration

	// The topic of the simulation
	GetTopic() string

	// The category of the simulation
	GetCategory() string

	RunSimulation(*sync.WaitGroup, *types.Config, chan *Debate, chan error)
}
