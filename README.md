# AWS Config Manager

This is a small CLI application to help me manage AWS Config files.

## Usage

Create a blank pair of `config` and `credentials` files.
```shell
awscm fresh
```

Save your current `config` and `credentials` files into a named pair.
```shell
awscm save work
```

Replace your current `config` and `credentials` files with a named pair.
```shell
awscm use work
```

Delete a named `config` and `credentials` pair.
```shell
awscm delete work
```