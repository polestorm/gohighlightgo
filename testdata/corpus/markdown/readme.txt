[//]: # (REQ: LANG-181 LANG-182 LANG-183 LANG-184 LANG-185 LANG-186)
# gohighlightgo

Setext heading
==============

Text with `code`, ``code with ` inside``, an unmatched ` backtick, a [link](https://example.com/a)
and <https://example.com> or <user@example.com>, escaped \*stars\* and &amp; entity.

> Quote with **bold** text
> > nested quote

- item one
* item two
  1. ordered
  2) other
- [ ] task
- [x] done task

[ref]: https://example.com "Title"

---

```go
func main() { fmt.Println("hi") }
```

~~~
plain code block <&>
~~~

````markdown
```go
not highlighted as Go inside Markdown
```
````

<div class="note">HTML block</div>

```unknown
kept as a string
