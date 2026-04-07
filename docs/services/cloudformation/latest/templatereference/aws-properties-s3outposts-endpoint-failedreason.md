This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Outposts::Endpoint FailedReason

The failure reason, if any, for a create or delete endpoint operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ErrorCode" : String,
  "Message" : String
}

```

### YAML

```yaml

  ErrorCode: String
  Message: String

```

## Properties

`ErrorCode`

The failure code, if any, for a create or delete endpoint operation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Message`

Additional error details describing the endpoint failure and recommended action.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::S3Outposts::Endpoint

NetworkInterface
