'use strict';

angular.module('eg.goal').factory('Article', Article);

Article.$inject = ['$resource'];

function Article($resource) {

    var Article = $resource('/article', {}, {
        query: {
            url: "/article/list",
            isArray: true
        }
    });

    return {
        instantiate: instantiate,
        New: New
    };

    /////////////

    function instantiate(raw) {
        return new Report(raw);
    }

    function New() {
        return new Article
    }
}



