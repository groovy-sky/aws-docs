Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## SupportedProtocols     : string     in package    - [Aws](package-aws.md)

Priority ordered collection of supported AWS protocols.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html\#toc)

#### Cases  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html\#toc-cases)

[CBOR](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#enumcase_CBOR)
= 'smithy-rpc-v2-cbor' [EC2](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#enumcase_EC2)
= 'ec2' [JSON](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#enumcase_JSON)
= 'json' [QUERY](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#enumcase_QUERY)
= 'query' [REST\_JSON](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#enumcase_REST_JSON)
= 'rest-json' [REST\_XML](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#enumcase_REST_XML)
= 'rest-xml'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html\#toc-methods)

[isSupported()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#method_isSupported)
: bool Check if a protocol is valid.

### Cases  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html\#cases)

#### JSON  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html\#enumcase_JSON)

#### CBOR  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html\#enumcase_CBOR)

#### REST\_JSON  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html\#enumcase_REST_JSON)

#### REST\_XML  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html\#enumcase_REST_XML)

#### QUERY  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html\#enumcase_QUERY)

#### EC2  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html\#enumcase_EC2)

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html\#methods)

#### isSupported()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html\#method_isSupported)

Check if a protocol is valid.

`
    public
            static        isSupported(string $protocol) : bool`

##### Parameters

$protocol
: string

##### Return values

bool
—

True if the protocol is supported, otherwise false.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Cases](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#toc-cases)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#toc-methods)
- Cases
  - [JSON](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#enumcase_JSON)
  - [CBOR](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#enumcase_CBOR)
  - [REST\_JSON](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#enumcase_REST_JSON)
  - [REST\_XML](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#enumcase_REST_XML)
  - [QUERY](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#enumcase_QUERY)
  - [EC2](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#enumcase_EC2)
- Methods
  - [isSupported()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#method_isSupported)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.SupportedProtocols.html#top)
