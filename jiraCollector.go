package GoJiraTools

import (
	http "net/http"
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

	jc := JiraConnection{}

	if len(protocol) > 0 {
		jc.protocol = protocol
	} else {
		jc.protocol = "http"
	}

	jc.host = host
	if port != 0 {
		jc.port = port
	} else {
		jc.port = 80
	}

	jc.username = username
	jc.password = password

	jc.apiVersion = apiVersion

	jc.baseURI = fmt.Sprintf("%s://%s:%d/%d", jc.protocol, jc.host, jc.port, jc.apiVersion);

	return &jc

}

func makeRequest(jc *JiraConnection) (*http.Request, error) {
	r, err := http.NewRequest("GET", jc.baseURI, nil);

	if(err != nil) {
		return nil, err
	}

	return r, nil
}

func fetchIssue(jc *JiraConnection, issueId string) {

}

