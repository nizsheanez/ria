'use strict';

angular.module('eg.goal').factory('Article', ['$resource', function ($resource) {

    var Article = $resource('/article', {}, {
        query : {
            url: "/article/list",
            isArray: true
        }
    });

    var service = {
        instantiate: function(raw) {
            return new Report(raw);
        },
        New: function() {
            return new Article
        }
    };
    return service;
}]);

