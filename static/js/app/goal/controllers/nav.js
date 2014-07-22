'use strict';

angular.module('eg.goal').controller('NavigationCtrl', ['$scope', '$location', function ($scope, $location) {
    $scope.activeTab = function (tab) {
        return tab === $location.path().split('/')[1];
    };
}]);
