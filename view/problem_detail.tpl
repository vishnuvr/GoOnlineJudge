{{define "content"}}
  {{with .Detail}}
    <h1>{{.Title}}</h1>
    <div id="problemInfo" class="rfloat">
      <div class="limit">
        <div class="key">Time Limit</div>
        <div class="value">{{.Time}}ms<br></div>
      </div>
      <div class="limit">
        <div class="key">Memory Limit</div>
        <div class="value">{{.Memory}}kB<br></div>
      </div>
      <div class="checker">
        <div class="key">Judge Program</div>
        <div class="value">
          <span title="纯文本对比">{{ShowSpecial .Special}}</span>
        </div>
      </div>
      <div class="checker">
        <div class="key">Ratio</div>
        <div class="value">
          <span title="纯文本对比">{{ShowRatio .Solve .Submit}}({{.Solve}}/{{.Submit}})</span>
        </div>
      </div>
    </div>
    <div id="problemContent">
      <p class="submitRes-1"><b>Description:</b></p>
      <p>{{.Description}}</p>
      <p class="submitRes-1"><b>Input:</b></p>
      <p>{{.Input}}</p>
      <p class="submitRes-1"><b>Output:</b></p>
      <p>{{.Output}}</p>
      <b class="submitRes-1">Sample Input:</b>
      <pre>{{.In}}</pre>
      <b class="submitRes-1">Sample Output:</b>
      <pre>{{.Out}}</pre>
      {{if .Hint}}
        <p class="submitRes-1"><b>Hint:</b></p>
        <p>{{.Hint}}</p>
      {{end}}
      {{if .Source}}
        <p class="submitRes-1"><b>Source:</b></p>
        <p>{{.Source}}</p>
      {{end}}
    </div>
    <hr>
  {{end}}
{{end}}






