Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## Transfer        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

Transfers files from the local filesystem to S3 or from S3 to the local
filesystem.

This class does not support copying from the local filesystem to somewhere
else on the local filesystem or from one S3 bucket to another.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html\#toc-interfaces)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html#method___construct)
: mixed When providing the $source argument, you may provide a string referencing
the path to a directory on disk to upload, an s3 scheme URI that contains
the bucket and key (e.g., "s3://bucket/key"), or an \\Iterator object
that yields strings containing filenames that are the path to a file on
disk or an s3 scheme URI. The bucket portion of the s3 URI may be an S3
access point ARN. The "/key" portion of an s3 URI is optional.[promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html#method_promise)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Transfers the files.[transfer()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html#method_transfer)
: mixed Transfers the files synchronously.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html\#method___construct)

When providing the $source argument, you may provide a string referencing
the path to a directory on disk to upload, an s3 scheme URI that contains
the bucket and key (e.g., "s3://bucket/key"), or an \\Iterator object
that yields strings containing filenames that are the path to a file on
disk or an s3 scheme URI. The bucket portion of the s3 URI may be an S3
access point ARN. The "/key" portion of an s3 URI is optional.

`
    public
                    __construct(S3ClientInterface $client, string|Iterator $source, string $dest[, array<string|int, mixed> $options = [] ]) : mixed`

When providing an iterator for the $source argument, you must also
provide a 'base\_dir' key value pair in the $options argument.

The $dest argument can be the path to a directory on disk or an s3
scheme URI (e.g., "s3://bucket/key").

The options array can contain the following key value pairs:

- base\_dir: (string) Base directory of the source, if $source is an
iterator. If the $source option is not an array, then this option is
ignored.
- before: (callable) A callback to invoke before each transfer. The
callback accepts a single argument: Aws\\CommandInterface $command.
The provided command will be either a GetObject, PutObject,
InitiateMultipartUpload, or UploadPart command.
- after: (callable) A callback to invoke after each transfer promise is fulfilled.
The function is invoked with three arguments: the fulfillment value, the index
position from the iterable list of the promise, and the aggregate
promise that manages all the promises. The aggregate promise may
be resolved from within the callback to short-circuit the promise.
- mup\_threshold: (int) Size in bytes in which a multipart upload should
be used instead of PutObject. Defaults to 20971520 (20 MB).
- concurrency: (int, default=5) Number of files to upload concurrently.
The ideal concurrency value will vary based on the number of files
being uploaded and the average size of each file. Generally speaking,
smaller files benefit from a higher concurrency while larger files
will not.
- debug: (bool) Set to true to print out debug information for
transfers. Set to an fopen() resource to write to a specific stream
rather than writing to STDOUT.

##### Parameters

$client
: [S3ClientInterface](class-aws-s3-s3clientinterface.md)

Client used for transfers.

$source
: string\|Iterator

Where the files are transferred from.

$dest
: string

Where the files are transferred to.

$options
: array<string\|int, mixed>
= \[\]

Hash of options.

#### promise()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html\#method_promise)

Transfers the files.

`
    public
                    promise() : PromiseInterface`

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### transfer()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html\#method_transfer)

Transfers the files synchronously.

`
    public
                    transfer() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html#method___construct)
  - [promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html#method_promise)
  - [transfer()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html#method_transfer)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html#top)
