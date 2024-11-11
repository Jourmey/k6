
## 本地调试
./k6 run examples/http_get.js


## 增加额外输出
./k6 run  --out json=test.json  --out csv=test_results.csv  examples/http_3.js
