package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type TypeProduct int

const (
	quanAo TypeProduct = iota
	dienTu
	giayDep
)

type Product struct {
	name        string
	color       Color
	size        Size
	typeProduct TypeProduct
}

type Filter struct{}

func (f *Filter) filterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func (f *Filter) filterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

/*
- Bạn có thể thấy hàm filterByColor và filterBySize đang vi phạm nguyên tắc mở đóng tại vì
Nguyên tắc OCP quy định rằng các thực thể phần mềm (lớp, module, hàm, v.v.) nên "mở để mở rộng, đóng để sửa đổi".
Điều này có nghĩa là bạn nên có thể thêm chức năng mới mà không cần thay đổi mã hiện có.
Trong đoạn trên, nếu bạn muốn thêm các tiêu chí lọc mới (ví dụ: lọc theo TypeProduct),
bạn sẽ phải sửa đổi lớp Filter và thêm các phương thức mới. Điều này vi phạm nguyên tắc OCP.
*/

//Thực hiện nguyên tắc OCP

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

type TypeProductpecification struct {
	typeProduct TypeProduct
}

func (spec TypeProductpecification) IsSatisfied(p *Product) bool {
	return p.typeProduct == spec.typeProduct
}

type AndSpecification struct {
	first, second, third Specification
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) && spec.second.IsSatisfied(p) && spec.third.IsSatisfied(p)
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	aothun := Product{"ao-thun-cut-tay", green, small, quanAo}
	maylanh := Product{"may-lanh-to-si-ba", green, large, dienTu}
	depbitis := Product{"dep-deo-ngang-bitis", red, large, giayDep}

	products := []Product{aothun, maylanh, depbitis}
	// fmt.Print("Green products (old):\n")
	// f := Filter{}
	// for _, v := range f.filterByColor(products, green) {
	// 	fmt.Printf(" - %s is green\n", v.name)
	// }

	fmt.Print("Green products (new):\n")
	greenSpec := ColorSpecification{green}

	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}

	typeSpec := TypeProductpecification{dienTu}

	largeGreenSpec := AndSpecification{largeSpec, greenSpec, typeSpec}
	
	fmt.Print("Large blue items:\n")
	for _, v := range bf.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green and dientu\n", v.name)
	}
}
