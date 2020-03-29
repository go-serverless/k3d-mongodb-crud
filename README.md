# k3d-mongodb-crud

Building Serverless CRUD services in Go with OpenFaaS, Arkade, MongoDB and k3d

> This is just a learning playground

Check out the blog post [here](https://dev.to/wingkwong/building-serverless-crud-services-in-go-with-openfaas-arkade-mongodb-and-k3d-1d93)

### Create
```
curl http://127.0.0.1:8080/function/k3d-mongodb-crud --request POST
{"success":true, "message": "4 records have been inserted"}
```

Go to MongoDB, check the collection
```
> db.foo.find()
{ "_id" : ObjectId("5e8014f10c4812773ee77f16"), "bar" : "bar", "baz" : "baz" }
{ "_id" : ObjectId("5e8014f10c4812773ee77f17"), "bar" : "bar1", "baz" : "baz1" }
{ "_id" : ObjectId("5e8014f10c4812773ee77f18"), "bar" : "bar2", "baz" : "baz2" }
{ "_id" : ObjectId("5e8014f10c4812773ee77f19"), "bar" : "bar3", "baz" : "baz3" }
```

### Read
```
curl http://127.0.0.1:8080/function/k3d-mongodb-crud --request GET
[{"Bar":"bar","Baz":"baz"},{"Bar":"bar1","Baz":"baz1"},{"Bar":"bar2","Baz":"baz2"},{"Bar":"bar3","Baz":"baz3"}]
```

### Update 
```
curl http://127.0.0.1:8080/function/k3d-mongodb-crud --request PUT
{"success":true, "message": "bar1 has been updated to bar1-updated"}
```

Go to MongoDB, check the collection
```
> db.foo.find()
{ "_id" : ObjectId("5e8014f10c4812773ee77f16"), "bar" : "bar", "baz" : "baz" }
{ "_id" : ObjectId("5e8014f10c4812773ee77f17"), "bar" : "bar1-updated", "baz" : "baz1" }
{ "_id" : ObjectId("5e8014f10c4812773ee77f18"), "bar" : "bar2", "baz" : "baz2" }
{ "_id" : ObjectId("5e8014f10c4812773ee77f19"), "bar" : "bar3", "baz" : "baz3" }
```

### Delete
```
curl http://127.0.0.1:8080/function/k3d-mongodb-crud --request DELETE
{"success":true, "message": "bar1 has been deleted"}
```

```
> db.foo.find()
{ "_id" : ObjectId("5e8014f10c4812773ee77f17"), "bar" : "bar1-updated", "baz" : "baz1" }
{ "_id" : ObjectId("5e8014f10c4812773ee77f18"), "bar" : "bar2", "baz" : "baz2" }
{ "_id" : ObjectId("5e8014f10c4812773ee77f19"), "bar" : "bar3", "baz" : "baz3" }
```
