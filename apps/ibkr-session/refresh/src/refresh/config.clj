(ns refresh.config
    (:gen-class)
    (:require
        [cambium.core :as log]))

(defn- get-env [env]
    (let [target (System/getenv env)]
        (if (not target) (throw (new Exception (+ env " env is required"))) target)))

(defn get-ibkr-gateway []
    (get-env "IBKR_GATEWAY_TARGET"))

(defn get-client-timeout-ms []
    (Integer/parseInt (get-env "HTTP_CLIENT_TIMEOUT_MS")))

(defn get-socket-timeout-ms []
    (Integer/parseInt (get-env "HTTP_SOCKET_TIMEOUT_MS")))
