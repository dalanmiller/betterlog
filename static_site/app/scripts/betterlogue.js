'use strict';

var Scotty = angular.module('Scotty',[]);

Scotty.controller( "LogCtrl", ["$scope", "$http", function($scope, $http) {
    
    $scope.ordering = "num";

 	console.log("Attempting to get JSON and parse the classes file");
 	$http({method: 'GET', url: 'classes.json'}).
 		success(function(d){
			d.forEach(function(t){
				t.keywords = [t.title, t.professor, t.subject, t.num].join(" ");
			});
			$scope.classes = d;
  		});
}]);

Scotty.controller("ClassCtrl", ["$scope", function($scope){

	console.log($scope.ClassCtrl(ev));

}]);
