package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

const (
	stockDSN = "root:passwd@tcp(127.0.0.1:3307)/shop_product_stock"
	orderDSN = "root:passwd@tcp(127.0.0.1:3308)/shop_order"
	bagDSN   = "root:passwd@tcp(127.0.0.1:3309)/user_money_bag"
)

func main() {

	// 库存的连接
	stockDb, err := sql.Open("mysql", stockDSN)
	if err != nil {
		panic(err.Error())
	}

	err = stockDb.Ping()
	if err != nil {
		fmt.Println(err)
	}

	defer stockDb.Close()

	//订单的连接
	orderDb, err := sql.Open("mysql", orderDSN)
	if err != nil {
		panic(err.Error())
	}
	defer orderDb.Close()

	//钱包的连接
	moneyDb, err := sql.Open("mysql", bagDSN)
	if err != nil {
		panic(err.Error())
	}
	defer moneyDb.Close()

	// 生成xid(如果在同一个数据库，子事务不能使用相同xid)
	xid := strconv.FormatInt(time.Now().UnixMilli(), 10)
	fmt.Println(xid)
	//如果后续执行过程有报错，那么回滚所有子事务
	defer func() {
		if err := recover(); err != nil {

			stockDb.Exec(fmt.Sprintf("XA ROLLBACK '%s'", xid))
			orderDb.Exec(fmt.Sprintf("XA ROLLBACK '%s'", xid))
			moneyDb.Exec(fmt.Sprintf("XA ROLLBACK '%s'", xid))

			fmt.Println("事务回滚")
		}
	}()

	// 第一阶段 Prepare

	// stage 1
	//库存 子事务启动

	{
		if ret, err := stockDb.Exec(fmt.Sprintf("XA START '%s'", xid)); err != nil {
			fmt.Printf("stockDb XA START '%s'\r\n", xid)
			fmt.Println(err.Error())
			panic(err.Error())
		} else {
			fmt.Println("###1:  ", ret)
		}

		//扣除库存，这里省略了数据行锁操作
		if ret, err := stockDb.Exec("update product_stock set stock=stock-1 where id =1"); err != nil {
			fmt.Printf("stockDb 扣库存 %s\r\n", xid)
			fmt.Println(err.Error())
			panic(err.Error())
		} else {
			fmt.Println("###2:  ", ret)
		}
		//事务执行结束
		if ret, err := stockDb.Exec(fmt.Sprintf("XA END '%s'", xid)); err != nil {
			fmt.Printf("stockD XA END '%s'\r\n", xid)
			fmt.Println(err.Error())
			panic(err.Error())
		} else {
			fmt.Println("###3:  ", ret)
		}

		//设置库存任务为Prepared状态
		if ret, err := stockDb.Exec(fmt.Sprintf("XA PREPARE '%s'", xid)); err != nil {
			fmt.Printf("stockD XA PREPARE %s\r\n", xid)
			fmt.Println(err.Error())
			panic(err.Error())
		} else {
			fmt.Println("###4:  ", ret)
		}
	}

	// stage 2
	// 订单 子事务启动

	{
		if ret, err := orderDb.Exec(fmt.Sprintf("XA START '%s'", xid)); err != nil {
			fmt.Printf("orderDb XA START '%s'\r\n", xid)
			fmt.Println(err.Error())
			panic(err.Error())
		} else {
			fmt.Println("###2-1:  ", ret)
		}

		//创建订单

		if ret, err := orderDb.Exec("insert shop_order(pid) value (?)", 2); err != nil {
			fmt.Printf("insert shop_order(id,pid) value (1,2)\r\n")
			fmt.Println(err.Error())
			panic(err.Error())
		} else {
			fmt.Println("###2-2:  ", ret)
		}
		//事务执行结束
		if ret, err := orderDb.Exec(fmt.Sprintf("XA END '%s'", xid)); err != nil {
			fmt.Printf("XA END '%s'\r\n", xid)
			fmt.Println(err.Error())
			panic(err.Error())
		} else {
			fmt.Println("###2-3:  ", ret)
		}

		//设置任务为Prepared状态
		if ret, err := orderDb.Exec(fmt.Sprintf("XA PREPARE '%s'", xid)); err != nil {
			fmt.Printf("XA PREPARE '%s'\r\n", xid)
			fmt.Println(err.Error())
			panic(err.Error())
		} else {
			fmt.Println("###2-4:  ", ret)
		}
	}

	// stage 3
	// 钱包 子事务启动
	{
		if _, err = moneyDb.Exec(fmt.Sprintf("XA START '%s'", xid)); err != nil {
			panic(err.Error())
		}
		//扣减用户账户现金，这里省略了数据行锁操作
		if _, err = moneyDb.Exec("update user_money_bag set money=money-47 where id =2"); err != nil {
			panic(err.Error())
		}
		//事务执行结束
		if _, err = moneyDb.Exec(fmt.Sprintf("XA END '%s'", xid)); err != nil {
			panic(err.Error())
		}
		//设置任务为Prepared状态
		if _, err = moneyDb.Exec(fmt.Sprintf("XA PREPARE '%s'", xid)); err != nil {
			panic(err.Error())
		}
	}

	// 第二阶段 commit
	{
		// 在这时，如果链接断开、Prepared状态的XA事务仍旧在MySQL存在
		//任意一个链接调用XA RECOVER都能够看到这三个没有最终提交的事务
		//--------
		//第二阶段 运行到这里没有任何问题
		//那么执行 commit
		//--------
		if _, err = stockDb.Exec(fmt.Sprintf("XA COMMIT '%s'", xid)); err != nil {
			panic(err.Error())
		}
		if _, err = orderDb.Exec(fmt.Sprintf("XA COMMIT '%s'", xid)); err != nil {
			panic(err.Error())
		}
		if _, err = moneyDb.Exec(fmt.Sprintf("XA COMMIT '%s'", xid)); err != nil {
			panic(err.Error())
		}
		//到这里全部流程完毕

		fmt.Println("分布式事务XA")
	}
}
