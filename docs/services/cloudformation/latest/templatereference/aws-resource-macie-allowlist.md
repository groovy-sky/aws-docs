This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Macie::AllowList

The `AWS::Macie::AllowList` resource specifies an allow list. In Amazon Macie, an allow list defines specific text or a text pattern for Macie to ignore when it inspects data sources for sensitive data. If data
matches text or a text pattern in an allow list, Macie doesn’t report the
data in sensitive data findings or sensitive data discovery results, even if the data
matches the criteria of a custom data identifier or a managed data identifier. You can
create and use allow lists in all the AWS Regions where Macie is currently available except the Asia Pacific (Osaka) Region.

Macie supports two types of allow lists:

- **Predefined text** \- For this type of list
( `S3WordsList`), you create a line-delimited plaintext file that
lists specific text to ignore, and you store the file in an Amazon Simple Storage Service (Amazon S3) bucket. You then configure settings
for Macie to access the list in the bucket.

This type of list typically contains specific words, phrases, and other kinds
of character sequences that aren’t sensitive, aren't likely to change, and don’t
necessarily adhere to a common pattern. If you use this type of list, Macie doesn't report occurrences of text that exactly match a
complete entry in the list. Macie treats each entry in the list
as a string literal value. Matches aren't case sensitive.

- **Regular expression** \- For this type of list
( `Regex`), you specify a regular expression that defines a text
pattern to ignore. Unlike an allow list with predefined text, you store the
regex and all other list settings in Macie.

This type of list is helpful if you want to specify text that isn’t sensitive
but varies or is likely to change while also adhering to a common pattern. If
you use this type of list, Macie doesn't report occurrences of
text that completely match the pattern defined by the list.

For more information, see [Defining sensitive data exceptions with\
allow lists](../../../macie/latest/user/allow-lists.md) in the _Amazon Macie User_
_Guide_.

An `AWS::Macie::Session` resource must exist for an AWS account before you can create an `AWS::Macie::AllowList`
resource for the account. Use a [DependsOn\
attribute](../userguide/aws-attribute-dependson.md) to ensure that an `AWS::Macie::Session` resource is
created before other Macie resources are created for an account. For
example, `"DependsOn": "Session"`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Macie::AllowList",
  "Properties" : {
      "Criteria" : Criteria,
      "Description" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Macie::AllowList
Properties:
  Criteria:
    Criteria
  Description: String
  Name: String
  Tags:
    - Tag

```

## Properties

`Criteria`

The criteria that specify the text or text pattern to ignore. The criteria can be the
location and name of an Amazon S3 object that lists specific text to ignore
( `S3WordsList`), or a regular expression ( `Regex`) that
defines a text pattern to ignore.

_Required_: Yes

_Type_: [Criteria](aws-properties-macie-allowlist-criteria.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A custom description of the allow list. The description can contain 1-512
characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A custom name for the allow list. The name can contain 1-128 characters.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to the allow list.

For more information, see [Resource tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-macie-allowlist-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the `AllowList`. For example,
`{ "Ref": "AllowList" }`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the allow list.

`Id`

The unique identifier for the allow list.

`Status`

The current status of the allow list, which indicates whether Amazon Macie
can access and use the list's criteria. If the list's criteria specify a regular
expression ( `Regex`), this value is typically `OK`. Macie can compile the expression. If the list's criteria specify an Amazon S3 object ( `S3WordsList`), possible values are:

- `OK` \- Macie can retrieve and parse the contents of
the object.

- `S3_OBJECT_ACCESS_DENIED` \- Macie isn't allowed to
access the object or the object is encrypted with a customer managed AWS KMS key that Macie isn't allowed to use. Check the
bucket policy and other permissions settings for the bucket and the object. If
the object is encrypted, also ensure that it's encrypted with a key that Macie is allowed to use.

- `S3_OBJECT_EMPTY` \- Macie can retrieve the object
but the object doesn't contain any content. Ensure that the object contains the
correct entries. Also ensure that the list's criteria specify the correct bucket
and object names.

- `S3_OBJECT_NOT_FOUND` \- The object doesn't exist in Amazon S3. Ensure that the list's criteria specify the correct bucket and
object names.

- `S3_OBJECT_OVERSIZE` \- Macie can retrieve the
object. However, the object contains too many entries or its storage size
exceeds the quota for an allow list. Try breaking the list into multiple files
and ensure that each file doesn't exceed any quotas. Then configure list
settings in Macie for each file.

- `S3_THROTTLED` \- Amazon S3 throttled the request to
retrieve the object. Wait a few minutes and then try again.

- `S3_USER_ACCESS_DENIED` \- Amazon S3 denied the request to
retrieve the object. If the specified object exists, you're not allowed to
access it or it's encrypted with an AWS KMS key that you're not
allowed to use. Work with your AWS administrator to ensure that
the list's criteria specify the correct bucket and object names, and you have
read access to the bucket and the object. If the object is encrypted, also
ensure that it's encrypted with a key that you're allowed to use.

- `UNKNOWN_ERROR` \- A transient or internal error occurred when
Macie attempted to retrieve or parse the object. Wait a few
minutes and then try again. A list can also have this status if it's encrypted
with a key that Amazon S3 and Macie can't access or
use.

For more information, see [Allow list options and\
requirements](../../../macie/latest/user/allow-lists-options.md) in the _Amazon Macie User_
_Guide_.

## Examples

The following examples demonstrate how to declare an
`AWS::Macie::AllowList` resource.

- [Creating an allow list that specifies a regular expression](#aws-resource-macie-allowlist--examples--Creating_an_allow_list_that_specifies_a_regular_expression)

- [Creating an allow list that specifies predefined text](#aws-resource-macie-allowlist--examples--Creating_an_allow_list_that_specifies_predefined_text)

### Creating an allow list that specifies a regular expression

This example creates an allow list that uses a regular expression to specify a
text pattern to ignore. The list is designed to ignore specific email addresses
that a custom data identifier might otherwise detect and report in sensitive data findings: all email
addresses for the _example.com_ domain.

#### JSON

```json

{
    "Type": "AWS::Macie::AllowList",
    "DependsOn": "Session",
    "Properties": {
        "Criteria": {
            "Regex": "[a-z]@example.com"
        },
        "Description": "Ignores email addresses for Example Corp.",
        "Name": "MyAllowList",
        "Tags": [
            {
                "Key": "Stack",
                "Value": "Production"
            }
        ]
    }
}
```

#### YAML

```yaml

Type: 'AWS::Macie::AllowList'
DependsOn: Session
Properties:
  Criteria:
    Regex: '[a-z]@example.com'
  Description: Ignores email addresses for Example Corp.
  Name: MyAllowList
  Tags:
    - Key: Stack
      Value: Production
```

### Creating an allow list that specifies predefined text

This example creates an allow list that specifies a list of predefined text to
ignore. The list is designed to ignore specific phone numbers that a managed data
identifier might otherwise detect and report in sensitive data findings: all public
phone numbers for a company named Example Corp. The list is stored in an Amazon S3 object
named `AllowLists/Macie/MyList.txt`. The object is stored in an S3 bucket
named `
                            amzn-s3-demo-bucket
                        `.

#### JSON

```json

{
    "Type": "AWS::Macie::AllowList",
    "DependsOn": "Session",
    "Properties": {
        "Criteria": {
            "S3WordsList": {
                "BucketName": "amzn-s3-demo-bucket",
                "ObjectKey": "AllowLists/Macie/MyList.txt"
            }
        },
        "Description": "Lists all public phone numbers for Example Corp.",
        "Name": "MyAllowList",
        "Tags": [
            {
                "Key": "Stack",
                "Value": "Production"
            }
        ]
    }
}
```

#### YAML

```yaml

Type: 'AWS::Macie::AllowList'
DependsOn: Session
Properties:
  Criteria:
    S3WordsList:
      BucketName: amzn-s3-demo-bucket
      ObjectKey: AllowLists/Macie/MyList.txt
  Description: Lists all public phone numbers for Example Corp.
  Name: MyAllowList
  Tags:
    - Key: Stack
      Value: Production
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Macie

Criteria

All content copied from https://docs.aws.amazon.com/.
