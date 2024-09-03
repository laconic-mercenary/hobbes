from flask import Flask, request, make_response
from cloudevents.http import CloudEvent, to_structured

import os
import requests
import json

app = Flask(__name__)

@app.route('/health', methods=['GET', 'POST'])
def health_check():
    return make_response({'status':'ok'})

@app.route('/channel', methods=['POST'])
def send_ce_to_hobbes_channel():
    data = json.loads(request.data)
    attributes = {
        "type": "org.hobbes.event",
        "source": "http://events.hobbes.svc.cluster.local",
    }
    event = CloudEvent(attributes, data)
    headers, body = to_structured(event)
    try:
        requests.post(os.environ['K_SINK'], data=body, headers=headers, timeout=2.5)
    except ValueError as e:
        app.logger.error(e)
        return 'invalid request', 400
    except Exception as e:
        app.logger.error(e)
        return 'server error', 500
    return make_response({'forward':'ok', 'target':os.environ['K_SINK']})

@app.route('/', methods=['POST'])
def receive_event():
    print(request.json)
    app.logger.info(request.json)
    return "ok"

### main ###
if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8080)
