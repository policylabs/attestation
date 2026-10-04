// SPDX-FileCopyrightText: Copyright 2025 The Policy Labs Project Contributors
// SPDX-License-Identifier: Apache-2.0

package attestation

// Verification is a minimal abstraction of an identity/signature verification
// results object. This is mostly a placeholder to mark where verification data
// fits in the attestation framework.
type Verification interface {
	// GetVerified returns a bool indicating if the verification was successful.
	GetVerified() bool

	// MatchesIdentity gets a value that an implementation can turn into an
	// identity and returns a bool indicating if the signature verification
	// matches it.
	MatchesIdentity(any) bool
}
