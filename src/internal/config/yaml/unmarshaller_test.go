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
	"testing"

	"github.com/jdheim/launchee/internal/lctx"
	"github.com/jdheim/launchee/internal/test/stub"
)

func TestUnmarshalConfigs(t *testing.T) {
	lctx.LoggerImpl = stub.LoggerStub{}
	ConfigPathImpl = stub.ConfigPathValidStub{}
	_, err := UnmarshalConfigs()
	if err != nil {
		t.Errorf("UnmarshalConfigs() = %v", err)
	}
}
