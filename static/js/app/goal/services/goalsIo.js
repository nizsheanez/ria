'use strict';
// services in angular.js is lazy, but socket must be opened ASAP
// i know about .run() on service, but it's critical time
(function () {
    var socket = new JsonWebSocket({
        url: {
            ws: 'ws://' + document.domain + ':8080/'
        },
        root: 'js/websocket/',
        pushHandler: function () {
        },
        errorHandler: function (status, error) {
        }
    });
/*
    var socket = new WebSocketConnection2({
        url: {
            ws: 'ws://' + document.domain + ':8080/'
        },
        root: 'js/websocket/'
    });
*/
    angular.module('eg.goal').service('egSocket', ['$q', '$rootScope', 'alertService', function ($q, $rootScope, alertService) {
        return AngularSocketDecorator(socket, $rootScope, alertService);
    }]);

    angular.module('eg.goal').factory('$socketResource', ['$q', '$rootScope', 'egSocket', 'alertService', function ($q, $rootScope, egSocket, alertService) {

        var DEFAULT_ACTIONS = {
            'get': {
                url: 'view'
            },
            'save': {
                url: 'save'
            },
            'query': {
                url: 'api/goals/v1/user',
                isArray: true
            },
            'delete': {
                url: 'delete'
            }
        };

        var noop = angular.noop,
            forEach = angular.forEach,
            extend = angular.extend,
            copy = angular.copy,
            isFunction = angular.isFunction;


        var Resource = function (data) {
            copy(data || {}, this);
        }

        function resourceFactory(url, paramDefaults, actions) {
            actions = extend({}, DEFAULT_ACTIONS, actions);

            var val = new Resource;

            forEach(actions, function (action, name) {

                var value = action.isArray ? [] : new Resource();

                Resource[name] = function (url, params, callback) {
                    var promise = egSocket.send(url, params, function (response) {
                        var data = response.data,
                            promise = value.$promise;

                        if (data) {
                            // Need to convert action.isArray to boolean in case it is undefined
                            // jshint -W018
                            if (angular.isArray(data) !== (!!action.isArray)) {
                                throw $resourceMinErr('badcfg', 'Error in resource configuration. Expected ' +
                                    'response to contain an {0} but got an {1}',
                                    action.isArray ? 'array' : 'object', angular.isArray(data) ? 'array' : 'object');
                            }
                            // jshint +W018
                            if (action.isArray) {
                                value.length = 0;
                                forEach(data, function (item) {
                                    value.push(new Resource(item));
                                });
                            } else {
                                copy(data, value);
                                value.$promise = promise;
                            }
                        }

                        value.$resolved = true;

                        response.resource = value;

                        callback(response)

                        return response;
                    });
                    // we are creating instance / collection
                    // - set the initial promise
                    // - return the instance / collection
                    value.$promise = promise;
                    value.$resolved = false;
                    return value;
                };

                Resource.prototype['$' + name] = function (params, success, error) {
                    if (isFunction(params)) {
                        error = success;
                        success = params;
                        params = {};
                    }

                    var result = Resource[name](action.url, params, success);
                    return result.$promise || result;
                };
            });

            return val;
        }

        return resourceFactory;
    }]);

})();