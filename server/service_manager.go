package server

import (
	"context"
	"fmt"
	"time"

	"github.com/ProtonFundationGroup/seventies/v2/module/chat"
	"github.com/ProtonFundationGroup/seventies/v2/module/wallet"
	"github.com/ProtonFundationGroup/seventies/v2/server/pb-go/api"
)

var _ api.APIServiceServer = (*ServiceManager)(nil)

type ServiceManager struct {
	walletSvc *wallet.Service
	chatSvc   *chat.Service
}

func newServiceManager() *ServiceManager {
	return &ServiceManager{
		walletSvc: &wallet.Service{},
		chatSvc:   &chat.Service{},
	}
}

func (mgr *ServiceManager) TestMethod(ctx context.Context, req *api.MyRequest) (*api.MyResponse, error) {
	timestamp := time.Unix(req.Timestamp, 0).Format("2006-01-02 15:04:05")
	rspMsg := fmt.Sprintf("【%v】the method `%v` of service `%v` was invoked, with parameters = %v",
		timestamp, req.Method, req.ServiceName, req.Parameters)

	return &api.MyResponse{
		RspMsg: rspMsg,
	}, nil
}
