'use strict';
angular.module('repoFinder').factory('searchService', function($http) {
  return {
    github: function(keyword, callback) {
      var kw, url;
      kw = String(keyword).replace(/\s/g, "+");
      if (kw.length === 0) {
        kw = "Hello+World";
      }
      url = "https://api.github.com/search/repositories";
      return $http.get(url + "?q=" + kw + "&per_page=100").success(callback);
    },
    bitbucket: function(keyword, callback) {
      var kw, url;
      kw = String(keyword).replace(/\s/g, "+");
      if (kw.length === 0) {
        kw = "Hello+World";
      }
      url = "http://hi-ougi.com:8080/BbSearch/";
      return $http.get(url + kw).success(callback);
    },
    gitorious: function(keyword, callback) {
      var kw, url;
      kw = String(keyword).replace(/\s/g, "+");
      if (kw.length === 0) {
        kw = "Hello+World";
      }
      url = "http://hi-ougi.com:8080/GrSearch/";
      return $http.get(url + kw).success(callback);
    }
  };
});
