This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventSourceMapping DocumentDBEventSourceConfig

Specific configuration settings for a DocumentDB event source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CollectionName" : String,
  "DatabaseName" : String,
  "FullDocument" : String
}

```

### YAML

```yaml

  CollectionName: String
  DatabaseName: String
  FullDocument: String

```

## Properties

`CollectionName`

The name of the collection to consume within the database. If you do not specify a collection, Lambda consumes all collections.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `57`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

The name of the database to consume within the DocumentDB cluster.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FullDocument`

Determines what DocumentDB sends to your event stream during document update operations. If set to UpdateLookup, DocumentDB sends a delta describing the changes, along with a copy of the entire document. Otherwise, DocumentDB sends only a partial document that contains the changes.

_Required_: No

_Type_: String

_Allowed values_: `UpdateLookup | Default`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DestinationConfig

Endpoints
