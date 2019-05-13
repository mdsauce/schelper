SCHelper Spec
=============
schelper will read in a sauce connect log of any kind (-v or -vv or no v/vv).  It will compare each line in the log against a known list of `Known Problems` and their respective representations in logs.  Partial matches will be labeled as such just to bring them to the operator's attention.

After a `Problem` is identified the log line locations and timestamp will be recorded as part of that `Problem`.  A category for the `Disruption` will be assigned.  

Then at the end of the program a meta assessment will fit the `Disruption` summary into the stages of the `Lifecycle` of Sauce Connect.

### Useful Objects
A `Problem` is the `Location` & `Disruption`.

Problem:
Location [Line Entries, Timestamp from Log]
Disruption [Category, Next Steps]

A `Disruption` is a `Category` of pre-defined problem with the associated `Next Steps`. 

A `Known Problem` is from the hardcoded list of log entries from past tickets.  Invisible to the operator til runtime.

A `Lifecycle` is a list of line locations for each point in the Sauce Connect Lifecycle.
- Startup
- Maki Allocation
- Running
- Keep Alive
- Shutdown