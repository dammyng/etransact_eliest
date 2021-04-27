#Eliest intergration doc

- Registrations

```shell
    curl --location --request POST 'http://52.178.164.225/v1/register' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "msisdn": "080xxxxx123",
        "yob":"xxxx"
    }'
```

- Get details

```shell
    curl --location --request GET 'http://52.178.164.225/v1/details/080xxxxx123' \
```

- Fund account

```shell
    curl --location --request POST 'http://52.178.164.225/v1/fund' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "msisdn": "080xxxxx123",
        "amount":0.1
    }'
```
