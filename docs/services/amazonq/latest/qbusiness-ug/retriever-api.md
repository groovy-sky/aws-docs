# Creating a retriever for an Amazon Q Business application using APIs

You can't create an Amazon Q Business retriever using the AWS Management Console. If you use
the console, Amazon Q Business creates a retriever for you when you create an
Amazon Q Business index. If you're using an Amazon Kendra index as retriever, adding
an Amazon Kendra index will also add a retriever.

###### Note

You can only add an Amazon Kendra index as retriever if you have existing Amazon Kendra
indexes.

When you use the APIs, you must create a retriever for your Amazon Q Business
index separately.

API actionAPI descriptionRelevant User Guide topic[CreateRetriever](../api-reference/api-createretriever.md)Creates an Amazon Q Business or Amazon Kendra retriever

- [Creating an Amazon Q Business\
retriever](native-retriever.md)

- [Creating an Amazon Kendra\
retriever](add-kendra-retriever.md)

[DeleteRetriever](../api-reference/api-deleteretriever.md)Deletes an Amazon Q Business or Amazon Kendra retriever

- [Deleting an Amazon Q Business\
retriever](supported-native-retriever-actions.md#delete-native-retriever)

- [Deleting an Amazon Kendra\
retriever](supported-kendra-retriever-actions.md#delete-kendra-retriever)

[GetRetriever](../api-reference/api-getretriever.md)Gets information about an existing Amazon Q Business or Amazon Kendra
retriever

- [Getting Amazon Q Business\
retriever properties](supported-native-retriever-actions.md#describe-native-retriever)

- [Getting Amazon Kendra\
retriever properties](supported-kendra-retriever-actions.md#describe-kendra-retriever)

[ListRetrievers](../api-reference/api-listretrievers.md)Lists existing Amazon Q Business or Amazon Kendra retrievers

- [Listing retrievers](supported-native-retriever-actions.md#list-native-retriever)

- [Getting Amazon Kendra\
retriever properties](supported-kendra-retriever-actions.md#list-kendra-retriever)

[UpdateRetriever](../api-reference/api-updateretriever.md)Updates an existing Amazon Q Business or Amazon Kendra
retriever

- [Updating an Amazon Q Business\
retriever](supported-native-retriever-actions.md#update-native-retriever)

- [Updating an Amazon Kendra\
retriever](supported-kendra-retriever-actions.md#update-kendra-retriever)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Creating an index

Connecting a data source

All content copied from https://docs.aws.amazon.com/.
