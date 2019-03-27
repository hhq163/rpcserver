package main

import "context"

type User struct{}

//扣钱接口
func (u *User) DeductMoney(ctx context.Context, req *pb.DeductMoneyRequest) (*pb.DeductMoneyResponse, error) {
	if(req.hall_id <= 0 || req.user_id <= 0){
		return &pb.DeductMoneyResponse{
			result_code : -1,
			desc : "hall_id or user_id <=0",
			restult : {
				amount: 0.0,
				order_sn: ""
			}
		}, fmt.Errorf("param is wrong!")
	}

	pushmongo(func(){
		mgo := mgoSession.Copy()
		defer mgo.Close()
		doc := bson.M{
			"order_sn":      grid,
			"uid":           playdata.userID,
			"agent_id":      playdata.agentID,
			"hall_id":       playdata.hallID,
			"user_name":     playdata.account,
			"type":          itype,
			"amount":        totalBetValue + withholdingMoneys,
			"status":        4,
			"user_money":    nowmoney,
			"desc":          roundno,
			"admin_user":    playdata.account,
			"admin_user_id": playdata.userID,
			"cash_no":       roundno,
			"add_time":      time.Now(),
			"pKey":          pKey,
		}

		if err := session.DB("").C("cash_record").Insert(doc); err != nil {
			docjosn, _ := json.Marshal(doc)
			slog.ErrorDB("cash_record failed err:", err, string(docjosn))
		}
		pushmysql(func(){
			tablename := GetTableName(hall_id)
			result, err := mysqlDB.Exec(fmt.Sprintf("UPDATE %s SET password=?, password_md=? WHERE uid=? AND password=? LIMIT 1", tablename), strNewPwd, strNewPwd, s.userid, strOldPwd)
			if err != nil {
				slog.ErrorDB(err)
				return
			}
			rows, err := result.RowsAffected()
			if err != nil {
				slog.ErrorDB(err)
			}
			if rows < 0 {
				
			}

		})

	})
}

//加钱接口
func (u *User) Recharge(ctx context.Context, req *pb.RechargeRequest) (*pb.RechargeResponse, error) {

}


