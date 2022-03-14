Strip JSON Comments
-------------------
Strips single line `//` and block `/* */` comments from .jsonc text:

```
echo '[1, /* strip this comment */ 2, 3]' | ./strip-jsonc 
[1,                          2, 3]
```

### Comment characters are replaced with spaces
To preserve line and column locations of the remaining JSON text.

### Block comments do not nest
JSONC doesn't have a spec but its block comments have been [described](https://code.visualstudio.com/docs/languages/json#_json-with-comments) as behaving like JavaScript, i.e. they do not nest.

### Trailing commas are not removed
Some JSONC parsers support trailing commas, but as `strip-jsonc` is line-oriented, it does not have an option to remove trailing commas.

### Test input
This repo contains a test input file called `test.jsonc` which exercises the different cases of embedding comments in JSON.

### Perl
The `strip-jsonc` script is written in Perl and designed to be fast and compatible with any version of Perl 5.

### Go
There is a golang variant of the strip-jsonc script which has the same behavior as the Perl script but can be 3-4x faster for large inputs. It can be compiled with go build:

```
go build gostrip-jsonc.go
```
