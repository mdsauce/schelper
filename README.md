# SCHelper
Feed in a Sauce Connect Proxy log, get back general info and next steps to troubleshoot any problems. -v, -vv, or zero verbosity, doesn't matter.  More data is better in general though.  

### Starting
`$ schelper sclog ~/absolute/or/relative/path/to/sc.log`


### Adding Disruptions
A `Disruption` is currently 

```
type Disruption struct {
	Category     string
	GeneralSteps string
	Info         string
}
```

Go to `disruptions.go` and add a new local variable to the package with this format.  Disruptions should be general like what an organ is to a body.

```
var <local name of disruption> = Disruption{Category: "Overarching category.  Avoid conflicts as much as you can.", GeneralSteps: "In general, what should you do when you see this Category of problem?", Info: "The WHY? of this category.  What's going on?  Any background info?  Anything important or useful?"}

```

### Adding Known Problems
A `Known Problem` is currently

```
type KnownProblem struct {
	Name       string
	Disruption Disruption
	Logs       []byte
	NextSteps  string
}
```

Go to `knownProblems.go` and add an entry to the `AllProblems()` function.

```
var <name of problem> = KnownProblem{Name: "name-of-problem-w-hyphens", Disruption: localDisruptionVariable, Logs: []byte("add as much detail as possible here.  One big string with each word separated by spaces EXACTLY as it appears in the log.  NO LINEBREAKS!!! NO NEWLINES!!! Just one long continuous string"), NextSteps: `Anything with the '`' is a string literal.  Will literally be printed out as it appears here.  Feel free to play w/ formatting.`}
// Append your problem to []AllProblems (slice, like an array)
AllProblems = append(AllProblems, <name of problem>)
```
Add a Test to `comparison_test.go` if you want to be sure it works as intended.  Same format for every test.
