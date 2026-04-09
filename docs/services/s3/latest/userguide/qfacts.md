# Amazon S3 multipart upload limits

Multipart upload allows you to upload a single object as a set of parts. Each part is a
contiguous portion of the object's data. After all parts of your object are uploaded, Amazon S3
assembles these parts and creates the object. In general, when your object size reaches 100
MB, you should consider using multipart uploads instead of uploading the object in a single
operation. For more information about multipart uploads, see [Uploading and copying objects using multipart upload in Amazon S3](mpuoverview.md).

The following table provides multipart upload core specifications. These include maximum
object size, maximum number of parts, maximum part size, and more. There is no minimum size
limit on the last part of your multipart upload.

ItemSpecificationMaximum object size48.8 TiB Maximum number of parts per upload10,000Part numbers1 to 10,000 (inclusive)Part size5 MiB to 5 GiB. There is no minimum size limit on the last part of your multipart
upload.Maximum number of parts returned for a list parts request1000 Maximum number of multipart uploads returned in a list multipart uploads
request1000

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Upload an object through multipart upload and
verify its data integrity

Making conditional requests

All content copied from https://docs.aws.amazon.com/.
