- battle tested → reliable, effective, and proven in real-world production over time
- binge watch → watching something in one sitting, often without practicing—e.g., coding tutorials
- data aggregation → collecting and combining data from many sources into one summary
- data orchestration → coordinating and managing data tasks, pipelines, and their order
- data migration → migration means switching, moving or switching data from an old system or schema to a newer one
- data integration → connecting data from different systems so they work together
- integrity → being whole, correct, and not corrupted
- data integrity → data is correct, complete, consistent, and can be trusted)
- succinct → clear and concise
- pass by value → value gets copied
- pass by reference → value does not get copied; reference is passed
- gaurd clauses → early returns, a style of writng codes
```go
func div(dividend, divisor int) (int, error) {
	// gaurd clause
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend/divisor, nil
}
```

🧠 Quick memory tip
Aggregation = gather
Integration = connect
Orchestration = control the flow
