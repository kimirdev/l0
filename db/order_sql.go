package db

import (
	"l0/models"

	"github.com/jmoiron/sqlx"
)

type OrderSql struct {
	db *sqlx.DB
}

func NewOrderSql(db *sqlx.DB) *OrderSql {
	return &OrderSql{db: db}
}

func (s *OrderSql) GetAll() ([]models.OrderDTO, error) {
	var data []models.OrderDTO
	orderQuery := `select * from orders o`

	if err := s.db.Select(&data, orderQuery); err != nil {
		return nil, err
	}

	for i := 0; i < len(data); i++ {
		var delivery models.DeliveryDTO
		deliveryQuery := `select * from deliveries d where d.id = $1`

		if err := s.db.Get(&delivery, deliveryQuery, data[i].DeliveryId); err != nil {
			return nil, err
		}
		data[i].Delivery = delivery

		var payment models.PaymentDTO
		paymentQuery := `select * from payments p where p.id = $1`
		if err := s.db.Get(&payment, paymentQuery, data[i].PaymentId); err != nil {
			return nil, err
		}
		data[i].Payment = payment

		itemsOrderQuery := `select io.order_id, io.item_id from items_orders io where io.order_id = $1`

		var itemsOrders []models.ItemsOrdersDTO
		if err := s.db.Select(&itemsOrders, itemsOrderQuery, data[i].OrderId); err != nil {
			return nil, err
		}

		items := make([]models.ItemDTO, 0)
		for _, io := range itemsOrders {
			itemQuery := `select * from items i where i.id = $1`

			var item models.ItemDTO

			if err := s.db.Get(&item, itemQuery, io.ItemId); err != nil {
				return nil, err
			}
			items = append(items, item)
		}

		data[i].Items = items
	}
	return data, nil
}

func (s *OrderSql) Create(data models.OrderDTO) error {

	tx, err := s.db.Begin()

	if err != nil {
		return err
	}

	deliveryQuery := `INSERT INTO deliveries
	(name, phone, zip, city, address, region, email)
	VALUES ($1, $2, $3, $4, $5, $6, $7) 
	RETURNING id`

	rowDelivery := s.db.QueryRow(deliveryQuery, data.Delivery.Name, data.Delivery.Phone, data.Delivery.Zip, data.Delivery.City,
		data.Delivery.Address, data.Delivery.Region, data.Delivery.Email)

	var deliveryId int
	err = rowDelivery.Scan(&deliveryId)

	if err != nil {
		tx.Rollback()
		return err
	}

	paymentQuery := `INSERT INTO payments
	(transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id`

	rowPayment := s.db.QueryRow(paymentQuery, data.Payment.Transaction, data.Payment.RequestID, data.Payment.Currency,
		data.Payment.Provider, data.Payment.Amount, data.Payment.PaymentDt, data.Payment.Bank, data.Payment.DeliveryCost,
		data.Payment.GoodsTotal, data.Payment.CustomFee)

	var paymentId int
	err = rowPayment.Scan(&paymentId)

	if err != nil {
		tx.Rollback()
		return err
	}

	orderQuery := `INSERT INTO orders
	(order_uid, track_number, entry, local, internal_signature, customer_id, delivery_service, 
	shardkey, sm_id, date_created, oof_shard, delivery_id, payment_id)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	RETURNING id`

	rowOrder := s.db.QueryRow(orderQuery, data.OrderUID, data.TrackNumber, data.Entry, data.Local, data.InternalSignature,
		data.CustomerID, data.DeliveryService, data.Shardkey, data.SmID, data.DateCreated, data.OofShard, deliveryId, paymentId)

	var orderId int
	err = rowOrder.Scan(&orderId)

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range data.Items {
		itemQuery := `INSERT INTO items
		(chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id`

		rowItem := s.db.QueryRow(itemQuery, item.ChrtId, item.TrackNumber, item.Price, item.Rid,
			item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)

		var itemId int
		err = rowItem.Scan(&itemId)

		if err != nil {
			tx.Rollback()
			return err
		}

		orderItemQuery := `INSERT INTO items_orders
		(order_id, item_id)
		VALUES ($1, $2)`

		_, err := s.db.Exec(orderItemQuery, orderId, itemId)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (s *OrderSql) GetById(uid string) (models.OrderDTO, error) {
	var data models.OrderDTO

	orderQuery := `SELECT o.order_uid, o.track_number, o.entry, o.local, o.internal_signature,
	o.customer_id, o.delivery_service, o.shardkey, o.sm_id, o.date_created, o.oof_shard, 
	d.name, d.phone, d.zip, d.city, d.address, d.region, d.email, 
	p.transaction, p.request_id, p.currency, p.provider, p.amount, p.payment_dt,
	p.bank, p.delivery_cost, p.goods_total, p.custom_fee
		FROM orders o
		INNER JOIN deliveries d ON d.id = o.delivery_id
		INNER JOIN payments p ON p.id = o.payment_id
		WHERE o.order_uid = $1`

	rowOrder := s.db.QueryRow(orderQuery, uid)

	rowOrder.Scan(&data.OrderUID, &data.TrackNumber, &data.Entry, &data.Local, &data.InternalSignature,
		&data.CustomerID, &data.DeliveryService, &data.Shardkey, &data.SmID, &data.DateCreated, &data.OofShard,
		&data.Delivery.Name, &data.Delivery.Phone, &data.Delivery.Zip, &data.Delivery.City, &data.Delivery.Address, &data.Delivery.Region,
		&data.Delivery.Email, &data.Payment.Transaction, &data.Payment.RequestID, &data.Payment.Currency, &data.Payment.Provider,
		&data.Payment.Amount, &data.Payment.PaymentDt, &data.Payment.Bank, &data.Payment.DeliveryCost, &data.Payment.GoodsTotal,
		&data.Payment.CustomFee,
	)

	itemQuery := `SELECT i.chrt_id, i.track_number, i.price, i.rid, i.name, i.sale, i.size, i.total_price,
	i.nm_id, i.brand, i.status FROM items i
	INNER JOIN items_orders io ON io.item_id = i.id
	INNER JOIN orders o ON o.id = io.order_id
	WHERE o.order_uid = $1`

	rowItem, err := s.db.Query(itemQuery, uid)

	if err != nil {
		return data, err
	}

	data.Items = make([]models.ItemDTO, 0)
	for rowItem.Next() {
		var item models.ItemDTO
		err = rowItem.Scan(
			&item.ChrtId, &item.TrackNumber, &item.Price, &item.Rid, &item.Name,
			&item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status,
		)
		if err != nil {
			return data, err
		}
		data.Items = append(data.Items, item)
	}

	return data, nil
}
