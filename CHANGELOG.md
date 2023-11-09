CHANGELOG
===

This is the CHANGELOG in stable version.

v2.1.0 (2023-11-09)
===

### Updates

* Bump up Go version to 1.21

### Fixes

* Change module path to `/v2`

v2.0.0 (2023-07-29)
===

### ⚠ BREAKING CHANGES
This version is not compatible with v1.0.0.

If your application uses the `DeleteMemberInput.ScreenName` field, you must update your code to use the `DeleteMemberInput.ScreenNameOrEmail` field instead in this version.

### Fixes

* Fix the `DeleteMemberInput.ScreenName` field to `DeleteMemberInput.ScreenNameOrEmail` field. ([858555](https://github.com/michimani/go-esa/commit/88585557b7ff38dfbd32019ce3d7e51411e98ba5))

v1.0.0 (2022-03-27)
====

* release 🚀
