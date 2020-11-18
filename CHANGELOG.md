# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- The project structure to support more than one version of a framework
  - moved resty v2 middleware to [here](https://github.com/Wr4thon/requestid/tree/v1.0.0/resty/v2)
  - moved atreugo v10 middleware to [here](https://github.com/Wr4thon/requestid/tree/v1.0.0/atreugo/v10)
  - moved atreugo v11 middleware to [here](https://github.com/Wr4thon/requestid/tree/v1.0.0/atreugo/v11)
  - moved echo v4 middleware to [here](https://github.com/Wr4thon/requestid/tree/v1.0.0/echo/v4)

### Added

- Added support for Atreugo v10
- Added support for Atreugo v11

## [0.3.0] - 2020-11-16

### Removed

- Remove support for atreugo v10

## [0.2.0] - 2020-11-05

### Added

- Initial support for atreugo v10
- golangci-lint configuration
- LICENSE file (Apache 2.0)
- description and example code
- added AUTHORS file
- this changelog

### Changed

- comments on echo handler

## [0.1.0] - 2020-11-05

### Added

- Initial support for resty v2
- Initial support for echo v4

[Unreleased]: https://github.com/Wr4thon/requestid/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/Wr4thon/requestid/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/Wr4thon/requestid/releases/tag/v0.1.0
