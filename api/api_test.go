package api

import (
    "testing"
)

func TestFragment(t *testing.T){
    tables := []struct {
		key string
		nBlocks uint
		err error
	}{
		{"../data/test1", 10, nil},
		{"../data/test2", 5, nil},
		{"../data/test3", 11, nil},
		{"../data/test4", 15, nil},
		{"../data/test5", 1, nil},
	}

    for _, table := range tables {
		nBlocks, err := Fragment(table.key)
		if nBlocks != table.nBlocks {
			t.Errorf("Fragmentation of file [%s] was incorrect, got: %d blocks, want: %d blocks.", table.key, nBlocks, table.nBlocks)
		}
        if err != nil {
            t.Errorf("Error encountered")
        }
	}
}

func TestCollect(t *testing.T){
    tables := []struct {
        key string
        size uint
        nBlocks uint
        blocks []string
    }{
        {"../data/test1", 41458199, 10, []string{"../data/test10", "../data/test11", "../data/test12", "../data/test13", "../data/test14", "../data/test15", "../data/test16", "../data/test17", "../data/test18", "../data/test19"}},
        {"../data/test2", 18398648, 5, []string{"../data/test20", "../data/test21", "../data/test22", "../data/test23", "../data/test24"}},
        {"../data/test3", 44334495,  11, []string{"../data/test30", "../data/test31", "../data/test32", "../data/test33", "../data/test34", "../data/test35", "../data/test36", "../data/test37", "../data/test38", "../data/test39", "../data/test310"}},
        {"../data/test4", 60738203, 15, []string{"../data/test40", "../data/test41", "../data/test42", "../data/test43", "../data/test44", "../data/test45", "../data/test46", "../data/test47", "../data/test48", "../data/test49", "../data/test410", "../data/test411", "../data/test412", "../data/test413", "../data/test414"}},
        {"../data/test5", 1440054,   1, []string{"../data/test50"}},
    }

    for _, table := range tables {
		nBlocks, blocks := Collect(table.key, table.size)
		if nBlocks != table.nBlocks {
			t.Errorf("Number of blocks in file [%s] was incorrect, got: %d blocks, want: %d blocks.", table.key, nBlocks, table.nBlocks)
		}
        for i, block := range blocks {
            if block != table.blocks[i] {
                t.Errorf("Blocks do not match, got: %s, want: %s", block, table.blocks[i])
            }
        }
	}
}
