{{define "content"}}
<h1>Status List</h1>
<form accept-charset="UTF-8" id="search_form">
User: <input id="search_uid" name="search_uid" size="20" type="text" value="{{.SearchUid}}">
Problem: <input id="search_pid" name="search_pid" size="10" type="text" value="{{.SearchPid}}">
Result: <select id="search_judge" name="search_judge">
  <option value="0" {{if .SearchJudge0}}selected{{end}}>All</option>
  <option value="1" {{if .SearchJudge1}}selected{{end}}>Pending</option>
  <option value="2" {{if .SearchJudge2}}selected{{end}}>Running & Judging</option>
  <option value="3" {{if .SearchJudge3}}selected{{end}}>Accept</option>
  <option value="4" {{if .SearchJudge4}}selected{{end}}>Compile Error</option>
  <option value="5" {{if .SearchJudge5}}selected{{end}}>Runtime Error</option>
  <option value="6" {{if .SearchJudge6}}selected{{end}}>Wrong Answer</option>
  <option value="7" {{if .SearchJudge7}}selected{{end}}>Time Limit Exceeded</option>
  <option value="8" {{if .SearchJudge8}}selected{{end}}>Memory Limit Exceeded</option>
  <option value="9" {{if .SearchJudge9}}selected{{end}}>Output Limit Exceeded</option>
</select>
Language: <select id="search_language" name="search_language">
  <option value="0" {{if .SearchLanguage0}}selected{{end}}>All</option>
  <option value="1" {{if .SearchLanguage1}}selected{{end}}>C</option>
  <option value="2" {{if .SearchLanguage2}}selected{{end}}>C++</option>
  <option value="3" {{if .SearchLanguage3}}selected{{end}}>Java</option>
</select>
<input name="commit" type="submit" value="Go">
</form>

<div class="pagination">
  {{$current := .CurrentPage}}
  {{$url := .URL}}
  {{if .IsPreviousPage}}
  <a href="{{$url}}/page/{{NumSub .CurrentPage 1}}">Prev</a>
  {{else}}
  <span>Prev</span>
  {{end}}

  {{if .IsPageHead}}
    {{with .PageHeadList}}
      {{range .}}
        {{if NumEqual . $current}}
          <span>{{.}}</span>
        {{else}}
          <a href="{{$url}}/page/{{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}

  {{if .IsPageMid}}
  ...
    {{with .PageMidList}}
      {{range .}}
        {{if NumEqual . $current}}
          <span>{{.}}</span>
        {{else}}
          <a href="{{$url}}/page/{{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}

  {{if .IsPageTail}}
  ...
    {{with .PageTailList}}
      {{range .}}
        {{if NumEqual . $current}}
          <span>{{.}}</span>
        {{else}}
          <a href="{{$url}}/page/{{.}}">{{.}}</a>
        {{end}}
      {{end}}
    {{end}}
  {{end}}

  {{if .IsNextPage}}
  <a href="{{$url}}/page/{{NumAdd .CurrentPage 1}}">Next</a>
  {{else}}
  <span>Next</span>
  {{end}}
</div>

<table id="contest_list">
  <thead>
    <tr>
      <th class="header">ID</th>
      <th class="header">User</th>
      <th class="header">Problem</th>
      <th class="header">Result</th>
      <th class="header">Time</th>
      <th class="header">Memory</th>
      <th class="header">Language</th>
      <th class="header">Code Length</th>
      <th class="header">Submit Time</th>
    </tr>
  </thead>
  <tbody>
    {{with .Solution}}  
      {{range .}} 
        {{if ShowStatus .Status}} 
          <tr>
            <td>{{.Sid}}</td>
            <td>{{.Uid}}</td>
            <td><a href="/problem/detail/pid/{{.Pid}}">{{.Pid}}</a></td>
            <td>{{ShowJudge .Judge}}</td>
            <td>{{.Time}}ms</td>
            <td>{{.Memory}}kB</td>
            <td>{{ShowLanguage .Language}}<a href="/status/code/sid/{{.Sid}}"> <a href="/status/code/sid/{{.Sid}}">[view]</a></td>
            <td>{{.Length}}B</td>
            <td>{{.Create}}</td>
          </tr>
        {{end}}
      {{end}}  
    {{end}}
  </tbody>
</table>
<script type="text/javascript">
  $('#search_form').submit( function(e) {
    e.preventDefault();
    var uid = $('#search_uid').val();
    var pid = $('#search_pid').val();
    var judge = $('#search_judge').val();
    var language = $('#search_language').val();
    var url = '/status/list';
    if (uid != '')
      url += '/uid/' + uid;
    if (pid != '')
      url += '/pid/' + pid;
    if (judge > 0)
      url += '/judge/' + judge;
    if (language > 0)
      url += '/language/' + language;
    window.location.href = url;
  });
  </script>
{{end}}
