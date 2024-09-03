(ns refresh.core
  (:gen-class)
  (:require
    [cambium.core :as log]
    [refresh.http :as rhttp]
    [refresh.util :as rutil]
    [refresh.config :as rcfg]))

(defn- is-session-ok []
  (log/debug "refresh.core/is-session-ok")
  (= (rhttp/request (rcfg/get-ibkr-gateway)) 200))

(defn- check-ibkr-session []
  (log/debug "refresh.core/check-ibkr-session")
  (log/info {:session (is-session-ok)} "Response"))

(defn -main
  "Main entrypoint"
  [& args]
  (rutil/init-logging)
  (log/debug "refresh.core/main")
  (log/info "Started")
  (check-ibkr-session)
  (log/info "Finished")
)
