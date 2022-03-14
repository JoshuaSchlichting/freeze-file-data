# freeze-file-data

This is a command line application used for generating a JSON map file of a directory and all of its files (and the details for that file). The program is written Go and utilizes Go Channels to write the output file *while* reading scanning the content of the specified directory.

# Usage
> NOTE: The following instructions assume you are in the root directory of this project.
1. Build the application
```bash
chmod +x scripts/build.sh
./scripts/build.sh
# Newly created executable file is in ./bin/
```

2. Execute the application
```bash
./bin/freeze-file-data -h
> Usage of ./bin/freeze-file-data_macos_x86_64:
>   -R    Inspect directory recursively. Does nothing when target is a file.
>   -target string
>         File or directory to inspect.
```

### Example:
```bash
./bin/freeze-file-data -target ~ -R
> 2022/03/14 00:14:29 target: /Users/josh
> 2022/03/14 00:14:29 Scanning '/Users/josh'...
```
### Output:
The application creates a JSON file in the target directory, called `describeFiles.json`. Following the example above, we could check our output with `cat ~/describeFiles.json`.