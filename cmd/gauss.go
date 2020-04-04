/*
Copyright © 2020 Kirill Denisov <kirill.denisov700@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kirilldenisov/rndm/random"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// gaussCmd represents the gauss command
var gaussCmd = &cobra.Command{
	Use:   "gauss ([mean] [sd] | \"std\")",
	Short: "Generates numbers from normal distribution",
	Long: `Generates numbers from normal distribution with given mean and standart deviation.
	Or simply type \"st\" to use standart normal distribution (N(0,1)).
	If you want to copy-paste the result to your python or R script/notebook,
	there is --format flag available.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		home, _ := os.UserHomeDir()
		viper.SetConfigFile(home + "/.rndm.yaml")
		configerr := viper.ReadInConfig()
		if configerr != nil {
			return fmt.Errorf("provide your api key with " + loginCmd.Use)
		}

		if len(args) < 1 {
			cmd.Help()
			return fmt.Errorf("pass some args, boy")
		}

		format, _ := cmd.Flags().GetString("format")
		number, _ := cmd.Flags().GetInt("number")
		precision, _ := cmd.Flags().GetInt("precision")
		token := viper.GetString("token")
		a := &random.API{token}

		if args[0] == "std" {
			set, err := a.Normal(format, 0, 1, number, precision)
			if err != nil {
				return err
			}
			fmt.Println(set)
			return nil
		}

		mean, wrongArgOne := strconv.ParseFloat(args[0], 32)
		sd, wrongArgTwo := strconv.ParseFloat(args[1], 32)
		if wrongArgOne != nil || wrongArgTwo != nil {
			return fmt.Errorf("wrong arguments, boy")
		}

		set, err := a.Normal(format, float32(mean), float32(sd), number, precision)
		if err != nil {
			return err
		}
		fmt.Println(set)
		return nil
	},
}

func init() {
	gaussCmd.Flags().StringP("format", "f", "", "Formats your set of number. Available formats are:\n \"break\" - every number on a new line\n \"py\" - Python list\n \"r\" - R vector")
	gaussCmd.Flags().IntP("number", "n", 10, "Number of numbers to generate")
	gaussCmd.Flags().IntP("precision", "p", 3, "Number of places to the right of the decimal point")
	rootCmd.AddCommand(gaussCmd)
}
