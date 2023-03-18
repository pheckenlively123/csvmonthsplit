# csvmonthsplit

## Introduction

This is a small utility for splitting up the CSV history files generated
by Binance US for tax purposes.  Binance US limits report exports to 10
per month, so doing monthly exports of a past year is impossible in a
single month.  This utility takes a full year export, and it breaks it
up into monthly files.

## Build

Building `csvmonthsplit` requires the Go SDK version 1.20.1 or later.
Provided you have the SDK properly installed (google for how to do that),
run the commands below in the repo.

```
cd cmd/csvmonthsplit
go build
```

## Execution

Run `csvmonthsplit` with the `-h` flag as demonstrated below to see a usage declaration.

```
$ ./csvmonthsplit -h
Usage of ./csvmonthsplit:
  -infile string
        Input file.
```

Running `csvmonthsplit` on a Binance US export file should provide the results below.

```
$ ./csvmonthsplit -infile 202303051719_2021_all.csv
$ ls 202303051719_2021_all*
202303051719_2021_all-2021-01.csv  202303051719_2021_all2021-03.csv   202303051719_2021_all-2021-06.csv  202303051719_2021_all-2021-09.csv  202303051719_2021_all-2021-12.csv
202303051719_2021_all-2021-02.csv  202303051719_2021_all-2021-04.csv  202303051719_2021_all-2021-07.csv  202303051719_2021_all-2021-10.csv  202303051719_2021_all.csv
202303051719_2021_all-2021-03.csv  202303051719_2021_all-2021-05.csv  202303051719_2021_all-2021-08.csv  202303051719_2021_all-2021-11.csv
```

`csvmonthsplit` creates new monthly files using the full year file as
a starting point.  It appends the year and month in `YYYY-MM` format.
If these files exist in the directory already, *they will be overwritten.*