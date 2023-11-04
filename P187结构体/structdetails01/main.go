package main

import "fmt"

type Point struct{
	x int
	y int
}

type Rect struct{
	leftUp, rightDown Point
}

// 结构体
type Rect2 struct{
	leftUp, rightDown *Point
}

func main(){
	r1 := Rect{Point{1, 2}, Point{3, 4}}

	// 打印地址
	fmt.Printf("r1.leftUp.x    地址 = %p\n", &r1.leftUp.x)
	fmt.Printf("r1.leftUp.y    地址 = %p\n", &r1.leftUp.y)
	fmt.Printf("r1.rightDown.x 地址 = %p\n", &r1.rightDown.x)
	fmt.Printf("r1.rightDown.x 地址 = %p\n", &r1.rightDown.y)
	fmt.Println()


	// r2是有两个 *Point类型， 这两个 *Point类型的本身地址也是连续的
	// 但是他们指向的地址不一定是连续的

	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}

	// 他们本身的地址，本身的地址是连续的
	fmt.Printf("r2.leftUp    本身 地址 = %p\n", &r2.leftUp)
	fmt.Printf("r2.rightDown 本身 地址 = %p\n", &r2.rightDown)
	fmt.Println()

	// 他们指向的地址,指向的地址不一定是连续的
	fmt.Printf("r2.leftUp    指向 地址 = %p\n", r2.leftUp)
	fmt.Printf("r2.rightDown 指向 地址 = %p\n", r2.rightDown)
	
}