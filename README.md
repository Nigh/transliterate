Go Transliteration
===

> A fancy way to map Unicode to ASCII and more.

This repo/fork is based off the work of [alexsergivan](https://github.com/alexsergivan) and the amazing work at [transliterator](https://github.com/alexsergivan/transliterator).

Differences from the original:
* ~50-70% decrease in time/op (speed increase)
* Zero allocations for "creating" a new "replacer" (previously called Transliterator)
* Single allocation during transliteration (using pooled buffers)
* Cleaned up package
* Benchmarks
* New API
* **The datasets ar NOT "safe" for modification and use by multiple clients**
  * To ensure "safety", create deep copies of the transliterate-data and transliterate-lang datasets
  * The default maps are global shared variables

```text
// Benchmarks on The Egg Russian Text

// sudoless/transliteration
Benchmark_Transliterate-8   	   10000	    114579 ns/op	    6529 B/op	       1 allocs/op
21k   allocated objects
66MB  allocated space

// alexsergivan/transliterator
Benchmark_Transliterate-8   	    4774	    251253 ns/op	   51320 B/op	      16 allocs/op
58k   allocated objects
237MB allocated space
```


Transliteration takes Unicode text and converts to ASCII characters.

For now, only these languages have specific transliteration rules:
DE, DA, EO, RU, BG, SV, HU, HR, SL, SR, NB, UK, MK, CA, BS.
For other languages, general ASCII transliteration rules will be applied. Also, this package supports adding custom
transliteration rules for your specific use-case. Please check the examples section below.

Installation
---

```
go get github.com/Nigh/transliterate
```

Examples
---

// TODO: Finish writing examples, also write Go style docs and examples


Simple, default, transliteration

```go

```

Adding new lang overwrites

```go

```

Creating a custom replacer

```go

```
