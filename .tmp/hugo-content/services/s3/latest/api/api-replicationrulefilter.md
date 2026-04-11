# ReplicationRuleFilter

A filter that identifies the subset of objects to which the replication rule applies. A
`Filter` must specify exactly one `Prefix`, `Tag`, or an
`And` child element.

## Contents

**And**

A container for specifying rule filters. The filters determine the subset of objects to which the
rule applies. This element is required only if you specify more than one filter. For example:

- If you specify both a `Prefix` and a `Tag` filter, wrap these filters in
an `And` tag.

- If you specify a filter based on multiple tags, wrap the `Tag` elements in an
`And` tag.

Type: [ReplicationRuleAndOperator](api-replicationruleandoperator.md) data type

Required: No

**Prefix**

An object key name prefix that identifies the subset of objects to which the rule applies.

###### Important

Replacement must be made for object keys containing special characters (such as carriage returns) when using
XML requests. For more information, see [XML related object key constraints](../userguide/object-keys.md#object-key-xml-related-constraints).

Type: String

Required: No

**Tag**

A container for specifying a tag key and value.

The rule applies only to objects that have the tag in their tag set.

Type: [Tag](api-tag.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/replicationrulefilter.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/replicationrulefilter.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/replicationrulefilter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationRuleAndOperator

ReplicationTime

All content copied from https://docs.aws.amazon.com/.
