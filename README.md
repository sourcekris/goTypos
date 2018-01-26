# goTypos
A library to generate typographical errors for an abitrary string.

## Types

```
type typos struct {
  Original string
  Typos []string
}
```

## Usage

```
import "github.com/sourcekris/goTypos"

t := typos.NewTypos("someword")
t.ReverseLetter()
fmt.Println(t.Typos)
```

## Background
This library uses methods ported from Python and originally featured in:
* https://github.com/ncouture/TypoGenerator