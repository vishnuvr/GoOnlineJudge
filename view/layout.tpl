<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>{{.Title}}</title>
    <link href="/static/css/style.css" rel="stylesheet" type="text/css" />
    <!--[if lt IE 8]><link href="/static/css/style-lt8.css" media="screen" rel="stylesheet" type="text/css" /><![endif]-->
    <!--[if lt IE 7]><link href="/static/css/style-lt7.css" media="screen" rel="stylesheet" type="text/css" /><![endif]-->
    <script src="/static/js/jquery.min.js" type="text/javascript"></script>
    <script src="/static/js/action.js" type="text/javascript"></script>
    <!-- link href='http://fonts.googleapis.com/css?family=Droid+Sans:400,700' rel='stylesheet' type='text/css' -->
    {{if .IsCode}}
      <link href="/static/prettify/prettify.css" rel="stylesheet" type="text/css" />
      <script src="/static/prettify/prettify.js" type="text/javascript"></script>
    {{end}}
  </head>
  <body {{if .IsCode}}onload="prettyPrint()"{{end}}>
    <div class="container">
      <div id="pageHeader">
        <div id="logo" class="lfloat">
          <a href="/"><img alt="Logo" src="/static/img/logo.png" /></a>
        </div>
        <div id="headerInfo" class="rfloat">
          {{if .IsCurrentUser}}
            {{if .IsShowAdmin}}<a href="/admin/">[Admin]</a>{{end}}
            <a href="/user/detail/uid/{{.CurrentUser}}">[{{.CurrentUser}}]</a>
            <a class="user_signout" href="#">[Sign Out]</a>
          {{else}}
            {{if .IsUserSignIn}}[Sign In]{{else}}<a href="/user/signin">[Sign In]</a>{{end}}
            {{if .IsUserSignUp}}[Sign Up]{{else}}<a href="/user/signup">[Sign Up]</a>{{end}}
          {{end}}
        </div>
        <hr> 
        </div>
        <div id="navibar" class="span-3">
        <ul>
          <li>{{if .IsNews}}<span>Home</span>{{else}}<a href="/">Home</a>{{end}}</li>
          <li>{{if .IsProblem}}<span>Problem</span>{{else}}<a href="/problem/list">Problem</a>{{end}}</li>
          <li>{{if .IsStatus}}<span>Status</span>{{else}}<a href="/status/list">Status</a>{{end}}</li>
          <li>{{if .IsRanklist}}<span>Ranklist</span>{{else}}<a href="/ranklist">Ranklist</a>{{end}}</li>
          <li>{{if .IsContest}}<span>Contest</span>{{else}}<a href="/contest/list">Contest</a>{{end}}</li>
          {{if .IsContestDetail}}
            <div id="psnavi">
              <ul>
                <li>{{if .IsContestProblem}}<span>Problem</sapn>{{else}}<a href="/contest/problem/list/cid/{{.Cid}}">Problem</a>{{end}}</li>
                <li>{{if .IsContestStatus}}<span>Status</sapn>{{else}}<a href="/contest/status/list/cid/{{.Cid}}">Status</a>{{end}}</li>
                <li>{{if .IsContestRanklist}}<span>Ranklist</sapn>{{else}}<a href="/contest/ranklist/cid/{{.Cid}}">Ranklist</a>{{end}}</li>
              </ul>
            </div>
          {{end}}
          {{if .IsCurrentUser}}
            <li><a href="/user/settings">Settings</a></li>
            {{if .IsSettings}}
            <div id="psnavi">
              <ul>
                <li>{{if .IsSettingsDetail}}<span>Detail</span>{{else}}<a href="/user/detail/uid/{{.CurrentUser}}">Detail</a>{{end}}</li>
                <li>{{if .IsSettingsEdit}}<span>Edit Info</span>{{else}}<a href="/user/edit">Edit Info</a>{{end}}</li>
                <li>{{if .IsSettingsPassword}}<span>Edit Password</span>{{else}}<a href="/user/pagepassword">Edit Password</a>{{end}}</li>
              </ul>
            </div>
            {{end}}
          {{end}}
        </ul>
      </div>
      <div id="body" class="span-22 last">
        {{template "content" .}}
      </div>
      <div id="pageFooter" class="center">
        <hr class="nomarginbottom">
        <div id="footerContainer">
          <div class="center">Copyright (C) 2013-2014 Zhejiang Gongshang University ACM Club</div>
        </div>
      </div>
    </div>
    <script type="text/javascript">
    $('.user_signout').on('click', function(e) {
      e.preventDefault();
      $.ajax({
        type:'POST',
        url:'/user/logout',
        data:$(this).serialize(),
        error: function() {
          alert('Sign Out Failed.');
        },
        success: function() {
          window.location.href = '/user/signin';
        }
      });
    });
    </script>
  </body>
</html>

