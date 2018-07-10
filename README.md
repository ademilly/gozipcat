# gozipcat

gozipcat is a command line utility unarchiving and concatenating zip archives content.

## install

```bash
    $ go get github.com/ademilly/gozipcat
    .
```

## usage

```bash
    $ gozipcat -h
    Usage of gozipcat:
        -prefix string
                prefix to use in output filenames
        -root string
                path to directory containing zip archives
    $ ls examples/
    archive_one.zip archive_two.zip
    $ unzip -l examples/archive_one.zip
    Archive:  /path/to/examples/archive_one.zip
      Length      Date    Time    Name
    ---------  ---------- -----   ----
        .......................   hello.csv
        .......................   world.csv
    ---------                     -------
        .......................   2 files
    $ unzip -l examples/archive_two.zip
    Archive:  /path/to/examples/archive_two.zip
      Length      Date    Time    Name
    ---------  ---------- -----   ----
        .......................   hello.csv
        .......................   world.csv
    ---------                     -------
        .......................   2 files
    $ gozipcat -root examples -prefix myprefix
    $ ls examples/
    archive_one.zip archive_two.zip myprefix_hello.csv myprefix_world.csv
```
