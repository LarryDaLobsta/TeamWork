# 📘 Go Templates: Global Namespace & Block Pitfalls

## Context
During development of this project, I ran into repeated issues where page content would “bleed” into other pages when using Go templates with Fiber.

Examples of symptoms:
- The login dashboard rendering on the index page
- The signup page rendering dashboard content
- Pages changing depending on parse order
- Correct handlers rendering incorrect HTML

These issues were **not caused by Fiber routing or handlers**, but by how Go templates work.

This document captures the key lessons learned so I don’t repeat these mistakes in future projects.

---

## 🔑 Core Concept: Go Templates Are Globally Scoped

In Go’s `html/template` system:

- **All templates are loaded into a single global namespace**
- `{{define "name"}}` creates a **global identifier**
- There is **no file-level scoping**
- The last template parsed with a given name wins

This means:

```gotemplate
{{define "content"}} ... {{end}}
```

defined in multiple files **will overwrite each other**, even if they appear in different directories.

---

## ⚠️ Why `{{block "content"}}` Caused Problems

The idiomatic Go pattern looks like this:

```gotemplate
{{block "content" .}}{{end}}
```

However, this only works safely when:
- You control template parse order
- Only one page defines `"content"`
- Templates are not all loaded at once

With Fiber’s template engine:
- All templates are parsed globally
- Multiple pages defining `"content"` collide
- Parse order is non-obvious
- Content “bleeds” between pages

This makes shared block names unsafe in non-trivial apps.

---

## ❌ What *Not* to Do

- Do **not** define `"content"` in multiple page templates
- Do **not** expect blocks to be scoped per page
- Do **not** try to dynamically select templates:
  ```gotemplate
  {{template .BodyTemplate .}} // ❌ not allowed
  ```
- Do **not** render `"base"` directly from handlers

---

## ✅ The Reliable Solution: Explicit Layout Composition

To avoid collisions entirely, use **explicit composition** instead of blocks.

### Layout is split into two templates:

#### `layout_start.html`
```gotemplate
{{define "layout_start"}}
<!doctype html>
<html>
  <head>
    <meta charset="UTF-8" />
    <title>{{.Title}}</title>
  </head>
  <body>
    {{template "header" .}}
    <main>
{{end}}
```

#### `layout_end.html`
```gotemplate
{{define "layout_end"}}
    </main>
  </body>
</html>
{{end}}
```

---

## 🧱 Page Templates (No Shared Blocks)

Each page explicitly composes its layout:

### `createuser.html`
```gotemplate
{{define "createuser"}}
  {{template "layout_start" .}}

  <h1>Create User</h1>
  {{template "create_user_form" .}}

  {{template "layout_end" .}}
{{end}}
```

### `index.html`
```gotemplate
{{define "index"}}
  {{template "layout_start" .}}
    <h1>Home</h1>
  {{template "layout_end" .}}
{{end}}
```

### `logindashboard.html`
```gotemplate
{{define "logindashboard"}}
  {{template "layout_start" .}}
    <h1>Dashboard</h1>
  {{template "layout_end" .}}
{{end}}
```

---

## 🧠 Why This Works

- No shared `"content"` name
- No `block` overrides
- No dynamic template names
- Deterministic rendering
- Compatible with Fiber’s global template loader
- Scales cleanly as pages increase

This approach trades a small amount of repetition for:
- clarity
- correctness
- debuggability

---

## 🧩 What Should Be Global Templates

Safe to keep global:
- `header`
- `footer`
- navigation
- small UI partials (forms, alerts, etc.)

Unsafe to share globally:
- page body content
- layout blocks like `"content"`
- anything page-specific

---

## 📝 Final Takeaway

> **Explicit composition beats clever abstraction in Go templates.**

When using Go templates in real applications (especially with Fiber), **assume everything is global** and design accordingly.

This avoids:
- mysterious bugs
- render order issues
- accidental overrides
- wasted debugging time
