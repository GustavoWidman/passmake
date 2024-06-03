package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/GustavoWidman/passmaker/src/utils"
)

func GenerateCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "passmaker",
		Short: "A simple password generator",
		Long: `Generate a random password

Description:
  This command generates a random password with the specified length and charset.
  The length of the password can be configured using the flag --length or -l.
  Note that length is a unsigned 8-bit integer, so the maximum value is 255 and the minimum is 1 (0 throws an error).
  The charset can be configured using the flags --lowercase, --uppercase, --numbers, --symbols or -L, -U, -N, -S respectively.
  Setting at least one of the flags will override the default charset, in which case
  lowercase, uppercase and numbers or LUN will be overwritten to the specified charset.
  If the flag --all (or -A) is set, all characters will be included in the charset.
  The flag --charset can be used to override the charset with a custom string, it has precedence over the other flags.`,
		Example: `  passmake		# Generate a password with length 16
  passmake -l 32	# Generate a password with length 32
  passmake -L		# Generate a password with length 16 and only lowercase characters
  passmake -l 12 -UN	# Generate a password with length 12 and only uppercase and numbers
  passmake -LUSl 64	# Generate a password with length 64 and all characters except numbers
`,
		Args: cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if cmd.Flag("length").Changed && utils.Length == 0 {
				cmd.Help()
				cmd.Println("\nError: length must be greater than 0")
				return
			}

			if cmd.Flag("charset").Changed {
				charset_string := utils.RemoveDuplicateLetters(utils.CharsetOverride)
				password := utils.GeneratePassword(utils.Length, charset_string)
				cmd.Println(password)
				return
			}

			charset, err := utils.FlagsToCharset(cmd)

			if err != nil {
				cmd.Help()
				cmd.Println("\nError: " + err.Error())
				return
			}

			charset_string := utils.CharsetToString(charset)

			password := utils.GeneratePassword(utils.Length, charset_string)
			cmd.Println(password)
		},
	}

	utils.AddFlags(cmd)

	return cmd
}

func Execute() {
	if err := GenerateCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
