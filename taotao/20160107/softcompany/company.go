package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	om := &OrderMarket{}
	om.Init() //初始化一个订单市场
	cm := &CoderMarket{}
	cm.Init() //初始化一个程序员市场

	var coder GoCoder
	var err1 error
	flag := false
	totalMoney := 100.0 //单位：w
	for {
		if totalMoney > 1000.0 || totalMoney < 0 {
			fmt.Println("创业结束. totalMoney:%v", totalMoney)
			return
		}
		fmt.Println("新一轮工作开始\n")

		var ordertime int
		i := 0
	Here:
		order, err := om.GetOrder()
		if err != nil {
			fmt.Printf("第 ％d 月找订单失败\n", i)
			totalMoney -= 0.1 //无订单时，每月消耗1000元
			if i == 2 {
				fmt.Println("连续三个月无订单，创业失败\n")
				return
			}
			i++
			goto Here
		} else {
			ordertime = order.OrderTime

		}

		//招收程序员
		ordertime -= rand.Int() % 3 //假设最多三个月可招收一名程序员，在订单中耗时
		if !flag {
			coder, err1 = cm.GetCoder()
			if err1 != nil {
				fmt.Println("没有找到程序员，退回订单后继续\n")
				om.ReturnOrder(order)
				totalMoney -= order.OrderMoney * 2
				continue
			}
		}

		sf, err := coder.WriteCode(order)
		if err != nil { //编码失败
			fmt.Printf("编码超时失败\n")
			totalMoney -= order.OrderMoney * 2
			totalMoney -= coder.Salary
			continue
		}
		om.DeliverOrder(order, sf)
		totalMoney += order.OrderMoney

		//支付程序员工资
		moneyToCoder := coder.Salary + order.OrderMoney/10.0
		coder.GetMoney(moneyToCoder)
		totalMoney -= moneyToCoder

		if rand.Int()%100 > 50 { //程序员50%的概率留下
			flag = true

		}

	}
}

type Order struct {
	OrderName  string
	OrderTime  int     //月
	OrderMoney float64 //万

}

type OrderMarket struct {
	orderPool []Order
}

func (om *OrderMarket) Init() {
	rand.Seed(time.Now().Unix())
	var orderPool []Order
	for i := 0; i < 10000; i++ {
		order := Order{
			OrderMoney: float64(rand.Int()%50 + 50), //订单为50w+
			OrderName:  "order:" + strconv.Itoa(i+1),
			OrderTime:  rand.Int() % 24,
		}
		orderPool = append(orderPool, order)
	}

	om.orderPool = orderPool
}

func (om *OrderMarket) GetOrder() (order Order, err error) {
	if rand.Int()%100 < 20 {
		return order, errors.New("no order")
	}

	order = om.orderPool[rand.Int()%len(om.orderPool)]
	fmt.Printf("get an order:%v money:%vw  ordertime:%vmonth \n", order.OrderName, order.OrderMoney, order.OrderTime)
	return
}

func (om *OrderMarket) ReturnOrder(order Order) {
	om.orderPool = append(om.orderPool, order)
	fmt.Println("ret an no complete order:", order.OrderName)
}

func (om *OrderMarket) DeliverOrder(order Order, sf SoftWare) {
	fmt.Println("deliver order success. software info:", sf.Name) // 打印软件名字
}

type SoftWare struct {
	Name string
}

type CoderMarket struct {
	coderPool []GoCoder
}

type Coder interface {
	WriteCode(order Order) (SoftWare, error)
	GetMoney(money float64)
}

type GoCoder struct {
	Name   string
	Salary float64
}

func (gc *GoCoder) WriteCode(order Order) (sf SoftWare, err error) {

	fmt.Printf("begin to write code for %v\n", order.OrderName)
	var i int = rand.Int() % (order.OrderTime + 3)
	if i > order.OrderTime {
		return sf, errors.New("too lang time to code")
	}
	time.Sleep(time.Second * 2) //睡眠表示在写代码，用时在订单时间＋3个月进行随机数，若超过订单时间，则返回开发失败
	fmt.Println("write code ok.")

	sf.Name = fmt.Sprintf("hello, world! by %v", gc.Name)
	return
}

func (gc *GoCoder) GetMoney(money float64) {
	fmt.Printf("%v get money:%vw\n", gc.Name, money)
}

func (cm *CoderMarket) Init() {

	rand.Seed(time.Now().Unix())
	var coderpool []GoCoder
	for i := 0; i < 1000; i++ {
		coder := GoCoder{
			Name:   "coder:" + strconv.Itoa(i),
			Salary: float64(rand.Int()%5) + rand.Float64(), //工资收入随机，工资收入上限5w

		}
		coderpool = append(coderpool, coder)
	}

	cm.coderPool = coderpool
}

func (cm *CoderMarket) GetCoder() (coder GoCoder, err error) {
	index := rand.Int() % len(cm.coderPool)
	coder = cm.coderPool[index]
	fmt.Printf("get a coder:%v", coder.Name)
	return
}
