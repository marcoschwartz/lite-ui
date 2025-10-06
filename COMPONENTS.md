# Component Reference

Complete API documentation for all Lite UI components.

## Table of Contents

- [Forms](#forms)
  - [Button](#button)
  - [TextInput](#textinput)
  - [TextArea](#textarea)
  - [Select](#select)
  - [Checkbox](#checkbox)
  - [Radio](#radio)
  - [Switch](#switch)
  - [Toggle](#toggle)
  - [FileInput](#fileinput)
  - [Rating](#rating)
  - [FormGroup](#formgroup)
- [Layout](#layout)
  - [Card](#card)
  - [Sidebar](#sidebar)
  - [Navbar](#navbar)
  - [Modal](#modal)
  - [Drawer](#drawer)
  - [Tabs](#tabs)
  - [Divider](#divider)
  - [Hero](#hero)
  - [Footer](#footer)
  - [Collapse](#collapse)
- [Data Display](#data-display)
  - [Table](#table)
  - [Pagination](#pagination)
  - [Badge](#badge)
  - [Avatar](#avatar)
  - [Timeline](#timeline)
  - [Stepper](#stepper)
  - [List](#list)
  - [Accordion](#accordion)
  - [Carousel](#carousel)
  - [Calendar](#calendar)
  - [EventCalendar](#eventcalendar)
  - [ChatBubble](#chatbubble)
- [Feedback](#feedback)
  - [Alert](#alert)
  - [Progress](#progress)
  - [Loading](#loading)
  - [Tooltip](#tooltip)
  - [Toast](#toast)
  - [Skeleton](#skeleton)
  - [RadialProgress](#radialprogress)
- [Navigation](#navigation)
  - [Dropdown](#dropdown)
  - [ActionMenu](#actionmenu)
  - [Menu](#menu)
  - [Breadcrumbs](#breadcrumbs)
- [Date & Time](#date--time)
  - [DatePicker](#datepicker)
  - [Countdown](#countdown)
  - [Slider](#slider)
- [Utilities](#utilities)
  - [DarkModeToggle](#darkmodetoggle)
  - [Kbd](#kbd)
  - [FAB](#fab)
  - [Icons](#icons)

---

## Forms

### Button

Button component with multiple variants and loading state support.

**API**

```go
type ButtonVariant string

const (
    ButtonPrimary   ButtonVariant = "primary"
    ButtonSecondary ButtonVariant = "secondary"
    ButtonDanger    ButtonVariant = "danger"
    ButtonSuccess   ButtonVariant = "success"
)

templ Button(text string, variant ButtonVariant, attrs templ.Attributes)
templ LoadingButton(text string, variant ButtonVariant, loadingVar string, attrs templ.Attributes)
```

**Usage**

```go
// Basic button
@ui.Button("Save", ui.ButtonPrimary, templ.Attributes{
    "type": "submit",
})

// Loading button
@ui.LoadingButton("Submit", ui.ButtonPrimary, "isLoading", templ.Attributes{
    "x-data": "{ isLoading: false }",
    "@click": "isLoading = true",
})
```

---

### TextInput

Text input field with label and error state support.

**API**

```go
templ TextInput(
    label string,
    inputType string,  // "text", "email", "password", etc.
    name string,
    placeholder string,
    required bool,
    attrs templ.Attributes,
)
```

**Usage**

```go
@ui.TextInput("Email", "email", "email", "Enter your email", true, templ.Attributes{})
```

---

### TextArea

Multi-line text input.

**API**

```go
templ TextArea(
    label string,
    name string,
    placeholder string,
    rows int,
    required bool,
    attrs templ.Attributes,
)
```

**Usage**

```go
@ui.TextArea("Description", "description", "Enter description...", 4, false, templ.Attributes{})
```

---

### Select

Dropdown select with custom styling.

**API**

```go
type SelectOption struct {
    Value string
    Label string
}

templ Select(
    label string,
    name string,
    placeholder string,
    required bool,
    options []SelectOption,
)
```

**Usage**

```go
@ui.Select("Country", "country", "Select a country", false, []ui.SelectOption{
    {Value: "us", Label: "United States"},
    {Value: "uk", Label: "United Kingdom"},
})
```

---

### Checkbox

Checkbox with label.

**API**

```go
templ Checkbox(
    label string,
    name string,
    checked bool,
    attrs templ.Attributes,
)
```

**Usage**

```go
@ui.Checkbox("Accept terms", "terms", false, templ.Attributes{})
```

---

### Radio

Radio button input.

**API**

```go
templ Radio(
    label string,
    name string,
    value string,
    checked bool,
    attrs templ.Attributes,
)
```

**Usage**

```go
@ui.Radio("Option 1", "choice", "option1", false, templ.Attributes{})
@ui.Radio("Option 2", "choice", "option2", true, templ.Attributes{})
```

---

### Switch

Toggle switch component.

**API**

```go
templ Switch(
    name string,
    label string,
    description string,
    checked bool,
    attrs templ.Attributes,
)
```

**Usage**

```go
@ui.Switch("notifications", "Enable notifications", "Receive push notifications", true, templ.Attributes{})
```

---

### Toggle

Alternative toggle component.

**API**

```go
templ Toggle(
    label string,
    name string,
    checked bool,
    attrs templ.Attributes,
)
```

**Usage**

```go
@ui.Toggle("Dark Mode", "theme", false, templ.Attributes{})
```

---

### FileInput

File upload input with drag-and-drop support.

**API**

```go
templ FileInput(
    label string,
    name string,
    accept string,
    required bool,
    attrs templ.Attributes,
)
```

**Usage**

```go
@ui.FileInput("Upload Image", "photo", "image/*", false, templ.Attributes{})
```

---

### Rating

Star rating component.

**API**

```go
templ Rating(
    name string,
    maxStars int,
    currentRating int,
    attrs templ.Attributes,
)
```

**Usage**

```go
@ui.Rating("rating", 5, 3, templ.Attributes{})
```

---

### FormGroup

Wrapper for form fields with label and error message.

**API**

```go
templ FormGroup(
    label string,
    required bool,
    error string,
    input templ.Component,
)
```

**Usage**

```go
@ui.FormGroup("Email", true, "", ui.TextInput("", "email", "email", "Enter email", true, templ.Attributes{}))
```

---

## Layout

### Card

Versatile card component with header and stat variants.

**API**

```go
type CardConfig struct {
    // Header variant
    Title        string
    HeaderAction templ.Component

    // Stat variant
    Variant    string // "stat"
    StatLabel  string
    StatValue  string
    StatIcon   templ.Component
    StatColor  string // "indigo", "green", "purple", etc.

    // Layout
    Padding string // "default", "none"
}

templ Card(config CardConfig, content templ.Component)
```

**Usage**

```go
// Basic card
@ui.Card(ui.CardConfig{}, CardContent())

// Card with header
@ui.Card(ui.CardConfig{
    Title: "User Profile",
    HeaderAction: EditButton(),
}, ProfileContent())

// Stat card
@ui.Card(ui.CardConfig{
    Variant: "stat",
    StatLabel: "Total Users",
    StatValue: "1,234",
    StatIcon: ui.IconUsers("indigo"),
    StatColor: "indigo",
}, nil)
```

---

### Sidebar

Responsive sidebar with mobile drawer.

**API**

```go
templ Sidebar(
    header templ.Component,
    content templ.Component,
    footer templ.Component,
)
```

**Usage**

```go
@ui.Sidebar(
    SidebarHeader(),
    SidebarNav(),
    SidebarFooter(),
)
```

---

### Navbar

Top navigation bar.

**API**

```go
templ Navbar(logo templ.Component, content templ.Component)
```

**Usage**

```go
@ui.Navbar(Logo(), NavItems())
```

---

### Modal

Modal dialog overlay.

**API**

```go
templ Modal(
    id string,
    title string,
    content templ.Component,
)
```

**Usage**

```go
// Open with Alpine.js event
@ui.Button("Open Modal", ui.ButtonPrimary, templ.Attributes{
    "@click": "window.dispatchEvent(new CustomEvent('open-modal', { detail: 'confirm_modal' }))",
})

@ui.Modal("confirm_modal", "Confirm Action", ConfirmContent())
```

---

### Drawer

Slide-out drawer panel.

**API**

```go
templ Drawer(
    id string,
    title string,
    content templ.Component,
)
```

**Usage**

```go
@ui.Drawer("settings_drawer", "Settings", SettingsContent())
```

---

### Tabs

Tabbed content with smooth transitions.

**API**

```go
type Tab struct {
    ID      string
    Label   string
    Content templ.Component
}

templ Tabs(id string, tabs []Tab)
```

**Usage**

```go
@ui.Tabs("product-tabs", []ui.Tab{
    {ID: "overview", Label: "Overview", Content: OverviewContent()},
    {ID: "specs", Label: "Specifications", Content: SpecsContent()},
    {ID: "reviews", Label: "Reviews", Content: ReviewsContent()},
})
```

---

### Divider

Visual separator component.

**API**

```go
templ Divider(
    variant string, // "horizontal" or "vertical"
    text string,    // Optional text in the middle
)
```

**Usage**

```go
@ui.Divider("horizontal", "") // Simple horizontal line
@ui.Divider("horizontal", "OR") // With label
@ui.Divider("vertical", "") // Vertical divider
```

---

### Hero

Hero section component for landing pages.

**API**

```go
templ Hero(title string, description string, actions templ.Component, backgroundImage string)
templ HeroWithImage(title string, description string, actions templ.Component, imageUrl string)
templ HeroMinimal(title string, description string, actions templ.Component)
```

**Usage**

```go
// Full-width hero with gradient background
@ui.Hero("Welcome", "Build amazing apps", HeroButtons(), "")

// Hero with side image
@ui.HeroWithImage("Get Started", "Easy setup in minutes", HeroButtons(), "https://...")

// Minimal hero
@ui.HeroMinimal("Simple & Clean", "No distractions", HeroButtons())
```

---

### Footer

Website footer component with multiple layouts.

**API**

```go
type FooterSection struct {
    Title string
    Links []FooterLink
}

type FooterLink struct {
    Label string
    Href  string
}

templ Footer(sections []FooterSection, copyright string)
templ FooterSimple(copyright string, links []FooterLink)
templ FooterWithSocial(sections []FooterSection, copyright string, socialLinks templ.Component)
```

**Usage**

```go
// Multi-column footer
@ui.Footer([]ui.FooterSection{
    {
        Title: "Product",
        Links: []ui.FooterLink{
            {Label: "Features", Href: "/features"},
            {Label: "Pricing", Href: "/pricing"},
        },
    },
}, "© 2025 Company")

// Simple footer
@ui.FooterSimple("© 2025 Company", []ui.FooterLink{
    {Label: "Privacy", Href: "/privacy"},
    {Label: "Terms", Href: "/terms"},
})
```

---

### Collapse

Simple collapsible content section.

**API**

```go
templ Collapse(title string, content templ.Component, defaultOpen bool)
```

**Usage**

```go
@ui.Collapse("Details", DetailsContent(), false)
@ui.Collapse("Expanded Section", ExpandedContent(), true)
```

---

## Data Display

### Table

Responsive table component.

**API**

```go
templ Table(headers []string, content templ.Component)
templ TableRow(content templ.Component)
templ TableCell(content templ.Component)
```

**Usage**

```go
@ui.Table([]string{"Name", "Email", "Role", "Actions"}, UserRows())

templ UserRows() {
    <tr>
        <td class="px-6 py-4">John Doe</td>
        <td class="px-6 py-4">john@example.com</td>
        <td class="px-6 py-4">@ui.Badge("Admin", "red")</td>
        <td class="px-6 py-4">
            @ui.ActionMenu([]ui.ActionMenuItem{
                {Label: "Edit", Href: "/edit"},
                {Label: "Delete", Href: "/delete", Danger: true, Divider: true},
            })
        </td>
    </tr>
}
```

---

### Pagination

Full-featured pagination component with HTMX support.

**API**

```go
type PaginationConfig struct {
    CurrentPage int    // Current page (1-indexed)
    TotalPages  int    // Total number of pages
    TotalItems  int    // Total number of items
    PerPage     int    // Items per page
    BasePath    string // Base URL path (e.g., "/users")
}

templ Pagination(config PaginationConfig)
```

**Usage**

```go
@ui.Pagination(ui.PaginationConfig{
    CurrentPage: 2,
    TotalPages:  10,
    TotalItems:  100,
    PerPage:     10,
    BasePath:    "/users",
})
```

**Features**
- Page numbers display correctly (1, 2, 3... not 0, 1, 2)
- Smart page range with ellipsis for large page counts
- HTMX integration for SPA-like navigation
- Previous/Next buttons with disabled states
- Responsive design

---

### Badge

Status and label badges.

**API**

```go
templ Badge(text string, color string)
```

**Colors:** `"gray"`, `"red"`, `"yellow"`, `"green"`, `"blue"`, `"indigo"`, `"purple"`, `"pink"`

**Usage**

```go
@ui.Badge("Active", "green")
@ui.Badge("Pending", "yellow")
@ui.Badge("Error", "red")
```

---

### Avatar

User avatar component.

**API**

```go
templ Avatar(src string, alt string, size string)
```

**Sizes:** `"xs"`, `"sm"`, `"md"`, `"lg"`, `"xl"`

**Usage**

```go
@ui.Avatar("https://i.pravatar.cc/150?img=1", "User", "md")
```

---

### Timeline

Event timeline component.

**API**

```go
type TimelineItem struct {
    Title       string
    Description string
    Time        string
    Icon        templ.Component
    Color       string
}

templ Timeline(items []TimelineItem)
```

**Usage**

```go
@ui.Timeline([]ui.TimelineItem{
    {
        Title: "Order placed",
        Description: "Your order was placed successfully",
        Time: "2 hours ago",
        Icon: ui.IconCheck("green"),
        Color: "green",
    },
    {
        Title: "Processing",
        Description: "Your order is being processed",
        Time: "1 hour ago",
        Icon: ui.IconClock(),
        Color: "blue",
    },
})
```

---

### Stepper

Multi-step progress indicator.

**API**

```go
type StepperStep struct {
    Label       string
    Description string
    Status      string // "complete", "current", "upcoming"
}

templ Stepper(steps []StepperStep)
```

**Usage**

```go
@ui.Stepper([]ui.StepperStep{
    {Label: "Account", Description: "Create account", Status: "complete"},
    {Label: "Profile", Description: "Add details", Status: "current"},
    {Label: "Confirm", Description: "Review & confirm", Status: "upcoming"},
})
```

---

### List

Versatile list component with multiple variants.

**API**

```go
type ListItem struct {
    Content  templ.Component
    Icon     templ.Component
    Href     string
    Active   bool
    Disabled bool
}

templ List(items []ListItem, variant string)
```

**Variants:** `"default"`, `"bordered"`, `"divided"`

**Usage**

```go
@ui.List([]ui.ListItem{
    {Content: ListContent("Home", "Dashboard"), Icon: ui.IconHome(), Href: "/", Active: true},
    {Content: ListContent("Settings", "Preferences"), Icon: ui.IconSettings(), Href: "/settings"},
}, "divided")
```

---

### Accordion

Collapsible accordion component.

**API**

```go
type AccordionItem struct {
    ID      string
    Title   string
    Content templ.Component
}

templ Accordion(items []AccordionItem, allowMultiple bool)
```

**Usage**

```go
@ui.Accordion([]ui.AccordionItem{
    {ID: "item1", Title: "Section 1", Content: Section1()},
    {ID: "item2", Title: "Section 2", Content: Section2()},
}, false)
```

---

### Carousel

Image/content carousel.

**API**

```go
type CarouselItem struct {
    ID      string
    Content templ.Component
    Image   string
    Alt     string
}

templ Carousel(id string, items []CarouselItem)
```

**Usage**

```go
@ui.Carousel("product-carousel", []ui.CarouselItem{
    {ID: "1", Image: "https://example.com/image1.jpg", Alt: "Product 1"},
    {ID: "2", Image: "https://example.com/image2.jpg", Alt: "Product 2"},
})
```

---

### Calendar

Simple calendar component.

**API**

```go
templ Calendar(currentMonth string, days []CalendarDay)
```

**Usage**

```go
@ui.Calendar("January 2025", GetDaysForMonth())
```

---

### EventCalendar

Full-featured event calendar.

**API**

```go
type CalendarEvent struct {
    ID        string
    Title     string
    StartTime string
    EndTime   string
    AllDay    bool
    Color     string
}

templ EventCalendar(events []CalendarEvent)
```

**Usage**

```go
@ui.EventCalendar([]ui.CalendarEvent{
    {ID: "1", Title: "Meeting", StartTime: "09:00", EndTime: "10:00", Color: "blue"},
})
```

---

### ChatBubble

Chat message bubble component for messaging interfaces.

**API**

```go
templ ChatBubble(
    direction string,  // "sent" or "received"
    message string,
    time string,
    avatar templ.Component,
)

templ ChatBubbleWithImage(
    direction string,
    message string,
    imageUrl string,
    time string,
    avatar templ.Component,
)

templ ChatContainer(messages templ.Component)
```

**Usage**

```go
@ui.ChatBubble("sent", "Hello!", "10:30 AM", ui.Avatar("...", "User", "sm"))
@ui.ChatBubble("received", "Hi there!", "10:31 AM", ui.Avatar("...", "User", "sm"))

// With image
@ui.ChatBubbleWithImage("sent", "Check this out", "https://...", "10:32 AM", ui.Avatar("...", "User", "sm"))

// Full chat container
@ui.ChatContainer(ChatMessages())
```

---

## Feedback

### Alert

Alert message with variants and dismissible option.

**API**

```go
templ Alert(
    variant string, // "info", "success", "warning", "error"
    title string,
    message string,
    dismissible bool,
)
```

**Usage**

```go
@ui.Alert("success", "Success", "Changes saved successfully", true)
@ui.Alert("error", "Error", "An error occurred", false)
```

---

### Progress

Progress bar component.

**API**

```go
templ Progress(
    value int,
    max int,
    size string,    // "sm", "md", "lg"
    color string,   // "indigo", "green", "red", etc.
    showLabel bool,
)
```

**Usage**

```go
@ui.Progress(75, 100, "md", "indigo", true)
```

---

### Loading

Loading spinner component.

**API**

```go
templ Loading(size string, color string)
```

**Sizes:** `"sm"`, `"md"`, `"lg"`

**Usage**

```go
@ui.Loading("md", "indigo")
```

---

### Tooltip

Tooltip component with positioning.

**API**

```go
templ Tooltip(
    content string,
    position string, // "top", "bottom", "left", "right"
    trigger templ.Component,
)
```

**Usage**

```go
@ui.Tooltip("Click to save", "top", ui.Button("Save", ui.ButtonPrimary, templ.Attributes{}))
```

---

### Toast

Toast notification component for temporary messages.

**API**

```go
// Add ToastContainer once to your layout
templ ToastContainer()

// Trigger from button or code
window.dispatchEvent(new CustomEvent('show-toast', {
    detail: {
        message: 'Operation completed',
        variant: 'success',  // 'success', 'error', 'warning', 'info'
        title: 'Success',
        duration: 3000  // milliseconds (optional)
    }
}))
```

**Usage**

```go
// Add to your layout (once)
@ui.ToastContainer()

// Trigger from Alpine.js/HTMX
@ui.Button("Save", ui.ButtonPrimary, templ.Attributes{
    "@click": "window.dispatchEvent(new CustomEvent('show-toast', { detail: { message: 'Saved!', variant: 'success' } }))",
})
```

---

### Skeleton

Loading placeholder components with pulse animation.

**API**

```go
templ SkeletonText(size string, lines int)     // "sm", "md", "lg"
templ SkeletonCircle(size string)              // "sm", "md", "lg", "xl"
templ SkeletonBox(width string, height string) // Custom dimensions
templ SkeletonCard()                            // Pre-built card skeleton
templ SkeletonTable(rows int)                  // Table skeleton
```

**Usage**

```go
// Text loading
@ui.SkeletonText("md", 3)

// Avatar loading
@ui.SkeletonCircle("md")

// Card loading
@ui.SkeletonCard()

// Table loading
@ui.SkeletonTable(5)
```

---

### RadialProgress

Circular progress indicator.

**API**

```go
templ RadialProgress(
    value int,
    max int,
    size string,  // "sm", "md", "lg", "xl"
    color string, // "indigo", "green", "blue", "red", etc.
    showLabel bool,
)
```

**Usage**

```go
@ui.RadialProgress(75, 100, "md", "indigo", true)
```

---

## Navigation

### Dropdown

Dropdown menu component.

**API**

```go
type DropdownItem struct {
    Label string
    Href  string
}

templ Dropdown(label string, items []DropdownItem)
```

**Usage**

```go
@ui.Dropdown("Options", []ui.DropdownItem{
    {Label: "Edit", Href: "/edit"},
    {Label: "Delete", Href: "/delete"},
})
```

---

### ActionMenu

Row action menu for tables (uses fixed positioning to avoid clipping).

**API**

```go
type ActionMenuItem struct {
    Label   string
    Href    string
    Danger  bool // Red text for destructive actions
    Divider bool // Show divider before this item
}

templ ActionMenu(items []ActionMenuItem)
```

**Usage**

```go
@ui.ActionMenu([]ui.ActionMenuItem{
    {Label: "Edit", Href: "/edit/1"},
    {Label: "View Details", Href: "/view/1"},
    {Label: "Delete", Href: "/delete/1", Danger: true, Divider: true},
})
```

---

### Menu

Vertical or horizontal navigation menu.

**API**

```go
type MenuItem struct {
    Label  string
    Href   string
    Icon   templ.Component
    Active bool
    Badge  string
}

templ Menu(items []MenuItem, variant string)
```

**Variants:** `"vertical"`, `"horizontal"`, `"compact"`

**Usage**

```go
@ui.Menu([]ui.MenuItem{
    {Label: "Home", Href: "/", Icon: ui.IconHome(), Active: true},
    {Label: "Settings", Href: "/settings", Icon: ui.IconSettings(), Badge: "2"},
}, "vertical")
```

---

### Breadcrumbs

Breadcrumb navigation.

**API**

```go
type BreadcrumbItem struct {
    Label string
    Href  string
}

templ Breadcrumbs(items []BreadcrumbItem)
```

**Usage**

```go
@ui.Breadcrumbs([]ui.BreadcrumbItem{
    {Label: "Home", Href: "/"},
    {Label: "Products", Href: "/products"},
    {Label: "Details", Href: ""},
})
```

---

## Date & Time

### DatePicker

Date, time, and datetime picker with custom UI.

**API**

```go
templ DatePicker(
    label string,       // Field label
    name string,        // Form field name
    value string,       // Initial value (YYYY-MM-DD, HH:MM, or YYYY-MM-DDTHH:MM)
    inputType string,   // "date", "time", or "datetime-local"
    required bool,      // Is field required
    attrs templ.Attributes,
)
```

**Usage**

```go
// Date picker only
@ui.DatePicker("Deadline", "deadline", "2025-10-15", "date", true, templ.Attributes{})

// Time picker only
@ui.DatePicker("Meeting Time", "meeting_time", "14:30", "time", false, templ.Attributes{})

// Date and time picker
@ui.DatePicker("Appointment", "appointment", "2025-10-15T14:30", "datetime-local", true, templ.Attributes{})
```

**Form Submission**

The DatePicker properly submits values with forms (including HTMX):

```go
<form hx-post="/tasks">
    @ui.DatePicker("Deadline", "deadline", "", "date", true, templ.Attributes{})
    @ui.Button("Submit", ui.ButtonPrimary, templ.Attributes{"type": "submit"})
</form>
```

---

### Countdown

Live countdown timer with days, hours, minutes, and seconds.

**API**

```go
templ Countdown(
    targetDate string, // ISO 8601 format: "2025-12-31T23:59:59"
    size string,       // "sm", "md", "lg"
)
```

**Usage**

```go
// Large countdown
@ui.Countdown("2025-12-31T23:59:59", "lg")

// Medium countdown
@ui.Countdown("2025-06-15T12:00:00", "md")

// Small countdown
@ui.Countdown("2025-03-01T09:00:00", "sm")
```

---

### Slider

Range slider component.

**API**

```go
templ Slider(
    label string,
    name string,
    min int,
    max int,
    value int,
    step int,
    attrs templ.Attributes,
)
```

**Usage**

```go
@ui.Slider("Volume", "volume", 0, 100, 50, 1, templ.Attributes{})
```

---

## Utilities

### DarkModeToggle

Dark mode toggle switch.

**API**

```go
templ DarkModeToggle()
```

**Usage**

```go
@ui.DarkModeToggle()
```

**Note:** Requires dark mode script in HTML (see README.md)

---

### Kbd

Keyboard shortcut display component.

**API**

```go
templ Kbd(key string)
templ KbdCombo(keys []string)
templ KbdHelp(description string, keys []string)
```

**Usage**

```go
// Single key
@ui.Kbd("Ctrl")
@ui.Kbd("Enter")

// Key combination
@ui.KbdCombo([]string{"Ctrl", "Shift", "P"})

// With description
@ui.KbdHelp("Save", []string{"Ctrl", "S"})
@ui.KbdHelp("Copy", []string{"Ctrl", "C"})
```

---

### FAB

Floating Action Button and Speed Dial.

**API**

```go
templ FAB(icon templ.Component, attrs templ.Attributes)

type SpeedDialAction struct {
    Label string
    Icon  templ.Component
    Attrs templ.Attributes
}

templ SpeedDial(actions []SpeedDialAction)
```

**Usage**

```go
// Simple FAB
@ui.FAB(ui.IconPlus(), templ.Attributes{
    "@click": "alert('Add new')",
})

// Speed Dial with multiple actions
@ui.SpeedDial([]ui.SpeedDialAction{
    {
        Label: "Create Post",
        Icon: ui.IconPlus(),
        Attrs: templ.Attributes{"@click": "createPost()"},
    },
    {
        Label: "Upload File",
        Icon: ui.IconFolder(),
        Attrs: templ.Attributes{"@click": "uploadFile()"},
    },
})
```

---

### Icons

Common icon set.

**Available Icons**

```go
templ IconUsers(color string)
templ IconCheck(color string)
templ IconCurrency(color string)
templ IconTrending(color string)
templ IconPlus()
templ IconChevronDown()
templ IconDotsVertical()
templ IconCheckCircle()
templ IconClock()
templ IconBell()
templ IconMenu()
templ IconHome()
templ IconChart()
templ IconFolder()
templ IconSettings()
```

**Usage**

```go
@ui.IconHome()
@ui.IconUsers("indigo")
@ui.IconCheck("green")
```

---

## Common Patterns

### Form with Validation

```go
templ SignupForm() {
    <form class="space-y-4">
        @ui.FormGroup("Email", true, "",
            ui.TextInput("Email", "email", "email", "Enter your email", true, templ.Attributes{}))

        @ui.FormGroup("Password", true, "",
            ui.TextInput("Password", "password", "password", "Enter password", true, templ.Attributes{}))

        @ui.Checkbox("I accept the terms", "terms", false, templ.Attributes{})

        @ui.Button("Sign Up", ui.ButtonPrimary, templ.Attributes{"type": "submit"})
    </form>
}
```

### Dashboard Layout

```go
templ Dashboard() {
    @ui.Sidebar(Header(), Navigation(), Footer())
    @ui.Navbar(Logo(), NavItems())

    <main class="lg:ml-64 pt-16 p-6">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
            @ui.Card(ui.CardConfig{
                Variant: "stat",
                StatLabel: "Total Users",
                StatValue: "1,234",
                StatIcon: ui.IconUsers("indigo"),
                StatColor: "indigo",
            }, nil)

            @ui.Card(ui.CardConfig{
                Variant: "stat",
                StatLabel: "Revenue",
                StatValue: "$12,345",
                StatIcon: ui.IconCurrency("green"),
                StatColor: "green",
            }, nil)
        </div>

        @ui.Card(ui.CardConfig{Title: "Recent Activity"}, ActivityTable())
    </main>
}
```

### Data Table with Actions

```go
templ UsersTable(users []User) {
    @ui.Card(ui.CardConfig{Title: "Users"}, UsersContent(users))
}

templ UsersContent(users []User) {
    @ui.Table([]string{"Name", "Email", "Role", "Status", "Actions"}, UserRows(users))

    <div class="px-6 py-4">
        @ui.Pagination(ui.PaginationConfig{
            CurrentPage: 1,
            TotalPages: 10,
            TotalItems: 100,
            PerPage: 10,
            BasePath: "/users",
        })
    </div>
}

templ UserRows(users []User) {
    for _, user := range users {
        <tr>
            <td class="px-6 py-4">{ user.Name }</td>
            <td class="px-6 py-4">{ user.Email }</td>
            <td class="px-6 py-4">{ user.Role }</td>
            <td class="px-6 py-4">@ui.Badge(user.Status, "green")</td>
            <td class="px-6 py-4">
                @ui.ActionMenu([]ui.ActionMenuItem{
                    {Label: "Edit", Href: fmt.Sprintf("/users/%d/edit", user.ID)},
                    {Label: "View", Href: fmt.Sprintf("/users/%d", user.ID)},
                    {Label: "Delete", Href: fmt.Sprintf("/users/%d/delete", user.ID), Danger: true, Divider: true},
                })
            </td>
        </tr>
    }
}
```

---

## Tips & Best Practices

### Preventing Flash (x-cloak)

Always include this CSS in your stylesheet:

```css
[x-cloak] {
    display: none !important;
}
```

Components that use `x-cloak`: Modal, Drawer, Dropdown, ActionMenu, Tabs, Accordion, Carousel, Tooltip, Select

### Performance

For tables with many rows:
- Use server-side pagination
- Implement virtual scrolling for 100+ rows
- Load action menus on demand

### Dark Mode

All components support dark mode out of the box. Just add the dark mode script and toggle component.

### HTMX Integration

Components work seamlessly with HTMX:

```go
<div
    hx-get="/users"
    hx-target="#user-list"
    hx-swap="innerHTML"
    hx-trigger="load"
>
    @ui.Loading("md", "indigo")
</div>
```

### Responsive Design

Most components are mobile-first and responsive by default. The Sidebar and Table components have special mobile optimizations.
