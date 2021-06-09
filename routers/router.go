// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"BookStore/controllers/admincontroller"
	"BookStore/controllers/authorcontrolcontroller"
	"BookStore/controllers/bookcontroller"
	"BookStore/controllers/categorycontroller"
	"BookStore/controllers/commoncontroller"
	"BookStore/controllers/customercontroller"
	"BookStore/controllers/ordercontroller"
	"BookStore/controllers/paymentcontroller"
	"BookStore/controllers/permissioncontroller"
	"BookStore/controllers/publishercontroller"
	"BookStore/controllers/rolecontroller"
	"BookStore/controllers/rolepermissioncontrolcontroller"
	"BookStore/controllers/rolepermissioncontroller"
	"BookStore/controllers/roleusercontroller"
	"BookStore/controllers/vouchercontroller"
	"BookStore/controllers/websocketcontroller"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1/api/",
		beego.NSNamespace("/account/admins",
			beego.NSInclude(
				&admincontroller.AdminController{},
			),
		),
		beego.NSNamespace("/account/customer",
			beego.NSInclude(
				&customercontroller.CustomerController{},
			),
		),
		beego.NSNamespace("/orders",
			beego.NSInclude(
				&ordercontroller.OrderController{},
			),
		),
		beego.NSNamespace("/wsorders",
			beego.NSInclude(
				&websocketcontroller.OrderbookWSController{},
			),
		),
		beego.NSNamespace("/payments",
			beego.NSInclude(
				&paymentcontroller.PaymentController{},
			),
		),
		beego.NSNamespace("/vouchers",
			beego.NSInclude(
				&vouchercontroller.VoucherController{},
			),
		),
		beego.NSNamespace("/authorization/roles",
			beego.NSInclude(
				&rolecontroller.RoleController{},
			),
		),
		beego.NSNamespace("/authorization/role-user",
			beego.NSInclude(
				&roleusercontroller.RoleUserController{},
			),
		),
		beego.NSNamespace("/authorization/permissions",
			beego.NSInclude(
				&permissioncontroller.PermissionController{},
			),
		),
		beego.NSNamespace("/authorization/role-permissions",
			beego.NSInclude(
				&rolepermissioncontroller.RolePermissionController{},
			),
		),
		beego.NSNamespace("/authorization/author-controls",
			beego.NSInclude(
				&authorcontrolcontroller.AuthorControlController{},
			),
		),
		beego.NSNamespace("/authorization/role-permission-controls",
			beego.NSInclude(
				&rolepermissioncontrolcontroller.RolePermsControlController{},
			),
		),
		beego.NSNamespace("/books",
			beego.NSInclude(
				&bookcontroller.BookController{},
			),
		),
		beego.NSNamespace("/categories",
			beego.NSInclude(
				&categorycontroller.CategoryController{},
			),
		),
		beego.NSNamespace("/publishers",
			beego.NSInclude(
				&publishercontroller.PublisherController{},
			),
		),
		beego.NSNamespace("/common",
			beego.NSInclude(
				&commoncontroller.CommonController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
