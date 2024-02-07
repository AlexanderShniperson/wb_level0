#!/bin/bash
cd $(dirname $0)
for i in {1..1000}
do
    order_id="$(date +%s)$RANDOM"
    result="$(cat ./data.json | sed "s/{{ORDER_ID}}/$(echo $order_id)/g"; echo)"
    echo $result | nats pub databus &
done