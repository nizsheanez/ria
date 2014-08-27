'use strict';

angular.module('eg.goal').controller('NewsCtrl', function ($q, $http, $scope, $routeParams, $location, $modal, Tpl, User) {

    $scope.tpl = Tpl;
    $scope.users = [];

    User.getAll(function (result) {
        $scope.users = result
    });
});


angular.module('eg.goal').controller('ArticleCreateCtrl', function ($scope, $location, Tpl, User) {

    $scope.tpl = Tpl;

    $scope.article = User.getInstance();
    $scope.article.$get({id:1}, function(r) {
        $scope.article = r
    });

    $scope.submit = function() {
        if ($scope.articleForm.$valid) {
            // Submit as normal
        } else {
            return false;
        }
    }

});

