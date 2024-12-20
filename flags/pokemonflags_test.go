package flags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupPokemonFlagSet(t *testing.T) {
	// Call the function to get the flag set and types flag
	pokeFlags, typesFlag := SetupPokemonFlagSet()

	// Assertions
	assert.NotNil(t, pokeFlags, "Flag set should not be nil")
	assert.Equal(t, "pokeFlags", pokeFlags.Name(), "Flag set name should be 'pokeFlags'")

	// Check types flag
	assert.NotNil(t, typesFlag, "Types flag should not be nil")
	assert.Equal(t, bool(false), *typesFlag, "Types flag name should be 'types'")
}
