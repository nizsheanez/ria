'use strict';

angular.module('eg.goal').controller('NewsCtrl', function ($q, $http, $scope, $socketResource, $routeParams, $location, $modal, Tpl, User, Category, Goal, Conclusion, Modal) {

    $scope.tpl = Tpl;
    $scope.users = [];

    User.getAll(function (result) {
        console.log(result)
    });

//    $scope.focusGoal = false;
//    $scope.setFocus = function (goal) {
//        $scope.focusGoal = goal;
//    };
//
//    $scope.inFocus = function (goal) {
//        return goal.id === $scope.focusGoal.id;
//    }
//
//    $scope.defaultPlaceholder = 'Сделано';
//
//    $scope.save = function (model) {
//        model.$save();
//    };
//
//    if ($location.path() === '/') {
//        $scope.day = 'today';
//    } else if ($location.path() === '/yesterday') {
//        $scope.day = 'yesterday';
//    }
//
//    $scope.location = $location;
//
//    $scope.editGoalModal = Modal.editGoalModal;
//    $scope.addGoalModal = Modal.addGoalModal;
//    $scope.backLogModal = Modal.backLogModal;
//
//    $scope.complete = function (goal) {
//        goal.completed = 1;
//        goal.$save();
//    };
//
//    $scope.fail = function (goal) {
//        goal.completed = 2;
//        goal.$save();
//    };
//
});

