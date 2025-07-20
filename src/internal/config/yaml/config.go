/*
 * © 2025-2025 JDHeim
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

package yaml

import (
	"strings"

	"github.com/jdheim/launchee/internal/config/frontend"
)

type config struct {
	Title     string
	Shortcuts []*shortcut
}

type shortcut struct {
	Name    string
	Icon    string
	Command string
	Url     string
	Patch   string `yaml:"$patch"`
}

func newConfigWithoutShortcuts(title string) *config {
	return &config{
		Title: title,
	}
}

func (yc *config) toFrontendConfig() *frontend.Config {
	frontendConfig := frontend.NewConfig(len(yc.Shortcuts))
	yc.overrideFrontendConfig(frontendConfig)
	return frontendConfig
}

func (yc *config) overrideFrontendConfig(config *frontend.Config) {
	if yc.Title != "" {
		config.UI.Nav.Title = yc.Title
	}
	if frontendShortcuts := yc.toFrontendShortcuts(); frontendShortcuts != nil {
		config.Shortcuts = frontendShortcuts
	}
}

func (yc *config) toFrontendShortcuts() []*frontend.Shortcut {
	shortcutCount := len(yc.Shortcuts)
	if shortcutCount == 0 {
		return nil
	}
	frontendShortcuts := make([]*frontend.Shortcut, shortcutCount)
	for i := range yc.Shortcuts {
		frontendShortcuts[i] = yc.toFrontendShortcut(i)
	}
	return frontendShortcuts
}

func (yc *config) toFrontendShortcut(i int) *frontend.Shortcut {
	return &frontend.Shortcut{
		Id:      i,
		Name:    yc.Shortcuts[i].Name,
		Icon:    frontend.NewIcon(yc.Shortcuts[i].Icon),
		Command: strings.Fields(yc.Shortcuts[i].Command),
		Url:     yc.Shortcuts[i].Url,
	}
}

func (s *shortcut) isPatchMode() bool {
	return s.Patch != "" && (s.Patch == patchDelete || s.Patch == patchMerge)
}
