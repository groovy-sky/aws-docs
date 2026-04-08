Menu

- [Aws](namespace-aws.md)
- [Sns](namespace-aws-sns.md)

## MessageValidator        in package    - [Aws](package-aws.md)

This class is maintained in [aws-php-sns-message-validator](https://github.com/aws/aws-php-sns-message-validator)

Uses openssl to verify SNS messages to ensure that they were sent by AWS.

### Table of Contents  [header link](class-aws-sns-messagevalidator-toc.md)

#### Constants  [header link](class-aws-sns-messagevalidator-toc-constants.md)

[SIGNATURE\_VERSION\_1](class-aws-sns-messagevalidator-constant-signature-version-1.md)
= '1' [SIGNATURE\_VERSION\_2](class-aws-sns-messagevalidator-constant-signature-version-2.md)
= '2'

#### Methods  [header link](class-aws-sns-messagevalidator-toc-methods.md)

[\_\_construct()](class-aws-sns-messagevalidator-method-construct.md)
: mixed Constructs the Message Validator object and ensures that openssl is
installed.[getStringToSign()](class-aws-sns-messagevalidator-method-getstringtosign.md)
: string Builds string-to-sign according to the SNS message spec.[isValid()](class-aws-sns-messagevalidator-method-isvalid.md)
: bool Determines if a message is valid and that is was delivered by AWS. This
method does not throw exceptions and returns a simple boolean value.[validate()](class-aws-sns-messagevalidator-method-validate.md)
: mixed Validates a message from SNS to ensure that it was delivered by AWS.

### Constants  [header link](class-aws-sns-messagevalidator-constants.md)

#### SIGNATURE\_VERSION\_1  [header link](class-aws-sns-messagevalidator-constant-signature-version-1.md)

`
    public
        mixed
    SIGNATURE_VERSION_1
    = '1'
`

#### SIGNATURE\_VERSION\_2  [header link](class-aws-sns-messagevalidator-constant-signature-version-2.md)

`
    public
        mixed
    SIGNATURE_VERSION_2
    = '2'
`

### Methods  [header link](class-aws-sns-messagevalidator-methods.md)

#### \_\_construct()  [header link](class-aws-sns-messagevalidator-method-construct.md)

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

#### getStringToSign()  [header link](class-aws-sns-messagevalidator-method-getstringtosign.md)

Builds string-to-sign according to the SNS message spec.

`
    public
                    getStringToSign(Message $message) : string`

##### Parameters

$message
: [Message](class-aws-sns-message.md)

Message for which to build the string-to-sign.

##### Tags  [header link](class-aws-sns-messagevalidator-method-getstringtosign-tags.md)

link[http://docs.aws.amazon.com/sns/latest/gsg/SendMessageToHttp.verify.signature.html](../../../../services/sns/latest/gsg/sendmessagetohttp-verify-signature.md)

##### Return values

string

#### isValid()  [header link](class-aws-sns-messagevalidator-method-isvalid.md)

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

#### validate()  [header link](class-aws-sns-messagevalidator-method-validate.md)

Validates a message from SNS to ensure that it was delivered by AWS.

`
    public
                    validate(Message $message) : mixed`

##### Parameters

$message
: [Message](class-aws-sns-message.md)

Message to validate.

##### Tags  [header link](class-aws-sns-messagevalidator-method-validate-tags.md)

throws[InvalidSnsMessageException](class-aws-sns-exception-invalidsnsmessageexception.md)

If the cert cannot be retrieved or its
source verified, or the message
signature is invalid.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-sns-messagevalidator-toc-constants.md)
  - [Methods](class-aws-sns-messagevalidator-toc-methods.md)
- Constants
  - [SIGNATURE\_VERSION\_1](class-aws-sns-messagevalidator-constant-signature-version-1.md)
  - [SIGNATURE\_VERSION\_2](class-aws-sns-messagevalidator-constant-signature-version-2.md)
- Methods
  - [\_\_construct()](class-aws-sns-messagevalidator-method-construct.md)
  - [getStringToSign()](class-aws-sns-messagevalidator-method-getstringtosign.md)
  - [isValid()](class-aws-sns-messagevalidator-method-isvalid.md)
  - [validate()](class-aws-sns-messagevalidator-method-validate.md)

[Back To Top](class-aws-sns-messagevalidator-top.md)
