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
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Icon struct {
	Path   string
	Bytes  []byte
	Base64 string
}

const defaultMimeType = "image/png"

var mimeTypes = map[string]string{
	"apng": "image/apng",
	"avif": "image/avif",
	"bmp":  "image/bmp",
	"gif":  "image/gif",
	"ico":  "image/x-icon",
	"jpeg": "image/jpeg",
	"jpg":  "image/jpeg",
	"png":  "image/png",
	"svg":  "image/svg+xml",
	"tif":  "image/tiff",
	"tiff": "image/tiff",
	"webp": "image/webp",
}

func NewIcon(path string) *Icon {
	newIcon := &Icon{
		Path: path,
	}
	newIcon.encodeToBase64()
	return newIcon
}

func NewIconWithBytes(bytes []byte) *Icon {
	newIcon := &Icon{
		Bytes: bytes,
	}
	newIcon.encodeToBase64()
	return newIcon
}

func (i *Icon) encodeToBase64() {
	if len(i.Bytes) == 0 {
		if i.Path != "" {
			if i.Bytes = readFile(i.Path); len(i.Bytes) == 0 {
				i.Base64 = defaultBase64()
				return
			}
		} else {
			i.Base64 = defaultBase64()
			return
		}
	}
	mimeType := getMimeType(i.Path)
	base64Str := base64.StdEncoding.EncodeToString(i.Bytes)
	i.Base64 = fmt.Sprintf("data:%s;base64,%s", mimeType, base64Str)
}

func readFile(path string) []byte {
	if _, err := os.Stat(path); err != nil {
		return nil
	}
	bytes, _ := os.ReadFile(path)
	return bytes
}

func defaultBase64() string {
	return fmt.Sprintf("data:%s;base64,", defaultMimeType)
}

func IsValidIcon(path string) bool {
	_, exists := mimeTypes[getExtension(path)]
	return exists
}

func SupportedExtensions() string {
	extensions := make([]string, 0, len(mimeTypes))
	for extension := range mimeTypes {
		extensions = append(extensions, extension)
	}
	sort.Strings(extensions)
	return fmt.Sprint(strings.Join(extensions, ", "))
}

func getMimeType(path string) string {
	if mimeType, exists := mimeTypes[getExtension(path)]; exists {
		return mimeType
	}
	return defaultMimeType
}

func getExtension(path string) string {
	return strings.ToLower(strings.TrimPrefix(filepath.Ext(path), "."))
}
