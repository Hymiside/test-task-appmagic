# Тестовое задание на Junior Go Developer

## Как запустить сервис
1. Склонируйте репозиторий
2. Установите зависимости
```
~$ go mod vendor
```
3. Установите переменные окружения в файле `.env`
4. Запустите файл `cmd/main.go`

## Методы
1. GasPerMonth - возвращает сколько было потрачено gas помесячно
```
TYPE: GET
URL: /api/gas-per-month
```

2. PricePerDay - возвращает среднюю цену gas за день
```
TYPE: GET
URL: /api/price-per-day
```

3. HourlyPrice - возвращает частотное распределение цены по часам
```
TYPE: GET
URL: /api/hourly-price
```

4. SumAllPeriod - возвращает сколько заплатили за весь период
```
TYPE: GET
URL: /api/sum-all-period
```
