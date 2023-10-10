## ⚡️ Quick start

- Rename `.env.example` to `.env`
- Run project by this command:
  ```bash
  # setup pg db
  $ docker-compose up --detach

  # build css and go and run it
  $ make run
  ```

## 🗄 Project structure

- `/cmd` codebase entrypoint to init server, jobs, etc

- `/internal` core codebase and business logic
  -  `/app` composing internal apps
  - `/locator` composing services and/or repositories
  - `/modules` business/domain logic
  - `/route` endpoint routers
  - `/server` composing server logic, http, ws etc...
  - `/templates` html templates for rendering
- `/pkg` shared packages that can be used by multiple apps
  - `/config` app configuration and env reading
  - `/db` all db connections, eg. `pg`, `redis`, etc...