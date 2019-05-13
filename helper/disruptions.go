package helper

// Disruption represents a Category of problem and
// any general next steps
type Disruption struct {
	Category     string
	GeneralSteps string
	Info         string
}

// LocalDNS represents problems with DNS servers hosted on the customer's intranet
var LocalDNS = Disruption{Category: "Intranet DNS", GeneralSteps: "Find IP Address of local DNS servers that were used by looking for phrase 'PROXY found DNS server'.", Info: "The Local DNS server Could not be found or could not be used to resolve domain names."}
