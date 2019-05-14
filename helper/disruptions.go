package helper

// Disruption represents a Category of problem and
// any general next steps
type Disruption struct {
	Category     string
	GeneralSteps string
	Info         string
}

// LocalDNS represents problems with DNS servers hosted on the customer's intranet
var localDNS = Disruption{Category: "Intranet DNS", GeneralSteps: "Find IP Address of local DNS servers that were used by looking for phrase 'PROXY found DNS server'.", Info: "The Local DNS server Could not be found or could not be used to resolve domain names."}

var kgpConnection = Disruption{Category: "KGP Maki Connection", GeneralSteps: "Ensure the tunnel can stay up for at least 24 hours via Looker.", Info: "This should only happen briefly, ideally never, except at the end of a tunnel's lifecycle. If it happens repeatedly there is a proble connecting the SC Client to the Maki Tunnel VM."}
