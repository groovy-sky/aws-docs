This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CodeSecurityIntegration

Creates a code security integration with a source code repository provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::InspectorV2::CodeSecurityIntegration",
  "Properties" : {
      "CreateIntegrationDetails" : CreateDetails,
      "Name" : String,
      "Tags" : {Key: Value, ...},
      "Type" : String,
      "UpdateIntegrationDetails" : UpdateDetails
    }
}

```

### YAML

```yaml

Type: AWS::InspectorV2::CodeSecurityIntegration
Properties:
  CreateIntegrationDetails:
    CreateDetails
  Name: String
  Tags:
    Key: Value
  Type: String
  UpdateIntegrationDetails:
    UpdateDetails

```

## Properties

`CreateIntegrationDetails`

Contains details required to create a code security integration with a specific
repository provider.

_Required_: No

_Type_: [CreateDetails](aws-properties-inspectorv2-codesecurityintegration-createdetails.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the code security integration.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_$:.]*$`

_Minimum_: `1`

_Maximum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to apply to the code security integration.

_Required_: No

_Type_: Object of String

_Pattern_: `^.{2,127}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of repository provider for the integration.

_Required_: No

_Type_: String

_Allowed values_: `GITLAB_SELF_MANAGED | GITHUB`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdateIntegrationDetails`

The updated integration details specific to the repository provider type.

_Required_: No

_Type_: [UpdateDetails](aws-properties-inspectorv2-codesecurityintegration-updatedetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the code security integration.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the code security integration.

`AuthorizationUrl`

The URL used to authorize the integration with the repository provider.

`CreatedAt`

The timestamp when the code security integration was created.

`LastUpdatedAt`

The timestamp when the code security integration was last updated.

`Status`

The current status of the integration.

`StatusReason`

The reason for the current status of the code security integration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WeeklySchedule

CreateDetails

All content copied from https://docs.aws.amazon.com/.
