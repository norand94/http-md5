#!/usr/bin/env bash
curl --data 'request={"Id": -10, "Text" : "test_string"}' localhost:8888
curl --data 'request={"Text" : "test_string"}' localhost:8888
curl --data 'request={"Id": 10}' localhost:8888
curl --data 'request={"Id": 10, "Text" : "test_strrterdgerww4rwmfkmf.wkfwfekmfwefkmwefnelnqmqewfmwelfnqlnqwefneqwmqefnwelqwqwefqnqwqfqlqwqwnqejfnqwdnqefnwefqldmefwwgwgwjfhjwfkuhwefhsewhgrjergjkwglwegljkwdnqfjnqwldwqfjnqekjfing"}' localhost:8888
curl --data 'request={foobar}' localhost:8888

curl --data 'request={"Id": 10, "Text" : "test_string"}' localhost:8888 #Нормальный запрос
echo " /n"