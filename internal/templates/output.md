{{ if gt (len .Lint.Issues) 0 }}
## Lint
{{ if gt (len .Lint.Report.Warnings) 0 }}
Warnings: 
<details open>
<pre>{{ range .Lint.Report.Warnings }}{{.Tag}}: {{.Text}}
{{ end }} 
</pre>
</details>
{{ end }}


Issues:
<table>
    <tr>
        <th>Linter</th>
        <th>Issue</th>
        <th>File</th> 
    </tr>
    {{ range .Lint.Issues }}<tr>
        <td>{{ .FromLinter }}</td>
        <td><code>{{ .Text }}</code></td>
        <td>
          <a href="{{ (link .Pos.Filename .Pos.Line) }}">
            {{ .Pos.Filename }}:{{ .Pos.Line }}:{{ .Pos.Column }}
          </a>
        </td>
</tr>{{ end }}
</table>
{{ end }}


## Coverage
{{ $totalStatements := .TotalCoverage.Statements }}
{{ $totalCovered := .TotalCoverage.Covered }}
{{ $totalCoverage := (percent $totalCovered $totalStatements) }}
{{ $uncoveredCode := (trim (substract 100.00 $totalCoverage)) }}
```{{(coverPie .TotalCoverage)}}
```

<table>
<tr>
    <th>📦 Package</th>
    <th>Coverage</th> 
</tr>
{{- range $key, $value := .PackageCoverage }}
{{if ne $key "*"}}
<tr>
    <td>{{$key}}</td>
{{ $totalStatements := .Statements }}
{{ $totalCovered := .Covered }}
{{ $totalCoverage := printf "%.2f%%" (percent $totalCovered $totalStatements) }}
<td>{{$totalCoverage}}</td>
</tr> 
{{end}}
{{- end}}
</table>

## Tests
```{{(testPie .TestsResult)}}
```
<table>
    <tr>
        <th>📦 Package</th>
        <th>Passed</th>
        <th>Skipped</th>
        <th>Failed</th>
        <th>Duration</th>
    </tr>
{{- range $key, $value := .TestsResult }}
    <tr>
    <td>{{(pkgBadge $value)}} {{$key}}</td>
    <td>{{(pkgPassedCount $value)}}</td>
    <td>{{(pkgSkippedCount $value)}}</td>
    <td>{{(pkgPassedCount $value)}}</td>
    <td>{{.Elapsed}}</td>
    </tr>
    <tr>
    <td colspan="5">{{if eq (len $value.Tests) 0}}
No tests found
{{end}}
{{- range $key, $value := .Tests }}
        {{(testBadge $value)}} <code>{{$value.Name}}</code>
<details {{(detailOpened $value.Output)}}><pre><code>{{(testOutput $value)}}
</code></pre></details>
{{- end }}
{{- end }}</td>
    </tr>
</table> 