package command

import "testing"

func TestGetDiceData(t *testing.T) {
	// Assemble
	cases := []struct {
		Name              string
		CommandInput      string
		IsError           bool
		ExpectedDieCount  int
		ExpectedDiceSides int
	}{
		{
			Name:              "valid",
			CommandInput:      "4d8",
			IsError:           false,
			ExpectedDieCount:  4,
			ExpectedDiceSides: 8,
		},
		{
			Name:              "valid - large count",
			CommandInput:      "10000d8",
			IsError:           false,
			ExpectedDieCount:  10000,
			ExpectedDiceSides: 8,
		},
		{
			Name:              "valid - large sides",
			CommandInput:      "4d10000",
			IsError:           false,
			ExpectedDieCount:  4,
			ExpectedDiceSides: 10000,
		},
		{
			Name:              "invalid - no sides",
			CommandInput:      "4d",
			IsError:           true,
			ExpectedDieCount:  -1,
			ExpectedDiceSides: -1,
		},
		{
			Name:              "invalid - no count",
			CommandInput:      "d8",
			IsError:           true,
			ExpectedDieCount:  -1,
			ExpectedDiceSides: -1,
		},
		{
			Name:              "invalid - no delimiter",
			CommandInput:      "4 8",
			IsError:           true,
			ExpectedDieCount:  -1,
			ExpectedDiceSides: -1,
		},
	}

	for _, c := range cases {
		// Act
		dieCount, diceSides, err := getDiceData(c.CommandInput)

		if c.IsError {
			if err == nil {
				t.Errorf("%s - expected an error, but got none", c.Name)
				continue
			}
		} else {
			if err != nil {
				t.Errorf("%s - expected no error, but got %v", c.Name, err)
				continue
			}
			if c.ExpectedDieCount != dieCount {
				t.Errorf("%s - expected dieCount of %d, but got %d", c.Name, c.ExpectedDieCount, dieCount)
				continue
			}
			if c.ExpectedDiceSides != diceSides {
				t.Errorf("%s - expected diceSides of %d, but got %d", c.Name, c.ExpectedDiceSides, diceSides)
				continue
			}
		}
	}
}
