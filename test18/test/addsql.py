
import os


for i in range(1000,2000):
    os.system("curl -H \"Content-Type:application/json\" -X POST --data \'{\"user-name\": \"testcreate"+str(i)+"\",\"user-password\": \"111111\",\"user-nickName\": \"gogogo\"}\' http://localhost:8080/users/")
