package main

import (
    "os"
    "flag"
    "io"
    "bufio"
    "fmt"
    "time"
    "strconv"
    "encoding/csv"
)

var rdfPath = flag.String("path", "", "Specify rdf data path")

const SEED = 0xc70f6907

const CAP = 100000

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

    lVRecord := make([]string, 2)
    rVRecord := make([]string, 2)
    eRecord := make([]string, 3)
    exists := make(map[int64]bool)   
    for {
        if lineNum % 100000 == 0 {
            fmt.Printf("hava read lines: %d\n", lineNum)
        }

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
        
        if len(exists) >= CAP {
            exists = make(map[int64]bool)
        }

        vid := line[0]
        lVRecord[0] = vid
        lVRecord[1] = line[0]
        if _, ok := exists[vid]; !ok {
            exists[vid] = true
            vWriter.Write(lVRecord)
        }

        
        vid = line[2]
        rVRecord[0] = vid
        rVRecord[1] = line[2]
        if _, ok := exists[vid]; !ok {
            exists[vid] = true
            vWriter.Write(rVRecord)
        }
    
        eRecord[0] = lVRecord[0]
        eRecord[1] = rVRecord[0]
        eRecord[2] = line[1]
        eWriter.Write(eRecord)
    }

    return nil
}
