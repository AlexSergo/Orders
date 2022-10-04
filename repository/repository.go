package repository

import (
	"L0/database"
	"L0/models"
	"database/sql"
)

const (
	order_insert   = `INSERT INTO orders(uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	delivery_isert = `INSERT INTO delivery(name, phone, zip, city, address, region, email, order_uid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	payment_insert = `INSERT INTO payment(transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, order_uid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	item_insert    = `INSERT INTO item(chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, order_uid) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	order_select    = `select * from orders`
	delivery_select = `select name,phone,zip,city,address,region,email from delivery where order_uid = $1`
	payment_select  = `select transaction,request_id,currency,provider,amount,payment_dt,bank,delivery_cost,goods_total,custom_fee from payment where order_uid = $1`
	item_select     = `select chrt_id,track_number,price,rid,name,sale,size,total_price,nm_id,brand,status from item where order_uid = $1`
)

func Insert(order models.Order) {
	db := database.Get()

	db.QueryRow(order_insert, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated, order.OofShard)
	db.QueryRow(delivery_isert, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email, order.OrderUID)
	db.QueryRow(payment_insert, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee, order.OrderUID)

	for _, item := range order.Items {
		db.QueryRow(item_insert, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status, order.OrderUID)
	}
}

func SelectOrders() []models.Order {
	orders := make([]models.Order, 0)
	database.Connect()
	db := database.Get()

	order_response, err := db.Query(order_select)
	if err != nil {
		panic(err)
	}
	defer order_response.Close()

	for order_response.Next() {
		order := models.Order{}

		order_response.Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale,
			&order.InternalSignature, &order.CustomerID, &order.DeliveryService,
			&order.ShardKey, &order.SmID, &order.DateCreated, &order.OofShard)

		order.Delivery = getDelivery(db, order.OrderUID)
		order.Payment = getPayment(db, order.OrderUID)
		order.Items = getItems(db, order.OrderUID)

		orders = append(orders, order)
	}

	return orders
}

func getDelivery(db *sql.DB, order_uid string) models.Delivery {
	delivery := models.Delivery{}

	delivery_response, err := db.Query(delivery_select, order_uid)
	if err != nil {
		panic(err)
	}
	defer delivery_response.Close()

	for delivery_response.Next() {
		delivery_response.Scan(&delivery.Name, &delivery.Phone, &delivery.Zip,
			&delivery.City, &delivery.Address, &delivery.Region, &delivery.Email)
	}

	return delivery
}

func getPayment(db *sql.DB, order_uid string) models.Payment {
	payment := models.Payment{}

	payment_response, err := db.Query(payment_select, order_uid)
	if err != nil {
		panic(err)
	}
	defer payment_response.Close()

	for payment_response.Next() {
		payment_response.Scan(&payment.Transaction, &payment.RequestID, &payment.Currency,
			&payment.Provider, &payment.Amount, &payment.PaymentDT, &payment.Bank,
			&payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee)
	}

	return payment
}

func getItems(db *sql.DB, order_uid string) []models.Item {
	items := make([]models.Item, 0)

	item_response, err := db.Query(item_select, order_uid)
	if err != nil {
		panic(err)
	}
	defer item_response.Close()

	for item_response.Next() {
		i := models.Item{}
		item_response.Scan(&i.ChrtID, &i.TrackNumber, &i.Price, &i.Rid, &i.Name,
			&i.Sale, &i.Size, &i.TotalPrice, &i.NmID, &i.Brand, &i.Status)
		items = append(items, i)
	}

	return items
}
