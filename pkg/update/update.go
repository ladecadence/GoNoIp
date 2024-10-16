package update

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ladecadence/GoNoIp/pkg/config"
)

type NoIPCode string

const (
	NoIPCodeGood     NoIPCode = "Good"
	NoIPCodeNoChg    NoIPCode = "No Change"
	NoIPCodeNoHost   NoIPCode = "No Host"
	NoIPCodeBadAuth  NoIPCode = "Bad Auth"
	NoIPCodeBadAgent NoIPCode = "Bad Agent"
	NoIPCodeDonator  NoIPCode = "No Donator"
	NoIPCodeAbuse    NoIPCode = "Abuse"
	NoIPCodeC911     NoIPCode = "C911"
	NoIPCodeUnknown  NoIPCode = "Unknown"

	UserAgentString string = "Ladecadence.net NoIP/0.2 zako@ladecadence.net"
)

func Update(host config.Host) NoIPCode {

	// create update url
	requestURL := host.UpdateUrl + fmt.Sprintf("?hostname=%s&offline=%s", host.Hostname, host.Offline)
	if host.IP != "" {
		requestURL += fmt.Sprintf("&myip=%s", host.IP)
	}

	// create request
	request, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return NoIPCodeUnknown
	}

	// add auth and headers and send request
	request.SetBasicAuth(host.Username, host.Password)
	request.Header.Set("User-Agent", UserAgentString)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return NoIPCodeUnknown
	}

	// check response
	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return NoIPCodeUnknown
	}
	code := strings.Split(strings.TrimSpace(string(resBody)), " ")

	switch code[0] {
	case "good":
		{
			return NoIPCodeGood
		}
	case "nochg":
		{
			return NoIPCodeNoChg
		}
	case "nohost":
		{
			return NoIPCodeNoHost
		}
	case "badauth":
		{
			return NoIPCodeBadAuth
		}
	case "badagent":
		{
			return NoIPCodeBadAgent
		}
	case "!donator":
		{
			return NoIPCodeDonator
		}
	case "abuse":
		{
			return NoIPCodeAbuse
		}
	case "911":
		{
			return NoIPCodeC911
		}
	default:
		{
			return NoIPCodeUnknown
		}
	}
}
