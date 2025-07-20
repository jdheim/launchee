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
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jdheim/launchee/build"
)

func TestNewUI(t *testing.T) {
	testCases := map[string]struct {
		in   int
		want func() *UI
	}{
		"zero": {
			0, func() *UI {
				return newDefaultWantUI()
			},
		},
		"one": {
			1, func() *UI {
				return newDefaultWantUI()
			},
		},
		"three": {
			3, func() *UI {
				defaultWantUI := newDefaultWantUI()
				defaultWantUI.Content.IconColumns = 3
				defaultWantUI.Content.IconsPerRow = 5
				return defaultWantUI
			},
		},
		"five": {
			5, func() *UI {
				defaultWantUI := newDefaultWantUI()
				defaultWantUI.Content.IconColumns = 5
				defaultWantUI.Content.IconsPerRow = 5
				return defaultWantUI
			},
		},
		"fifteen": {
			15, func() *UI {
				defaultWantUI := newDefaultWantUI()
				defaultWantUI.Content.IconColumns = 15
				defaultWantUI.Content.IconsPerRow = 15
				return defaultWantUI
			},
		},
		"hundred": {
			100, func() *UI {
				defaultWantUI := newDefaultWantUI()
				defaultWantUI.Content.IconColumns = 20
				defaultWantUI.Content.IconsPerRow = 20
				return defaultWantUI
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := NewUI(testCase.in)
			if diff := cmp.Diff(testCase.want(), got); diff != "" {
				t.Errorf("NewUI(%d) = diff -want +got\n%s", testCase.in, diff)
			}
		})
	}
}

func TestWidth(t *testing.T) {
	testCases := map[string]struct {
		in   int
		want int
	}{
		"zero":    {0, 282},
		"one":     {1, 282},
		"three":   {3, 282},
		"five":    {5, 282},
		"fifteen": {15, 802},
		"hundred": {100, 1062},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			testUI := NewUI(testCase.in)
			got := testUI.Width()
			if got != testCase.want {
				t.Errorf("NewUI(%d).Width() = %d, want %d", testCase.in, got, testCase.want)
			}
		})
	}
}

func TestHeight(t *testing.T) {
	testCases := map[string]struct {
		in   int
		want int
	}{
		"zero":    {0, 105},
		"one":     {1, 105},
		"three":   {3, 105},
		"five":    {5, 105},
		"fifteen": {15, 105},
		"hundred": {100, 313},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			testUI := NewUI(testCase.in)
			got := testUI.Height(testCase.in)
			if got != testCase.want {
				t.Errorf("NewUI(%[1]d).Height(%[1]d) = %d, want %d", testCase.in, got, testCase.want)
			}
		})
	}
}

func newDefaultWantUI() *UI {
	return &UI{
		Nav: &Nav{
			Title:      "Launchee",
			AppIcon:    NewIconWithBytes(build.GetAppIconBytes()),
			IconSize:   23,
			IconUrl:    "https://launchee.jdheim.com",
			MenuHeight: 8,
		},
		Content: &Content{
			IconColumns: 1,
			IconsPerRow: 5,
			IconSize:    32,
			Margin:      5,
		},
	}
}
