package apiresp

import (
	"testing"

	"github.com/ammmnia/tools/utils/jsonutil"
	"github.com/openimsdk/protocol/relation"
	"github.com/openimsdk/protocol/wrapperspb"
)

func TestName(t *testing.T) {
	resp := &ApiResponse{
		ErrCode: 1234,
		ErrMsg:  "test",
		ErrDlt:  "4567",
		Data: &relation.UpdateFriendsReq{
			OwnerUserID:   "123456",
			FriendUserIDs: []string{"1", "2", "3"},
			Remark:        wrapperspb.String("1234567"),
		},
	}
	data, err := resp.MarshalJSON()
	if err != nil {
		panic(err)
	}
	t.Log(string(data))

	var rReso ApiResponse
	rReso.Data = &relation.UpdateFriendsReq{}

	if err := jsonutil.JsonUnmarshal(data, &rReso); err != nil {
		panic(err)
	}

	t.Logf("%+v\n", rReso)
}
