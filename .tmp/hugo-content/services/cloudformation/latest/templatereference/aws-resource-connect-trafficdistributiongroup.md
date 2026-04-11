This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::TrafficDistributionGroup

Information about a traffic distribution group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::TrafficDistributionGroup",
  "Properties" : {
      "Description" : String,
      "InstanceArn" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Connect::TrafficDistributionGroup
Properties:
  Description: String
  InstanceArn: String
  Name: String
  Tags:
    - Tag

```

## Properties

`Description`

The description of the traffic distribution group.

_Required_: No

_Type_: String

_Pattern_: `(^[\S].*[\S]$)|(^[\S]$)`

_Minimum_: `1`

_Maximum_: `250`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceArn`

The Amazon Resource Name (ARN).

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov):connect:[a-z]{2}-[a-z]+-[0-9]{1}:[0-9]{1,20}:instance/[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

_Minimum_: `1`

_Maximum_: `250`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the traffic distribution group.

_Required_: Yes

_Type_: String

_Pattern_: `(^[\S].*[\S]$)|(^[\S]$)`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags used to organize, track, or control access for this resource. For example,
{"tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](aws-properties-connect-trafficdistributiongroup-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the traffic distribution group name. For example:

`{ "Ref": "myTrafficDistributionGroupName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IsDefault`

Describes whether this is the default traffic distribution group.

`Status`

The status of the traffic distribution group.

`TrafficDistributionGroupArn`

The Amazon Resource Name (ARN) of the traffic distribution group.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
