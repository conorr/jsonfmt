jsonfmt
=======

A formatting utility for JSON/JSONP
----

`jsonfmt` is a formatting utility for files containing JSON or JSONP data. Given a file, jsonfmt will 

Given a file of JSON or JSONP, jsonfmt prints the indented form to stdout, where it can be piped into pager or a file.

```
$ cat example.json
{"fruits":["apple","orange","banana"],"veggies":["lettuce","carrots","celery"]}
$ jsonfmt example.json
{
    "fruits": [
        "apple",
        "orange",
        "banana"
    ],
    "veggies": [
        "lettuce",
        "carrots",
        "celery"
    ]
}
```

If passed the `--sort`/`-s` flag, jsonfmt will sort keys alphabetically:

```
$ jsonfmt --sort example.json
{
    "fruits": [
        "apple",
        "orange",
        "banana"
    ],
    "veggies": [
        "lettuce",
        "carrots",
        "celery"
    ]
}
```

If passed the `--replace`/`-r` flag, jsonfmt will overwrite the source file with its formatted version.

jsonfmt automatically detects and handles JSONP. For example:

```
$ cat example.js
SOME.CALLBACK({"apples":true,"oranges":true,"pineapples":false})
$ jsonfmt example.js
$ cat example.js
SOME.CALLBACK({
    "apples": true,
    "oranges": true,
    "pineapples": false
})
```
