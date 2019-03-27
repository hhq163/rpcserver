package main

import "context"

type User struct{}

//扣款接口
func (u *User) DeductMoney(ctx context.Context, req *pb.DeductMoneyRequest) (*pb.DeductMoneyResponse, error) {

}

//充值接口
func (u *User) Recharge(ctx context.Context, req *pb.RechargeRequest) (*pb.RechargeResponse, error) {

}
