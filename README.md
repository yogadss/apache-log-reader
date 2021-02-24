#Description
CLI tools to get last n minutes from apache_format log files under specific directory

# System Requirement

## Development

### Must Have
- `go v1.13.x` 

# Usage
`./analytics <flags> -t=<last_n_minutes> -d=<full_path_to_directory> -f=<file_prefix>`

#Arguments
- `last_n_minutes` to determinethe last n minutes of the log will be displayed
- `full_path_to_directory` full path to log directory to be scanned
- `file_prefic` (optional) determine prefix file to be scanned. Default is `http`

#Flags
- `verbose` to display info level SYSTEM LOG

# Build
`go build -o analytics`