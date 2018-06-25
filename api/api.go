package api

import (
    "io"
    "os"
    "blocks/block"
)


func Fragment(key string) (nBlocks uint, err error) {
    file, err := os.Open(key)
    if err != nil {
        return 0, err
    }
    defer file.Close()

    for {
        dataRead := make([]byte, block.BlockSize)
        bytesRead, err := file.Read(dataRead)
        if err != nil {
            if err != io.EOF {
                return 0, err
            }
            break
        }

        newID := block.GenerateID(key, nBlocks)
        newBlock := block.New(newID, dataRead[0:bytesRead])
        newBlock.Save()
        nBlocks++
    }
    return nBlocks, nil
}

func Collect(key string, size uint) (nBlocks uint, blockIDs []string) {
    nBlocks = block.CalcBlocks(size)
    blockIDs = make([]string, nBlocks)
    for id := uint(0); id < nBlocks; id++ {
        blockIDs[id] = block.GenerateID(key, id)
    }
    return nBlocks, blockIDs
}

func Combine(blockIDs []string, key string) (err error) {
    newFile, err := os.Create(key)
    if err != nil {
        return err
    }
    defer newFile.Close()

    for _, blockID := range blockIDs {
        block, err := block.Get(blockID)
        if err != nil {
            return err
        }
        _, err = newFile.Write(block.Data())
        if err != nil {
            return err
        }
    }
    return nil
}
