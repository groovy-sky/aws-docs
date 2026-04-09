# Amazon Q Business Confluence (Cloud) data source connector field mappings

To improve retrieved results and customize the end user chat experience, Amazon Q Business enables you to map document attributes from your data sources to
fields in your Amazon Q index.

Amazon Q offers two kinds of attributes to map to index fields:

- **Reserved or default** – Reserved attributes are
based on document attributes that commonly occur in most data. You can use
reserved attributes to map commonly occurring document attributes in your data
source to Amazon Q index fields.

- **Custom** – You can create custom attributes to
map document attributes that are unique to your data to Amazon Q
index fields.

When you connect Amazon Q to a data source, Amazon Q
automatically maps specific data source document attributes to fields within an Amazon Q index. If a document attribute in your data source doesn't have a
attribute mapping already available, or if you want to map additional document
attributes to index fields, use the custom field mappings to specify how a data source
attribute maps to an Amazon Q index field. You create field mappings by
editing your data source after your application environment and retriever are created.

To learn more about document attributes and how they work in Amazon Q, see
[Document attributes and types in Amazon Q](doc-attributes-types.md).

###### Important

Filtering using document attributes in chat is only supported through the
API.

The Amazon Q
Confluence connector supports the following entities and the associated
reserved and custom attributes.

###### Important

If you map any Confluence (Cloud) field to Amazon Q document title and document body fields,
Amazon Q will generate responses from data in the document title and body.

###### Supported entities and field mappings

- [Space](#confluence-field-mappings-space)

- [Page](#confluence-field-mappings-page)

- [Blog](#confluence-field-mappings-blog)

- [Comment](#confluence-field-mappings-comment)

- [Attachment](#confluence-field-mappings-attachment)

## Space

Confluence field nameIndex field nameDescriptionData typespaceNamecf\_sp\_document\_titleCustomStringitemType\_categoryDefaultStringurl\_source\_uriDefaultStringspaceKeycf\_space\_keyCustomStringdescriptioncf\_descriptionCustomStringspaceTypecf\_typeCustomString

## Page

Confluence field nameIndex field nameDescriptionData typetitle\_cf\_page\_document\_titleCustomStringauthors\_authorsDefaultString listcreatedDate\_created\_atDefaultDatemodifiedDate\_last\_updated\_atDefaultDatelabelscf\_labelsCustomString listversioncf\_versionCustomLong (numeric)itemType\_categoryDefaultStringspaceKeycf\_space\_keyCustomStringspaceNamecf\_space\_nameCustomStringurl\_source\_uriDefaultStringstatuscf\_statusCustomStringparentIdcf\_parent\_idCustomString

## Blog

Confluence field nameIndex field nameDescriptionData typetitlecf\_bg\_document\_titleCustomStringauthor\_authorsDefaultString listpublishedDate\_created\_atDefaultDatelabels\_source\_uriDefaultStringversioncf\_versionCustomLong (numeric)itemType\_categoryCustomStringspaceKeycf\_space\_keyCustomStringmodifiedDate\_last\_updated\_atDefaultDatespaceNamecf\_space\_nameCustomStringstatuscf\_statusCustomStringurl\_source\_uriDefaultStringparentIdcf\_parent\_idCustomString

## Comment

Confluence field nameIndex field nameDescriptionData typetitlecf\_cmt\_document\_titleCustomStringauthor\_authorsDefaultString listcreatedDate\_created\_atDefaultDateversioncf\_versionCustomLong (numeric)itemType\_categoryDefaultStringspaceKeycf\_space\_keyCustomStringspaceNamecf\_space\_nameCustomStringcontentTypecf\_content\_typeCustomStringurl\_source\_uriDefaultStringparentIdcf\_parent\_idCustomStringstatuscf\_statusCustomString

## Attachment

Confluence field nameIndex field nameDescriptionData typefileNamecf\_attachment\_document\_titleCustomStringauthor\_authorsDefaultString listcreatedDate\_created\_atDefaultDatelabelscf\_labelsCustomString listversioncf\_versionCustomLong (numeric)itemType\_categoryDefaultStringspaceKeycf\_space\_keyCustomStringcontentTypecf\_content\_typeCustomStringmodifiedDate\_last\_updated\_atDefaultDatefileSizecf\_file\_sizeCustomLong (numeric)fileTypecf\_attachment\_file\_typeCustomStringspaceNamecf\_space\_nameCustomStringdocumentId\_document\_idDefaultString listurl\_source\_uriDefaultStringparentIdcf\_parent\_idCustomStringattachmentCommentcf\_attachment\_commentCustomStringstatuscf\_statusCustomString

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ACL crawling

IAM role

All content copied from https://docs.aws.amazon.com/.
