{{- /* gotype: github.com/mrbrist/poebin/internal/r2.Build */ -}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <link rel="icon" type="image/x-icon" href="/assets/favicon.ico">
    <title>Level {{ .Data.Build.Level }} {{ .Data.Build.AscendClassName }}</title>
</head>
<body>
    {{ .Data.Items.ItemSets }}

</body>
</html>
