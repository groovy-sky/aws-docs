Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## DocModel        in package    - [Aws](package-aws.md)

Encapsulates the documentation strings for a given service-version and
provides methods for extracting the desired parts related to a service,
operation, error, or shape (i.e., parameter).

### Table of Contents  [header link](class-aws-api-docmodel-toc.md)

#### Methods  [header link](class-aws-api-docmodel-toc-methods.md)

[\_\_construct()](class-aws-api-docmodel-method-construct.md)
: mixed [getErrorDocs()](class-aws-api-docmodel-method-geterrordocs.md)
: null\|string Retrieves documentation about an error.[getOperationDocs()](class-aws-api-docmodel-method-getoperationdocs.md)
: null\|string Retrieves documentation about an operation.[getServiceDocs()](class-aws-api-docmodel-method-getservicedocs.md)
: null\|string Retrieves documentation about the service.[getShapeDocs()](class-aws-api-docmodel-method-getshapedocs.md)
: null\|string Retrieves documentation about a shape, specific to the context.[toArray()](class-aws-api-docmodel-method-toarray.md)
: array<string\|int, mixed> Convert the doc model to an array.

### Methods  [header link](class-aws-api-docmodel-methods.md)

#### \_\_construct()  [header link](class-aws-api-docmodel-method-construct.md)

`
    public
                    __construct(array<string|int, mixed> $docs) : mixed`

##### Parameters

$docs
: array<string\|int, mixed>

##### Tags  [header link](class-aws-api-docmodel-method-construct-tags.md)

throwsRuntimeException

#### getErrorDocs()  [header link](class-aws-api-docmodel-method-geterrordocs.md)

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

#### getOperationDocs()  [header link](class-aws-api-docmodel-method-getoperationdocs.md)

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

#### getServiceDocs()  [header link](class-aws-api-docmodel-method-getservicedocs.md)

Retrieves documentation about the service.

`
    public
                    getServiceDocs() : null|string`

##### Return values

null\|string

#### getShapeDocs()  [header link](class-aws-api-docmodel-method-getshapedocs.md)

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

#### toArray()  [header link](class-aws-api-docmodel-method-toarray.md)

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
  - [Methods](class-aws-api-docmodel-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-api-docmodel-method-construct.md)
  - [getErrorDocs()](class-aws-api-docmodel-method-geterrordocs.md)
  - [getOperationDocs()](class-aws-api-docmodel-method-getoperationdocs.md)
  - [getServiceDocs()](class-aws-api-docmodel-method-getservicedocs.md)
  - [getShapeDocs()](class-aws-api-docmodel-method-getshapedocs.md)
  - [toArray()](class-aws-api-docmodel-method-toarray.md)

[Back To Top](class-aws-api-docmodel-top.md)
