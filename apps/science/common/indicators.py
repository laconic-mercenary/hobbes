from ta.trend import EMAIndicator
from ta.trend import SMAIndicator
from ta.trend import MACD
from ta.trend import ADXIndicator
from ta.momentum import RSIIndicator
from ta.momentum import ultimate_oscillator
from ta.volume import MFIIndicator
import pandas as pd

def rsi(data, span):
  data["rsi"] = RSIIndicator(data["adjclose"], window = span, fillna = True).rsi()
  return data

def ult(data, span, weight):
  data["ult"] = ultimate_oscillator(
    data["high"], data["low"], data["adjclose"],
    span['short'], span['middle'], span['long'],
    weight['short'], weight['middle'], weight['long'],
    fillna = True
  )
  return data

def mfi(data, span):
  data["mfi"] = MFIIndicator(data["high"], data["low"], data["adjclose"], data["volume"], window = span, fillna = True).money_flow_index()
  return data

def ema(data, spans):
  for span in spans:
    name = 'ema_{span}'.format(span=span)
    data[name] = EMAIndicator(data["adjclose"], window = span, fillna = True).ema_indicator()
  return data

def sma(data, span):
  name = 'sma_{span}'.format(span=span)
  data[name] = SMAIndicator(data["adjclose"], window = span, fillna = True).sma_indicator()
  return data

def sma_by_deviation_rate(data, span):
  copy = data.copy()
  copy["sma_stdCenter"] = SMAIndicator(data["sma_deviation"], window = span, fillna = True).sma_indicator()
  return copy

def macd(data, spans):
  data["macd"] = MACD(data["adjclose"], window_slow = spans['slow'], window_fast = spans['fast'], window_sign = spans['sign'], fillna = True).macd()
  data["macd_signal"] = MACD(data["adjclose"], window_slow = spans['slow'], window_fast = spans['fast'], window_sign = spans['sign'], fillna = True).macd_signal()
  return data.drop(data.index[[0,1]]).reset_index(drop=True)

def adx_di(data, span):
  data["adx"] = ADXIndicator(data["high"], data["low"], data["adjclose"], span, fillna = True).adx()
  # data["adx_pos"] = ADXIndicator(data["high"], data["low"], data["adjclose"], span, fillna = True).adx_pos()
  # data["adx_neg"] = ADXIndicator(data["high"], data["low"], data["adjclose"], span, fillna = True).adx_neg()
  return data.drop(data.index[0:25]).reset_index(drop=True)

# rsi_trend = get_trendline(rsi.index.values, rsi['rsi_14'].values)
# def get_trendline(index, data, order=1):
#   coeffs = np.polyfit(index, list(data), order)
#   slope = coeffs[-2]
#   return float(slope)