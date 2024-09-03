import yaml
import os

def simulate(data, ema_golden_crosses, ema_dead_crosses):

  yml = "{}/../../config.yml".format(os.path.dirname(os.path.abspath(__file__)))

  with open(yml, 'r') as file:
    settings = yaml.safe_load(file)

  result = 0
  order_details = []
  cost = 0
  earnings = 0

  for dead_cross in ema_dead_crosses:
    previous_gold_cross_date = min([i[0] for i in ema_golden_crosses if i[0] <= dead_cross[0]], key=lambda x: abs(x - dead_cross[0]), default=None)

    if (previous_gold_cross_date == None):
      continue

    for i in ema_golden_crosses:
      if (i[0] == previous_gold_cross_date):
        previous_gold_cross = i
    
    if (data.loc[previous_gold_cross[1]]['adx'] > settings['adx_di']['adx_trend_demarcation']):

      # buy_price = (data.loc[previous_gold_cross[1] + 1]['low'] + data.loc[previous_gold_cross[1] + 1]['high']) / 2
      # sell_price = (data.loc[dead_cross[1] + 1]['low'] + data.loc[dead_cross[1] + 1]['high']) / 2

      buy_price = data.loc[previous_gold_cross[1] + 1]['open']
      sell_price = data.loc[dead_cross[1] + 1]['open']

      diff = sell_price - buy_price
      result += diff

      change_percentage = sell_price / buy_price
      sign = "-" if change_percentage < 1 else "+"

      cost += buy_price
      earnings += sell_price

      order_detail = {
        'Symbol': data.loc[previous_gold_cross[1]]['ticker'],
        'Buy_price': buy_price,
        'Buy_day': data.loc[previous_gold_cross[1] + 1]['date'],
        'Sell_price': sell_price,
        'Sell_day': data.loc[dead_cross[1] + 1]['date'],
        'Change(%)': ('{}{}').format(sign, change_percentage)
      }

      order_details.append(order_detail)

  return result, order_details, cost, earnings
