## ⚡️ Quick start

- Rename `.env.example` to `.env`
- Run project by this command:
  ```bash
  make docker.run
  ```

## 🗄 Project structure

### Directories

1. **/cmd** entry points.

2. **/config** has structures which contains service config.

3. **/db** has seeders and method for connecting to the database.

4. **/deploy** contains the container (Docker) package configuration and template(docker-compose) for project deployment.

5. **/development** includes Docker and docker-compose files for setup linter.

6. **/migrations** has files for run migrations.

7. **/models** includes structures describing data models.

8. **/repositories** contains methods for selecting entities from the database.

9. **/requests** has structures describing the parameters of incoming requests, and validator.

10. **/responses** includes structures describing the parameters of outgoing response.

11. **/server** is the main project folder. This folder contains the executable server.go.

12. **/server/builders** contains builders for initializing entities.

13. **/server/handlers** contains request handlers.

14. **/server/routes** has a file for configuring routes.

15. **/services** contains methods for creating entities.

16. **/tests**  includes tests and test data.
