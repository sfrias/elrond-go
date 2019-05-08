package mock

import (
	"github.com/ElrondNetwork/elrond-go-sandbox/data"
)

// ForkDetectorMock is a mock implementation for the ForkDetector interface
type ForkDetectorMock struct {
	AddHeaderCalled                 func(header data.HeaderHandler, hash []byte, isProcessed bool) error
	RemoveHeadersCalled             func(nonce uint64, hash []byte)
	CheckForkCalled                 func() (bool, uint64)
	GetHighestFinalBlockNonceCalled func() uint64
	ProbableHighestNonceCalled      func() uint64
}

// AddHeader is a mock implementation for AddHeader
func (f *ForkDetectorMock) AddHeader(header data.HeaderHandler, hash []byte, isProcessed bool) error {
	return f.AddHeaderCalled(header, hash, isProcessed)
}

// RemoveHeaders is a mock implementation for RemoveHeaders
func (f *ForkDetectorMock) RemoveHeaders(nonce uint64, hash []byte) {
	f.RemoveHeadersCalled(nonce, hash)
}

// CheckFork is a mock implementation for CheckFork
func (f *ForkDetectorMock) CheckFork() (bool, uint64) {
	return f.CheckForkCalled()
}

// GetHighestFinalBlockNonce is a mock implementation for GetHighestFinalBlockNonce
func (f *ForkDetectorMock) GetHighestFinalBlockNonce() uint64 {
	return f.GetHighestFinalBlockNonceCalled()
}

// GetProbableHighestNonce is a mock implementation for GetProbableHighestNonce
func (f *ForkDetectorMock) ProbableHighestNonce() uint64 {
	return f.ProbableHighestNonceCalled()
}
