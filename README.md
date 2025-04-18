# GrantMaster

![GrantMaster Logo](assets/logo.png)

**GrantMaster** is a lightweight, Go-based tool that reads structured YAML/JSON policies and dynamically grants the appropriate database permissions in any database. Designed for efficiency, it leverages existing modules to minimize complexity while ensuring precise and secure access control.

## Features

- ✅ Parses YAML/JSON to define user roles and permissions  
- ✅ Grants, revokes, and manages database privileges dynamically  
- ✅ Supports integration with multiple database systems  
- ✅ Uses existing Go modules for efficiency and reliability  
- ✅ Provides logging and auditability for access management  

## GitOps Integration

**GrantMaster** can be easily integrated into **GitOps workflows**, allowing access policies to be managed declaratively through version-controlled repositories. This enables:  

- 🔹 Automated and consistent database access provisioning  
- 🔹 Auditability and traceability of access changes  
- 🔹 Infrastructure-as-code (IaC) best practices for security and compliance  

## Getting Started

### Installation

```sh
go install github.com/welibekov/grantmaster@latest
```

### Build

```sh
make
```

### Run tests

```sh
make runtest
```

Or for specific postgres version

```sh
POSTGRES_DOCKER_IMAGE=postgres:14 make runtest
```

### Example of policy struct

```yaml
- username: david.gilmour
  roles:
    - write_all
- username: jimi.hendrix
  roles:
    - read_all
    - write_all
```

### Then apply policy like
```sh
gm fakegres policy apply policy.yaml
```

### Example of role struct

```yaml
- name: song_write
  schemas:
    - schema: song
      grants:
        - usage
        - select
    - schema: lyrics
      grants:
        - usage
        - select

- name: song_read
  schemas:
    - schema: song
      grants:
        - usage
        - create
```

### Then apply role like
```sh
gm fakegres role apply role.yaml
```


### Supported postgres versions
```
Postgres:9
Postgres:10
Postgres:11
Postgres:12
Postgres:13
Postgres:14
Postgres:15
Postgres:16
Postgres:17
```
