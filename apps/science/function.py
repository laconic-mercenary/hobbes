# python3 -m pdb function.py '["GS"]' 50
# pyenv global 3.8.0
# eval "$(pyenv init -)"
# cd apps/science
# python function.py 'symbol' '["MARA"]' 50
# python function.py 'market' 'sp500_dow'

import sys
import json
from datetime import datetime, date, timedelta
import yaml
import numpy as np
import common.fetch_stock_info as fetch_stock_info
import common.get_list_of_symbols as get_list_of_symbols
import common.indicators as indicators
import common.ema.deviations as deviations
import common.ema.crosses as crosses
import common.ema.determine_appropriateness as determine_appropriateness
# import common.ema.plots as plots
# import pdb

def get_previous_time(days_before = -60):
  previous_time = datetime.now() + timedelta(days=days_before)
  str_previous_time = ('{}/{}/{}').format(previous_time.year, str(previous_time.month).zfill(2), str(previous_time.day).zfill(2))
  return str_previous_time

def build_output(crosses_buy_tickers, crosses_sell_tickers, deviations_buy_tickers, deviations_sell_tickers):
  buy_tickers = np.unique(crosses_buy_tickers + deviations_buy_tickers)
  sell_tickers = np.unique(crosses_sell_tickers + deviations_sell_tickers)
  return buy_tickers, sell_tickers

def print_results(buys_label, buys, sells_label, sells):
  buys = buys.tolist() if type(buys) is not list else buys
  sells = sells.tolist() if type(sells) is not list else sells
  buys_out = json.dumps({buys_label : buys})
  sells_out = json.dumps({sells_label : sells})
  print(buys_out)
  print(sells_out)
  
def function():
  tickers = get_list_of_symbols.listed_symbols(sys.argv[1], sys.argv[2])

  with open('config.yml', 'r') as file:
    settings = yaml.safe_load(file)

  today = date.today().strftime('%Y/%m/%d')
  str_previous_time = get_previous_time(settings['data']['span'])

  crosses_buy_tickers = []
  crosses_sell_tickers = []
  deviations_buy_tickers = []
  deviations_sell_tickers = []
  
  for ticker in tickers:
    try:
      raw_data = fetch_stock_info.get_data(ticker, start_date = str_previous_time, end_date = today)
    except:
      continue

    history = raw_data.reset_index().rename(columns={"index": "date"})

    # Skip shares that are cheaper than the min price
    if (history.iloc[-1]["adjclose"] < settings['price']['min']):
      continue

    # Calculate EMA
    # history = indicators.ema(history, settings['ema']['span'])

    # Calculate ADX/DI
    # try:
    #   history = indicators.adx_di(history, settings['adx_di']['span'])
    # except ValueError:
    #   continue
    # except IndexError:
    #   continue

    # Calculate SMA deviation rate
    try:
      history = deviations.sma(history, settings['sma']['deviation_span'])
    except:
      continue

    # Calculate Golden crosses / Dead crosses
    # ema_golden_crosses, ema_dead_crosses = crosses.ema_crosses(history, settings['ema']['span'])

    # Determine if a buy is apporopriate from ema cross standpoint
    # ema_buy_signal = determine_appropriateness.buy_ema_cross(history, ema_golden_crosses)
    # if (ema_buy_signal == True):
    #   crosses_buy_tickers.append(ticker)

    # Determine if a sell is apporopriate from ema cross standpoint
    # ema_sell_signal = determine_appropriateness.sell_ema_cross(ema_dead_crosses)
    # if (ema_sell_signal == True):
    #   crosses_sell_tickers.append(ticker)

    # Determine if a buy is apporopriate from a sma deviation rate standpoint
    sma_buy_signal = determine_appropriateness.buy_sma_deviation(history.tail(1))
    if (sma_buy_signal == True):
      deviations_buy_tickers.append(ticker)

    # Determine if a sell is apporopriate from a sma deviation rate standpoint
    sma_sell_signal = determine_appropriateness.sell_sma_deviation(history.tail(1))
    if (sma_sell_signal == True):
      deviations_sell_tickers.append(ticker)

  buy_tickers, sell_tickers = build_output(crosses_buy_tickers, crosses_sell_tickers, deviations_buy_tickers, deviations_sell_tickers)

  print_results('regular_buy_tickers', buy_tickers, 'regular_sell_tickers', sell_tickers)
  print_results('crosses_buy_tickers', crosses_buy_tickers, 'crosses_sell_tickers', crosses_sell_tickers)
  print_results('deviations_buy_tickers', deviations_buy_tickers, 'deviations_sell_tickers', deviations_sell_tickers)

def main():
  function()
  
if __name__ == "__main__":
  main()