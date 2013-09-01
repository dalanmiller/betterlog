function BetterlogueCtrl($scope) {
  
  //$scope.classList = $.get('class_list.json', function(doc){return JSON.parse(doc)});

  $scope.filter = function() {
    //$scope.todos.push({text:$scope.todoText, done:false});
    //$scope.todoText = '';
  };

  $scope.update = function() {
    console.log("wtf")
    console.log($scope.query)
  }

}

