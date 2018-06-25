package block

import (
    "io/ioutil"
    "os"
    "fmt"
)

const BlockSize = 4194304

type Block struct {
    id string
    data []byte
}

func New(id string, data []byte) Block {
    return Block{id:id, data:data}
}

func Get(id string) (Block, error) {
    data, err := ioutil.ReadFile(id)
    if err != nil {
        return New(id, data), err
    }

    return New(id, data), nil
}

func (b Block) Save() error {
    file, err := os.Create(b.id)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.Write(b.data)
    if err != nil {
        return err
    }

    return nil
}

func (b Block) Data() []byte {
    return b.data
}

func CalcBlocks(fileSize uint) uint {
    if fileSize % BlockSize  == 0 {
        return fileSize / BlockSize
    }
    return fileSize / BlockSize + 1
}

func GenerateID(key string, id uint) string {
    return fmt.Sprintf("%s%d", key, id)
}
