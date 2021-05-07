# Siem Data Producer

This utility to produce data on any tcp or udp destinations. built in golang

<b><u>Build Status: <br><br><img style="align=center" src="https://gitlab.com/yjagdale/siem-data-producer/badges/master/pipeline.svg"/>

# Env variable support

| Variable | Default Value | Possible Values |
| :---: | :---: | :---: | 
| LOG_LEVEL | <b>INFO | <b>DEBUG</b> - To Run application in debug logging mode <br> ERROR - To Run application in Error logging mode |
| DB_PATH | <b>/home/app <br> ./ | inside Container: Default path will be /home/app<br> On System it will be present working dir|

## Deployment

___

### Docker based deployment

___

```
docker run --rm --name siem-data-producer registry.gitlab.com/yjagdale/siem-data-producer:latest
```

With Persistence:

```
docker run --rm --name siem-data-producer -v /storage:/storage -e DB_PATH=/storage registry.gitlab.com/yjagdale/siem-data-producer:latest
```

In Debug Mode:

```
docker run --rm --name siem-data-producer -e LOG_LEVEL=debug -v /storage:/storage -e DB_PATH=/storage registry.gitlab.com/yjagdale/siem-data-producer:latest
```

### Api Doc

___
#### Configuration
Get:
```langurage:shell
curl --location --request GET 'http://localhost:8080/configuration/'
```