jsonfmt
=======

JSON formatter utility
----

`jsonfmt` takes a file containing JSON data and formats it.

```
$cat example.json
{"fruits":["apple","orange","banana"],"veggies":["lettuce","carrots","celery"]}
$jsonfmt example.json
$cat example.json
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
