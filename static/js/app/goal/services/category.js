'use strict';

angular.module('eg.goal').factory('Category', Category);

Category.$inject = ['$socketResource'];

function Category($socketResource) {

    var Category = $socketResource('goalCategory');

    var categories = [];

    var service = {
        set: set,
        getAll: getAll
    };
    return service;

    //////////////////

    function set(categoriesArray) {
        angular.forEach(categoriesArray, function (val) {
            var cat = new Category(val);
            categories.push(cat);
        });
    }

    function getAll() {
        return categories;
    }
}

