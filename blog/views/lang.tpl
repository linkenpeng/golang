<!DOCTYPE html>
<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8"> 
</head>

<body>	

	<div>{{.Hi}}<div>  
	<div>{{.Bye}}<div>  
	<hr>
	<div>{{i18n .Lang .Hi}}<div>  
	<div>{{i18n .Lang .Bye}}<div>  
	<div>{{i18n .Lang "about"}}<div>  
	<div>{{i18n .Lang "about.about"}}<div>  
</body>
</html>