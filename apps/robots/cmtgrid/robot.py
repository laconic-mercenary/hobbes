#!/bin/python

import config
import util
import finance

class Robot:
    def __init__(self, base_price=config.get_base_price(), level_amount=config.get_level_amount(), grid_levels=config.get_grid_levels(), api=finance.get_default()):
        assert level_amount > 0.0
        assert grid_levels > 0
        assert grid_levels % 2 == 0
        assert api is not None
        assert base_price > 0.0
        self.api = api
        self.build_grid(base_price, level_amount, grid_levels)

    def build_grid(self, base_price, level_amount, grid_levels=10):
        self.grid = { 
            0: base_price,
            "levels": grid_levels,
            "current": {
                "buy-order" : {
                    "order-id": None
                },
                "sell-order" : {
                    "order-id": None
                },
                "index": 1
            }
        }
        for i in range(1, int(grid_levels / 2) + 1):
            self.grid[i] = base_price + (level_amount * i)
            self.grid[-1 * i] = base_price - (level_amount * i)
            self.say("GRID: level:{}, sell:{}, buy:{}".format(i, self.grid[i], self.grid[-i]))
        ## DIRTY
        self.grid["current"]["buy-order"]["_grid"] = self.grid
        self.grid["current"]["sell-order"]["_grid"] = self.grid
        ## END DIRTY
    
    def work(self):
        self.say("I'm starting now!")
        while True:
            self.loop()

    def say(self, message):
        print("[Robot, {}] says >>> {}".format(type(self.api).__name__, message))

    def currentGridPosition(self):
        return self.grid["current"]

    def isBuyStopOrderFilled(self, current):
        return self.api.QueryOrderStatus(current["buy-order"]) == "closed"

    def isSellStopOrderFilled(self, current):
        return self.api.QueryOrderStatus(current["sell-order"]) == "closed"

    def isBuyStopOrderOpen(self, current):
        return self.api.QueryOrderStatus(current["buy-order"]) == "open"

    def isSellStopOrderOpen(self, current):
        return self.api.QueryOrderStatus(current["sell-order"]) == "open"

    def isBuyStopOrderPlaced(self, current):
        return self.api.QueryOrderStatus(current["sell-order"]) in ["pending", "open", "closed", "canceled", "expired"]

    def isSellStopOrderPlaced(self, current):
        return self.api.QueryOrderStatus(current["sell-order"]) in ["pending", "open", "closed", "canceled", "expired"]

    def moveNextInGrid(self):
        index = self.grid["current"]["index"]
        self.grid["current"] = {
            "buy-order" : {
                "order-id": None
            },
            "sell-order" : {
                "order-id": None
            },
            "index": index + 1
        }
        self.say("!! Moving to next GRID position !! ({})".format(self.grid["current"]["index"]))

    def checkGridBoundary(self):
        current_price = self.api.GetCurrentPrice()
        levels = int(self.grid["levels"] / 2)
        if current_price > self.grid[levels] or current_price < self.grid[-levels]:
            raise Exception("grid broken")
        self.say("Current Price is {}".format(current_price))

    def reEstablishBuySell(self):
        if not self.isBuyStopOrderPlaced(self.currentGridPosition()):
            self.say("Seems no BUY order is placed")
            if not self.isSellStopOrderPlaced(self.currentGridPosition()):
                self.say("Seems no SELL order is placed")
                self.placeBuyStopOrder(self.currentGridPosition())
                self.placeSellStopOrder(self.currentGridPosition())
            else:
                self.say("There is a SELL order, but no BUY order - better cancel the SELL order")
                self.cancelSellStopOrder(self.currentGridPosition())

    def placeBuyStopOrder(self, current):
        # make sure to do stop-loss
        index = current["index"]
        buy_price = self.grid[-index]
        self.say("!! BUY stop order at {} placed (index {})".format(buy_price, -index))
        id = self.api.AddBuyStopOrder(buy_price)
        current["buy-order"]["order-id"] = id

    def placeSellStopOrder(self, current):
        index = current["index"]
        sell_price = self.grid[index]
        self.say("!! SELL stop order at {} placed (index {})".format(sell_price, index))
        id = self.api.AddSellStopOrder(sell_price)
        current["sell-order"]["order-id"] = id

    def cancelSellStopOrder(self, current):
        id = current["sell-order"]["order-id"]
        self.say("!! SELL stop order cancelled")
        self.api.CancelSellStopOrder(id)
        current["sell-order"]["order-id"] = None

    def loop(self):
        self.show_grid()
        self.sleep()
        if self.isBuyStopOrderFilled(self.currentGridPosition()):
            if self.isSellStopOrderFilled(self.currentGridPosition()):
                self.moveNextInGrid()
            elif self.isSellStopOrderOpen(self.currentGridPosition()):
                self.checkGridBoundary()
            else:
                self.reEstablishBuySell()
        elif self.isBuyStopOrderOpen(self.currentGridPosition()):
            self.checkGridBoundary()
        else:
            self.reEstablishBuySell()

    def sleep(self):
        sleep_time = config.get_seeking_sleep_time()
        self.say("Let me rest for {} seconds".format(sleep_time))
        util.sleep(sleep_time)

    def show_grid(self):
        sellsmsg = "sells:"
        buyssmsg = "buys :"
        for i in range(1, int(self.grid["levels"] / 2) + 1):
            sellsymbol = "?"
            buysymbol = "?"
            if i < self.grid["current"]["index"]:
                sellsymbol = "X"
                buysymbol = "X"
            elif i == self.grid["current"]["index"]:
                if self.grid["current"]["buy-order"]["order-id"] is not None:
                    buysymbol = "*"
                if self.grid["current"]["sell-order"]["order-id"]is not None:
                    sellsymbol = "*"
            sellsmsg += "{}-{},".format(self.grid[i], sellsymbol)
            buyssmsg += "{}-{},".format(self.grid[-i], buysymbol)
        print("{}{}{}".format(buyssmsg, "\n", sellsmsg))
