Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## S3UriParser        in package    - [Aws](package-aws.md)

Extracts a region, bucket, key, and and if a URI is in path-style

### Table of Contents  [header link](class-aws-s3-s3uriparser-toc.md)

#### Methods  [header link](class-aws-s3-s3uriparser-toc-methods.md)

[parse()](class-aws-s3-s3uriparser-method-parse.md)
: array<string\|int, mixed> Parses a URL or S3 StreamWrapper Uri (s3://) into an associative array
of Amazon S3 data including:

### Methods  [header link](class-aws-s3-s3uriparser-methods.md)

#### parse()  [header link](class-aws-s3-s3uriparser-method-parse.md)

Parses a URL or S3 StreamWrapper Uri (s3://) into an associative array
of Amazon S3 data including:

`
    public
                    parse(string|UriInterface $uri) : array<string|int, mixed>`

- bucket: The Amazon S3 bucket (null if none)
- key: The Amazon S3 key (null if none)
- path\_style: Set to true if using path style, or false if not
- region: Set to a string if a non-class endpoint is used or null.

##### Parameters

$uri
: string\| [UriInterface](class-psr-http-message-uriinterface.md)

##### Tags  [header link](class-aws-s3-s3uriparser-method-parse-tags.md)

throwsInvalidArgumentException\| [InvalidArnException](class-aws-arn-exception-invalidarnexception.md)

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-s3-s3uriparser-toc-methods.md)
- Methods
  - [parse()](class-aws-s3-s3uriparser-method-parse.md)

[Back To Top](class-aws-s3-s3uriparser-top.md)
