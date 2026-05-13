# Changelog

## 0.1.0 (2026-05-13)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/pops-maellard/stainless-cli/compare/v0.0.1...v0.1.0)

### Features

* allow `-` as value representing stdin to binary-only file parameters in CLIs ([bd9e5aa](https://github.com/pops-maellard/stainless-cli/commit/bd9e5aae866c06b073999c01a7488236b2f0746b))
* better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url` ([5a1b6b2](https://github.com/pops-maellard/stainless-cli/commit/5a1b6b296a4adc815c3dea1ecedce5fc95ed2043))
* binary-only parameters become CLI flags that take filenames only ([995e18a](https://github.com/pops-maellard/stainless-cli/commit/995e18af84db501716d3fe0cfff9ccabfc0aee49))
* **cli:** add `--raw-output`/`-r` option to print raw (non-JSON) strings ([a08ee81](https://github.com/pops-maellard/stainless-cli/commit/a08ee81c28e3a48c546e0043171acf93585cc955))
* **cli:** alias parameters in data with `x-stainless-cli-data-alias` ([5a050c0](https://github.com/pops-maellard/stainless-cli/commit/5a050c084544c445f8a8b07b9dc118382b294eb7))
* **cli:** send filename and content type when reading input from files ([b3bb874](https://github.com/pops-maellard/stainless-cli/commit/b3bb8741697cb3b989e88ba680c14114cc7f09c8))
* set CLI flag constant values automatically where `x-stainless-const` is set ([9b3e4e6](https://github.com/pops-maellard/stainless-cli/commit/9b3e4e60a07449fbe3b9d8b25b5e478b4aa95661))
* support passing path and query params over stdin ([8f25560](https://github.com/pops-maellard/stainless-cli/commit/8f255603143d547576c3c9e8d6aa6b4b87bfcccf))


### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([7e67223](https://github.com/pops-maellard/stainless-cli/commit/7e67223d2f1605dfdbcdb701a165d7b5dba9e749))
* better support passing client args in any position ([9d310cb](https://github.com/pops-maellard/stainless-cli/commit/9d310cbba0b3bde4daabca4afc19d239234d6f07))
* cli no longer hangs when stdin is attached to a pipe with empty input ([d5aec1b](https://github.com/pops-maellard/stainless-cli/commit/d5aec1bfe7faf6ff8832495ac483af44f82c5f0e))
* **cli:** correctly load zsh autocompletion ([7587e8a](https://github.com/pops-maellard/stainless-cli/commit/7587e8a80e9cbbe005f1ebd6a8779817a10bf32f))
* fall back to main branch if linking fails in CI ([a875ca4](https://github.com/pops-maellard/stainless-cli/commit/a875ca4a50460fbad07b500be66ecf5b366318e1))
* fix for failing to drop invalid module replace in link script ([09bb43c](https://github.com/pops-maellard/stainless-cli/commit/09bb43cf8bc28a1e166454d68306165bd4d58873))
* fix for off-by-one error in pagination logic ([3251314](https://github.com/pops-maellard/stainless-cli/commit/325131433a07ce1214a496bc4517eb75960c24b3))
* fix quoting typo ([0ec9709](https://github.com/pops-maellard/stainless-cli/commit/0ec9709ae9aa1c3b4d72a9b89849fd6cb4160664))
* flags for nullable body scalar fields are strictly typed ([8e2f181](https://github.com/pops-maellard/stainless-cli/commit/8e2f181d0c87a1ca48d785fa70811b4c8edf8132))
* handle empty data set using `--format explore` ([66f35bf](https://github.com/pops-maellard/stainless-cli/commit/66f35bf361e37953bf22c0ec6a5f4c332c7db172))
* improve linking behavior when developing on a branch not in the Go SDK ([889e585](https://github.com/pops-maellard/stainless-cli/commit/889e585cddff904ff87da4ba80f61b42324ecbb8))
* improved workflow for developing on branches ([cbe48ee](https://github.com/pops-maellard/stainless-cli/commit/cbe48ee6c8d8eb50d6f5142cc07510dc918c09d8))
* no longer require an API key when building on production repos ([f3ba29c](https://github.com/pops-maellard/stainless-cli/commit/f3ba29cd7e2f3c192a11fdb78165e78e4804e565))
* only set client options when the corresponding CLI flag or env var is explicitly set ([6fb3342](https://github.com/pops-maellard/stainless-cli/commit/6fb334252697914226c6880c90a3d67b824377ef))
* use `RawJSON` when iterating items with `--format explore` in the CLI ([f3d8ddb](https://github.com/pops-maellard/stainless-cli/commit/f3d8ddb6633d4ed04b0eb310aec9edad9fb9bfa6))


### Chores

* add documentation for ./scripts/link ([31bda43](https://github.com/pops-maellard/stainless-cli/commit/31bda43ed52f68ef7bafa3ba07afd561d41318f3))
* **ci:** skip lint on metadata-only changes ([bbfe937](https://github.com/pops-maellard/stainless-cli/commit/bbfe93749b5c10c61d86da9b62a1041f44033b63))
* **ci:** support manually triggering release workflow ([c443404](https://github.com/pops-maellard/stainless-cli/commit/c4434043a3bd8813deb69be0c9c3938c801b7124))
* **cli:** additional test cases for `ShowJSONIterator` ([23d186d](https://github.com/pops-maellard/stainless-cli/commit/23d186d5b0efc6cc57692bbeb5ed3afd56a3efb7))
* **cli:** fall back to JSON when using default "explore" with non-TTY ([cbe09ca](https://github.com/pops-maellard/stainless-cli/commit/cbe09cacda503dcd41c8e91145c5fa0362281a16))
* **cli:** let `--format raw` be used in conjunction with `--transform` ([e8b6490](https://github.com/pops-maellard/stainless-cli/commit/e8b64907b2e9fdb77fe91eefe17ec7058773e685))
* **cli:** switch long lists of positional args over to param structs ([36ed6f3](https://github.com/pops-maellard/stainless-cli/commit/36ed6f3de141af2e2fbddb3d1905fe3e2b694748))
* **cli:** use `ShowJSONOpts` as argument to `formatJSON` instead of many positionals ([bdd2b96](https://github.com/pops-maellard/stainless-cli/commit/bdd2b96e9a566c9f22598574b88360d6fb6cdab5))
* configure new SDK language ([afc9984](https://github.com/pops-maellard/stainless-cli/commit/afc9984ccfb28ae42cbfdc3e3936b68a7d7b95de))
* configure new SDK language ([9a259f5](https://github.com/pops-maellard/stainless-cli/commit/9a259f5e726a8d609ae9e8233dba6461f15be030))
* configure new SDK language ([c0c5e0d](https://github.com/pops-maellard/stainless-cli/commit/c0c5e0d00d8a243f005f4ea50ec334c3bb6fe90e))
* **internal:** codegen related update ([659a143](https://github.com/pops-maellard/stainless-cli/commit/659a143e2ca0d321fa71516a8d3a2b2ec3923642))
* **internal:** more robust bootstrap script ([1245b81](https://github.com/pops-maellard/stainless-cli/commit/1245b81872bc9004a2835469a229684e98ef6dc4))
* **internal:** tweak CI branches ([bd6ff93](https://github.com/pops-maellard/stainless-cli/commit/bd6ff933806da8b54b3935e0786eefa520612719))
* **internal:** update gitignore ([9d613f0](https://github.com/pops-maellard/stainless-cli/commit/9d613f068592633bfe7da26cce57b0593c7293c5))
* mark all CLI-related tests in Go with `t.Parallel()` ([666acff](https://github.com/pops-maellard/stainless-cli/commit/666acffe140a8f21c72a823b55cfffe9ae6d99bc))
* modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary ([26a1e9c](https://github.com/pops-maellard/stainless-cli/commit/26a1e9c69b7203abfebb83a62d774bf92697dc8c))
* omit full usage information when missing required CLI parameters ([3583848](https://github.com/pops-maellard/stainless-cli/commit/3583848491e2557c8e9694ef596191b77ed8ccae))
* redact api-key headers in debug logs ([37e526a](https://github.com/pops-maellard/stainless-cli/commit/37e526a8a99572aee18367fde4334aa0d0720d42))
* switch some CLI Go tests from `os.Chdir` to `t.Chdir` ([66743be](https://github.com/pops-maellard/stainless-cli/commit/66743bea9efb6627dab9bb50c0091dec57033997))
* update SDK settings ([fcb986a](https://github.com/pops-maellard/stainless-cli/commit/fcb986af99df55168170136f33cda0f73c991fb7))
