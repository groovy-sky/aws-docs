Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)
- [ErrorParser](namespace-aws-api-errorparser.md)

## RestJsonErrorParser     extends [AbstractErrorParser](class-aws-api-errorparser-abstracterrorparser.md)   in package    - [Aws](package-aws.md)       Uses  [JsonParserTrait](class-aws-api-errorparser-jsonparsertrait.md)

Parses JSON-REST errors.

### Table of Contents  [header link](class-aws-api-errorparser-restjsonerrorparser-toc.md)

#### Methods  [header link](class-aws-api-errorparser-restjsonerrorparser-toc-methods.md)

[\_\_construct()](class-aws-api-errorparser-restjsonerrorparser-method-construct.md)
: mixed [\_\_invoke()](class-aws-api-errorparser-restjsonerrorparser-method-invoke.md)
: mixed

### Methods  [header link](class-aws-api-errorparser-restjsonerrorparser-methods.md)

#### \_\_construct()  [header link](class-aws-api-errorparser-restjsonerrorparser-method-construct.md)

`
    public
                    __construct([Service|null $api = null ][, JsonParser|null $parser = null ]) : mixed`

##### Parameters

$api
: [Service](class-aws-api-service.md) \|null
= null$parser
: JsonParser\|null
= null

#### \_\_invoke()  [header link](class-aws-api-errorparser-restjsonerrorparser-method-invoke.md)

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
  - [Methods](class-aws-api-errorparser-restjsonerrorparser-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-api-errorparser-restjsonerrorparser-method-construct.md)
  - [\_\_invoke()](class-aws-api-errorparser-restjsonerrorparser-method-invoke.md)

[Back To Top](class-aws-api-errorparser-restjsonerrorparser-top.md)
