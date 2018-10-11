/*
@Time : 2018/10/10 16:49 
@Author : kenny zhu
@File : server.micro.go
@Software: GoLand
@Others:
*/
package example

import (
	"github.com/kennyzhu/go-os/dbservice/tools/example/categories"

	"context"
	"fmt"
)

type Say struct {
}

/**
/ 区别：1：没有exampleServiceServer示例结构（可以有）
2：函数GetAlphaTime(ctx xcontext.Context, in *GetAlphaTimeRequest) (result *MyTime, err error)
	类比函数GetAlphaTime(context.Context, *GetAlphaTimeRequest, *MyTime) error,
	函数代码内容一样的
3: 注意：函数传输参数指针变量值不能修改，内存是外部分配好的.
 */
/*
func NewExampleServiceServer() *exampleServiceServer {
	return &exampleServiceServer{}
}
func (s *exampleServiceServer) GetAlphaTime(ctx xcontext.Context, in *GetAlphaTimeRequest) (result *MyTime, err error) {
	result = new(MyTime)
	aux := GetAlphaTime()
	result = &aux
	return
}*/
func NewSay() *Say {
	return &Say{}
}

func (s *Say) GetAlphaTime(ctx context.Context, in *GetAlphaTimeRequest, result *MyTime) error {
	// result = new(MyTime)
	aux := GetAlphaTime()
	*result = aux
	fmt.Println("GetAlphaTime Enter:")
	return nil
}
func (s *Say) GetDurationForLength(ctx context.Context, in *GetDurationForLengthRequest, result *MyDuration) error {
	// result = new(MyDuration)
	aux := GetDurationForLength(in.Arg1)
	*result = *aux
	return nil
}
func (s *Say) GetDurationForLengthCtx(ctx context.Context, in *GetDurationForLengthCtxRequest, result *MyDuration) (err error) {
	// result = new(MyDuration)
	var aux*MyDuration
	aux, err = GetDurationForLengthCtx(ctx, in.Arg1)
	*result = *aux
	return
}
func (s *Say) GetOmegaTime(ctx context.Context, in *GetOmegaTimeRequest, result *MyTime) (err error) {
	// result = new(MyTime)
	var aux*MyTime
	aux, err = GetOmegaTime()
	*result = *aux
	return
}

//
func (s *Say) GetPhone(ctx context.Context, in *GetPhoneRequest, result *Product) (err error) {
	// result = new(Product)
	fmt.Printf("%p\n", result)
	// result = new(Product)
	aux := GetPhone()
	*result = *aux
	result.CategoryID = 100
	fmt.Printf("%p\n", result)
	fmt.Println(*result)
	return
}
func (s *Say) RandomCategory(ctx context.Context, in *RandomCategoryRequest, result *categories.CategoryOptions) (err error) {
	// result = new(categories.CategoryOptions)
	aux := RandomCategory()
	*result = aux
	return
}
func (s *Say) RandomNumber(ctx context.Context, in *RandomNumberRequest, result *RandomNumberResponse) (err error) {
	// result = new(RandomNumberResponse)
	result.Result1 = RandomNumber(in.Arg1, in.Arg2)
	return
}

