# AMPEL attestation framework

This repository contains the [AMPEL](https://github.com/carabiner-dev/ampel)
attestation framework. Our framework builds on top of the [in-toto attestation
formats](https://github.com/in-toto/attestation). In a sens _it is_ the in-toto
attestation framework as it is 100% compatible, only defining extra methods
that make it possible to use attested data in AMPEL's policy enforcement
context.

At this time, the framework only contains interface definitions for the
attestation components, the query system and systems to store and retrieve them.
Code here is intended to have a minimal dependency footprint. All implementations
are stored in other repositories and projects (concrete pointers TBD).

## Attestation Components

The framework splits the attestation stack in the following components:

#### Envelope

The envelope abstracts any wrapper (real or conceptual) around an attestation.
Usually envelopes provide strong identity or cryptogrpahic safeguards on the
attestation data.

Sometimes, envelopes wrap other wrappers. Ideally, parsers will stop at the
outermost layer and squash any internal ones. In general, the envelope format
is not relevant on the policy enforcement context. After parsing an attestation's
envelope, the only preserved data is the signature/identity verification
information.

Envelope examples include DSSE, sigstore bundles, etc. Envelopes can also be 
conceptual, for example an attestation paired with an esternal signature file
can be coerced into an Envelope providing verification data.

#### Statement

The statement is a 1:1 map of an in-toto statement with addutional methods. Just
as in vanila in-toto, a Statement's data is split between its subjects and a
predicate. It has a Type and a Predicate type.

At this time, ampel only supports in-toto statements and while there are no
plans to support other formats, anything that can be coerced into a
predicate+subjects shape could potentially be used as a statement.

#### Subject

The subject data structure is also a 1:1 wrapper on the in-toto resource descriptor.
It captures the hashes, name and/or URI of whetever the claims in the attestation
are about.

#### Predicate

The predicate captures the claim data. Just as in-toto, the predicate is expected
to be of any format, it has a PredicateType which is used to filter and categorize
attested data.

In AMPEL, the predicate is the main work unit. Signature verification, attestation
sorting and filtering, all are done before the runtime has access to the data.
This is why, the Predicate in AMPEL is overloaded with extra data such as its
origin, its siganture verification data, etc. When evaluating a policy, the
runtime does not have access to the subject or envelope. Only the predicate but
it is conveniently loaded with the applicable subject and signature data in case
the policy needs them.

## Attestation Repositories

Attestation repositories store and retrieve attestation data. Anything that
implements the `Repository` interface can be used as a repository, from simple
implementations like a directory reader to more complex ones like OCI registries,
git repos or specialized storage systems.

If repositories don't wish to implement or expose all capabilities they can
implement only the `Fetcher` the `Storer` interfaces.

### Specialized Vaariants

As we identify useful capabilities, more fetcher and storer types will be defined.
This allows implementations,such as the
[Carabiner collector](https://github.com/policylabs/collector), to leverage
optimized capabilities from the repository implementaion backends.

At present we have defined the `FetcherBySubject`, `FetcherByPredicateType` and
`FetcherByPredicateTypeAndSubject` specialized fetchers.

## Attestation Queries

Once a repository is queried and returns a set of attestations, AMPEL can use
attestation queries to sort or filter attestation sets returned from repositories.

Queries operate at the Envelope level and filter attestations based on their
properties. AMPEL makes extensive use of queries and while the heavy lift of
selecting approriate attestation is expected to be performed by the repository,
AMPEL uses queries to refine the statements what will eventually be exposed to
the runtime engines for evaluation. This means AMPEL can overcome the limitation
of simple repository drivers (for example the JSONL driver) to provide missing
filtering capabilities.

---

## Copyright

The contents of this repository are Copyright by The Policy Labs Project Contributors and
released under the Apache-2.0 license. Feel free to contribute patches, comments
or bug reports.
