# Changelog

## [0.12.0](https://github.com/goexts/generic/compare/v0.11.0...v0.12.0) (2025-11-21)


### Features

* **maps:** add FirstKey and FirstKeyBy functions with comprehensive tests and examples ([2e425c2](https://github.com/goexts/generic/commit/2e425c230523704905260168b3da0d9fcdbac482))
* **maps:** add FirstValue and FirstValueBy functions with comprehensive tests and examples ([6479945](https://github.com/goexts/generic/commit/6479945907ba055f8eec5aea3a56ffe12e43a872))
* **maps:** add Random, RandomKey, RandomValue and FirstKeyOrRandom, FirstValueOrRandom functions with tests and examples ([ec6f3e8](https://github.com/goexts/generic/commit/ec6f3e831c62a7d5056900602f1d15955605b211))
* **maps:** improve random functions test stability by adding predictable outputs in examples ([e3bb3b0](https://github.com/goexts/generic/commit/e3bb3b0208aeac0938b792617b8ceb86cfe28c0d))
* **res:** add generic Pair type with utility methods and comprehensive tests ([f3b6616](https://github.com/goexts/generic/commit/f3b66168296b52c6ba4280a3cc07ece288cdc981))

## [0.11.0](https://github.com/goexts/generic/compare/v0.10.0...v0.11.0) (2025-11-06)


### Features

* **promise:** add All and Then helper functions and improve error handling ([8dddb64](https://github.com/goexts/generic/commit/8dddb64386f43a89005d85398e9eaf6b9a3317fb))

## [0.10.0](https://github.com/goexts/generic/compare/v0.9.0...v0.10.0) (2025-10-31)


### Features

* **maps:** Add KV and KVs convenient functions and simplify type inference ([9569432](https://github.com/goexts/generic/commit/956943284d50d2feac4d01ccd0aa3f5e7f10d974))

## [0.9.0](https://github.com/goexts/generic/compare/v0.8.0...v0.9.0) (2025-10-30)


### Features

* **examples:** add String method to Product struct for better string representation ([305428a](https://github.com/goexts/generic/commit/305428a014e8f319a5858a2cc0c5cb81944bcb66))

## [0.8.0](https://github.com/goexts/generic/compare/v0.7.0...v0.8.0) (2025-10-20)


### Features

* **maps:** add ToSliceWith and FromSliceWithIndex functions with tests for filtered map conversions ([d958f6f](https://github.com/goexts/generic/commit/d958f6f68a4ea9380389c0836485643b0a1271c6))

## [0.7.0](https://github.com/goexts/generic/compare/v0.6.0...v0.7.0) (2025-10-20)


### Features

* **maps:** improve Filter function logic and refactor tests for better clarity and coverage ([7066ec0](https://github.com/goexts/generic/commit/7066ec0da392263500ee8335134fe10de8ba2d8e))
* **maps:** remove outdated comment in TestToKVs test function ([65958d6](https://github.com/goexts/generic/commit/65958d66be69f938a8d42777b4f09e5800a27d90))
* **maps:** remove unused sortKV function and clean up test function parameters ([8063b28](https://github.com/goexts/generic/commit/8063b2894f1214d200319bc16e34b9888bb1da57))
* **maps:** rename and restructure map utility functions (Filter-&gt;Exclude, KVsToMap-&gt;FromKVs, etc.) and add Concat/ConcatWith functions ([c6f3fd6](https://github.com/goexts/generic/commit/c6f3fd6229c3b1cd53fbd59d7073c49ccaeaf50a))
* **maps:** rename functions to more concise names (MergeFunc-&gt;MergeWith, MergeMaps-&gt;Concat, etc.) and update tests ([7f7c859](https://github.com/goexts/generic/commit/7f7c8596d042031e1f78ae1892bd6db09ad03da4))

## [0.6.0](https://github.com/goexts/generic/compare/v0.5.0...v0.6.0) (2025-10-11)


### Features

* **docs:** add slices_example.go and update API documentation with examples and improved descriptions ([c652a63](https://github.com/goexts/generic/commit/c652a635b7b82ef8753def09e408cf745e594b40))
* **slices:** add FilterIncluded and FilterExcluded functions with benchmarks and tests ([495c227](https://github.com/goexts/generic/commit/495c227bb74d202da6a8d48545501d59615f9b00))

## [0.5.0](https://github.com/goexts/generic/compare/v0.4.0...v0.5.0) (2025-09-28)


### Features

* **configure:** add OptionSet and OptionSetE constructors and move tests to apply_test.go ([e9195a3](https://github.com/goexts/generic/commit/e9195a35396a56562bb84e73d1e5ca132667a86e))
* **configure:** add type-safe constructors New, NewWith, NewE, NewWithE and rename New to NewAny ([e8e6b16](https://github.com/goexts/generic/commit/e8e6b1678c8c09285a9d9a08f608b3e437cb806d))


### Bug Fixes

* **configure:** correct spelling of "Endeavor" in test case ([3ce7744](https://github.com/goexts/generic/commit/3ce77445f008eda8d877e90fc839c53233127d7a))

## [0.4.0](https://github.com/goexts/generic/compare/v0.3.0...v0.4.0) (2025-09-05)


### Features

* **cmp:** add comparison functions and improve existing ones ([a7d9ac9](https://github.com/goexts/generic/commit/a7d9ac92bb8476c8a12c5905a082612738d946fe))
* **configure:** enhance option handling with reflection and new utilities ([ab1d208](https://github.com/goexts/generic/commit/ab1d208f7942c5d3f7240ed9386e5b5fd5912de1))
* **promise:** implement Promise type and async functionality ([aaed4ff](https://github.com/goexts/generic/commit/aaed4ffd4915ff92c774596a2757a4fec0b80992))
* **promise:** implement Promise type and async functionality ([6367f35](https://github.com/goexts/generic/commit/6367f359ecae84c12327c1d618e50be79a89c503))
* **slices:** generate adapter code for bytes and runes ([e1518b3](https://github.com/goexts/generic/commit/e1518b300839f94c118b1d452158e430b0861b7a))

## [0.3.0](https://github.com/goexts/generic/compare/v0.2.6...v0.3.0) (2025-08-29)


### Features

* **cmp:** add comparison functions and improve existing ones ([a7d9ac9](https://github.com/goexts/generic/commit/a7d9ac92bb8476c8a12c5905a082612738d946fe))
* **generic:** add ChooseFunc utility ([a1666f4](https://github.com/goexts/generic/commit/a1666f4ba72c73dae6e54c0e76edd62ab46d319a))
* **promise:** implement Promise type and async functionality ([aaed4ff](https://github.com/goexts/generic/commit/aaed4ffd4915ff92c774596a2757a4fec0b80992))
* **promise:** implement Promise type and async functionality ([6367f35](https://github.com/goexts/generic/commit/6367f359ecae84c12327c1d618e50be79a89c503))
* **slices:** generate adapter code for bytes and runes ([e1518b3](https://github.com/goexts/generic/commit/e1518b300839f94c118b1d452158e430b0861b7a))
