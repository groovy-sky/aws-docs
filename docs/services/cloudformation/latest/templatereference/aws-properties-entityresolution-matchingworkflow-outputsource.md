This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::MatchingWorkflow OutputSource

A list of `OutputAttribute` objects, each of which have the fields
`Name` and `Hashed`. Each of these objects selects a column to be
included in the output table, and whether the values of the column should be hashed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplyNormalization" : Boolean,
  "CustomerProfilesIntegrationConfig" : CustomerProfilesIntegrationConfig,
  "KMSArn" : String,
  "Output" : [ OutputAttribute, ... ],
  "OutputS3Path" : String
}

```

### YAML

```yaml

  ApplyNormalization: Boolean
  CustomerProfilesIntegrationConfig:
    CustomerProfilesIntegrationConfig
  KMSArn: String
  Output:
    - OutputAttribute
  OutputS3Path: String

```

## Properties

`ApplyNormalization`

Normalizes the attributes defined in the schema in the input data. For example, if an
attribute has an `AttributeType` of `PHONE_NUMBER`, and the data in
the input table is in a format of 1234567890, AWS Entity Resolution will normalize this field
in the output to (123)-456-7890.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomerProfilesIntegrationConfig`

Property description not available.

_Required_: No

_Type_: [CustomerProfilesIntegrationConfig](aws-properties-entityresolution-matchingworkflow-customerprofilesintegrationconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KMSArn`

Customer KMS ARN for encryption at rest. If not provided, system will use an AWS Entity Resolution managed KMS key.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):kms:.*:[0-9]+:.*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Output`

A list of `OutputAttribute` objects, each of which have the fields
`Name` and `Hashed`. Each of these objects selects a column to be
included in the output table, and whether the values of the column should be hashed.

_Required_: Yes

_Type_: Array of [OutputAttribute](aws-properties-entityresolution-matchingworkflow-outputattribute.md)

_Minimum_: `0`

_Maximum_: `750`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputS3Path`

The S3 path to which AWS Entity Resolution will write the output table.

_Required_: No

_Type_: String

_Pattern_: `^s3://([^/]+)/?(.*?([^/]+)/?)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutputAttribute

ProviderProperties

All content copied from https://docs.aws.amazon.com/.
