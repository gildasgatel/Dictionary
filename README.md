# Dictionary

Dictionary is a Golang project that provides a dictionary manager through a Command-Line Interface (CLI). Badger is used for data persistence, ensuring a reliable and efficient database.

![This is an image](https://github.com/gildasgatel/Dictionary/blob/master/book.jpg)


## Features

- **CRUD:** Complete implementation of CRUD (Create, Read, Update, Delete) operations to manage dictionary entries.
- **CLI Interface:** Simple usage through a command-line interface to interact with the dictionary.

## Usage

To use the program, execute the following command in your terminal:

```bash
./Dictionary list
```

## Available Commands:
Add an entry to the dictionary:
```bash
./Dictionary add <key> <value>
```
Read an entry from the dictionary:
```bash
./Dictionary get <key>
```
Update an entry in the dictionary:
```bash
./Dictionary update <key> <new_value>
```
Delete an entry from the dictionary:
```bash
./Dictionary delete <key>
```

## Compatibility
- Windows: The application is compatible with Windows systems.
- Mac: The application is compatible with Mac OS.
- Linux: The application is compatible with Linux distributions.
- Android: The application can be used on Android using compatible terminals or emulators.
