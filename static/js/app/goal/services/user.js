'use strict';

angular.module('eg.goal').factory('User', ['$rootScope', '$socketResource', '$resource', 'Category', 'Goal', 'Conclusion', function ($rootScope, $socketResource, $resource, Category, Goal, Conclusion) {

    var User = $socketResource('user');

//    var User = $resource('http://'+document.domain+'/ws/user', null,
//        {
//            'get': {method: 'GET'},
//            'save': {method: 'POST'},
//            'query': {method: 'GET', isArray: true},
//            'remove': {method: 'DELETE'},
//            'delete': {method: 'DELETE'}
//        }
//    );

    var userModel = new User;
    userModel.$get([1]);


    var service = {
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
        getAll: function () {
            return User;
        }
    };
    return service;
}]);
