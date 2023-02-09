# ZINC INDEXER

Simple app to insert files with the next pattern to zinc db.

```
Message-ID: string
Date: string
From: string
To: string
Subject: string
Mime-Version: string
Content-Type: string
Content-Transfer-Encoding: string
X-From: string
X-To: string
X-cc: string
X-bcc: string
X-Folder: string
X-Origin: string
X-FileName: string

message content string
```

## How to use.

### Docker

```
    docker build . -t indexer_image
    docker run indexer_image
```

### Compile and use

If you want to use binary source, then you should run commands below.

```
go build -o indexer_name main.go
./indexer_name {folder_path}
```

### Use dev or prod environment.

If you want to use dev environment you could use

```
go run tags dev main.go {folder_path} # dev env
go run tags prod main.go {folder_path} # prod env
```