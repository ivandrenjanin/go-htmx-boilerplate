## ⚡️ Quick start

- Rename `.env.example` to `.env`
- Run project by this command:
  ```bash
  make docker.run
  ```

## 🗄 Project structure

### /app

**Folder with business logic only**. This directory doesn't care about what database driver we're using.

- `/app/controller` folder for functional controller (used in routes)
- `/app/dto` Data Transfer Objects(DTO) folder for transforming data before it is sent to API clients
- `/app/model` folder for describeing business models and methods
- `/app/repository` folder for performing database operations for models

### /cmd
**Main applications for this project.**

The directory name for each application should match the name of the executable you want to have (e.g., `/cmd/server` `/cmd/cron`).
Don't put a lot of code in the application directory. If you think the code can be imported and used in other projects,
then it should live in the `/pkg` directory.

### /docs

**Folder with API Documentation.**

This directory contains config files for auto-generated API Docs by Swagger, screenshots
and any other documents related to this project.

### /pkg

**Folder with project-specific functionality.** This directory contains all the project-specific code tailored only for the business use case.

- `/pkg/config` folder for configuration functions
- `/pkg/middleware` folder for middlewares (Fiber built-in and ours)
- `/pkg/route` folder for describeing routes of the project
- `/pkg/validator` folder with validation functions

### /platform

**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual project,
like setting up the database, logger instance and storing migrations, seeds(demo data).

- `/platform/database` folder with database setup functions (by default, PostgreSQL)
- `/platform/logger` folder with better logger setup functions (by default, Logrus)
- `/platform/migrations` folder with migration files (used with [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool)
- `/platform/seeds` folder with demo data for application rapid setup. mostly **sql** scripts
