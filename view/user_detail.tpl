{{define "content"}}
  <h1>User Detail</h1>
  {{with .Detail}}
  <table>
    <tbody>
        <tr>
          <th>Handle</th>
          <td>{{.Uid}}</td>
        </tr>
        <tr>
          <th>Nick</th>
          <td>{{.Nick}}</td>
        </tr>
        <tr>
          <th>Email</th>
          <td>{{.Mail}}</td>
        </tr>
        <tr>
          <th>Motto</th>
          <td>{{.Motto}}</td>
        </tr>
    </tbody>
  </table>
  {{end}}
{{end}}