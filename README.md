# Blog Aggregator

## Usage

Check out the [documentation](https://github.com/dagregi/blog-aggregator/tree/master/docs) on how to use the API.

## Installation

Make sure you have the following tools installed on your machine.

-   [Goose (for DB migrations)](https://github.com/pressly/goose)

## Running the project

Firstly make sure you have a PostgreSQL database running on your machine.

Then create a database with the name you want _(`blogator` is the default)_ and run the migrations.

```bash
make migrate/up
```

After that, you can run the project with the following command:

```bash
make run
```

## Running the tests

To run the tests, you can use the following command:

```bash
make test
```
