jsonfmt
=======

A fast JSON formatting utility
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

You can use the `--sort`/`-s` flag to sort keys alphabetically:

```
$ cat fruits.json
{"bananas":2,"apples":5,"pineapples":1,"mangoes":3}
$ jsonfmt fruits.json --sort
$ cat fruits.json
{
    "apples": 5,
    "bananas": 2,
    "mangoes": 3,
    "pineapples": 1
}
```

It also handles JSONP automatically. For example:

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
