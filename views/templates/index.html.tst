<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <link rel="stylesheet" href="/static/css/index.css">
    <title>MyAlbum</title>
</head>
<body>

{{range $index, $value := .pagers.Before}}
<li><a href="{{$value}}">{{$value}}</a> </li>
{{end}}
<li>{{ .pagers.Current }}</li>
{{range $index, $value := .pagers.After}}
<li><a href="{{$value}}">{{$value}}</a></li>
{{end}}
</body>
</html>
