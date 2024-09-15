# Changelog

## v1.2.5
* Send the proper DNS rcode as a part of the DNS response/answer

## v1.2.4

### Added
* Query logging is now disabled by default, can optionally be enabled using the `query_logging` config option.

## v1.2.3

### Added
* Added support for macOS (Darwin) AMD64 and ARM64 platforms.
* Added support for the Windows ARM64 platform.

### Fixed
* Re-introduced support for the Windows AMD64 platform, mistakingly not included in previous release.

## v1.2.2

### Added
* Added a workflow for making releases. ([#3](https://github.com/oddmario/simple-dns-server/pull/3))
* Added a workflow for compiling commits and pull requests ([#3](https://github.com/oddmario/simple-dns-server/pull/3))

### Changed
* Various improvements to how Docker builds work. ([#3](https://github.com/oddmario/simple-dns-server/pull/3))
