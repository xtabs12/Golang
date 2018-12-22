Package RcMath

import (
	"fmt"
)

//Factorial Function
func FactoriInt32(value int32) int32 {
	var result int32
	result = 1
	if value > 15 {
		fmt.Printf("%i\n","error!!! -> Overload")
		return 1
	} else {
		for i:=int32(0);i<=value-1;i++ {
			result = result * (i+1)
		}
		return result
	}
}

func FactoriInt64(value int64) int64 {
	var result int64
	result = 1
	if value > 20 {
		fmt.Printf("%i\n","error!!! -> Overload")
		return 1
	} else if value > 0 {
		for i:=int64(1);i<=value;i++ {
			result = result * i
			fmt.Printf("%s %s %s\n",i ,"-,-",result)
		}
		return result
	}
	return 1	
}
//end Factorial Function

