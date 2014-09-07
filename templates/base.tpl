{{ define "base" }}
<!DOCTYPE html>
<html>
<head>
	<title>{{ .Title }} - WorldWeather</title>
	<link rel="stylesheet" type="text/css" href="/assets/css/main.css">
</head>
<body>
	<div class='container'>
		<div class='header'>
			<div class='siteTitle'>
				World Weather
			</div>
			<nav>
				<a href='/'>
					Home
				</a>
				<span clas='dot'>•</span>
				<a href='/about'>
					About
				</a>
				<span clas='dot'>•</span>
				<a href='http://golang.org'>
					Go Lang
				</a>
			</nav>
		</div>
		<div class='content'>
			{{ template "body" .}}

		</div>
		<div class='footer'>
			&copy; 2014 Akshay Shekher
		</div>
	</div>
	{{ template "scripts" .}}
</body>
</html>
{{ end }}
