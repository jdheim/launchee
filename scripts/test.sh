#!/usr/bin/env bash
# STARTS ALL THE TESTS

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

readonly TEST_COVERAGE_FILENAME="go-test-coverage"

usage() {
  cat << EOF
Usage: $(basename "$0") [OPTION]

Starts all the tests

OPTIONS:
  -c                     Generate Test Coverage html report
  -o                     Generate and open Test Coverage html report
EOF
  exit 1
}

main() {
  cd ..
  readOptions "$@"
  cd src
  testAll
}

readOptions() {
  while getopts ":hco" option; do
    case "${option}" in
      c) isGenerateTestCoverageHtmlReport="true";;
      o) isGenerateTestCoverageHtmlReport="true"
        isOpenTestCoverageHtmlReport="true";;
      h|?) usage ;;
    esac
  done
}

testAll() {
  step "Start all the tests"
  go test ./... -cover -coverprofile="${TEST_COVERAGE_FILENAME}.out"
  local testCoverageRawLine testCoverageLine testCoverage
  testCoverageRawLine="$(go tool cover -func="${TEST_COVERAGE_FILENAME}.out" | grep -i "total:")"
  echo "${testCoverageRawLine}" | awk '{print $1, $3}'
  testCoverage=$(echo "${testCoverageRawLine}" | awk '{print $3}')
  ../scripts/common/createDynamicJSONForBadgeWithPercent.sh "${TEST_COVERAGE_FILENAME}" "go" "Test Coverage" "${testCoverage}"
  if [[ "${isGenerateTestCoverageHtmlReport:-false}" == "true" ]]; then
    step "Generate Test Coverage html report"
    go tool cover -html="${TEST_COVERAGE_FILENAME}.out" -o "${TEST_COVERAGE_FILENAME}.html"
    if [[ "${isOpenTestCoverageHtmlReport:-false}" == "true" ]]; then
      step "Open Test Coverage html report"
      xdg-open "${TEST_COVERAGE_FILENAME}.html"
    fi
  fi
}

main "$@"
