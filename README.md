# go-boringwoz

Package **boringwoz** provides a way of creating `"boring_wozniak"` style human-friendly **random-names**, for the Go programming language.

_"boring wozniak"_ style **random-names** are of the following form:

	<ADJECTIVE> "_" <SURNAME>

For example:

* `admiring_agnesi`
* `zen_zhukovsky`
* `nostalgic_kang`
* `kind_edison`
* `zealous_cori`
* `great_franklin`
* `goofy_jang`
* `youthful_jemison`
* `zealous_tafazoli`
* `fervent_arbabian`
* `reverent_holder`
* `beautiful_lamport`
* `serene_hawking`
* `jolly_gould`
* `gifted_eliasi`
* `sleepy_goldwasser`
* `adoring_albattani`
* `pedantic_kazerouni`
* `priceless_chen`

`"boring_wozniak"` style **random-names** are used as an alternative to, for example, naming things using UUIDs.
It is thought that, `"boring_wozniak"` style **random-names** are easier to recognize (for humans) than, for example, UUIDs.

This package provides both the classic algorithm (with 108 adjectives and 149 surnames), and a new algorithm (with more adjectives and more surnames).

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-boringwoz

[![GoDoc](https://godoc.org/github.com/reiver/go-boringwoz?status.svg)](https://godoc.org/github.com/reiver/go-boringwoz)

## Examples

Here is an example of generating a **random-name**:

```go
import "github.com/reiver/go-boringwoz"

//

name boringwoz.RandomName()
```

And, if you want to use the classic algorith, you can do with with the following:

```go
import "github.com/reiver/go-boringwoz"

//

name boringwoz.RandomNameClassic()
```

## Import

To import package **boringwoz** use `import` code like the following:
```
import "github.com/reiver/go-boringwoz"
```

## Installation

To install package **boringwoz** do the following:
```
GOPROXY=direct go get github.com/reiver/go-boringwoz
```

## Author

Package **boringwoz** was written by [Charles Iliya Krempeaux](http://reiver.link)
