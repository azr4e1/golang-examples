import socket

sckt = socket.socket(family=socket.AF_UNIX)
sckt.connect('/tmp/echo.sock')
sckt.send(bytes("ciao".encode('utf-8')))
sckt.recv(1024)
sckt.close()
