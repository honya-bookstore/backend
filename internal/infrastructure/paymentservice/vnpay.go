package paymentservice

import (
	"context"
	"errors"
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
	param application.VerifyIPNVNPayParam,
	getOrder func(ctx context.Context, orderID uuid.UUID) (*domain.Order, error),
	onSuccess func(ctx context.Context, order *domain.Order) error,
	onFailure func(ctx context.Context, order *domain.Order) error,
) (code, message string, err error) {
	codeEnum := govnpayerrors.IPNCodeTransactionSuccess
	defer func() {
		code = codeEnum.ToString()
		message = codeEnum.Message()
	}()
	tnxRef, err := uuid.Parse(param.TxnRef)
	if err != nil {
		codeEnum = govnpayerrors.IPNCodeOtherErrors
		err = multierror.Append(domain.ErrInvalid, err)
		return code, message, err
	}
	verifyIPNRequest := &govnpaymodels.VerifyIPNRequest{
		TxnRef:            param.TxnRef,
		Amount:            param.Amount,
		ResponseCode:      param.ResponseCode,
		SecureHash:        param.SecureHash,
		TmnCode:           param.TmnCode,
		TransactionStatus: param.TransactionStatus,
		BankTranNo:        param.BankTranNo,
		CardType:          param.CardType,
		BankCode:          param.BankCode,
		OrderInfo:         param.OrderInfo,
		PayDate:           param.PayDate,
		TransactionNo:     param.TransactionNo,
		HashSecret:        v.srvCfg.VNPSecureSecret,
		HashAlgo:          govnpayhelper.HashAlgo(v.srvCfg.VNPHashAlgo),
	}
	if ok, err := govnpay.VerifyIPN(verifyIPNRequest); !ok || err != nil {
		codeEnum = govnpayerrors.IPNCodeInvalidSignature
		err = multierror.Append(domain.ErrInvalid, errors.New("invalid signature or tmn code"), err)
		return code, message, err
	}
	order, err := getOrder(ctx, tnxRef)
	if err != nil {
		codeEnum = govnpayerrors.IPNCodeOrderNotFound
		return code, message, err
	}
	amount, err := strconv.ParseInt(param.Amount, 10, 64)
	if err != nil {
		codeEnum = govnpayerrors.IPNCodeInvalidAmount
		err = multierror.Append(domain.ErrInvalid, err)
		return code, message, err
	}
	if amount/100 != order.TotalAmount {
		codeEnum = govnpayerrors.IPNCodeInvalidAmount
		err = multierror.Append(domain.ErrInvalid, onFailure(ctx, order))
		return code, message, err
	}
	if err := onSuccess(ctx, order); err != nil {
		codeEnum = govnpayerrors.IPNCodeOtherErrors
		err = multierror.Append(domain.ErrInternal, err)
		return code, message, err
	}
	return code, message, err
}
