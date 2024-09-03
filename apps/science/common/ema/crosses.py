import numpy as np

def ema_crosses(data, spans):
  ema_long = 'ema_{ema_max}'.format(ema_max=max(spans))
  ema_short = 'ema_{ema_min}'.format(ema_min=min(spans))
  golden = []
  dead = []
  for i, item in data.iterrows():
    if (i == 0):
      continue
    # Golden Cross: (short ema is higher than long EMA) and (the previous short ema is lower than the previous long ema)
    if (item[ema_short] > item[ema_long]) & (data.iloc[i - 1][ema_short] <= data.iloc[i - 1][ema_long]):
      golden.append([item[0], i])
    # Dead Cross: (short ema is lower than long EMA) and (the previous short ema is higher than the previous long ema)
    if (item[ema_short] < item[ema_long]) & (data.iloc[i - 1][ema_short] >= data.iloc[i - 1][ema_long]):
      dead.append([item[0], i])     
  return np.array(golden), np.array(dead)