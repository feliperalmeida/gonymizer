package gonymizer

import (
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/require"
)

var proc = []ProcessorDefinition{
	{
		Name:     "",
		Max:      0,
		Min:      0,
		Variance: 0,
		Comment:  "",
	},
}

var cMap = ColumnMapper{
	TableSchema:     "",
	TableName:       "",
	ColumnName:      "",
	DataType:        "",
	ParentSchema:    "",
	ParentTable:     "",
	ParentColumn:    "",
	OrdinalPosition: 4,
	IsNullable:      false,
	Processors:      proc,
	Comment:         "",
}

func TestProcessorFunc(t *testing.T) {
}

func TestProcessorAlphaNumericScrambler(t *testing.T) {
	var alphaTest ColumnMapper

	output, err := ProcessorAlphaNumericScrambler(&cMap, "AsDfG10*&")
	require.Nil(t, err)
	require.NotEqual(t, output, "AsDfG0*&")

	alphaTest.ParentSchema = "test_schema"
	alphaTest.ParentTable = "test_table"
	alphaTest.ParentColumn = "test_column"

	outputA, err := ProcessorAlphaNumericScrambler(&alphaTest, "My name is Mud - Pr1mUs")
	require.Nil(t, err)
	outputB, err := ProcessorAlphaNumericScrambler(&alphaTest, "My name is Mud - 111111")
	require.Nil(t, err)
	outputC, err := ProcessorAlphaNumericScrambler(&alphaTest, "My name is Mud - Pr1mUs")
	require.Nil(t, err)

	// outputA != outputB
	require.NotEqual(t, outputA, outputB)
	// outputA === outputC
	require.Equal(t, outputA, outputC)
}

func TestProcessorAddress(t *testing.T) {
	output, err := ProcessorAddress(&cMap, "1234 Testing Lane")
	require.Nil(t, err)
	require.NotEqual(t, output, "1234 Testing Lane")

	output, err = ProcessorAddress(&cMap, "")
	require.Nil(t, err)
	require.NotEqual(t, output, "")
}

func TestProcessorCity(t *testing.T) {
	output, err := ProcessorCity(&cMap, "Rick and Morty Ville")
	require.Nil(t, err)
	require.NotEqual(t, output, "Rick and Morty Ville")

	output, err = ProcessorCity(&cMap, "")
	require.Nil(t, err)
	require.NotEqual(t, output, "")
}

func TestProcessorEmailAddress(t *testing.T) {
	output, err := ProcessorEmailAddress(&cMap, "rick@morty.example.com")
	require.Nil(t, err)
	require.NotEqual(t, output, "rick@morty.example.com")

	output, err = ProcessorEmailAddress(&cMap, "")
	require.Nil(t, err)
	require.NotEqual(t, output, "")
}

func TestProcessorFirstName(t *testing.T) {
	output, err := ProcessorFirstName(&cMap, "RickMortyRick")
	require.Nil(t, err)
	require.NotEqual(t, output, "RickMortyRick")

	output, err = ProcessorFirstName(&cMap, "")
	require.Nil(t, err)
	require.NotEqual(t, output, "")
}

func TestProcessorFullName(t *testing.T) {
	output, err := ProcessorFullName(&cMap, "Morty & Rick")
	require.Nil(t, err)
	require.NotEqual(t, output, "Morty & Rick")

	output, err = ProcessorFullName(&cMap, "")
	require.Nil(t, err)
	require.NotEqual(t, output, "")
}

func TestProcessorIdentity(t *testing.T) {
	output, err := ProcessorIdentity(&cMap, "Hi Rick!")
	require.Nil(t, err)
	require.Equal(t, output, "Hi Rick!")

	output, err = ProcessorIdentity(&cMap, "")
	require.Nil(t, err)
	require.Equal(t, output, "")
}

func TestProcessorLastName(t *testing.T) {
	output, err := ProcessorLastName(&cMap, "Bye Rick!")
	require.Nil(t, err)
	require.NotEqual(t, output, "Bye Rick!")

	output, err = ProcessorLastName(&cMap, "")
	require.Nil(t, err)
	require.NotEqual(t, output, "")
}

func TestProcessorPhoneNumber(t *testing.T) {
	output, err := ProcessorPhoneNumber(&cMap, "+18885551212")
	require.Nil(t, err)
	require.NotEqual(t, output, "+18885551212")

	output, err = ProcessorPhoneNumber(&cMap, "")
	require.Nil(t, err)
	require.NotEqual(t, output, "")
}

func TestProcessorState(t *testing.T) {
	output, err := ProcessorState(&cMap, "Antarctica")
	require.Nil(t, err)
	require.NotEqual(t, output, "Antarctica")

	output, err = ProcessorState(&cMap, "")
	require.Nil(t, err)
	require.NotEqual(t, output, "")
}

func TestProcessorUserName(t *testing.T) {
	output, err := ProcessorUserName(&cMap, "Ricky and Julian")
	require.Nil(t, err)
	require.NotEqual(t, output, "Ricky and Julian")

	output, err = ProcessorUserName(&cMap, "")
	require.Nil(t, err)
	require.NotEqual(t, output, "")
}

func TestProcessorZip(t *testing.T) {
	output, err := ProcessorZip(&cMap, "00000-00")
	require.Nil(t, err)
	require.NotEqual(t, output, "00000-00")

	output, err = ProcessorZip(&cMap, "")
	require.Nil(t, err)
	require.NotEqual(t, output, "")
}

func TestProcessorRandomDate(t *testing.T) {
	var failBoats = []string{
		"I AM THE FAIL BOAT!",
		"01.01.1970",
		"01/01/1970",
		"1970.01.01",
		"1970/01/01",
		"",
	}

	output, err := ProcessorRandomDate(&cMap, "1970-01-01")
	require.Nil(t, err)
	require.NotEqual(t, output, "1970-01-01")

	for _, tst := range failBoats {
		output, err = ProcessorRandomDate(&cMap, tst)
		require.NotNil(t, err)
		require.NotEqual(t, output, nil)
	}
}

func TestProcessorRandomUUID(t *testing.T) {
	var testUUID uuid.UUID

	testUUID, err := uuid.NewUUID()
	require.Nil(t, err)

	output, err := ProcessorRandomUUID(&cMap, testUUID.String())
	require.Nil(t, err)
	require.NotEqual(t, output, testUUID)

	if val, found := UUIDMap[testUUID]; found {
		if val == testUUID {
			t.Fatalf("UUIDs match\t%s <=> %s", testUUID.String(), val.String())
		}
	} else {
		t.Fatalf("Unable to find UUID '%s' in the UUID map!", output)
	}
	output, err = ProcessorRandomUUID(&cMap, "")
	require.NotNil(t, err)
	require.Equal(t, output, "")
}

func TestProcessorScrubString(t *testing.T) {
	output, err := ProcessorScrubString(&cMap, "Ricky and Julian")
	require.Nil(t, err)
	require.NotEqual(t, output, "Ricky and Julian")

	output, err = ProcessorScrubString(&cMap, "")
	require.Nil(t, err)
	require.Equal(t, output, "")
}

func TestRandomizeUUID(t *testing.T) {
	tempUUID := uuid.New().String()
	output, err := ProcessorRandomUUID(&cMap, tempUUID)
	require.Nil(t, err)
	require.NotEqual(t, output, tempUUID)

	output, err = ProcessorRandomUUID(&cMap, "")
	require.NotNil(t, err)
	require.Equal(t, output, "")
}
