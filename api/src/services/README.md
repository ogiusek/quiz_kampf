services should have file structure:
- services/$service.go
this file is for access.
it allows to globaly access service.
remember to have way to mock and configure service.

- services/$service/**.go
this is optinal for service implementation if not using one

- services/($service)_test/**.go
this is optional for tests.
if you app has more than one file in services/$service please dont put there tests.
