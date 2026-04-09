# FilterRule

Specifies the Amazon S3 object key name to filter on. An object key name is the name assigned to an
object in your Amazon S3 bucket. You specify whether to filter on the suffix or prefix of the object key
name. A prefix is a specific string of characters at the beginning of an object key name, which you can
use to organize objects. For example, you can start the key names of related objects with a prefix, such
as `2023-` or `engineering/`. Then, you can use `FilterRule` to find
objects in a bucket with key names that have the same prefix. A suffix is similar to a prefix, but it is
at the end of the object key name instead of at the beginning.

## Contents

**Name**

The object key name prefix or suffix identifying one or more objects to which the filtering rule
applies. The maximum length is 1,024 characters. Overlapping prefixes and suffixes are not supported.
For more information, see [Configuring Event Notifications](../dev/notificationhowto.md) in the _Amazon S3 User Guide_.

Type: String

Valid Values: `prefix | suffix`

Required: No

**Value**

The value that the filter searches for in object key names.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/filterrule.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/filterrule.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/filterrule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExistingObjectReplication

GetBucketMetadataConfigurationResult

All content copied from https://docs.aws.amazon.com/.
