from statistics import stdev
import common.indicators as indicators
import matplotlib.pyplot as plt

def sma(data, deviation_span):
  # fetch sma
  data = indicators.sma(data, deviation_span)
  sma_name = 'sma_{sma_span}'.format(sma_span=deviation_span)
  doubled = deviation_span * 2

  # calculate deviation rate
  data["sma_deviation"] = data['adjclose'] / data[sma_name] * 100 - 100
  data = data.drop(data.index[0]).reset_index(drop=True)

  # calculate center
  data["sma_stdCenter"] = data['sma_deviation'].rolling(doubled).mean()
  data = data.drop(data.index[0:50]).reset_index(drop=True)

  # calculate standard deviation
  stdVal = stdev(data["sma_deviation"])

  # Standarization
  data["sma_plusDev"] = data["sma_stdCenter"] + stdVal * 2
  data["sma_minusDev"] = data["sma_stdCenter"] - stdVal * 2

  return data
  