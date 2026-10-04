// SPDX-FileCopyrightText: Copyright 2025 The Policy Labs Project Contributors
// SPDX-License-Identifier: Apache-2.0

package attestation

import (
	"context"
	"errors"
)

var (
	ErrFetcherMethodNotImplemented = errors.New("fetching method not implemented")
	ErrStorerMethodNotImplemented  = errors.New("storing method not implemented")
)

// Repository is an abstraction of a system that can store attestations, serve
// them or both.
// Repositories can express their capabilities by implementing the Storer and/or
// Fetcher interfaces and their specialized variants.
type Repository interface{} //nolint:iface // To be implemented

// AttestationFetcher is the the trait that repositories that can fetch
// attestations must implement
type Fetcher interface {
	Fetch(context.Context, FetchOptions) ([]Envelope, error)
}

// FetcherBySubject is a fetcher that can filter natively by subject hashes.
type FetcherBySubject interface {
	FetchBySubject(context.Context, FetchOptions, []Subject) ([]Envelope, error)
}

// FetcherByPredicateType is a fetcher that can filter natively by predictae types.
type FetcherByPredicateType interface {
	FetchByPredicateType(context.Context, FetchOptions, []PredicateType) ([]Envelope, error)
}

// FetcherByPredicateTypeAndSubject is a fetcher that can filter natively by
// predictae types and subject.
type FetcherByPredicateTypeAndSubject interface {
	FetchByPredicateTypeAndSubject(context.Context, FetchOptions, []PredicateType, []Subject) ([]Envelope, error)
}

// Storer is a repository that can store attestations
type Storer interface {
	Store(context.Context, StoreOptions, []Envelope) error
}

// StoreOptions control how attestations are retrieved from a Fetcher. All
// repositories implementing the Fetcher interface are expected to honor FetchOptions.
type FetchOptions struct {
	// Maximum number of attestations to return
	Limit int

	// Limit data reads to this size (bytes per attestation)
	MaxReadSize int64

	// Apply query after fetching attestations
	Query *Query
}

// StoreOptions control how attestations are stored in the storer. All repositories
// implementing the Storer interface are expected to honor StoreOptions.
type StoreOptions struct{}
