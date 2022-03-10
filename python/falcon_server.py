import falcon

class Simple(object):
    def on_get(self,req,resp):
        resp.status = falcon.HTTP_200
        resp.set_header("Access-Control-Allow-Origin",'*')
        resp.set_header('Access-Control-Allow-Origin', '*')
        resp.set_header('Access-Control-Allow-Methods', '*')
        resp.set_header('Access-Control-Allow-Headers', '*')
        resp.body = "Hello World!"
        return



App = falcon.App()

# Routes
simple = Simple()

App.add_route('/simple',simple)
