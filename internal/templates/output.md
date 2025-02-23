## Lint

<table>
    <tr>
        <th>Linter</th>
        <th>Issue</th>
        <th>File</th>
        <th>Column</th>
    </tr>
    {{ range .Lint.Issues }}
        <tr>
            <td>{{ .FromLinter }}</td>
            <td>{{ .Text }}</td>
            <td>{{ .File }}</td>
            <td>{{ .Pos.Filename }}:{{ .Pos.Line }}:{{ .Pos.Column }}</td>
        </tr>
    {{ end }}
</table>


## Coverage
```mermaid
%%{init: {"theme":"base","themeVariables":{"fontFamily":"monospace","pieSectionTextSize":"24px","darkMode":true,"pie1":"#2da44e","pie2":"#cf222e","pie3":"#dbab0a"}}}%%
pie
    "Covered": 1
    "Uncovered": 2
```


## Tests
```mermaid
%%{init: {"theme":"base","themeVariables":{"fontFamily":"monospace","pieSectionTextSize":"24px","darkMode":true,"pie1":"#2da44e","pie2":"#cf222e","pie3":"#dbab0a"}}}%%
pie
    "Skip": 2
    "Fail": 1
    "Pass": 3
```
<table>
    <tr>
        <th>Package</th>
        <th>Passed</th>
        <th>Skipped</th>
        <th>Duration</th>
    </tr>
</table> 