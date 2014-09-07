'use strict';

angular.module('eg.goal').controller('GoalCtrl', GoalCtrl);

GoalCtrl.$inject = ['$scope', '$location', 'Tpl', 'User', 'Category', 'Goal', 'Conclusion', 'Modal'];

function GoalCtrl($scope, $location, Tpl, User, Category, Goal, Conclusion, Modal) {

    $scope.tpl = Tpl;
    $scope.keys = [];
    $scope.goals = [];
    $scope.categories = [];
    $scope.conclusions = [];


    User.get(function () {
        $scope.categories = Category.getAll();
        $scope.conclusions = Conclusion.getAll();
        $scope.goals = Goal.getAll();
    });

    $scope.focusGoal = false;
    $scope.setFocus = function (goal) {
        $scope.focusGoal = goal;
    };

    $scope.inFocus = function (goal) {
        return goal.id === $scope.focusGoal.id;
    };

    $scope.defaultPlaceholder = 'Сделано';

    $scope.save = function (model) {
        model.$save();
    };

    if ($location.path() === '/') {
        $scope.day = 'today';
    } else if ($location.path() === '/yesterday') {
        $scope.day = 'yesterday';
    }

    $scope.location = $location;

    $scope.editGoalModal = Modal.editGoalModal;
    $scope.addGoalModal = Modal.addGoalModal;
    $scope.backLogModal = Modal.backLogModal;

    $scope.complete = function (goal) {
        goal.completed = 1;
        goal.$save();
    };

    $scope.fail = function (goal) {
        goal.completed = 2;
        goal.$save();
    };


    showScreen();
}

angular.module('eg.goal').controller('GoalEditModalCtrl', GoalEditModalCtrl);

GoalEditModalCtrl.$inject = ['$scope', '$modalInstance', 'params', 'Category'];

function GoalEditModalCtrl($scope, $modalInstance, params, Category) {

    $scope.goal = params.goal;
    $scope.categories = Category.getAll();

    $scope.html = params.html;

    $scope.ok = function () {
        $modalInstance.close(params.goal);
    };

    $scope.cancel = function () {
        $modalInstance.dismiss('cancel');
    };
}


angular.module('eg.goal').controller('BacklogModalCtrl', BacklogModalCtrl);

BacklogModalCtrl.$inject = ['$scope', '$modalInstance', 'params', 'Goal', 'Modal'];

function BacklogModalCtrl($scope, $modalInstance, params, Goal, Modal) {

    $scope.goals = Goal.getAll();
    $scope.category = params.category;
    $scope.html = params.html;

    $scope.editGoalModal = Modal.editGoalModal;

    $scope.ok = function () {
        $modalInstance.close(params.goal);
    };

    $scope.cancel = function () {
        $modalInstance.dismiss('cancel');
    };
}
