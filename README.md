# Awesome auth

This project is part of [awesome-project]().

This is the service responsible for authentication users and acts as an SSO.

### Requirements
**Taskfile**: an alternative and more modern tool to Makefile
To install it:
```shell
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin
```

### Improvements
- Cron job for revoking token existed more than 24 hours
- Testing
- Rate limiting
- Service level monitoring
