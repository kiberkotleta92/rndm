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
	"rndm/random"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// intCmd represents the int command
var intCmd = &cobra.Command{
	Use:   "int [min] [max]",
	Short: "Generates integer from range",

	RunE: func(cmd *cobra.Command, args []string) error {
		home, _ := os.UserHomeDir()
		viper.SetConfigFile(home + "/.rndm.yaml")
		configerr := viper.ReadInConfig()
		if configerr != nil {
			return fmt.Errorf("provide your api key with " + loginCmd.Use)
		}

		if len(args) < 2 {
			cmd.Help()
			return fmt.Errorf("pass some args, boy")
		}

		min, wrongArgOne := strconv.Atoi(args[0])
		max, wrongArgTwo := strconv.Atoi(args[1])
		if wrongArgOne != nil || wrongArgTwo != nil {
			return fmt.Errorf("wrong arguments, boy")
		}
		if min > max {
			return fmt.Errorf("boy, are you sure you know what min and max stands for?")
		}
		token := viper.GetString("token")
		a := &random.API{token}
		i, err := a.Integer(min, max)
		if err != nil {
			return err
		}
		fmt.Println(i)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(intCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// intCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// intCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
