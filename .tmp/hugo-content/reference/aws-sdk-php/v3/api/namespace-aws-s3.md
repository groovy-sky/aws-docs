Menu

- [Aws](namespace-aws.md)

## S3

### Table of Contents  [header link](namespace-aws-s3-toc.md)

#### Namespaces  [header link](namespace-aws-s3-namespaces.md)

[Crypto](namespace-aws-s3-crypto.md)[Exception](namespace-aws-s3-exception.md)[Parser](namespace-aws-s3-parser.md)[RegionalEndpoint](namespace-aws-s3-regionalendpoint.md)[S3Transfer](namespace-aws-s3-s3transfer.md)[UseArnRegion](namespace-aws-s3-usearnregion.md)

#### Interfaces  [header link](namespace-aws-s3-toc-interfaces.md)

[S3ClientInterface](class-aws-s3-s3clientinterface.md)\*\*Amazon Simple Storage Service\*\* client.

#### Classes  [header link](namespace-aws-s3-toc-classes.md)

[BatchDelete](class-aws-s3-batchdelete.md)Efficiently deletes many objects from a single Amazon S3 bucket using an
iterator that yields keys. Deletes are made using the DeleteObjects API
operation.[MultipartCopy](class-aws-s3-multipartcopy.md)[MultipartUploader](class-aws-s3-multipartuploader.md)Encapsulates the execution of a multipart upload to S3 or Glacier.[ObjectCopier](class-aws-s3-objectcopier.md)Copies objects from one S3 location to another, utilizing a multipart copy
when appropriate.[ObjectUploader](class-aws-s3-objectuploader.md)Uploads an object to S3, using a PutObject command or a multipart upload as
appropriate.[PostObject](class-aws-s3-postobject.md)[PostObjectV4](class-aws-s3-postobjectv4.md)Encapsulates the logic for getting the data for an S3 object POST upload form[S3Client](class-aws-s3-s3client.md)Client used to interact with \*\*Amazon Simple Storage Service (Amazon S3)\*\*.[S3MultiRegionClient](class-aws-s3-s3multiregionclient.md)\*\*Amazon Simple Storage Service\*\* multi-region client.[S3UriParser](class-aws-s3-s3uriparser.md)Extracts a region, bucket, key, and and if a URI is in path-style[StreamWrapper](class-aws-s3-streamwrapper.md)Amazon S3 stream wrapper to use "s3://<bucket>/<key>" files with PHP
streams, supporting "r", "w", "a", "x".[Transfer](class-aws-s3-transfer.md)Transfers files from the local filesystem to S3 or from S3 to the local
filesystem.

#### Traits  [header link](namespace-aws-s3-toc-traits.md)

[CalculatesChecksumTrait](class-aws-s3-calculateschecksumtrait.md)[MultipartUploadingTrait](class-aws-s3-multipartuploadingtrait.md)[S3ClientTrait](class-aws-s3-s3clienttrait.md)A trait providing S3-specific functionality. This is meant to be used in
classes implementing \\Aws\\S3\\S3ClientInterface

```

```

×

**On this page**

- Table Of Contents
  - [Interfaces](namespace-aws-s3-toc-interfaces.md)
  - [Classes](namespace-aws-s3-toc-classes.md)
  - [Traits](namespace-aws-s3-toc-traits.md)

[Back To Top](namespace-aws-s3-top.md)
