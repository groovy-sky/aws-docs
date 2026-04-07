Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## PostObject        in package    - [Aws](package-aws.md)

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html\#tags)

deprecated

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method___construct)
: mixed Constructs the PostObject.[getBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method_getBucket)
: string Gets the bucket name.[getClient()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method_getClient)
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)Gets the S3 client.[getFormAttributes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method_getFormAttributes)
: array<string\|int, mixed> Gets the form attributes as an array.[getFormInputs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method_getFormInputs)
: array<string\|int, mixed> Gets the form inputs as an array.[getJsonPolicy()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method_getJsonPolicy)
: string Gets the raw JSON policy.[setFormAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method_setFormAttribute)
: mixed Set a form attribute.[setFormInput()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method_setFormInput)
: mixed Set a form input.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html\#method___construct)

Constructs the PostObject.

`
    public
                    __construct(S3ClientInterface $client, string $bucket, array<string|int, mixed> $formInputs, string|array<string|int, mixed> $jsonPolicy) : mixed`

##### Parameters

$client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)

Client used with the POST object

$bucket
: string

Bucket to use

$formInputs
: array<string\|int, mixed>

Associative array of form input
fields.

$jsonPolicy
: string\|array<string\|int, mixed>

JSON encoded POST policy document.
The policy will be base64 encoded
and applied to the form on your
behalf.

#### getBucket()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html\#method_getBucket)

Gets the bucket name.

`
    public
                    getBucket() : string`

##### Return values

string

#### getClient()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html\#method_getClient)

Gets the S3 client.

`
    public
                    getClient() : S3ClientInterface`

##### Return values

[S3ClientInterface](class-aws-s3-s3clientinterface.md)

#### getFormAttributes()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html\#method_getFormAttributes)

Gets the form attributes as an array.

`
    public
                    getFormAttributes() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getFormInputs()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html\#method_getFormInputs)

Gets the form inputs as an array.

`
    public
                    getFormInputs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getJsonPolicy()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html\#method_getJsonPolicy)

Gets the raw JSON policy.

`
    public
                    getJsonPolicy() : string`

##### Return values

string

#### setFormAttribute()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html\#method_setFormAttribute)

Set a form attribute.

`
    public
                    setFormAttribute(string $attribute, string $value) : mixed`

##### Parameters

$attribute
: string

Form attribute to set.

$value
: string

Value to set.

#### setFormInput()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html\#method_setFormInput)

Set a form input.

`
    public
                    setFormInput(string $field, string $value) : mixed`

##### Parameters

$field
: string

Field name to set

$value
: string

Value to set.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method___construct)
  - [getBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method_getBucket)
  - [getClient()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method_getClient)
  - [getFormAttributes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method_getFormAttributes)
  - [getFormInputs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method_getFormInputs)
  - [getJsonPolicy()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method_getJsonPolicy)
  - [setFormAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method_setFormAttribute)
  - [setFormInput()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#method_setFormInput)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html#top)
