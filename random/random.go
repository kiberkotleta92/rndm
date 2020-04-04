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

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//format formats set of numbers to common array types for easy copy-pasting
func format(format string, arr []float64) string {
	var stringArray []string
	for _, val := range arr {
		s := fmt.Sprintf("%g", val)
		stringArray = append(stringArray, s)
	}
	switch format {
	case "break":
		return strings.Join(stringArray, "\n")
	case "py":
		return "[" + strings.Join(stringArray, ", ") + "]"
	case "r":
		return "c(" + strings.Join(stringArray, ", ") + ")"
	}
	return strings.Join(stringArray, " ")
}

//sendRequest uses fasthttp to send a request to random.org, returns data field of response
func sendRequest(name string, params interface{}) ([]interface{}, error) {
	var responseStruct RandomResponse
	var err error

	requestStruct := RandomRequest{
		Jsonrpc: "2.0",
		Method:  name,
		Params:  params,
	}

	requestBody, err := requestStruct.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("request json marshal: %s", err)
	}
	req := bytes.NewReader(requestBody)

	resp, err := http.Post("https://api.random.org/json-rpc/1/invoke", "", req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %s", err)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)

	err = responseStruct.UnmarshalJSON(respBody)
	if err != nil {
		return nil, fmt.Errorf("response json unmarshal: %s", err)
	}

	return responseStruct.Result.Random.Data, nil
}

//Coin flips coins and returns ascii heads or tails
func (a *API) Coin() (string, error) {
	params := GenerateInteger{
		APIKey: a.Token,
		Max:    1,
		Min:    0,
		N:      1,
	}

	res, err := sendRequest("generateIntegers", params)
	if err != nil {
		return "", fmt.Errorf("connection error: %s", err)
	}
	if len(res) < 1 {
		return "", fmt.Errorf("index problems")
	}
	if res[0].(float64) == 1 {
		return HEADS, nil
	} else {
		return TAILS, nil
	}
}

//Integer returns integer from a given range
func (a *API) Integer(min, max int) (string, error) {
	params := GenerateInteger{
		APIKey: a.Token,
		Max:    max,
		Min:    min,
		N:      1,
	}

	res, err := sendRequest("generateIntegers", params)
	if err != nil {
		return "", fmt.Errorf("connection error: %s", err)
	}
	if len(res) < 1 {
		return "", fmt.Errorf("index problems")
	}
	s := strconv.Itoa(int(res[0].(float64)))
	return s, nil
}

//Normal returns formated set of n numbers from normal distribution with
// given mean and standart deviation
func (a *API) Normal(frmt string, mean, sd float32, n, digits int) (string, error) {
	arr := make([]float64, 0, n)
	if digits == 0 {
		digits = 1
	}
	params := GenerateGaussians{
		APIKey: a.Token,
		Mean:   mean,
		SD:     sd,
		N:      n,
		Digits: digits,
	}
	res, err := sendRequest("generateGaussians", params)
	if err != nil {
		return "", fmt.Errorf("connection error: %s", err)
	}
	if len(res) < 1 {
		return "", fmt.Errorf("index problems")
	}
	for _, val := range res {
		arr = append(arr, val.(float64))
	}
	return format(frmt, arr), nil
}
