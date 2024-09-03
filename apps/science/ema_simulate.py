# python ema_simulate.py 'symbol' '["GS", "APPL"]' 50
# python ema_simulate.py 'market' 'sp500_dow'
# python3 -m pdb ema_simulate.py '["GS"]' 50
# pyenv global 3.8.0
# eval "$(pyenv init -)"

import sys
from datetime import datetime, date, timedelta
import yaml
import pandas as pd
import common.fetch_stock_info as fetch_stock_info
import common.get_list_of_symbols as get_list_of_symbols
import common.indicators as indicators
import common.ema.crosses as crosses
import simulate.ema.simulate_buy_sell as simulate_buy_sell

def get_previous_time(days_before = -60):
  previous_time = datetime.now() + timedelta(days=days_before)
  str_previous_time = ('{}/{}/{}').format(previous_time.year, str(previous_time.month).zfill(2), str(previous_time.day).zfill(2))
  return str_previous_time
  
def function():
  tickers = get_list_of_symbols.listed_symbols(sys.argv[1], sys.argv[2])

  with open('config.yml', 'r') as file:
    settings = yaml.safe_load(file)

  today = date.today().strftime('%Y/%m/%d')
  str_previous_time = get_previous_time(settings['data']['span'])

  order_object = pd.DataFrame(columns=['Symbol', 'Buy_price', 'Buy_day', 'Sell_price', 'Sell_day', 'Change(%)'])
  orders = []
  result_all = 0
  
  for ticker in tickers:
    try:
      raw_data = fetch_stock_info.get_data(ticker, start_date = str_previous_time, end_date = today)
    except (KeyError, AssertionError):
      continue

    history = raw_data.reset_index().rename(columns={"index": "date"})

    # Skip shares that are cheaper than the min price
    if (history.iloc[-1]["adjclose"] < settings['price']['min']):
      continue
      
    # EMA
    history = indicators.ema(history, settings['ema']['span'])

    # ADX/DI
    try:
      history = indicators.adx_di(history, settings['adx_di']['span'])
    except ValueError:
      continue
    except IndexError:
      continue

    # Golden crosses / Dead crosses
    ema_golden_crosses, ema_dead_crosses = crosses.ema_crosses(history, settings['ema']['span'])

    print(('ema_golden_crosses: \n {} \n ema_dead_crosses: \n {}').format(ema_golden_crosses, ema_dead_crosses))

    result, order_details, cost, earnings = simulate_buy_sell.simulate(history, ema_golden_crosses, ema_dead_crosses)

    orders.extend(order_details)
    result_all += result

  order_object = order_object.append(orders, True)

  print(('orders: \n {} \n result: \n {} \n cost: \n {} \n earnings: \n {}').format(order_object, result_all, cost, earnings))

def main():
  function()
  
if __name__ == "__main__":
  main()