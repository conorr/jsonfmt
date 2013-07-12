jsonfmt
=======

JSON formatter utility
----

`jsonfmt` takes a file containing JSON data and outputs a formatted, readable version to stdout, where it can be grepped, piped to another file, etc. This is especially useful for minified JSON files.

```
$cat example.json
{"fruits":["apple","orange","banana"],"veggies":["lettuce","carrots","celery"]}
$jsonfmt example.json
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

`jsonfmt` can also handle files containing JSONP, such as a REST API response saved to a file. A JSONP response is wrapped in a callback and, though it is valid JavaScript, it is not valid JSON. However `jsonfmt` automatically detects JSONP and handles it accordingly:

```
$cat example.js
SOME_JSONP_CALLBACK({"fruits":["apple","orange","banana"],"veggies":["lettuce","carrots","celery"]})
$jsonfmt example.js
SOME_JSONP_CALLBACK({
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
})
```

`jsonfmt` is written in Go, so it's really fast!
