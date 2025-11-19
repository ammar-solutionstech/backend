-# API Coverage Inventory
+
+Generated on 2025-11-13 from the ITaaS `public` schema and updated after implementing the latest controllers and routes.
+
+## Legend
+- `Covered` lists controllers/services/routes already implemented.
+- `Gaps` highlights remaining work (none when marked `—`).
+
+## Tables
+
+### Brand
+- Columns: `id (integer, not null)`, `name (character, not null)`
+- Covered: Generic CRUD via `controllers/resource_factories.go` and `routes/inventory.go` (`/api/brands`)
+- Gaps: —
+
+### City
+- Columns: `id`, `name`, `country_id`
+- Covered: Generic CRUD via `/api/cities`
+- Gaps: —
+
+### Contact
+- Columns: `id`, `name`, `country_id`, `mobile_number`, `phone_number`, `website`, `email`
+- Covered: Generic CRUD via `/api/contacts`
+- Gaps: —
+
+### Country
+- Columns: `id`, `name`, `code`
+- Covered: Generic CRUD via `/api/countries`; linked to users through `models.Country`
+- Gaps: —
+
+### Documents
+- Columns: `id`, `name`, `document_size`, `picture`, `document_type`, `equipment_id`, `supplier_id`, `help_desk_id`
+- Covered: Generic CRUD via `/api/documents`; binary payload accepted as base64
+- Gaps: —
+
+### Equipment
+- Columns: `id`, `name`, `description`, `state`, `serial_number`, `equipment_type_id`, `model_id`, `production_date`, `ip_v4`, `ip_v6`, `mac_address`, `operating_system_id`, `location_id`, `proccessors`, `country_of_region`, `user_id`, `warrantly_start_date`, `warrantly_end_date`, `supplier_id`
+- Covered: Generic CRUD via `/api/equipment`
+- Gaps: —
+
+### Equipment_Software
+- Columns: `software_id`, `equipment_id`, `software_version`, `software_size`, `license`
+- Covered: Relationship endpoints under `/api/equipment/{equipmentId}/software` (controllers `equipment_relationships_controller.go`)
+- Gaps: —
+
+### Equipment_User_History
+- Columns: `equipment_id`, `user_id`, `start_date`, `end_date`, `comment`
+- Covered: `/api/equipment/{equipmentId}/user-history` endpoints with history validation
+- Gaps: —
+
+### Equipment_type
+- Columns: `id`, `name`, `decsription`
+- Covered: Generic CRUD via `/api/equipment-types`
+- Gaps: —
+
+### Help_desk
+- Columns: `id`, `name`, `create_date`, `help_desk_type_id`, `portal_user_id`, `description`, `state`, `resolve_date`, `help_desk_id`, `project_id`
+- Covered: Generic CRUD via `/api/help-desk`
+- Gaps: —
+
+### Help_desk_rating
+- Columns: `id`, `name`, `rate`, `help_desk_id`
+- Covered: Generic CRUD via `/api/help-desk/ratings`
+- Gaps: —
+
+### Help_desk_team
+- Columns: `id`, `name`, `user_manager_id`, `description`, `is_active`, `created_at`, `updated_at`, `deleted_at`
+- Covered: Generic CRUD via `/api/help-desk/teams` plus member management `/api/help-desk/teams/{teamId}/members`
+- Gaps: —
+
+### Help_desk_transaction
+- Columns: `id`, `name`, `transaction_type_id`, `date_time`, `time`, `help_desk_id`
+- Covered: Generic CRUD via `/api/help-desk/transactions`
+- Gaps: —
+
+### Help_desk_transaction_user
+- Columns: `user_id`, `transaction_id`
+- Covered: Relationship endpoints `/api/help-desk/transactions/{transactionId}/users`
+- Gaps: —
+
+### Help_desk_type
+- Columns: `id`, `name`, `description`, `team_id`, `is_active`, `created_at`, `updated_at`, `deleted_at`
+- Covered: Dedicated controller + service (`controllers/help_desk_type_controller.go`)
+- Gaps: —
+
+### Location
+- Columns: `id`, `name`, `zip_code`, `state`, `building_number`, `room_number`, `latitude`, `longitude`, `country_id`, `city_id`, `location_map`
+- Covered: Generic CRUD via `/api/locations`
+- Gaps: —
+
+### Maintenance
+- Columns: `id`, `name`, `start_date`, `end_date`, `contact_id`
+- Covered: Generic CRUD via `/api/maintenance`
+- Gaps: —
+
+### Menu
+- Columns: `id`, `name`, `title`, `uri`
+- Covered: Generic CRUD via `/api/menus` plus menu-role management `/api/menus/{menuId}/roles`
+- Gaps: —
+
+### Model
+- Columns: `id`, `name`, `brand_id`
+- Covered: Generic CRUD via `/api/models`
+- Gaps: —
+
+### Operating_System
+- Columns: `id`, `name`, `description`, `version`, `architectures`
+- Covered: Generic CRUD via `/api/operating-systems`
+- Gaps: —
+
+### Permission
+- Columns: `id`, `name`, `module`
+- Covered: Legacy controller (`controllers/permission_controller.go`)
+- Gaps: Align with generic service pattern later
+
+### Role
+- Columns: `id`, `name`, `description`
+- Covered: Legacy controller (`controllers/role_controller.go`) + permission management endpoints
+- Gaps: Align with generic service pattern later
+
+### Software
+- Columns: `id`, `software_name`, `category_id`, `license_exp_date`
+- Covered: Generic CRUD via `/api/software`
+- Gaps: —
+
+### Software_category
+- Columns: `id`, `name`
+- Covered: Generic CRUD via `/api/software-categories`
+- Gaps: —
+
+### Transaction_type
+- Columns: `id`, `name`, `expected_time`
+- Covered: Generic CRUD via `/api/help-desk/transaction-types`
+- Gaps: —
+
+### User
+- Columns: `id`, `first_name`, `latest_name`, `father_name`, `nationality_id`, `dob`, `work_email`, `private_email`, `password`, `work_mobile`, `private_mobile`, `id_number`, `id_type`, `contact_id`, `active`
+- Covered: Legacy `controllers/user_controller.go` including role/permission helpers
+- Gaps: Potential future refactor onto generic/service layer
+
+### equipment_help_desk
+- Columns: `equipment_id`, `help_desk_id`
+- Covered: `/api/equipment/{equipmentId}/help-desk` endpoints (link/unlink)
+- Gaps: —
+
+### menu_role
+- Columns: `menu_id`, `role_id`
+- Covered: `/api/menus/{menuId}/roles` management
+- Gaps: —
+
+### role_permission
+- Columns: `role_id`, `permission_id`
+- Covered: Role controller add/remove endpoints
+- Gaps: —
+
+### team_members
+- Columns: `user_id`, `help_desk_team_id`
+- Covered: `/api/help-desk/teams/{teamId}/members`
+- Gaps: —
+
+### user_help_desk
+- Columns: `user_id`, `help_desk_id`
+- Covered: `/api/help-desk/{helpDeskId}/participants`
+- Gaps: —
+
+### user_role
+- Columns: `user_id`, `role_id`
+- Covered: User controller role assignment endpoints
+- Gaps: —
+
+### user_special_permission
+- Columns: `user_id`, `permission_id`
+- Covered: User controller permission assignment endpoints
+- Gaps: —
+
+## Notes
+- Generic CRUD is delivered by `controllers/resource_controller.go` and `services/generic_service.go`, with concrete registrations defined in the `routes` package.
+- Relationship endpoints rely on dedicated services (`services/equipment_service.go`, `services/helpdesk_service.go`, `services/navigation_service.go`) to keep database writes centralized.
+- Legacy controllers (User, Role, Permission) continue to operate but can be migrated to the generic stack later for consistency.
