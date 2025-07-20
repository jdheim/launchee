#!/usr/bin/env bash
# STARTS LAUNCHEE IN DEV MODE

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

Starts Launchee in DEV mode

OPTIONS:
  -d                     Starts Docs in DEV mode
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
  start "${projectVersion}"
}

readOptions() {
  while getopts ":hd" option; do
    case "${option}" in
      d) startDocs;;
      h|?) usage ;;
    esac
  done
}

startDocs() {
  step "Start Docs in DEV mode"
  cd docs
  pnpm start
  exit $?
}

start() {
  local projectVersion="${1}"
  step "Start Launchee ${projectVersion} in DEV mode"
  run wails dev -extensions "go,yml" \
    -tags "webkit2_41" \
    -ldflags "-X github.com/jdheim/launchee/cmd.appVersion=${projectVersion}"
}

main "$@"
