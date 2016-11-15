package main

import (
"fmt"
"math"
)


var memory [8388608]byte
var all_inf_about_pages []Page
var adress *byte
var how_much_memory_needed int
var one_MB int = 1048576
var new_size int


type Page struct {
	flag_of_joint bool
	size_of_blocks int
	point_of_start *byte
	full_blocks int
	array_of_free_blocks []*byte

}

func main() {
	fmt.Println("All inf about pages")
	init_pages()
	fmt.Println(all_inf_about_pages)

	fmt.Println("ALLOC")
	how_much_memory_needed = 15
	addr, quantity, numb_of_page := mem_alloc(how_much_memory_needed)
	print_inf()


	for i:=0;i<70000;i++{
		how_much_memory_needed = 15
		addr, quantity, numb_of_page = mem_alloc(how_much_memory_needed)

	}


	fmt.Println("ALLOC")
	how_much_memory_needed = 15
	addr, quantity, numb_of_page = mem_alloc(how_much_memory_needed)
	print_inf()

	fmt.Println("ALLOC")
	how_much_memory_needed = 38
	addr, quantity, numb_of_page = mem_alloc(how_much_memory_needed)
	print_inf()

	fmt.Println("ALLOC")
	how_much_memory_needed = 1048577
	addr, quantity, numb_of_page = mem_alloc(how_much_memory_needed)
	print_inf()

	fmt.Println("FREE")
	mem_free(addr,quantity,numb_of_page)
	print_inf()

	fmt.Println("ALLOC")
	how_much_memory_needed = 32
	addr, quantity,numb_of_page = mem_alloc(how_much_memory_needed)
	print_inf()

	fmt.Println("ALLOC")
	how_much_memory_needed = 65
	addr, quantity,numb_of_page = mem_alloc(how_much_memory_needed)
	print_inf()

	fmt.Println("ALLOC")
	how_much_memory_needed =129
	addr, quantity, numb_of_page = mem_alloc(how_much_memory_needed)
	print_inf()

	fmt.Println("REALLOC")
	new_size = 1025
	mem_realloc(new_size, addr,quantity,numb_of_page)
	print_inf()

	fmt.Println("ALLOC")
	how_much_memory_needed =257
	addr, quantity ,numb_of_page = mem_alloc(how_much_memory_needed)
	print_inf()

	fmt.Println("ALLOC")
	how_much_memory_needed =513
	addr, quantity ,numb_of_page = mem_alloc(how_much_memory_needed)
	print_inf()


	fmt.Println("REALLOC")
	new_size = 1025
	mem_realloc(new_size, addr,quantity,numb_of_page)
	print_inf()

	fmt.Println("ALLOC")
	how_much_memory_needed =32
	addr, quantity ,numb_of_page = mem_alloc(how_much_memory_needed)
	print_inf()


	fmt.Println("FREE")
	mem_free(addr,quantity,numb_of_page)
	print_inf()

	fmt.Println("ALLOC")
	how_much_memory_needed =512
	addr, quantity ,numb_of_page = mem_alloc(how_much_memory_needed)
	print_inf()

	fmt.Println("REALLOC")
	new_size = 1025
	addr, quantity ,numb_of_page = mem_realloc(new_size, addr,quantity,numb_of_page)
	print_inf()

	fmt.Println("FREE")
	mem_free(addr,quantity,numb_of_page)
	print_inf()

}


func init_pages(){

	pages := len(memory) / one_MB   // devide for 1 MByte

	for i:=0; i<pages ;i++{
		adress = &memory[i*one_MB] //адрес начала нового блока
		all_inf_about_pages = append(all_inf_about_pages, Page{flag_of_joint: false, size_of_blocks: 0 , point_of_start: adress, full_blocks: 0})
	}

}


func mem_alloc(how_much_memory_needed int) (adress *byte, quantity_of_blocks int, number_of_page int) {


	flag_for_changing_new_page := false

	how_much_memory_needed = to_2(how_much_memory_needed)
	pages := len(memory) / one_MB   // devide for 1 megaByte



	if how_much_memory_needed> one_MB { //значит, что не помещается на 1 страницу
		count := 0
		index_of_last_el := -1
		new := how_much_memory_needed/ one_MB //сколько страниц нужно
		for i:=0;i<pages;i++{

			if all_inf_about_pages[i].size_of_blocks == 0{
				count += 1
			}
			if count >= new{
				index_of_last_el = i
				break
			}
		}

		if index_of_last_el != -1{

			for i:=index_of_last_el-new+1;i<index_of_last_el+1;i++{

				all_inf_about_pages[i].array_of_free_blocks = append(all_inf_about_pages[i].array_of_free_blocks,all_inf_about_pages[i].point_of_start)

				all_inf_about_pages[i].size_of_blocks = how_much_memory_needed
				all_inf_about_pages[i].full_blocks = 1
				all_inf_about_pages[i].flag_of_joint = true
			}

			adress = all_inf_about_pages[index_of_last_el-new+1].array_of_free_blocks[0]

			for i:=index_of_last_el-new+1;i<index_of_last_el+1;i++{ //фор для того, чтоб убрать выдееленную память из информации о страницах

				all_inf_about_pages[i].array_of_free_blocks = all_inf_about_pages[i].array_of_free_blocks[1:]
			}

		}
		return adress, new, index_of_last_el-new+1
	}



	for i:=0;i<pages;i++{
		if all_inf_about_pages[i].size_of_blocks == how_much_memory_needed{
			flag_for_changing_new_page = true

			if one_MB /int(how_much_memory_needed) > all_inf_about_pages[i].full_blocks+1 { //если в страницу помещается еще одна задача
				all_inf_about_pages[i].full_blocks = all_inf_about_pages[i].full_blocks + 1

				adress = all_inf_about_pages[i].array_of_free_blocks[1]
				all_inf_about_pages[i].array_of_free_blocks = all_inf_about_pages[i].array_of_free_blocks[1:]

				return adress, 1, i
			}else{
				flag_for_changing_new_page = false //если не помещается
			}

		}
	}
	if flag_for_changing_new_page == false{ //если страницы с нужным размером не было еще или она заполнилась

		flag_for_err:= false

		for i:=0;i<pages;i++{
			if all_inf_about_pages[i].size_of_blocks == 0{


				all_inf_about_pages[i].array_of_free_blocks = append(all_inf_about_pages[i].array_of_free_blocks,all_inf_about_pages[i].point_of_start)


				for free_block:=1;free_block<(one_MB/how_much_memory_needed);free_block++{
					all_inf_about_pages[i].array_of_free_blocks = append(all_inf_about_pages[i].array_of_free_blocks,&memory[i+how_much_memory_needed*free_block])
				}

				all_inf_about_pages[i].size_of_blocks = how_much_memory_needed
				all_inf_about_pages[i].full_blocks = 1
				flag_for_err = true


				adress = all_inf_about_pages[i].array_of_free_blocks[0]
				all_inf_about_pages[i].array_of_free_blocks = all_inf_about_pages[i].array_of_free_blocks[1:]
				return adress, 1, i
			}
		}
		if flag_for_err == false{
			fmt.Println("Error. We can't give you memory, because memory is full")

		}
	}

	return adress, 0, -1

}


func to_2(value int) int { 											//по 4 байта

	if value<=16 {
		return 16
	}

	two_degrees:= math.Log2(float64(value))
	new_value:=1

	if int(two_degrees*1000000) == int(two_degrees)*1000000{
		for i:=0;i<int(two_degrees);i++{
			new_value = new_value*2
		}
		return new_value
	}else {
		for i:=0;i<int(two_degrees)+1;i++{
			new_value = new_value*2
		}
		return new_value
	}

	return 16
}


func mem_free(adress *byte, quantity_of_blocks int, number_of_page int)  {
	if quantity_of_blocks == 1 && number_of_page != -1{

		all_inf_about_pages[number_of_page].array_of_free_blocks = append(all_inf_about_pages[number_of_page].array_of_free_blocks, adress)
		all_inf_about_pages[number_of_page].full_blocks = all_inf_about_pages[number_of_page].full_blocks-1
		if all_inf_about_pages[number_of_page].full_blocks == 0{
			all_inf_about_pages[number_of_page].size_of_blocks = 0
		}
	}
	if quantity_of_blocks > 1 && number_of_page != -1{

		for i:=number_of_page;i<number_of_page+quantity_of_blocks;i++{
			all_inf_about_pages[i].array_of_free_blocks = append(all_inf_about_pages[i].array_of_free_blocks, adress)
			all_inf_about_pages[i].full_blocks = all_inf_about_pages[i].full_blocks-1
			all_inf_about_pages[i].size_of_blocks = 0
			all_inf_about_pages[i].flag_of_joint = false

		}
	}
}


func mem_realloc(new_size int, adress *byte, quantity_of_blocks int, number_of_page int)(adress_new *byte, quantity_of_blocks_new int, nubm int){

	var buf []byte
	var new1 int
	var new1_for_buf_new int

	for i:=0; i<=number_of_page ;i++{
		new1 = new1 + all_inf_about_pages[i].size_of_blocks*(quantity_of_blocks+2)
	}

	for i:= new1;i < new1+all_inf_about_pages[number_of_page].size_of_blocks;i++{
		buf = append(buf,memory[i])
	}

	if quantity_of_blocks == 1 && number_of_page != -1{
		mem_free(adress, quantity_of_blocks, number_of_page)
		old_size_of_block := all_inf_about_pages[number_of_page].size_of_blocks
		adress_new,quantity_of_blocks_new,nubm = mem_alloc(new_size)

		if nubm < 0{
			adress_new,quantity_of_blocks_new,nubm = mem_alloc(old_size_of_block)
			for i:= new1;i < new1+all_inf_about_pages[number_of_page].size_of_blocks;i++{
				memory[i] = buf[i-new1]
			}
		}else {
			for i:=0; i<=nubm ;i++{
				new1_for_buf_new = new1_for_buf_new + all_inf_about_pages[i].size_of_blocks*(quantity_of_blocks+2)
			}
			for i:= new1_for_buf_new;i < new1_for_buf_new+old_size_of_block;i++{
				memory[i] = buf[i-new1_for_buf_new]
			}
		}

	}
	return adress_new, quantity_of_blocks_new, nubm
}

func print_inf() {

	for i:=0;i<len(all_inf_about_pages);i++{
		fmt.Println(all_inf_about_pages[i].flag_of_joint, " ; ",all_inf_about_pages[i].full_blocks, " ; ",all_inf_about_pages[i].point_of_start, " ; ",all_inf_about_pages[i].size_of_blocks, " ; ")
	}
	fmt.Println("---------------------------------------")
}
