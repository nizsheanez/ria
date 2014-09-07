'use strict';

angular.module('eg.goal').factory('alertService', alertService);

alertService.$inject = ['$rootScope'];

function alertService($rootScope) {
    // create an array of alerts available globally
    $rootScope.alerts = [];

    var alertService = {
        add: add,
        closeAlert: closeAlert
    };

    $rootScope.$on('notification:add', function (event, data) {
        alertService.add(data.type, data.msg);
    });

    return alertService;

    function add(type, msg) {
        $rootScope.alerts.push({'type': type, 'msg': msg});
    }

     function closeAlert(index) {
        $rootScope.alerts.splice(index, 1);
    }

}