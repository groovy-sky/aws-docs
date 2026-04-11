# Delete

Container for the objects to delete.

## Contents

**Objects**

The object to delete.

###### Note

**Directory buckets** \- For directory buckets, an object
that's composed entirely of whitespace characters is not supported by the `DeleteObjects`
API operation. The request will receive a `400 Bad Request` error and none of the objects
in the request will be deleted.

Type: Array of [ObjectIdentifier](api-objectidentifier.md) data types

Required: Yes

**Quiet**

Element to enable quiet mode for the request. When you add this element, you must set its value to
`true`.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/delete.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/delete.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/delete.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultRetention

DeletedObject

All content copied from https://docs.aws.amazon.com/.
