# GOTASK

Tiny project aiming at creating a CLI task tracker in Golang

## Purpose

1. Practice Golang
1. Keep track of tasks

## Roadmap
1. ~~Add~~
2. ~~List all~~
3. Progress
4. Regress
5. List by

## Features
### Add tasks
```console
[user@host]$ gotask add "Publish code"
New task "Publish code" added to do @ id=4
```

### List all tasks
```console
[user@host]$ gotask list
+---------------------------------------------------------------------+
|id| Description  |   Status    |    Created       |    Updated       |
|--|--------------|-------------|------------------|------------------|
| 1| Update doc   | todo        | 01-03-2026 09:23 |                  | # Never updated
| 2| Write code   | doing       | 04-03-2026 10:13 | 05-03-2026 17:45 | # Set in progress at UpdateAt
| 3| Write tests  | done        | 05-03-2026 11:34 | 01-04-2026 19:43 | # Done at UpdateAt
| 4| Publish code | todo        | 01-04-2026 09:13 |                  |
+---------------------------------------------------------------------+
```

### Mark progress
```console
[user@host]$ gotask progress 4
[todo] task "Publish code" progress to [doing]
```

### List tasks by status
```console
[user@host]$ gotask list doing
+---------------------------------------------------------------------+
|id| Description  |   Status    |    Created       |    Updated       |
|--|--------------|-------------|------------------|------------------|
| 2| Write code   | doing       | 04-03-2026 10:13 | 05-03-2026 17:45 |
| 4| Publish code | doing       | 01-04-2026 09:13 | 01-04-2026 10:30 | 
+---------------------------------------------------------------------+
```

### Mark regress
```console
[user@host]$ gotask regress 4
[doing] task "Publish code" regress to [todo]
```

### List task by id
```console
[user@host]$ gotask list 4
+---------------------------------------------------------------------+
|id| Description  |   Status    |    Created       |    Updated       |
|--|--------------|-------------|------------------|------------------|
| 4| Publish code | todo        | 01-04-2026 09:13 | 01-04-2026 10:33 | 
+---------------------------------------------------------------------+
```


## Storing data
For convenience reasons, the data will be stored in a JSON file with the following structure.

```json
[
    {
        "id": 1,
        "Description": "Update doc",
        "Status": 0,
        "Created": 1772353380,
        "UpdatedAt": -1
    },
    {
        "id": 2,
        "Description": "Write code",
        "Status": 1,
        "Created": 1772615580,
        "UpdatedAt": 1772729100
    }
]
```

Note : todo=0 / doing=1 / done=2