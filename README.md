# cq

`cq` is a small command line utility for parsing and filtering CSV data.

## Installation
TODO

## Usage

Given a file, `data.csv`:
```csv
name,age,email
bob,25,bob@bob.com
alice,23,
brian,73,brian@bob.com
```

We can run the following to filter the data:
```sh
cat data.csv | cq -filter='age>=25' -output=json -sort=age
```

And see the following output:
```json
[
  {
    "age": "25",
    "email": "bob@bob.com",
    "name": "bob"
  },
  {
    "age": "73",
    "email": "brian@bob.com",
    "name": "brian"
  }
]
```

```sh
cat mydata.csv | cq --output=json 
```
