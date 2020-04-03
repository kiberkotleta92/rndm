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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// coinCmd represents the coin command
var coinCmd = &cobra.Command{
	Use:   "coin",
	Short: "Flips coin, nothing more",
	RunE: func(cmd *cobra.Command, args []string) error {
		home, _ := os.UserHomeDir()
		viper.SetConfigFile(home + "/.rndm.yaml")
		configerr := viper.ReadInConfig()
		if configerr != nil {
			return fmt.Errorf("provide your api key with " + loginCmd.Use)
		}
		token := viper.GetString("token")
		a := &random.API{token}

		c, err := a.Coin()
		if err != nil {
			return fmt.Errorf("check connection or if your key is valid (%s)", err)
		}
		fmt.Println(c)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(coinCmd)
}
