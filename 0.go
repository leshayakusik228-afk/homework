package main

func main() {
	var res1 int64
	//var res2 int64
	var res3 int64
	var mul1 int64
	//	var mul2 int64
	var mul3 int64
	res1, mul1 = function(128, 128)
	//res2, mul2 = function(256, 256)
	res3, mul3 = function(512, 512)
	println(res1, mul1, res3, mul3)
}
func function(x int64, y int64) (int64, int64) {
	sum := x + y
	mul := x * y
	return sum, mul

}
