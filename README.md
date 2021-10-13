# OneTimePad

## About

OneTimePad or `onepad` is a simple tool that encrypts and decrypts files using a so called [one time pad][wiki-otp]. That means that a key is randomly generated and then applied to the file. Without this key it is virtually impossible to decrypt the file.

## Usage

After [installing](#installation) `onepad` you can use it as follows.

`onepad <filename> <keypath / keyfile>`

The first argument (`filename`) is the file you want to encrypt or decrypt.

The second argument (`keypath / keyfile`) can be two types of inputs:

- a directory where the key will be stored
- a file ending in `.key` where the key corresponding to the encrypted file is

Depending on the type of input, the given file will either be encrypted (a `keypath` is given) or decrypted (a `.key` file is given).

When encrypting, the key will be stored in the given directory with this format: `<filename>.key`.

When decrypting, the key file will be deleted after decrypting.

## Installation

You can either download one of the prebuild executables or download the source code and compile it yourself [here][git-release].

To build from source, just run `go build onepad.go`.

To install it run `go install onepad.go`.

[wiki-otp]: https://de.wikipedia.org/wiki/One-Time-Pad
[git-release]: https://github.com/Flobii/one-time-pad/releases/
