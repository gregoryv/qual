# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).


## [Unreleased]
### Changed

- Assert() removed, resulted in Not so readable output

## [0.2.1] 2018-12-15
### Added

- Print calculated fix duration and estimated done date

### Fixed

- Assert() generates generic messages when source code is not available

## [0.2.0] - 2018-08-19
### Added

- Assert() func for readable assertions

### Changed

- LineLength takes tabSize to take visual width into account

### Fixed

- Failed output referred to qual test, missing t.Helper

## [0.1.0] - 2018-04-21
### Added

- High and Standard test helpers for quick assessment
