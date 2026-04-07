This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::EnvironmentActions

The details about the specified action configured for an environment. For example, the
details of the specified console links for an analytics tool that is available in this
environment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::EnvironmentActions",
  "Properties" : {
      "Description" : String,
      "DomainIdentifier" : String,
      "EnvironmentIdentifier" : String,
      "Identifier" : String,
      "Name" : String,
      "Parameters" : AwsConsoleLinkParameters
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::EnvironmentActions
Properties:
  Description: String
  DomainIdentifier: String
  EnvironmentIdentifier: String
  Identifier: String
  Name: String
  Parameters:
    AwsConsoleLinkParameters

```

## Properties

`Description`

The environment action description.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainIdentifier`

The Amazon DataZone domain ID of the environment action.

_Required_: No

_Type_: String

_Pattern_: `^dzd[-_][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentIdentifier`

The environment ID of the environment action.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_-]{1,36}$`

_Minimum_: `1`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Identifier`

The ID of the environment action.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,36}$`

_Minimum_: `1`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the environment action.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w -]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

The parameters of the environment action.

_Required_: No

_Type_: [AwsConsoleLinkParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-environmentactions-awsconsolelinkparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainId` and
`EnvironmentActionId`, which uniquely identifies the environment action. For example: `{ "Ref":
    "MyEnvironmentAction" }` for the resource with the logical ID `MyEnvironmentAction`,
`Ref` returns `DomainId|EnvironmentActionId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DomainId`

The Amazon DataZone domain ID of the environment action.

`EnvironmentId`

The environment ID of the environment action.

`Id`

The ID of the environment action.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EnvironmentParameter

AwsConsoleLinkParameters
