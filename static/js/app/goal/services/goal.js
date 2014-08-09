'use strict';

angular.module('eg.goal').factory('Goal', ['$socketResource', 'Report', function ($socketResource, Report) {
    var Goal = $socketResource('goal');

    var goals = [];
    var service = {
        instantiate: function(raw) {
            var goal = new Goal(raw);
            goal['today'].report = Report.instantiate(raw['today'].report);
            goal['yesterday'].report = Report.instantiate(raw['yesterday'].report);
            return goal;
        },
        set: function(goalsArray) {
            angular.forEach(goalsArray, function (raw) {
                goals.push(service.instantiate(raw));
            });
        },
        getAll: function () {
            return goals;
        },
        add: function(newGoalData) {
            var goal = Goal.instantiate(newGoalData);
            goal.$save(function(response) {
                goals.push(Goal.instantiate(response));
            });
        }
    };
    return service;
}]);

'use strict';

angular.module('eg.goal').factory('Conclusion', ['$socketResource', function ($socketResource) {
    var Conclusion = $socketResource('conclusion');

    var conclusions = [];
    var service = {
        set: function(goalsArray) {
            angular.forEach(goalsArray, function (val, key) {
                conclusions[key] = new Conclusion(val);
            });
        },
        getAll: function () {
            return conclusions;
        }
    };
    return service;
}]);

