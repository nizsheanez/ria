
angular.module('eg.goal').directive('ngFocus', ngFocus);

function ngFocus() {
    var FOCUS_CLASS = "ng-focused";
    return {
        restrict: 'A',
        require: 'ngModel',
        link: link
    };

    /////////////////

    function link(scope, element, attrs, ctrl) {
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

angular.module('eg.goal').directive('egSmartErrors', egSmartErrors);

egSmartErrors.$inject = ['$interval']

function egSmartErrors($interval) {
    var CHECK_DIRTY = false;

    return {
        restrict: 'A',
        require: '^form',
        link: link
    };

    ////////////////

    function link($scope, el, attrs, formCtrl) {
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

angular.module('eg.goal').directive('ngVisible', ngVisible);

function ngVisible() {
    return function (scope, element, attr) {
        scope.$watch(attr.ngVisible, function (visible) {
            element.css('visibility', visible ? 'visible' : 'hidden');
        });
    };
}