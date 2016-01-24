package main

const BaseTmpl = `
{{ define "base" }}
<html class="" lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Scattr</title>
    <link rel="stylesheet" href="/assets/css/foundation.min.css">
    <link href="/assets/css/foundation-icons.css" rel="stylesheet" type="text/css">
    <link href="/assets/css/font-awesome.min.css" rel="stylesheet" type="text/css">
    <link href="/assets/css/codemirror.css" rel="stylesheet" type="text/css">
    <link href="/assets/css/scattr.css" rel="stylesheet" type="text/css">

    <meta class="foundation-mq">
</head>
<body>

    <div class="top-bar">
        <div class="row">
            <div class="top-bar-left">
                <ul class="dropdown menu" data-dropdown-menu="" role="menubar" data-dropdownmenu="1bvtxt-dropdownmenu">
                    <li class="menu-text" role="menuitem" tabindex="0" id="logo">Scattr</li>
                </ul>
            </div>
        </div>
    </div>

    <br>

    <div class="row">
      {{ template "content" . }}
    </div>
    <div class="row column">
        <hr>
        <ul class="menu">
            <li>Contribute to Scattr</li>
            <li><a href="http://github.com" id="ghlogo"><i class="fa fa-github"></i></a></li>
        </ul>
    </div>
    <script src="/assets/js/jquery-2.1.4.min.js"></script>
    <script src="/assets/js/foundation.js"></script>
    <script src="/assets/js/codemirror.js"></script>
    <script src="/assets/js/mode.js"></script>
    <script src="/assets/js/scattr.js"></script>
    <script>
        $(document).foundation();
    </script>
</body>
</html>
{{ end }}
`
const ConfigTemp = `
{{ define "title"}}<title>Configurations</title>{{ end }}
{{ define "content" }}
<div class="row">
  <div class="medium-7 large-6 columns">
    <div class="callout secondary">
      <h5>Add URL</h5>
      <form method="post" action="/add/">
        <input type="url" name="url" placeholder="http://" required>
        <input type="submit" value="Add" class="button">
      </form>
    </div>
    <hr/>
    <h5>URLs <small>List of urls where the requests will scatter</small></h5>
    <div >
          {{range $i, $elem := .OutUrls}}

            <div class="callout">
              <p>{{$elem}}</p>
              <a class="close-button" aria-label="Dismiss alert" href="/delete/{{$i}}">
                <span aria-hidden="false"><i class="fa fa-trash"></i></span>
              </a>
            </div>

          {{end}}
    </div>


  </div>

  <div class="medium-7 large-6 columns">
    <h5>SCRIPT <small>Use this script to control scattering</small></h5>
    <div class="callout">
      <form method="post" action="/savescript/">
      	<textarea  id="code" rows="20" cols="35" name="script"></textarea></br>
        <input type="submit" value="Save" class="small primary button">
      </form>
    </div>
  </div>

</div>
{{ end }}
`
