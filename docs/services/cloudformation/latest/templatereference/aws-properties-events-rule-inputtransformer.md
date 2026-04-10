This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule InputTransformer

Contains the parameters needed for you to provide custom input to a target based on one or
more pieces of data extracted from the event.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputPathsMap" : {Key: Value, ...},
  "InputTemplate" : String
}

```

### YAML

```yaml

  InputPathsMap:
    Key: Value
  InputTemplate: String

```

## Properties

`InputPathsMap`

Map of JSON paths to be extracted from the event. You can then insert these in the
template in `InputTemplate` to produce the output you want to be sent to the
target.

`InputPathsMap` is an array key-value pairs, where each value is a valid JSON
path. You can have as many as 100 key-value pairs. You must use JSON dot notation, not bracket
notation.

The keys cannot start with "AWS."

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputTemplate`

Input template where you specify placeholders that will be filled with the values of the
keys from `InputPathsMap` to customize the data sent to the target. Enclose each
`InputPathsMaps` value in brackets: < _value_ \>

If `InputTemplate` is a JSON object (surrounded by curly braces), the following
restrictions apply:

- The placeholder cannot be used as an object key.

The following example shows the syntax for using `InputPathsMap` and
`InputTemplate`.

` "InputTransformer":`

`{`

`"InputPathsMap": {"instance": "$.detail.instance","status":
        "$.detail.status"},`

`"InputTemplate": "<instance> is in state <status>"`

`}`

To have the `InputTemplate` include quote marks within a JSON string, escape
each quote marks with a slash, as in the following example:

` "InputTransformer":`

`{`

`"InputPathsMap": {"instance": "$.detail.instance","status":
        "$.detail.status"},`

`"InputTemplate": "<instance> is in state \"<status>\""`

`}`

The `InputTemplate` can also be valid JSON with varibles in quotes or out, as
in the following example:

` "InputTransformer":`

`{`

`"InputPathsMap": {"instance": "$.detail.instance","status":
        "$.detail.status"},`

`"InputTemplate": '{"myInstance": <instance>,"myStatus": "<instance> is
        in state \"<status>\""}'`

`}`

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Transform input into a string](#aws-properties-events-rule-inputtransformer--examples--Transform_input_into_a_string)

- [Transform input into JSON](#aws-properties-events-rule-inputtransformer--examples--Transform_input_into_JSON)

- [Transform input and substitute a variable](#aws-properties-events-rule-inputtransformer--examples--Transform_input_and_substitute_a_variable)

### Transform input into a string

The following example takes `instance` and `state` and outputs a string.

#### JSON

```json

"InputTransformer": {
    "InputPathsMap": {
       "instance": "$.detail.instance-id",
       "state": "$.detail.state"
    },
    "InputTemplate": "\"instance <instance> is in <state>\"\n  }\n"
}
```

#### YAML

```yaml

InputTransformer:
  InputPathsMap:
    "instance" : "$.detail.instance-id"
    "state" : "$.detail.state"
  InputTemplate: |
    "instance <instance> is in <state>"
```

### Transform input into JSON

The following example takes `instance` and `state` and outputs JSON that includes strings and variables.

#### JSON

```json

"InputTransformer": {
    "InputPathsMap": {
       "instance": "$.detail.instance-id",
       "state": "$.detail.state"
    },
    "InputTemplate": "{\n   \"instance\" : <instance>,\n   \"state\" : <state>,\n   \"instanceStatus\": \"instance \\\"<instance>\\\" is in <state>\"\n}\n"
}
```

#### YAML

```yaml

InputTransformer:
  InputPathsMap:
    "instance" : "$.detail.instance-id"
    "state" : "$.detail.state"
  InputTemplate: |
    {
       "instance" : <instance>,
       "state" : <state>,
       "instanceStatus": "instance \"<instance>\" is in <state>"
    }
```

### Transform input and substitute a variable

The following example defines a variable `RootDomainName`. It then takes `instance` and
`state`, substitutes `RootDomainName` for `domain`, and outputs JSON.

#### JSON

```json

"Parameters": {
    "RootDomainName": {
        "Description": "Domain name to use",
        "Type": "String",
        "Default": "testdomain.com"
    }
},
```

#### JSON

```json

"InputTransformer": {
    "InputPathsMap": {
       "instance": "$.detail.instance-id",
       "state": "$.detail.state"
    },
    "InputTemplate": {
        "Fn::Sub": [
            "{\n   \"domain\": \"${Domain}\",\n   \"instance\" : <instance>,\n   \"state\" : <state>,\n   \"instanceStatus\": \"instance \\\"<instance>\\\" is in <state>\"\n} - ( Domain: !Ref RootDomainName )\n"
        ]
    }
}
```

#### YAML

```yaml

Parameters:
  RootDomainName:
    Description: Domain name to use
    Type: String
    Default: testdomain.com
```

#### YAML

```yaml

InputTransformer:
  InputPathsMap:
    "instance" : "$.detail.instance-id"
    "state" : "$.detail.state"
  InputTemplate: !Sub
    - |
      {
         "domain": "${Domain}",
         "instance" : <instance>,
         "state" : <state>,
         "instanceStatus": "instance \"<instance>\" is in <state>"
      }
    - Domain: !Ref RootDomainName
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpParameters

KinesisParameters

All content copied from https://docs.aws.amazon.com/.
