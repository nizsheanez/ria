'use strict';

angular.module('eg.goal').factory('alertService', function ($rootScope) {
    var alertService = {};


    // create an array of alerts available globally
    $rootScope.alerts = [];

    alertService.add = function (type, msg) {
        $rootScope.alerts.push({'type': type, 'msg': msg});
    };

    alertService.closeAlert = function (index) {
        $rootScope.alerts.splice(index, 1);
    };

    $rootScope.$on('notification:add', function (event, data) {
        alertService.add(data.type, data.msg);
    });

    return alertService;
});