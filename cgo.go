package main


//int Add(int x, int y){
//	return x+y;
//}
import "C"

func CallC(x C.int,y C.int)  C.int{
	return  C.Add(x,y)
}