# DeletedObject

Information about the deleted object.

## Contents

**DeleteMarker**

Indicates whether the specified object version that was permanently deleted was (true) or was not
(false) a delete marker before deletion. In a simple DELETE, this header indicates whether (true) or not
(false) the current version of the object is a delete marker. To learn more about delete markers, see
[Working with delete\
markers](../userguide/deletemarker.md).

###### Note

This functionality is not supported for directory buckets.

Type: Boolean

Required: No

**DeleteMarkerVersionId**

The version ID of the delete marker created as a result of the DELETE operation. If you delete a
specific object version, the value returned by this header is the version ID of the object version
deleted.

###### Note

This functionality is not supported for directory buckets.

Type: String

Required: No

**Key**

The name of the deleted object.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**VersionId**

The version ID of the deleted object.

###### Note

This functionality is not supported for directory buckets.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/deletedobject.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/deletedobject.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/deletedobject.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Delete

DeleteMarkerEntry

All content copied from https://docs.aws.amazon.com/.
