'use strict';

angular.module('eg.goal').factory('User', ['$rootScope', '$socketResource', '$resource', 'Category', 'Goal', 'Conclusion', function ($rootScope, $socketResource, $resource, Category, Goal, Conclusion) {

    var User = $socketResource('user');
    var User = $resource('http://blog3.ru/v2/user', null,
        {
            'get': {method: 'GET'},
            'save': {method: 'POST'},
            'query': {method: 'GET', isArray: true},
            'remove': {method: 'DELETE'},
            'delete': {method: 'DELETE'}
        }
    );
    User.get().$promise.then(function(user) {
        console.log(user);
    });

    var service = {
        instantiate: function(raw) {
            var user = new User(raw);
            Category.set(raw.categories);
            Goal.set(raw.goals);
            Conclusion.set(raw.conclusions);
            return user;
        } ,
        get: function (callback) {
            var user = service.instantiate(storage.init);

            callback();
        },
        getAll: function () {
            return User;
        }
    };
    return service;
}]);

