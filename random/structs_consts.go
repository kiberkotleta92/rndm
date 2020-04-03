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

package random

//API stores random.org api key
type API struct {
	Token string
}

//RandomResponse is self explanatory, need it for fast (un)marshalling with easyjson
type RandomResponse struct {
	ID      int64  `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		AdvisoryDelay int64 `json:"advisoryDelay"`
		BitsLeft      int64 `json:"bitsLeft"`
		BitsUsed      int64 `json:"bitsUsed"`
		Random        struct {
			CompletionTime string        `json:"completionTime"`
			Data           []interface{} `json:"data"`
		} `json:"random"`
		RequestsLeft int64 `json:"requestsLeft"`
	} `json:"result"`
}

//RandomRequest is self explanatory, need it for fast (un)marshalling with easyjson
type RandomRequest struct {
	ID      int64       `json:"id"`
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

//structs for spicific methods' parameters

//GenerateInteger represents params of generateIntegers method
type GenerateInteger struct {
	APIKey      string `json:"apiKey"`
	Max         int    `json:"max"`
	Min         int    `json:"min"`
	N           int    `json:"n"`
	Replacement bool   `json:"replacement"`
}

//GenerateGaussians represents params of generateGaussians method
type GenerateGaussians struct {
	APIKey string  `json:"apiKey"`
	Mean   float32 `json:"mean"`
	SD     float32 `json:"standardDeviation"`
	N      int     `json:"n"`
	Digits int     `json:"significantDigits"`
}

//HEADS taken from https://www.asciiart.eu
const HEADS = `
               ,,==="""""""===,,
           ,==""' |\ |   /\   '""==,
        ,="'|\    | \|  /__\   /\  '"=,
      /"    |,"\  |  | /'  '\ /  )     "\
    /"  ,"  |                 '\/    /|  "\
   /'  |   ,                       /",|   '\
  /'   ",/"                           |    '\
 /'      I=I=I               ,d8ba,___      '\
/'     I=8=8=8=I_I_          88888P"""       '\
|   xXXXXXXXXXXXXXXXxIxx    ,888"             |
| ~XXXXXXXXXXXXXXX~-~-~-~-~ d888~-~-~-~-~-~-~ |
| ~-~-~-~-~-~-~-~-,aad888ba,8888,-~-~-~-~-~-~ |
| ~-~-~-~-~-~-,ad888888888888888b-~-~-~-~-~-~ |
\ ~-~-~-~-~,ad8888888888888888888-~-~-~-~-~-~ /
'\ -~-~-~-~-~-~-~-~-~-~-~-~-~-~-~-~-~-~-~-~- /'
 '\ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~,-,~~~~~ /'
  '\    /"\         1 9 9 4        \ /\    /'
   '\  "\,/'                   |\   '\ '  /'
     "\      /""\   |    |     |,'\     /"
       '"=,_ \__/   |__  |__   |    ,="' 
          '""=,__             __,=""'  
               '""=========""''
`

//TAILS taken from https://www.asciiart.eu
const TAILS = `
             _.oood"""""""booo._
         _.o""      _____    * ""o._
       oP"  _.ooo""""   """"o|o*_* "Yo
     o8   oP                 | |"._* '8o
    d'  o8'_.--._            | |/  ,\* 'b
   d'  d'.' __   ".          | |: (( '\
  8'  d'/,-"  '.   :         | |  ||\_/* '8
 8   8'|/      :   :    |)   _ |  || |'|   8
,8  8          :  :   /)| \ || |\_|| | |8  8.
8' ,8         /  :    " /_) |':' | | | |8. '8
8  8'        /  /       _ _='  \ ' __   __  8
8  8        /  /        \|__ |  | |  | | 8| 8
8  8.      /  /         ||   |  | |-:' | 8| 8
8. '8    ,' ,'       __/ |__ |__| |  \ |__|,8
'8  8  ,' ,'      _ /     __ . . . . . .8LL8'
 8   8"   '------'/(    ,'  '.'. | | ,-|8  8
  8.(_________dd_/  \__/ '  0|'.': |: (8 ,8
   Y.  Y.                    | :/| |,\|* .P
    Y.  "8.          .,o     | | |,|"*  ,P
     "8.  "Yo_               | |p|"* ,8"
       "Y_   '"ooo.__   __.oo|"* * _P"
         ''"oo_     """""    * _oo""'
              '"""boooooood"""'
`
