Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)
- [ErrorParser](namespace-aws-api-errorparser.md)

## JsonRpcErrorParser     extends [AbstractErrorParser](class-aws-api-errorparser-abstracterrorparser.md)   in package    - [Aws](package-aws.md)       Uses  [JsonParserTrait](class-aws-api-errorparser-jsonparsertrait.md)

Parsers JSON-RPC errors.

### Table of Contents  [header link](class-aws-api-errorparser-jsonrpcerrorparser-toc.md)

#### Methods  [header link](class-aws-api-errorparser-jsonrpcerrorparser-toc-methods.md)

[\_\_construct()](class-aws-api-errorparser-jsonrpcerrorparser-method-construct.md)
: mixed [\_\_invoke()](class-aws-api-errorparser-jsonrpcerrorparser-method-invoke.md)
: mixed

### Methods  [header link](class-aws-api-errorparser-jsonrpcerrorparser-methods.md)

#### \_\_construct()  [header link](class-aws-api-errorparser-jsonrpcerrorparser-method-construct.md)

`
    public
                    __construct([Service|null $api = null ][, JsonParser|null $parser = null ]) : mixed`

##### Parameters

$api
: [Service](class-aws-api-service.md) \|null
= null$parser
: JsonParser\|null
= null

#### \_\_invoke()  [header link](class-aws-api-errorparser-jsonrpcerrorparser-method-invoke.md)

`
    public
                    __invoke(ResponseInterface $response[, CommandInterface|null $command = null ]) : mixed`

##### Parameters

$response
: [ResponseInterface](class-psr-http-message-responseinterface.md)$command
: [CommandInterface](class-aws-commandinterface.md) \|null
= null
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-api-errorparser-jsonrpcerrorparser-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-api-errorparser-jsonrpcerrorparser-method-construct.md)
  - [\_\_invoke()](class-aws-api-errorparser-jsonrpcerrorparser-method-invoke.md)

[Back To Top](class-aws-api-errorparser-jsonrpcerrorparser-top.md)
