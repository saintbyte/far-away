package templates

const (
	IndexTemplate = `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>{{ title }} – Far away</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no" />
    <meta name="format-detection" content="telephone=no" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="MobileOptimized" content="176" />
    <meta name="HandheldFriendly" content="True" />
    <meta name="robots" content="index, follow" />
    <meta property="og:type" content="article">
    <meta property="og:title" content="{{ title }}">
    <meta property="og:description" content="{{ description }}">
    <meta property="og:image" content="/i/far-away.jpg">
    <meta property="og:site_name" content="Far away">
    <meta property="article:author" content="">
    <meta name="twitter:card" content="summary">
    <meta name="twitter:title" content="{{ title }}">
    <meta name="twitter:description" content="{{ description }}">
    <meta name="twitter:image" content="/i/far-away.jpg">
    {% if slug %}<link rel="canonical" href="[domain][slug]" />{% endif %}

    <link rel="icon" type="image/png" href="/favicon/favicon-96x96.png" sizes="96x96" />
    <link rel="icon" type="image/svg+xml" href="/favicon/favicon.svg" />
    <link rel="shortcut icon" href="/favicon/favicon.ico" />
    <link rel="apple-touch-icon" sizes="180x180" href="/favicon/apple-touch-icon.png" />
    <meta name="apple-mobile-web-app-title" content="Far away" />
    <link rel="manifest" href="/favicon/site.webmanifest" />
    <script type="module" crossorigin src="/static/js/index.js"></script>
    <link rel="stylesheet" crossorigin href="/static/assets/index.css">
</head>
<body>
<div id="app"></div>
</body>
</html>
`
	PageTemplate = `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>{{ title }} – Far away</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no" />
    <meta name="format-detection" content="telephone=no" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="MobileOptimized" content="176" />
    <meta name="HandheldFriendly" content="True" />
    <meta name="robots" content="index, follow" />
    <meta property="og:type" content="article">
    <meta property="og:title" content="{{ title }}">
    <meta property="og:description" content="{{ description }}">
    <meta property="og:image" content="/i/far-away.jpg">
    <meta property="og:site_name" content="Far away">
    <meta property="article:author" content="">
    <meta name="twitter:card" content="summary">
    <meta name="twitter:title" content="{{ title }}">
    <meta name="twitter:description" content="{{ description }}">
    <meta name="twitter:image" content="/i/far-away.jpg">
    {% if slug %}<link rel="canonical" href="[domain][slug]" />{% endif %}

    <link rel="icon" type="image/png" href="/favicon/favicon-96x96.png" sizes="96x96" />
    <link rel="icon" type="image/svg+xml" href="/favicon/favicon.svg" />
    <link rel="shortcut icon" href="/favicon/favicon.ico" />
    <link rel="apple-touch-icon" sizes="180x180" href="/favicon/apple-touch-icon.png" />
    <meta name="apple-mobile-web-app-title" content="Far away" />
    <link rel="manifest" href="/favicon/site.webmanifest" />
    <link rel="stylesheet" crossorigin href="/static/assets/index.css">
</head>
<body>
<div id="app">
<div class="article-header">
        <div class="article-header-title">
            {{ title }}
        </div>
        <div class="article-header-author">
            <small>{{ author }}</small>
        </div>
</div> 
<div id="editor">{% autoescape off %}{{ text }}{% endautoescape %}</div>
</div>
</body>
</html>
`
)
