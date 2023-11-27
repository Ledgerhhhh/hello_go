namespace go hello.ledger

struct HelloReq{
    1: required string Id (api.query="id");
    2: optional i32 age;
    3: required list<string> hobbies
}

struct HelloResp{
    1:string RespBody;
}

service HelloService{
    HelloResp HelloMethod(1:HelloReq request)(api.get="/hello")
}


struct HelloReq2{
    1: required string Id (api.query="id");
    2: optional i32 age;
    3: required list<string> hobbies
}

struct HelloResp2{
    1:string RespBody;
}

service HelloService2{
    HelloResp HelloMethod2(1:HelloReq request)(api.get="/hello2")
}

struct HelloReq3{
    1: required string Id (api.query="id");
    2: optional i32 age;
    3: required list<string> hobbies
}

struct HelloResp3{
    1:string RespBody;
}

service HelloService3{
    HelloResp HelloMethod3(1:HelloReq request)(api.get="/hello3")
}


