# Evi - The encrypted editing layer

[![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](LICENSE)

**Evi** is a layer that sits inbetween encryption and your text editor to provide a seamless experience of instantaneous security while you write your documents.

When opening a file, it decrypts it using a user-specified key. The program then sends you over to your default text editor to easily edit the file. After you exit the editor, the file is then easily secured with defaulted AES-256 encryption.

```ruby
$ evi test.txt
:: Encryption key:
:: [d]etails   [a]lgorithm   [m]ode
>> supersecret

$ echo "hello" > test.txt
$ cat test.txt
%,y^�X⏭5:��M��t���{43+�;2��D
```

## Installation

### Installing

```bash
go install github.com/notwithering/evi@latest
```

### Testing

```bash
go run github.com/notwithering/evi@latest file.txt
```

## Current Options

> [!NOTE]
> More options are being added such as different algorithms, modes, etc.

### Algorithms

- AES-256

### Modes

- GCM
