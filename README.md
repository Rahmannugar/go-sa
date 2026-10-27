# DSA in Go

Go implementations of data structures and algorithms exercises from a [Boot.dev YouTube course](https://www.youtube.com/watch?v=DMeD8trbj6A). The original exercises are taught in Python.

## Exercises

| Directory | What it covers |
| --- | --- |
| [`find-minimum`](find-minimum/) | Find the smallest integer in a slice. |
| [`total-sum`](total-sum/) | Add up the integers in a slice. |
| [`get-exponent`](get-exponent/) | Calculate values using averages and exponents. |
| [`logarithms`](logarithms/) | Calculate a score using a base-2 logarithm. |
| [`big-O-notation`](big-O-notation/) | Explore linear, quadratic, constant, and logarithmic time complexity. |
| [`sorting`](sorting/) | Sort influencers by a calculated vanity score. |
| [`bubble-sort`](bubble-sort/) | Sort integers by repeatedly swapping adjacent values. |
| [`insertion-sort`](insertion-sort/) | Sort integers by inserting each value into its position. |
| [`merge-sort`](merge-sort/) | Sort integers by splitting and merging slices. |
| [`quick-sort`](quick-sort/) | Sort integers in place by partitioning around a pivot. |

Each directory has a standalone `main.go` with example inputs. To run an exercise from the repository root, use:

```sh
go run ./find-minimum
go run ./total-sum
go run ./get-exponent
go run ./logarithms
go run ./big-O-notation
go run ./sorting
go run ./bubble-sort
go run ./insertion-sort
go run ./merge-sort
go run ./quick-sort
```

