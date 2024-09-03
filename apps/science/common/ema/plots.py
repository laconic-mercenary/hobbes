import matplotlib.pyplot as plt

def describe_candles_ema_crosses_adx(data, name, ema_spans, ema_golden_crosses, ema_dead_crosses):
  for i in range(len(data)):
    # data.iloc[i, 0] => datetime
    # data.iloc[i, 1] => open
    # data.iloc[i, 2] => high
    # data.iloc[i, 3] => low
    # data.iloc[i, 4] => close
    plt.vlines(x = data.iloc[i, 0], ymin = data.iloc[i, 3], ymax = data.iloc[i, 2], color = 'black', linewidth = 1)
    # close > high = Bull day
    if data.iloc[i, 4] > data.iloc[i, 1]:
        plt.vlines(x = data.iloc[i, 0], ymin = data.iloc[i, 1], ymax = data.iloc[i, 4], color = 'green', linewidth = 4)
    # close < high = Bear day
    if data.iloc[i, 4] < data.iloc[i, 1]:
        plt.vlines(x = data.iloc[i, 0], ymin = data.iloc[i, 4], ymax = data.iloc[i, 1], color = 'red', linewidth = 4) 
    # close = high = Doji day 
    if data.iloc[i, 4] == data.iloc[i, 1]:
        plt.vlines(x = data.iloc[i, 0], ymin = data.iloc[i, 4], ymax = data.iloc[i, 1], color = 'black', linewidth = 4)

  # plot EMA
  ema_max = 'ema_{ema_max}'.format(ema_max=max(ema_spans))
  ema_min = 'ema_{ema_min}'.format(ema_min=min(ema_spans))
  plt.plot(data['date'], data[ema_max], label = ema_max)
  plt.plot(data['date'], data[ema_min], label = ema_min)

  # plot EMA crosses
  for i in range(len(ema_golden_crosses)):
    plt.plot(ema_golden_crosses[i][0], ema_golden_crosses[i][1], "o", markersize = 5, markerfacecolor='g', markeredgecolor='b', label='Golden Cross' if i == 0 else '')
  for i in range(len(ema_dead_crosses)):
    plt.plot(ema_dead_crosses[i][0], ema_dead_crosses[i][1], "o", markersize = 5, markerfacecolor='r', markeredgecolor='b', label='Dead Cross' if i == 0 else '')
  
  plt.legend()
  plt.grid()
  plt.title(name)
  plt.show()