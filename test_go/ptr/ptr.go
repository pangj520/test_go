package ptr

import (
	"fmt"
	"unsafe"
)

type Admin struct {
	Name string
	Age  int
}

func PtrConvertTo() {
	admin := Admin{
		Age:  18,
		Name: "seekload",
	}
	ptr := &admin
	age := (*string)(unsafe.Pointer(ptr)) // 1
	fmt.Printf("name%p\n, ptr:%p", age, ptr)
	*age = "60"

	fmt.Println(*ptr)

	age2 := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + unsafe.Offsetof(ptr.Name))) // 2
	*age2 = "name"
	fmt.Printf("age%p", age2)
	fmt.Println(*ptr)
}
