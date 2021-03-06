'use strict';

angular.module('eg.goal').factory('Tpl', Tpl);

function Tpl() {
    var tplBase = '/static/js//app/goal/';

    var service = {
        modal: {
            edit: tplBase + 'views/modal/edit.html',
            backLog: tplBase + 'views/modal/back_log.html'
        },
        grid: tplBase + 'views/goal_grid.html',
        goals: tplBase + 'views/goals.html',
        sidebar: tplBase + 'views/sidebar.html',
        news: {
            grid: tplBase + 'views/news_grid.html'
        }
    };
    return service;
}