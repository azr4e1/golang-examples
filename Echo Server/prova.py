import socket

sckt = socket.socket(family=socket.AF_INET)
sckt.connect(('127.0.0.1', 9999))
sckt.send(bytes("ciao".encode('utf-8')))
sckt.recv(1024)
sckt.close()
