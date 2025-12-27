# loto

Proposing lottery ticket candidates for Japan ([Takarakuji](https://ja.wikipedia.org/wiki/%E5%AE%9D%E3%81%8F%E3%81%98)).
Applicable to "Loto" or "Numbers".

This tool is purely a complete random pick;
it does not analyze or suggest candidates, nor does it guarantee winning.

Think of it as a simple joke tool for a kind of enjoyment.

## usage

This is the standard usage.  
Below is an example proposing 10 candidates for Loto6.

```bash
loto -n 10 loto6
```

Available argument names can be viewed with `loto list`.

The following arguments are valid:

- loto6
- loto7
- miniloto
- numbers3
- numbers4

## help

```
Usage:
  loto [flags]
  loto [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        Displays the available argument names

Flags:
  -h, --help         help for loto
  -n, --length int   Specify the number of lottery results to pick (default 5)
```
