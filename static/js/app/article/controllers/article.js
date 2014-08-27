'use strict';

angular.module('eg.goal').controller('ArticleCreateCtrl', function ($scope, $location, Tpl, Article) {

    $scope.tpl = Tpl;

    $scope.article = Article.New();

    $scope.submit = function() {
        if ($scope.articleForm.$valid) {
            $scope.article.$save()
        } else {
            return false;
        }
    }

});

