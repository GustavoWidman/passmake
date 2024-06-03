package utils

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	Length          uint8
	Lowercase       bool
	Uppercase       bool
	Numbers         bool
	Symbols         bool
	All             bool
	CharsetOverride string
)

func FlagsToCharset(cmd *cobra.Command) (CharacterSet, error) {
	charset := CharacterSet{
		Lowercase,
		Uppercase,
		Numbers,
		Symbols,
	}

	if cmd.Flag("lowercase").Changed ||
		cmd.Flag("uppercase").Changed ||
		cmd.Flag("numbers").Changed ||
		cmd.Flag("symbols").Changed {
		if !Lowercase && !Uppercase && !Numbers && !Symbols {
			return charset, errors.New(
				"empty charset, please enable at least one of the following: lowercase, uppercase, numbers, symbols",
			)
		}

		charset = CharacterSet{
			Lowercase: cmd.Flag("lowercase").Changed && Lowercase,
			Uppercase: cmd.Flag("uppercase").Changed && Uppercase,
			Numbers:   cmd.Flag("numbers").Changed && Numbers,
			Symbols:   cmd.Flag("symbols").Changed && Symbols,
		}
	}

	if cmd.Flag("all").Changed {
		if !All {
			return charset, errors.New(
				"empty charset, please enable at least one of the following: lowercase, uppercase, numbers, symbols",
			)
		}

		charset = CharacterSet{
			Lowercase: true,
			Uppercase: true,
			Numbers:   true,
			Symbols:   true,
		}
	}

	return charset, nil
}

func AddFlags(cmd *cobra.Command) {
	cmd.Flags().Uint8VarP(
		&Length, "length", "l", 16, "The length of the password",
	)
	cmd.Flags().BoolVarP(
		&Lowercase, "lowercase", "L", true, "Use lowercase characters",
	)
	cmd.Flags().BoolVarP(
		&Uppercase, "uppercase", "U", true, "Use uppercase characters",
	)
	cmd.Flags().BoolVarP(
		&Numbers, "numbers", "N", true, "Use numbers",
	)
	cmd.Flags().BoolVarP(
		&Symbols, "symbols", "S", false, "Use symbols",
	)
	cmd.Flags().BoolVarP(
		&All, "all", "A", false, "Use all characters",
	)
	cmd.Flags().StringVar(
		&CharsetOverride, "charset", "", "Overrides all other charset flags, picking at random from the given string",
	)
}
