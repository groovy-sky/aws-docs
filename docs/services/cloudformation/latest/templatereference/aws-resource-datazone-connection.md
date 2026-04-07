This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection

In Amazon DataZone, a connection enables you to connect your resources (domains, projects, and environments) to
external resources and services.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataZone::Connection",
  "Properties" : {
      "AwsLocation" : AwsLocation,
      "Description" : String,
      "DomainIdentifier" : String,
      "EnableTrustedIdentityPropagation" : Boolean,
      "EnvironmentIdentifier" : String,
      "Name" : String,
      "ProjectIdentifier" : String,
      "Props" : ConnectionPropertiesInput,
      "Scope" : String
    }
}

```

### YAML

```yaml

Type: AWS::DataZone::Connection
Properties:
  AwsLocation:
    AwsLocation
  Description: String
  DomainIdentifier: String
  EnableTrustedIdentityPropagation: Boolean
  EnvironmentIdentifier: String
  Name: String
  ProjectIdentifier: String
  Props:
    ConnectionPropertiesInput
  Scope: String

```

## Properties

`AwsLocation`

The location where the connection is created.

_Required_: No

_Type_: [AwsLocation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-connection-awslocation.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

Connection description.

_Required_: No

_Type_: String

_Pattern_: `^[\S\s]*$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainIdentifier`

The ID of the domain where the connection is created.

_Required_: Yes

_Type_: String

_Pattern_: `^dzd[_-][a-zA-Z0-9_-]{1,36}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnableTrustedIdentityPropagation`

Specifies whether the trusted identity propagation is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnvironmentIdentifier`

The ID of the environment where the connection is created.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the connection.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w][\w\.\-\_]*$`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProjectIdentifier`

The ID of the project where you want to list connections.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Props`

Connection props.

_Required_: No

_Type_: [ConnectionPropertiesInput](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-datazone-connection-connectionpropertiesinput.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

The scope of the connection.

_Required_: No

_Type_: String

_Allowed values_: `DOMAIN | PROJECT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string containing pipe-separated `DomainId` and
`ConnectionId`, which uniquely identifies a connection. For example: `{ "Ref": "MyConnection"
    }` for the resource with the logical ID `MyConnection`, `Ref` returns
`DomainId|ConnectionId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConnectionId`

The ID of the connection.

`DomainId`

The domain ID of the connection.

`DomainUnitId`

The domain unit ID of the connection.

`EnvironmentId`

The ID of the environment.

`EnvironmentUserRole`

The environment user role.

`ProjectId`

The ID of the project.

`Type`

The type of the connection.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon DataZone

AmazonQPropertiesInput
