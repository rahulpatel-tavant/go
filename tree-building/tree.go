package tree
import (
  "sort"
  "fmt"
)

type Node struct {
  ID int
  Children []*Node
}

type Record struct {
  ID int
  Parent int
}

type ByID []Record
func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }

func Build(records []Record) (*Node, error) {
  if len(records) <= 0 {
    return nil, nil
  }
  sort.Sort(ByID(records))
  nodes := make([]*Node, len(records))

  for index, record := range records {
    nodes[index] = &Node{ID: record.ID}
    if index == 0 && (record.ID != 0 || record.Parent != 0) {
      return nil, fmt.Errorf("InValid record")
    } else if index == 0 {
      continue
    } else if index != 0 && (index != record.ID || record.ID <= record.Parent) {
      return nil, fmt.Errorf("InValid record")
    }

    parent := nodes[record.Parent]
    if parent != nil {
      parent.Children = append(parent.Children, nodes[index])
    }
  }
  return nodes[0], nil
}

