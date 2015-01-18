jsonfmt
=======

A formatting utility for JSON
----

`jsonfmt` is a formatting utility for files containing JSON data.

Given a file containing JSON or JSONP, `jsonfmt` formats it and writes it to stdout.

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

##### `--replace`/`-r`
Format the file in-place by replacing it with its formatted contents.
    
##### `--sort`/`-s`
Sort keys recursively by alphabet
