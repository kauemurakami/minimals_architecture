[![🇺🇸 English](https://img.shields.io/badge/🇺🇸-English-orange?style=flat&labelColor=f5f5f5)](README.md)
[![🇧🇷 Português](https://img.shields.io/badge/🇧🇷-Português-green?style=flat&labelColor=f5f5f5)](README.pt.md)
[![🇪🇸 Español](https://img.shields.io/badge/🇪🇸-Español-red?style=flat&labelColor=f5f5f5)](README.es.md)

# minimals_architecture

[![→ PROPOSAL](https://img.shields.io/badge/→_PROPOSAL-6f42c1?style=for-the-badge)](README.md)
[![→ FRONTEND](https://img.shields.io/badge/→_FRONTEND-02569B?style=for-the-badge)](FRONTEND.md)
[![→ BACKEND](https://img.shields.io/badge/→_BACKEND-0F766E?style=for-the-badge)](BACKEND.md)
[![→ EMBEDDED](https://img.shields.io/badge/→_EMBEDDED-D97706?style=for-the-badge)](EMBEDDED.md)

> **Clean principles. Minimal ceremony. Clear responsibilities.**

Minimal Architecture is a responsibility-driven approach to organizing software around **Clean Architecture principles**, while keeping the implementation practical, readable, and adaptable to different technologies.

The idea came from **GetX Pattern**, which solved a real organization problem for a period of time in the Flutter community: developers wanted a predictable way to structure their applications without having to reinvent the same project organization every time.

That experience led to a broader question:

> **What if the same idea could be tested across different platforms, runtimes, and application structures while staying aligned with Clean Architecture?**

That experiment became **Minimal Architecture**.

It is not a framework architecture, and it is not tied to Flutter, GetX, Dart, Node.js, Go, or any specific technology.

It is a proposal for making Clean Architecture **easier to implement consistently**.

---

## The Problem

Clean Architecture is often explained through diagrams, layers, dependency rules, and abstractions.

The principles are valuable.

The problem usually starts when developers try to turn those principles directly into folders and files.

Different projects end up with structures such as:

```text
controllers/
services/
repositories/
helpers/
utils/
managers/
providers/
usecases/
models/
```

The names may look organized, but the actual responsibilities can still be mixed.

A controller may contain business rules.

A service may access a database.

A repository may contain presentation logic.

A helper may actually be an authentication service.

A DTO may construct domain objects directly.

A feature may contain several unrelated responsibilities simply because they did not have an obvious place to go.

This is where architecture becomes difficult to maintain.

> **The problem is not that a file has a responsibility. The problem is that responsibilities that should be separated are often placed together in the same file, folder, or module.**

Minimal Architecture tries to make those boundaries **visible and intuitive**.

---

# What Minimal Architecture Is

Minimal Architecture is a **structure for standardizing the implementation of Clean Architecture**.

The folder organization does **not** define Clean Architecture.

Clean Architecture is defined by principles such as:

- separation of concerns
- dependency direction
- business logic independence
- controlled coupling
- explicit boundaries
- replaceable infrastructure
- testability

The folders simply make those principles easier to see and follow.

Think of the structure as a **map**.

It does not create the architecture.

It makes the architecture **legible**.

---

# The Pillars

| Pillar | Idea |
| --- | --- |
| **Responsibility** | Every piece of code should have a clear reason to exist and an obvious owner. |
| **Separation** | Different responsibilities should not be mixed just because they are related. |
| **Boundaries** | Dependencies should cross explicit boundaries instead of spreading through the application. |
| **Minimalism** | Do not create abstractions, layers, or files without a real responsibility behind them. |
| **Predictability** | Similar responsibilities should have predictable places across projects and technologies. |
| **Adaptability** | The concepts remain stable while implementation details can change between Flutter, Go, Node.js, embedded systems, and other environments. |

---

# The Structure

A generic frontend-oriented example can look like this:

```text
lib/
│
├── core/
│   ├── config/
│   ├── constants/
│   ├── errors/
│   └── theme/
│
├── common/
│   ├── extensions/
│   ├── functions/
│   ├── masks/
│   └── utils/
│
├── domain/
│   ├── models/
│   ├── enums/
│   ├── types/
│   └── usecases/
│
├── data/
│   ├── dto/
│   ├── providers/
│   ├── repositories/
│   └── services/
│
├── features/
│   ├── feature_a/
│   │   ├── components/
│   │   ├── repository/
│   │   ├── controller/
│   │   └── screen/
│   │
│   └── components/
│
├── routes/
│   ├── router.dart
│   ├── routes.dart
│   └── middlewares/
│
└── main.dart
```

**This is a reference structure, not a mandatory checklist.**

A project does not become better by having every folder.

If a responsibility does not exist, the corresponding folder or file does not need to exist.

For example:

```text
features/
└── profile/
    ├── controller/
    └── page/
```

can be completely valid.

There is no reason to create a repository, use case, service, DTO, and provider only because the architecture diagram contains those concepts.

---

# 1. `core/`

`core` contains concerns that belong to the application as a whole, especially application-wide infrastructure and configuration.

```text
core/
├── config/
├── constants/
├── errors/
└── theme/
```

### `config/`

Application configuration.

Examples:

```text
environment.dart
api_config.dart
database_config.dart
app_config.dart
```

Typical responsibilities:

- environment configuration
- API configuration
- database configuration
- application-wide configuration

### `constants/`

Application-wide constants.

Examples:

```text
storage_keys.dart
api_constants.dart
app_constants.dart
```

Constants should remain constants. This should not become a generic dumping ground.

### `errors/`

Application-level error definitions and error handling infrastructure.

Examples:

```text
app_exception.dart
network_exception.dart
validation_exception.dart
error_handler.dart
```

### `theme/`

Frontend-specific application appearance.

Examples:

```text
app_theme.dart
color_scheme.dart
text_theme.dart
```

A backend does not need `theme/`.

That is intentional.

> **Minimal Architecture standardizes responsibilities, not identical folders across every technology.**

---

# 2. `common/`

`common` contains functionality that is genuinely reusable across unrelated parts of the application.

```text
common/
├── extensions/
├── functions/
├── masks/
└── utils/
```

The important word is **genuinely**.

`common` should not become:

```text
things-we-did-not-know-where-to-put/
```

If something belongs specifically to authentication, users, payments, or another feature, it should remain close to that responsibility.

### `extensions/`

Generic language or framework extensions.

Example:

```dart
extension StringExtensions on String {
  bool get isNotEmptyOrWhitespace => trim().isNotEmpty;
}
```

### `functions/`

Small generic functions that do not represent a business operation.

Examples:

```text
format_date.dart
debounce.dart
parse_currency.dart
```

### `masks/`

Reusable frontend input masks.

Examples:

```text
phone_mask.dart
document_mask.dart
currency_mask.dart
```

### `utils/`

Generic utilities.

Examples:

```text
logger.dart
date_utils.dart
response_utils.dart
```

If a utility starts knowing too much about a feature, it probably does not belong here.

---

# 3. `domain/`

`domain` contains the application's business concepts.

It should describe **what the application is about**, without depending unnecessarily on HTTP, databases, UI frameworks, or external representations.

```text
domain/
├── models/
├── enums/
├── types/
└── usecases/
```

### `models/`

Domain models.

Example:

```dart
class User {
  final String id;
  final String name;

  const User({
    required this.id,
    required this.name,
  });
}
```

A domain model should represent the application's concept of a user.

It does not have to mirror an API response.

### `enums/`

Closed sets of meaningful domain values.

Examples:

```text
UserRole
OrderStatus
ContactType
PaymentStatus
```

### `types/`

`types/` is for domain concepts that deserve to be represented as **their own type**, especially when a value starts carrying behavior, rules, validation, or multiple representations.

A type does not have to be a simple alias or primitive wrapper.

It can be a class with its own internal value, operations, validation, and representations.

Examples:

```text
UserId
Email
PhoneNumber
Money
Coordinates
```

A good example is `Money`.

Instead of passing a raw `double` around and creating formatting functions everywhere:

```dart
double price = 129.90;
```

the domain can represent the concept itself:

```dart
class Money {
  final double value;

  const Money(this.value);

  Money operator +(Money other) {
    return Money(value + other.value);
  }

  Money operator *(double multiplier) {
    return Money(value * multiplier);
  }

  String get formatted => ...;
}
```

The important part is not the exact implementation. The important part is that the **value and the behavior that belongs to that value stay together**.

The same `Money` value could represent:

```text
129.90
$129.90
US$ 129.90
R$ 129,90
€129.90
```

depending on the context in which it is presented, while internally still maintaining the numeric value required for operations such as:

```text
addition
subtraction
multiplication
comparison
conversion
```

This avoids spreading domain-specific formatting and rules across unrelated helpers and UI files:

```text
formatMoney()
formatDollar()
formatEuro()
formatReal()
parseMoney()
calculatePrice()
```

Instead, the concept is concentrated in the type that actually represents it.

This is particularly useful when a value starts becoming more complex than a primitive:

```text
Primitive
   ↓
double

Recognized concept
   ↓
Money

Value + rules + behavior + representations
   ↓
A domain type
```

Types therefore provide a way to **give a meaningful identity to complex values** without turning every small concept into a large architecture layer.

A type should be introduced when the concept itself has enough meaning or behavior to justify owning its representation.

> **If a value has its own rules, behavior, validation, or representations, consider making the value a type instead of scattering that logic around the application.**

---

# 4. `data/`

`data` deals with external representations, persistence, communication, and infrastructure that provides or stores application data.

```text
data/
├── dto/
├── providers/
├── repositories/
└── services/
```

### `dto/`

DTOs represent external or persistence-oriented data.

A common flow is:

```text
External representation
        ↓
      DTO
        ↓
    Domain Model
```

For example:

```dart
class UserDTO {
  final String id;
  final String name;

  factory UserDTO.fromJson(Map<String, dynamic> json) {
    // external representation
  }

  User toModel() {
    // domain representation
  }
}
```

DTOs are useful when the external representation differs from the domain representation.

They are **not mandatory for every model**.

There is no value in creating a DTO, mapper, parser, and converter when all of them represent exactly the same thing without protecting a meaningful boundary.

### `providers/`

Providers communicate directly with external sources.

Examples:

```text
ApiProvider
DatabaseProvider
StorageProvider
FirebaseProvider
FileProvider
```

The provider answers:

> **How do I communicate with this external source?**

It should not become the place where business rules live.

### `repositories/`

Repositories provide a boundary around data access.

For example:

```text
UserRepository
    ↓
 ┌───────────────┐
 │ API           │
 │ Database      │
 │ Cache         │
 └───────────────┘
```

A repository is useful when it reduces coupling, coordinates sources, or gives the application a meaningful data-access boundary.

It is **optional**.

> **A repository should exist because it solves a problem — not because Clean Architecture diagrams usually contain one.**

### `services/`

Services own infrastructure or stateful responsibilities that are better represented as a service than as a generic helper.

Examples:

```text
AuthService
StorageService
NotificationService
TokenService
AnalyticsService
```

For example, token creation and validation are naturally authentication responsibilities:

```text
AuthService
├── createToken()
└── validateToken()
```

That is much clearer than placing authentication behavior in:

```text
helpers/
auth_helper.dart
```

---

# 5. `features/`

`features` organizes the application by functionality.

Examples:

```text
features/
├── auth/
├── users/
├── posts/
├── profile/
└── checkout/
```

This is where the architecture becomes concrete.

A feature should contain the pieces necessary to implement that capability.

It does **not** need to contain every architectural layer.

---

## Frontend-specific responsibilities

Frontend applications naturally introduce presentation concepts that do not apply to every environment.

A feature may look like:

```text
features/
├── profile/
│   ├── components/
│   │   ├── profile_header.dart
│   │   └── profile_stats.dart
│   ├── repository/
│   ├── controller/
│   └── screen/
│
└── components/
    ├── app_button.dart
    └── app_dialog.dart
```

The important distinction is between **feature components** and **shared components**.

### `features/*/components/`

These are components that belong specifically to one feature.

Examples:

```text
ProfileHeader
ProfileStats
PostComposer
CheckoutSummary
```

A `ProfileHeader` that only makes sense inside the profile feature should stay inside:

```text
features/profile/components/
```

This keeps the feature cohesive and prevents feature-specific UI from becoming a global dependency.

### `features/components/`

These are globally reusable frontend components shared by multiple features or modules.

Examples:

```text
AppButton
AppDialog
AppTextField
AppLoading
AppEmptyState
```

For example:

```text
features/
├── auth/
│   └── components/
│       └── login_form.dart
│
├── profile/
│   └── components/
│       └── profile_header.dart
│
└── components/
    ├── app_button.dart
    ├── app_dialog.dart
    └── app_loading.dart
```

The rule is simple:

> **If the component belongs to a feature, keep it in that feature. If it is genuinely reusable across unrelated features, put it in the shared `features/components/` area.**

This prevents the common frontend problem where every component is placed in a single global folder, even when most of them are actually feature-specific.

It also avoids the opposite problem: duplicating the same reusable component across several modules.

### `screen/`

A screen is a user-facing entry point of a feature.

Examples:

```text
LoginScreen
HomeScreen
ProfileScreen
CheckoutScreen
```

A screen primarily handles presentation and interaction.

It should not become the place for database queries, API implementation, or large business workflows.

The name may vary by framework:

```text
Screen
Page
View
Route
```

The responsibility is what matters.

### `controller/`

Controllers coordinate presentation state and user interactions.

For example:

```text
Screen
 ↓
Controller
 ↓
UseCase
 ↓
Repository
```

The exact implementation depends on the frontend technology.

A controller should not become a dumping ground for API access, persistence, and unrelated business rules.

### `repository/`

A feature-level repository can group data needs specific to that feature.

Example:

```text
ProfileRepository
├── getProfile()
├── updateProfile()
└── getProfileStats()
```

Again: **optional**.

If the feature does not have a meaningful repository responsibility, do not create one just to complete the folder structure.

---

## Backend-specific responsibilities

The same feature concept works in a backend, but the concrete files naturally change.

For example:

```text
features/
└── users/
    ├── controller.go
    ├── routes.go
    └── functions/
        ├── get_user.go
        ├── create_user.go
        └── delete_user.go
```

A backend controller is an HTTP boundary.

A possible flow is:

```text
HTTP
 ↓
Route
 ↓
Controller
 ↓
Function / UseCase
 ↓
Repository
 ↓
Provider
 ↓
Database
```

But a simple endpoint may only need:

```text
Route
 ↓
Controller
 ↓
Function
 ↓
Provider
```

That is still a valid Minimal Architecture implementation.

---

# 6. `routes/`

`routes` defines how external navigation or requests enter the application.

Frontend example:

```text
routes/
├── router.dart
├── routes.dart
└── middlewares/
```

Typical responsibilities:

### `router`

Defines the application's routing mechanism.

### `routes`

Defines route names and route configuration.

### `middlewares`

Defines policies that need to run before entering a route.

Examples:

```text
authenticated.dart
guest.dart
```

A middleware can answer:

> **Is this navigation/request allowed to enter this route?**

It should not become the entire authentication system.

---

# 7. `main.dart`

`main.dart` is the application entry point.

Its job is to initialize and start the application.

For example:

```text
Load configuration
        ↓
Initialize dependencies
        ↓
Initialize services
        ↓
Start application
```

It should not become a dumping ground for application logic.

---

# Clean Architecture + Minimalism

Minimal Architecture follows Clean Architecture principles while intentionally avoiding architecture ceremony.

The goal is not:

```text
More layers
More interfaces
More abstractions
More files
```

The goal is:

```text
Clear boundaries
Clear ownership
Controlled dependencies
Minimal coupling
```

A project with three well-separated responsibilities can be cleaner than a project with twenty abstractions.

The number of folders is not an architectural quality metric.

---

# Responsibility-Driven Decisions

When deciding where code belongs, ask:

| Question | Purpose |
| --- | --- |
| **What does this code actually do?** | Identify its real responsibility. |
| **Who owns that responsibility?** | Find the correct boundary. |
| **Is it feature-specific or generic?** | Decide between `features` and shared areas. |
| **Does it represent business knowledge?** | Consider `domain`. |
| **Does it communicate with an external source?** | Consider `data`. |
| **Does it coordinate a meaningful operation?** | Consider a use case. |
| **Does it reduce data coupling?** | Consider a repository. |
| **Does it represent external data?** | Consider a DTO. |
| **Is it truly generic?** | Consider `common`. |
| **Is it application-wide infrastructure/configuration?** | Consider `core`. |

This prevents the classic architecture trap:

> **"I need somewhere to put this file, so I will create a new layer."**

Instead:

> **"This code has a responsibility. Where should that responsibility live?"**

---

# What Is Optional?

Minimal Architecture deliberately makes several concepts optional.

| Concept | Required? | When it makes sense |
| --- | --- | --- |
| DTO | No | External/persistence representation differs from the domain. |
| Repository | No | A data boundary or source coordination provides real value. |
| Use Case | No | An operation contains meaningful application/business orchestration. |
| Service | No | A distinct infrastructure/state/external responsibility exists. |
| Interface | No | An abstraction solves a real coupling, substitution, or contract problem. |
| Mapper | No | Conversion is complex enough to justify a dedicated boundary. |
| Feature layer | Yes, conceptually | Functionality should be grouped by capability, but its internal files are flexible. |

> **If a layer does not have a responsibility, it should not exist just to make the architecture look complete.**

---

# Architecture Across Technologies

Minimal Architecture is designed to survive changes in technology.

The concepts remain stable while the implementation adapts.

### Flutter

```text
lib/
├── core/
├── common/
├── domain/
├── data/
├── features/
├── routes/
└── main.dart
```

### Go

```text
internal/
├── core/
├── common/
├── domain/
├── data/
├── features/
└── routes/
```

### Node.js

```text
src/
├── core/
├── common/
├── domain/
├── data/
├── features/
└── routes/
```

The goal is not to force every ecosystem to use exactly the same files.

For example:

```text
Frontend
→ Page / Component / Controller

Backend
→ Route / Controller / Handler / Function

Infrastructure
→ Provider / Service
```

The names and mechanics may change.

The responsibility does not.

---

# Why Folder Organization Still Matters

A common argument is:

> "Clean Architecture is about principles, so folders don't matter."

Technically, the principles are more important than the folders.

But in real projects, **readability is an architectural advantage**.

A predictable structure helps developers:

- find code faster
- understand dependencies
- recognize responsibilities
- onboard new contributors
- review changes
- avoid accidental coupling
- keep features cohesive
- resist the temptation to create random shared layers

So folders do not *define* Clean Architecture.

They make the architecture **visible enough to be followed consistently**.

That distinction is fundamental.

---

# The Real Goal

Minimal Architecture is not trying to make every project look identical.

It is trying to make the **reason behind the organization** understandable.

A developer should be able to look at a project and quickly understand:

```text
Where is the business concept?

Where is external data handled?

Where does this feature live?

Where does this request enter the system?

Where is application-wide infrastructure?

Where is truly shared code?

Where does this operation belong?

What should this code NOT know about?
```

When those answers are obvious, the architecture starts working for the team instead of becoming another thing the team has to maintain.

---

# The Minimal Architecture Rule

> **Do not organize code by what files you are expected to have. Organize it by the responsibilities the application actually has.**

And:

> **Do not add abstraction because Clean Architecture has a name for it. Add abstraction when it creates a useful boundary.**

And most importantly:

> **Keep responsibilities separate, keep dependencies controlled, and keep the ceremony as small as possible.**

---

# From GetX Pattern to Minimal Architecture

Minimal Architecture has its roots in the experience of **GetX Pattern**.

GetX Pattern was created around a practical problem in the Flutter community: providing a predictable project structure so developers could stop reinventing the organization of their applications.

That idea was useful enough to raise a bigger question:

> **Could a predictable structure help developers implement Clean Architecture more consistently across different environments?**

Minimal Architecture is the result of exploring that question.

The focus moved from a specific framework to a broader architectural model:

```text
Framework-specific structure
        ↓
Responsibility-driven structure
        ↓
Cross-platform architectural concepts
```

The goal is not to replace Clean Architecture.

It is to make its implementation more **practical, readable, predictable, and minimal**.

---

# In One Sentence

> **Minimal Architecture is a responsibility-driven way to make Clean Architecture easier to implement, easier to read, and harder to accidentally break — without turning architecture into ceremony.**

**Clear responsibilities.  
Predictable structure.  
Minimal ceremony.**
