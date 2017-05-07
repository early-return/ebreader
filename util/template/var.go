package template

var (
	htmlTemplate = `
<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <title>EBReader - {{.Title}}</title>
    <style type="text/css">
        body {
            padding: 0;
            margin: 0;
        }
        
        body body {
            margin-left: 100px;
        }
        
        nav {
            position: fixed;
            top: 0;
            left: 0;
            bottom: 0;
            width: 300px;
            overflow: auto;
            font-family: sans-serif;
        }
        
        nav .title {
            margin: 0;
            margin-bottom: 20px;
            padding: 15px;
            border-bottom: 1px solid #AAA;
        }
        
        nav ul {
            list-style: none;
            margin-left: 20px;
            padding: 0;
        }
        
        nav a {
            text-decoration: none;
            color: #333;
            display: block;
            padding: 8px 15px;
        }
        
        nav a:hover {
            text-decoration: underline;
        }
        
        .nav-level1 {
            margin: 0;
        }
        
        #frame {
            position: fixed;
            left: 320px;
            right: 0;
            top: 0;
            bottom: 0;
        }
        
        #frame iframe {
            border: 0;
            width: 100%;
            height: 100%;
        }
    </style>
</head>

<body>
    <nav>
        <p class="title">{{.Title}}</p>
        <ul class="nav-level1">
			{{range .Navs}}
            	<li><a target="content" href="{{.Src.URL}}">{{.Title}}</a></li>
				{{if .SubNavs}}
					<ul class="nav-level2">
						{{range .SubNavs}}
        					<li><a target="content" href="{{.Src.URL}}">{{.Title}}</a></li>
							{{if .SubNavs}}
								<ul class="nav-level3">
									{{range .SubNavs}}
										<li><a target="content" href="{{.Src.URL}}">{{.Title}}</a></li>
									{{end}}
								</ul>
							{{end}}
						{{end}}
					</ul>
				{{end}}
			{{end}}
		</ul>
    </nav>
    <div id="frame">
        <iframe name="content" src="{{.Navs|getFirstSrc}}" seamless>
        </iframe>
    </div>
</body>

</html>
	`
)
