'use strict';

angular.module('eg.goal').factory('User', ['$rootScope', '$resource', 'Category', 'Goal', 'Conclusion', function ($rootScope, $resource, Category, Goal, Conclusion) {

    var User = $resource('/user', {}, {
        'query' : {
            url: "/user/list",
            isArray: true
        }
    });

    var service = {
        getInstance: function() {
            return new User();
        },
        instantiate: function (raw) {
            var user = new User(raw);
            Category.set(raw.categories);
            Goal.set(raw.goals);
            Conclusion.set(raw.conclusions);
            return user;
        },
        get: function (callback) {
            var user = service.instantiate(storage.init);

            callback(user);
        },
        getAll: function (cb) {
            User.query(cb);
        }
    };
    return service;
}]);
