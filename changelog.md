# Changelog
This project adheres to semantic versioning.
All notable changes to this project will be documented in this file.

## [0.5.0-dev]

- Add field LineLength.IncludeGenerated, standard test now ignores
  generated files. I.e. those containing "DO NOT EDIT"
- Move func LineLength to type LineLength with method Test

## [0.4.3] 2023-12-15

- Update dependencies

## [0.4.2] 2021-12-27

- Update dependencies

## [0.4.1] 2021-06-10

- Update dependencies

## [0.4.0] 2019-11-05
### Added

- FuncHeight checking max number of lines per function

### [0.3.2] 2019-09-25
### Fixed

- Exclude vendor on all platforms

## [0.3.1] 2019-04-11
### Fixed

- LineLength failures show actual test that failed not this package

## [0.3.0] 2019-04-01
### Changed

- Type T defines only needed funcs for this package
- CyclomaticComplexity logs the total duration to fix
- LineLength only shows which line to trim and by how much
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
