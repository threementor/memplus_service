package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["memplus_service/controllers:AnkiCardController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:AnkiCardController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:AnkiCardController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:AnkiCardController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:AnkiCardController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:AnkiCardController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:AnkiCardController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:AnkiCardController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:AnkiCardController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:AnkiCardController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:AnkiDeckController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:AnkiDeckController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:AnkiDeckController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:AnkiDeckController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:AnkiDeckController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:AnkiDeckController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:AnkiDeckController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:AnkiDeckController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:AnkiDeckController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:AnkiDeckController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:CardController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:CardController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:CardController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:CardController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:CardController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:CardController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:CardController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:CardController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:CardController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:CardController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:CardController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:CardController"],
		beego.ControllerComments{
			Method: "Forget",
			Router: `/:id/forget`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:CardController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:CardController"],
		beego.ControllerComments{
			Method: "UpdateNote",
			Router: `/:id/note`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:CardController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:CardController"],
		beego.ControllerComments{
			Method: "Remember",
			Router: `/:id/remember`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:CardController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:CardController"],
		beego.ControllerComments{
			Method: "Soso",
			Router: `/:id/soso`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:HistoryController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:HistoryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:HistoryController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:HistoryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:HistoryController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:HistoryController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:HistoryController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:HistoryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:HistoryController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:HistoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "AddCardToDeck",
			Router: `/:id/create/card`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "GetReadyTasks",
			Router: `/:id/ready_cards`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KlgDirController"],
		beego.ControllerComments{
			Method: "GetRootDirs",
			Router: `/roots`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeBasesController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeBasesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeBasesController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeBasesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeBasesController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeBasesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeBasesController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeBasesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeBasesController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeBasesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeRelationsController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeRelationsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeRelationsController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeRelationsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeRelationsController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeRelationsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeRelationsController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeRelationsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeRelationsController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:KnowledgeRelationsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:LoopController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:LoopController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:LoopController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:LoopController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:LoopController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:LoopController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:LoopController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:LoopController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:LoopController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:LoopController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:NoteController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:NoteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:NoteController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:NoteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:NoteController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:NoteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:NoteController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:NoteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:NoteController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:NoteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:ProductController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:ProductController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:ProductController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:ProductController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:ProductController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:ProductController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:ProductController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:ProductController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:ProductController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:ProductController"],
		beego.ControllerComments{
			Method: "SubmitTrade",
			Router: `/:id/submit_trade`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:TradeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:TradeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Pay",
			Router: `/:id/pay`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Notify",
			Router: `/alipay_notify`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:TradeController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:TradeController"],
		beego.ControllerComments{
			Method: "Return",
			Router: `/alipay_return`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "ChangePwd",
			Router: `/change_pwd/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["memplus_service/controllers:UserController"] = append(beego.GlobalControllerRouter["memplus_service/controllers:UserController"],
		beego.ControllerComments{
			Method: "Status",
			Router: `/status/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
