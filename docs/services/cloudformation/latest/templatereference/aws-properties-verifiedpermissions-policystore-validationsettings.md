This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::PolicyStore ValidationSettings

A structure that contains Cedar policy validation settings for the policy store. The
validation mode determines which validation failures that Cedar considers serious enough
to block acceptance of a new or edited static policy or policy template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mode" : String
}

```

### YAML

```yaml

  Mode: String

```

## Properties

`Mode`

The validation mode currently configured for this policy store. The valid values
are:

- **OFF** – Neither Verified Permissions nor Cedar
perform any validation on policies. No validation errors are reported by either
service.

- **STRICT** – Requires a schema to be present in
the policy store. Cedar performs validation on all submitted new or updated
static policies and policy templates. Any that fail validation are rejected and
Cedar doesn't store them in the policy store.

###### Important

If `Mode=STRICT` and the policy store doesn't contain a schema, Verified Permissions rejects all static policies and policy templates because there is no
schema to validate against.

To submit a static policy or policy template without a schema, you must turn off
validation.

_Required_: Yes

_Type_: String

_Allowed values_: `OFF | STRICT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::VerifiedPermissions::PolicyTemplate
