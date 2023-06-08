package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "github.com/beevik/etree"
)

var (
    // key : filename-id
    ConfigCache = make(map[string]LineMap)
    // key : filename-groupid
    GroupCache = make(map[string]ItemList)
)

type LineMap map[string]string
type ItemList []LineMap

func main() {
    matches, _ := filepath.Glob(GetCWD() + "/config/*.xml")
    for _, file := range matches {
        err := parse(file)
        if err != nil {
            sprintf := fmt.Sprintf("error xml config file %s %s ", file, err)
            panic(sprintf)
        }
    }
    println("GroupCache len ", len(GroupCache))
    println("ConfigCache len ", len(ConfigCache))
}

func parse(filepath string) error {
    index := strings.LastIndex(filepath, "/")
    filename := filepath[index+1 : len(filepath)-4]

    doc := etree.NewDocument()
    err := doc.ReadFromFile(filepath)
    if err != nil {
        fmt.Println("read err document xml", filepath)
        return err
    }
    root := doc.Root()
    ele := root
    if root.Tag != "Group" {
        for _, ele = range root.ChildElements() {
            parseInside(filename, ele)
        }
    } else {
        parseInside(filename, ele)
    }
    return nil
}

func parseInside(filename string, ele *etree.Element) {
    switch ele.Tag {
    case "Group":
        groupName := ele.Attr[0].Value
        l := ItemList{}
        for _, line := range ele.ChildElements() {
            m := LineMap{}
            for _, item := range line.Attr {
                m[item.Key] = item.Value
            }
            ConfigCache[filename+"-"+m["id"]] = m
            l = append(l, m)
        }
        GroupCache[filename+"-"+groupName] = l
    case "ItemSpec":
        m := LineMap{}
        for _, item := range ele.Attr {
            m[item.Key] = item.Value
        }
        ConfigCache[filename+"-"+m["id"]] = m
    }
}

func GetCWD() string {
    cwd, _ := os.Getwd()
    return cwd
}
