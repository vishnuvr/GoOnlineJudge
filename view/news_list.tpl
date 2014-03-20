{{define "content"}}
<h2 class="compact">News List</h2>
{{with .News}}
	{{range .}}
		{{if ShowStatus .Status}}
			<p class="news">
				<span class="flag"></span>
				<span class="date">{{.Create}}</span>		
				<br>{{.Title}}
				<a href="/news/detail/nid/{{.Nid}}">[Details]</a>
			</p>
		{{end}}
	{{end}}
{{end}}
{{end}}
