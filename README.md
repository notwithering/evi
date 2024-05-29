# Evi - The encrypted editing layer

[![MIT License](https://img.shields.io/badge/License-MIT-a10b31)](LICENSE)

**Evi** is a layer that sits inbetween encryption and your text editor to provide a seamless experience of instantaneous security while you write your documents.

When opening a file, it decrypts it using a user-specified key. The program then sends you over to your default text editor to easily edit the file. After you exit the editor, the file is then easily secured with AES-256 encryption.

```ruby
$ evi test.txt
:: Encryption key:
:: [d]etails   [e]ncryption
>> *******

$ echo "hello" > test.txt
$ cat test.txt
%,y^�X⏭5:��M��t���{43+�;2��D
```

## Installation

> [!WARNING]
> Losing your key will result in you not being able to decrypt your file.

### Installing

```bash
go install github.com/notwithering/evi@latest
```

### Testing

```bash
go run github.com/notwithering/evi@latest file.txt
```