var map, autocomplete, places;

function initialize() {
	setupMapsAndPlaces();
}

function ajaxPOST(url, data, cb) {
	var r = new XMLHttpRequest();
	r.open("POST", url, true);
	r.setRequestHeader("Content-type","application/x-www-form-urlencoded");
	r.onreadystatechange = function () {
		if (r.readyState != 4 || r.status != 200) return;
		cb(JSON.parse(r.response));
	};
	r.send(data);
}


function fetchWeather(location, cb) {
	ajaxPOST("/weather", "city="+escape(location.name), cb);
}

/**
 * result = {
 *    cityName: "City" ,
 *    values: {
 *        Temperature: 20,
 *        Humidity: 80
 *    }
 * }
 */
 function addResultBox(result, id) {
 	var el = document.createElement('div');
 	el.id = id;
 	el.className = 'result'

 	var heading = document.createElement('h3');
 	heading.innerText = result.CityName;
 	el.appendChild(heading);


 	var result_content = document.createElement('div');
 	result_content.className = 'result-contents';
 	el.appendChild(result_content);

 	var result_heading = document.createElement('div');
 	result_heading.className = 'result-names';
 	result_content.appendChild(result_heading);

 	var result_values = document.createElement('div');
 	result_values.className = 'result-values';
 	result_content.appendChild(result_values);

 	for(var item in result.Values) {
 		if (result.Values[item] == "<null>") {
 			continue;
 		}
 		var head =  document.createElement('div');
 		head.innerText = item;
 		result_heading.appendChild(head);

 		var val =  document.createElement('div');
 		val.innerText = result.Values[item];
 		result_values.appendChild(val);
 	}

 	document.getElementById('results').appendChild(el);
 }
/**
 * setup google maps and places API
 */
 function setupMapsAndPlaces() {
 	var mapOptions = {
 		center: new google.maps.LatLng(28.61, 77.23),
 		zoom: 1,
 		mapTypeId: google.maps.MapTypeId.HYBRID
 	};
 	map = new google.maps.Map(document.getElementById("map-canvas"),
 		mapOptions);
 	autocomplete = new google.maps.places.Autocomplete((document.getElementById('cityName')),
 		{ types: ['(cities)'] });

 	places = new google.maps.places.PlacesService(map);

 	google.maps.event.addListener(autocomplete, 'place_changed', onPlaceChanged);

 }

/**
 * google places API, place changed callback
 */
function onPlaceChanged() {
 	var place = autocomplete.getPlace();
 	if (place.geometry) {
		// send location to the API
		//sendLocation(place);
		fetchWeather(place, function(data) {
			var uid = uuid();
			addResultBox(data, uid);
			var marker = new google.maps.Marker({
				position: place.geometry.location,
				map: map,
				title: data.CityName,
				uid: uid
			});

			google.maps.event.addListener(marker, 'click', function() {
				//map.setZoom(8);
				//map.setCenter(marker.getPosition());
				addHoverClass(uid);
			});
		});
	} else {
		document.getElementById('cityName').placeholder = 'Enter a city';
	}
}

function addHoverClass(id) {
	var el = document.getElementById(id);
	el.classList.add('hover');
	setTimeout(function() {
		el.classList.remove('hover');
	}, 1500);
}
function uuid() {
	function s4() {
		return Math.floor((1 + Math.random()) * 0x10000)
		.toString(16)
		.substring(1);
	}
	return function() {
		return s4() + s4() + '-' + s4() + '-' + s4() + '-' +
		s4() + '-' + s4() + s4() + s4();
	}();
}
/**
 * calls the function initialize when the dom is ready
 */
google.maps.event.addDomListener(window, 'load', initialize);
