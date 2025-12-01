// Реализуйте структуру данных LRU Cache с фиксированной емкостью. 
// Когда cache заполнен и нужно добавить новый элемент, удаляется самый давно использованный элемент.

// Требования:

// Используйте dummy nodes для головы и хвоста двусвязного списка

// Все операции должны работать за O(1)

// Реализуйте методы:

// Get(key int) int - возвращает значение по ключу, обновляет порядок использования

// Put(key, value int) - добавляет или обновляет пару ключ-значение

// При переполнении удаляйте самый старый элемент

package main

import "fmt"

type Node struct {
	key int
	value int
	prev *Node
	next *Node
}

type LRUCache struct {
	cap int
	cache map[int]*Node
	head *Node
	tail *Node
}

func Constructor(cap int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		cap: cap,
		head: head,
		tail: tail,
		cache: make(map[int]*Node),
	}
}

func (l *LRUCache) RemoveNode(node *Node) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}

func (l *LRUCache) AddNode(node *Node) {
	head := l.head
	node.prev = head
	node.next = head.next
	head.next = node
	node.next.prev = node
}

func (l *LRUCache) MoveToHead(node *Node) {
	l.RemoveNode(node)
	l.AddNode(node)
}

func (l *LRUCache) Get(key int) int {
	res, ok := l.cache[key]
	if ok {
		l.MoveToHead(res)
		return res.value
	}
	return -1
}

func (l *LRUCache) RemoveTail() *Node{
	node := l.tail.prev
	l.RemoveNode(node)
	return node
}

func (l *LRUCache) Put(key, value int) {
	if node, ok := l.cache[key]; ok {
		node.value = value
		l.MoveToHead(node)
		return
	}
	nodeNew := &Node{key:key, value: value}
	l.cache[key] = nodeNew
	l.AddNode(nodeNew)
	if l.cap < len(l.cache) {
		tail := l.RemoveTail()
		delete(l.cache, tail.key)
	}
}

func (l *LRUCache) String() string{
	res := "LRU Cache(newest -> oldest)"
	start := l.head.next
	for start != l.tail {
		res+=fmt.Sprintf("(%d:%d)", start.key, start.value )
		start = start.next
	}
	return res
}

func (l *LRUCache) Size() int {
	return len(l.cache)
}

func main() {
  fmt.Println("=== LRU Cache с Dummy Nodes ===")

  lru := Constructor(3)

  fmt.Println("Инициализирован пустой cache:")
  fmt.Println(lru.String())

  operations := []struct {
    op   string
    key  int
    val  int
    desc string
  }{
    {"PUT", 1, 100, "Добавляем (1:100)"},
    {"PUT", 2, 200, "Добавляем (2:200)"},
    {"PUT", 3, 300, "Добавляем (3:300)"},
    {"GET", 2, 0, "Читаем ключ 2 (должен стать самым новым)"},
    {"PUT", 4, 400, "Добавляем (4:400) - должен вытолкнуть 1"},
    {"GET", 1, 0, "Пытаемся прочитать вытолкнутый ключ 1"},
    {"PUT", 5, 500, "Добавляем (5:500) - должен вытолкнуть 3"},
  }

  for _, op := range operations {
    fmt.Printf("\n--- %s ---\n", op.desc)

    switch op.op {
    case "PUT":
      lru.Put(op.key, op.val)
      fmt.Printf("PUT(%d, %d)\n", op.key, op.val)
    case "GET":
      result := lru.Get(op.key)
      fmt.Printf("GET(%d) = %d\n", op.key, result)
    }

    fmt.Printf("Состояние: %s\n", lru.String())
    fmt.Printf("Размер: %d/%d\n", lru.Size(), lru.cap)
  }

  fmt.Println("\n=== Тест обновления существующего ключа ===")
  lru2 := Constructor(2)

  lru2.Put(1, 1)
  lru2.Put(2, 2)
  fmt.Println("После добавления двух элементов:")
  fmt.Println(lru2.String())

  lru2.Put(1, 10)
  fmt.Println("После обновления ключа 1:")
  fmt.Println(lru2.String())

  lru2.Put(3, 3)
  fmt.Println("После добавления ключа 3:")
  fmt.Println(lru2.String())
}