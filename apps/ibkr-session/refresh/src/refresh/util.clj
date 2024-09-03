(ns refresh.util
    (:gen-class)
    (:require
        [cambium.codec :as codec]
        [cambium.core  :as log]
        [cambium.logback.json.flat-layout :as flat]))

(defn init-logging []
    (flat/set-decoder! codec/destringify-val))
