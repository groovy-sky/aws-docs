Menu

- [Aws](namespace-aws.md)

## S3

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.html\#toc)

#### Namespaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.html\#namespaces)

[Crypto](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.crypto.html)[Exception](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.exception.html)[Parser](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.parser.html)[RegionalEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.regionalendpoint.html)[S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)[UseArnRegion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.usearnregion.html)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.html\#toc-interfaces)

[S3ClientInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html)\*\*Amazon Simple Storage Service\*\* client.

#### Classes  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.html\#toc-classes)

[BatchDelete](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.BatchDelete.html)Efficiently deletes many objects from a single Amazon S3 bucket using an
iterator that yields keys. Deletes are made using the DeleteObjects API
operation.[MultipartCopy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartCopy.html)[MultipartUploader](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploader.html)Encapsulates the execution of a multipart upload to S3 or Glacier.[ObjectCopier](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectCopier.html)Copies objects from one S3 location to another, utilizing a multipart copy
when appropriate.[ObjectUploader](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.ObjectUploader.html)Uploads an object to S3, using a PutObject command or a multipart upload as
appropriate.[PostObject](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObject.html)[PostObjectV4](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.PostObjectV4.html)Encapsulates the logic for getting the data for an S3 object POST upload form[S3Client](class-aws-s3-s3client.md)Client used to interact with \*\*Amazon Simple Storage Service (Amazon S3)\*\*.[S3MultiRegionClient](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3MultiRegionClient.html)\*\*Amazon Simple Storage Service\*\* multi-region client.[S3UriParser](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3UriParser.html)Extracts a region, bucket, key, and and if a URI is in path-style[StreamWrapper](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.StreamWrapper.html)Amazon S3 stream wrapper to use "s3://<bucket>/<key>" files with PHP
streams, supporting "r", "w", "a", "x".[Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.Transfer.html)Transfers files from the local filesystem to S3 or from S3 to the local
filesystem.

#### Traits  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.html\#toc-traits)

[CalculatesChecksumTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.CalculatesChecksumTrait.html)[MultipartUploadingTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.MultipartUploadingTrait.html)[S3ClientTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html)A trait providing S3-specific functionality. This is meant to be used in
classes implementing \\Aws\\S3\\S3ClientInterface

```

```

×

**On this page**

- Table Of Contents
  - [Interfaces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.html#toc-interfaces)
  - [Classes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.html#toc-classes)
  - [Traits](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.html#toc-traits)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.html#top)
