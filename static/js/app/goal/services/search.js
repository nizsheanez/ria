'use strict';

angular.module('eg.goal', []).filter('search', function () {
    return function (items, name) {
        var arrayToReturn = [];
        for (var i = 0; i < items.length; i++) {
            if (items[i].name != name) {
                arrayToReturn.push(items[i]);
            }
        }

        return arrayToReturn;
    };
});