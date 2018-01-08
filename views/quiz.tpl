<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <meta name="description" content="">
    <meta name="author" content="">
    <link rel="icon" href="/static/favicon.ico">

    <title>Total Recall - Choose Your Destiny!</title>

    <!-- Bootstrap core CSS -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <!-- Custom styles for this template -->
    <link href="/static/css/main.css" rel="stylesheet">

    <script src="/static/js/jquery-3.2.1.min.js"></script>
    <script src="/static/js/main.js"></script>

</head>

<body>
<div class="container">
    <div class="page-header">
        <h1>Choose Your Destiny!</h1>
        <p class="lead">Flawless Victory!</p>
    </div>

    <h3>Your Answer</h3>

    <div id="quizForm">
        <div class="row">
            <div class="panel panel-success">
                <div class="panel-heading"><h3 class="panel-title">YES</h3></div>
                <div class="panel-body">
                    <form method="post" name="yes">
                        <button type="submit" id="yesBtn" name="btn" value="yes" class="btn btn-success btn-lg">YES</button>
                    </form>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="panel panel-danger">
                <div class="panel-heading"><h3 class="panel-title">NO</h3></div>
                <div class="panel-body">
                    <form method="post" name="no">
                        <div class="col-xs-11">
                            <div class="checkbox">
                                <label class="pull-right" >
                                    <input type="checkbox" id="enableNoBtnCheckbox" value="enable" aria-label="Activate NO button">Activate NO button
                                </label>
                            </div>
                        </div>
                        <div class="col-xs-1">
                            <button type="submit" id="noBtn" name="btn" value="no" class="btn btn-danger pull-right" disabled="disabled">NO</button>
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>
</body>
</html>
