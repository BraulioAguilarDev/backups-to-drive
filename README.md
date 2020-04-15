# Upload file to Google Drive

Go tool to upload files

## Prerequisites

Create a new Cloud Platform project & generate `credentials.json` for [Drive API](https://www.iperiusbackup.net/en/how-to-enable-google-drive-api-and-get-client-credentials/)

## Usage

```shell
$ go run *.go -i ${file} -o ${description}  -f ${folder}
```

| Flag | Required | Description          |
|------|----------|----------------------|
| -i   | True     |     My file          |
| -o   | False    |     Description file |
| -f   | False    |     Folder           |

## TODO
Cron


[Was done with love](http://brauliodev.com/)
