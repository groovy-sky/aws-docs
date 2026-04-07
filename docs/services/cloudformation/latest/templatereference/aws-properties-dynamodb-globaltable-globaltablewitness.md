This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable GlobalTableWitness

The witness Region for the MRSC global table. A MRSC global table can be configured with
either three replicas, or with two replicas and one witness.

The witness must be in a different Region than the replicas and within the same Region
set:

- US Region set: US East (N. Virginia), US East (Ohio), US West (Oregon)

- EU Region set: Europe (Ireland), Europe (London), Europe (Paris), Europe
(Frankfurt)

- AP Region set: Asia Pacific (Tokyo), Asia Pacific (Seoul), Asia Pacific
(Osaka)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Region" : String
}

```

### YAML

```yaml

  Region: String

```

## Properties

`Region`

The name of the AWSRegion that serves as a witness for the MRSC global
table.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GlobalSecondaryIndex

KeySchema
