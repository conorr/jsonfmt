jsonfmt
=======

A formatting utility for JSON
----

`jsonfmt` is a formatting utility for files containing JSON data.

Given a file, `jsonfmt` formats the JSON data and writes it to stdout, where it can be piped to a pager or directed into a file.

```
$ cat example.json
{"veggies":["lettuce","carrots","celery"],"fruits":["apple","orange","banana"]}
$ jsonfmt example.json
{
    "veggies": [
        "lettuce",
        "carrots",
        "celery"
    ],
    "fruits": [
        "apple",
        "orange",
        "banana"
    ]
}
```

`jsonfmt` automatically detects and handles JSON data wrapped in a JavaScript callback (JSONP). This is useful if you have a large JSONP response you've saved off in a file:

```
$ cat jsonp_example.js
SOME.CALLBACK({"veggies":["lettuce","carrots","celery"],"fruits":["apple","orange","banana"]})
$ jsonfmt jsonp_example.js
SOME.CALLBACK({
    "veggies": [
        "lettuce",
        "carrots",
        "celery"
    ],
    "fruits": [
        "apple",
        "orange",
        "banana"
    ]
})
```

#### Options

The `--replace`/`-r` option can be used to format the file in-place by replacing it with its formatted contents. Additionally, the `--sort`/`-s` flag can be used to recursively sort all keys alphabetically:

```
$ jsonfmt --sort example.json
{
    "fruits": [
        "apple",
        "banana",
        "orange"
    ],
    "veggies": [
        "carrots",
        "celery",
        "lettuce"
    ]
}
```
