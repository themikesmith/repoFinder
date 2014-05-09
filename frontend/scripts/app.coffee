'use strict'

angular.module('repoFinder', [])
    .config ($routeProvider, $httpProvider) ->
        $routeProvider
            .when '/',
                templateUrl: 'views/main.html'
            .otherwise
                redirectTo: '/'
    .controller 'mainCtrl', ( $scope, $http )->
        $scope.resShow = false
        $scope.searchRes = []
        $scope.url = "https://api.github.com/search/repositories"
        $scope.keyword = "Hello Git!"
        $scope.doSearch = ->
            $http.get( $scope.url + "?q=" + $scope.keyword )
            .success ( data )->
                $scope.resShow = true
                $scope.count = data.total_count
                $scope.searchRes = []
                for item in data.items
                    $scope.searchRes.push
                        name:   item.full_name
                        url:    item.html_url
                        avatar: item.owner.avatar_url

