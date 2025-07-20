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

package yaml

import (
	"github.com/jdheim/launchee/internal/lctx"
)

const (
	patchDelete  = "delete"
	patchMerge   = "merge"
	patchReplace = "replace"
)

func (yc *config) sanitize() *config {
	if yc == nil {
		return nil
	} else if yc.Shortcuts == nil {
		return yc
	}
	sanitizedShortcuts := make([]*shortcut, 0, len(yc.Shortcuts))
	processed := make(map[string]bool)
	for _, shortcut := range yc.Shortcuts {
		if !processed[shortcut.Name] && !shortcut.isPatchMode() {
			processed[shortcut.Name] = true
			sanitizedShortcuts = append(sanitizedShortcuts, shortcut)
		}
	}
	yc.Shortcuts = sanitizedShortcuts
	return yc
}

func (yc *config) merge(other *config) *config {
	if yc == nil || other == nil {
		return yc
	}
	merged := newConfigWithoutShortcuts(yc.Title)
	if other.Title != "" {
		merged.Title = other.Title
	}
	if len(other.Shortcuts) != 0 {
		merged.Shortcuts = yc.mergeShortcuts(other)
	} else {
		merged.Shortcuts = yc.Shortcuts
	}
	return merged
}

func (yc *config) mergeShortcuts(other *config) []*shortcut {
	mergedShortcuts := make([]*shortcut, 0, len(yc.Shortcuts)+len(other.Shortcuts))
	otherShortcuts := toShortcutMapByName(other.Shortcuts)
	processed := make(map[string]bool)

	lctx.LogInfo("----- Configuration merge started -----")
	for _, shortcut := range yc.Shortcuts {
		if otherShortcut, found := otherShortcuts[shortcut.Name]; found {
			processed[shortcut.Name] = true
			if otherShortcut.Patch == patchDelete {
				lctx.LogInfof("Deleting %+v", otherShortcut)
				continue
			} else if otherShortcut.Patch == patchMerge {
				mergedShortcut := shortcut.merge(otherShortcut)
				lctx.LogInfof("Overridding with %+v", mergedShortcut)
				mergedShortcuts = append(mergedShortcuts, mergedShortcut)
				continue
			}
			lctx.LogInfof("Replacing with %+v", otherShortcut)
			mergedShortcuts = append(mergedShortcuts, otherShortcut)
		} else {
			lctx.LogInfof("Adding %+v", shortcut)
			mergedShortcuts = append(mergedShortcuts, shortcut)
		}
	}

	for _, otherShortcut := range other.Shortcuts {
		if !processed[otherShortcut.Name] && otherShortcut.Patch != patchDelete && otherShortcut.Patch != patchMerge {
			lctx.LogInfof("Adding not processed one %+v", otherShortcut)
			processed[otherShortcut.Name] = true
			mergedShortcuts = append(mergedShortcuts, otherShortcut)
		}
	}
	lctx.LogInfo("----- Configuration merge finished -----")
	return mergedShortcuts
}

func toShortcutMapByName(shortcuts []*shortcut) map[string]*shortcut {
	shortcutMap := make(map[string]*shortcut)
	for _, shortcut := range shortcuts {
		shortcutMap[shortcut.Name] = shortcut
	}
	return shortcutMap
}

func (s *shortcut) merge(other *shortcut) *shortcut {
	s.Patch = patchMerge
	if other.Icon != "" {
		s.Icon = other.Icon
	}
	if other.Command != "" {
		s.Command = other.Command
		if other.CommandArgs != "" {
			s.CommandArgs = other.CommandArgs
		}
		s.Url = ""
	} else if other.Url != "" {
		s.Command = ""
		s.CommandArgs = ""
		s.Url = other.Url
	}
	return s
}
