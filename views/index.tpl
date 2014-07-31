<!DOCTYPE html>
<html lang="en" ng-app="eg.goal">
<head>
    <meta charset="utf-8"/>
    <title>{{ .Title }}</title>
    {{ .Css | str2html }}
</head>
    <script>
        var AUTOBAHN_DEBUG = true;
    </script>
<body>
<div class="main-container ">
    <div class="container">
        <div>
            <alert ng-repeat="alert in alerts" type="alert.type" close="closeAlert($index)"></alert>
        </div>
        <ng-view/>
    </div>
</div>


<script>
    storage = {{ .Storage }}
</script>

{{ if .Debug }}
<script src="//:35729/livereload.js"></script>
{{ end }}

{{ .Js | str2html }}
</body>
</html>
