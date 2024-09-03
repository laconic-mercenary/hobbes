import json
import os
import random

class IBKRTest:
    def __init__(self):
        self.stock = "NET"
        self.interval = 1
        self.span = 1
        self.data = {}
        self.index = 0
        with open('{}/test/data/{}_{}m_{}d.json'.format(os.getcwd(), self.stock, self.interval, self.span), 'r') as myfile:
            data=myfile.read()
            self.data = json.loads(data)
        self.current_price = self.GetCurrentPrice()

    def QueryOrderStatus(self, order):
        if order["order-id"] is None:
            return ""
        grid = order["_grid"] ## cheating
        if order["order-id"] == grid["current"]["buy-order"]["order-id"]:
            index = grid["current"]["index"]
            current_buy_price = grid[-index]
            if self.current_price < current_buy_price:
                order["_closed_price"] = self.current_price
                order["_closed"] = True
                print("!! buy-order TRIGGERED")
                return "closed"
            order["_open"] = True
            return "open"
        elif order["order-id"] == grid["current"]["sell-order"]["order-id"]:
            index = grid["current"]["index"]
            current_sell_price = grid[index]
            if self.current_price > current_sell_price:
                order["_closed_price"] = self.current_price
                order["_closed"] = True
                print("!! sell-order TRIGGERED")
                return "closed"
            order["_open"] = True
            return "open"
        else:
            raise Exception()

    def GetCurrentPrice(self):
        self.current_price = float(self.data["data"][self.index]["c"])
        self.index = self.index + 1
        return self.current_price

    def AddBuyStopOrder(self, buy_price):
        return str(random.randint(10000, 99999))

    def AddSellStopOrder(self, sell_price):
        return str(random.randint(10000, 99999))

    def CancelSellStopOrder(self, id):
        pass
