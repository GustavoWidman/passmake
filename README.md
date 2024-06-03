# passmake

## Description

This is a simple CLI written in Go that generates a random password. The password is generated using the `crypto/rand` package. The password length and character set can be customized using the flags. There are predefined characters sets that can be mixed and matched to create a custom character set or you can create your own character set, see `Usage` for more information.

## Dependencies

The following Go packages were used in this project:

- [cobra](https://cobra.dev/)

These next dependencies are system dependencies that you will need to install in order to build the project.

- [git](https://git-scm.com/)
- [go](https://go.dev/)

Please consult your package manager's documentation on how to install these packages. If you are like me and you use Arch Linux (btw), you can install them with the following command:

```bash
sudo pacman -S git go
```

Note: I'm pretty sure `git` come default with arch btw because i already had it installed but make sure it is installed nonetheless.

## Installation

Download the executable from the [releases page](https://github.com/GustavoWidman/video2gif/releases/) and add it to your PATH or just throw it in `/usr/local/bin` if you feel like it. You can also build the executable yourself, see `Manual Build` for more information.

## AUR

If you are using Arch Linux (btw), you can install the package from the AUR. The package is called `passmake` and you can install it with your favorite AUR helper. Here is an example using `yay`:

```bash
yay -S passmake
```

## Manual Build

If you want to build the executable yourself, you use the `go build` command to build an executable yourself. This will create an executable called `passmake` in the current directory. You can then add it to your PATH or just throw it in `/usr/local/bin` if you feel like it.

```bash
go mod tidy # This will download the dependencies
go build -o passmake -ldflags "-s -w" src/main.go # This will build the executable to a file called `passmake`
```

When building the executable, the `-ldflags "-s -w"` flag is used to strip the debug information from the executable. This will make the executable smaller, more secure and faster overall.
After building the executable, you can add it to your PATH or just throw it in `/usr/local/bin` if you feel like it.

## Usage

After installed and added to your PATH, you can check the help message by running the following command:

```bash
$ passmake --help
Generate a random password

Description:
  This command generates a random password with the specified length and charset.
  The length of the password can be configured using the flag --length or -l.
  Note that length is a unsigned 8-bit integer, so the maximum value is 255 and the minimum is 1 (0 throws an error).
  The charset can be configured using the flags --lowercase, --uppercase, --numbers, --symbols or -L, -U, -N, -S respectively.
  Setting at least one of the flags will override the default charset, in which case
  lowercase, uppercase and numbers or LUN will be overwritten to the specified charset.
  If the flag --all (or -A) is set, all characters will be included in the charset.
  The flag --charset can be used to override the charset with a custom string, it has precedence over the other flags.

Usage:
  passmake [flags]

Examples:
  passmake              # Generate a password with length 16
  passmake -l 32        # Generate a password with length 32
  passmake -L           # Generate a password with length 16 and only lowercase characters
  passmake -l 12 -UN    # Generate a password with length 12 and only uppercase and numbers
  passmake -LUSl 64     # Generate a password with length 64 and all characters except numbers


Flags:
  -A, --all              Use all characters
      --charset string   Overrides all other charset flags, picking at random from the given string
  -h, --help             help for passmake
  -l, --length uint8     The length of the password (default 16)
  -L, --lowercase        Use lowercase characters (default true)
  -N, --numbers          Use numbers (default true)
  -S, --symbols          Use symbols
  -U, --uppercase        Use uppercase characters (default true)
```

This will show you the help message with all the available options. You can then use the command to generate a password with the desired length and charset.

## Examples

Here are some examples of how you can use the CLI:

```bash
# Generate a password with the default settings (length 16, lowercase, uppercase and numbers).

$ passmake
```

```bash
# Generate a password with length 32 and the default charset (LUN).

$ passmake -l 32
```

```bash
# Generate a password with length 128 and only lowercase characters.

$ passmake -Ll 128
```

More examples can be found in the `Usage` section or simply by calling the help menu with the `--help` or `-h` flag.

## License

This project is licensed under the CC0 1.0 Universal License - see the [LICENSE](LICENSE) file for details.
