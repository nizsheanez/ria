'use strict';

angular.module('eg.goal').controller('NavigationCtrl', NavigationCtrl);

NavigationCtrl.$inject = ['$scope', '$location'];

function NavigationCtrl($scope, $location) {
    $scope.activeTab = function (tab) {
        return tab === $location.path().split('/')[1];
    };
}
