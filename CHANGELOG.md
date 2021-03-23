# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

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
