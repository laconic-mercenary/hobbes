FROM python:3.8-slim

ENV ROBOT_GRID_BASEPRICE="0.0"
ENV ROBOT_FINANCIAL_ENGINE_MODULE="test"
ENV ROBOT_FINANCIAL_ENGINE_CLASSNAME="IBKRTest"

RUN mkdir /app

WORKDIR /app

COPY requirements.txt .

RUN pip install -r requirements.txt

COPY . .

ENTRYPOINT [ "python", "main.py" ]