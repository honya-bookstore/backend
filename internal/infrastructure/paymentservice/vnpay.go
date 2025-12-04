package paymentservice

import (
	"context"

	"backend/config"
	"backend/internal/application"
	"backend/internal/domain"

	"github.com/hashicorp/go-multierror"
	"github.com/lamphusy/go-vnpay/govnpay"
	govnpaymodels "github.com/lamphusy/go-vnpay/model"
)

type VNPay struct {
	srvCfg config.Server
}

var _ application.OrderPaymentService = (*VNPay)(nil)

func (o *VNPay) GetPaymentURL(ctx context.Context, param application.GetPaymentURLParam) (string, error) {
	url, err := govnpay.GetPaymentURL(&govnpaymodels.GetPaymentURLRequest{
		Version:        govnpay.Version210,
		TmnCode:        o.srvCfg.VNPTMNCode,
		ReturnURL:      param.ReturnURL,
		Amount:         param.Order.TotalAmount,
		OrderInfo:      "",
		TxnRef:         param.Order.ID.String(),
		CreateDate:     param.Order.CreatedAt,
		IpAddr:         "0.0.0.0",
		HashSecret:     o.srvCfg.VNPSecureSecret,
		HashAlgo:       o.srvCfg.VNPHashAlgo,
		InitPaymentURL: o.srvCfg.VNPURL,
	})
	if err != nil {
		return "", multierror.Append(nil, domain.ErrInternal, err)
	}
	return url, nil
}

func (a *VNPay) VerifyIPN() error
