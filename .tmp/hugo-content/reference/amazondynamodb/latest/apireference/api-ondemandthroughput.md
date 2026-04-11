# OnDemandThroughput

Sets the maximum number of read and write units for the specified on-demand table. If
you use this parameter, you must specify `MaxReadRequestUnits`,
`MaxWriteRequestUnits`, or both.

## Contents

###### Note

In the following list, the required parameters are described first.

**MaxReadRequestUnits**

Maximum number of read request units for the specified table.

To specify a maximum `OnDemandThroughput` on your table, set the value of
`MaxReadRequestUnits` as greater than or equal to 1. To remove the
maximum `OnDemandThroughput` that is currently set on your table, set the
value of `MaxReadRequestUnits` to -1.

Type: Long

Required: No

**MaxWriteRequestUnits**

Maximum number of write request units for the specified table.

To specify a maximum `OnDemandThroughput` on your table, set the value of
`MaxWriteRequestUnits` as greater than or equal to 1. To remove the
maximum `OnDemandThroughput` that is currently set on your table, set the
value of `MaxWriteRequestUnits` to -1.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/ondemandthroughput.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/ondemandthroughput.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/ondemandthroughput.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LocalSecondaryIndexInfo

OnDemandThroughputOverride

All content copied from https://docs.aws.amazon.com/.
