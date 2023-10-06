package service

import (
	"Hannon-app/app/config"
	"Hannon-app/features/rents"
	"Hannon-app/features/users"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type RentService struct {
	rentData rents.RentDataInterface
	userData users.UserDataInterface
	config   *config.AppConfig
}

// Callback implements rents.RentServiceInterface.
func (service *RentService) Callback(input rents.RentCore) error {
	err := service.rentData.Callback(input)
	return err
}

// AcceptPayment implements rents.RentServiceInterface.
func (service *RentService) AcceptPayment(id uint, userid uint) error {
	rent, err := service.rentData.GetById(id)
	if rent.Status == "waiting payment" || rent.UserID != userid && err != nil {
		return errors.New("checkout failed, please pay previous invoice")
	}

	year, month, day := time.Now().Date()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()

	invoiceNum := fmt.Sprint("invoice/", rent.ID, "/", year, month, day, hour, minute, second)
	invoiceNum = strings.ReplaceAll(invoiceNum, " ", "")
	rent.InvoiceNumber = invoiceNum
	xendit.Opt.SecretKey = service.config.SecretKeyXendit

	userData, err := service.userData.SelectById(userid)
	if err != nil {
		return errors.New("user not found")
	}

	customer := xendit.InvoiceCustomer{
		GivenNames: userData.Name,
		Email:      userData.Email,
	}

	NotificationType := []string{"email"}

	customerNotificationPreference := xendit.InvoiceCustomerNotificationPreference{
		InvoiceCreated: NotificationType,
		InvoicePaid:    NotificationType,
		InvoiceExpired: NotificationType,
	}

	data := invoice.CreateParams{
		ExternalID:                     rent.InvoiceNumber,
		Amount:                         float64(rent.TotalPrice),
		Description:                    "Invoice Demo #123",
		InvoiceDuration:                3000,
		Customer:                       customer,
		CustomerNotificationPreference: customerNotificationPreference,
		Currency:                       "IDR",
	}

	resp, err := invoice.Create(&data)

	var paymentUrl = resp.InvoiceURL
	rent.PaymentLink = &paymentUrl
	rent.Status = "waiting payment"
	rent.IDXendit = resp.ID

	err = service.rentData.UpdatebyId(rent.ID, rent)

	return err
}

// ReadById implements rents.RentServiceInterface.
func (service *RentService) ReadById(id uint) (rents.RentCore, error) {
	result, err := service.rentData.GetById(id)
	return result, err
}

// UpdatebyId implements rents.RentServiceInterface.
func (service *RentService) UpdatebyId(id uint, input rents.RentCore) error {
	err := service.rentData.UpdatebyId(id, input)
	return err
}

// Add implements rents.RentServiceInterface.
func (service *RentService) Add(input rents.RentCore) error {
	err := service.rentData.Create(input)
	return err
}

func New(repo rents.RentDataInterface, userRepo users.UserDataInterface, config *config.AppConfig) rents.RentServiceInterface {
	return &RentService{
		rentData: repo,
		userData: userRepo,
		config:   config,
	}
}
