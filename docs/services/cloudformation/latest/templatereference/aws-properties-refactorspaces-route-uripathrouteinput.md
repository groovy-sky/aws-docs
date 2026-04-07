This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RefactorSpaces::Route UriPathRouteInput

The configuration for the URI path route type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActivationState" : String,
  "AppendSourcePath" : Boolean,
  "IncludeChildPaths" : Boolean,
  "Methods" : [ String, ... ],
  "SourcePath" : String
}

```

### YAML

```yaml

  ActivationState: String
  AppendSourcePath: Boolean
  IncludeChildPaths: Boolean
  Methods:
    - String
  SourcePath: String

```

## Properties

`ActivationState`

If set to `ACTIVE`, traffic is forwarded to this route’s service after the
route is created.

_Required_: Yes

_Type_: String

_Allowed values_: `INACTIVE | ACTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AppendSourcePath`

If set to `true`, this option appends the source path to the service URL
endpoint.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IncludeChildPaths`

Indicates whether to match all subpaths of the given source path. If this value is
`false`, requests must match the source path exactly before they are forwarded to
this route's service.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Methods`

A list of HTTP methods to match. An empty list matches all values. If a method is present,
only HTTP requests using that method are forwarded to this route’s service.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourcePath`

This is the path that Refactor Spaces uses to match traffic. Paths must start with
`/` and are relative to the base of the application. To use path parameters in
the source path, add a variable in curly braces. For example, the resource path {user}
represents a path parameter called 'user'.

_Required_: No

_Type_: String

_Pattern_: `^(/([a-zA-Z0-9._:-]+|\{[a-zA-Z0-9._:-]+\}))+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::RefactorSpaces::Service
