# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

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