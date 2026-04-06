# Understand error codes in the Amazon Q Business Gmail connector

The following table provides information about error codes you may see for the
Gmail connector and suggested resolutions.

Error codeError messageSuggested resolutionGML-5001There was a problem while retrieving directory.There was a problem while retrieving directory because of incorrect
credentials. Provide correct credentials and try again.GML-5002There was a problem while retrieving user specific Gmail
object.There was a problem while retrieving user specific Gmail object
because of incorrect credentials. Provide correct credentials and try
again.GML-5003Connection lost - A problem occurred while validating
credentials.Connection was lost due to invalid credentials. Provide correct
credentials and try again.GML-5004There was a problem while retrieving the user list because the API
was not responding.There was a problem while retrieving the user list because the API
was not responding. Try again.GML-5100There was a problem while retrieving repository configurations.
Repository configurations may be empty or incorrect.Repository configurations should not be empty or incorrect. Provide
valid details for repository configurations.GML-5101There was a problem while retrieving message entity from repository
configurations. No message entity found in repository
configurations.Message entity should not be empty. Check if message entity is
present in repository configurations and provide the same if not
present.GML-5102There was a problem while retrieving attachment entity from
repository configurations. No attachment entity found in repository
configurations.Attachment entity should not be empty. Check if attachment entity is
present in repository configurations and provide the same if not
present.GML-5103There was a problem while retrieving field mappings for message
entity from repository configurations. Field mappings may be empty or
incorrect.Field mappings should not be empty or incorrect. Provide proper field
mappings for message entity in repository configurations.GML-5104There was a problem while retrieving field mappings for attachment
entity from repository configurations. Field mappings may be empty or
incorrect.Field mappings should not be empty or incorrect. Provide proper field
mappings for message entity in repository configurations.GML-5105There was a problem while retrieving field mapping values for message
entity. Field mapping values may be empty or incorrect.Field mappings values should not be empty or incorrect. Provide
proper field mapping values for message entity in repository
configurations.GML-5106There was a problem while retrieving field mapping values for
attachment entity. Field mapping values may be empty or
incorrect.Field mappings values should not be empty or incorrect. Provide
proper field mapping values for message entity in repository
configurations.GML-5107There was a problem while parsing before/after date filter value.
Before/After date format may be incorrect.Provide correct before/after date format. E.g.
yyyy-MM-ddTHH:mm:ssZ.GML-5108There was a problem while retrieving client email id. Client email id
may be empty or incorrect.The client email id should not be empty or incorrect. Provide correct
client email id.GML-5109There was a problem while retrieving admin account email id. Admin
account email id may be empty or incorrect.The admin account email id should not be empty or incorrect. Provide
correct admin account email id.GML-5110There was a problem while retrieving private key. Private key may be
empty or incorrect.The private key should not be empty or incorrect. Provide correct
private key.GML-5111One or more of the provided filter regex are invalid.Provide correct regex value in filter fields.GML-5200There was a problem while retrieving Gmail items.There was a problem while retrieving Gmail items because user is not
provided. Ensure that user is not empty.GML-5201There was a problem while retrieving the message body because the API
was not responding.There was a problem while retrieving the message body because the API
was not responding. Try again.GML-5202There was a problem while retrieving the message subject because the
API was not responding.There was a problem while retrieving the message subject because the
API was not responding. Try again.GML-5203There was a problem while retrieving the attachment because the API
was not responding.There was a problem while retrieving the attachment because the API
was not responding. Try again.GML-5204There was a problem while retrieving the message metadata because the
API was not responding.There was a problem while retrieving the message metadata because the
API was not responding. Try again.GML-5205There was a problem while retrieving the attachment metadata because
the API was not responding.There was a problem while retrieving the attachment metadata because
the API was not responding. Try again.GML-5206There was a problem while retrieving the message because the API was
not responding.There was a problem while retrieving the message because the API was
not responding. Try again.GML-5500Connection timed out - API is not responding. The threshold number of
API calls has been exceeded.Timeout exception occurred due to API not responding. The threshold
number of API hits has been exceeded. Try again.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IAM role

Google Calendar (Preview)
