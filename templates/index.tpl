{{ define "body"}}
<div class='searchRow'>
	<input type='text' id='cityName' placeholder='Enter a city'>
	<!--<button>Search</button> -->
</div>
<div id='results'>
</div>
<div class='mapArea'>
	<div id='map-canvas'></div>
</div>

{{ end }}

{{ define "scripts"}}
<script type="text/javascript" src="https://maps.googleapis.com/maps/api/js?libraries=places&key=AIzaSyDuBoo4JowSuneYMgiNEEVh72zD_hKZVzs"></script>
<script type="text/javascript" src="/assets/js/main.js"></script>
{{ end }}
