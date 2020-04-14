package pubkeyConverter

import (
	"encoding/hex"
	"fmt"

	"github.com/ElrondNetwork/elrond-go/data/state"
)

// hexPubkeyConverter encodes or decodes provided public key as/from hex
type hexPubkeyConverter struct {
	len int
}

// NewHexPubkeyConverter returns a hexPubkeyConverter instance
func NewHexPubkeyConverter(addressLen int) (*hexPubkeyConverter, error) {
	if addressLen < 1 {
		return nil, fmt.Errorf("%w when creating hex address converter, addressLen should have been greater than 0",
			state.ErrInvalidAddressLength)
	}
	if addressLen%2 == 1 {
		return nil, fmt.Errorf("%w when creating hex address converter, addressLen should have been an even number",
			state.ErrInvalidAddressLength)
	}

	return &hexPubkeyConverter{
		len: addressLen,
	}, nil
}

// Decode converts the provided public key string as hex decoded bytes
func (ppc *hexPubkeyConverter) Decode(humanReadable string) ([]byte, error) {
	buff, err := hex.DecodeString(humanReadable)
	if err != nil {
		return nil, err
	}

	if len(buff) != ppc.len {
		return nil, fmt.Errorf("%w when converting to address, expected length %d, received %d",
			state.ErrWrongSize, ppc.len, len(buff))
	}

	return buff, nil
}

// Encode converts the provided bytes in a form that this converter can decode. In this case it will encode to hex
func (ppc *hexPubkeyConverter) Encode(pkBytes []byte) string {
	return hex.EncodeToString(pkBytes)
}

// CreateAddressFromString creates an address container based on the provided string
func (ppc *hexPubkeyConverter) CreateAddressFromString(humanReadable string) (state.AddressContainer, error) {
	buff, err := ppc.Decode(humanReadable)
	if err != nil {
		return nil, err
	}

	return state.NewAddress(buff), nil
}

// CreateAddressFromBytes creates an address container based on the provided public key bytes
func (ppc *hexPubkeyConverter) CreateAddressFromBytes(pkBytes []byte) (state.AddressContainer, error) {
	if len(pkBytes) != ppc.len {
		return nil, fmt.Errorf("%w when converting to address, expected length %d, received %d",
			state.ErrWrongSize, ppc.len, len(pkBytes))
	}

	return state.NewAddress(pkBytes), nil
}

// Len returns the decoded address length
func (ppc *hexPubkeyConverter) Len() int {
	return ppc.len
}

// IsInterfaceNil returns true if there is no value under the interface
func (ppc *hexPubkeyConverter) IsInterfaceNil() bool {
	return ppc == nil
}