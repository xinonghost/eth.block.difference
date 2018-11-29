/**
 * Find Ethereum block number difference between few sources.
 */

package main

import (
        "fmt"
        "net/http"
        "io/ioutil"
        "encoding/json"
        "bytes"
        "strconv"
        "math"
)

type Response struct {
    Jsonrpc string
    Id int
    Result string
}

func main() {
    var etherscan float64 = float64(getFromEtherscan())
    var infura float64 = float64(getFromInfura())

    fmt.Println("Block number difference is ", math.Abs(etherscan - infura))
}

func getFromEtherscan() int {
    resp, _ := http.Get("https://api.etherscan.io/api?module=proxy&action=eth_blockNumber")

    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    
    return getBlockNumber(body)
}

func getFromInfura() int {
    var url = "https://mainnet.infura.io/v3/63206be06ce54120b6cad3e60a677267"
    var requestBody = []byte(`{"jsonrpc":"2.0","method":"eth_blockNumber","params": [],"id":1}`)

    req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, _ := client.Do(req)

    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    return getBlockNumber(body)
}

func getBlockNumber(bytes []byte) int {
    var response Response

    json.Unmarshal(bytes, &response)

    number, _ := strconv.ParseInt(response.Result, 0, 64)

    return int(number);
}