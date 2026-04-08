Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## PostObjectV4        in package    - [Aws](package-aws.md)       Uses  [SignatureTrait](class-aws-signature-signaturetrait.md)

Encapsulates the logic for getting the data for an S3 object POST upload form

##### Tags  [header link](class-aws-s3-postobjectv4-tags.md)

link[http://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectPOST.html](../../../../services/s3/latest/api/restobjectpost.md)link[http://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-post-example.html](../../../../services/s3/latest/api/sigv4-post-example.md)

### Table of Contents  [header link](class-aws-s3-postobjectv4-toc.md)

#### Methods  [header link](class-aws-s3-postobjectv4-toc-methods.md)

[\_\_construct()](class-aws-s3-postobjectv4-method-construct.md)
: mixed Constructs the PostObject.[getBucket()](class-aws-s3-postobjectv4-method-getbucket.md)
: string Gets the bucket name.[getClient()](class-aws-s3-postobjectv4-method-getclient.md)
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)Gets the S3 client.[getFormAttributes()](class-aws-s3-postobjectv4-method-getformattributes.md)
: array<string\|int, mixed> Gets the form attributes as an array.[getFormInputs()](class-aws-s3-postobjectv4-method-getforminputs.md)
: array<string\|int, mixed> Gets the form inputs as an array.[setFormAttribute()](class-aws-s3-postobjectv4-method-setformattribute.md)
: mixed Set a form attribute.[setFormInput()](class-aws-s3-postobjectv4-method-setforminput.md)
: mixed Set a form input.

### Methods  [header link](class-aws-s3-postobjectv4-methods.md)

#### \_\_construct()  [header link](class-aws-s3-postobjectv4-method-construct.md)

Constructs the PostObject.

`
    public
                    __construct(S3ClientInterface $client, string $bucket, array<string|int, mixed> $formInputs[, array<string|int, mixed> $options = [] ][, mixed $expiration = '+1 hours' ]) : mixed`

The options array accepts the following keys:

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

$options
: array<string\|int, mixed>
= \[\]

Policy condition options

$expiration
: mixed
= '+1 hours'

Upload expiration time value. By
default: 1 hour valid period.

##### Tags  [header link](class-aws-s3-postobjectv4-method-construct-tags.md)

link[http://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html](../../../../services/s3/latest/api/sigv4-query-string-auth.md)

#### getBucket()  [header link](class-aws-s3-postobjectv4-method-getbucket.md)

Gets the bucket name.

`
    public
                    getBucket() : string`

##### Return values

string

#### getClient()  [header link](class-aws-s3-postobjectv4-method-getclient.md)

Gets the S3 client.

`
    public
                    getClient() : S3ClientInterface`

##### Return values

[S3ClientInterface](class-aws-s3-s3clientinterface.md)

#### getFormAttributes()  [header link](class-aws-s3-postobjectv4-method-getformattributes.md)

Gets the form attributes as an array.

`
    public
                    getFormAttributes() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getFormInputs()  [header link](class-aws-s3-postobjectv4-method-getforminputs.md)

Gets the form inputs as an array.

`
    public
                    getFormInputs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### setFormAttribute()  [header link](class-aws-s3-postobjectv4-method-setformattribute.md)

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

#### setFormInput()  [header link](class-aws-s3-postobjectv4-method-setforminput.md)

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
  - [Methods](class-aws-s3-postobjectv4-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-postobjectv4-method-construct.md)
  - [getBucket()](class-aws-s3-postobjectv4-method-getbucket.md)
  - [getClient()](class-aws-s3-postobjectv4-method-getclient.md)
  - [getFormAttributes()](class-aws-s3-postobjectv4-method-getformattributes.md)
  - [getFormInputs()](class-aws-s3-postobjectv4-method-getforminputs.md)
  - [setFormAttribute()](class-aws-s3-postobjectv4-method-setformattribute.md)
  - [setFormInput()](class-aws-s3-postobjectv4-method-setforminput.md)

[Back To Top](class-aws-s3-postobjectv4-top.md)
