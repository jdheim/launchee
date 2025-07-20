/*
 * © 2025-2025 JDHeim.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package frontend

import (
	"strings"
	"testing"
)

func TestNewIconEmpty(t *testing.T) {
	testCases := map[string]struct {
		in   string
		want *Icon
	}{
		"empty":   {"", newDefaultWantIcon("")},
		"invalid": {"invalid.png", newDefaultWantIcon("invalid.png")},
		"valid":   {"../../../build/appicon.png", newDefaultWantIcon("../../../build/appicon.png")},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := NewIcon(testCase.in)
			if got.Path != testCase.want.Path {
				t.Errorf("NewIcon(%q) = %q, want %q", testCase.in, got.Path, testCase.want.Path)
			}
			if name == "valid" {
				if len(got.Bytes) == 0 {
					t.Errorf("NewIcon(%q) = empty bytes", testCase.in)
				}
				if !strings.HasPrefix(got.Base64, testCase.want.Base64) || got.Base64 == testCase.want.Base64 {
					t.Errorf("NewIcon(%q) = %q, want prefix %q", testCase.in, got.Base64, testCase.want.Base64)
				}
			} else {
				if len(got.Bytes) != 0 {
					t.Errorf("NewIcon(%q) = not empty bytes", testCase.in)
				}
				if got.Base64 != testCase.want.Base64 {
					t.Errorf("NewIcon(%q) = %q, want %q", testCase.in, got.Base64, testCase.want.Base64)
				}
			}
		})
	}
}

func TestIsValidIcon(t *testing.T) {
	testCases := map[string]struct {
		in   string
		want bool
	}{
		"empty": {"", false},
		"doc":   {"test.doc", false},
		"pdf":   {"test.pdf", false},
		"apng":  {"test.apng", true},
		"avif":  {"test.avif", true},
		"bmp":   {"test.bmp", true},
		"gif":   {"test.gif", true},
		"ico":   {"test.ico", true},
		"jpg":   {"test.jpg", true},
		"jpeg":  {"test.jpeg", true},
		"png":   {"test.png", true},
		"svg":   {"test.svg", true},
		"tif":   {"test.tif", true},
		"tiff":  {"test.tiff", true},
		"webp":  {"test.webp", true},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := IsValidIcon(testCase.in)
			if got != testCase.want {
				t.Errorf("IsValidIcon(%q) = %t, want %t", testCase.in, got, testCase.want)
			}
		})
	}
}

func TestSupportedExtensions(t *testing.T) {
	got := SupportedExtensions()
	want := "apng, avif, bmp, gif, ico, jpeg, jpg, png, svg, tif, tiff, webp"
	if got != want {
		t.Errorf("SupportedExtensions() = %q, want %q", got, want)
	}
}

func newDefaultWantIcon(path string) *Icon {
	return &Icon{
		Path:   path,
		Base64: "data:image/png;base64,",
	}
}
