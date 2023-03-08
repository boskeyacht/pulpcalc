package sets

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"time"

	"github.com/baribari2/pulp-calculator/simulator"
	"github.com/baribari2/pulp-calculator/simulator/sets/enneagram"
	"gopkg.in/yaml.v3"
)

// Parses the yaml config file and returns a SimulationSet
// Go polymorphism is a PITA -_-
func NewSimulationSetsFromFile(path string) ([]simulator.SimulationSet, error) {
	s := make(map[interface{}]interface{}, 0)
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(f, &s)
	if err != nil {
		return nil, err
	}

	if ok := validateYAML(s); !ok {
		return nil, fmt.Errorf("invalid yaml")
	}

	sets := make([]simulator.SimulationSet, len(s["sets"].([]interface{})))

	for i, set := range s["sets"].([]interface{}) {
		switch set.(map[string]interface{})["simulation_type"] {
		case "enneagram":
			sets[i] = NewEnneagramSetFromYAML(set.(map[string]interface{}))
			break
		case "age":
			// sets[i] = NewAgeSetFromYAML(set.(map[string]interface{}))
			break
		case "registration":
			// sets[i] = NewRegistrationSetFromYAML(set.(map[string]interface{}))
			break
		default:
			return nil, fmt.Errorf("invalid simulation type")
		}
	}

	return sets, nil
}

func NewEnneagramSetFromYAML(yml map[string]interface{}) simulator.SimulationSet {
	d := time.Duration(time.Duration(yml["duration"].(int)) * time.Second)
	fl := make([]float64, len(yml["distribution"].([]interface{})))

	for i, f := range yml["distribution"].([]interface{}) {
		fl[i] = f.(float64)
	}

	es := enneagram.NewEnneagramSet(
		yml["simulation_size"].(int),
		nil,
		fl,
		d,
		yml["topic"].(string),
		yml["category"].(string),
		yml["depth"].(int),
	)

	return es
}

// In order to dynamically create a SimulationSet from a yaml file, we pass a map[string]interface{},
// which means we need to validate the data in the yaml file. (We dynamically create the SimulationSet
// because we want to be able to create a SimulationSet for any type of simulation for concurrent execution)
func validateYAML(yml map[interface{}]interface{}) bool {
	if array := reflect.ValueOf(yml["sets"]).Kind(); array != reflect.Slice {
		fmt.Println("sets", array)

		return false
	}

	for _, set := range yml["sets"].([]interface{}) {
		if str := reflect.ValueOf(set.(map[string]interface{})["simulation_type"]).Kind(); str != reflect.String {
			fmt.Println("simulation_type", str)

			return false
		}

		if str := reflect.ValueOf(set.(map[string]interface{})["simulation_size"]).Kind(); str != reflect.Int {
			fmt.Println("simulation_size", str)

			return false
		}

		if str := reflect.ValueOf(set.(map[string]interface{})["duration"]).Kind(); str != reflect.Int {
			fmt.Println("duration", str)

			return false
		}

		if str := reflect.ValueOf(set.(map[string]interface{})["topic"]).Kind(); str != reflect.String {
			fmt.Println("topic", str)

			return false
		}

		if str := reflect.ValueOf(set.(map[string]interface{})["category"]).Kind(); str != reflect.String {
			fmt.Println("category", str)

			return false
		}

		if str := reflect.ValueOf(set.(map[string]interface{})["distribution"]).Kind(); str != reflect.Slice {
			return false
		}

		for _, dist := range set.(map[string]interface{})["distribution"].([]interface{}) {
			if str := reflect.ValueOf(dist).Kind(); str != reflect.Float64 {
				fmt.Println("distribution", str)

				return false
			}
		}
	}

	return true
}
