## Assignment 1: show your lasted “structure project” or standard of init project structure and defined convention

My lasted project is an APIs service base on Nest.js framework. let see the brief of my project structure.

```json
root
|-package.json
|-tsconfig.json
|-configuration // contains environment's constant
|-src
|   |-server.ts // root file
|   |-assets
|   |   |- static assets eg: PDF
|   |-auth // auth guard layer
|   |   |-guards
|   |   |   |-auth.guard.ts
|   |   |   |-auth.guard.spec.ts
|   |-base // main logics is here
|   |   |-v1 // defined APIs version
|   |   |   |-module folders
|   |   |   |   |-controller
|   |   |   |   |   |-controller.ts
|   |   |   |   |   |-controller.spec.ts
|   |   |   |   |-dto
|   |   |   |   |   |-dto.ts
|   |   |   |   |-interface
|   |   |   |   |   |-interface.ts
|   |   |   |   |-service
|   |   |   |   |   |-service.ts
|   |   |   |   |   |-service.spec.ts
|   |   |   |   |-module.ts
|   |   |-v2
|   |-common
|   |   |-configuration // global configuration eg: TTL, date format
|   |   |-exception // exception layer handler
|   |   |-intereptor // interceptor layer handler
|   |   |-middleware // middleware layer handler
|   |-entities // DB's entities
|   |-shared // contains helper functions
```

### Common conventions

1. Using kebab-case for file/folder structure
2. Common constants used in services keep within `src/common/configuration` file
3. Common functions keep in `src/shared`
4. Changing Request or Response API's structures which is breaking-change must be a new endpoint with increment version eg: `/v1/users` -> `/v2/users`
5. A service should contains related logic and not too large otherwise separate to a new one. eg: `payment service`, `payment-history service`
6. Mock code must not entangle with production code. eg: implement mocked module and re-routing it from `middleware` layer
