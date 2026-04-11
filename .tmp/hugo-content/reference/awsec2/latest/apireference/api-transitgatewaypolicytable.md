# TransitGatewayPolicyTable

Describes a transit gateway policy table.

## Contents

**creationTime**

The timestamp when the transit gateway policy table was created.

Type: Timestamp

Required: No

**state**

The state of the transit gateway policy table

Type: String

Valid Values: `pending | available | deleting | deleted`

Required: No

**TagSet.N**

he key-value pairs associated with the transit gateway policy table.

Type: Array of [Tag](api-tag.md) objects

Required: No

**transitGatewayId**

The ID of the transit gateway.

Type: String

Required: No

**transitGatewayPolicyTableId**

The ID of the transit gateway policy table.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewaypolicytable.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewaypolicytable.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewaypolicytable.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayPolicyRuleMetaData

TransitGatewayPolicyTableAssociation
