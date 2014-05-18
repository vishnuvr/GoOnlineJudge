{{define "content"}}
<h1>Admin - User Online List</h1>
<table id="user_list">
	<thead>
		<tr>
		    <th class="header">Uid</th>    
		    <th class="header">Status</th>
		    <th class="header">Last Operate</th>
		    <th class="header">Down Line</th>
		</tr>
	</thead>
	<tbody>
		{{with .User}}
			{{range .}}
				<tr>
					<td><a href="#">{{.Uid}}</a></td>
					<td><a href="#">{{if ShowStatus .Status}}Available{{else}}Reserved{{end}}</a></td>
					<td><a href="#">{{.Last}}</a></td>
					<td><a href="#">Down Line</a></td>
				</tr>
			{{end}}
		{{end}}
	</tbody>
</table>
{{end}}