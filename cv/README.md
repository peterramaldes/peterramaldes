# Peter A. Ramaldes, Curriculum Vitae

All content lives in a single [data.yml](data.yml). `index.html` is generated
from it via [templates](tmpl/index.html) using Go's `html/template`:

```
go run gen.go
```

`gen.go` reads `data.yml`, then walks `tmpl/`: a subdirectory (e.g.
`tmpl/index.html/`) is rendered as a set of templates glob'd together and
written to a file named after that subdirectory; a lone file is rendered
on its own the same way.

Projects, volunteer history, and accolades aren't populated yet — add them as
they come up.

Once finished, the intent is that the PDF version is created by printing from
Firefox with all the extras disabled, ensuring it contains transferable
fonts.

The structure and tooling here (YAML data + Go `html/template` generator) is
adapted from [rwxrob/rwxrob](https://github.com/rwxrob/rwxrob) (Apache-2.0
licensed), with all personal content replaced.
