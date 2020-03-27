from xmlrpc.server import SimpleXMLRPCServer
from xmlrpc.server import SimpleXMLRPCRequestHandler

# Restrict to a particular path.
class RequestHandler(SimpleXMLRPCRequestHandler):
    rpc_paths = ('/RPC2',)

# Create server
with SimpleXMLRPCServer(('localhost', 8000),
                        requestHandler=RequestHandler) as server:                
    server.register_introspection_functions()

    def join_string(dic):
        first = dic['First']
        second = dic['Second']
        return {'Joined':first+second}
    
    server.register_function(join_string, 'join')

    server.serve_forever()