package paymentservice

import (
	"context"
	"strconv"

	"backend/config"
	"backend/internal/application"
	"backend/internal/domain"

	"github.com/hashicorp/go-multierror"
	govnpayerrors "github.com/lamphusy/go-vnpay/error"
	"github.com/lamphusy/go-vnpay/govnpay"
	govnpaymodels "github.com/lamphusy/go-vnpay/model"
)

type VNPay struct {
	srvCfg *config.Server
}

var _ application.VNPayPaymentService = (*VNPay)(nil)

func ProvideVNPay(srvCfg *config.Server) *VNPay {
	return &VNPay{srvCfg: srvCfg}
}

func (o *VNPay) GetPaymentURL(ctx context.Context, param application.GetPaymentURLVNPayParam) (string, error) {
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

func (a *VNPay) VerifyIPN(ctx context.Context, param application.VerifyVNPayIPNParam) (code string, message string) {
	if param.Data.SecureHash != a.srvCfg.VNPSecureSecret || param.Data.TmnCode != a.srvCfg.VNPTMNCode {
		code = govnpayerrors.MerchantRespInvalidSignature.ToString()
		message = govnpayerrors.MerchantRespInvalidSignature.Message()
		return
	}
	amount, err := strconv.ParseInt(param.Data.Amount, 10, 64)
	if err != nil {
		code = govnpayerrors.MerchantRespInvalidAmount.ToString()
		message = govnpayerrors.MerchantRespInvalidAmount.Message()
		return
	}
	if amount != param.Order.TotalAmount {
		code = govnpayerrors.MerchantRespInvalidAmount.ToString()
		message = govnpayerrors.MerchantRespInvalidAmount.Message()
		return
	}
	code = govnpayerrors.IPNCodeTransactionSuccess.ToString()
	message = govnpayerrors.IPNCodeTransactionSuccess.Message()
	return
}
