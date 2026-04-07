This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket RecordExpiration

The journal table record expiration settings for a journal table in an S3 Metadata configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Days" : Integer,
  "Expiration" : String
}

```

### YAML

```yaml

  Days: Integer
  Expiration: String

```

## Properties

`Days`

If you enable journal table record expiration, you can set the number of days to retain your
journal table records. Journal table records must be retained for a minimum of 7 days. To set
this value, specify any whole number from `7` to `2147483647`. For example,
to retain your journal table records for one year, set this value to `365`.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expiration`

Specifies whether journal table record expiration is enabled or disabled.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QueueConfiguration

RedirectAllRequestsTo
