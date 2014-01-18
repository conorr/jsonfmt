jsonfmt
=======

JSON formatting utility
----

`jsonfmt` takes a file containing JSON data and formats it.

```
$ cat example.json
{"fruits":["apple","orange","banana"],"veggies":["lettuce","carrots","celery"]}
$ jsonfmt example.json
$ cat example.json
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

It can also handle JSONP format. For example:

```
$ cat example.js
SOMECALLBACK({"foo":"bar"})
$ jsonfmt example.js
$ cat example.js
SOMECALLBACK({
    "foo": "bar"
})
```
