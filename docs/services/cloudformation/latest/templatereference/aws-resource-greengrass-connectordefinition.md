This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ConnectorDefinition

The `AWS::Greengrass::ConnectorDefinition` resource represents a connector
definition for AWS IoT Greengrass. Connector definitions are used to organize your
connector definition versions.

Connector definitions can reference multiple connector definition versions. All
connector definition versions must be associated with a connector definition. Each
connector definition version can contain one or more connectors.

###### Note

When you create a connector definition, you can optionally include an initial
connector definition version. To associate a connector definition version later, create
an [`AWS::Greengrass::ConnectorDefinitionVersion`](../userguide/aws-resource-greengrass-connectordefinitionversion.md) resource and
specify the ID of this connector definition.

After you create the connector definition version that contains the connectors you
want to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Greengrass::ConnectorDefinition",
  "Properties" : {
      "InitialVersion" : ConnectorDefinitionVersion,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Greengrass::ConnectorDefinition
Properties:
  InitialVersion:
    ConnectorDefinitionVersion
  Name: String
  Tags:
    - Tag

```

## Properties

`InitialVersion`

The connector definition version to include when the connector definition is created. A
connector definition version contains a list of [`connector`](../userguide/aws-properties-greengrass-connectordefinition-connector.md) property types.

###### Note

To associate a connector definition version after the connector definition is
created, create an [`AWS::Greengrass::ConnectorDefinitionVersion`](../userguide/aws-resource-greengrass-connectordefinitionversion.md) resource and
specify the ID of this connector definition.

_Required_: No

_Type_: [ConnectorDefinitionVersion](aws-properties-greengrass-connectordefinition-connectordefinitionversion.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the connector definition.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Application-specific metadata to attach to the connector definition. You can use tags in
IAM policies to control access to AWS IoT Greengrass resources. You
can also use tags to categorize your resources. For more information, see [Tagging Your\
AWS IoT Greengrass Resources](../../../greengrass/v1/developerguide/tagging.md) in the _AWS IoT Greengrass Version 1 Developer Guide_.

This `Json` property type is processed as a map of key-value pairs. It uses
the following format, which is different from most `Tags` implementations in
CloudFormation templates.

```json

"Tags": {
    "KeyName0": "value",
    "KeyName1": "value",
    "KeyName2": "value"
}
```

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the connector definition, such as
`1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the `ConnectorDefinition`, such as
`arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/connectors/1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

`Id`

The ID of the `ConnectorDefinition`, such as
`1234a5b6-78cd-901e-2fgh-3i45j6k178l9`.

`LatestVersionArn`

The ARN of the last `ConnectorDefinitionVersion` that was added to the
`ConnectorDefinition`, such as `arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/connectors/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2`.

`Name`

The name of the `ConnectorDefinition`, such as
`MyConnectorDefinition`.

## Examples

### Connector Definition Snippet

The following snippet defines a connector definition resource with an initial
version that contains a connector.

For an example of a complete template, see the [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md) resource.

#### JSON

```json

"TestConnectorDefinition": {
    "Type": "AWS::Greengrass::ConnectorDefinition",
    "Properties": {
        "Name": "DemoTestConnectorDefinition",
        "InitialVersion": {
            "Connectors": [
                {
                    "Id": "Connector1",
                    "ConnectorArn": {
                        "Fn::Join": [
                            ":",
                            [
                                "arn:aws:greengrass",
                                {
                                    "Ref": "AWS::Region"
                                },
                                ":/connectors/SNS/versions/1"
                            ]
                        ]
                    },
                    "Parameters": {
                        "DefaultSNSArn": {
                            "Fn::Join": [
                                ":",
                                [
                                    "arn:aws:sns",
                                    {
                                        "Ref": "AWS::Region"
                                    },
                                    {
                                        "Ref": "AWS::AccountId"
                                    },
                                    "defaultSns"
                                ]
                            ]
                        }
                    }
                }
            ]
        }
    }
}

```

#### YAML

```yaml

TestConnectorDefinition:
  Type: 'AWS::Greengrass::ConnectorDefinition'
  Properties:
    Name: DemoTestConnectorDefinition
    InitialVersion:
      Connectors:
        - Id: Connector1
          ConnectorArn: !Join
            - ':'
            - - 'arn:aws:greengrass'
              - !Ref 'AWS::Region'
              - ':/connectors/SNS/versions/1'
          Parameters:
            DefaultSNSArn: !Join
              - ':'
              - - 'arn:aws:sns'
                - !Ref 'AWS::Region'
                - !Ref 'AWS::AccountId'
                - defaultSns
```

## See also

- [CreateConnectorDefinition](../../../../reference/greengrass/v1/apireference/createconnectordefinition-post.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS IoT Greengrass

Connector

All content copied from https://docs.aws.amazon.com/.
