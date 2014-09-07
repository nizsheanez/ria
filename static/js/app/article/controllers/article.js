'use strict';

angular.module('eg.goal').controller('ArticleCreateCtrl', ArticleCreateCtrl);

ArticleCreateCtrl.$inject = ['Article'];

function ArticleCreateCtrl(Article) {

    /* jshint validthis: true */
    var vm = this;
    vm.article = Article.New();
    vm.submit = formSubmit;

    //////////

    function formSubmit() {
        if (vm.$scope.articleForm.$valid) {
            vm.article.$save()
        } else {
            return false;
        }
    }
}


