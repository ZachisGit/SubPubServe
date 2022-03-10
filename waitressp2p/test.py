from waitress import serve
import falcon_server

serve(falcon_server.App, host="0.0.0.0", port=8081)
