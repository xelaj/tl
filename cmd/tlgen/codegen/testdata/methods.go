package somepkg

// Registers a validated phone number in the system.
type AccountDeleteAccountParams struct {
	Reason string
}

func (*AccountDeleteAccountParams) CRC() uint32 { return 0x418d4e0b }

// Registers a validated phone number in the system.
func (c *Client) AccountDeleteAccount(reason string) (bool, error) {
	responseData, err := c.MakeRequest(&AccountDeleteAccountParams{Reason: reason})
	if err != nil {
		return false, errors.Wrap(err, "sending AccountDeleteAccount")
	}

	resp, ok := responseData.(bool)
	if !ok {
		panic("got invalid response type: " + reflect.TypeOf(responseData).String())
	}
	return resp, nil
}


