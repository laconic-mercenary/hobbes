import common.fetch_stock_info as fetch_stock_info
import ast
import pandas as pd

def listed_symbols(kind, name):
  if (kind == "symbol"):
    return ast.literal_eval(name)
  elif (kind == "market"):
    if (name == "sp500"):
      return fetch_stock_info.tickers_sp500()
    elif (name == "dow"):
      return fetch_stock_info.tickers_dow()
    elif (name == "sp500_dow"):
      tickers = fetch_stock_info.tickers_sp500()
      dow = fetch_stock_info.tickers_dow()
      tickers.extend(dow)
      return tickers
    elif (name == "nasdaq"):
      return get_nasdaq()
    elif (name == "all"):
      tickers = fetch_stock_info.tickers_sp500()
      dow_tickers = fetch_stock_info.tickers_dow()
      nasdaq = get_nasdaq()
      tickers.extend(dow_tickers)
      tickers.extend(nasdaq)
      return tickers

def get_nasdaq():
  nasdaq = pd.read_csv('http://ftp.nasdaqtrader.com/dynamic/SymDir/nasdaqlisted.txt', sep='|')
  nasdaq_normal = nasdaq[nasdaq['Financial Status']=='N']
  nasdaq_normal = nasdaq_normal[nasdaq_normal['Test Issue']=='N']
  nasdaq_tickers = nasdaq_normal[nasdaq_normal['ETF']=='N'][['Symbol']]
  return nasdaq_tickers['Symbol'].tolist()