package consts

const (
	ADMIN_MODEL              = "Admin"
	AUTHOR_MODEL             = "Author"
	AWARD_MODEL              = "Award"
	BOOK_AUTHOR_MODEL        = "Book_Author"
	BOOK_AWARD_MODEL         = "Book_Award"
	BOOK_MODEL               = "Book"
	BOOK_SPECIFICATION_MODEL = "Book_Specification"
	CATEGORY_MODEL           = "Category"
	COMMENT_MODEL            = "Comment"
	CUSTOMER_MODEL           = "Customer"
	DELIVERY_MODEL           = "Delivery"
	FORMAT_MODEL         = "Format"
	ORDER_DETAIL_MODEL   = "Order_Detail"
	ORDER_MODEL          = "Order"
	PAYMENT_MODEL        = "Payment"
	PUBLISHER_MODEL      = "Publisher"
	SUBTOPIC_MODEL       = "Subtopic"
	TOPIC_MODEL          = "Topic"
	VOUCHER_MODEL        = "Voucher"
	ROLE_MODEL           = "Role"
	USER_ROLE_MODEL = "UserRole"
	PERMISSION_MODEL     = "Permission"
	ROUTER_PAGE_MODEL    = "RouterPage"
	AUTHOR_CONTROL_MODEL = "AuthorControl"
	BOOK_TYPE_MODEL = "BookType"
	ORDER_VOUCHER_MODEL = "OrderVoucher"
)

const (
	ACCESS_SECRET = "access_book_store_secret_login"
	REFRESH_SECRET = "refresh_book_store_secret_login"
	SESSION_USER = "SessionUser"
)
var MAP_BOOK_TYPE = map[int]string{
	1: "Non series",
	2: "Series",
}