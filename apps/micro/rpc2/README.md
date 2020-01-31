各服务生成模型命令：

inventory-srv
```
$  micro new --namespace=mu.micro.book --type=srv --alias=inventory github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/inventory-srv
```
order-web

```
$  micro new --namespace=mu.micro.book --type=web --alias=order github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/orders-web
```
order-srv
```
$  micro new --namespace=mu.micro.book --type=srv --alias=order github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/orders-srv
```
payment-web
```
$  micro new --namespace=mu.micro.book --type=web --alias=payment github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/payment-web
```
payment-srv
```
$  micro new --namespace=mu.micro.book --type=srv --alias=payment github.com/elitecodegroovy/gnetwork/apps/micro/rpc2/payment-srv
```