package typos

import (
  "strings"
)

const (
  alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
  vowels = "aeiouy"
)

type typos struct {
  Original string
  Typos []string
}

func NewTypos(w string) *typos {
  return &typos{
    Original: w,
  }
}

func (t *typos) AddTypo(s string) {
  if s == t.Original {
    return
  }

  for _, b := range t.Typos {
    if s == b {
      return
    }
  }

  t.Typos = append(t.Typos, s)
}

func (t *typos) InsertedKey() {
  for i := 0; i < len(t.Original); i++ {
    for _, r := range alphabet {
      t.AddTypo(t.Original[:i+1] + string(r) + t.Original[i+1:])
    }
  }
}

func (t *typos) SkipLetter() {
  for i := 1; i < len(t.Original)+1; i++ {
    t.AddTypo(t.Original[:i-1] + t.Original[i:])
  }
}

func (t *typos) DoubleLetter() {
  for i := 1; i < len(t.Original)+1; i++ {
    t.AddTypo(t.Original[:i] + string(t.Original[i-1]) + t.Original[i:])
  }
}

func (t *typos) ReverseLetter() {
  for i := 0; i < len(t.Original); i+=2 {
    if len(t.Original[i:]) >= 2 {
      t.AddTypo(t.Original[:i] + string(t.Original[i+1]) + 
                string(t.Original[i]) + t.Original[i+2:])
    }
  }
}

func (t *typos) WrongVowel() {
  for i := 0; i < len(t.Original); i++ {
    if strings.ContainsAny(vowels, string(t.Original[i])) {
      for _, y := range vowels {
        wordbytes := []byte(t.Original)
        wordbytes[i] = byte(y)
        t.AddTypo(string(wordbytes))
      }
    }
  }
}

func (t *typos) WrongKey() {
  for i := 0; i < len(t.Original); i++ { 
    for _, r := range alphabet {
      t.AddTypo(t.Original[:i] + string(r) + t.Original[i+1:])
    }
  }
}

func (t *typos) AllTypos() {
  t.InsertedKey()
  t.SkipLetter()
  t.DoubleLetter()
  t.ReverseLetter()
  t.WrongVowel()
  t.WrongKey()
}