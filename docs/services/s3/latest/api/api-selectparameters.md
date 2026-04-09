# SelectParameters

###### Important

Amazon S3 Select is no longer available to new customers. Existing customers of Amazon S3 Select can
continue to use the feature as usual. [Learn more](http://aws.amazon.com/blogs/storage/how-to-optimize-querying-your-data-in-amazon-s3)

Describes the parameters for Select job types.

Learn [How to\
optimize querying your data in Amazon S3](http://aws.amazon.com/blogs/storage/how-to-optimize-querying-your-data-in-amazon-s3) using [Amazon Athena](../../../athena/latest/ug/what-is.md), [S3 Object Lambda](../userguide/transforming-objects.md), or client-side
filtering.

## Contents

**Expression**

###### Important

Amazon S3 Select is no longer available to new customers. Existing customers of Amazon S3 Select can
continue to use the feature as usual. [Learn more](http://aws.amazon.com/blogs/storage/how-to-optimize-querying-your-data-in-amazon-s3)

The expression that is used to query the object.

Type: String

Required: Yes

**ExpressionType**

The type of the provided expression (for example, SQL).

Type: String

Valid Values: `SQL`

Required: Yes

**InputSerialization**

Describes the serialization format of the object.

Type: [InputSerialization](api-inputserialization.md) data type

Required: Yes

**OutputSerialization**

Describes how the results of the Select job are serialized.

Type: [OutputSerialization](api-outputserialization.md) data type

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/selectparameters.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/selectparameters.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/selectparameters.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SelectObjectContentEventStream

ServerSideEncryptionByDefault

All content copied from https://docs.aws.amazon.com/.
