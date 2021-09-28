# Siem Data Producer

This utility to produce data on any tcp or udp destinations. built in golang

<b><u>Build Status: <br><br><img style="align=center" src="https://gitlab.com/yjagdale/siem-data-producer/badges/master/pipeline.svg"/></u>

# API Documentation:
You can get api documentations as swagger.

http://"your server ip":"port"/swagger/index.html
<i><br><b><u><a href="http://localhost:8082/swagger/index.html"> local url </a></u></b></i> 


# Env variable support

| Variable | Default Value | Possible Values |
| :---: | :---: | :---: | 
| LOG_LEVEL | <b>INFO | <b>DEBUG</b> - To Run application in debug logging mode <br> ERROR - To Run application in Error logging mode |
| DB_PATH | <b>/storage/database <br>or<br> ./ | <b><u>Inside Container:</b></u> Default path will be /storage/app<br> <b><u>On System:</b></u> it will be present working dir|

## Deployment

### Docker based deployment

___

```
docker run --rm --name siem-data-producer registry.gitlab.com/yjagdale/siem-data-producer:latest
```

With Persistence:

```
docker run --rm -p 8082:8082 --name siem-data-producer -v /storage:/storage -e DB_PATH=/storage registry.gitlab.com/yjagdale/siem-data-producer:latest
```

In Debug Mode:

```
docker run --rm -d -p 8082:8082 --name siem-data-producer -e LOG_LEVEL=debug -v /storage:/storage -e DB_PATH=/storage registry.gitlab.com/yjagdale/siem-data-producer:latest
```

### Api Doc

___
#### Configuration
Get:
```langurage:shell
curl --location --request GET 'http://localhost:8080/configuration/'
```

Configuration will hold all the override which you want to apply in line while producing.

To Create:
```json
{
  "override_key": "date_format_10",
  "override_values": [
    "FUNCTION::DATE::Jan 2 15:04:05"
  ]
}
```

Override Key: will be the string that will be replaced from the actual log line.
Override_value: value that needs to be substituted in place of override key in actual log line
    if value doest start with `FUNCTION` then value will be substituted as is in actual log line(no manipulation to specified value)
ex. 
    log line: `this is log line with format_1`
    override: 
    ```json
        {
            "override_key": "format_1",
            "override_values": [
                "value1",
                "value2",
            ]
        }
    ```
    output line which goes to destination is 
    ```
    this is log line with value1
    ```
    or
    ```
        this is log line with value2
    ```