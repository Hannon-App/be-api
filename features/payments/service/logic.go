package service

import (
	"Hannon-app/app/config"
	"Hannon-app/features/payments"
	"fmt"
	"strings"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/virtualaccount"
)

type PaymentService struct {
	PaymentData payments.PaymentDataInterface
	config      *config.AppConfig
}

// AddVA implements payments.PaymentServiceInterface.
func (service *PaymentService) AddVA(input payments.VirtualAccountObjectCore) error {
	xendit.Opt.SecretKey = service.config.SecretKeyXendit

	year, month, day := time.Now().Date()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	nsecond := time.Now().Nanosecond()

	externalID := fmt.Sprint("va-", year, month, day, hour, minute, second, nsecond)
	externalID = strings.ReplaceAll(externalID, " ", "")
	data := virtualaccount.CreateFixedVAParams{
		ExternalID: externalID,
		BankCode:   "BCA",
		Name:       "Hannon",
	}

	resp, err := virtualaccount.CreateFixedVA(&data)
	if err != nil {
		fmt.Println("Error occurred:", err.Error())
		return err
	}

	if resp != nil {
		fmt.Printf("created fixed va: %+v\n", resp)
	} else {
		fmt.Println("Response is nil")
	}
	mappedResp := payments.VirtualAccountObjectCore{
		ID:              resp.ID,
		ExternalID:      resp.ExternalID,
		OwnerID:         resp.OwnerID,
		MerchantCode:    resp.MerchantCode,
		BankCode:        resp.BankCode,
		AccountNumber:   resp.AccountNumber,
		Currency:        resp.Currency,
		IsSingleUse:     resp.IsSingleUse,
		IsClosed:        resp.IsClosed,
		ExpectedAmount:  resp.ExpectedAmount,
		SuggestedAmount: resp.SuggestedAmount,
		ExpirationDate:  resp.ExpirationDate,
		Name:            resp.Name,
		Description:     resp.Description,
		Status:          resp.Status,
	}
	errXen := service.PaymentData.CreateVA(mappedResp)
	if errXen != nil {
		fmt.Println("Failed to save in database:", err)
		return err
	}

	return nil
}

func New(repo payments.PaymentDataInterface, config *config.AppConfig) payments.PaymentServiceInterface {
	return &PaymentService{
		PaymentData: repo,
		config:      config,
	}
}
