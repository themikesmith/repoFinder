'use strict';
angular.module('repoFinder').controller('mainCtrl', function($scope, $http, searchService) {
  $scope.keyword = "";
  $scope.init = function() {
    return $scope.res = {
      gitCount: 0,
      bbCount: 0,
      grCount: 0,
      repo: []
    };
  };
  $scope.search = searchService;
  $scope.resShow = false;
  $scope.sortOpts = [
    {
      opt: 0,
      desc: "Relevance"
    }, {
      opt: 1,
      desc: "Date (Newest to Oldest)"
    }, {
      opt: 2,
      desc: "Date (Oldest to Newest)"
    }
  ];
  $scope.sortOption = $scope.sortOpts[0];
  $scope.changeSort = function(opt) {
    var total_count;
    $scope.sortOption = opt;
    switch (opt.opt) {
      case 0:
        total_count = $scope.res.gitCount + $scope.res.bbCount + $scope.res.grCount;
        return $scope.res.repo.sort(function(a, b) {
          var score_a, score_b;
          score_a = (a.rank / total_count) * a.weight;
          score_b = (b.rank / total_count) * b.weight;
          return score_a - score_b;
        });
      case 1:
        return $scope.res.repo.sort(function(a, b) {
          return b.update - a.update;
        });
      case 2:
        return $scope.res.repo.sort(function(a, b) {
          return a.update - b.update;
        });
    }
  };
  $scope.$watch('res', function(nv, ov) {
    if (nv !== ov) {
      return $scope.changeSort($scope.sortOption);
    }
  }, true);
  return $scope.doSearch = function() {
    $scope.init();
    $scope.resShow = true;
    $scope.search.github($scope.keyword, function(data) {
      var item, rank, _i, _len, _ref, _results;
      $scope.res.gitCount = data.total_count;
      rank = 1;
      _ref = data.items;
      _results = [];
      for (_i = 0, _len = _ref.length; _i < _len; _i++) {
        item = _ref[_i];
        $scope.res.repo.push({
          author: String(item.full_name).split("/")[0],
          reponame: String(item.full_name).split("/")[1],
          rank: rank,
          name: item.full_name,
          url: item.html_url,
          avatar: item.owner.avatar_url,
          description: item.description,
          update: Date.parse(item.updated_at),
          icon: "fa fa-github fa-fw",
          from: "Github",
          weight: 1 / Math.log(data.total_count)
        });
        _results.push(rank++);
      }
      return _results;
    });
    $scope.search.bitbucket($scope.keyword, function(data) {
      var item, rank, _i, _len, _ref, _results;
      $scope.res.bbCount = data.total_count;
      rank = 1;
      _ref = data.items;
      _results = [];
      for (_i = 0, _len = _ref.length; _i < _len; _i++) {
        item = _ref[_i];
        $scope.res.repo.push({
          author: String(item.Title).split("/")[0],
          reponame: String(item.Title).split("/")[1],
          rank: rank,
          name: item.Title,
          url: item.Url,
          avatar: item.Avatar,
          description: item.Description,
          update: Date.parse(item.Date),
          icon: "fa fa-bitbucket",
          from: "Bitbucket",
          weight: 1 / Math.log(data.total_count)
        });
        _results.push(rank++);
      }
      return _results;
    });
    return $scope.search.gitorious($scope.keyword, function(data) {
      var item, rank, _i, _len, _ref, _results;
      $scope.res.grCount = data.total_count;
      rank = 1;
      _ref = data.items;
      _results = [];
      for (_i = 0, _len = _ref.length; _i < _len; _i++) {
        item = _ref[_i];
        $scope.res.repo.push({
          author: String(item.Title).split("/")[0],
          reponame: String(item.Title).split("/")[1],
          rank: rank,
          name: item.Title,
          url: item.Url,
          avatar: "/images/Gitorious.png",
          description: item.Description,
          update: Date.parse(item.Date),
          icon: "",
          from: "Gitorious",
          weight: 1 / Math.log(data.total_count)
        });
        _results.push(rank++);
      }
      return _results;
    });
  };
});
