import xmlrpc.client

s = xmlrpc.client.ServerProxy('http://localhost:8000', verbose=True)
print(s.join({'First':'foo', 'Second':'bar'}))
