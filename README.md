# Typography

This is a module that replaces some characters and sequences of characters often
to be seen in text files with their typographycally correct counterparts. Right
now it's replacing:

- three dots (...) with an ellipsis (…)
- two dashes (--) with an ndash (–)
- three dashes (---) with an mdash (—)
- double quotes ("") with double guillements («»)
- single quotes ('') with single guillemets (‹›)
- single quotes within words (don't) with an apostrophe (don’)

## TODO

- Multiple sets of quotation marks should be supported, and the client must be
  able to choose the set he likes to use.
    - English style: “” and ‘’
    - German style: „“ and ‚‘
    - German book style: »« and ›‹
