This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EntityResolution::PolicyStatement

Adds a policy statement object. To retrieve a list of existing policy statements, use
the `GetPolicy` API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EntityResolution::PolicyStatement",
  "Properties" : {
      "Action" : [ String, ... ],
      "Arn" : String,
      "Condition" : String,
      "Effect" : String,
      "Principal" : [ String, ... ],
      "StatementId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EntityResolution::PolicyStatement
Properties:
  Action:
    - String
  Arn: String
  Condition: String
  Effect: String
  Principal:
    - String
  StatementId: String

```

## Properties

`Action`

The action that the principal can use on the resource.

For example, `entityresolution:GetIdMappingJob`,
`entityresolution:GetMatchingJob`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Arn`

The Amazon Resource Name (ARN) of the resource that will be accessed by the
principal.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):entityresolution:[a-z]{2}-[a-z]{1,10}-[0-9]:[0-9]{12}:((schemamapping|matchingworkflow|idmappingworkflow|idnamespace)/[a-zA-Z_0-9-]{1,255})$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Condition`

A set of condition keys that you can use in key policies.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `40960`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Effect`

Determines whether the permissions specified in the policy are to be allowed
( `Allow`) or denied ( `Deny`).

###### Important

If you set the value of the `effect` parameter to `Deny` for
the `AddPolicyStatement` operation, you must also set the value of the
`effect` parameter in the `policy` to `Deny` for the
`PutPolicy` operation.

_Required_: No

_Type_: String

_Allowed values_: `Allow | Deny`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Principal`

The AWS service or AWS account that can access the
resource defined as ARN.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatementId`

A statement identifier that differentiates the statement from others in the same
policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::EntityResolution::SchemaMapping
