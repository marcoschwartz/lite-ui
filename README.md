# Lite UI

A lightweight, independent UI component library for Go + Templ applications.

Built with [Templ](https://templ.guide/), [Tailwind CSS](https://tailwindcss.com/), [HTMX](https://htmx.org/), and [Alpine.js](https://alpinejs.dev/).

## Features

✨ **41 Production-Ready Components** - Complete UI toolkit for modern web apps
🎨 **Tailwind CSS v4** - Modern utility-first styling with dark mode
⚡ **HTMX Integration** - Smooth SPA-like experiences
🔧 **Type-Safe** - Full Go type safety with Templ
📦 **Zero Dependencies** - Only depends on `github.com/a-h/templ`
🌙 **Dark Mode** - Built-in dark mode support
📱 **Responsive** - Mobile-first design

## Quick Links

- 📖 **[Component Reference](COMPONENTS.md)** - Complete API documentation for all 41 components
- 🎨 **[Live Demo](examples/demo-app)** - Interactive component examples
- 🚀 **[Quick Start](#quick-start)** - Get started in minutes

## Installation

```bash
go get github.com/marcoschwartz/lite-ui@latest
```

## Quick Start

### 1. Import Components

```go
package main

import (
    ui "github.com/marcoschwartz/lite-ui/components"
)

templ MyPage() {
    @ui.Button("Click me", ui.ButtonPrimary, templ.Attributes{})
}
```

### 2. Set Up Tailwind CSS

Create an `input.css` file with the following content:

```css
@import "tailwindcss";

/* Alpine.js x-cloak - Prevents flash of content before Alpine loads */
[x-cloak] {
	display: none !important;
}
```

Configure Tailwind:

```javascript
// tailwind.config.js
export default {
  content: [
    './templates/**/*.templ',
    './node_modules/github.com/marcoschwartz/lite-ui/components/**/*.templ'
  ],
  theme: {
    extend: {}
  }
}
```

Build your CSS:

```bash
npx @tailwindcss/cli -i ./static/css/input.css -o ./static/css/output.css --minify
```

**Important:** The `[x-cloak]` CSS rule is required to prevent flashing of dropdowns, modals, and other Alpine.js components before the JavaScript loads.

### 3. Include HTMX and Alpine.js

```html
<script src="https://unpkg.com/htmx.org@2.0.4"></script>
<script defer src="https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"></script>
```

## Usage Example

```go
package main

import ui "github.com/marcoschwartz/lite-ui/components"

templ MyPage() {
    // Button
    @ui.Button("Click me", ui.ButtonPrimary, templ.Attributes{})

    // Card
    @ui.Card(ui.CardConfig{Title: "User Profile"}, ProfileContent())

    // Table with actions
    @ui.Table([]string{"Name", "Email", "Actions"}, UserRows())
}

templ UserRows() {
    <tr>
        <td class="px-6 py-4">John Doe</td>
        <td class="px-6 py-4">john@example.com</td>
        <td class="px-6 py-4">
            @ui.ActionMenu([]ui.ActionMenuItem{
                {Label: "Edit", Href: "/edit/1"},
                {Label: "Delete", Href: "/delete/1", Danger: true, Divider: true},
            })
        </td>
    </tr>
}
```

For detailed component API documentation, see **[COMPONENTS.md](COMPONENTS.md)**.

## Important Notes

### x-cloak CSS (Required)

**You must include this CSS** to prevent component flashing:

```css
[x-cloak] {
    display: none !important;
}
```

Without this, dropdowns, modals, and other interactive components will briefly flash before Alpine.js loads.

### Dark Mode

Add this script to enable dark mode:

```html
<script>
    const theme = localStorage.getItem('theme') ||
        (window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light');
    if (theme === 'dark') document.documentElement.classList.add('dark');
</script>
```

Then use the toggle: `@ui.DarkModeToggle()`

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details.

## Credits

Built by [Marco Schwartz](https://github.com/marcoschwartz)

Powered by:
- [Templ](https://templ.guide/) - Type-safe HTML templating for Go
- [Tailwind CSS](https://tailwindcss.com/) - Utility-first CSS framework
- [HTMX](https://htmx.org/) - High power tools for HTML
- [Alpine.js](https://alpinejs.dev/) - Lightweight JavaScript framework
