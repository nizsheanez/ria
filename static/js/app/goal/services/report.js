'use strict';

angular.module('eg.goal').factory('Report', Report);

Report.$inject = ['$socketResource'];

function Report($socketResource) {

    var Report = $socketResource('report');

    var service = {
        instantiate: instantiate
    };
    return service;

    //////////////////
    function instantiate(raw) {
        return new Report(raw);
    }
}

