package test

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	url := "http://www.goinggo.net/index.xml"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tTest 0:\tWhen checking %q for status code %d", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to make the Get call : %v", ballotX, err)
			}
			t.Log("\t", checkMark, "Should be able to make the Get call.")

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t%s\tShould receive a %d status code.", checkMark, statusCode)
			} else {
				t.Errorf("\t%s\tShould receive a %d status code : %d", ballotX, statusCode, resp.StatusCode)
			}
		}
	}
}
