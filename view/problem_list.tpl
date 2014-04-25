{{define "content"}}
<h1>Problem List</h1>

<form accept-charset="UTF-8" id="search_form">
Search: <input id="search" name="search" size="30" type="text" value="{{.SearchValue}}">
<select id="option" name="option">
  <option value="pid" {{if .SearchPid}}selected{{end}}>ID</option>
  <option value="title" {{if .SearchTitle}}selected{{end}}>Title</option>
  <option value="source" {{if .SearchSource}}selected{{end}}>Source</option>
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
      <th class="header">Title</th>
      <th class="header">Ratio</th>
    </tr>
  </thead>
  <tbody>
    {{$time := .Time}}
    {{with .Problem}}  
      {{range .}} 
        {{if ShowStatus .Status}}
          {{if ShowExpire .Expire $time}}
            <tr>
              <td>{{.Pid}}</td>
              <td><a href="/problem/detail/pid/{{.Pid}}">{{.Title}}</a></td>
              <td>{{ShowRatio .Solve .Submit}} ({{.Solve}}/{{.Submit}})</td>
            </tr>
          {{end}}
        {{end}}
      {{end}}  
    {{end}}
  </tbody>
</table>

<script type="text/javascript">
  $('#search_form').submit( function(e) {
    e.preventDefault();
    var value = $('#search').val();
    var key = $('#option').val();
    window.location.href = '/problem/list/'+key+'/'+value;
  });
  </script>
{{end}}
