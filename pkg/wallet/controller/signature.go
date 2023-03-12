package controller

import "fmt"

// Sign signs the data with the given account.
func (w *DIDControllerImpl) Sign(data []byte) ([]byte, error) {
	return w.primaryAccount.Sign(data)
}

// Verify verifies the signature with the given account.
func (w *DIDControllerImpl) Verify(data, sig []byte) (bool, error) {
	return w.primaryAccount.Verify(data, sig)
}

// SignWithAccount signs the data with the given account.
func (w *DIDControllerImpl) SignWithAccount(data []byte, accountName string) ([]byte, error) {
	accs, err := w.ListAccounts()
	if err != nil {
		return nil, err
	}
	for _, acc := range accs {
		if acc.Name() == accountName {
			return acc.Sign(data)
		}
	}
	return nil, fmt.Errorf("account %s not found", accountName)
}

// VerifyWithAccount verifies the signature with the given account.
func (w *DIDControllerImpl) VerifyWithAccount(data, sig []byte, accountName string) (bool, error) {
	accs, err := w.ListAccounts()
	if err != nil {
		return false, err
	}
	for _, acc := range accs {
		if acc.Name() == accountName {
			return acc.Verify(data, sig)
		}
	}
	return false, fmt.Errorf("account %s not found", accountName)
}
