# create node
curl http://127.0.0.1:32767/apulisEdge/api/node/createNode -d @createNode.json|jq .

#  list all nodes
curl http://127.0.0.1:32767/apulisEdge/api/node/listNodes -d @listNodes.json|jq .

# describe node
curl http://127.0.0.1:32767/apulisEdge/api/node/desNode -d @desNode.json|jq .
