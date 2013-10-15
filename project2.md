#Project 2


##Bootstrap

I've used Bootstrap to rapidly design and build the HTML structure of the site. As well, I use a lot of the responsive capabilities that Bootstrap provides in order to provide a baseline mobile friendly experience without too too much development time. 

The table was the biggest challenge but ultimately it's in what I argue to be the most effective style in order for mobile users to parse the data in the tables.

##Angular

Angular provides nearly all the interactivity on the site and has built in functions that enabled the kind of functionality and simplicity I wanted that the site to have for all users. 

Specifically the most important part is the filtering of table rows based on input from the user.

	<div class="col-md-12">
	<form>
	  <div class="form-group">
	    <input ng-model="query" ng-change="update()" id="search" type="search" class="form-control input-lg" placeholder="Search or filter by any information you see below"></div>
	</form>
	<div class="table-responsive">
	  <table class="table table-condensed">
	    <thead>
	      <th ng-click="ordering = 'num'; reverse=!reverse">
	        <a href="">Course #</a>
	      </th>
	      <th ng-click="ordering = 'title'; reverse=!reverse">
	        <a href="">Course Title</a>
	      </th>
	      <th ng-click="ordering = 'professor'; reverse=!reverse">
	        <a href="">Faculty</a>
	      </th>
	      <th class="header-subject" ng-click="ordering = 'subject'; reverse=!reverse">
	        <a href="">Subject</a>
	      </th>
	    </thead>
	    <tbody ng-repeat="class in classes | filter:query | orderBy:ordering:reverse">
	      <tr class="class">
	        <td>
	          <a href="{{class.url}}" target="_blank">{{class.num}}</a>
	        </td>
	        <td>{{class.title}}</td>
	        <td>{{class.professor}}</td>
	        <td class="class-subject">{{class.subject}}</td>
	      </tr>
	    </tbody>
	  </table>
	</div>


##Moment & jQuery

I use the moment.js library to display information about when the user was last on the site. If the user has never been on the site, then it just displays the current day. The last bullt point on the page is located with jQuery and then the text is replaced depending on whether or not it can find the specified 'last' key in the localStorage.

	$(document).ready(function(){

		if(window.localStorage){
		
			if(window.localStorage.getItem('last') !== null){
				var diff = moment(localStorage.getItem('last')).fromNow();
				
				$("#lastTime").html("<strong>The last time we saw you was "+diff+".</strong>");

				console.log('Setting new time');
				localStorage.setItem('last', moment().format());
			} else {

				$('#lastTime').html("We've never see you before but today is "+moment().format('dddd'));
				localStorage.setItem('last', moment().format());
			}
		
		}
	});

!(https://www.dropbox.com/s/xatc8k6itqbrizy/Screenshot%202013-10-15%2019.29.40.png)

##Retinajs 

Although I can't test this on a retina device since I don't own one. But I have included a secondary image of the scotty dog with the approved @2x suffix in the filename so that if it is viewed on a retina screen that it should transition normally. 

##Joyride & jQuery

Also if you have not been to the site, the Joyride sequence will begin by first showing you the title and explaining what Scotty is, and then leading you to the input field and explaining that you can input whatever you'd like to filter the table rows below when searching for a class. Joyride is designed a jQuery plugin and requires that jQuery has also been loaded. 

###Setting up the Joyride 'ride'

	<ol id="tourGuide" style="display:none;">
		<li data-id="title">Hey! Welcome to Scotty! An easier way to parse the Heinz class database.<br/><br/></li>
		<li data-id="search">Searching for anything here and it will filter out all the rows down below.<br/><br/></li>
	</ol>
	<script>
	$(window).load(function() {
		$("#tourGuide").joyride({
			"tipLocation":'bottom'
		});
	});
	</script>



