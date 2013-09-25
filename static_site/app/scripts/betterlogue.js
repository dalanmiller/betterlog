function BetterlogueCtrl($scope, $http) {
  
  	$scope.test = ["1","2","3"];
  	
  	// $http({method: 'GET', url: '/someUrl'}).
	  // success(function(data, status, headers, config) {
	  //   // this callback will be called asynchronously
	  //   // when the response is available
	  // }).
	  // error(function(data, status, headers, config) {
	  //   // called asynchronously if an error occurs
	  //   // or server returns response with an error status.
  	// });

 	console.log("Attempting to get JSON and parse the classes file");
 	$http({method: 'GET', url: 'scripts/classes.json'}).
 		success(function(d){
			d.forEach(function(t){

				t.keywords = [t.title, t.professor, t.subject, t.num].join(" ");
				t.visible = "none";
			});

			$scope.classes = d;

  		});

  
  $scope.debugClasses = function(){
  	console.log(classes);
  }

}

