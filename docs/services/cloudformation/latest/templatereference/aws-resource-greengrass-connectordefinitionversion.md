This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ConnectorDefinitionVersion

The `AWS::Greengrass::ConnectorDefinitionVersion` resource represents a
connector definition version for AWS IoT Greengrass. A connector definition version
contains a list of connectors.

###### Note

To create a connector definition version, you must specify the ID of the connector
definition that you want to associate with the version. For information about creating a
connector definition, see [`AWS::Greengrass::ConnectorDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinition.html).

After you create a connector definition version that contains the connectors you want
to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Greengrass::ConnectorDefinitionVersion",
  "Properties" : {
      "ConnectorDefinitionId" : String,
      "Connectors" : [ Connector, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Greengrass::ConnectorDefinitionVersion
Properties:
  ConnectorDefinitionId: String
  Connectors:
    - Connector

```

## Properties

`ConnectorDefinitionId`

The ID of the connector definition associated with this version. This value is a
GUID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Connectors`

The connectors in this version. Only one instance of a given connector can be added to
the connector definition version at a time.

_Required_: Yes

_Type_: Array of [Connector](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrass-connectordefinitionversion-connector.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the connector definition
version, such as `arn:aws:greengrass:us-east-1:123456789012:/greengrass/definition/connectors/1234a5b6-78cd-901e-2fgh-3i45j6k178l9/versions/9876ac30-4bdb-4f9d-95af-b5fdb66be1a2`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Connector Definition Version Snippet

The following snippet defines connector definition and connector definition
version resources. The connector definition version references the connector
definition and contains a connector.

For an example of a complete template, see the [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html) resource.

#### JSON

```json

"TestConnectorDefinition": {
    "Type": "AWS::Greengrass::ConnectorDefinition",
    "Properties": {
        "Name": "DemoTestConnectorDefinition"
    }
},
"TestConnectorDefinitionVersion": {
    "Type": "AWS::Greengrass::ConnectorDefinitionVersion",
    "Properties": {
        "ConnectorDefinitionId": {
            "Ref": "TestConnectorDefinition"
        },
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
```

#### YAML

```yaml

TestConnectorDefinition:
  Type: 'AWS::Greengrass::ConnectorDefinition'
  Properties:
    Name: DemoTestConnectorDefinition
TestConnectorDefinitionVersion:
  Type: 'AWS::Greengrass::ConnectorDefinitionVersion'
  Properties:
    ConnectorDefinitionId: !Ref TestConnectorDefinition
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

- [CreateConnectorDefinitionVersion](https://docs.aws.amazon.com/greengrass/v1/apireference/createconnectordefinitionversion-post.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConnectorDefinitionVersion

Connector
