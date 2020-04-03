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
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rndm",
	Short: "Abusing random.org API",
	Long: `
 __________________________________________________________
|                                                          |     Q
|       .-.             .-.       .-.           .-.        |  ___|\_.-,
|      (_) )-.            ;  :   (_) )-.         .;|/:     S\ Q~\___ \|
|        .:   \         .;:  :     .:   \       .;   :     |(   )o 5) Q
|       .::.   )       .;' \ :    .:'    \     .;    :     |\\  \_ ()
|     .-:. ':-'    .:'.;    \:  .-:.      ).:'.;     :     | \'. _'/'.
|    (_/     ':._.(__.'      '.(_/  '----'(__.'      '.   .-. '-(  x< \
|                                             ,o         /\, '.  )  /'\\
|____________________________________________ \'.__.----/ .'\  '.-'/   \\
  by kirilldenisov                             '---'q__/.'__ ;    /     \\_
                                                    '---'     '--'       '"'

rndm is the most joyful and useful cli app you'll ever see! 
True random -> true fun!
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
