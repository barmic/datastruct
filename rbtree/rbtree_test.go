package rbtree

import "testing"

func TestInserttionRbtree(t *testing.T) {
    set := NewRBTree()
    elem := 12
    set.add(elem)
    if set.size() != 1 || !set.contains(elem) {
        t.Errorf("L'insertion ne fonctionne pas")
    }
}

func TestContainsRbtree(t *testing.T) {
    set := NewRBTree()
    set.add(12)
    set.add(13)
    set.add(14)
    set.add(15)
    if set.size() != 4 || !set.contains(12) || !set.contains(13) || !set.contains(14) || !set.contains(15) || set.contains(24) || set.contains(2) {
        t.Errorf("Le contains ne fonctionne pas")
    }
}

func TestNotContainsRbtree(t *testing.T) {
    set := NewRBTree()
    elem := 12
    set.add(elem)
    if set.size() != 1 || set.contains(13) {
        t.Errorf("Le not contains ne fonctionne pas")
    }
}

func TestSuppressionRbtree(t *testing.T) {
    tree := NewRBTree()
    elem := 12
    elem2 := 13
    elem3 := 14
    tree.add(elem)
    tree.add(elem2)
    tree.add(elem3)
    tree.remove(elem)
    if tree.size() != 2 || tree.contains(elem) {
        t.Errorf("La suppression ne fonctionne pas")
    }
}

func TestSuppressionDupplicatRbtree(t *testing.T) {
    tree := NewRBTree()
    elem := 12
    tree.add(elem)
    tree.add(elem)
    if tree.size() != 1 || !tree.contains(elem) {
        t.Errorf("La suppression des doublons ne fonctionne pas")
    }
}
