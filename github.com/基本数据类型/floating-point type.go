package main;

import "fmt";
import "math";

func main(){
	fmt.Printf("%f\n", math.Pi);
	fmt.Printf("%.2f\n", math.Pi);
	fmt.Printf("%.4f\n", math.Pi);

	fmt.Println(math.MaxFloat32);
	fmt.Println(math.MaxFloat64);
}
