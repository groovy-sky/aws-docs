Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)
- [ErrorParser](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.api.errorparser.html)

## JsonRpcErrorParser     extends [AbstractErrorParser](class-aws-api-errorparser-abstracterrorparser.md)   in package    - [Aws](package-aws.md)       Uses  [JsonParserTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ErrorParser.JsonParserTrait.html)

Parsers JSON-RPC errors.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ErrorParser.JsonRpcErrorParser.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ErrorParser.JsonRpcErrorParser.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ErrorParser.JsonRpcErrorParser.html#method___construct)
: mixed [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ErrorParser.JsonRpcErrorParser.html#method___invoke)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ErrorParser.JsonRpcErrorParser.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ErrorParser.JsonRpcErrorParser.html\#method___construct)

`
    public
                    __construct([Service|null $api = null ][, JsonParser|null $parser = null ]) : mixed`

##### Parameters

$api
: [Service](class-aws-api-service.md) \|null
= null$parser
: JsonParser\|null
= null

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ErrorParser.JsonRpcErrorParser.html\#method___invoke)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ErrorParser.JsonRpcErrorParser.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ErrorParser.JsonRpcErrorParser.html#method___construct)
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ErrorParser.JsonRpcErrorParser.html#method___invoke)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ErrorParser.JsonRpcErrorParser.html#top)
