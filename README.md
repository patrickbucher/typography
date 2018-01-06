# Typography

This is a module that replaces some characters and sequences of characters often
to be seen in text files with their typographycally correct counterparts. Right
now it's replacing:

- three dots (...) with an ellipsis (…)
- two dashes (--) with an ndash (–)
- three dashes (---) with an mdash (—)
- double quotes ("") with nice double quotes
- single quotes ('') with nice single quotes
- single quotes within words (don't) with an apostrophe (don’t)

## Quote Styles

Four quote styles (`typography.QuoteStyle`) are supported:

- English style: “” (double) and ‘’ (single)
    * `typography.QuoteStyle.English`
- German style: „“ (double) and ‚‘ (single)
    * `typography.QuoteStyle.German`
- Guillemets: «» (double) and ‹› (single)
    * `typography.QuoteStyle.Guillemets`
- Reverse Guillemets: »« (double) and ›‹ (single)
    * `typography.QuoteStyle.ReverseGuillemets`

## Client

I implemented a client that takes text from `os.Stdin`, beautifies it, and
writes it to `os.Stdout`. It can be found in the `cmd` folder (`beautify.go`).
It supports the quote styles mentioned above using the flags `-e` (English
style), `-d` (German style), `-g` (Guillemets) and `-r` (reverse Guillemets).
