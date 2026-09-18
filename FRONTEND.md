
[![🇺🇸 English](https://img.shields.io/badge/🇺🇸-English-orange?style=flat&labelColor=f5f5f5)](FRONTEND.md)
[![🇧🇷 Português](https://img.shields.io/badge/🇧🇷-Português-green?style=flat&labelColor=f5f5f5)](FRONTEND.pt.md)
[![🇪🇸 Español](https://img.shields.io/badge/🇪🇸-Español-red?style=flat&labelColor=f5f5f5)](FRONTEND.es.md)

# minimals_architecture — Frontend

[![→ PROPOSAL](https://img.shields.io/badge/→_PROPOSAL-6f42c1?style=for-the-badge)](README.md)
[![→ FRONTEND](https://img.shields.io/badge/→_FRONTEND-02569B?style=for-the-badge)](FRONTEND.md)
[![→ BACKEND](https://img.shields.io/badge/→_BACKEND-0F766E?style=for-the-badge)](BACKEND.md)
[![→ EMBEDDED](https://img.shields.io/badge/→_EMBEDDED-D97706?style=for-the-badge)](EMBEDDED.md)

### Frontend Generic Proposal

```text
src/
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
│   └── utils/
│       ├── formatters/
│       ├── .../
│       └── masks/
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
│       ├── db/
│       ├── .../
│       └── api/
│   ├── repositories/
│   └── services/
│
├── features/
│   ├── login/
│   │   ├── components/
│   │       └── login_form
│   │   ├── repository
│   │   ├── controller
│   │   └── screen
│   │
│   └── components/
│       └── app_button
│
├── routes/
│   ├── router.dart
│   ├── routes.dart
│   └── middlewares/
│
└── main.dart
```

### Flutter Example Proposal
```text
lib/
│
├── core/
│   ├── config/
│       ├── environment.dart
│       └── app_config.dart
│   ├── constants/
│       └── keys.dart
│   ├── errors/
│       ├── app_exceptions.dart
│       └── error_handler.dart
│   └── theme/
│       ├── app_colors.dart
│       └── app_theme.dart
│
├── common/
│   ├── extensions/
│       └── is_valid_email.dart
│   ├── functions/
│       └── debounce.dart
│   └── utils/
│       ├── formatters/
│           └── phone_format.dart
│       ├── .../
│       └── masks/
│           └── phone_mask.dart
│
├── domain/
│   ├── models/
│       └── user_model.dart
│   ├── enums/
│       └── user_role.dart
│   ├── types/
│       └── user_contact.dart
│   └── usecases/
│       ├── ...
│       └── auth_usecase.dart
│
├── data/
│   ├── dto/
│       └── auth_response_dto.dart
│   ├── providers/
│       ├── db/
│       ├── storage/
│           └── auth_storage.dart
│       └── api/
│           ├── user_api.dart
│           ├── .../
│           └── auth_api.dart
│   ├── repositories/
│       ├── user_repository.dart
│       └── auth_repository.dart
│   └── services/
│       └── auth_service.dart
│
├── features/
│   ├── login/
│   │   ├── widgets/
│   │       └── login_form.dart
│   │   ├── repository.dart
│   │   ├── controller.dart
│   │   └── page.dart
│   │
│   ├── home/
│   │   ├── widgets/
│   │   ├── controller.dart
│   │   ├── repository.dart
│   │   └── page.dart
│   │
│   └── widgets/
│       └── app_button.dart
│
├── routes/
│   ├── router.dart
│   ├── routes.dart
│   └── middlewares/
│       └── authenticated.dart
│
└── main.dart
```
