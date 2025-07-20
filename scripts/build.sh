#!/bin/bash
# BUILDS LAUNCHEE

#
# © 2025-2025 JDHeim
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

Builds Launchee

OPTIONS:
  -j                     Dry-run JReleaser release
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
  build "${projectVersion}"
}

readOptions() {
  while getopts ":hj#" option; do
    case "${option}" in
      j) dryRunJReleaserRelease;;
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

build() {
  local projectVersion="${1}"
  step "Build Launchee ${projectVersion}"
  if [[ "${buildForJdvm:-false}" == "true" ]]; then
    local jdvmAppIcon="../../jdvm/wayland-icon-fix/icons/apps/jdvm.png"
    if [[ ! -f "${jdvmAppIcon}" ]]; then
      echo -e "${ERROR} Missing ${jdvmAppIcon}"
      exit 1
    fi
    cp -v "build/appicon.png" "build/appicon.png.bak"
    cp -v "${jdvmAppIcon}" "build/appicon.png"
  fi
  run wails build -clean \
    -tags "webkit2_41" \
    -ldflags "-X github.com/jdheim/launchee/cmd.appVersion=${projectVersion} \
              -X github.com/jdheim/launchee/cmd.buildForJdvm=${buildForJdvm:-false}"
  if [[ "${buildForJdvm:-false}" == "true" ]]; then
    mv -v "build/appicon.png.bak" "build/appicon.png"
    cp -v "build/bin/launchee" "../../jdvm/src/jdvm-docker/src/main/docker/scripts/launchee"
  fi
}

main "$@"
