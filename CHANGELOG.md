# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

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