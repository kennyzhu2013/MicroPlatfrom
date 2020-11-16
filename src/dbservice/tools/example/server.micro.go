package example

import (
	"context"
	categories2 "github.com/kennyzhu/go-os/src/dbservice/tools/example/categories"
)

type exampleServiceServer struct {
}

func NewExampleServiceServer() *exampleServiceServer {
	return &exampleServiceServer{}
}
func (s *exampleServiceServer) GetAlphaTime(ctx context.Context, in *GetAlphaTimeRequest, result *MyTime) (err error) {
	tempResult := new(MyTime)
	aux := GetAlphaTime()
	tempResult = &aux
	*result = *tempResult
	return
}
func (s *exampleServiceServer) GetDurationForLength(ctx context.Context, in *GetDurationForLengthRequest, result *MyDuration) (err error) {
	tempResult := new(MyDuration)
	tempResult = GetDurationForLength(in.Arg1)
	*result = *tempResult
	return
}
func (s *exampleServiceServer) GetDurationForLengthCtx(ctx context.Context, in *GetDurationForLengthCtxRequest, result *MyDuration) (err error) {
	tempResult := new(MyDuration)
	tempResult, err = GetDurationForLengthCtx(ctx, in.Arg1)
	*result = *tempResult
	return
}
func (s *exampleServiceServer) GetOmegaTime(ctx context.Context, in *GetOmegaTimeRequest, result *MyTime) (err error) {
	tempResult := new(MyTime)
	tempResult, err = GetOmegaTime()
	*result = *tempResult
	return
}
func (s *exampleServiceServer) GetPhone(ctx context.Context, in *GetPhoneRequest, result *Product) (err error) {
	tempResult := new(Product)
	tempResult = GetPhone()
	*result = *tempResult
	return
}
func (s *exampleServiceServer) RandomCategory(ctx context.Context, in *RandomCategoryRequest, result *categories2.CategoryOptions) (err error) {
	tempResult := new(categories2.CategoryOptions)
	aux := RandomCategory()
	tempResult = &aux
	*result = *tempResult
	return
}
func (s *exampleServiceServer) RandomNumber(ctx context.Context, in *RandomNumberRequest, result *RandomNumberResponse) (err error) {
	tempResult := new(RandomNumberResponse)
	tempResult.Result1 = RandomNumber(in.Arg1, in.Arg2)
	*result = *tempResult
	return
}
