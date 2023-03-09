package service

type ServiceHandler interface {
	// This method is used to get the challenge response from the DID controller.
	BeginRegistration(aka string) ([]byte, error)

	// This is the method that will be called when the user clicks on the "Register" button.
	FinishRegistration(aka string, challengeResponse string) ([]byte, error)

	// This method is used to get the options for the assertion.
	BeginLogin(aka string) ([]byte, error)

	// This is the method that will be called when the user clicks the "Login" button on the login page.
	FinishLogin(aka string, challengeResponse string) ([]byte, error)
}
