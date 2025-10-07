# Changelog

All notable changes to Lite UI will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.13] - 2025-10-07

### Added
- Added SelectConfig with DefaultValue and OnChange support to Select component
- Added helper functions getInitialLabel() and optionsToJSON() for Select state initialization
- Added SelectWithConfig() function for advanced Select configuration

### Changed
- Removed fade transitions from Tabs component for instant content switching
- Select component now supports initial value selection and custom onChange handlers

### Fixed
- Fixed Select component x-cloak coverage on button span and SVG elements

## [0.1.12] - 2025-01-07

### Fixed
- Fixed Alpine.js flashing across multiple components by adding x-cloak to all dynamic elements
- Fixed Tabs component buttons flashing on page load
- Fixed Dropdown component button flashing on page load
- Fixed DarkMode toggle button and slider flashing on page load
- Fixed Select component options and chevron icon flashing
- Resolved cascade flashing issue where Dropdown at top of page caused ActionMenu at bottom to flicker

## [0.1.11] - 2025-01-07

### Fixed
- Fixed Alpine.js flashing in ActionMenu component by adding x-cloak to trigger button
- Prevents team member list and other ActionMenu instances from flashing on page load

## [1.0.0] - 2025-01-05

### Added
- Initial release of Lite UI
- 24 production-ready components
- Full Tailwind CSS v4 support
- Dark mode support with toggle component
- HTMX integration patterns
- Alpine.js interactivity
- Responsive design with mobile-first approach
- Type-safe Go + Templ implementation

### Components
- Forms: Button, TextInput, TextArea, Select, Checkbox, Switch, FormGroup
- Layout: Card, Sidebar, Navbar, Modal, Drawer, Tabs
- Data Display: Table, Pagination, Badge, Avatar, Timeline, Stepper
- Feedback: Progress, Slider
- Navigation: Dropdown
- Utilities: DarkModeToggle, Icons

### Documentation
- Comprehensive README with usage examples
- Component API documentation
- Quick start guide
- Tailwind configuration guide
- HTMX integration examples
