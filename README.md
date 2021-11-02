# Cron Expression Parser

## Description
Command line script which parses a cron string and
expands each field to show the times at which it will run.  
Parser support standart cron format, 
except special strings such as "@yearly"

## Example
For example, the following input argument:
```shell
*/15 0 1,15 * 1-5 /usr/bin/find
```
Should yield the following output:
```shell
minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
```

## Usage
The cron string will be passed to your application as a single argument.

```shell
~$ cron "*/15 0 1,15 * 1-5 /usr/bin/find"
```

## Testing
That application has tests that can be run by:
```shell
make test
```

Output should be like:
```shell
go test -race -cover ./...
ok      cron    0.407s  coverage: 68.3% of statements
```