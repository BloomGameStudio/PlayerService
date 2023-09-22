# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added 

- Write Error Handler
- Error Close Error Handler

- Level Model & Websocket
  - https://github.com/BloomGameStudio/PlayerService/issues/33 
  - This implements the `/level` websocket endpoint with full CRUD functionality for "Level".
    and adds "Level" read functionality to the `/player` websocket endpoint.
   
  - Implemented & Added:
    - Level Websocket Controller
    - Level Read Functionality in the `/player` Socket Endpoint
    - Level Reader
    - Level Writer
    - Level Handler
    - Level Model
    - Level Public Model
    - Level Model in Player Model
    - Level Model in Public Player Model
  
  - Note: This does not Implement "Level" Write Functionality in the `/player` Websocket Endpoint.

- Scale Websocket
  - https://github.com/BloomGameStudio/PlayerService/issues/46
  - This implements the `/scale` websocket endpoints supporting read/write operations.

### Changed

-   playerWriter Error Handling
-   Websocket timeout duration moved to the `WS_TIMEOUT_SECONDS` environment variable

## [0.1.0] - 16-09-2032 

### Added

- This CHANGELOG file.


[unreleased]: https://github.com/BloomGameStudio/PlayerService/compare/staging...dev
[0.1.0]: https://github.com/BloomGameStudio/PlayerService/releases/tag/0.1.0