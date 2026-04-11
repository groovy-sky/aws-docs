This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution OriginGroupFailoverCriteria

A complex data type that includes information about the failover criteria for an
origin group, including the status codes for which CloudFront will failover from the
primary origin to the second origin.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StatusCodes" : StatusCodes
}

```

### YAML

```yaml

  StatusCodes:
    StatusCodes

```

## Properties

`StatusCodes`

The status codes that, when returned from the primary origin, will trigger CloudFront to failover to the second origin.

_Required_: Yes

_Type_: [StatusCodes](aws-properties-cloudfront-distribution-statuscodes.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OriginGroup

OriginGroupMember

All content copied from https://docs.aws.amazon.com/.
