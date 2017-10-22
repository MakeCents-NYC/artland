package mastercard

import (
	"fmt"
	"io"
	"os/exec"
	"time"
)

var OAUTH_AUTHORIZATION string = "Authorization"
var APP_ID string = "4127"
var API_ENDPOINT string = "/labs/proxy/chain/api/v1/network/"

type MasterCard struct {
	keystoreloc  string
	protoloc     string
	userkey      string
	keyalias     string
	keystorepass string
}

func NewMasterCardAPI(keystoreloc, protoloc, userkey, keyalias, keystorepass string) (*MasterCard, error) {
	return &MasterCard{
		keystoreloc:  keystoreloc,
		protoloc:     protoloc,
		userkey:      userkey,
		keyalias:     keyalias,
		keystorepass: keystorepass,
	}, nil
}

func (m *MasterCard) GetApp(id string) (*AppInfo, error) {
	return &AppInfo{}, nil
}

func (m *MasterCard) CreateNode() error {
	// TODO: WIll not impl
	return nil
}

func (m *MasterCard) GetNode(addr string) error {
	// TODO: WIll not impl
	return nil
}

func (m *MasterCard) Invite() error {
	// TODO: WIll not impl
	return nil
}

func (m *MasterCard) Join() error {
	// TODO: WIll not impl
	return nil
}

func (m *MasterCard) Authorize(req *AuthorizationRequest) (*AuthorizationResponse, error) {
	return &AuthorizationResponse{}, nil
}

func (m *MasterCard) GetLastBlock() error {
	return nil
}

func (m *MasterCard) GetBlock(id string) error {
	return nil
}

func (m *MasterCard) PostTransaction() error {
	//endpoint := fmt.Sprintf("%sentry", API_ENDPOINT)
	subProcess := exec.Command("node", "util/app.js")

	stdin, err := subProcess.StdinPipe()
	if err != nil {
		return err
	}

	if _, err := io.WriteString(stdin, fmt.Sprintf("%s\n", m.keystoreloc)); err != nil {
		return err
	}
	if _, err := io.WriteString(stdin, fmt.Sprintf("%s\n", m.keystorepass)); err != nil {
		return err
	}
	if _, err := io.WriteString(stdin, fmt.Sprintf("%s\n", m.keyalias)); err != nil {
		return err
	}
	if _, err := io.WriteString(stdin, fmt.Sprintf("%s\n", m.userkey)); err != nil {
		return err
	}
	if _, err := io.WriteString(stdin, fmt.Sprintf("%s\n", m.protoloc)); err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)
	if _, err := io.WriteString(stdin, fmt.Sprintf("%s\n", m.protoloc)); err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)

	if _, err := io.WriteString(stdin, "2\n"); err != nil {
		return err
	}
	if _, err := io.WriteString(stdin, fmt.Sprintf("%s\n", "TODO")); err != nil {
		return err
	}

	return nil
}

func (m *MasterCard) GetTransaction(hash string) error {
	//endpoint := fmt.Sprintf("%sentry/%s", API_ENDPOINT, hash)
	subProcess := exec.Command("node", "util/app.js")

	stdin, err := subProcess.StdinPipe()
	if err != nil {
		return err
	}

	if _, err := io.WriteString(stdin, fmt.Sprintf("%s\n", m.keystoreloc)); err != nil {
		return err
	}
	if _, err := io.WriteString(stdin, fmt.Sprintf("%s\n", m.keystorepass)); err != nil {
		return err
	}
	if _, err := io.WriteString(stdin, fmt.Sprintf("%s\n", m.keyalias)); err != nil {
		return err
	}
	if _, err := io.WriteString(stdin, fmt.Sprintf("%s\n", m.userkey)); err != nil {
		return err
	}
	if _, err := io.WriteString(stdin, fmt.Sprintf("%s\n", m.protoloc)); err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)
	if _, err := io.WriteString(stdin, fmt.Sprintf("%s\n", m.protoloc)); err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)

	if _, err := io.WriteString(stdin, "3\n"); err != nil {
		return err
	}
	return nil
}

func (m *MasterCard) Settle() error {
	return nil
}

// Sign signs one or more (binary) values with the caller's node private key
func (m *MasterCard) Sign() error {
	return nil
}

// GetStatus gets general network status information
func (m *MasterCard) GetStatus() error {
	return nil
}
