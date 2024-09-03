FROM python:3.8-slim

ARG kind='market'

ARG target='all'

ENV kind=${kind}

ENV target=${target}

WORKDIR /code

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

CMD python function.py ${kind} ${target}
