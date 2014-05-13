'use strict';
angular.module('repoFinder', ['ngRoute']).config(function($routeProvider, $httpProvider, $locationProvider) {
  $routeProvider.when('/', {
    templateUrl: 'views/main.html'
  }).otherwise({
    redirectTo: '/'
  });
  return $locationProvider.html5Mode(true);
});
