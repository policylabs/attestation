// SPDX-FileCopyrightText: Copyright 2025 The Policy Labs Project Contributors
// SPDX-License-Identifier: Apache-2.0

package attestation

import "testing"

const (
	sha256Algo = "sha256"
	sha512Algo = "sha512"
	md5Algo    = "md5"

	sha256Val = "abc123"
	sha512Val = "def456"
	md5Val    = "ghi789"
)

type fakeSubject struct {
	name, uri string
	digest    map[string]string
}

func (fs *fakeSubject) GetName() string {
	return fs.name
}

func (fs *fakeSubject) GetUri() string {
	return fs.uri
}

func (fs *fakeSubject) GetDigest() map[string]string {
	return fs.digest
}

func TestSubjectsMatch(t *testing.T) {
	tests := []struct {
		name     string
		s1       Subject
		s2       Subject
		expected bool
	}{
		{
			name:     "single-identical",
			s1:       &fakeSubject{digest: map[string]string{sha256Algo: sha256Val}},
			s2:       &fakeSubject{digest: map[string]string{sha256Algo: sha256Val}},
			expected: true,
		},
		{
			name:     "identical-multiple-match",
			s1:       &fakeSubject{digest: map[string]string{sha256Algo: sha256Val, sha512Algo: sha512Val}},
			s2:       &fakeSubject{digest: map[string]string{sha256Algo: sha256Val, sha512Algo: sha512Val}},
			expected: true,
		},
		{
			name:     "hash-mismatch",
			s1:       &fakeSubject{digest: map[string]string{sha256Algo: sha256Val}},
			s2:       &fakeSubject{digest: map[string]string{sha256Algo: "xyz789"}},
			expected: false,
		},
		{
			name:     "no-common",
			s1:       &fakeSubject{digest: map[string]string{sha256Algo: sha256Val}},
			s2:       &fakeSubject{digest: map[string]string{sha512Algo: sha512Val}},
			expected: false,
		},
		{
			name:     "empty-s1",
			s1:       &fakeSubject{digest: map[string]string{}},
			s2:       &fakeSubject{digest: map[string]string{sha256Algo: sha256Val}},
			expected: false,
		},
		{
			name:     "partial-overlap-with-match",
			s1:       &fakeSubject{digest: map[string]string{sha256Algo: sha256Val, sha512Algo: sha512Val}},
			s2:       &fakeSubject{digest: map[string]string{sha256Algo: sha256Val, md5Algo: md5Val}},
			expected: true,
		},
		{
			name:     "partial-overlap-with-mismatch",
			s1:       &fakeSubject{digest: map[string]string{sha256Algo: sha256Val, sha512Algo: sha512Val}},
			s2:       &fakeSubject{digest: map[string]string{sha256Algo: "different", md5Algo: md5Val}},
			expected: false,
		},
		{
			name:     "s2-subset-s1-with-match",
			s1:       &fakeSubject{digest: map[string]string{sha256Algo: sha256Val, sha512Algo: sha512Val, md5Algo: md5Val}},
			s2:       &fakeSubject{digest: map[string]string{sha256Algo: sha256Val}},
			expected: true,
		},
		{
			name:     "multiple-common-with-mismatch",
			s1:       &fakeSubject{digest: map[string]string{sha256Algo: sha256Val, sha512Algo: sha512Val}},
			s2:       &fakeSubject{digest: map[string]string{sha256Algo: sha256Val, sha512Algo: "wrong"}},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubjectsMatch(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("SubjectsMatch() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
