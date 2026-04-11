Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## PostObject        in package    - [Aws](package-aws.md)

##### Tags  [header link](class-aws-s3-postobject-tags.md)

deprecated

### Table of Contents  [header link](class-aws-s3-postobject-toc.md)

#### Methods  [header link](class-aws-s3-postobject-toc-methods.md)

[\_\_construct()](class-aws-s3-postobject-method-construct.md)
: mixed Constructs the PostObject.[getBucket()](class-aws-s3-postobject-method-getbucket.md)
: string Gets the bucket name.[getClient()](class-aws-s3-postobject-method-getclient.md)
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)Gets the S3 client.[getFormAttributes()](class-aws-s3-postobject-method-getformattributes.md)
: array<string\|int, mixed> Gets the form attributes as an array.[getFormInputs()](class-aws-s3-postobject-method-getforminputs.md)
: array<string\|int, mixed> Gets the form inputs as an array.[getJsonPolicy()](class-aws-s3-postobject-method-getjsonpolicy.md)
: string Gets the raw JSON policy.[setFormAttribute()](class-aws-s3-postobject-method-setformattribute.md)
: mixed Set a form attribute.[setFormInput()](class-aws-s3-postobject-method-setforminput.md)
: mixed Set a form input.

### Methods  [header link](class-aws-s3-postobject-methods.md)

#### \_\_construct()  [header link](class-aws-s3-postobject-method-construct.md)

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

#### getBucket()  [header link](class-aws-s3-postobject-method-getbucket.md)

Gets the bucket name.

`
    public
                    getBucket() : string`

##### Return values

string

#### getClient()  [header link](class-aws-s3-postobject-method-getclient.md)

Gets the S3 client.

`
    public
                    getClient() : S3ClientInterface`

##### Return values

[S3ClientInterface](class-aws-s3-s3clientinterface.md)

#### getFormAttributes()  [header link](class-aws-s3-postobject-method-getformattributes.md)

Gets the form attributes as an array.

`
    public
                    getFormAttributes() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getFormInputs()  [header link](class-aws-s3-postobject-method-getforminputs.md)

Gets the form inputs as an array.

`
    public
                    getFormInputs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getJsonPolicy()  [header link](class-aws-s3-postobject-method-getjsonpolicy.md)

Gets the raw JSON policy.

`
    public
                    getJsonPolicy() : string`

##### Return values

string

#### setFormAttribute()  [header link](class-aws-s3-postobject-method-setformattribute.md)

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

#### setFormInput()  [header link](class-aws-s3-postobject-method-setforminput.md)

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
  - [Methods](class-aws-s3-postobject-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-postobject-method-construct.md)
  - [getBucket()](class-aws-s3-postobject-method-getbucket.md)
  - [getClient()](class-aws-s3-postobject-method-getclient.md)
  - [getFormAttributes()](class-aws-s3-postobject-method-getformattributes.md)
  - [getFormInputs()](class-aws-s3-postobject-method-getforminputs.md)
  - [getJsonPolicy()](class-aws-s3-postobject-method-getjsonpolicy.md)
  - [setFormAttribute()](class-aws-s3-postobject-method-setformattribute.md)
  - [setFormInput()](class-aws-s3-postobject-method-setforminput.md)

[Back To Top](class-aws-s3-postobject-top.md)
