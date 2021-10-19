#!/usr/bin/env python3

import socket
import gen.proto.python.echo.v1.echo_pb2 as echov1

HOST = '127.0.0.1'  # The server's hostname or IP address
PORT = 65432        # The port used by the server

'''Simple python client'''

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))
    s.sendall(b'Hello, world')
    echov1.EchoRequest
    data = s.recv(1024)

print('Received', repr(data))

