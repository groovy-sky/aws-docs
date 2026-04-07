This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Resolver::OutpostResolver

Creates a Amazon Route 53 Resolver on an Outpost.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53Resolver::OutpostResolver",
  "Properties" : {
      "InstanceCount" : Integer,
      "Name" : String,
      "OutpostArn" : String,
      "PreferredInstanceType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53Resolver::OutpostResolver
Properties:
  InstanceCount: Integer
  Name: String
  OutpostArn: String
  PreferredInstanceType: String
  Tags:
    - Tag

```

## Properties

`InstanceCount`

Amazon EC2 instance count for the Resolver on the Outpost.

_Required_: No

_Type_: Integer

_Minimum_: `4`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Name of the Resolver.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutpostArn`

The ARN (Amazon Resource Name) for the Outpost.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreferredInstanceType`

The Amazon EC2 instance type. If you specify this, you must also specify a value for the `OutpostArn`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A key value pair that helps you identify a Route 53 Resolver.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53resolver-outpostresolver-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns returns the `Id` of the Outpost Resolver.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN (Amazon Resource Name) for the Resolver on an Outpost.

`CreationTime`

The date and time that the Outpost Resolver was created, in Unix time format and Coordinated Universal Time (UTC).

`CreatorRequestId`

A unique string that identifies the request that created the Resolver endpoint.
The `CreatorRequestId` allows failed requests to be retried without the risk of running the operation twice.

`Id`

The ID of the Resolver on Outpost.

`ModificationTime`

The date and time that the Outpost Resolver was modified, in Unix time format and Coordinated Universal Time (UTC).

`Status`

Status of the Resolver.

Valid Values: CREATING \| OPERATIONAL \| UPDATING \| DELETING \| ACTION\_NEEDED \| FAILED\_CREATION \| FAILED\_DELETION.

`StatusMessage`

A detailed description of the Resolver.

## Examples

### Create a Resolver on Outpost

The following example creates a Amazon Route 53 Resolver on an AWS Outposts.

#### JSON

```json

{
"Type": "AWS::Route53Resolver::OutpostResolver",
"Properties": {
    "Name": "SampleOutpostResolver",
    "InstanceCount": 4,
    "OutpostArn": "arn:aws:outposts:us-west-2:123456789012:outpost/op-12345678901234567",
    "PreferredInstanceType": "m5.large",
    "Tags": [
        {
            "Key": "keyname1",
            "Value": "value1"
        },
        {
            "Key": "keyname2",
            "Value": "value2"
        }
      ]
    }
}
```

#### YAML

```yaml

Type: AWS::Route53Resolver::OutpostResolver
Properties:
  Name: SampleOutpostResolver
  InstanceCount: 4
  OutpostArn: arn:aws:outposts:us-west-2:123456789012:outpost/op-12345678901234567
  PreferredInstanceType: m5.large
  Tags:
    - Key: "keyname1"
      Value: "value1"
    - Key: "keyname2"
      Value: "value2"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
