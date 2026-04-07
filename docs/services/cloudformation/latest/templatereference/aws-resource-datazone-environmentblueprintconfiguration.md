This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::EnvironmentBlueprintConfiguration

The configuration details of an environment blueprint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::EnvironmentBlueprintConfiguration",
  "Properties" : {
      "DomainIdentifier" : String,
      "EnabledRegions" : [ String, ... ],
      "EnvironmentBlueprintIdentifier" : String,
      "EnvironmentRolePermissionBoundary" : String,
      "GlobalParameters" : {Key: Value, ...},
      "ManageAccessRoleArn" : String,
      "ProvisioningConfigurations" : [ ProvisioningConfiguration, ... ],
      "ProvisioningRoleArn" : String,
      "RegionalParameters" : [ RegionalParameter, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::EnvironmentBlueprintConfiguration
Properties:
  DomainIdentifier: String
  EnabledRegions:
    - String
  EnvironmentBlueprintIdentifier: String
  EnvironmentRolePermissionBoundary: String
  GlobalParameters:
    Key: Value
  ManageAccessRoleArn: String
  ProvisioningConfigurations:
    - ProvisioningConfiguration
  ProvisioningRoleArn: String
  RegionalParameters:
    - RegionalParameter

```

## Properties

`DomainIdentifier`

The identifier of the Amazon DataZone domain in which an environment blueprint exists.

_Required_: Yes

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnabledRegions`

The enabled AWS Regions specified in a blueprint configuration.

_Required_: Yes

_Type_: Array of String

_Maximum_: `16`

_Minimum_: `4 | 0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentBlueprintIdentifier`

The identifier of the environment blueprint.

In the current release, only the following values are supported: `DefaultDataLake` and
`DefaultDataWarehouse`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentRolePermissionBoundary`

The environment role permission boundary.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[^:]*:iam::(aws|\d{12}):policy/[\w+=,.@-]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalParameters`

Region-agnostic environment blueprint parameters.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManageAccessRoleArn`

The ARN of the manage access role.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[^:]*:iam::\d{12}:role(/[a-zA-Z0-9+=,.@_-]+)*/[a-zA-Z0-9+=,.@_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisioningConfigurations`

The provisioning configuration of a blueprint.

_Required_: No

_Type_: Array of [ProvisioningConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-environmentblueprintconfiguration-provisioningconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisioningRoleArn`

The ARN of the provisioning role.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[^:]*:iam::\d{12}:role(/[a-zA-Z0-9+=,.@_-]+)*/[a-zA-Z0-9+=,.@_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegionalParameters`

The regional parameters of the environment blueprint.

_Required_: No

_Type_: Array of [RegionalParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-environmentblueprintconfiguration-regionalparameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainId` and the
`EnvironmentBlueprintId`, which uniquely identifies the environment blueprint configuration. For example:
`{ "Ref": "MyBlueprintConfiguration" }` for the resource with the logical ID
`MyBlueprintConfiguration`, `Ref` returns `DomainId|BlueprintConfigurationId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp of when an environment blueprint was created.

`DomainId`

The identifier of the Amazon DataZone domain in which an environment blueprint exists.

`EnvironmentBlueprintId`

The identifier of the environment blueprint. This identifier should be used when creating environment
profiles.

`UpdatedAt`

The timestamp of when the environment blueprint was updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AwsConsoleLinkParameters

LakeFormationConfiguration
