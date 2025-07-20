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

const launchee7ConfigPath = "internal/test/stub/stub_config/launchee-7.yml"
const launcheeOverrideConfigPath = "internal/test/stub/stub_config/launchee-override.yml"
const launcheeEmptyConfigPath = "internal/test/stub/stub_config/launchee-empty.yml"
const invalidConfigPath = "/usr/bin/echo"
const notExistsConfigPath = "/tmp/not-exists.yml"

type ConfigPathValidStub struct{}

func (ConfigPathValidStub) GetSystemConfigPath() string {
	return launchee7ConfigPath
}

func (ConfigPathValidStub) GetUserConfigPath() string {
	return launcheeOverrideConfigPath
}

type ConfigPathInvalidStub struct{}

func (ConfigPathInvalidStub) GetSystemConfigPath() string {
	return invalidConfigPath
}

func (ConfigPathInvalidStub) GetUserConfigPath() string {
	return invalidConfigPath
}

type SystemConfigPathInvalidStub struct{}

func (SystemConfigPathInvalidStub) GetSystemConfigPath() string {
	return invalidConfigPath
}

func (SystemConfigPathInvalidStub) GetUserConfigPath() string {
	return launchee7ConfigPath
}

type UserConfigPathInvalidStub struct{}

func (UserConfigPathInvalidStub) GetSystemConfigPath() string {
	return launchee7ConfigPath
}

func (UserConfigPathInvalidStub) GetUserConfigPath() string {
	return invalidConfigPath
}

type ConfigPathNotExistsStub struct{}

func (ConfigPathNotExistsStub) GetSystemConfigPath() string {
	return notExistsConfigPath
}

func (ConfigPathNotExistsStub) GetUserConfigPath() string {
	return notExistsConfigPath
}

type SystemConfigPathNotExistsStub struct{}

func (SystemConfigPathNotExistsStub) GetSystemConfigPath() string {
	return notExistsConfigPath
}

func (SystemConfigPathNotExistsStub) GetUserConfigPath() string {
	return launchee7ConfigPath
}

type UserConfigPathNotExistsStub struct{}

func (UserConfigPathNotExistsStub) GetSystemConfigPath() string {
	return launchee7ConfigPath
}

func (UserConfigPathNotExistsStub) GetUserConfigPath() string {
	return notExistsConfigPath
}

type ConfigPathEmptyStub struct{}

func (ConfigPathEmptyStub) GetSystemConfigPath() string {
	return launcheeEmptyConfigPath
}

func (ConfigPathEmptyStub) GetUserConfigPath() string {
	return launcheeEmptyConfigPath
}

type SystemConfigPathEmptyStub struct{}

func (SystemConfigPathEmptyStub) GetSystemConfigPath() string {
	return launcheeEmptyConfigPath
}

func (SystemConfigPathEmptyStub) GetUserConfigPath() string {
	return launchee7ConfigPath
}

type UserConfigPathEmptyStub struct{}

func (UserConfigPathEmptyStub) GetSystemConfigPath() string {
	return launchee7ConfigPath
}

func (UserConfigPathEmptyStub) GetUserConfigPath() string {
	return launcheeEmptyConfigPath
}
