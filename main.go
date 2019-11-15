package main

// #cgo CFLAGS: -I${SRCDIR}/hash
// #cgo LDFLAGS: -lstdc++ -L${SRCDIR}/hash -lhash
// #include "bridge.h"
import "C"

import (
    "os"
    "flag"
    "io"
    "bufio"
    "fmt"
    "time"
    "encoding/csv"
)

var rdfPath = flag.String("path", "", "Specify rdf data path")

func main() {
    flag.Parse()
  
    if err := Read(*rdfPath); err != nil {
        panic(err)
    }
}

func Read(rdfPath string) error {
    srcFile, err := os.Open(rdfPath)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    var vFilePath = "./vertex.csv"
    vFile, err := os.Create(vFilePath)
    if err != nil {
        return err
    }
    defer vFile.Close()

    var eFilePath = "./edge.csv"
    eFile, err := os.Create(eFilePath)
    if err != nil {
        return err
    }
    defer eFile.Close()

    reader := csv.NewReader(bufio.NewReader(srcFile))
    vWriter := csv.NewWriter(bufio.NewWriter(vFile))
    eWriter := csv.NewWriter(bufio.NewWriter(eFile))
    defer vWriter.Flush()
    defer eWriter.Flush()

    now := time.Now()

    defer func() {
        fmt.Printf("Finish convert rdf data, consume time: %.2fs\n", time.Since(now).Seconds())
    }()

    lineNum, numErrorLines := 0, 0

    lastLV, lastRV := "rdf-converter", "rdf-converter"    
    for {
        //if lineNum % 100000 == 0 {
        //    fmt.Printf("hava read lines: %d\n", lineNum)
        //}

        line, err := reader.Read()
        if err == io.EOF {
            fmt.Printf("totalLines: %d, errorLines: %d\n", lineNum, numErrorLines)
            break
        }

        lineNum++

        if err != nil {
            numErrorLines++
            continue
        }

        if len(line) != 3 {
            numErrorLines++
            continue
        }

        vL := make([]string, 2)
        vL[0] = C.GoString(C.getHash(C.CString(line[0])))
        vL[1] = line[0]
        if line[0] != lastLV {
            vWriter.Write(vL)
            lastLV = line[0]
        }

        vR := make([]string, 2)
        vR[0] = C.GoString(C.getHash(C.CString(line[2])))
        vR[1] = line[2]
        if line[2] != lastRV {
            vWriter.Write(vR)
            lastRV = line[2]
        }
        
        E := make([]string, 3)
        E[0] = vL[0]
        E[1] = vR[0]
        E[2] = line[1]
        eWriter.Write(E)
    }

    return nil
}
