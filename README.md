# cwe

[![Go Reference](https://pkg.go.dev/badge/github.com/Preetam/cwe.svg)](https://pkg.go.dev/github.com/Preetam/cwe)

Go package of CWE IDs and metadata. The list is generated from a CSV from the [Comprehensive CWE Dictionary](https://cwe.mitre.org/data/definitions/2000.html).

### Example

Here's CWE-918:

```go
"CWE-918": {
	Name:                  "Server-Side Request Forgery (SSRF)",
	WeaknessAbstraction:   "Base",
	Status:                "Incomplete",
	Description:           "The web server receives a URL or similar request from an upstream component and retrievesthe contents of this URL, but it does not sufficiently ensure that the request is being sent to the expecteddestination.",
	ExtendedDescription:   "By providing URLs to unexpected hosts or ports, attackers can make it appear that theserver is sending the request, possibly bypassing access controls such as firewalls that prevent the attackers fromaccessing the URLs directly. The server can be used as a proxy to conduct port scanning of hosts in internalnetworks, use other URLs such as that can access documents on the system (using file://), or use other protocolssuch as gopher:// or tftp://, which may provide greater control over the contents of requests.",
	RelatedWeaknesses:     "::NATURE:ChildOf:CWE ID:441:VIEW ID:1000:ORDINAL:Primary::NATURE:ChildOf:CWE ID:610:VIEWID:1003:ORDINAL:Primary::",
	WeaknessOrdinalities:  "",
	ApplicablePlatforms:   "::LANGUAGE CLASS:Language-Independent:LANGUAGE PREVALENCE:Undetermined::TECHNOLOGY NAME:WebServer:TECHNOLOGY PREVALENCE:Undetermined::",
	BackgroundDetails:     "",
	AlternateTerms:        "::TERM:XSPA:DESCRIPTION:Cross Site Port Attack::",
	ModesOfIntroduction:   "::PHASE:Architecture and Design::PHASE:Implementation::",
	ExploitationFactors:   "",
	LikelihoodOfExploit:   "",
	CommonConsequences:    "::SCOPE:Confidentiality:IMPACT:Read Application Data::SCOPE:Integrity:IMPACT:ExecuteUnauthorized Code or Commands::",
	DetectionMethods:      "",
	PotentialMitigations:  "",
	ObservedExamples:      "::REFERENCE:CVE-2002-1484:DESCRIPTION:Web server allows attackers to request a URL fromanother server, including other ports, which allows proxied scanning.:LINK:https://cve.mitre.org/cgi-bin/cvenamecgi?name=CVE-2002-1484::REFERENCE:CVE-2004-2061:DESCRIPTION:CGI script accepts and retrieves incoming URLs:LINK:https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2004-2061::REFERENCE:CVE-2010-1637:DESCRIPTION:Web-basedmail program allows internal network scanning using a modified POP3 port number.:LINK:https://cve.mitre.org/cgi-bincvename.cgi?name=CVE-2010-1637::REFERENCE:CVE-2009-0037:DESCRIPTION:URL-downloading library automatically followsredirects to file:// and scp:// URLs:LINK:https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2009-0037::",
	FunctionalAreas:       "",
	AffectedResources:     "",
	TaxonomyMappings:      "",
	RelatedAttackPatterns: "",
	Notes:                 "::TYPE:Relationship:NOTE:CWE-918 (SSRF) and CWE-611 (XXE) are closely related, because theyboth involve web-related technologies and can launch outbound requests to unexpected destinations. However, XXE canbe performed client-side, or in other contexts in which the software is not acting directly as a server, so theServer portion of the SSRF acronym does not necessarily apply.::",
},
```
