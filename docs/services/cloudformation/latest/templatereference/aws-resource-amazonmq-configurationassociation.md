This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmazonMQ::ConfigurationAssociation

The `AWS::AmazonMQ::ConfigurationAssociation` resource Property description not available. for AmazonMQ.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AmazonMQ::ConfigurationAssociation",
  "Properties" : {
      "Broker" : String,
      "Configuration" : ConfigurationId
    }
}

```

### YAML

```yaml

Type: AWS::AmazonMQ::ConfigurationAssociation
Properties:
  Broker: String
  Configuration:
    ConfigurationId

```

## Properties

`Broker`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Configuration`

Returns information about all configurations.

_Required_: Yes

_Type_: [ConfigurationId](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amazonmq-configurationassociation-configurationid.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon MQ configurationassociation ID. For example:

`c-1234a5b6-78cd-901e-2fgh-3i45j6k178l9`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### ConfigurationAssociation

The following example creates an Amazon MQ for ActiveMQ ConfigurationAssociation.

#### JSON

```json

"ConfigurationAssociation1": {
		"Type": "AWS::AmazonMQ::ConfigurationAssociation",
		"Properties": {
			"Broker": {
				"Ref": "Broker1"
			},
			"Configuration": {
				"Id": {
					"Ref": "Configuration1"
				},
				"Revision": {
					"Fn::GetAtt": [
						"Configuration1",
						"Revision"
					]
				}
			}
		}
	}
```

#### YAML

```yaml

 ConfigurationAssociation1:
    Type: AWS::AmazonMQ::ConfigurationAssociation
    Properties:
      Broker: {Ref: Broker1}
      Configuration:
        Id: {Ref: Configuration1}
        Revision: {'Fn::GetAtt': [Configuration1, Revision]}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TagsEntry

ConfigurationId
