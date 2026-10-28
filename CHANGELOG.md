# Changelog

## [1.8.0](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.7.2...v1.8.0) (2026-10-28)


### Features

* Autocreate missing PTRs, handle notFound errors gracefully ([2949dd5](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/2949dd56e725e1fdff798a54bcbeb284d8e61dc9))
* **config:** allow configuring the WAPI base path ([2cfe3f3](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/2cfe3f3f300fbb60d9a50b96fc939c782bce68d1)), closes [#18](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/18)
* **infoblox:** add AAAA (IPv6) record support ([13c9812](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/13c98126a06b2f60d9f3c31615fd11f7e206e65e))
* **infoblox:** create PTR records for AAAA records ([7b77eb5](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/7b77eb50d508078dbce1df5a4f56f7a2bbfe5ffd))
* Reduce amount of requests to Infoblox API ([8c7c1eb](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/8c7c1eb3b26b46e5cf7d28d7f43f65c12a9dd098))
* replace caarlos0/env by alecthomas/kong, allow arguments as configuration ([#77](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/77)) ([ffc4444](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/ffc4444b8fae769d51d19b527e47a8faeb3eadbc))


### Bug Fixes

* bump dependencies ([#78](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/78)) ([2f02eed](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/2f02eed6066c9f960566398b3793297e4d8274ad))
* **deps:** update module github.com/miekg/dns to v1.1.73 ([42591be](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/42591beb184ba51f40c94cdfd872663b93d9bfad))
* **deps:** update module github.com/prometheus/client_golang to v1.25.0 ([5b14539](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/5b145396d361943f8f1f8b67ba05f3127799904e))
* **deps:** update module github.com/sirupsen/logrus to v1.10.2 ([6945136](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/69451364085f882847ba90e5de7c2682337bad73))
* **deps:** update module github.com/stretchr/testify to v1.12.1 ([a5903de](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/a5903de90ae7f0cb4924a8532d8b69b46ec1bf04))
* **deps:** update module sigs.k8s.io/external-dns to v0.24.0 ([#93](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/93)) ([2e88a4f](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/2e88a4ff146981d1f0db554170a2490aad31e93c))
* **infoblox:** compare IPv6 targets in canonical form ([1b615a4](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/1b615a4e04e9099b18151818dfd3b165ddad5874))
* **infoblox:** select the most specific reverse zone ([#82](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/82)) ([281f471](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/281f47194e26014f43fa974f63f3e588ade08c50))
* **infoblox:** skip TXT records colliding with a CNAME ([a21fcd8](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/a21fcd8413eb5886bd67483504bd590dfedfdd29)), closes [#23](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/23)
* **infoblox:** surface WAPI auth and timeout failures in the log ([de59be8](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/de59be83f5cafa5ff82dd3961659b47974a8cfdb)), closes [#50](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/50) [#19](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/19)
* **infoblox:** update existing CNAME instead of create plus delete ([e75d053](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/e75d053de3b520f94e17298f67867b5f4af92b33)), closes [#63](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/63)


### Miscellaneous Chores

* **deps:** update actions/checkout action to v7 ([59734fb](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/59734fb01e034afddb434207d196756566ee552c))
* **deps:** update actions/setup-go action to v7 ([dbc240d](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/dbc240dad8a330843ce9553f0aa3d7106878cec6))
* **deps:** update actions/setup-python action to v7 ([8f26967](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/8f2696758ad0bb43d45d022b9457004a1bd6e6b0))
* **deps:** update anchore/sbom-action action to v0.24.3 ([da5fe20](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/da5fe20b59e177e41a82118c2444cf073ecc73c8))
* **deps:** update dependency golangci/golangci-lint to v2.15.0 ([14b0ad7](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/14b0ad7cb690c0d3fabac9ad771d96dc7cf7fe80))
* **deps:** update docker/login-action action to v4 ([ce0f9e0](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/ce0f9e03be89cb5edcb3fad56f588285ad515ca7))
* **deps:** update docker/setup-buildx-action action to v4 ([7e35942](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/7e35942b72247b5b04cab6d41ba311cd051a4907))
* **deps:** update docker/setup-qemu-action action to v4 ([e181ae6](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/e181ae693e88d3c3e197fb29d9c75fdab9a7a5c5))
* **deps:** update golangci/golangci-lint-action action to v9 ([4a5164d](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/4a5164d5b9b0cca9b96c20db67c5c4a6c9562f09))
* **deps:** update googleapis/release-please-action action to v5 ([7ff0683](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/7ff0683b57cf37b82bd479c58382e2af58dd5314))
* **deps:** update goreleaser/goreleaser-action action to v7 ([5df25b8](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/5df25b8381515ef67c7b47c185cf3abd85ea8844))
* **deps:** update sigstore/cosign-installer action to v3.10.1 ([1796a13](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/1796a13b5b02a2f2ef5a963e85cb718dee884834))
* **deps:** update sigstore/cosign-installer action to v4 ([a1ada73](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/a1ada73c4f9b531431201de1149baaa71b6797fa))

## [1.7.2](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.7.1...v1.7.2) (2026-07-21)


### Miscellaneous Chores

* **deps:** bump golang.org/x/net from 0.28.0 to 0.38.0 ([9055e21](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/9055e21956d640cffe52e867985d70a883efe5d3))

## [1.7.1](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.7.0...v1.7.1) (2026-03-04)


### Miscellaneous Chores

* upgrade go to 1.26.0 ([#58](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/58)) ([6783698](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/6783698ee8f047d4db681e97acb0caebddd97e14))
* upgrade golangci-lint version ([21be821](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/21be821735ab2e8a7f677914f3d4ed1fc726ba9b))

## [1.7.0](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.6.0...v1.7.0) (2025-10-29)


### Features

* make use_ttl configurable via INFOBLOX_USE_TTL environment variable ([#56](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/56)) ([f6b8704](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/f6b8704b020f118ddbc901931c900cbad5dda08f))

## [1.6.0](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.5.0...v1.6.0) (2025-09-30)


### Features

* Add metrics for infoblox API calls ([a533e11](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/a533e112c3148f351fdc522af8ddf3296749a3d6))

## [1.5.0](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.4.4...v1.5.0) (2025-08-21)


### Features

* add NS record support ([#49](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/49)) ([f15d8b4](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/f15d8b403f9b0fea14750718aba8975016898f51))

## [1.4.4](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.4.3...v1.4.4) (2025-05-13)


### Bug Fixes

* Update README with new defaults ([a05beef](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/a05beefdfff9fb76b551680cf57e35f437ff248f))

## [1.4.3](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.4.2...v1.4.3) (2025-04-01)


### Bug Fixes

* continue on buildRecord error ([9f376ad](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/9f376ade7346f8c5d1fa508cbdc1117f3b1e673c))
* show errors in webhook ([323c425](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/323c4257a19a4035c9b6c8daf45102bb7eb9383a))

## [1.4.2](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.4.1...v1.4.2) (2024-12-29)


### Bug Fixes

* ptr search query ([#32](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/32)) ([bbcb30d](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/bbcb30d70b902506d15f3dc42b1f7daa3eb02739))

## [1.4.1](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.4.0...v1.4.1) (2024-12-14)


### Bug Fixes

* revert repo-name ([#30](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/30)) ([6454728](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/64547289e9505f2cbece2a2c95b7f9aa1a0a25e0))

## [1.4.0](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.3.0...v1.4.0) (2024-12-12)


### Features

* Add support for extensible attributes ([#26](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/26)) ([27b345d](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/27b345d7e7453d3bf9ab2d4b21627770e3ee84b5))


### Miscellaneous Chores

* **release:** release 1.3.0 ([#27](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/27)) ([dc845fe](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/dc845fe1af694868f1c7a31629236e3c15a92fa2))

## [1.3.0](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.2.0...v1.3.0) (2024-12-12)


### Features

* Add support for extensible attributes ([#26](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/26)) ([27b345d](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/27b345d7e7453d3bf9ab2d4b21627770e3ee84b5))

## [1.2.0](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.1.1...v1.2.0) (2024-11-06)


### Features

* Implement support for PTR records ([#22](https://github.com/AbsaOSS/external-dns-infoblox-webhook/issues/22)) ([ca1cbaf](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/ca1cbafb48379eecee06294f050f6f5cd2fec0cc))

## [1.1.1](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.1.0...v1.1.1) (2024-06-06)


### Bug Fixes

* test release ([350ffd0](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/350ffd0d05e6f8160cca72eb575013099c667e42))

## [1.1.0](https://github.com/AbsaOSS/external-dns-infoblox-webhook/compare/v1.0.0...v1.1.0) (2024-06-05)


### Features

* Infoblox paging support ([f457da8](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/f457da8a091c062b6c264157aca94519b2805eb8))

## 1.0.0 (2024-05-28)


### ⚠ BREAKING CHANGES

* Initial commit

### Features

* Initial commit ([5629414](https://github.com/AbsaOSS/external-dns-infoblox-webhook/commit/562941456f5a60be6c98de55aa3e5f54edf73e2f))
