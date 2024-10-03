package services

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/sabareeswaran-sakthivel/e-commerce/models"
	"github.com/sabareeswaran-sakthivel/e-commerce/repository"
)

type UploadCSVDataService struct {
	or  repository.OrdersRepository
	pr  repository.ProductsRepository
	ord repository.OrderDetailsRepository
	re  repository.RegionsRepository
	ca  repository.CategoriesRepository
	cr  repository.CustomersRepository
}

func (uc *UploadCSVDataService) UploadCSVData(fileName string) (string, error) {
	fmt.Println("Uploading CSV Data")
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records[1:] {
		var (
			orderId, quantitySold             int
			unitPrice, shippingCost, discount float64
		)

		customerId := record[2]

		orderId, err = strconv.Atoi(record[0])
		if err != nil {
			panic(err)
		}

		productId := record[1]

		quantitySold, err = strconv.Atoi(record[7])
		if err != nil {
			panic(err)
		}

		discount, err = strconv.ParseFloat(record[9], 64)
		if err != nil {
			panic(err)
		}

		shippingCost, err = strconv.ParseFloat(record[10], 64)
		if err != nil {
			panic(err)
		}

		unitPrice, err = strconv.ParseFloat(record[8], 64)
		if err != nil {
			panic(err)
		}

		regionName := record[5]
		region, err := uc.re.GetRegionByName(regionName)
		if err != nil {
			fmt.Println(err)
		}
		category := uc.createCategory(record[4])

		uc.createOrder(orderId, customerId, record[6], record[11], shippingCost, region.Id)
		uc.createOrderDetails(orderId, productId, quantitySold, discount)
		uc.createProduct(productId, record[3], "", category.Id, unitPrice)
		uc.createCustomer(customerId, record[12], record[13], record[14])
	}
	return "Data Uploaded Successfully", nil
}

func (uc *UploadCSVDataService) createOrder(id int, customerId string, dateOfSale string, paymentMethod string, shippingCost float64, regionId int) {
	order := models.Order{
		Id:            id,
		CustomerId:    customerId,
		DateOfSale:    dateOfSale,
		PaymentMethod: paymentMethod,
		ShippingCost:  shippingCost,
		RegionId:      regionId,
	}
	err := uc.or.Create(order)
	if err != nil {
		fmt.Println(err)
	}
}

func (uc *UploadCSVDataService) createProduct(id string, name string, description string, categoryId string, unitPrice float64) {
	product, err := uc.pr.GetByName(name)
	if err != nil {
		fmt.Println(err)
	}
	if product.Id == "" {
		product := models.Product{
			Id:          id,
			Name:        name,
			Description: description,
			CategoryId:  categoryId,
			UnitPrice:   unitPrice,
		}
		err := uc.pr.Create(product)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (uc *UploadCSVDataService) createCustomer(id string, name string, email string, address string) {
	customer, err := uc.cr.GetCustomerByEmail(email)
	if err != nil {
		fmt.Println(err)
	}
	if customer.Id == "" {
		customer := models.Customer{
			Id:       id,
			Name:     name,
			Email:    email,
			Address:  address,
			IsActive: true,
			Age:      nil,
		}
		err := uc.cr.Create(customer)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (uc *UploadCSVDataService) createOrderDetails(orderId int, productId string, quantitySold int, discount float64) {
	orderDetails := models.OrderDetails{
		OrderId:      orderId,
		ProductId:    productId,
		QuantitySold: quantitySold,
		Discount:     discount,
	}
	err := uc.ord.Create(orderDetails)
	if err != nil {
		fmt.Println(err)
	}
}

func (uc *UploadCSVDataService) createCategory(name string) models.Category {
	category, err := uc.ca.GetCategoryByName(name)
	if err != nil {
		fmt.Println(err)
	}
	if category.Id == "" {
		category.Id = uuid.New().String()
		category.Name = name
		err = uc.ca.Create(category)
		if err != nil {
			fmt.Println(err)
		}
	}
	return category
}
