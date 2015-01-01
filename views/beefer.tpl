<html>
<head>
    <title>Beefer!</title>
    <link rel="stylesheet" href="/static/css/bootstrap.css">
</head>
<body>
<div class="navbar navbar-default">
    <div class="navbar-header">
        <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-responsive-collapse">
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
        </button>
        <a class="navbar-brand" href="#">Beefer</a>
    </div>
    <div class="navbar-collapse collapse navbar-responsive-collapse">
        <ul class="nav navbar-nav">
            <li class="active"><a href="#">Home</a></li>
        </ul>
        <div class="navbar-right">
            <ul class="nav navbar-nav">
                {{if .User}}
                <li><p class="navbar-text">Hello, {{.User.Username}}</p></li>
                <li><a href="/user/logout">LOG OUT</a></li>
                {{else}}
                <li><a href="/user/login">LOG IN</a></li>
                <li><a href="/user/signup">SIGN UP</a></li>
                {{end}}
            </ul>
        </div>
    </div>
</div>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/2.1.3/jquery.min.js"></script>
<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.1/js/bootstrap.min.js"></script>
</body>
</html>