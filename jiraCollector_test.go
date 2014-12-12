package GoJiraTools

import (
	"fmt"
	"testing"
)

var PROTOCOL = "http"
var HOST = "localhost"
var PORT = 80
var USER = "testuser"
var PASS = "testpass"
var APIVERSION = 2
var SSL = false

var BASE_URI = fmt.Sprintf("%s://%s:%d/%d", PROTOCOL, HOST, PORT,APIVERSION)

func TestNewJiraConnection(t *testing.T) {
	jc := 	NewJiraConnection(PROTOCOL, 
			HOST, PORT, 
			USER,PASS,
			2,true);
	
	if(jc.protocol != PROTOCOL) {
		t.Errorf("protocol should be %S, not %s", PROTOCOL, jc.protocol)
	}

	if(jc.host != HOST) {
		t.Errorf("host should be %s, not %s", HOST, jc.host)
	}

	if(jc.baseURI != BASE_URI) {
		t.Errorf("base URI should be %s ---- not %s",BASE_URI,jc.baseURI)
	}

}

func TestNewJiraConnectionDefaults(t *testing.T){
	jc := NewJiraConnection("", HOST, 0, USER, PASS, APIVERSION, SSL)

	if(jc.protocol != PROTOCOL) {
		t.Errorf("protocol should be %s, not %s", PROTOCOL, jc.protocol)
	}
}