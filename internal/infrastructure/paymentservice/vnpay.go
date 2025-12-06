package paymentservice

import (
	"context"
	"fmt"
	"strconv"

	"backend/config"
	"backend/internal/application"
	"backend/internal/domain"

	govnpayerrors "github.com/electricilies/govnpay/error"
	"github.com/electricilies/govnpay/govnpay"
	govnpayhelper "github.com/electricilies/govnpay/helper"
	govnpaymodels "github.com/electricilies/govnpay/model"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

type VNPay struct {
	srvCfg *config.Server
}

var _ application.VNPayPaymentService = (*VNPay)(nil)

func ProvideVNPay(srvCfg *config.Server) *VNPay {
	return &VNPay{srvCfg: srvCfg}
}

func (v *VNPay) GetPaymentURL(
	ctx context.Context,
	param application.GetPaymentURLVNPayParam,
) (string, error) {
	url, err := govnpay.GetPaymentURL(&govnpaymodels.GetPaymentURLRequest{
		Version:        govnpay.Version210,
		TmnCode:        v.srvCfg.VNPTMNCode,
		ReturnURL:      param.ReturnURL,
		Amount:         param.Order.TotalAmount,
		OrderInfo:      "",
		TxnRef:         param.Order.ID.String(),
		CreateDate:     param.Order.CreatedAt,
		IpAddr:         "0.0.0.0",
		HashSecret:     v.srvCfg.VNPSecureSecret,
		HashAlgo:       govnpayhelper.HashAlgo(v.srvCfg.VNPHashAlgo),
		InitPaymentURL: v.srvCfg.VNPURL,
	})
	if err != nil {
		return "", multierror.Append(nil, domain.ErrInternal, err)
	}
	return url, nil
}

func (v *VNPay) VerifyIPN(
	ctx context.Context,
	param application.VerifyVNPayIPNParam,
	getOrder func(ctx context.Context, orderID uuid.UUID) (*domain.Order, error),
	onSuccess func(ctx context.Context, order *domain.Order) error,
	onFailure func(ctx context.Context, order *domain.Order) error,
) (code string, message string) {
	codeEnum := govnpayerrors.IPNCodeTransactionSuccess
	var err error
	defer func() {
		code = codeEnum.ToString()
		message = codeEnum.Message()
		if err != nil {
			message = fmt.Sprintf("%s: %s", message, err.Error())
		}
	}()
	if param.Data.SecureHash != v.srvCfg.VNPSecureSecret || param.Data.TmnCode != v.srvCfg.VNPTMNCode {
		codeEnum = govnpayerrors.IPNCodeInvalidSignature
		return
	}
	order, err := getOrder(ctx, param.OrderID)
	if err != nil {
		codeEnum = govnpayerrors.IPNCodeOrderNotFound
		return
	}
	amount, err := strconv.ParseInt(param.Data.Amount, 10, 64)
	if err != nil || amount != order.TotalAmount {
		codeEnum = govnpayerrors.IPNCodeInvalidAmount
		onFailure(ctx, order)
		return
	}
	if err := onSuccess(ctx, order); err != nil {
		codeEnum = govnpayerrors.IPNCodeOtherErrors
	}
	return
}
