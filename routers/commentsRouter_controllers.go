package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["memplus/controllers:AnkiCardController"] = append(beego.GlobalControllerRouter["memplus/controllers:AnkiCardController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:AnkiCardController"] = append(beego.GlobalControllerRouter["memplus/controllers:AnkiCardController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:AnkiCardController"] = append(beego.GlobalControllerRouter["memplus/controllers:AnkiCardController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:AnkiCardController"] = append(beego.GlobalControllerRouter["memplus/controllers:AnkiCardController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:AnkiCardController"] = append(beego.GlobalControllerRouter["memplus/controllers:AnkiCardController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:AnkiDeckController"] = append(beego.GlobalControllerRouter["memplus/controllers:AnkiDeckController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:AnkiDeckController"] = append(beego.GlobalControllerRouter["memplus/controllers:AnkiDeckController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:AnkiDeckController"] = append(beego.GlobalControllerRouter["memplus/controllers:AnkiDeckController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:AnkiDeckController"] = append(beego.GlobalControllerRouter["memplus/controllers:AnkiDeckController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:AnkiDeckController"] = append(beego.GlobalControllerRouter["memplus/controllers:AnkiDeckController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:CardController"] = append(beego.GlobalControllerRouter["memplus/controllers:CardController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:CardController"] = append(beego.GlobalControllerRouter["memplus/controllers:CardController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:CardController"] = append(beego.GlobalControllerRouter["memplus/controllers:CardController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:CardController"] = append(beego.GlobalControllerRouter["memplus/controllers:CardController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:CardController"] = append(beego.GlobalControllerRouter["memplus/controllers:CardController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:HistoryController"] = append(beego.GlobalControllerRouter["memplus/controllers:HistoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:HistoryController"] = append(beego.GlobalControllerRouter["memplus/controllers:HistoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:HistoryController"] = append(beego.GlobalControllerRouter["memplus/controllers:HistoryController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:HistoryController"] = append(beego.GlobalControllerRouter["memplus/controllers:HistoryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:HistoryController"] = append(beego.GlobalControllerRouter["memplus/controllers:HistoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "GetReadyTasks",
			Router: `/:id/ready_tasks`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "GetRootDirs",
			Router: `/roots`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KnowledgeBasesController"] = append(beego.GlobalControllerRouter["memplus/controllers:KnowledgeBasesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KnowledgeBasesController"] = append(beego.GlobalControllerRouter["memplus/controllers:KnowledgeBasesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KnowledgeBasesController"] = append(beego.GlobalControllerRouter["memplus/controllers:KnowledgeBasesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KnowledgeBasesController"] = append(beego.GlobalControllerRouter["memplus/controllers:KnowledgeBasesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KnowledgeBasesController"] = append(beego.GlobalControllerRouter["memplus/controllers:KnowledgeBasesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KnowledgeRelationsController"] = append(beego.GlobalControllerRouter["memplus/controllers:KnowledgeRelationsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KnowledgeRelationsController"] = append(beego.GlobalControllerRouter["memplus/controllers:KnowledgeRelationsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KnowledgeRelationsController"] = append(beego.GlobalControllerRouter["memplus/controllers:KnowledgeRelationsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KnowledgeRelationsController"] = append(beego.GlobalControllerRouter["memplus/controllers:KnowledgeRelationsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:KnowledgeRelationsController"] = append(beego.GlobalControllerRouter["memplus/controllers:KnowledgeRelationsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:LoopController"] = append(beego.GlobalControllerRouter["memplus/controllers:LoopController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:LoopController"] = append(beego.GlobalControllerRouter["memplus/controllers:LoopController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:LoopController"] = append(beego.GlobalControllerRouter["memplus/controllers:LoopController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:LoopController"] = append(beego.GlobalControllerRouter["memplus/controllers:LoopController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:LoopController"] = append(beego.GlobalControllerRouter["memplus/controllers:LoopController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TaskController"] = append(beego.GlobalControllerRouter["memplus/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TaskController"] = append(beego.GlobalControllerRouter["memplus/controllers:TaskController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TaskController"] = append(beego.GlobalControllerRouter["memplus/controllers:TaskController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TaskController"] = append(beego.GlobalControllerRouter["memplus/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TaskController"] = append(beego.GlobalControllerRouter["memplus/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TaskController"] = append(beego.GlobalControllerRouter["memplus/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Forget",
			Router: `/:id/forget`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TaskController"] = append(beego.GlobalControllerRouter["memplus/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Remember",
			Router: `/:id/remember`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TaskController"] = append(beego.GlobalControllerRouter["memplus/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Soso",
			Router: `/:id/soso`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus/controllers:TradeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus/controllers:TradeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Notify",
			Router: `/alipay_notify`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Return",
			Router: `/alipay_return`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Pay",
			Router: `/pay`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus/controllers:UserController"],
		beego.ControllerComments{
			Method: "ChangePwd",
			Router: `/change_pwd/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus/controllers:UserController"],
		beego.ControllerComments{
			Method: "Status",
			Router: `/status/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
