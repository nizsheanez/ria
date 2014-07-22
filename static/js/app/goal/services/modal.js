'use strict';

angular.module('eg.goal').factory('Modal', ['$modal', 'Tpl', 'Goal', function ($modal, Tpl, Goal) {
    var service = {
        editGoalModal: function (goal) {
            $.animationCount++;
            var promise = $modal.open({
                templateUrl: Tpl.modal.edit,
                controller: 'GoalEditModalCtrl',
                resolve: {
                    params: function () {
                        return {
                            goal: {
                                id: goal.id,
                                title: goal.title,
                                fk_goal_category: goal.fk_goal_category
                            },
                            html: {
                                modalClass: 'goal-edit-modal'
                            }
                        }
                    }
                }
            });

            promise.opened.then(function() {
                $.animationCount--;
            });
            promise.result.then(function (newGoal) {
                goal.title = newGoal.title;
                goal.fk_goal_category = newGoal.fk_goal_category;
                goal.$save();
            }, function () {
                //just dismiss
            });
            return promise;
        },
        addGoalModal: function (category) {
            $.animationCount++;
            var promise = $modal.open({
                    templateUrl: Tpl.modal.edit,
                    controller: 'GoalEditModalCtrl',
                    resolve: {
                        'params': function () {
                            return {
                                goal: {
                                    fk_goal_category: category.id
                                },
                                html: {
                                    modalClass: 'goal-add-modal'
                                }
                            }
                        }
                    }
                }
            );
            promise.opened.then(function() {
                $.animationCount--;
                console.log($.active === 0 && $.animationCount === 0 && $.socketCallCount === 0);
            });
            promise.result.then(function (newGoal) {
                    Goal.add(newGoal);
                }, function () {
                    //just dismiss
                });
        },
        backLogModal: function (category) {
            $.animationCount++;

            var promise = $modal.open({
                    templateUrl: Tpl.modal.backLog,
                    controller: 'BacklogModalCtrl',
                    resolve: {
                        'params': function () {
                            return {
                                category: category,
                                html: {
                                    modalClass: 'goal-back-log-modal'
                                }
                            }
                        }
                    }
                }
            );
            promise.opened.then(function() {
                $.animationCount--;
            });
        }
    };
    return service;
}])
;