Menu

- [Aws](namespace-aws.md)
- [Sns](namespace-aws-sns.md)

## MessageValidator        in package    - [Aws](package-aws.md)

This class is maintained in [aws-php-sns-message-validator](https://github.com/aws/aws-php-sns-message-validator)

Uses openssl to verify SNS messages to ensure that they were sent by AWS.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html\#toc-constants)

[SIGNATURE\_VERSION\_1](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#constant_SIGNATURE_VERSION_1)
= '1' [SIGNATURE\_VERSION\_2](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#constant_SIGNATURE_VERSION_2)
= '2'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#method___construct)
: mixed Constructs the Message Validator object and ensures that openssl is
installed.[getStringToSign()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#method_getStringToSign)
: string Builds string-to-sign according to the SNS message spec.[isValid()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#method_isValid)
: bool Determines if a message is valid and that is was delivered by AWS. This
method does not throw exceptions and returns a simple boolean value.[validate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#method_validate)
: mixed Validates a message from SNS to ensure that it was delivered by AWS.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html\#constants)

#### SIGNATURE\_VERSION\_1  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html\#constant_SIGNATURE_VERSION_1)

`
    public
        mixed
    SIGNATURE_VERSION_1
    = '1'
`

#### SIGNATURE\_VERSION\_2  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html\#constant_SIGNATURE_VERSION_2)

`
    public
        mixed
    SIGNATURE_VERSION_2
    = '2'
`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html\#method___construct)

Constructs the Message Validator object and ensures that openssl is
installed.

`
    public
                    __construct([callable $certClient = null ][, string $hostNamePattern = '' ]) : mixed`

##### Parameters

$certClient
: callable
= null

Callable used to download the certificate.
Should have the following function signature:
`function (string $certUrl) : string|false $certContent`

$hostNamePattern
: string
= ''

#### getStringToSign()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html\#method_getStringToSign)

Builds string-to-sign according to the SNS message spec.

`
    public
                    getStringToSign(Message $message) : string`

##### Parameters

$message
: [Message](class-aws-sns-message.md)

Message for which to build the string-to-sign.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html\#method_getStringToSign\#tags)

link[http://docs.aws.amazon.com/sns/latest/gsg/SendMessageToHttp.verify.signature.html](https://docs.aws.amazon.com/sns/latest/gsg/SendMessageToHttp.verify.signature.html)

##### Return values

string

#### isValid()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html\#method_isValid)

Determines if a message is valid and that is was delivered by AWS. This
method does not throw exceptions and returns a simple boolean value.

`
    public
                    isValid(Message $message) : bool`

##### Parameters

$message
: [Message](class-aws-sns-message.md)

The message to validate

##### Return values

bool

#### validate()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html\#method_validate)

Validates a message from SNS to ensure that it was delivered by AWS.

`
    public
                    validate(Message $message) : mixed`

##### Parameters

$message
: [Message](class-aws-sns-message.md)

Message to validate.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html\#method_validate\#tags)

throws[InvalidSnsMessageException](class-aws-sns-exception-invalidsnsmessageexception.md)

If the cert cannot be retrieved or its
source verified, or the message
signature is invalid.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#toc-methods)
- Constants
  - [SIGNATURE\_VERSION\_1](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#constant_SIGNATURE_VERSION_1)
  - [SIGNATURE\_VERSION\_2](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#constant_SIGNATURE_VERSION_2)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#method___construct)
  - [getStringToSign()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#method_getStringToSign)
  - [isValid()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#method_isValid)
  - [validate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#method_validate)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Sns.MessageValidator.html#top)
