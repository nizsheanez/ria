'use strict';

var lang = 'ru';
/*
 angular.module('MapApp', ['ui.router', 'ui.bootstrap', 'pascalprecht.translate']).value('lang', lang);

 angular.module('MapApp').config(['$translateProvider', function($translateProvider) {
 // add translation table
 $translateProvider.translations(translations);
 }]);
 */

var translations = [];


angular.module('eg.goal', [
    'eg.components',
    'ngRoute',
    'ngResource',
    'ui.bootstrap',
    'monospaced.elastic',
    'ngDebounce',
    'textAngular'
]);

angular.module('eg.goal').config(['msdElasticConfig', function (config) {
    config.append = '\n\n';
}]);

angular.module('eg.goal').config(['$routeProvider', '$locationProvider', function ($routeProvider, $locationProvider) {

    var dir = '/static/js/app/goal/views';

    $routeProvider
        .when('/', {templateUrl: dir + '/goals.html', controller: 'GoalCtrl'})
        .when('/yesterday', {templateUrl: dir + '/goals.html', controller: 'GoalCtrl'})
        .when('/active', {templateUrl: dir + '/goals.html', controller: 'GoalCtrl'})
        .when('/goal/:id', {templateUrl: dir + '/goals.html', controller: 'GoalCtrl'})
        .when('/news/', {templateUrl: dir + '/news.html', controller: 'NewsCtrl'});

    $locationProvider
        .html5Mode(true)
        .hashPrefix('!');

//    $urlRouterProvider.otherwise({
//        redirectTo: '/'
//    });
}]);


angular.module('eg.goal').run(function ($rootScope, $templateCache, $compile) {

    $rootScope.textAngularTools = {
        checkbox: {
            display: "<button type='button' ng-click='action()' ><i class='fa fa-check-square-o'></i></button>",
            action: function () {

                function insertNodeAtCursor(node) {
                    var sel, range, html;
                    if (window.getSelection) {
                        sel = window.getSelection();
                        if (sel.getRangeAt && sel.rangeCount) {
                            sel.getRangeAt(0).insertNode(node);
                        }
                    } else if (document.selection && document.selection.createRange) {
                        range = document.selection.createRange();
                        html = (node.nodeType == 3) ? node.data : node.outerHTML;
                        range.pasteHTML(html);
                    }
                }

                function placeCaretAtEnd(el) {
                    el.focus();
                    if (typeof window.getSelection != "undefined"
                        && typeof document.createRange != "undefined") {
                        var range = document.createRange();
                        range.setStartAfter(el);
                        range.setEndAfter(el);
                        var sel = window.getSelection();
                        sel.removeAllRanges();
                        sel.addRange(range);
                    } else if (typeof document.body.createTextRange != "undefined") {
                        var textRange = document.body.createTextRange();
                        textRange.moveToElementText(el);
                        textRange.collapse(false);
                        textRange.select();
                    }
                }


                var element = $compile('<input type="checkbox" eg-todo />')(this.$parent);

                this.$parent.displayElements.text[0].focus();
                insertNodeAtCursor(element[0]);
                placeCaretAtEnd(element[0])
                return this.$parent.wrapSelection("insertHTML", "&nbsp;");
            }
        },
        html: {
            display: "<button type='button' ng-click='action()' ng-class='displayActiveToolClass(active)'>&lg;</button>"
        },
        ul: {
            display: "<button type='button' ng-click='action()' ng-class='displayActiveToolClass(active)'><i class='fa fa-list-ul'></i></button>"
        },
        ol: {
            display: "<button type='button' ng-click='action()' ng-class='displayActiveToolClass(active)'><i class='fa fa-list-ol'></i></button>"
        },
        quote: {
            display: "<button type='button' ng-click='action()' ng-class='displayActiveToolClass(active)'><i class='fa fa-quote-right'></i></button>"
        },
        undo: {
            display: "<button type='button' ng-click='action()' ng-class='displayActiveToolClass(active)'><i class='fa fa-undo'></i></button>"
        },
        redo: {
            display: "<button type='button' ng-click='action()' ng-class='displayActiveToolClass(active)'><i class='fa fa-repeat'></i></button>"
        },
        bold: {
            display: "<button type='button' ng-click='action()' ng-class='displayActiveToolClass(active)'><i class='fa fa-bold'></i></button>"
        },
        justifyLeft: {
            display: "<button type='button' ng-click='action()' ng-class='displayActiveToolClass(active)'><i class='fa fa-align-left'></i></button>"
        },
        justifyRight: {
            display: "<button type='button' ng-click='action()' ng-class='displayActiveToolClass(active)'><i class='fa fa-align-right'></i></button>"
        },
        justifyCenter: {
            display: "<button type='button' ng-click='action()' ng-class='displayActiveToolClass(active)'><i class='fa fa-align-center'></i></button>"
        },
        italics: {
            display: "<button type='button' ng-click='action()' ng-class='displayActiveToolClass(active)'><i class='fa fa-italic'></i></button>"
        },
        clear: {
            display: "<button type='button' ng-click='action()' ng-class='displayActiveToolClass(active)'><i class='fa fa-ban'></i></button>"
        },
        insertImage: {
            display: "<button type='button' ng-click='action()' ng-class='displayActiveToolClass(active)'><i class='fa fa-picture-o'></i></button>"
        },
        insertLink: {
            display: "<button type='button' ng-click='action()' ng-class='displayActiveToolClass(active)'><i class='fa fa-chain'></i></button>"
        }
    };

    $rootScope.textAngularOpts = {
        toolbar: [
            ['checkbox', 'bold', 'italics', 'ul', 'ol', 'redo', 'undo', 'clear', 'insertImage', 'insertLink', 'html']
        ],
        classes: {
            toolbar: "btn-toolbar",
            toolbarGroup: "btn-group",
            toolbarButton: "btn",
            toolbarButtonActive: "active",
            textEditor: 'eg-form-control',
            htmlEditor: 'eg-form-control'
        }
    }
});

var showScreen = function () {
    $(document).ready(function () {
        var el = $('.first-load:first');
        el.remove();
    });
};
setTimeout(showScreen, 500);

