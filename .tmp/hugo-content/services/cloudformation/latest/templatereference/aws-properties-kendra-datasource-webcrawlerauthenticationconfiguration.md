This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource WebCrawlerAuthenticationConfiguration

Provides the configuration information to connect to websites that require user
authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BasicAuthentication" : [ WebCrawlerBasicAuthentication, ... ]
}

```

### YAML

```yaml

  BasicAuthentication:
    - WebCrawlerBasicAuthentication

```

## Properties

`BasicAuthentication`

The list of configuration information that's required to connect to and crawl a
website host using basic authentication credentials.

The list includes the name and port number of the website host.

_Required_: No

_Type_: Array of [WebCrawlerBasicAuthentication](aws-properties-kendra-datasource-webcrawlerbasicauthentication.md)

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TemplateConfiguration

WebCrawlerBasicAuthentication

All content copied from https://docs.aws.amazon.com/.
