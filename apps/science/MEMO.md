- Buy cross => buy order placed => there's a sell cross before the order is fulfilled => check if there's a buy order => cancel the placed buy order
- Sell cross => sell order placed => there's a buy cross before the order is fulfilled => check if there's a sell order =>  cancel the placed sell order
- Buy cross => buy order placed => there's a buy cross before the order is fulfilled => check if there's a buy order => if there is, ignore the second buy order, if there isn't place buy order
- Sell cross => sell order placed => there's a sell cross before the order is fulfilled => check if there's a sell order => if there is, ignore the second sell order, if there isn't place sell order
- Sell cross => check if we actually have that stock => sell order placed
- API fail => 5 min grace period

- Responsiblity
- science
    - rating and the list
- consumer
    - quantity, budget

- Pullers and consumers
   - sell cross consumer (interested in what positions we have)
   - buy cross consumer (interested in something?)
   - logging consumer (optional)
   - if sold event is the only thing we're interested, only one puller and we can go routine


- buy if
     - there's a cross on the 30 min candle price (run every 30 minutes except the morning)
     - don't sell if there's a dead cross on the same day on the same candle
     - adx > 20
 - Sell if
     - there's a cross on the 1 day candle price (run in the morning / ~ yesterday's stocks)
     - diviation rate?

# MACD
# history = indicators.macd(history, settings['macd']['span'])

# ULT
# history = indicators.ult(history, settings['ult']['span'], settings['ult']['weight'])

# RSI
# history = indicators.rsi(history, settings['rsi']['span'])

# MFI
# history = indicators.mfi(history, settings['mfi']['span'])

# Print
# print(('ticker: {} \n ema_golden_crosses: \n {} \n ema_dead_crosses: \n {}').format(ticker, ema_golden_crosses, ema_dead_crosses))

# Plotting example
# plots.describe_candles_ema_crosses_adx(history, ticker, settings['ema']['span'], ema_golden_crosses, ema_dead_crosses)
