#!/usr/bin/env bash
# BUILDS LAUNCHEE

#
# © 2025-2025 JDHeim.com
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

[[ -f "$(dirname "${BASH_SOURCE[0]}")/common/functions.sh" ]] && . "$(dirname "${BASH_SOURCE[0]}")/common/functions.sh"

usage() {
  cat << EOF
Usage: $(basename "$0") [OPTION]

Builds Launchee for Linux

OPTIONS:
  -j                     Dry-run JReleaser release
  -w                     Builds Launchee for Windows
EOF
  exit 1
}

main() {
  cd ..
  readOptions "$@"
  local projectVersion
  projectVersion="$(getProjectVersion)"
  scripts/common/updateVersion.sh "${projectVersion}"
  scripts/common/updateCopyright.sh
  cd src
  if [[ "${buildForWindows:-false}" == "true" ]]; then
    buildForWindows "${projectVersion}"
  else
    buildForLinux "${projectVersion}"
  fi
}

readOptions() {
  while getopts ":hjw#" option; do
    case "${option}" in
      j) dryRunJReleaserRelease;;
      w) buildForWindows="true";;
      \#) buildForJdvm="true";;
      h|?) usage;;
    esac
  done
}

dryRunJReleaserRelease() {
  step "Dry-run JReleaser release"
  if [[ -z "${GITHUB_TOKEN-}" ]]; then
    echo -e "${ERROR} The GITHUB_TOKEN env variable is not set"
    exit 1
  fi
  JRELEASER_GITHUB_TOKEN=${GITHUB_TOKEN-} run jreleaser release --dry-run --output-directory=target
  exit $?
}

# Requires github.com/wailsapp/go-webview2 in v1.0.21
buildForWindows() {
  local projectVersion="${1}"
  step "Build Launchee ${projectVersion} for Windows"
  run wails build -clean \
      -platform "windows" \
      -webview2 "embed" \
      -ldflags "-X github.com/jdheim/launchee/cmd.appVersion=${projectVersion} \
                -X github.com/jdheim/launchee/cmd.buildForJdvm=${buildForJdvm:-false}"
}

buildForLinux() {
  local projectVersion="${1}"
  step "Build Launchee ${projectVersion} for Linux"
  buildForJdvmPreAction
  run wails build -clean \
    -tags "webkit2_41" \
    -ldflags "-X github.com/jdheim/launchee/cmd.appVersion=${projectVersion} \
              -X github.com/jdheim/launchee/cmd.buildForJdvm=${buildForJdvm:-false}"
  buildForJdvmPostAction
}

buildForJdvmPreAction() {
  if [[ "${buildForJdvm:-false}" == "true" ]]; then
    local jdvmAppIcon="../../jdvm/wayland-icon-fix/icons/apps/jdvm.png"
    if [[ ! -f "${jdvmAppIcon}" ]]; then
      echo -e "${ERROR} Missing ${jdvmAppIcon}"
      exit 1
    fi
    cp -v "build/appicon.png" "build/appicon.png.bak"
    cp -v "${jdvmAppIcon}" "build/appicon.png"
  fi
}

buildForJdvmPostAction() {
  if [[ "${buildForJdvm:-false}" == "true" ]]; then
    mv -v "build/appicon.png.bak" "build/appicon.png"
    cp -v "build/bin/launchee" "../../jdvm/src/jdvm-docker/src/main/docker/scripts/launchee"
  fi
}

main "$@"
