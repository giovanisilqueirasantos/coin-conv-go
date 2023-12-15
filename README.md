# coin-conv-go
a simple backend developed in golang to convert one currency into another. Application developed for study purposes and mainly to pass interviews.

## mysql commit to up the database:
docker run --detach --name=coinconvgodb --env="MYSQL_ROOT_PASSWORD=rootpass" --env="MYSQL_PASSWORD=password" --env="MYSQL_USER=user" --env="MYSQL_DATABASE=coinconvgo" --publish 3306:3306 --volume=$(pwd)/init.sql:/docker-entrypoint-initdb.d/init.sql mysql:5.7

## route of the aplication

/exchange/:amount/:from/:to/:rate

## possible currencies to exchange

real
dollar
euro
btc

## example of request
/exchange/10/real/dollar/4.5

```json
{
	"simboloMoeda": "$",
	"valorConvertido": 45
}
```
