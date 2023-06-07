package valueobject

type OrderDeliveryMode string

const (
	OrderDeliveryModeLocal          OrderDeliveryMode = "DELIVERY_LOCAL"
	OrderDeliveryModePickup                           = "DELIVERY_PICKUP"
	OrderDeliveryModeRegionalPickup                   = "DELIVERY_REGIONAL_PICKUP"
	OrderDeliveryModeRegionalTodoor                   = "DELIVERY_REGIONAL_TODOOR"
)

type OrdersDeliveryType string

const (
	OrdersDeliveryTypePickup   OrdersDeliveryType = "PICKUP"
	OrdersDeliveryTypeDelivery                    = "DELIVERY"
)

type OrderPaymentMode string

const (
	OrderPaymentModePayWithCredit OrderPaymentMode = "PAY_WITH_CREDIT"
	OrderPaymentModePrepaid                        = "PREPAID"
)

type OrdersState string

const (
	OrdersStateNew           OrdersState = "NEW"
	OrdersStateSignRequired              = "SIGN_REQUIRED"
	OrdersStatePickup                    = "PICKUP"
	OrdersStateDelivery                  = "DELIVERY"
	OrdersStateKaspiDelivery             = "KASPI_DELIVERY"
	OrdersStateArchive                   = "ARCHIVE"
)

type OrdersStatus string

const (
	OrdersStatusApprovedByBank               OrdersStatus = "APPROVED_BY_BANK"
	OrdersStatusAcceptedByMerchant                        = "ACCEPTED_BY_MERCHANT"
	OrdersStatusCompleted                                 = "COMPLETED"
	OrdersStatusCancelled                                 = "CANCELLED"
	OrdersStatusCancelling                                = "CANCELLING"
	OrdersStatusKaspiDeliveryReturnRequested              = "KASPI_DELIVERY_RETURN_REQUESTED"
	OrdersStatusReturnAcceptedByMerchant                  = "RETURN_ACCEPTED_BY_MERCHANT"
	OrdersStatusReturned                                  = "RETURNED"
)
