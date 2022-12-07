package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(x int, y int) int  { return x + y }
func sub(x int, y int) int  { return x - y }
func mult(x int, y int) int { return x * y }
func dev(x int, y int) int  { return x / y }

func rim_number(rim_dig int) string {
	var a string
	var b string
	rim_diga := (rim_dig / 10)
	rim_digaa := (rim_diga * 10)
	rim_digb := (rim_dig - (rim_diga * 10))
	var rim_dig_for_a = map[int]string{
		10:  "x",
		20:  "xx",
		30:  "xxx",
		40:  "xl",
		50:  "l",
		60:  "lx",
		70:  "lxx",
		80:  "lxxx",
		90:  "xc",
		100: "c",
	}
	var rim_dig_for_b = map[int]string{
		1: "i",
		2: "ii",
		3: "iii",
		4: "iv",
		5: "v",
		6: "vi",
		7: "vii",
		8: "viii",
		9: "ix",
	}
	//fmt.Print("\n")
	//fmt.Println("rim_digb")
	//fmt.Println(rim_digb)
	////fmt.Println(len(rim_dig))
	//fmt.Println("rim_digb")
	//fmt.Println("rim_dig_for_a[rim_diga]")
	//fmt.Println(rim_dig_for_a[int(rim_digaa)]) /////
	//fmt.Println((rim_digaa))                   /////
	//fmt.Println("rim_dig_for_a[rim_diga]")
	//fmt.Println("rim_dig_for_b[rim_digb]")
	//fmt.Println(rim_dig_for_b[rim_digb])
	//fmt.Println("rim_dig_for_b[rim_digb]")
	//fmt.Print("\n")
	//fmt.Println("rim_dig")
	//fmt.Println(rim_dig)
	//fmt.Println("rim_dig")
	a = rim_dig_for_a[rim_digaa] /////
	b = rim_dig_for_b[rim_digb]
	fmt.Println(strings.ToUpper(a + b))
	return a
	return b
}

var f_add = "+"
var f_sub = "-"
var f_mult = "*"
var f_dev = "/"
var f_add_flag = 0
var f_sub_flag = 0
var f_mult_flag = 0
var f_dev_flag = 0
var err_flag = 0
var err_flag_rim = 0
var flag_rim = 0
var flag_rim2 = 0
var flag_dig = 0
var flag_dig2 = 0

func main() {
	//fmt.Printf("%T", flag_rim)
	//fmt.Println("Enter function + - * / ")
	reader_f := bufio.NewReader(os.Stdin)
	for {
		input_f, _ := reader_f.ReadString('\n')
		ffff := strings.ToLower(input_f)

		allinput := strings.TrimSpace(ffff) // out
		add_f := (allinput + "              ")
		f := add_f
		//fmt.Printf("add_f is: %v", add_f)
		//fmt.Print("|" + allinput + "|")
		//
		//fmt.Print("\n")
		//
		//fmt.Println(allinput)
		//fmt.Printf("len f \t\t%d   ", len(f))
		//fmt.Print("\n")
		//
		//fmt.Printf("len allinput \t%d  ", len(allinput))
		//fmt.Print("\n")
		//
		//replacedText := strings.Replace(f, " ", "^", 20)
		//fmt.Println(replacedText)

		// first dig
		ff := f[0]
		//fmt.Println("0: " + string(f[0]))
		ff1 := f[1]
		//fmt.Println("1: " + string(f[1]))
		ff2 := f[2]
		ff3 := f[3]
		ff4 := f[4]
		var xx int
		{
			if ff == '0' {
				xx = 0
				flag_dig = 1
			} else if ff == '1' && ff1 == ' ' {
				xx = 1
				flag_dig = 1
			} else if ff == '2' && ff1 == ' ' {
				xx = 2
				flag_dig = 1
			} else if ff == '3' && ff1 == ' ' {
				xx = 3
				flag_dig = 1
			} else if ff == '4' && ff1 == ' ' {
				xx = 4
				flag_dig = 1
			} else if ff == '5' && ff1 == ' ' {
				xx = 5
				flag_dig = 1
			} else if ff == '6' && ff1 == ' ' {
				xx = 6
				flag_dig = 1
			} else if ff == '7' && ff1 == ' ' {
				xx = 7
				flag_dig = 1
			} else if ff == '8' && ff1 == ' ' {
				xx = 8
				flag_dig = 1
			} else if ff == '9' && ff1 == ' ' {
				xx = 9
				flag_dig = 1
			} else if ff == '1' && ff1 == '0' && ff2 == ' ' {
				xx = 10
				flag_dig = 1
			} else if (ff != '0' && ff1 != '0') || (ff != '1' && ff1 != '0') {
				//fmt.Println("Wrong number 10 or x max")
				err_flag = 1
			} else {
				//fmt.Println("Please input 1 + 10 or i + x")
				err_flag = 1
			}

		}
		{
			if ff == 'i' && ff1 == ' ' {
				xx = 1
				flag_rim = 1
			} else if ff == 'i' && ff1 == 'i' && ff2 == ' ' {
				xx = 2
				flag_rim = 1
			} else if ff == 'i' && ff1 == 'i' && ff2 == 'i' && ff3 == ' ' {
				xx = 3
				flag_rim = 1
			} else if ff == 'i' && ff1 == 'v' && ff2 == ' ' {
				xx = 4
				flag_rim = 1
			} else if ff == 'v' && ff1 == ' ' {
				xx = 5
				flag_rim = 1
			} else if ff == 'v' && ff1 == 'i' && ff2 == ' ' {
				xx = 6
				flag_rim = 1
			} else if ff == 'v' && ff1 == 'i' && ff2 == 'i' && ff3 == ' ' {
				xx = 7
				flag_rim = 1
			} else if ff == 'v' && ff1 == 'i' && ff2 == 'i' && ff3 == 'i' && ff4 == ' ' {
				xx = 8
				flag_rim = 1
			} else if ff == 'i' && ff1 == 'x' && ff2 == ' ' {
				xx = 9
				flag_rim = 1
			} else if ff == 'x' && ff1 == ' ' {
				xx = 10
				flag_rim = 1
			} else {
				err_flag_rim = 1
				//fmt.Println("er first rim")
			}
		}

		// func_math+err DONE
		mm := f[2]
		//fmt.Println("2: " + string(f[2]))
		mm1 := f[3]
		//fmt.Println("3: " + string(f[3]))
		mm2 := f[4]
		mm3 := f[5]
		mm4 := f[6]
		// math func
		{
			if (mm == '+') || (mm1 == '+') || (mm2 == '+') || (mm3 == '+') || (mm4 == '+') {
				f_add_flag = 1
				f_sub_flag = 0
				f_mult_flag = 0
				f_dev_flag = 0

			} else if (mm == '-') || (mm1 == '-') || (mm2 == '-') || (mm3 == '-') || (mm4 == '-') {
				f_add_flag = 0
				f_sub_flag = 1
				f_mult_flag = 0
				f_dev_flag = 0

			} else if (mm == '*') || (mm1 == '*') || (mm2 == '*') || (mm3 == '*') || (mm4 == '*') {
				f_add_flag = 0
				f_sub_flag = 0
				f_mult_flag = 1
				f_dev_flag = 0

			} else if (mm == '/') || (mm1 == '/') || (mm2 == '/') || (mm3 == '/') || (mm4 == '/') {

				f_add_flag = 0
				f_sub_flag = 0
				f_mult_flag = 0
				f_dev_flag = 1

			} else {
				fmt.Println("Error: You typed wrong mark, please use / * - + ")
			}
		}
		//fmt.Println("math func: ")

		ll_min_1 := f[3]
		ll := f[4]
		//fmt.Println("4: " + string(f[4]))
		ll1 := f[5]
		//fmt.Println("5: " + string(f[5]))
		ll2 := f[6]
		//fmt.Println("6: " + string(f[6]))
		ll3 := f[7]
		//fmt.Println("7: " + string(f[7]))
		ll4 := f[8]
		ll5 := f[9]
		ll6 := f[10]
		ll7 := f[11]
		ll8 := f[12]
		//ll3 := f[7]
		//fmt.Println("7: " + string(f[7]))
		//fmt.Println(string(f[8]))
		//fmt.Println(string(f[9]))
		//fmt.Println(string(f[10]))
		var yy int

		//&& ll1 == ' ' && ll2 == ' ' && ll3 == ' '
		//&& ll2 == ' ' && ll3 == ' ' && ll4 == ' '
		{
			if (ll == '1' && ll1 == ' ' && ll2 == ' ') || (ll1 == '1' && ll2 == ' ' && ll3 == ' ') {
				yy = 1
				flag_dig2 = 1
				//fmt.Println("1")
			} else if (ll == '2' && ll1 == ' ' && ll2 == ' ') || (ll1 == '2' && ll2 == ' ' && ll3 == ' ') {
				yy = 2
				flag_dig2 = 1
				//fmt.Println("2")
			} else if (ll == '3' && ll1 == ' ' && ll2 == ' ') || (ll1 == '3' && ll2 == ' ' && ll3 == ' ') {
				yy = 3
				flag_dig2 = 1
				//fmt.Println("3")
			} else if (ll == '4' && ll1 == ' ' && ll2 == ' ') || (ll1 == '4' && ll2 == ' ' && ll3 == ' ') {
				yy = 4
				flag_dig2 = 1
				//fmt.Println("4")
			} else if (ll == '5' && ll1 == ' ' && ll2 == ' ') || (ll1 == '5' && ll2 == ' ' && ll3 == ' ') {
				yy = 5
				flag_dig2 = 1
				//fmt.Println("5")
			} else if (ll == '6' && ll1 == ' ' && ll2 == ' ') || (ll1 == '6' && ll2 == ' ' && ll3 == ' ') {
				yy = 6
				flag_dig2 = 1
				//fmt.Println("6")
			} else if (ll == '7' && ll1 == ' ' && ll2 == ' ') || (ll1 == '7' && ll2 == ' ' && ll3 == ' ') {
				yy = 7
				flag_dig2 = 1
				//fmt.Println("7")
			} else if (ll == '8' && ll1 == ' ' && ll2 == ' ') || (ll1 == '8' && ll2 == ' ' && ll3 == ' ') {
				yy = 8
				flag_dig2 = 1
				//fmt.Println("8")

			} else if (ll == '9' && ll1 == ' ' && ll2 == ' ') || (ll1 == '9' && ll2 == ' ' && ll3 == ' ') {
				yy = 9
				flag_dig2 = 1
				//fmt.Println("9")
			} else if (ll == '1' && ll1 == '0' && ll2 == ' ' && ll3 == ' ') ||
				(ll1 == '1' && ll2 == '0' && ll3 == ' ' && ll4 == ' ') {
				yy = 10
				flag_dig2 = 1
				//fmt.Println("is second ten")
			} else if (ll == '0') || (ll1 == '0') {
				yy = 0
				flag_dig2 = 1
				//fmt.Println("0   ***  ")
			} else if (ff != '0' && ff1 != '0') || (ff != '1' && ff1 != '0') {
				//fmt.Println("Wro ax")
				err_flag = 1
			} else {
				err_flag = 1
				//fmt.Println("Wro ax2222222222222222222")
			}
		}
		{
			if (ll_min_1 == ' ' && ll == 'i' && ll1 == ' ' && ll2 == ' ') ||
				(ll == ' ' && ll1 == 'i' && ll2 == ' ' && ll3 == ' ') ||
				(ll1 == ' ' && ll2 == 'i' && ll3 == ' ' && ll4 == ' ') ||
				(ll2 == ' ' && ll3 == 'i' && ll4 == ' ' && ll5 == ' ') {
				yy = 1
				flag_rim2 = 1
			} else if (ll_min_1 == ' ' && ll == 'i' && ll1 == 'i' && ll2 == ' ' && ll3 == ' ') ||
				(ll == ' ' && ll1 == 'i' && ll2 == 'i' && ll3 == ' ' && ll4 == ' ') ||
				(ll1 == ' ' && ll2 == 'i' && ll3 == 'i' && ll4 == ' ' && ll5 == ' ') ||
				(ll2 == ' ' && ll3 == 'i' && ll4 == 'i' && ll5 == ' ' && ll6 == ' ') {
				yy = 2
				flag_rim2 = 1
			} else if (ll_min_1 == ' ' && ll == 'i' && ll1 == 'i' && ll2 == 'i' && ll3 == ' ' && ll4 == ' ') ||
				(ll == ' ' && ll1 == 'i' && ll2 == 'i' && ll3 == 'i' && ll4 == ' ' && ll5 == ' ') ||
				(ll1 == ' ' && ll2 == 'i' && ll3 == 'i' && ll4 == 'i' && ll5 == ' ' && ll6 == ' ') ||
				(ll2 == ' ' && ll3 == 'i' && ll4 == 'i' && ll5 == 'i' && ll6 == ' ' && ll7 == ' ') {
				yy = 3
				flag_rim2 = 1
			} else if (ll_min_1 == ' ' && ll == 'i' && ll1 == 'v' && ll2 == ' ' && ll3 == ' ') ||
				(ll == ' ' && ll1 == 'i' && ll2 == 'v' && ll3 == ' ' && ll4 == ' ') ||
				(ll1 == ' ' && ll2 == 'i' && ll3 == 'v' && ll4 == ' ' && ll5 == ' ') ||
				(ll2 == ' ' && ll3 == 'i' && ll4 == 'v' && ll5 == ' ' && ll6 == ' ') {
				yy = 4
				flag_rim2 = 1
			} else if (ll_min_1 == ' ' && ll == 'v' && ll1 == ' ' && ll2 == ' ') ||
				(ll == ' ' && ll1 == 'v' && ll2 == ' ' && ll3 == ' ') ||
				(ll1 == ' ' && ll2 == 'v' && ll3 == ' ' && ll4 == ' ') ||
				(ll2 == ' ' && ll3 == 'v' && ll4 == ' ' && ll5 == ' ') {
				yy = 5
				flag_rim2 = 1
			} else if (ll_min_1 == ' ' && ll == 'v' && ll1 == 'i' && ll2 == ' ' && ll3 == ' ') ||
				(ll == ' ' && ll1 == 'v' && ll2 == 'i' && ll3 == ' ' && ll4 == ' ') ||
				(ll1 == ' ' && ll2 == 'v' && ll3 == 'i' && ll4 == ' ' && ll5 == ' ') ||
				(ll2 == ' ' && ll3 == 'v' && ll4 == 'i' && ll5 == ' ' && ll6 == ' ') {
				yy = 6
				flag_rim2 = 1
			} else if (ll_min_1 == ' ' && ll == 'v' && ll1 == 'i' && ll2 == 'i' && ll3 == ' ' && ll4 == ' ') ||
				(ll == ' ' && ll1 == 'v' && ll2 == 'i' && ll3 == 'i' && ll4 == ' ' && ll5 == ' ') ||
				(ll1 == ' ' && ll2 == 'v' && ll3 == 'i' && ll4 == 'i' && ll5 == ' ' && ll6 == ' ') ||
				(ll2 == ' ' && ll3 == 'v' && ll4 == 'i' && ll5 == 'i' && ll6 == ' ' && ll7 == ' ') {
				yy = 7
				flag_rim2 = 1
			} else if (ll_min_1 == ' ' && ll == 'v' && ll1 == 'i' && ll2 == 'i' && ll3 == 'i' && ll4 == ' ' && ll5 == ' ') ||
				(ll == ' ' && ll1 == 'v' && ll2 == 'i' && ll3 == 'i' && ll4 == 'i' && ll5 == ' ' && ll6 == ' ') ||
				(ll1 == ' ' && ll2 == 'v' && ll3 == 'i' && ll4 == 'i' && ll5 == 'i' && ll6 == ' ' && ll7 == ' ') ||
				(ll2 == ' ' && ll3 == 'v' && ll4 == 'i' && ll5 == 'i' && ll6 == 'i' && ll7 == ' ' && ll8 == ' ') {
				yy = 8
				flag_rim2 = 1
			} else if (ll_min_1 == ' ' && ll == 'i' && ll1 == 'x' && ll2 == ' ' && ll3 == ' ') ||
				(ll == ' ' && ll1 == 'i' && ll2 == 'x' && ll3 == ' ' && ll4 == ' ') ||
				(ll1 == ' ' && ll2 == 'i' && ll3 == 'x' && ll4 == ' ' && ll5 == ' ') ||
				(ll2 == ' ' && ll3 == 'i' && ll4 == 'x' && ll5 == ' ' && ll6 == ' ') {
				yy = 9
				flag_rim2 = 1
			} else if (ll_min_1 == ' ' && ll == 'x' && ll1 == ' ' && ll2 == ' ') ||
				(ll == ' ' && ll1 == 'x' && ll2 == ' ' && ll3 == ' ') ||
				(ll1 == ' ' && ll2 == 'x' && ll3 == ' ' && ll4 == ' ') ||
				(ll2 == ' ' && ll3 == 'x' && ll4 == ' ' && ll5 == ' ') {
				yy = 10
				flag_rim2 = 1
			} else {
				err_flag_rim = 1
			}
		}
		// errors

		//fmt.Println(string(err_flag) + "\terr_flag")
		//fmt.Println(string(err_flag_rim) + "\terr_flag_rim")

		// rim to dig output
		f_cut_first := f[:2]
		first_dig, _ := strconv.Atoi(f_cut_first)
		//secound_dig := f[1]
		//first_numb := (first_dig + secound_dig)
		//fmt.Print("first_numb ")
		//fmt.Print(f[:2])
		//fmt.Println("\n")
		if (int(first_dig)) >= 11 {
			fmt.Println("Error: too big number use 10 or x max")
			break

		}
		if err_flag == 1 && err_flag_rim == 1 {
			fmt.Println("Error: wrong format use 1 + 4 or v + iii")
			break

		}

		//fmt.Println("$$$$$$$$$$$$$$$$")
		//fmt.Println(xx)
		//fmt.Println(yy)
		if f_add_flag == 1 && err_flag == 0 && flag_dig == 1 && flag_dig2 == 1 && err_flag_rim == 1 {
			//fmt.Print("add result ")
			fmt.Println(add(xx, yy))
		}
		if f_add_flag == 1 && err_flag == 1 && flag_rim == 1 && flag_rim2 == 1 && err_flag_rim == 0 {
			//fmt.Print("add result rim ")
			rim_dig_addstr := (xx + yy)
			//rim_dig_add, _ := strconv.Atoi(rim_dig_addstr)
			rim_number(rim_dig_addstr)
			//fmt.Println(rim_dig_addstr)
			//fmt.Println(add(xx, yy))
		}
		if f_sub_flag == 1 && err_flag == 0 && flag_dig == 1 && flag_dig2 == 1 && err_flag_rim == 1 {
			//fmt.Print("sub result ")
			fmt.Println(sub(xx, yy))
		}
		if f_sub_flag == 1 && err_flag == 1 && flag_rim == 1 && flag_rim2 == 1 && err_flag_rim == 0 {
			if yy < xx {
				//fmt.Print("sub result rim ")
				rim_dig_addstr := (xx - yy)
				rim_number(rim_dig_addstr)
			} else {
				/////////// erorachka
				fmt.Println("Error: The roman number don`t has negative numbers")
				break

			}
		}
		if f_mult_flag == 1 && err_flag == 0 && flag_dig == 1 && flag_dig2 == 1 && err_flag_rim == 1 {
			//fmt.Print("mult result ")
			fmt.Println(mult(xx, yy))
		}
		if f_mult_flag == 1 && err_flag == 1 && flag_rim == 1 && flag_rim2 == 1 && err_flag_rim == 0 {
			//fmt.Print("mult result rim ")
			rim_dig_addstr := (xx * yy)
			rim_number(rim_dig_addstr)
		}
		if f_dev_flag == 1 && err_flag == 0 && flag_dig == 1 && flag_dig2 == 1 && err_flag_rim == 1 {
			//fmt.Print("dev result ")
			fmt.Println(dev(xx, yy))
		}
		if f_dev_flag == 1 && err_flag == 1 && flag_rim == 1 && flag_rim2 == 1 && err_flag_rim == 0 {
			fmt.Print("dev result rim ")
			rim_dig_addstr := (xx / yy)
			rim_number(rim_dig_addstr)
		}

		//fmt.Println("$$$$$$$$$$$$$$$$")
		f_add_flag = 0
		f_sub_flag = 0
		f_mult_flag = 0
		f_dev_flag = 0
		err_flag = 0
		err_flag_rim = 0
		flag_rim = 0
	}
}
