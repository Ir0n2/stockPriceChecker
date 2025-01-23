This script scrapes google finance for stock prices (and crypto!) and emails you if it's dropped to a price where you'd want to buy.

My idea is that the price usually goes back up, especially if it's a pretty essential company like texas instruments, or lockheed martin.

you basically got a getPrice() function that takes 4 parameters. 

stock symbol, seperator, the exchange, and the price you want to buy it for.

the stock symbol, is goog. For cyrpto it's XMR or BTC or whatever

the seperators are different for whether it's crypto or stocks. : for stock. - for crypto.

the exchange is usually NASDAQ but for crypto you put the currency you want to compare to, such as CAD or USD

this is how it looks. 

        getPrice("GOOG", ":", "NASDAQ", 200.0)
        getPrice("AAPL", ":", "NASDAQ", 140.7)
        getPrice("XMR", "-", "CAD", 100.0)
