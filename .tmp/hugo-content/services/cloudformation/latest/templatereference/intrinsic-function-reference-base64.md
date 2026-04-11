This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# `Fn::Base64`

The intrinsic function `Fn::Base64` returns the Base64 representation of the
input string. This function is typically used to pass encoded data to Amazon EC2 instances by way
of the `UserData` property.

## Declaration

### JSON

```json

{ "Fn::Base64" : valueToEncode }
```

### YAML

Syntax for the full function name:

```yaml

Fn::Base64: valueToEncode
```

Syntax for the short form:

```yaml

!Base64 valueToEncode
```

###### Note

If you use the short
form and immediately include another function in the `valueToEncode`
parameter, use the full function name for at least one of the functions. For example,
the following syntax isn't valid:

```yaml

!Base64 !Sub string
!Base64 !Ref logical_ID
```

Instead, use the full function name for at least one of the functions, as shown in
the following examples:

```yaml

!Base64
  "Fn::Sub": string

Fn::Base64:
  !Sub string
```

## Parameters

valueToEncode

The string value you want to convert to Base64.

## Return value:

The original string, in Base64 representation.

## Examples

### JSON

```json

{ "Fn::Base64" : "AWS CloudFormation" }
```

### YAML

```yaml

Fn::Base64: AWS CloudFormation
```

## Supported functions

You can use any function that returns a string inside the `Fn::Base64`
function.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Intrinsic functions

Fn::Cidr

All content copied from https://docs.aws.amazon.com/.
