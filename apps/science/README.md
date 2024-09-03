# HobbesAndFriends
*****
## Docker Build

- Build by the following commands
- Docker build's arguments are currently ```kind``` and ```target```. When you're adding arguments, don't forget to add to the following build command as well. 

```
docker build . --build-arg kind='market' --build-arg target='all'
docker image ls
docker run {image ID}

# change the env variable dynamically at runtime
docker run --env kind='market' --env target='nasdaq' {image ID}
```

## Prerequisites

* python 3.8.0

*****

## For day-to-day analysis:

### Usage
- Will return a list of stock(s) that had a EMA buy / sell signal on a designated day (default: previous day).

- By changing the `ema.crosses.buy.days_ago` and `ema.crosses.sell.days_ago` values in the `config.yml`, the signal days can be controled.

### Commands 

#### 1. By Individual Stocks

``` python function.py 'symbol' '["GS"]' ```
``` python function.py 'symbol' '["GS", "APPL"]' ```

####  2. By Markets

``` python function.py 'market' 'sp500' ```
``` python function.py 'market' 'dow' ```
``` python function.py 'market' 'sp500_dow' ```
``` python function.py 'market' 'nasdaq' ```
``` python function.py 'market' 'all' ```

**Note: The last two (`nasdaq` and `all`) will pull and analyze about 5000 ~ 5500 stocks, so it's VERY heavy.**

### Process flow
- Function runs every hour
  #### Sell
    - Every hour the current price is fetched
    - Calculates the ema crosses based on the fetched information
    - Issues the sell list
    - Consumer eats the issued sell list (from the following steps, a loop is expected)
    - ~Loop Starts~
      - Consumer checks if we currently own the ticker.
          - IF NOT: Skip the ticker.
          - IF YES: continue to the next step.
      - Consumer checks if, in the currently placed sell orders, the ticker is included.
          - IF INCLUDED: Skip the ticker.
          - IF NOT INCLUDED: continue to the next step.
      - Consumer checks if, in the currently placed buy orders, the ticker is included.
          - IF INCLUDED: consumer cancels the placed buy order of the ticker, and excludes it from the sell list (i.e. we don't place another sell order for that stock). After that, skip the ticker.
          - IF NOT INCLUDED: continue to the next step.
      - Consumer places a sell order.
  #### Buy
    - Every hour the current price is fetched
    - Calculates the ema crosses based on the fetched information
    - Issues the buy list
    - Consumer eats the issued buy list (from the following steps, a loop is expected)
    - ~Loop Starts~
      - Consumer checks if, in the currently placed buy orders, the ticker is included.
          - IF INCLUDED: Skip the ticker.
          - IF NOT INCLUDED: continue to the next step.
      - Consumer checks if, in the currently placed sell orders, the ticker is included.
          - IF INCLUDED: consumer cancels the placed sell order of the ticker, and excludes it from the buy list (i.e. we don't place another buy order for that stock, as we alrealy own the stock if the canceling of the sell ticker is successful). After that, skip the ticker.
          - IF NOT INCLUDED: continue to the next step.
      - Consumer adds to the final filtered list.
    - Consumer takes the final filtered list and calculates the quantity and price for each of the ticker in the filtered list.
    - Consumer places the buy orders based on the above step.

### Concerns
- By running the ema_simulate.py a few times, I've found out that detecting the cross as quickly as possible is the most crucial.
- The best performance was, quite obviously, when the user **bought a share at the lowest price on the day that had a golden cross, and sold at the highest price on the day that had a dead cross**.
- Even delaying to the next day had a relatively big impact on the performance. 
- However, for the current setup, we can only buy/sell on the next day. 

```
(current: Buy/Sell at next day open)
buy_price = data.loc[previous_gold_cross[2] + 1]['open']
sell_price = data.loc[dead_cross[2] + 1]['open']
=> At 2/18 ~ 5/13, the benefit is $620.110673904419

(Buy/Sell at next day high/low)
buy_price = data.loc[previous_gold_cross[2] + 1]['high']
sell_price = data.loc[dead_cross[2] + 1]['low']
=> At 2/18 ~ 5/13, the benefit is $7598.955770492554

(Buy/Sell at the same day close)
buy_price = data.loc[previous_gold_cross[2]]['close']
sell_price = data.loc[dead_cross[2]]['close']
=> At 2/18 ~ 5/13, the benefit is $481.8248586654663

(best performance)
buy_price = data.loc[previous_gold_cross[2]]['low']
sell_price = data.loc[dead_cross[2]]['high']
=> At 2/18 ~ 5/13, $12902.391907691956

```
- Hence...

# To Do
- Improve on the simulation conditions
    - Introduce budget system
    - Introduce rating system (i.e. choose tickers with better condition and buy in a bulk, as the commission is expensive)

*****

## For simulation (ema_simulate.py):
### Commands
#### 1. By Individual Stocks

``` python ema_simulate.py 'symbol' '["GS"]' ```
``` python ema_simulate.py 'symbol' '["GS", "APPL"]' ```

#### 2. By Markets

``` python ema_simulate.py 'market' 'sp500' ```
``` python ema_simulate.py 'market' 'dow' ```
``` python ema_simulate.py 'market' 'sp500_dow' ```
``` python ema_simulate.py 'market' 'nasdaq' ```
``` python ema_simulate.py 'market' 'all' ```

**Note: The last two (`nasdaq` and `all`) will pull and analyze about 5000 ~ 5500 stocks, so it's VERY heavy.**

### Usage

- Will return the results if the user bought and sold strictly according to the crosses.
   - the resulting overall profit (starting from 0)
   - order information (dataframe)
```
  Symbol   Buy_price    Buy_day  Sell_price   Sell_day            Change(%)
0         A  132.870596 2021-02-17  124.022304 2021-02-23  -0.9334066952379082
1         A  118.369099 2021-03-12  123.476397 2021-03-25   +1.043147226820111
2         A  122.045400 2021-03-29  136.588301 2021-05-03  +1.1191597631697416
3       AAL   23.396401 2021-03-31   23.401599 2021-04-12  +1.0002221913849076
```

- The simulated conditions is as follows (to be improved):

  - The user has unlimited quantity of money (i.e. The user would buy every single stock that had a EMA golden cross regardless of the available money at that particular time)
  - Everyday, the user would buy every stock that had a golden cross and sell every stock that had a dead cross.
  - The user would buy/sell one share for every stock. 
  - Buy price is the EMA golden cross' next day's open
  - Sell price is the EMA dead cross' next day's open
  - Does not consider stop loss
    