Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## SupportedProtocols     : string     in package    - [Aws](package-aws.md)

Priority ordered collection of supported AWS protocols.

### Table of Contents  [header link](class-aws-api-supportedprotocols-toc.md)

#### Cases  [header link](class-aws-api-supportedprotocols-toc-cases.md)

[CBOR](class-aws-api-supportedprotocols-enumcase-cbor.md)
= 'smithy-rpc-v2-cbor' [EC2](class-aws-api-supportedprotocols-enumcase-ec2.md)
= 'ec2' [JSON](class-aws-api-supportedprotocols-enumcase-json.md)
= 'json' [QUERY](class-aws-api-supportedprotocols-enumcase-query.md)
= 'query' [REST\_JSON](class-aws-api-supportedprotocols-enumcase-rest-json.md)
= 'rest-json' [REST\_XML](class-aws-api-supportedprotocols-enumcase-rest-xml.md)
= 'rest-xml'

#### Methods  [header link](class-aws-api-supportedprotocols-toc-methods.md)

[isSupported()](class-aws-api-supportedprotocols-method-issupported.md)
: bool Check if a protocol is valid.

### Cases  [header link](class-aws-api-supportedprotocols-cases.md)

#### JSON  [header link](class-aws-api-supportedprotocols-enumcase-json.md)

#### CBOR  [header link](class-aws-api-supportedprotocols-enumcase-cbor.md)

#### REST\_JSON  [header link](class-aws-api-supportedprotocols-enumcase-rest-json.md)

#### REST\_XML  [header link](class-aws-api-supportedprotocols-enumcase-rest-xml.md)

#### QUERY  [header link](class-aws-api-supportedprotocols-enumcase-query.md)

#### EC2  [header link](class-aws-api-supportedprotocols-enumcase-ec2.md)

### Methods  [header link](class-aws-api-supportedprotocols-methods.md)

#### isSupported()  [header link](class-aws-api-supportedprotocols-method-issupported.md)

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
  - [Cases](class-aws-api-supportedprotocols-toc-cases.md)
  - [Methods](class-aws-api-supportedprotocols-toc-methods.md)
- Cases
  - [JSON](class-aws-api-supportedprotocols-enumcase-json.md)
  - [CBOR](class-aws-api-supportedprotocols-enumcase-cbor.md)
  - [REST\_JSON](class-aws-api-supportedprotocols-enumcase-rest-json.md)
  - [REST\_XML](class-aws-api-supportedprotocols-enumcase-rest-xml.md)
  - [QUERY](class-aws-api-supportedprotocols-enumcase-query.md)
  - [EC2](class-aws-api-supportedprotocols-enumcase-ec2.md)
- Methods
  - [isSupported()](class-aws-api-supportedprotocols-method-issupported.md)

[Back To Top](class-aws-api-supportedprotocols-top.md)
