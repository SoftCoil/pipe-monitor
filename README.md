
# Pipe Monitor

Monitor pipe progress via output to standard error.

Similar to functionality provided by the Pipe Viewer (pv) command, except this command is designed to work in 
environments where a console is not available.

## Usage Synopsis
Run with --help flag for usage:
```
# pm --help
Usage: pm [--size SIZE] [--name NAME] [--format FORMAT] INPUT_FILE

Positional arguments:
  INPUT_FILE             Optional input file. If not provided input will be read from STDIN

Options:
  --size SIZE, -s SIZE   Size of input from STDIN. Ignored if using INPUT_FILE
  --name NAME, -n NAME   A NAME tag for this output. Will be pre-pended to default FORMAT string
  --format FORMAT, -f FORMAT
                         Output format string. Allowed keys: %name, %size, %time, %eta, %percent, %written, %buffered
  --help, -h             display this help and exit

```

## Binary Release
Binary releases are provided for 45 OS/CPU Architecture combinations. You can download the latest binary
release here: [https://github.com/SoftCoil/pipe-monitor/releases](https://github.com/SoftCoil/pipe-monitor/releases)

## Build From Source

Set up your local golang environment, then clone this repository and run the make command from the project root directory.

`$ cd pipe-monitor`  
`$ make`

The "pm" binary will be build and saved in the pipe-monitor/bin directory.

## Examples

Pipe Monitor can read from stdin, in which case it will not print out any size related data:
```
# cat /dev/zero | pm > zero_file
Processed 0 bytes. 0 bytes buffered. Running 0s
Processed 1601699840 bytes. 0 bytes buffered. Running 2s
Processed 1989455872 bytes. 20480 bytes buffered. Running 4s
Processed 2271846400 bytes. 24576 bytes buffered. Running 6s
Processed 2540335104 bytes. 36864 bytes buffered. Running 8s
^Csignal: interrupt

# ls -l zero_file
-rw-rw-r-- 1 root root 2609664000 Jan  2 19:23 zero_file
```

Or it can read from a file, in which case it will print out more detailed information that requires knowing the size 
beforehand:
```
# pm zero_file > zero_file2
Processed 0 bytes of 2609664000 (0% complete). 0 bytes buffered. Running 0s, eta: <unknown>
Processed 1673056256 bytes of 2609664000 (64% complete). 4665344 bytes buffered. Running 2s, eta: 1s
Processed 2008084480 bytes of 2609664000 (76% complete). 5181440 bytes buffered. Running 4s, eta: 1s
Processed 2285899776 bytes of 2609664000 (87% complete). 10481664 bytes buffered. Running 6s, eta: 1s
Processed 2598014976 bytes of 2609664000 (99% complete). 2453504 bytes buffered. Running 8s, eta: 0s
Processed 2609664000 bytes of 2609664000 (100% complete). 0 bytes buffered. Running 8s, eta: 0s
```

If you wish to read from stdin, but know the total size, you can provide it:
```
# pm --size 2609664000 zero_file > zero_file2
Processed 0 bytes of 2609664000 (0% complete). 0 bytes buffered. Running 0s, eta: <unknown>
Processed 1673056256 bytes of 2609664000 (64% complete). 4665344 bytes buffered. Running 2s, eta: 1s
Processed 2008084480 bytes of 2609664000 (76% complete). 5181440 bytes buffered. Running 4s, eta: 1s
Processed 2285899776 bytes of 2609664000 (87% complete). 10481664 bytes buffered. Running 6s, eta: 1s
Processed 2598014976 bytes of 2609664000 (99% complete). 2453504 bytes buffered. Running 8s, eta: 0s
Processed 2609664000 bytes of 2609664000 (100% complete). 0 bytes buffered. Running 8s, eta: 0s
```

All output can be controlled with a provided format string:
```
# pm --format "Copy file zero_file: %written completed. Eta is %eta"  zero_file > zero_file2
Copy file zero_file: 0 completed. Eta is <unknown>
Copy file zero_file: 1602396160 completed. Eta is 1s
Copy file zero_file: 1921900544 completed. Eta is 1s
Copy file zero_file: 2215444480 completed. Eta is 1s
Copy file zero_file: 2487836672 completed. Eta is 0s
Copy file zero_file: 2609664000 completed. Eta is 0s
```