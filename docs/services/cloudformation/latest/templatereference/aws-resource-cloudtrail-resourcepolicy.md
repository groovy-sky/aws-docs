This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::ResourcePolicy

Attaches a resource-based permission policy to a CloudTrail event data store, dashboard, or channel. For more information about resource-based policies, see
[CloudTrail resource-based policy examples](../../../awscloudtrail/latest/userguide/security-iam-resource-based-policy-examples.md)
in the _CloudTrail User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudTrail::ResourcePolicy",
  "Properties" : {
      "ResourceArn" : String,
      "ResourcePolicy" : Json
    }
}

```

### YAML

```yaml

Type: AWS::CloudTrail::ResourcePolicy
Properties:
  ResourceArn: String
  ResourcePolicy: Json

```

## Properties

`ResourceArn`

The Amazon Resource Name (ARN) of the CloudTrail event data store, dashboard, or channel attached to the resource-based policy.

Example event data store ARN format:
`arn:aws:cloudtrail:us-east-2:123456789012:eventdatastore/EXAMPLE-f852-4e8f-8bd1-bcf6cEXAMPLE`

Example dashboard ARN format: `arn:aws:cloudtrail:us-east-1:123456789012:dashboard/exampleDash`

Example channel ARN format:
`arn:aws:cloudtrail:us-east-2:123456789012:channel/01234567890`

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9._/\-:]+$`

_Minimum_: `3`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourcePolicy`

A JSON-formatted string for an AWS resource-based policy.

For example resource-based policies, see
[CloudTrail resource-based policy examples](../../../awscloudtrail/latest/userguide/security-iam-resource-based-policy-examples.md)
in the _CloudTrail User Guide_.

_Required_: Yes

_Type_: Json

_Minimum_: `1`

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function,
`Ref` returns the resource. The resource is a combination of the resource-based
policy document and the channel ARN.

## Examples

### Example

The following example creates a resource policy that allows AWS
account ID `111122223333` to call `PutAuditEvents` on the
channel defined as the resource ARN in the policy. For information about creating a
resource policy, see [AWS CloudTrail resource-based policy examples](../../../awscloudtrail/latest/userguide/security-iam-resource-based-policy-examples.md) in the
_AWS CloudTrail User Guide_.

#### JSON

```json

{
  "Type": "AWS:CloudTrail:ResourcePolicy",
  "Properties": {
    "ResourceArn": "arn:aws:cloudtrail:us-east-1:01234567890:channel/EXAMPLE8-0558-4f7e-a06a-43969EXAMPLE",
    "ResourcePolicy": "{ \"Version\": \"2012-10-17\", \"Statement\": [ { \"Sid\": \"DeliverEventsThroughChannel\", \"Effect\": \"Allow\", \"Principal\": { \"AWS\": [ \"arn:aws:iam::111122223333:root\" ] }, \"Action\":\"cloudtrail-data:PutAuditEvents\", \"Resource\": \"arn:aws:cloudtrail:us-east-1:01234567890:channel/EXAMPLE8-0558-4f7e-a06a-43969EXAMPLE\" } ] }"
  }
}

```

#### YAML

```yaml

  Type: AWS:CloudTrail:ResourcePolicy
  Properties:
    ResourceArn: "arn:aws:cloudtrail:us-east-1:01234567890:channel/EXAMPLE8-0558-4f7e-a06a-43969EXAMPLE"
    ResourcePolicy: "{ \"Version\": \"2012-10-17\", \"Statement\": [ { \"Sid\": \"DeliverEventsThroughChannel\", \"Effect\": \"Allow\", \"Principal\": { \"AWS\": [ \"arn:aws:iam::111122223333:root\" ] }, \"Action\":\"cloudtrail-data:PutAuditEvents\", \"Resource\": \"arn:aws:cloudtrail:us-east-1:01234567890:channel/EXAMPLE8-0558-4f7e-a06a-43969EXAMPLE\" } ] }"

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::CloudTrail::Trail

All content copied from https://docs.aws.amazon.com/.
