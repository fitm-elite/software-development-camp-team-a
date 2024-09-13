# Elebs
Elebs is a command-line tool designed to scrape and extract data from CSV files containing electric bills for KMUTNB Dormitory, Prachinburi campus. The tool automates the process of reading, filtering, and parsing electric bill data, Ideal for dormitory management and tenants, Elebs enables quick access to utility data and simplifies the workflow for handling multiple records at once. Its focus is on accuracy, speed, and ease of use for administrative tasks.

## Using

Prepare your environment for `elebs` script
```shell
# Elebs configuration file

# MinIO configuration
export MINIO_ENDPOINT=""
export MINIO_ACCESS_KEY_ID=""
export MINIO_SECRET_ACCESS_KEY=""

# Line Messaging API configuration
export LINE_CHANNEL_SECRET=""
export LINE_CHANNEL_ACCESS_TOKEN=""
```

Use the command to build the main package to `binary file`:
```shell
make main
```

Use `elebs` command to execute the script:
```shell
elebs <*.csv>
```
