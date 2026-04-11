This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Integration

The `AWS::Glue::Integration` resource specifies an AWS Glue zero-ETL integration from a data source to a target. For more information, see [zero-ETL integration supported by AWS Glue](../../../glue/latest/dg/zero-etl-using.md) and [integration structure](../../../glue/latest/dg/aws-glue-api-integrations.md) in the AWS Glue developer guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::Integration",
  "Properties" : {
      "AdditionalEncryptionContext" : {Key: Value, ...},
      "DataFilter" : String,
      "Description" : String,
      "IntegrationConfig" : IntegrationConfig,
      "IntegrationName" : String,
      "KmsKeyId" : String,
      "SourceArn" : String,
      "Tags" : [ Tag, ... ],
      "TargetArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Glue::Integration
Properties:
  AdditionalEncryptionContext:
    Key: Value
  DataFilter: String
  Description: String
  IntegrationConfig:
    IntegrationConfig
  IntegrationName: String
  KmsKeyId: String
  SourceArn: String
  Tags:
    - Tag
  TargetArn: String

```

## Properties

`AdditionalEncryptionContext`

An optional set of non-secret key–value pairs that contains additional contextual information for encryption. This can only be provided if `KMSKeyId` is provided.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `131072`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataFilter`

Selects source tables for the integration using Maxwell filter syntax.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the integration.

_Required_: No

_Type_: String

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegrationConfig`

The structure used to define properties associated with the zero-ETL integration. For more information, see [IntegrationConfig structure.](../../../glue/latest/dg/aws-glue-api-integrations.md#aws-glue-api-integrations-IntegrationConfig)

_Required_: No

_Type_: [IntegrationConfig](aws-properties-glue-integration-integrationconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IntegrationName`

A unique name for the integration.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The ARN of a KMS key used for encrypting the channel.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceArn`

The ARN for the source of the integration.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws:.*:.*:[0-9]+:.*`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata assigned to the resource consisting of a list of key-value pairs.

_Required_: No

_Type_: Array of [Tag](aws-properties-glue-integration-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetArn`

The ARN for the target of the integration.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws:.*:.*:[0-9]+:.*`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreateTime`

The time when the integration was created, in UTC.

`IntegrationArn`

The Amazon Resource Name (ARN) for the created integration.

`Status`

The status of the integration being created.

The possible statuses are:

- CREATING: The integration is being created.

- ACTIVE: The integration creation succeeds.

- MODIFYING: The integration is being modified.

- FAILED: The integration creation fails.

- DELETING: The integration is deleted.

- SYNCING: The integration is synchronizing.

- NEEDS\_ATTENTION: The integration needs attention, such as synchronization.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::IdentityCenterConfiguration

IntegrationConfig

All content copied from https://docs.aws.amazon.com/.
