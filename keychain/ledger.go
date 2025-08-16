// Copyright (C) 2019-2024, Lux Industries Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package keychain

import (
	"github.com/luxfi/ids"
)

// Ledger interface for the ledger wrapper
type Ledger interface {
	Version() (major, minor, patch uint32, err error)
	Address(displayHRP string, addressIndex uint32) (ids.ShortID, error)
	Addresses(addressIndices []uint32) ([]ids.ShortID, error)
	SignHash(hash []byte, addressIndices []uint32) ([][]byte, error)
	Sign(unsignedTxBytes []byte, addressIndices []uint32) ([][]byte, error)
	Disconnect() error
}
