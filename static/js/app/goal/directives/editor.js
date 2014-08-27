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

angular.module('eg.goal').directive('ngFocus', [function () {
    var FOCUS_CLASS = "ng-focused";
    return {
        restrict: 'A',
        require: 'ngModel',
        link: function (scope, element, attrs, ctrl) {
            ctrl.$focused = false;
            element.bind('focus', function (evt) {
                element.addClass(FOCUS_CLASS);
                scope.$apply(function () {
                    ctrl.$focused = true;
                });
            }).bind('blur', function (evt) {
                element.removeClass(FOCUS_CLASS);
                scope.$apply(function () {
                    ctrl.$focused = false;
                });
            });
        }
    }
}]);

angular.module('eg.goal').directive('egSmartErrors', function ($interval) {
    var CHECK_DIRTY = false;

    return {
        restrict: 'A',
        require: '^form',
        link: function ($scope, el, attrs, formCtrl) {
            var inputEl = el[0].querySelector("[name]");
            var inputNgEl = angular.element(inputEl);

            var errContainerEl = el[0].querySelector(".error-container");
            var errContainerNgEl = angular.element(errContainerEl);
            // get the name on the text box so we know the property to check
            // on the form controller
            var inputName = inputNgEl.attr('name');

            // only apply the has-error class after the user leaves the text box
            var fieldExpr = formCtrl.$name + '["' + inputName+ '"]';
            var validateExpr = fieldExpr + '.$invalid'
            if (CHECK_DIRTY) {
                validateExpr += ' && ' + fieldExpr + '.$dirty';
            }

            var doValidate = function() {
                var isInvalidValid = formCtrl[inputName].$invalid;
                if (CHECK_DIRTY) {
                    isInvalidValid = isInvalidValid && formCtrl[inputName].$dirty;
                }

                el.toggleClass('has-error', isInvalidValid);
                errContainerNgEl.toggleClass('hidden', !isInvalidValid);
                var watchCleaner = $scope.$watch(validateExpr, function(validity) {
                    if (!validity) {
                        el.toggleClass('has-error', false);
                        errContainerNgEl.toggleClass('hidden', true);
                        watchCleaner()
                    }
                });
            };

            el.parent().bind('submit', doValidate);
            inputNgEl.bind('blur', doValidate);
        }
    }
});

angular.module('eg.goal').directive('ngVisible', function () {
    return function (scope, element, attr) {
        scope.$watch(attr.ngVisible, function (visible) {
            element.css('visibility', visible ? 'visible' : 'hidden');
        });
    };
})