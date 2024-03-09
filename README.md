# Mathutil Package:

The mathutil package provides functionality for computing the average of a slice of numerical values, with a 5% trim to remove outliers.

## Package Use:

1. Import the mathutil package into your Go program:

`import ("github.com/username/mathutil")`

2. Call the Average function with a slice of numerical values:

`numbers := []float64{1.2, 3.4, 5.6, 7.8, 9.0}`
`avg := mathutil.Average(numbers)`
`fmt.Println("Average:", avg)`



