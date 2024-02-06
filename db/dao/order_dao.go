package dao

import (
	"context"
	"time"
	"wblevel0/db/entity"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderDao struct {
	dbPool *pgxpool.Pool
}

func NewOrderDao(dbPool *pgxpool.Pool) *OrderDao {
	return &OrderDao{
		dbPool: dbPool,
	}
}

func (e *OrderDao) GetAll() ([]*entity.OrderEntity, error) {
	query := `SELECT order_uid, track_number, entry, delivery_name, delivery_phone, delivery_zip, delivery_city, delivery_address,
	delivery_region, delivery_email, payment_transaction, payment_request_id, payment_currency, payment_provider,
	payment_amount, payment_payment_dt, payment_bank, payment_delivery_cost, payment_goods_total, payment_custom_fee,
	locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, oof_shard, date_created
	FROM "order"`
	rows, err := e.dbPool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*entity.OrderEntity, 0)

	for rows.Next() {
		var orderUID string
		var trackNumber string
		var entry string
		var deliveryName string
		var deliveryPhone string
		var deliveryZip string
		var deliveryCity string
		var deliveryAddress string
		var deliveryRegion string
		var deliveryEmail string
		var paymentTransaction string
		var paymentRequestId string
		var paymentCurrency string
		var paymentProvider string
		var paymentAmount int
		var paymentPaymentDt int
		var paymentBank string
		var paymentDeliveryCost int
		var paymentGoodsTotal int
		var paymentCustomFee int
		var locale string
		var internalSignature string
		var customerId string
		var deliveryService string
		var shardKey string
		var smId int
		var oofShard string
		var dateCreated time.Time
		err := rows.Scan(&orderUID, &trackNumber, &entry, &deliveryName, &deliveryPhone, &deliveryZip, &deliveryCity, &deliveryAddress,
			&deliveryRegion, &deliveryEmail, &paymentTransaction, &paymentRequestId, &paymentCurrency, &paymentProvider, &paymentAmount,
			&paymentPaymentDt, &paymentBank, &paymentDeliveryCost, &paymentGoodsTotal, &paymentCustomFee, &locale, &internalSignature,
			&customerId, &deliveryService, &shardKey, &smId, &oofShard, &dateCreated)
		if err != nil {
			return nil, err
		}
		order := &entity.OrderEntity{
			OrderUID:    orderUID,
			TrackNumber: trackNumber,
			Entry:       entry,
			Delivery: &entity.OrderDeliveryEntity{
				Name:    deliveryName,
				Phone:   deliveryPhone,
				Zip:     deliveryZip,
				City:    deliveryCity,
				Address: deliveryAddress,
				Region:  deliveryRegion,
				Email:   deliveryEmail,
			},
			Payment: &entity.OrderPaymentEntity{
				Transaction:  paymentTransaction,
				RequestId:    paymentRequestId,
				Currency:     paymentCurrency,
				Provider:     paymentProvider,
				Amount:       paymentAmount,
				PaymentDt:    paymentPaymentDt,
				Bank:         paymentBank,
				DeliveryCost: paymentDeliveryCost,
				GoodsTotal:   paymentGoodsTotal,
				CustomFee:    paymentCustomFee,
			},
			Locale:            locale,
			InternalSignature: internalSignature,
			CustomerId:        customerId,
			DeliveryService:   deliveryService,
			ShardKey:          shardKey,
			SmId:              smId,
			OofShard:          oofShard,
			DateCreated:       dateCreated,
		}
		result = append(result, order)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	orderIds := make([]string, len(result))

	for idx, order := range result {
		orderIds[idx] = order.OrderUID
	}

	orderItems, err := e.getOrderItems(orderIds)
	if err != nil {
		return nil, err
	}

	for _, order := range result {
		if items, ok := orderItems[order.OrderUID]; ok {
			order.Items = items
		}
	}

	return result, nil
}

func (e *OrderDao) getOrderItems(orderUids []string) (map[string][]*entity.OrderItemEntity, error) {
	query := `SELECT order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status
	FROM order_item
	WHERE order_uid = ANY($1)
	ORDER BY order_uid`

	rows, err := e.dbPool.Query(context.Background(), query, orderUids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string][]*entity.OrderItemEntity, 0)

	for rows.Next() {
		var orderUid string
		var chrtId int
		var trackNumber string
		var price int
		var rId string
		var name string
		var sale int
		var size string
		var totalPrice int
		var nmId int
		var brand string
		var status int
		err := rows.Scan(&orderUid, &chrtId, &trackNumber, &price, &rId, &name, &sale, &size, &totalPrice, &nmId, &brand, &status)
		if err != nil {
			return nil, err
		}

		orderItemEntity := &entity.OrderItemEntity{
			ChrtId:      chrtId,
			TrackNumber: trackNumber,
			Price:       price,
			RId:         rId,
			Name:        name,
			Sale:        sale,
			Size:        size,
			TotalPrice:  totalPrice,
			NmId:        nmId,
			Brand:       brand,
			Status:      status,
		}

		if orderItems, ok := result[orderUid]; ok {
			result[orderUid] = append(orderItems, orderItemEntity)
		} else {
			orderItems := make([]*entity.OrderItemEntity, 1)
			orderItems[0] = orderItemEntity
			result[orderUid] = orderItems
		}
	}

	return result, nil
}

func (e *OrderDao) Add(entity *entity.OrderEntity) error {
	ctx := context.Background()
	tx, err := e.dbPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	query := `INSERT INTO "order" (order_uid, track_number, entry, locale, internal_signature, customer_id, 
		delivery_service, shardkey, sm_id, oof_shard, delivery_name, delivery_phone, 
		delivery_zip, delivery_city, delivery_address, delivery_region, delivery_email, payment_transaction, 
		payment_request_id, payment_currency, payment_provider, payment_amount, payment_payment_dt, payment_bank, 
		payment_delivery_cost, payment_goods_total, payment_custom_fee, date_created)
		VALUES ($1, $2, $3, $4, $5, $6,
				$7, $8, $9::int, $10, $11, $12,
				$13, $14, $15, $16, $17, $18,
				$19, $20, $21, $22::int, $23::int, $24,
				$25::int, $26::int, $27::int, $28);`

	_, err = tx.Exec(ctx, query,
		entity.OrderUID, entity.TrackNumber, entity.Entry, entity.Locale, entity.InternalSignature, entity.CustomerId,
		entity.DeliveryService, entity.ShardKey, entity.SmId, entity.OofShard, entity.Delivery.Name, entity.Delivery.Phone, entity.Delivery.Zip,
		entity.Delivery.City, entity.Delivery.Address, entity.Delivery.Region, entity.Delivery.Email, entity.Payment.Transaction,
		entity.Payment.RequestId, entity.Payment.Currency, entity.Payment.Provider, entity.Payment.Amount, entity.Payment.PaymentDt,
		entity.Payment.Bank, entity.Payment.DeliveryCost, entity.Payment.GoodsTotal, entity.Payment.CustomFee, entity.DateCreated)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	for _, item := range entity.Items {
		query = `INSERT INTO order_item (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`
		_, err = tx.Exec(ctx, query, entity.OrderUID, item.ChrtId, item.TrackNumber, item.Price, item.RId, item.Name, item.Sale, item.Size,
			item.TotalPrice, item.NmId, item.Brand, item.Status)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	return nil
}
