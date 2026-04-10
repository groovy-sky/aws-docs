This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket RedirectAllRequestsTo

Specifies the redirect behavior of all requests to a website endpoint of an Amazon S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HostName" : String,
  "Protocol" : String
}

```

### YAML

```yaml

  HostName: String
  Protocol: String

```

## Properties

`HostName`

Name of the host where requests are redirected.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

Protocol to use when redirecting requests. The default is the protocol that is used in the original
request.

_Required_: No

_Type_: String

_Allowed values_: `http | https`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecordExpiration

RedirectRule

All content copied from https://docs.aws.amazon.com/.
