# iot_home_appliance
Automated home with IoT

### Testing aralto

on the server:
```
./arauto_server
```

on the client:
```
curl -H "Content-Type: application/json" -X POST -d '{"cmd": "alert-red","type":"alert","msg":"fail to compile"}' "http://edison.local:9999/"
```
