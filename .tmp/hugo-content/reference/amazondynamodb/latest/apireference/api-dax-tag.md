# Tag

A description of a tag. Every tag is a key-value pair. You can add up to 50 tags to
a single DAX cluster.

AWS-assigned tag names and values are automatically assigned the
`aws:` prefix, which the user cannot assign. AWS-assigned
tag names do not count towards the tag limit of 50. User-assigned tag names have the
prefix `user:`.

You cannot backdate the application of a tag.

## Contents

###### Note

In the following list, the required parameters are described first.

**Key**

The key for the tag. Tag keys are case sensitive. Every DAX cluster
can only have one tag with the same key. If you try to add an existing tag (same key),
the existing tag value will be updated to the new value.

Type: String

Required: No

**Value**

The value of the tag. Tag values are case-sensitive and can be null.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/tag.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/tag.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/tag.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SubnetGroup

Amazon DynamoDB Streams

All content copied from https://docs.aws.amazon.com/.
