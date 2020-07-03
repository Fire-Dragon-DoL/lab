package actions

import (
	"io"
)

// RunExample is a JSON object with Action set to `run`
func RunExample() io.Reader {
	return ActionExample(`{"Time":"2020-07-01T11:22:52.295096127-07:00","Action":"run","Package":"command-line-arguments","Test":"TestInteractive"}`)
}

// OutputPassExample is a JSON object with Action set to `pass`
func OutputPassExample() io.Reader {
	return ActionExample(`{"Time":"2020-07-01T11:22:52.295405039-07:00","Action":"output","Package":"command-line-arguments","Test":"TestInteractive/Nested_3/Nested_3_One","Output":"        --- PASS: TestInteractive/Nested_3/Nested_3_One (0.00s)\n"}`)
}
