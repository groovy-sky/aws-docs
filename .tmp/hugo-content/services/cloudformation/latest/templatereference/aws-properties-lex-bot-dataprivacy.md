This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lex::Bot DataPrivacy

By default, data stored by Amazon Lex is encrypted. The
`DataPrivacy` structure provides settings that determine
how Amazon Lex handles special cases of securing the data for your bot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChildDirected" : Boolean
}

```

### YAML

```yaml

  ChildDirected: Boolean

```

## Properties

`ChildDirected`

For each Amazon Lex bot created with the Amazon Lex Model Building Service,
you must specify whether your use of Amazon Lex is related to a website,
program, or other application that is directed or targeted, in whole or
in part, to children under age 13 and subject to the Children's Online
Privacy Protection Act (COPPA) by specifying `true` or
`false` in the `childDirected` field. By
specifying `true` in the `childDirected` field,
you confirm that your use of Amazon Lex **is**
related to a website, program, or other application that is directed or
targeted, in whole or in part, to children under age 13 and subject to
COPPA. By specifying `false` in the
`childDirected` field, you confirm that your use of Amazon Lex
**is not** related to a website,
program, or other application that is directed or targeted, in whole or
in part, to children under age 13 and subject to COPPA. You may not
specify a default value for the `childDirected` field that
does not accurately reflect whether your use of Amazon Lex is related to a
website, program, or other application that is directed or targeted, in
whole or in part, to children under age 13 and subject to COPPA. If
your use of Amazon Lex relates to a website, program, or other application
that is directed in whole or in part, to children under age 13, you
must obtain any required verifiable parental consent under COPPA. For
information regarding the use of Amazon Lex in connection with websites,
programs, or other applications that are directed or targeted, in whole
or in part, to children under age 13, see the [Amazon Lex\
FAQ](https://aws.amazon.com/lex/faqs).

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomVocabularyItem

DataSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
