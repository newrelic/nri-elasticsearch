# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

Unreleased section should follow [Release Toolkit](https://github.com/newrelic/release-toolkit#render-markdown-and-update-markdown)

## Unreleased

## v5.4.2 - 2025-07-29

### ⛓️ Dependencies
- Updated golang patch version to v1.24.5

## v5.4.1 - 2025-07-01

### ⛓️ Dependencies
- Updated golang version to v1.24.4

## v5.4.0 - 2025-01-21

### ⛓️ Dependencies
- Updated golang patch version to v1.23.5

### 🚀 Enhancements
- Add FIPS compliant packages

## v5.3.0 - 2024-10-08

### dependency
- Upgrade go to 1.23.2

### 🚀 Enhancements
- Upgrade integrations SDK so the interval is variable and allows intervals up to 5 minutes

## v5.2.6 - 2024-07-09

### ⛓️ Dependencies
- Updated golang version to v1.22.5

## v5.2.5 - 2024-05-14

### ⛓️ Dependencies
- Updated golang version to v1.22.3

## v5.2.4 - 2024-04-16

### ⛓️ Dependencies
- Updated golang version to v1.22.2

## v5.2.3 - 2024-02-27

### ⛓️ Dependencies
- Updated github.com/newrelic/infra-integrations-sdk to v3.8.2+incompatible

## v5.2.2 - 2023-10-31

### ⛓️ Dependencies
- Updated golang version to 1.21

## v5.2.1 - 2023-08-08

### ⛓️ Dependencies
- Updated golang to v1.20.7

## v5.2.0 - 2023-07-25

### 🚀 Enhancements
- bumped golang version pinning 1.20.6

## 5.1.0 (2023-06-06)
### Changed
- Upgrade Go version to 1.20

## 5.0.0 (2023-01-19)
### Removed
- Dropped support for Elastic v5 and v6 (both EOL)
- Removed metrics since these were deprecated from v7:
  - "threadpool.indexActive"
  - "threadpool.indexQueue"
  - "threadpool.indexRejected"
  - "threadpool.indexThreads" 
### Added
Added support for:
- Elasticsearch v7
- Elasticsearch v8

## 4.5.3 (2022-06-21)
### Changed
- Bump dependencies
### Added
Added support for more distributions:
- RHEL(EL) 9
- Ubuntu 22.04
- Amazon Linux 2022


## 4.5.2 (2022-05-02)
### Changed
- Bump dependencies
### Added
* Added elasticsearch logging file example by @marcelschlapfer in https://github.com/newrelic/nri-elasticsearch/pull/109


## 4.5.1 (2021-10-20)
### Added
Added support for more distributions:
- Debian 11
- Ubuntu 20.10
- Ubuntu 21.04
- SUSE 12.15
- SUSE 15.1
- SUSE 15.2
- SUSE 15.3
- Oracle Linux 7
- Oracle Linux 8

## 4.5.0 (2021-08-27)
### Added
Moved default config.sample to [V4](https://docs.newrelic.com/docs/create-integrations/infrastructure-integrations-sdk/specifications/host-integrations-newer-configuration-format/), added a dependency for infra-agent version 1.20.0

Please notice that old [V3](https://docs.newrelic.com/docs/create-integrations/infrastructure-integrations-sdk/specifications/host-integrations-standard-configuration-format/) configuration format is deprecated, but still supported.

## 4.4.0 (2021-08-05)
### Changed
- Updated the SDK
### Added
- Adde new flag tls_unsecure_skip_verify

## 4.3.6 (2021-06-09)
### Changed
- Support for ARM

## 4.3.5 (2021-04-20)
### Added
- option `master_only` affection `command: all` only: Collect cluster metrics on the elected master only
- Upgraded github.com/newrelic/infra-integrations-sdk to v3.6.7
- Switched to go modules
- Upgraded pipeline to go 1.16
- Replaced gometalinter with golangci-lint

## 4.3.4 (2021-03-23)
### Changed
- Add arm packages and binaries.

## 4.3.3 (2020-10-19)
### Changed
- Removed the index limit

## 4.3.0 (2020-04-27)
### Changed
- Updated the SDK to the latest version
### Added
- `ssl_alternative_hostname` argument to work around invalid hostname issues

## 4.2.0 (2019-11-18)
### Changed
- Renamed the integration executable from nr-elasticsearch to nri-elasticsearch in order to be consistent with the package naming. **Important Note:** if you have any security module rules (eg. SELinux), alerts or automation that depends on the name of this binary, these will have to be updated.

## 4.1.1 - 2019-10-16
### Fixed
- Windows installer GUIDs

## 4.1.0 - 2019-06-19
### Added
- Windows build

## 4.0.2 - 2019-06-19
### Fixed
- Misspelling "Dcoument" and "Aueue"
- Added missing metric activeSearchesInMilliseconds

## 4.0.1 - 2019-05-20
### Fixed
- Segfault on blank node ingests

## 4.0.0 - 2019-04-22
### Changed
- Prefixed namespace to provide better uniqueness
- Cluster name and environment are now identity attributes
### Added
- Added a reporting endpoint to the entities


## 3.0.0 - 2019-02-04
### Changed
- Updated definition version
- Changed nodes to use their name rather than their host as displayName

## 2.0.0 - 2019-01-09
### Changed
- Updated some metric names and fixed some descriptions in the spec.csv

## 1.2.0 - 2018-11-29
### Added
- Cluster name is now attached to allow entities within a cluster
### Changed
- Nodes are no identified by hostname/IP rather than ID

## 1.1.0 - 2018-11-28
### Added
- ClusterEnvironment argument to help further identify clusters
### Fixed
- Issue when loading a config yaml file that had nested objects. Theses objects are ignored for inventory data.

## 1.0.1 - 2018-11-20
### Changed
- Added host and IP to Nodes as additional attributes

## 1.0.0 - 2018-11-16
### Changed
- Updated to version 1.0.0

## 0.2.2 - 2018-11-07
### Changed
- Increased index limit to 500

## 0.2.1 - 2018-10-23
### Added
- Additional descriptions for config parameters
### Fixed
- Incorrect YAML in config file

## 0.2.0 - 2018-10-04
### Added
- Added index regex configuration parameter to allow limiting which indices are collected
- Added a hard limit on the number of indices to collect (100)

## 0.1.3 - 2018-09-25
### Added
- Added local hostname argument to allow for overriding "localhost" as the host from which to collect inventory data.

## 0.1.2 - 2018-09-14
### Changed
- Removed IP field from Node struct. It was not required as part of collection and could cause an error as the value could be a single string or an array of strings.

## 0.1.1 - 2018-09-13
### Added
- Implemented client authentication
- Implemented toggles for primaries and indices
- Added status code and error checking to client requests

## 0.1.0 - 2018-08-28
### Added
- Initial version: Includes Metrics and Inventory data
