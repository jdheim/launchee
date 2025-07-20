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

package stub

type ConfigPathValidStub struct{}

func (ConfigPathValidStub) GetSystemConfigPath() string {
	return "../../test/stub/stub_config/launchee-7.yml"
}

func (ConfigPathValidStub) GetUserConfigPath() string {
	return "../../test/stub/stub_config/launchee-override.yml"
}

type ConfigPathInvalidStub struct{}

func (ConfigPathInvalidStub) GetSystemConfigPath() string {
	return "/usr/bin/echo"
}

func (ConfigPathInvalidStub) GetUserConfigPath() string {
	return "/usr/bin/echo"
}
