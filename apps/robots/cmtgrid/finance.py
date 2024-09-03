from importlib import import_module

import config

def get_default():
    module = config.get_finance_module()
    classname = config.get_finance_module_classname()
    classtype = getattr(import_module(module), classname)
    return classtype()