package GoJiraTools

import (
	http "net/http"
	"bytes"
	"fmt"
)

// container of JiraConnect information
type JiraConnection struct {
	protocol   string
	host       string
	port       int
	username   string
	password   string
	apiVersion int
	strictSSL  bool
	baseURI    string
}

// builds a JiraConnection object with defaults for protocol and port.  
// also builds the baseURI.
func NewJiraConnection(protocol string, host string, port int, username string, 
		password string, apiVersion int, stringSSL bool) *JiraConnection {

	jc := JiraConnection{protocol, host, port, username, password, apiVersion, stringSSL,""}

	// default protocol to httep
	if len(protocol) > 0 {
		jc.protocol = protocol
	} else {
		jc.protocol = "http"
	}

	// default port to 80s
	if port != 0 {
		jc.port = port
	} else {
		jc.port = 80
	}

	// create the baseURI
	jc.baseURI = fmt.Sprintf("%s://%s:%d/%d", jc.protocol, jc.host, jc.port, jc.apiVersion);

	return &jc

}

func makeRequest(jc *JiraConnection, urlsuffix string, payload []byte) (*http.Request, error) {
	
	payloadBuffer := bytes.NewBuffer(payload)

	r, err := http.NewRequest("GET", jc.baseURI + urlsuffix, payloadBuffer);

	if(err != nil) {
		return nil, err
	}

	r.SetBasicAuth(jc.username, jc.password)

	return r, nil
}

func findIssue(jc *JiraConnection, issueId string) {
	urlPart := "/issue/" + issueId

	r, err := makeRequest(jc, urlPart, nil)

	if(err != nil || r == nil){
		return;
	}



}

