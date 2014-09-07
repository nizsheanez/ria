'use strict';

angular.module('eg.goal').factory('User', User);

User.$inject = ['$resource', 'Category', 'Goal', 'Conclusion'];

function User($resource, Category, Goal, Conclusion) {

    var User = $resource('/user', {}, {
        'query': {
            url: "/user/list",
            isArray: true
        }
    });

    var service = {
        New: New,
        instantiate: instantiate,
        get: get,
        getAll: getAll
    };
    return service;

    ////////////////////

    function New() {
        return new User();
    }

    function instantiate(raw) {
        var user = new User(raw);
        Category.set(raw.categories);
        Goal.set(raw.goals);
        Conclusion.set(raw.conclusions);
        return user;
    }

    function get(callback) {
        var user = service.instantiate(storage.init);

        callback(user);
    }

    function getAll(cb) {
        User.query(cb);
    }
}
