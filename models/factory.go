package models

import "BookStore/consts"

type Factory struct {
	AdminSV             RepoServiceIf
	AuthorSV            RepoServiceIf
	AwardSV             RepoServiceIf
	BookAuthorSV        RepoServiceIf
	BookAwardSV         RepoServiceIf
	BookSV              RepoServiceIf
	BookSpecificationSV RepoServiceIf
	CategorySV          RepoServiceIf
	CommentSV           RepoServiceIf
	CustomerSV          RepoServiceIf
	DeliverySV          RepoServiceIf
	FormatSV            RepoServiceIf
	OrderDetailSV       RepoServiceIf
	OrderSV             RepoServiceIf
	PaymentSV           RepoServiceIf
	PublisherSV         RepoServiceIf
	SubtopicSV          RepoServiceIf
	TopicSV             RepoServiceIf
	VoucherSV           RepoServiceIf
}

func (f *Factory) NewFactory(_type string) RepoServiceIf {
	switch _type {
	case consts.ADMIN_MODEL:
		return &Admin{}
	case consts.AUTHOR_MODEL:
		return &Author{}
	case consts.AWARD_MODEL:
		return &Award{}
	case consts.BOOK_AUTHOR_MODEL:
		return &BookTypeAuthor{}
	case consts.BOOK_AWARD_MODEL:
		return &BookTypeAward{}
	case consts.BOOK_MODEL:
		return &Book{}
	case consts.BOOK_SPECIFICATION_MODEL:
		return &BookDetail{}
	case consts.CATEGORY_MODEL:
		return &Category{}
	case consts.COMMENT_MODEL:
		return &Comment{}
	case consts.CUSTOMER_MODEL:
		return &Customer{}
	case consts.DELIVERY_MODEL:
		return &Delivery{}
	case consts.FORMAT_MODEL:
		return &Format{}
	case consts.ORDER_DETAIL_MODEL:
		return &OrderDetail{}
	case consts.ORDER_MODEL:
		return &Order{}
	case consts.PAYMENT_MODEL:
		return &Payment{}
	case consts.PUBLISHER_MODEL:
		return &Publisher{}
	case consts.TOPIC_MODEL:
		return &Topic{}
	case consts.VOUCHER_MODEL:
		return &Voucher{}
	case consts.ROLE_MODEL:
		return &Role{}
	case consts.USER_ROLE_MODEL:
		return &AdminRole{}
	case consts.PERMISSION_MODEL:
		return &Permission{}
	case consts.ROUTER_PAGE_MODEL:
		return &RolePermissions{}
	case consts.AUTHOR_CONTROL_MODEL:
		return &AuthorControl{}
	case consts.BOOK_TYPE_MODEL:
		return &BookType{}
	case consts.ORDER_VOUCHER_MODEL:
		return &OrderVouchers{}
	default:
		return nil
	}
}

func NewFactoryInstance() RepoFactoryIf {
	f := &Factory{}
	f.AdminSV = f.NewFactory(consts.ADMIN_MODEL)
	f.AuthorSV = f.NewFactory(consts.AUTHOR_MODEL)
	f.AwardSV = f.NewFactory(consts.AWARD_MODEL)
	f.BookAuthorSV = f.NewFactory(consts.BOOK_AUTHOR_MODEL)
	f.BookAwardSV = f.NewFactory(consts.BOOK_AWARD_MODEL)
	f.BookSV = f.NewFactory(consts.BOOK_MODEL)
	f.BookSpecificationSV = f.NewFactory(consts.BOOK_SPECIFICATION_MODEL)
	f.CategorySV = f.NewFactory(consts.CATEGORY_MODEL)
	f.CommentSV = f.NewFactory(consts.COMMENT_MODEL)
	f.CustomerSV = f.NewFactory(consts.CUSTOMER_MODEL)
	f.DeliverySV = f.NewFactory(consts.DELIVERY_MODEL)
	f.FormatSV = f.NewFactory(consts.FORMAT_MODEL)
	f.OrderDetailSV = f.NewFactory(consts.ORDER_DETAIL_MODEL)
	f.OrderSV = f.NewFactory(consts.ORDER_MODEL)
	f.PaymentSV = f.NewFactory(consts.PAYMENT_MODEL)
	f.PublisherSV = f.NewFactory(consts.PUBLISHER_MODEL)
	f.SubtopicSV = f.NewFactory(consts.SUBTOPIC_MODEL)
	f.TopicSV = f.NewFactory(consts.TOPIC_MODEL)
	f.VoucherSV = f.NewFactory(consts.VOUCHER_MODEL)
	return f
}
