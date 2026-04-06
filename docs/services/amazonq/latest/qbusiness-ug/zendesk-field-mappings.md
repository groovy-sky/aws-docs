# Zendesk data source connector field mappings

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

The Amazon Q
Zendesk connector supports the following entities and the associated reserved
and custom attributes.

###### Supported entities and field mappings

- [Tickets](#zendesk-field-mappings-tickets)

- [Ticket comments](#zendesk-field-mappings-ticket-comments)

- [Ticket comment attachment](#zendesk-field-mappings-ticket-comment-attachment)

- [Article](#zendesk-field-mappings-article)

- [Article comment](#zendesk-field-mappings-article-comment)

- [Article comment attachment](#zendesk-field-mappings-article-comment-attachment)

- [Community topic](#zendesk-field-mappings-community-topic)

- [Community post](#zendesk-field-mappings-community-post)

- [Community post comment](#zendesk-field-mappings-community-post-comment)

## Tickets

Amazon Q supports crawling [Zendesk Tickets](https://developer.zendesk.com/api-reference/ticketing/tickets/tickets) and offers the following ticket field
mappings.

Zendesk field nameIndex field nameDescriptionData typeticketChannelzd-channelCustomStringcategory\_categoryDefaultStringauthors\_authorsDefaultString listassigneezd\_assigneeCustomStringtagszd\_tagsCustomString liststatuszd\_statusCustomStringsourceUrl\_source\_uriDefaultStringcreatedAt\_created\_atDefaultDateupdatedAt\_last\_updated\_atDefaultDateorganizationNamezd\_organization\_nameCustomString

## Ticket comments

Amazon Q supports crawling [Zendesk Ticket Comments](https://developer.zendesk.com/api-reference/ticketing/tickets/ticket_comments) and offers the following ticket
comment field mappings.

Zendesk field nameIndex field nameDescriptionData typecategory\_categoryDefaultStringauthors\_authorsDefaultString liststatuszd\_statusCustomStringsourceUrl\_source\_uriDefaultStringcreatedAt\_created\_atDefaultDateupdatedAt\_last\_updated\_atDefaultDateorganizationNamezd\_organization\_nameCustomString

## Ticket comment attachment

Amazon Q supports crawling [Zendesk Ticket Comment Attachments](https://developer.zendesk.com/api-reference/ticketing/tickets/ticket-attachments) and offers the
following ticket comment attachment field mappings.

Zendesk field nameIndex field nameDescriptionData typecategory\_categoryDefaultStringauthors\_authorsDefaultString liststatuszd\_statusCustomStringsourceUrl\_source\_uriDefaultStringcreatedAt\_created\_atDefaultDateupdatedAt\_last\_updated\_atDefaultDateorganizationNamezd\_organization\_nameCustomString

## Article

Amazon Q supports crawling [Zendesk Articles](https://developer.zendesk.com/api-reference/help_center/help-center-api/articles) and offers the following article field
mappings.

Zendesk field nameIndex field nameDescriptionData typeauthors\_authorsDefaultString listlabelszd\_article\_labelsCustomString listsectionzd\_article\_sectionCustomString listsourceUrl\_source\_uriDefaultStringcreatedAt\_created\_atDefaultDateupdatedAt\_last\_updated\_atDefaultDate

## Article comment

Amazon Q supports crawling [Zendesk Article Comments](https://developer.zendesk.com/api-reference/help_center/help-center-api/article_comments) and offers the following
article comment field mappings.

Zendesk field nameIndex field nameDescriptionData typeauthors\_authorsDefaultString listlabelszd\_article\_labelsCustomString listsectionzd\_article\_sectionCustomString listsourceUrl\_source\_uriDefaultStringcreatedAt\_created\_atDefaultDateupdatedAt\_last\_updated\_atDefaultDate

## Article comment attachment

Amazon Q supports crawling [Zendesk Article Comment Attachments](https://developer.zendesk.com/api-reference/ticketing/tickets/ticket-attachments) and offers the
following article comment attachment field mappings.

Zendesk field nameIndex field nameDescriptionData typeauthors\_authorsDefaultString listlabelszd\_article\_labelsCustomString listfileNamezd\_file\_nameCustomStringfileType\_file\_typeDefaultStringfileSizezd\_file\_sizeCustomLong (numeric)sectionzd\_article\_sectionCustomString listsourceUrl\_source\_uriDefaultStringcreatedAt\_created\_atDefaultDateupdatedAt\_last\_updated\_atDefaultDate

## Community topic

Amazon Q supports crawling [Zendesk Community Topics](https://developer.zendesk.com/api-reference/help_center/help-center-templates/community_topic_page) and offers the following
community topic field mappings.

Zendesk field nameIndex field nameDescriptionData typetopicNamezd\_topic\_nameCustomStringsourceUrl\_source\_uriDefaultStringcreatedAt\_created\_atDefaultDateupdatedAt\_last\_updated\_atDefaultDatecategory\_categoryDefaultString

## Community post

Amazon Q supports crawling [Zendesk Community Posts](https://developer.zendesk.com/api-reference/help_center/help-center-templates/community_post_page) and offers the following
community post field mappings.

Zendesk field nameIndex field nameDescriptionData typepostNamezd\_post\_nameCustomStringtopicNamezd\_topic\_nameCustomStringsourceUrl\_source\_uriDefaultStringcreatedAt\_created\_atDefaultDateupdatedAt\_last\_updated\_atDefaultDatecategory\_categoryDefaultString

## Community post comment

Amazon Q supports crawling [Zendesk Community Post Comments](https://developer.zendesk.com/api-reference/help_center/help-center-api/post_comments) and offers the following
community post comment field mappings.

Zendesk field nameIndex field nameDescriptionData typepostNamezd\_post\_nameCustomStringtopicNamezd\_topic\_nameCustomStringsourceUrl\_source\_uriDefaultStringcreatedAt\_created\_atDefaultDateupdatedAt\_last\_updated\_atDefaultDatecategory\_categoryDefaultString

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ACL crawling

IAM role
