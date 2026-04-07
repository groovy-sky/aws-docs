Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## DocModel        in package    - [Aws](package-aws.md)

Encapsulates the documentation strings for a given service-version and
provides methods for extracting the desired parts related to a service,
operation, error, or shape (i.e., parameter).

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html#method___construct)
: mixed [getErrorDocs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html#method_getErrorDocs)
: null\|string Retrieves documentation about an error.[getOperationDocs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html#method_getOperationDocs)
: null\|string Retrieves documentation about an operation.[getServiceDocs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html#method_getServiceDocs)
: null\|string Retrieves documentation about the service.[getShapeDocs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html#method_getShapeDocs)
: null\|string Retrieves documentation about a shape, specific to the context.[toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html#method_toArray)
: array<string\|int, mixed> Convert the doc model to an array.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html\#method___construct)

`
    public
                    __construct(array<string|int, mixed> $docs) : mixed`

##### Parameters

$docs
: array<string\|int, mixed>

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html\#method___construct\#tags)

throwsRuntimeException

#### getErrorDocs()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html\#method_getErrorDocs)

Retrieves documentation about an error.

`
    public
                    getErrorDocs(string $error) : null|string`

##### Parameters

$error
: string

Name of the error

##### Return values

null\|string

#### getOperationDocs()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html\#method_getOperationDocs)

Retrieves documentation about an operation.

`
    public
                    getOperationDocs(string $operation) : null|string`

##### Parameters

$operation
: string

Name of the operation

##### Return values

null\|string

#### getServiceDocs()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html\#method_getServiceDocs)

Retrieves documentation about the service.

`
    public
                    getServiceDocs() : null|string`

##### Return values

null\|string

#### getShapeDocs()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html\#method_getShapeDocs)

Retrieves documentation about a shape, specific to the context.

`
    public
                    getShapeDocs(string $shapeName, string $parentName, string $ref) : null|string`

##### Parameters

$shapeName
: string

Name of the shape.

$parentName
: string

Name of the parent/context shape.

$ref
: string

Name used by the context to reference the shape.

##### Return values

null\|string

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html\#method_toArray)

Convert the doc model to an array.

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html#method___construct)
  - [getErrorDocs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html#method_getErrorDocs)
  - [getOperationDocs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html#method_getOperationDocs)
  - [getServiceDocs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html#method_getServiceDocs)
  - [getShapeDocs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html#method_getShapeDocs)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DocModel.html#top)
