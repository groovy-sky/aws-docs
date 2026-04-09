# ReplicationRuleAndOperator

A container for specifying rule filters. The filters determine the subset of objects to
which the rule applies. This element is required only if you specify more than one filter.

For example:

- If you specify both a `Prefix` and a `Tag` filter, wrap
these filters in an `And` element.

- If you specify a filter based on multiple tags, wrap the `Tag` elements
in an `And` element.

## Contents

**Prefix**

An object key name prefix that identifies the subset of objects that the rule applies
to.

Type: String

Required: No

**Tags**

An array of tags that contain key and value pairs.

Type: Array of [S3Tag](api-control-s3tag.md) data types

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/replicationruleandoperator.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/replicationruleandoperator.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/replicationruleandoperator.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationRule

ReplicationRuleFilter

All content copied from https://docs.aws.amazon.com/.
