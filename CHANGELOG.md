# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.0.3] - 2023-11-17

* Update dependencies.

## [0.0.3] - 2023-10-21

* **Introduce `filesystem.CreateFile` function.**
* **Introduce `filesystem.WriteContentTo` function.**
* **Introduce `filesystem.StreamContentTo` function.**
* **Introduce `filesystem.CreateDirectory` function.**
* Accept argument in certain functions to overwrite default behavior.
    * Permission mode.
    * Creating non-existing directory structure when creating new file.
    * Overwriting existing files.

## [0.0.2] - 2023-10-12

* **Introduce `filesystem.CheckIfExists` function.**
* Add `YAMLDecode` function to `filesystem.Content` type.
* Replace `Streamer` interface with `io.ReadCloser`.

## [0.0.1] - 2023-10-10

* Release to GitHub

[0.0.4]: https://github.com/SevenOfSpades/go-wrapped-filesystem/releases/tag/v0.0.4
[0.0.3]: https://github.com/SevenOfSpades/go-wrapped-filesystem/releases/tag/v0.0.3
[0.0.2]: https://github.com/SevenOfSpades/go-wrapped-filesystem/releases/tag/v0.0.2
[0.0.1]: https://github.com/SevenOfSpades/go-wrapped-filesystem/releases/tag/v0.0.1