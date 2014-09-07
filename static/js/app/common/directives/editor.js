angular.module('eg.goal').directive('egEditor', egEditor);

egEditor.$inject = ['$debounce'];

function egEditor($debounce) {

    return {
        restrict: 'E',
        require: '^ngModel',
        scope: {
            ngModel: '=',
            ngChange: '&',
            ngClass: '@',
            ngFocus: '&',
            submodel: '=',
            placeholder: '&',
            fg: '&'
        },
        template: '<div text-angular ng-model="ngModel" ng-change="onChange()"></div>',
        link: link
    };

    /////////////
    function link($scope, element, attrs) {
        var editor = element.children('div');

        if (!$scope.ngModel) {
            $scope.ngModel = '<div>&nbsp;</div>';
        }

        //css
        editor
            .addClass('eg-editor');

        element.closest('eg-panel').find('header .editor-controls').append(element.find('.btn-toolbar').addClass('eg-editor-toolbar'));

        $scope.onChange = $debounce($scope.ngChange, 1000);
    }
}

angular.module('eg.goal').directive('egPanel', egPanel)

function egPanel() {
    return {
        restrict: 'E',
        replace: false,
        transclude: true,
        scope: {
        },
        link: link
    };

    //////////////

    function link($scope, element, attrs) {
        var div = $('<div>').append(element.children());
        element.append(div);
        element.addClass('eg-panel');
        div.addClass('panel panel-default');
        div.find('header').addClass('panel-heading');
    }
}

angular.module('eg.goal').directive('egTodo', egTodo);

function egTodo() {
    return {
        restrict: 'A',
        scope: false,
        link: link
    };

    ////////////////

    function link($scope, element, attrs) {
        element.change(function () {
            element.attr('checked', element.is(':checked') ? 'checked' : false);
        });
    }
}
