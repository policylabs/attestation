// SPDX-FileCopyrightText: Copyright 2025 The Policy Labs Project Contributors
// SPDX-License-Identifier: Apache-2.0

package attestation

import (
	"errors"
	"io"
)

var ErrNotCorrectFormat = errors.New("data does not conform to parser's format")

// EnvelopeParser abstracts an object that reads data and on the other end
// returns an attestation envelope. It is the EnvelopeParser's job to verify
// any signatures or other cryptographic material protecting the contained
// attestation.
type EnvelopeParser interface {
	ParseStream(r io.Reader) ([]Envelope, error)
	FileExtensions() []string
}

// StatementParser is an object that parses data and returns a statement.
type StatementParser interface {
	Parse([]byte) (Statement, error)
}

// The predicate parser reads a predicate's data and returns an object that
// implements the Predicate interface.
type PredicateParser interface {
	Parse([]byte) (Predicate, error)
	SupportsType(...PredicateType) bool
}
