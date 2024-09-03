import os

__ENV_GRID_BASEPRICE = "ROBOT_GRID_BASEPRICE"
__ENV_GRID_LEVELAMOUNT = "ROBOT_GRID_LEVELAMOUNT"
__ENV_FIN_ENGINE_MODULE = "ROBOT_FINANCIAL_ENGINE_MODULE"
__ENV_FIN_ENGINE_CLASSNAME = "ROBOT_FINANCIAL_ENGINE_CLASSNAME"

def __get_or_die(env):
    assert env in os.environ
    return os.environ[env]

def get_configs():
    return os.environ

def get_grid_levels():
    return 10

def get_base_price():
    return float(__get_or_die(__ENV_GRID_BASEPRICE))

def get_level_amount():
    return float(__get_or_die(__ENV_GRID_LEVELAMOUNT))

def get_finance_module():
    mod = __get_or_die(__ENV_FIN_ENGINE_MODULE)
    return mod

def get_finance_module_classname():
    classname = __get_or_die(__ENV_FIN_ENGINE_CLASSNAME)
    assert classname.isalpha()
    return classname

def get_seeking_sleep_time():
    return 2