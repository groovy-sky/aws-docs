This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# `Fn::Length`

The intrinsic function `Fn::Length` returns the number of elements within an
array or an intrinsic function that returns an array.

###### Important

You must use the [AWS::LanguageExtensions transform](transform-aws-languageextensions.md) to use the
`Fn::Length` intrinsic function.

## Declaration

### JSON

```json

{ "Fn::Length" : IntrinsicFunction }
```

```json

{ "Fn::Length" : Array }
```

### YAML

```yaml

Fn::Length : IntrinsicFunction
```

```yaml

Fn::Length : Array
```

## Parameters

`IntrinsicFunction`

The intrinsic function that returns an array that you want to return a
number of elements from.

`Array`

The array you want to return the number of elements from.

## Return value

The number of elements in the intrinsic function that returns an array or in the array
passed to the intrinsic function.

## Examples

### Return the number of elements in an intrinsic function that returns an array

This example snippet returns the number of elements in an intrinsic function that
returns an array. The function returns 3.

#### JSON

```json

{
//...
    "Transform": "AWS::LanguageExtensions"
    //...
        "Fn::Length" : {
            "Fn::Split": ["|", "a|b|c"]
        }
//...
}
```

#### YAML

```yaml

Transform: 'AWS::LanguageExtensions'
#...
  Fn::Length:
    !Split ["|", "a|b|c"]
#...
```

### Return the number of elements in a Ref intrinsic function that refers to a list parameter type

This example snippet returns the number of elements in a `Ref`
intrinsic function that refers to a list parameter type. If the parameter with the
name `ListParameter` is a list with 3 elements, the function returns
3.

#### JSON

```json

{
//...
    "Transform": "AWS::LanguageExtensions"
    //...
        "Fn::Length": {
            "Ref": "ListParameter"
        }
//...
}
```

#### YAML

```yaml

Transform: 'AWS::LanguageExtensions'
#...
  Fn::Length:
    !Ref ListParameter
#...
```

### Return the number of elements in an array

This example snippet returns the number of elements in the array passed to the
intrinsic function. The function returns 3.

#### JSON

```json

{
//...
    "Transform": "AWS::LanguageExtensions"
    //...
        "Fn::Length": [
            1,
            {"Ref": "ParameterName"},
            3
        ]
//...
}
```

#### YAML

```yaml

Transform: 'AWS::LanguageExtensions'
#...
  Fn::Length:
    - 1
    - !Ref ParameterName
    - 3
#...
```

## Supported functions

You can use the following functions in the `Fn::Length` intrinsic function
or array:

- `Condition Functions`

- `Fn::Base64`

- `Fn::FindInMap`

- `Fn::Join`

- `Fn::Length`

- `Fn::Select`

- `Fn::Split`

- `Fn::Sub`

- `Fn::ToJsonString`

- `Ref`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Fn::Join

Fn::Select

All content copied from https://docs.aws.amazon.com/.
