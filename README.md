# Lite UI

A lightweight, independent UI component library for Go + Templ applications.

Built with [Templ](https://templ.guide/), [Tailwind CSS](https://tailwindcss.com/), [HTMX](https://htmx.org/), and [Alpine.js](https://alpinejs.dev/).

## Features

✨ **25 Production-Ready Components** - Buttons, cards, forms, modals, tables, and more
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
- `ActionMenu` - Row action menu (for tables)

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

### Table with Action Menu

Use the `ActionMenu` component for table row actions. It automatically handles positioning to avoid clipping:

```go
<tr>
    <td class="px-6 py-4">John Doe</td>
    <td class="px-6 py-4">john@example.com</td>
    <td class="px-6 py-4">
        @ui.ActionMenu([]ui.ActionMenuItem{
            {Label: "Edit", Href: "/edit/1"},
            {Label: "View Details", Href: "/view/1"},
            {Label: "Delete", Href: "/delete/1", Danger: true, Divider: true},
        })
    </td>
</tr>
```

The `ActionMenu` uses fixed positioning to prevent clipping by table containers.

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

### ActionMenu

```go
type ActionMenuItem struct {
    Label   string
    Href    string
    Danger  bool // Style as dangerous action (red)
    Divider bool // Show divider before this item
}

templ ActionMenu(items []ActionMenuItem)
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

Lite UI components use Alpine.js for interactivity and include `x-cloak` attributes to prevent content flashing before Alpine initializes.

**Required CSS:** You must include this CSS rule in your stylesheet (see Setup step 2):

```css
[x-cloak] {
	display: none !important;
}
```

**How it works:** Components with dropdown menus, modals, drawers, accordions, tabs, carousels, and tooltips use `x-cloak` to remain hidden until Alpine.js loads. This ensures a smooth, flash-free user experience.

**Without this CSS rule:** Users will briefly see hidden content (dropdowns, modals, etc.) before Alpine.js initializes and hides them, creating an unpleasant visual flash.

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
