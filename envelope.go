// SPDX-FileCopyrightText: Copyright 2025 The Policy Labs Project Contributors
// SPDX-License-Identifier: Apache-2.0

package attestation

// EnvelopeVerificationOptions are values used to verify signed envelopes that
// may be useful to all envelop implementations.
type EnvelopeVerificationOptions struct {
	// AllowUnsigned causes the verification to fail if an envelope is not signed.
	AllowUnsigned bool
}

// EnvVerOptsConvertable is an interface that any verification material that
// can be converted to an EnvelopeVerificationOptions can implement to get
// data from the verifier.
type EnvVerOptsConvertable interface {
	ToEnvelopeVerificationOptions() EnvelopeVerificationOptions
}

// Envelope is a construct that wraps a statement, its signatures and all the
// verification material. The goal of this abstraction is to get a single
// interface to verify statements, even when all the bits amy be in separate
// files.
type Envelope interface {
	GetStatement() Statement
	GetPredicate() Predicate
	GetSignatures() []Signature
	GetCertificate() Certificate
	GetVerification() Verification
	Verify(...any) error
}
