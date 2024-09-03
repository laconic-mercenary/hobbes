(ns refresh.http
    (:gen-class)
    (:require
        [cambium.core :as log]
        [clj-http.client :as client]
        [slingshot.slingshot :as ss]
        [refresh.config :as cfg]))

(defn- post-json [target json headers connect-timeout-ms socket-timeout-ms]
    (log/debug "refresh.http/post-json")
    (client/post target
        { :body json
          :throw-exceptions false
          :headers headers
          :content-type :json
          :socket-timeout socket-timeout-ms
          :connection-timeout connect-timeout-ms
          :accept :json }))

(defn- log-succeeded []
    (log/info { :statusCode 200 } "Refresh-Succeeded"))

(defn- log-unauthorized []
    (log/warn { :statusCode 401 } "Unauthorized"))

(defn- log-server-err []
    (log/warn { :statusCode 500 } "Server-Error"))

(defn- log-unknown [status]
    (log/error { :statusCode status } "Unknown-Status-Received"))

;; https://github.com/dakrone/clj-http

(defn request [target]
    (log/debug "refresh.http/request")
    (let [response (post-json target "{}" {"X-Sender-Source" "hobbes"} (cfg/get-client-timeout-ms) (cfg/get-socket-timeout-ms))]
         (cond 
            (= (:status response) 200) (log-succeeded)
            (= (:status response) 401) (log-unauthorized)
            (= (:status response) 500) (log-server-err)
            :default (log-unknown (:status response)))
        (:status response)))
