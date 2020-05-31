<p align="center">
  <img alt="piglatin logo" src="assets/pig.png" height="150"/>
  <h3 align="center">piglatin</h3>
  <p align="center">Pig Latin translator</p>
</p>

---

`piglatin` is a fast [Pig Latin](https://en.wikipedia.org/wiki/Pig_Latin) translator written in Go. 

## Badges
[![dnozdrin](https://circleci.com/gh/dnozdrin/piglatin.svg?style=shield)](https://circleci.com/gh/dnozdrin/piglatin)
[![License](https://img.shields.io/github/license/dnozdrin/piglatin)](/LICENSE)
[![Release](https://img.shields.io/github/release/dnozdrin/piglatin.svg)](https://github.com/dnozdrin/piglatin/releases/latest)

### The algorithm of work
- Ensures proper capitalization
- Correct upper case and lower case formatting
- Correctly translates "qu" (e.g., ietquay instead of uietqay)
- Differentiates between "Y" as vowel and "Y" as consonant (e.g. yellow = ellowyay and style = ylestay)
- Correctly translates contractions
- Hyphenated words are treated as two words
- Words may consist of alphabetic characters only (A-Z and a-z)
- All punctuation, numerals, symbols and whitespace are not modified

### Supported languages
| Language | Key | Notes |
|---|---|---|
|English|`en`|is used implicitly by default|

### Usage
- Run `go run build` to build the project.
- Run `./piglatin`, enter the text that must be translated to the Pig Latin.
- Enter `:q` on a new line to exit.

Or
- Run `./piglatin < input.txt > output.txt` to translate the text in files.
