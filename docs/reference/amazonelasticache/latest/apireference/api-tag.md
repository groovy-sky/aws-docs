# Tag

A tag that can be added to an ElastiCache cluster or replication group. Tags are
composed of a Key/Value pair. You can use tags to categorize and track all your
ElastiCache resources, with the exception of global replication group. When you add or
remove tags on replication groups, those actions will be replicated to all nodes in the
replication group. A tag with a null Value is permitted.

## Contents

###### Note

In the following list, the required parameters are described first.

**Key**

The key for the tag. May not be null.

Type: String

Required: No

**Value**

The tag's value. May be null.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/tag.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/tag.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/tag.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SubnetOutpost

TimeRangeFilter

All content copied from https://docs.aws.amazon.com/.
