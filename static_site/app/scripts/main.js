console.log('\'Allo \'Allo!');


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
