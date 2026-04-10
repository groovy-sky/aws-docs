This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3ObjectLambda::AccessPoint ObjectLambdaConfiguration

A configuration used when creating an Object Lambda Access Point.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedFeatures" : [ String, ... ],
  "CloudWatchMetricsEnabled" : Boolean,
  "SupportingAccessPoint" : String,
  "TransformationConfigurations" : [ TransformationConfiguration, ... ]
}

```

### YAML

```yaml

  AllowedFeatures:
    - String
  CloudWatchMetricsEnabled: Boolean
  SupportingAccessPoint: String
  TransformationConfigurations:
    - TransformationConfiguration

```

## Properties

`AllowedFeatures`

A container for allowed features. Valid inputs are `GetObject-Range`,
`GetObject-PartNumber`, `HeadObject-Range`, and
`HeadObject-PartNumber`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchMetricsEnabled`

A container for whether the CloudWatch metrics configuration is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportingAccessPoint`

Standard access point associated with the Object Lambda Access Point.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransformationConfigurations`

A container for transformation configurations for an Object Lambda Access Point.

_Required_: Yes

_Type_: Array of [TransformationConfiguration](aws-properties-s3objectlambda-accesspoint-transformationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContentTransformation

PublicAccessBlockConfiguration

All content copied from https://docs.aws.amazon.com/.
