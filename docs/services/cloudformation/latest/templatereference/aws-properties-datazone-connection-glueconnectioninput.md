This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection GlueConnectionInput

The AWS Glue connecton input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AthenaProperties" : {Key: Value, ...},
  "AuthenticationConfiguration" : AuthenticationConfigurationInput,
  "ConnectionProperties" : {Key: Value, ...},
  "ConnectionType" : String,
  "Description" : String,
  "MatchCriteria" : String,
  "Name" : String,
  "PhysicalConnectionRequirements" : PhysicalConnectionRequirements,
  "PythonProperties" : {Key: Value, ...},
  "SparkProperties" : {Key: Value, ...},
  "ValidateCredentials" : Boolean,
  "ValidateForComputeEnvironments" : [ String, ... ]
}

```

### YAML

```yaml

  AthenaProperties:
    Key: Value
  AuthenticationConfiguration:
    AuthenticationConfigurationInput
  ConnectionProperties:
    Key: Value
  ConnectionType: String
  Description: String
  MatchCriteria: String
  Name: String
  PhysicalConnectionRequirements:
    PhysicalConnectionRequirements
  PythonProperties:
    Key: Value
  SparkProperties:
    Key: Value
  ValidateCredentials: Boolean
  ValidateForComputeEnvironments:
    - String

```

## Properties

`AthenaProperties`

The Amazon Athena properties of the AWS Glue connection.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\u0020-\uD7FF\uE000-\uFFFF\t]*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticationConfiguration`

The authentication configuration of the AWS Glue connection.

_Required_: No

_Type_: [AuthenticationConfigurationInput](aws-properties-datazone-connection-authenticationconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionProperties`

The connection properties of the AWS Glue connection.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionType`

The connection type of the AWS Glue connection.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the AWS Glue connection.

_Required_: No

_Type_: String

_Pattern_: `^[\u0020-\uD7FF\uE000-\uFFFF\r\n\t]*$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchCriteria`

The match criteria of the AWS Glue connection.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the AWS Glue connection.

_Required_: No

_Type_: String

_Pattern_: `^[\u0020-\uD7FF\uE000-\uFFFF\t]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhysicalConnectionRequirements`

The physical connection requirements for the AWS Glue connection.

_Required_: No

_Type_: [PhysicalConnectionRequirements](aws-properties-datazone-connection-physicalconnectionrequirements.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PythonProperties`

The Python properties of the AWS Glue connection.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\u0020-\uD7FF\uE000-\uFFFF\t]*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SparkProperties`

The Spark properties of the AWS Glue connection.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\u0020-\uD7FF\uE000-\uFFFF\t]*$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidateCredentials`

Speciefies whether to validate credentials of the AWS Glue
connection.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidateForComputeEnvironments`

Speciefies whether to validate for compute environments of the AWS Glue
connection.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectionPropertiesInput

GlueOAuth2Credentials

All content copied from https://docs.aws.amazon.com/.
