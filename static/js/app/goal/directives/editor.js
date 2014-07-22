angular.module('eg.goal').directive('egEditor', function ($debounce) {

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
        link: function ($scope, element, attrs) {
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
    };
});

angular.module('eg.goal').directive('egPanel', function ($debounce) {
    return {
        restrict: 'E',
        replace: false,
        transclude: true,
        scope: {
        },
        link: function ($scope, element, attrs) {
            var div = $('<div>').append(element.children());
            element.append(div);
            element.addClass('eg-panel');
            div.addClass('panel panel-default');
            div.find('header').addClass('panel-heading');
        }
    };
});

angular.module('eg.goal').directive('egTodo', function () {
    return {
        restrict: 'A',
        scope: false,
        link: function ($scope, element, attrs) {
            element.change(function () {
                element.attr('checked', element.is(':checked') ? 'checked' : false);
            });
        }
    };
});

