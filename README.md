# Lite UI

A lightweight, independent UI component library for Go + Templ applications.

Built with [Templ](https://templ.guide/), [Tailwind CSS](https://tailwindcss.com/), [HTMX](https://htmx.org/), and [Alpine.js](https://alpinejs.dev/).

## Features

✨ **24 Production-Ready Components** - Buttons, cards, forms, modals, tables, and more
🎨 **Tailwind CSS v4** - Modern utility-first styling with dark mode support
⚡ **HTMX Integration** - Smooth SPA-like experiences with server-side rendering
🔧 **Type-Safe** - Full Go type safety with Templ
📦 **Zero Dependencies** - Only depends on `github.com/a-h/templ`
🌙 **Dark Mode** - Built-in dark mode toggle and support
📱 **Responsive** - Mobile-first design with responsive utilities

## Components

### Forms
- `Button` - Primary, secondary, danger, success variants
- `TextInput` - Text, email, password, etc.
- `TextArea` - Multi-line text input
- `Select` - Dropdown select with options
- `Checkbox` - Checkbox with label
- `Switch` - Toggle switch
- `FormGroup` - Form field wrapper

### Layout
- `Card` - Versatile card with header, stat variants
- `Sidebar` - Responsive sidebar with mobile drawer
- `Navbar` - Top navigation bar
- `Modal` - Modal dialog overlay
- `Drawer` - Slide-out drawer
- `Tabs` - Tabbed content

### Data Display
- `Table` - Responsive table with mobile cards
- `Pagination` - Full-featured pagination
- `Badge` - Status and label badges
- `Avatar` - User avatars
- `Timeline` - Event timeline
- `Stepper` - Multi-step progress indicator

### Feedback
- `Progress` - Progress bar
- `Slider` - Range slider

### Navigation
- `Dropdown` - Dropdown menu

### Utilities
- `DarkModeToggle` - Theme switcher
- `Icons` - Common icon set

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

Copy `static/css/input.css` to your project and configure Tailwind:

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

### 3. Include HTMX and Alpine.js

```html
<script src="https://unpkg.com/htmx.org@2.0.4"></script>
<script defer src="https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"></script>
```

## Usage Examples

### Button

```go
@ui.Button("Save", ui.ButtonPrimary, templ.Attributes{
    "type": "submit",
})
```

### Card with Header

```go
@ui.Card(ui.CardConfig{
    Title: "User Profile",
    HeaderAction: EditButton(),
}, CardContent())
```

### Stat Card

```go
@ui.Card(ui.CardConfig{
    Variant: "stat",
    StatLabel: "Total Users",
    StatValue: "1,234",
    StatIcon: ui.IconUsers("indigo"),
    StatColor: "indigo",
}, nil)
```

### Modal

```go
@ui.Modal("confirm_modal", "Confirm Action",
    TriggerButton(),
    ModalContent()
)
```

### Sidebar with Custom Types

The sidebar uses a `Project` type. If you have your own project type, create an adapter:

```go
// Convert your type to ui.Project
func ToUIProject(p MyProject) ui.Project {
    return ui.Project{
        ID:   p.ID,
        Name: p.Name,
        Slug: p.Slug,
    }
}

// Use in template
@ui.Sidebar(
    items,
    userName,
    userEmail,
    userAvatar,
    ToUIProject(currentProject),
    ToUIProjects(projects),
)
```

### Table with Pagination

```go
@ui.Table(
    []string{"Name", "Email", "Status"},
    TableRows(users),
)

@ui.Pagination(ui.PaginationConfig{
    CurrentPage: page,
    TotalPages:  totalPages,
    TotalItems:  totalItems,
    PerPage:     10,
    BasePath:    "/users",
})
```

## Component API

### Button

```go
type ButtonVariant string

const (
    ButtonPrimary   ButtonVariant = "primary"
    ButtonSecondary ButtonVariant = "secondary"
    ButtonDanger    ButtonVariant = "danger"
    ButtonSuccess   ButtonVariant = "success"
)

templ Button(text string, variant ButtonVariant, attrs templ.Attributes)
```

### Card

```go
type CardConfig struct {
    // Header
    Title        string
    HeaderAction templ.Component

    // Stat variant
    Variant    string // "stat"
    StatLabel  string
    StatValue  string
    StatIcon   templ.Component
    StatColor  string // "indigo", "green", "purple", etc.

    // Layout
    Padding string // "none" for no padding
}

templ Card(config CardConfig, content templ.Component)
```

### Sidebar

```go
type Project struct {
    ID   int
    Name string
    Slug string
}

type SidebarItem struct {
    Label  string
    Href   string
    Icon   templ.Component
    Active bool
    Badge  string
}

templ Sidebar(
    items []SidebarItem,
    userName string,
    userEmail string,
    userAvatar string,
    currentProject Project,
    projects []Project,
)
```

### Pagination

```go
type PaginationConfig struct {
    CurrentPage int
    TotalPages  int
    TotalItems  int
    PerPage     int
    BasePath    string // e.g., "/users"
}

templ Pagination(config PaginationConfig)
```

## Dark Mode

Lite UI includes automatic dark mode support using Tailwind's dark mode classes.

Add the dark mode script to your HTML:

```html
<script>
    const theme = localStorage.getItem('theme') ||
        (window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light');
    if (theme === 'dark') {
        document.documentElement.classList.add('dark');
    }
</script>
```

Use the dark mode toggle component:

```go
@ui.DarkModeToggle()
```

## Alpine.js Integration

Lite UI includes built-in support for Alpine.js `x-cloak` to prevent content flashing before Alpine initializes. The `input.css` file already includes the necessary styles:

```css
[x-cloak] {
	display: none !important;
}
```

Components like Modal, Drawer, and Sidebar automatically use `x-cloak` to ensure smooth loading without visual flashing.

## HTMX Integration

Components are designed to work seamlessly with HTMX for dynamic interactions:

```go
<a
    href="/users"
    hx-get="/users"
    hx-target="body"
    hx-swap="outerHTML transition:true"
    hx-push-url="true"
>
    Users
</a>
```

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
