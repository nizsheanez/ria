package conf

import (
	"github.com/astaxie/beego"
	"ria/components"
)
var css string
var js string

func init() {
	assets := AppAsset()
	var err error = nil

	css, err = assets.Css();
	js, err = assets.Js();

	if err != nil {
		beego.Error(err)
	}
	if err != nil {
		beego.Error(err)
	}
}

func GetCss() string {
	return css
}

func GetJs() string {
	return js
}


func AppAsset() *components.Asset {
	jquery := components.NewAsset().
		SetBaseUrl("static/vendor").
		SetJs(
			"jquery/dist/jquery.min.js",
		)

	angular := components.NewAsset().
		SetBaseUrl("static/vendor").
		SetJs(
			"angular/angular.js",
			"angular-resource/angular-resource.js",
			"angular-route/angular-route.min.js",
			// "angular-ui-router/release/angular-ui-router.min.js",
			"angular-translate/angular-translate.js",
		)

	angularUi := components.NewAsset().
		SetBaseUrl("static/vendor").
		SetCss(
			"angular-ui/build/angular-ui.min.css",
		).
		SetJs(
			"angular-ui/build/angular-ui.min.js",
			"angular-ui-sortable/src/sortable.js",
		)

	uiBootstrap := components.NewAsset().
		SetBaseUrl("static/vendor/angular-bootstrap").
		SetJs(
			"ui-bootstrap.min.js",
			"ui-bootstrap-tpls.min.js",
		)

	angularElastic := components.NewAsset().
		SetBaseUrl("static/vendor/angular-elastic").
		SetJs(
			"elastic.js",
		)

	angularUiUtils := components.NewAsset().
		SetBaseUrl("static/vendor/angular-ui-utils").
		SetJs(
			"ui-utils.min.js",
		)

//	when := components.NewAsset().
//		SetBaseUrl("static/vendor/when").
//		SetJs(
//		"when.js",
//		)

	autobahn := components.NewAsset().
		SetBaseUrl("static/js/common").
		SetJs(
			"autobahn2.js",
			"socketResource.js",
		)
//		.SetDependencies(when)


	sockjs := components.NewAsset().
		SetBaseUrl("static/vendor/sockjs").
		SetJs(
			"sockjs.js",
		)

	app := components.NewAsset().
		SetBaseUrl("static").
		SetCss(
			"assets/build/css/site.css",
		).
		SetJs(
			"js/common/fixes.js",
			"js/common/debug.js",
			"js/common/components.js",
			"js/app/app.js",
			"js/app/goal/services/tpl.js",
			"js/app/goal/services/modal.js",
			"js/app/goal/services/user.js",
			"js/app/goal/services/category.js",
			"js/app/goal/services/report.js",
			"js/app/goal/services/goal.js",
			"js/app/goal/controllers/goal.js",
			"js/app/goal/controllers/news.js",
			"js/app/goal/controllers/nav.js",
			"js/app/goal/directives/editor.js",
			"js/app/goal/services/alert.js",
		).
		SetDependencies(
			jquery,
			angular,
			autobahn,
			sockjs,
			angularUi,
			uiBootstrap,
			angularElastic,
			angularUiUtils,
			textArea(),
		)

	return app
}

func textArea() *components.Asset {
	angularSanitize := components.NewAsset().
		SetBaseUrl("static/vendor/angular-sanitize").
		SetJs(
			"angular-sanitize.js",
		)

	fontAwesome := components.NewAsset().
		SetBaseUrl("static/vendor/components-font-awesome").
		SetCss(
			"css/font-awesome.min.css",
		)

	textAngular := components.NewAsset().
		SetBaseUrl("static/vendor/textAngular").
		SetJs(
			"textAngular.js",
		).
		SetDependencies(
			angularSanitize,
			fontAwesome,
		)

	return textAngular
}
