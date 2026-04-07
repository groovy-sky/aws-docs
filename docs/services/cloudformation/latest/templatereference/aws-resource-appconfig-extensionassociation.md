This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppConfig::ExtensionAssociation

When you create an extension or configure an AWS authored extension, you
associate the extension with an AWS AppConfig application, environment, or
configuration profile. For example, you can choose to run the `AWS AppConfig
            deployment events to Amazon SNS` AWS authored extension and receive notifications on an Amazon SNS
topic anytime a configuration deployment is started for a specific application. Defining
which extension to associate with an AWS AppConfig resource is called an
_extension association_. An extension association is a specified
relationship between an extension and an AWS AppConfig resource, such as an
application or a configuration profile. For more information about extensions and
associations, see [Extending\
workflows](../../../appconfig/latest/userguide/working-with-appconfig-extensions.md) in the _AWS AppConfig User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppConfig::ExtensionAssociation",
  "Properties" : {
      "ExtensionIdentifier" : String,
      "ExtensionVersionNumber" : Integer,
      "Parameters" : {Key: Value, ...},
      "ResourceIdentifier" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppConfig::ExtensionAssociation
Properties:
  ExtensionIdentifier: String
  ExtensionVersionNumber: Integer
  Parameters:
    Key: Value
  ResourceIdentifier: String
  Tags:
    - Tag

```

## Properties

`ExtensionIdentifier`

The name, the ID, or the Amazon Resource Name (ARN) of the extension.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExtensionVersionNumber`

The version number of the extension. If not specified, AWS AppConfig uses the
maximum version of the extension.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parameters`

The parameter names and values defined in the extensions. Extension parameters marked
`Required` must be entered for this field.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceIdentifier`

The ARN of an application, configuration profile, or environment.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Adds one or more tags for the specified extension association. Tags are metadata that
help you categorize resources in different ways, for example, by purpose, owner, or
environment. Each tag consists of a key and an optional value, both of which you define.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appconfig-extensionassociation-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns information about the extension association.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the extension defined in the association.

`ExtensionArn`

The ARN of the extension defined in the association.

`Id`

The system-generated ID for the association.

`ResourceArn`

The ARNs of applications, configuration profiles, or environments defined in the
association.

## Examples

### AWS AppConfig create extension association example

An extension association defines which extension to associate with an AWS AppConfig resource. An extension association is a specified relationship between an
extension and an AWS AppConfig resource, such as an application or a configuration
profile. The following example creates an extension association between an application
called MyDependentApplication and an extension called MyAmazingExtension.

#### JSON

```json

{
  "Resources": {
    "BasicExtensionAssociation": {
      "Type": "AWS::AppConfig::ExtensionAssociation",
      "Properties": {
        "ExtensionIdentifier": "MyAmazingExtension",
        "ExtensionVersionNumber": 0,
        "ResourceIdentifier": [
          "/",
          [
            "arn:${AWS::Partition}:appconfig:${AWS::Region}:${AWS::AccountId}:application",
            "MyApplication"
          ]
         ],
        "Parameters": {
          "MyTestParam": "My test parameter value"
        },
        "Tags": [
          {
            "Key": "ExtAssociation",
            "Value": "Test"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  BasicExtensionAssociation:
    Type: AWS::AppConfig::ExtensionAssociation
    Properties:
      ExtensionIdentifier: !ImportValue MyAmazingExtension
      ExtensionVersionNumber: 0
      ResourceIdentifier: !Join
        - '/'
        - - !Sub 'arn:${AWS::Partition}:appconfig:${AWS::Region}:${AWS::AccountId}:application'
          - !ImportValue MyApplication
      Parameters:
        MyTestParam: "My test parameter value"
      Tags:
        - Key: ExtAssociation
          Value: Test
```

## See also

- [AWS AppConfig](../../../appconfig/latest/userguide/what-is-appconfig.md)

- [Working with\
AWS AppConfig extensions](../../../appconfig/latest/userguide/working-with-appconfig-extensions.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
