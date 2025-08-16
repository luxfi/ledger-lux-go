// Copyright (C) 2019-2024, Lux Industries Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package keychain

import "github.com/luxfi/ids"

// Signer interface for signing operations
type Signer interface {
	Sign(message []byte) ([]byte, error)
	SignHash(hash []byte) ([]byte, error)
	Address() ids.ShortID
}

// Keychain interface for managing signers
type Keychain interface {
	// Get returns the Signer for the given address
	Get(address ids.ShortID) (Signer, bool)
	// List returns all addresses in the keychain
	List() []ids.ShortID
}