jsonfmt
=======

JSON formatter utility
-------

`jsonfmt` takes a minified JSON file -- for example, a large REST response saved to a file -- and outputs a formatted, readable version to stdout, where it can be piped to a file, grepped, etc.

```
$cat example.json
{"foo":"bar"}
$jsonfmt example.json
{
    "foo": "bar"
}
```
