package headersCache

import "github.com/pkg/errors"

// ErrHeaderNotFound signals that the header that was searched was not found in the pool
var ErrHeaderNotFound = errors.New("cannot find header in cache")

// ErrInvalidHeadersCacheParameter signals that parameters for items cache are invalid
var ErrInvalidHeadersCacheParameter = errors.New("invalid items cache parameters")
