This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::TrafficMirrorFilter

Specifies a Traffic Mirror filter.

A Traffic Mirror filter is a set of rules that defines the traffic to mirror.

By default, no traffic is mirrored. To mirror traffic, use [AWS::EC2::TrafficMirrorFilterRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrorfilterrule.html) to add Traffic Mirror rules to the filter.
The rules you add define what traffic gets mirrored.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::TrafficMirrorFilter",
  "Properties" : {
      "Description" : String,
      "NetworkServices" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::TrafficMirrorFilter
Properties:
  Description: String
  NetworkServices:
    - String
  Tags:
    - Tag

```

## Properties

`Description`

The description of the Traffic Mirror filter.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkServices`

The network service traffic that is associated with the Traffic Mirror filter.

Valid values are `amazon-dns`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to assign to a Traffic Mirror filter.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-trafficmirrorfilter-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Traffic Mirror filter.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Create a traffic mirror filter

This is a filter that you can use when you create a traffic mirror session. This
filter also configures mirroring of Amazon DNS network services.

#### JSON

```json

{
  "SampleTrafficMirrorFilter": {
    "Type": "AWS::EC2::TrafficMirrorFilter",
    "Properties": {
      "Description": "Example traffic mirror filter",
      "NetworkServices": [
        "amazon-dns"
      ],
      "Tags": [
        {
          "Key": "Name",
          "Value": "SampleFilter"
        }
      ]
    }
  }
}

```

#### YAML

```yaml

SampleTrafficMirrorFilter:
  Type: "AWS::EC2::TrafficMirrorFilter"
  Properties:
    Description: "Example traffic mirror filter"
    NetworkServices:
      - "amazon-dns"
    Tags:
    - Key: "Name"
      Value: "SampleFilter"

```

## See also

- [Traffic mirror\
filters and filter rules](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-filters.html) in _Traffic_
_Mirroring_

- [CreateTrafficMirrorFilter](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateTrafficMirrorFilter.html) in the _Amazon EC2 API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::SubnetRouteTableAssociation

Tag
