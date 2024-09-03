from datetime import datetime, date
import yaml
import os

yml = "{}/../../config.yml".format(os.path.dirname(os.path.abspath(__file__)))
with open(yml, 'r') as file:
  settings = yaml.safe_load(file)

def buy_ema_cross(data, golden_crosses):
  buy_signal = False
  for cross in golden_crosses:
    # Is the cross today
    if (datetime.strftime(date.today(), '%Y-%m-%d') == datetime.strftime(cross[0], '%Y-%m-%d')):
        # ADX: Strength of the trend
        if (data.loc[cross[1]]['adx'] > settings['adx_di']['adx_trend_demarcation']):
          buy_signal = True
  return buy_signal

def sell_ema_cross(dead_crosses):
  sell_signal = False
  for cross in dead_crosses:
    # Is the cross today
    if (datetime.strftime(date.today(), '%Y-%m-%d') == datetime.strftime(cross[0], '%Y-%m-%d')):
      sell_signal = True
  return sell_signal

def buy_sma_deviation(data):
  buy_signal = False

  mean = (data["sma_minusDev"] - data["sma_stdCenter"]) / 100
  acceptable_max = data["sma_minusDev"] - mean

  if ((acceptable_max > data["sma_deviation"]).bool()):
    buy_signal = True
  return buy_signal

def sell_sma_deviation(data):
  sell_signal = False

  mean = (data["sma_plusDev"] - data["sma_stdCenter"]) / 100
  acceptable_max = data["sma_plusDev"] - mean

  if ((acceptable_max < data["sma_deviation"]).bool()):
    sell_signal = True
  return sell_signal

