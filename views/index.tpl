<!DOCTYPE html>
<html lang="en" ng-app="eg.goal">
<head>
    <meta charset="utf-8"/>
    <title>{{ .Title }}</title>
    {{ .Css | str2html }}
</head>
<body>

<script>
    storage = {{ .Storage | str2html }}
</script>

{{ if .Debug }}
<script src="//:35729/livereload.js"></script>
{{ end }}

{{ .Js | str2html }}
<script>
    $(document).ready(function() {
        $.get('/site/auto', {}, function(auto) {
            $.post('/site/html', {a:auto}, function(a,b,c) {
                console.log(a,b,c);
            },'json');
        })
    });
</script>

</body>
</html>
