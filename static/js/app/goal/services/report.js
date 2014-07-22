'use strict';

angular.module('eg.goal').factory('Report', ['$socketResource', function ($socketResource) {

    var Report = $socketResource('report');

    var service = {
        instantiate: function(raw) {
            return new Report(raw);
        }
    };
    return service;
}]);

