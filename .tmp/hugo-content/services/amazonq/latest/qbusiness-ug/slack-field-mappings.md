# Slack data source connector field mappings

To improve retrieved results and customize the end user chat experience, Amazon Q Business enables you to map document attributes from your data sources to fields
in your Amazon Q index.

Amazon Q offers two kinds of attributes to map to index fields:

- **Reserved or default** – Reserved attributes are
based on document attributes that commonly occur in most data. You can use reserved
attributes to map commonly occurring document attributes in your data source to
Amazon Q index fields.

- **Custom** – You can create custom attributes to map
document attributes that are unique to your data to Amazon Q index
fields.

When you connect Amazon Q to a data source, Amazon Q automatically
maps specific data source document attributes to fields within an Amazon Q index.
If a document attribute in your data source doesn't have a attribute mapping already
available, or if you want to map additional document attributes to index fields, use the
custom field mappings to specify how a data source attribute maps to an Amazon Q
index field. You create field mappings by editing your data source after your application
and retriever are created.

To learn more about document attributes and how they work in Amazon Q, see
[Document attributes and types in Amazon Q](doc-attributes.md).

###### Important

Filtering using document attributes in chat is only supported through the API.

The Amazon Q Slack connector supports the following field
mappings:

Slack field nameIndex field nameDescriptionData typesizesl\_gen\_sizeCustomLong (numeric)emojissl\_gen\_emojisCustomString listtitlesl\_gen\_titleCustomStringauthors\_authorsDefaultString listurl\_source\_uriDefaultStringcategorysl\_gen\_categoryCustomStringcreated\_at\_created\_atDefaultDatelast\_updated\_at\_last\_updated\_atDefaultDatemsg\_channel\_idsl\_message\_channel\_idCustomStringmsg\_channel\_namesl\_msg\_channel\_nameCustomString

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ACL crawling

IAM role

All content copied from https://docs.aws.amazon.com/.
