This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::Channel

Contains information about a returned CloudTrail channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudTrail::Channel",
  "Properties" : {
      "Destinations" : [ Destination, ... ],
      "Name" : String,
      "Source" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CloudTrail::Channel
Properties:
  Destinations:
    - Destination
  Name: String
  Source: String
  Tags:
    - Tag

```

## Properties

`Destinations`

One or more event data stores to which events arriving through a channel will be logged.

_Required_: No

_Type_: Array of [Destination](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudtrail-channel-destination.html)

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the channel.

_Required_: No

_Type_: String

_Pattern_: `(^[a-zA-Z0-9._\-]+$)`

_Minimum_: `3`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The name of the partner or external event source. You cannot change this name after you create the
channel. A maximum of one channel is allowed per source.

A source can be either `Custom` for all valid non-AWS
events, or the name of a partner event source. For information about the source names for available partners, see [Additional information about integration partners](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store-integration.html#cloudtrail-lake-partner-information) in the CloudTrail User Guide.

_Required_: No

_Type_: String

_Pattern_: `(.*)`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudtrail-channel-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function,
`Ref` returns the resource name.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ChannelArn`

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the CloudTrail channel,
such as `arn:aws:cloudtrail:us-east-2:123456789012:channel/01234567890`.

## Examples

### Example

The following example creates a channel for a CloudTrail Lake integration
with an event source outside of AWS. For information about
CloudTrail Lake integrations, see [Create an integration with an event source outside of AWS](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store-integration.html)
in the _AWS CloudTrail User Guide_.

#### JSON

```json

{
    "Parameters": {
    	"Name" : String,
      	"EventDataStoreArn" : String,
      	"Source" : String
    },
    "Resources": {
        "myChannel": {
            "Type": "AWS::CloudTrail::Channel",
            "Properties": {
                "Name": {
                    "Ref": "Name"
                },
                "Source": {
                    "Ref": "Source"
                },
                "Destinations": [
                    {
                        "Type": "EVENT_DATA_STORE",
                        "Location": "{
                        	"Ref": "arn:aws:cloudtrail:us-east-1:12345678910:eventdatastore/EXAMPLE-f852-4e8f-8bd1-bcf6cEXAMPLE"
                    	}
                    }
                ],
                "Tags": [
                    {
                        "Key": "TagKeyIntTest",
                        "Value": "TagValueIntTest"
                    },
                    {
                        "Key": "TagKeyIntTest2",
                        "Value": "TagValueIntTest2"
                    }
                ]
            }
        }
    },
    "Outputs": {
        "myChannelArn": {
            "Description": "The channel ARN",
            "Value": {
                "Fn::GetAtt": [
                    "myChannel",
                    "arn:aws:cloudtrail:us-east-1:01234567890:channel/EXAMPLE8-0558-4f7e-a06a-43969EXAMPLE"
                ]
            }
        }
    }
}

```

#### YAML

```yaml

Parameters:
  Name:
    Type: String
  EventDataStoreArn:
    Type: String
  Source:
    Type: String
Resources:
  myChannel:
    Type: AWS::CloudTrail::Channel
    Properties:
      Name: !Ref Name
      Source: !Ref Source
      Destinations:
        - Type: "EVENT_DATA_STORE"
          Location: !Ref arn:aws:cloudtrail:us-east-1:12345678910:eventdatastore/EXAMPLE-f852-4e8f-8bd1-bcf6cEXAMPLE
      Tags:
        - Key: "TagKeyIntTest"
          Value: "TagValueIntTest"
        - Key: "TagKeyIntTest2"
          Value: "TagValueIntTest2"
Outputs:
  myChannelArn:
    Description: The channel ARN
    Value:
      'Fn::GetAtt':
        - myChannel
        - arn:aws:cloudtrail:us-east-1:01234567890:channel/EXAMPLE8-0558-4f7e-a06a-43969EXAMPLE

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS CloudTrail

Destination
