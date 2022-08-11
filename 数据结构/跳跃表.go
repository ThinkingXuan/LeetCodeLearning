package main

import (
	"fmt"
	"math/rand"
	"time"
)

// https://lotabout.me/2018/skip-list/
//https://blog.csdn.net/crayhl/article/details/108239128
// https://blog.csdn.net/qq_35475714/article/details/115555051
const (
	maxHeight = 64
	rate      = 0.25
	name      = "zksip"
)

type zSkipList struct {
	head   *zSkipListNode
	tail   *zSkipListNode
	length int
	level  int
}

type zSkipLevel struct {
	forward *zSkipListNode
	span    int
}

type zSkipListNode struct {
	ele      string
	score    float32
	backward *zSkipListNode
	level    []zSkipLevel
}

var zlist *zSkipList

func createNode(num int) *zSkipListNode {
	node := new(zSkipListNode)
	node.level = make([]zSkipLevel, num)
	return node
}

func createZSkipList() {
	zlist = new(zSkipList)
	zlist.head = createNode(maxHeight)
	zlist.level = 1
	zlist.tail = nil
}

//随机获取层数
func getLevel() int {
	level := 1
	//return 3
	for i := 0; i < maxHeight; i++ {
		num := rand.Intn(maxHeight)
		if num < rate*maxHeight {
			return level
		}
		level += 1
	}
	return level
}

func insert(ele string, score float32) {
	//创建该节点
	level := getLevel()
	node := createNode(level)
	fmt.Printf("当前层数%d,值=%.1f\n", level, score)
	node.ele = ele
	node.score = score

	//update数组记录插入节点的前一个节点bef（bef节点为需要修改的节点）
	update := make([]*zSkipLevel, zlist.level)
	//rank数组记录从head到bef节点的累计跨度
	rank := make([]int, zlist.level)
	//第一次插入
	if zlist.length == 0 {
		for i := level - 1; i >= 0; i-- {
			zlist.head.level[i].forward = node
			zlist.head.level[i].span = 1
		}
		zlist.tail = node
		zlist.level = level
	} else {
		p := zlist.head
		rankSum := 0
		//从最高层开始找，寻找每一层的bef节点
		for i, k := zlist.level-1, 0; i >= 0; {
			//当前节点的next节点为空，或者next节点的score值比新节点值大，则
			//当前节点即为bef节点
			if (p.level[i].forward == nil) || (score < p.level[i].forward.score) {
				update[i] = &p.level[i]
				rank[k] = rankSum
				i--
				k++
			} else {
				rankSum += p.level[i].span
				//当前层还未寻找到bef节点，当前层rank值继续累加
				p = p.level[i].forward
			}
		}
		temp := rank[len(rank)-1]
		var rankIndex int
		for i := 0; i < level && i < len(update); i++ {
			rankIndex = len(rank) - 1 - i
			//更新节点的next节点
			node.level[i].forward = update[i].forward
			//更新节点的span
			node.level[i].span = update[i].span - (temp - rank[rankIndex])

			update[i].forward = node
			update[i].span = temp - rank[rankIndex] + 1
		}

		//当新建节点层数大于zskip的最大层数，更新head到node节点的跨度
		for i := level; i > zlist.level; i-- {
			zlist.head.level[i-1].forward = node
			zlist.head.level[i-1].span = temp + 1
		}
		//当新建节点层数小于zskip的最大层数，update数组中level~zlist.level层数节点的跨度+1
		for i := level; i < zlist.level; i++ {
			update[i].span++
		}
	}

	if zlist.level < level {
		zlist.level = level
	}
	zlist.length += 1
}

func printLevelDetail(head zSkipLevel, layer int) {

	headDesc := fmt.Sprintf("%d:[head],%d -> ", layer+1, head.span)
	desc := fmt.Sprintf("%15s", headDesc)
	for i := 0; i < head.span-1; i++ {
		desc += fmt.Sprintf("%15s", "-> ")
	}
	p := head.forward
	for p != nil {
		item := fmt.Sprintf("[%s,%.1f],%d -> ", p.ele, p.score, p.level[layer].span)
		itemSpan := ""
		for i := 0; i < p.level[layer].span-1; i++ {
			itemSpan += fmt.Sprintf("%15s", "-> ")
		}
		desc += fmt.Sprintf("%15s", item) + itemSpan
		p = p.level[layer].forward
	}
	fmt.Println(desc)
}

//结构化打印跳表数据结构
func printZskipDetail() {
	for i := zlist.level - 1; i >= 0; i-- {
		printLevelDetail(zlist.head.level[i], i)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//创建空跳表
	createZSkipList()
	//插入元素
	insert("A", 10)
	insert("B", 12)
	insert("C", 13)
	insert("D", 11)
	insert("E", 12.5)
	//格式化打印
	printZskipDetail()
}
