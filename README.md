# GOTASK

Tiny project aiming to create a CLI task tracker in Golang

## Purpose

1. Practice Golang
1. Keep track of tasks

## Features
### Add tasks
```console
[user@host]$ gotask-cli add "Publish code"
New task "Publish code" added to do @ id=4
```

### List all tasks
```console
[user@host]$ gotask-cli list
+----+--------------+--------+---------------------+---------+
| ID | DESCRIPTION  | STATUS | CREATED             | UPDATED |
+----+--------------+--------+---------------------+---------+
|  1 | Update doc   | todo   | 2026-04-03 22:34:40 |         |
|  2 | Write code   | todo   | 2026-04-03 22:38:26 |         |
|  3 | Write tests  | todo   | 2026-04-03 22:38:45 |         |
|  4 | Publish code | todo   | 2026-04-03 22:38:56 |         |
+----+--------------+--------+---------------------+---------+
```

### Mark progress
```console
[user@host]$ gotask-cli progress 4
[todo] task "Publish code" progress to [doing]
```

### List tasks by status
```console
[user@host]$ gotask-cli list doing
+----+--------------+--------+---------------------+---------------------+
| ID | DESCRIPTION  | STATUS | CREATED             | UPDATED             |
+----+--------------+--------+---------------------+---------------------+
|  4 | Publish code | doing  | 2026-04-03 22:38:56 | 2026-04-03 22:44:36 |
+----+--------------+--------+---------------------+---------------------+
```

### Mark regress
```console
[user@host]$ gotask-cli regress 4
[doing] task "Publish code" regress to [todo]
```

## Storing data
For convenience reasons, the data will be stored in a JSON file with the following structure.

```json
[
	{
		"Id": 1,
		"Description": "Update doc",
		"Status": 0,
		"Created": 1775248480,
		"Updated": -1
	},
	{
		"Id": 2,
		"Description": "Write code",
		"Status": 0,
		"Created": 1775248706,
		"Updated": -1
	},
	{
		"Id": 3,
		"Description": "Write tests",
		"Status": 0,
		"Created": 1775248725,
		"Updated": -1
	},
	{
		"Id": 4,
		"Description": "Publish code",
		"Status": 0,
		"Created": 1775248736,
		"Updated": 1775249113
	}
]
```

Note : todo=0 / doing=1 / done=2