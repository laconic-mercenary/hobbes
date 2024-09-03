(defproject refresh "0.1.0-SNAPSHOT"
  :description "IBKR gateway/clientportal session refresh app"
  :url "http://example.com/FIXME"
  :license {:name "EPL-2.0 OR GPL-2.0-or-later WITH Classpath-exception-2.0"
            :url "https://www.eclipse.org/legal/epl-2.0/"}
  :dependencies [[org.clojure/clojure         "1.10.1"]
                 [cambium/cambium.core           "1.1.0"]
                 [cambium/cambium.codec-cheshire "1.0.0"]
                 [cambium/cambium.logback.json   "0.4.4"]
                 [clj-http                       "3.12.0"]
                 [slingshot                      "0.12.2"]]
  :main ^:skip-aot refresh.core
  :target-path "target/%s"
  :profiles {:uberjar {:aot :all
                       :jvm-opts ["-Dclojure.compiler.direct-linking=true"]}})
