'use strict';

angular.module('eg.goal').factory('Goal', Goal);

Goal.$inject = ['$socketResource', 'Report']

function Goal($socketResource, Report) {
    var Goal = $socketResource('goal');

    var goals = [];
    var service = {
        instantiate: instantiate,
        set: set,
        getAll: getAll,
        add: add
    };
    return service;

    ///////////////

    function instantiate(raw) {
        var goal = new Goal(raw);
        goal['today'].report = Report.instantiate(raw['today'].report);
        goal['yesterday'].report = Report.instantiate(raw['yesterday'].report);
        return goal;
    }

    function set(goalsArray) {
        angular.forEach(goalsArray, function (raw) {
            goals.push(service.instantiate(raw));
        });
    }

    function getAll() {
        return goals;
    }

    function add(newGoalData) {
        var goal = Goal.instantiate(newGoalData);
        goal.$save(function (response) {
            goals.push(Goal.instantiate(response));
        });
    }
}

'use strict';

angular.module('eg.goal').factory('Conclusion', Conclusion);

Conclusion.$inject = ['$socketResource']

function Conclusion($socketResource) {
    var Conclusion = $socketResource('conclusion');

    var conclusions = [];
    var service = {
        set: set,
        getAll: getAll
    };
    return service;

    function set(goalsArray) {
        angular.forEach(goalsArray, function (val, key) {
            conclusions[key] = new Conclusion(val);
        });
    }

    function getAll() {
        return conclusions;
    }
}

