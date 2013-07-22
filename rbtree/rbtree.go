package rbtree

import "log"

const (
    red   = iota  // c0 == 0
    black = iota  // c1 == 1
)

type rbtree struct {
    color int
    right *rbtree
    left *rbtree
    value *int
    father *rbtree
}

func NewRBTree() *rbtree {
    return &rbtree{black, nil, nil, nil, nil}
}

func (t *rbtree) add(elem int) {
    node := t.containsNode(elem)
    if node.value == nil {
        node.value = &elem
    } else if *node.value != elem {
        color := black
        if node.color == black {
            color = red
        }
        if elem < *node.value {
            node.left = &rbtree{color, nil, nil, &elem, node}
        } else {
            node.right = &rbtree{color, nil, nil, &elem, node}
        }
    }
}

func (t *rbtree) remove(elem int) bool {
    note := t.containsNode(elem)
    if node.value != nil && *node.value == elem && node.father != nil {
        isright := (*node.father).right == node
        if node.left == nil && node.right == nil { // On supprime une feuille
            if isright {
                (*node.father).right = nil
            } else {
                (*node.father).left = nil
            }
        } else if node.left == nil ^ node.right == nil {
            singlechild := node.left
            if singlechild == nil {
                singlechild = node.right
            }
            if isright {
                (*node.father).right = singlechild
            } else {
                (*node.father).left = singlechild
            }
        } else {
            
        }
    }
    return false;
}

func (t *rbtree) size() int {
    if t.value == nil {
        return 0
    }
    left, right := 0, 0
    if t.left != nil {
        left = t.left.size()
    }
    if t.right != nil {
        right = t.right.size()
    }
    return 1 + right + left
}

func (t *rbtree) contains(elem int) bool {
    node := t.containsNode(elem)
    log.Printf("elem = %d, *node.value = %d\n", elem, *node.value)
    return elem == *node.value
}

func (t *rbtree) containsNode(elem int) *rbtree {
    node := t
    if t.value != nil && *t.value != elem {
        if t.left != nil && elem < *t.value {
            node = t.left.containsNode(elem)
        } else if t.right != nil {
            node = t.right.containsNode(elem)
        }
    }
    return node
}
