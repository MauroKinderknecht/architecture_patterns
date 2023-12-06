bring cloud;
bring ex;
bring util;

// import code from javascript
class UUIDv5 {
    extern "./uuid.js" pub static inflight generateUUIDv4(): str;
}

// define a rest api (API Gateway in AWS)
let api = new cloud.Api();

// define a rest api (DynamoDB in AWS)
let itemsDb = new ex.Table({
    name: "items",
    primaryKey: "id",
    columns: {
        "id" => ex.ColumnType.STRING,
        "name" => ex.ColumnType.STRING,
        "price" => ex.ColumnType.NUMBER,
    }
});

api.post("/items", inflight (req: cloud.ApiRequest): cloud.ApiResponse => {
    let body = Json.parse(req.body ?? "");
    let id = UUIDv5.generateUUIDv4();

    if let existingItem = itemsDb.tryGet(id) {
        return cloud.ApiResponse {
            status: 409,
            headers: {
                "Content-Type" => "text/plain"
            },
            body: "Item already exists"
        };
    }

    let item = Json {
        id: id,
        name: body.get("name").asStr(),
        email: body.get("email").asStr(),
        password: body.get("price").asNum(),
    };

    itemsDb.insert(id, item);
    return cloud.ApiResponse {
        status: 201,
        headers: {
            "Content-Type" => "application/json"
        },
        body: Json.stringify(item)
    };
});

api.put("/items/{id}", inflight (req: cloud.ApiRequest): cloud.ApiResponse => {
    let id = req.vars.get("id");
    let body = Json.parse(req.body ?? "");

    if let existingItem = itemsDb.tryGet(id) {
        let item = Json {
            id: id,
            name: body.get("name").asStr(),
            email: body.get("email").asStr(),
            password: body.get("price").asNum(),
        };

        itemsDb.update(id, item);
        return cloud.ApiResponse {
            status: 200,
            headers: {
                "Content-Type" => "application/json"
            },
            body: Json.stringify(item)
        };
    }

    return cloud.ApiResponse {
        status: 404,
        headers: {
            "Content-Type" => "text/plain"
        },
        body: "Not found"
    };
});

api.get("/items", inflight (req: cloud.ApiRequest): cloud.ApiResponse => {
    let users = itemsDb.list();

    return cloud.ApiResponse {
        status: 200,
        headers: {
            "Content-Type" => "application/json"
        },
        body: Json.stringify(users)
    };
});

api.get("/items/{id}", inflight (req: cloud.ApiRequest): cloud.ApiResponse => {
    let id = req.vars.get("id");

    if let existingItem = itemsDb.tryGet(id) {
        return cloud.ApiResponse {
            status: 200,
            headers: {
                "Content-Type" => "application/json"
            },
            body: Json.stringify(existingItem)
        };
    }

    return cloud.ApiResponse {
        status: 404,
        headers: {
            "Content-Type" => "text/plain"
        },
        body: "Not found"
    };
});

api.delete("/items/{id}", inflight (req: cloud.ApiRequest): cloud.ApiResponse => {
    let id = req.vars.get("id");

    if let existingItem = itemsDb.tryGet(id) {
        itemsDb.delete(id);
        return cloud.ApiResponse {
            status: 200,
            headers: {
                "Content-Type" => "text/plain"
            },
            body: "Deleted"
        };
    }

    return cloud.ApiResponse {
        status: 404,
        headers: {
            "Content-Type" => "text/plain"
        },
        body: "Not found"
    };
});