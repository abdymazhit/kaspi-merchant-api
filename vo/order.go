package vo

type OrderDeliveryMode string

const (
	OrderDeliveryModeLocal          OrderDeliveryMode = "DELIVERY_LOCAL"
	OrderDeliveryModePickup         OrderDeliveryMode = "DELIVERY_PICKUP"
	OrderDeliveryModeRegionalPickup OrderDeliveryMode = "DELIVERY_REGIONAL_PICKUP"
	OrderDeliveryModeRegionalTodoor OrderDeliveryMode = "DELIVERY_REGIONAL_TODOOR"
)

type OrdersDeliveryType string

const (
	OrdersDeliveryTypePickup   OrdersDeliveryType = "PICKUP"
	OrdersDeliveryTypeDelivery OrdersDeliveryType = "DELIVERY"
)

type OrderPaymentMode string

const (
	OrderPaymentModePayWithCredit OrderPaymentMode = "PAY_WITH_CREDIT"
	OrderPaymentModePrepaid       OrderPaymentMode = "PREPAID"
)

type OrdersState string

const (
	OrdersStateNew           OrdersState = "NEW"
	OrdersStateSignRequired  OrdersState = "SIGN_REQUIRED"
	OrdersStatePickup        OrdersState = "PICKUP"
	OrdersStateDelivery      OrdersState = "DELIVERY"
	OrdersStateKaspiDelivery OrdersState = "KASPI_DELIVERY"
	OrdersStateArchive       OrdersState = "ARCHIVE"
)

type OrdersStatus string

const (
	OrdersStatusApprovedByBank               OrdersStatus = "APPROVED_BY_BANK"
	OrdersStatusAcceptedByMerchant           OrdersStatus = "ACCEPTED_BY_MERCHANT"
	OrdersStatusCompleted                    OrdersStatus = "COMPLETED"
	OrdersStatusCancelled                    OrdersStatus = "CANCELLED"
	OrdersStatusCancelling                   OrdersStatus = "CANCELLING"
	OrdersStatusKaspiDeliveryReturnRequested OrdersStatus = "KASPI_DELIVERY_RETURN_REQUESTED"
	OrdersStatusReturnAcceptedByMerchant     OrdersStatus = "RETURN_ACCEPTED_BY_MERCHANT"
	OrdersStatusReturned                     OrdersStatus = "RETURNED"
)
