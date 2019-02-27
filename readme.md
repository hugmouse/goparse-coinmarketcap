# crypto.go

Скрипт, который генерирует CSV таблицу, а так же JSON с данных, полученных от сайта coinmarketcap.com

# Формат
Скрипт генерирует CSV со следующими данными:

| ID | Name    | Symbol | Market capacity(USD) | Market capacity(BTC) | Price (USD)   | Price (BTC) | Circulating Supply (USD) | Volume (USD)  | Volume (BTC)  | Change (1h) | Change (24h) | Change (7d) |
|----|---------|--------|----------------------|----------------------|---------------|-------------|--------------------------|---------------|---------------|-------------|--------------|-------------|
| 1  | Bitcoin | BTC    | 67517711508.3        | 17557625.0           | 3845.49228659 | 1.0         | 17557625.0               | 9931066155.01 | 2582891.72508 | 0.215428    | -1.03931     | 0.0473774   |

# Использование
Скачать под подходящую платформу исполняемый файл из [/executables/](https://github.com/hugmouse/Rustling-on-Golang/tree/master/coinmarketcap-parser/executables) и запустить.

Ниже пример с `wget`:

```bash
# Данный пример скачивает версию для linux, с 64-битной архитектурой: 

wget https://github.com/hugmouse/Rustling-on-Golang/raw/master/coinmarketcap-parser/executables/crypto.go-linux-amd64

# Делаем файл исполняемым:
chmod +x crypto.go-linux-amd64

# Запускаем:
./crypto.go-linux-amd64

# Output example: 
# Amount of cryptocurrencies: 2090
# Scraping finished, check file "cryptocoinmarketcap.csv" and "cryptocoinmarketcap.json" for results
```
# Todo
- Сделать нормальную поддержку вывода JSON (в данный момент не поддерживается многомерный JSON)
- Отрефакторить код
