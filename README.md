# SCHelper (Sauce Connect Helper)
Feed in a Sauce Connect Proxy log, get back general info and next steps to troubleshoot any problems. -v, -vv, or zero verbosity, doesn't matter.  More data is better in general though.  

[![CircleCI](https://circleci.com/gh/mdsauce/schelper/tree/master.svg?style=svg)](https://circleci.com/gh/mdsauce/schelper/tree/master)

## Installation
1. `brew tap mdsauce/schelper https://github.com/mdsauce/schelper-brew`
2. `brew install schelper`

Using Homebrew run the above two commands.  Should take about thirty seconds if you have Homebrew installed.

If you don't have Homebrew or a Mac then `go get mdsauce/schelper` should work.  Cloning/forking the repo then doing a `go build` or `go install` will also work as long as you have golang installed.

## Usage Guide
`$ schelper read ~/absolute/or/relative/path/to/sc.log`

`$ schelper -h` or `--help` to get more info about your options and default settings.
`$ schelper sclog -h` or `--help` to get more info about specifying the SC log.

Add a `-v` or `--verbose` flag to stop suppression of redundant output.  May be messy.

## Contributing
If you wish to send pull requests adding more "known problems" that I missed read below.

### Adding Known Problems
A `Known Problem` is currently

```
type KnownProblem struct {
	Name       string
	Logs       []byte
	NextSteps  string
}
```

Go to `knownProblems.go` and add an entry to the `AllProblems()` function.

```
var <name of problem> = KnownProblem{Name: "name-of-problem-w-hyphens", Logs: []byte("add as much detail as possible here.  One big string with each word separated by spaces EXACTLY as it appears in the log.  NO LINEBREAKS!!! NO NEWLINES!!! Just one long continuous string"), NextSteps: `Anything with the '`' is a string literal.  Will literally be printed out as it appears here.  Feel free to play w/ formatting.`}
// Append your problem to []AllProblems (slice, like an array)
AllProblems = append(AllProblems, <name of problem>)
```
Add a Test to `comparison_test.go` if you want to be sure it works as intended.  Same format for every test.
