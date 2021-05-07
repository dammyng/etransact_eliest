#Eliest intergration doc

- Registrations

```shell
    curl --location --request POST 'https://eliestlotto.biz/v1/register' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "msisdn": "080xxxxx123",
        "yob":"xxxx"
    }'
```

- Get details

```shell
    curl --location --request GET 'https://eliestlotto.biz/v1/details/080xxxxx123' \
```

- Fund account

```shell
    curl --location --request POST 'https://eliestlotto.biz/v1/fund' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "msisdn": "080xxxxx123",
        "amount":0.1
    }'
```

---

### Get available games 

```shell
    curl --location --request GET 'https://eliestlotto.biz/v1/games/list' \
    --header 'Content-Type: application/json' \
```

### Play Game

- Would return the result of the game play instantly.

```shell
curl --location --request POST 'https://eliestlotto.biz/v1/play' \
--header 'Content-Type: application/json' \
--data-raw '{
    "msisdn": "080xxxxx123",
    "game_id":"1",
    "guess":"34#5"
}'
```

### Validate winning code

```shell
curl --location --request POST 'https://eliestlotto.biz/v1/wins/validate' \
--header 'Content-Type: application/json' \
--data-raw '{
    "msisdn": "080xxxxx123",
    "code":"9509925"
}'
```

**nb: The validity of the code can only be checked once to process customer transfer to bank request. A second request would return an invalid code error. If the bank deposit should fail call the `transfer failed` endpoint to revet the status of the code. if the Bank deposit is successful call the `transfer success` endpoint.**

### Winning deposit transfer fails callback

```shell
curl --location --request POST 'https://eliestlotto.biz/v1/wins/transfer_failed' \
--header 'Content-Type: application/json' \
--data-raw '{
    "msisdn": "080xxxxx123",
    "code":"9509925"
}'
```

### Winning deposit transfer success callback

```shell
curl --location --request POST 'https://eliestlotto.biz/v1/wins/transfer_success' \
--header 'Content-Type: application/json' \
--data-raw '{
    "msisdn": "080xxxxx123",
    "code":"9509925"
}'
```

### Transfer winning to agent

```shell
curl --location --request POST 'https://eliestlotto.biz/v1/wins/transfer_to_agent' \
--header 'Content-Type: application/json' \
--data-raw '{
    "msisdn": "080xxxxx123",
    "code":"9509925",
    "agent":"23480xxxxxx123"
}'
```

**nb: You do not have to verify the winning code before making this call, verification is done from eliest end.**


### Fund account with voucher
```shell
curl --location --request POST 'https://eliestlotto.biz/v1/fund/voucher' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code":"xxxxxxx",
    "msisdn":"xxxxx"
}'
```