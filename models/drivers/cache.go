package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.1" )

func U1zBrToTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p3KJYrlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qO3V0qJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iiZ9gpk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQmdwlVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FyKTZ11zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grwpk7T2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mPRn8dM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wjeCNxvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x8uOhvwNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8Vs3jqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fCyeaNA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxcMUSs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B0WfBs29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MhOpJo45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9kNclVKlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o2R02DTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDzHWhX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DYpZEJVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWWHO2R1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fp4YuDgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FpdADzHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEEEQv2oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P5hzzel5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func alZDSOE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZoEDjgRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4yh3UheYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EF2RgvPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKbcUGHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xQ1HtrWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ye0Wul6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80mcrjSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WjF4qaUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suqmaaOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n3WnxOouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TC58aZXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aOHoBYSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umxl8x35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2fhjk06DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1B6bCrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KeaNnIn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eONlLIpUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ldg72bVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8gyQn64eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qWVLb2q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqtl9BuzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z6lV2X2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XNfc81gAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QPi7OmuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ciuJygdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YoipeqOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iMFF6VjDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTUKQN5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6UxNcMCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqaFvy09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Ut5BL5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Of8q16EMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HhjaT7VFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gR0IaFWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNE7DSkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y7IoZwBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsWOIKcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pClsXd9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5cStQiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSlHCEM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rcsGoAFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lLd3LRMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E02XU0TEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKPUmQo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3pDFll3zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LvsyXE3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func auUk2CdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TjcWxs1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gZxIbIeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1qMPegeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eGmXMdZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZtI8JgmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aolttdGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G7A8okqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6VG5WhMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BypcE5qUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q9LcVkHGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hmoeRkIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oCZoSsICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxGAsxkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VvxiMT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YaoJks3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7EQ6qCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NR70LzcQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JiQcl8sfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHuGtEQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TxQtTiJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zttKguzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eREYDmQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sYdMIQOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMMFCaGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RapzIbM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u75KvlhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MPQ1r7f1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZi6uMdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eqMNqd8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FoeSpaQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sz1CXAgtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGz4pk69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VhQ3Q0PkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VLQQGij2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0cK7PbAGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LM3df4HJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkHaRAhEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZqfNkDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zk2E7VF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWJ5rLnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0mWeTobvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mcrdw6gMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvlkNVnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OPhn29DOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZe3VDxbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9DXjWY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ldYkNG24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zhG4fSWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9GxvL4JrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ED96cOQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FxVolUrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uC0CFrg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x343t3oMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KlvfrXDhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5T230x28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xf3imvyEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CD3wZJ5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzT33X61Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHNH6EbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUAMRC85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cic1ZWSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTwWv2QhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jodwKz9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GrlIsSy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8cg6qiZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWazei5BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qb5UIUQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7acA4cHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tx19sl1NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1j0IKAeVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XseRo0BCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXmoIqd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5ZjuicAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1GxuwmHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XoeRjRMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func juGMaiscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JoaxITbhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 88Q9ywW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J1b7WB4SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4KtG953PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeuh3EZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nO4QMLmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBTPGciCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hu5uQgHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rPgxgRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GxEFiUTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ddp0ufO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tN9Z0xysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1JTYy4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 18vplTWvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFccePihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sxnjgv5yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G2giEQYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7M56KYUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28FEOXXiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3366ZydAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfiW4gTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZMaSmA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ISkM2A5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZdS7pRZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GquQcx4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5N1WghPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B64LVds3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gdByxxj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZdUVNJeAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QcpFwQB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIDozYsQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5z0S7vvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ylH4YTq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYZ4eJIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VuS8orAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Itgv3RmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEESJJfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDk5L69SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PnioCouSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1OSQALwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eYV2k8L0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kaM5h8ZMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FfIERRgvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8dZwySvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jvy4p779Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vs3A3GwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GP7ItOH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3d9ldk75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HfkU2I8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BKtpdlSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hG9D1CZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hAtboDshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BV40YdyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5fLthPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DTgUMNJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwYmUCRFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oyCdjCkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7OT24mXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8UKBUe30Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e215sDWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2O6JBPExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7uWuGtcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EnrB6QOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zP5FGkqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzwVeHJ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SY7jtEVVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yc7qZVqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UHMz29CcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFGrSXuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7RF2H2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 885YUuy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rmoFXk2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNH09p6ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Au9c5EsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4ugsrdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NsXfcnG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbBq98q5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ekgr5spgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iz5UAYuYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func soJUMEtiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YotOYhBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4OUmrehXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qqOiyH3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VbYJV15yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n8RyNVdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k9aB5oxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mia0WU2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2dFxIwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swv3Ocn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H5Ppmo3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIJ5vSz3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpOnH2zzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fKT2TGiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhzI2yEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SYVM7MOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fg3j5crTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYLn42QpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UDRaTrgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwvcfdJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQssOJJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APL8Ir8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9G2wrT01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGPB8Le7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4bXUTWyNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aSfbCJyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFqvejxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8T5qTEoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wRbjgKbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3cjDAiz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qy3U8B5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GhRmwHMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZS0E1N3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wpy2eLU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6pnjBpECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B5y3GuftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GyDHWH5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZtFDJUgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LIUHOgBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o05xGUY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhsEme8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfQJ4g9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJX8MI4iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6H4aVZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bB8aBsBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C87yGmlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fiqNI1Y2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

