# supplementary\_time\_series

###### Note

If you cannot locate the supplementary\_time\_series data entity, your instance might be using an older data model version. You can contact AWS Support to upgrade your data model version or create a new data connection.

**Primary key (PK)**

The table below lists the colum names that are uniquely identified in the data entity.

NameColumnforecast\_supplementary\_time\_seriesid

The table below lists the column names supported by the data entity:

ColumnData typeRequiredDescription

id

string

Yes

Unique identifier with each supplementary data entry.

product\_id 2

string

No

Unique identifier for a specific product. Corresponds to product\_id in the outbound\_order\_line dataset.

product\_group\_id

string

No

Product hierarchy or grouping.

order\_date

timestamp

Yes1

The timestamp indicating the date and time when the date for the respective time-series was recorded.

channel\_id

string

No

Unique identifier for a specific product. Corresponds to product\_id in the outbound\_order\_line dataset.

customer\_tpartner\_id 2

string

No

Unique identifier for a specific user. Corresponds to customer\_tpartner\_id field in outbound\_order\_line dataset.

site\_id 2

string

No

Unique identifier for a specific site or location.

ship\_to\_site\_id 2

string

No

Unique identifier for a specific site or location. This corresponds to the _ship\_to\_site\_id_ in the _outbound\_order\_line_ dataset.

ship\_to\_site\_address\_zip

string

No

Postal code of _ship\_to\_site\_id_.

geo\_id 2

string

No

Geographical hierarchy ID.

ship\_from\_site\_id 2

string

No

Corresponds to the _ship\_from\_site\_id_ in the _outbound\_order\_line_ dataset.

ship\_from\_site\_address\_zip

string

No

Postal code of _ship\_from\_site\_id_.

time\_series\_name

string

Yes

The _time\_series\_name_ must start with a letter, should be 2 to 56 characters long, and can contain letters, numbers, and underscores. No other special characters are allowed.

time\_series\_value

string

Yes

Value corresponding to the specific time series. This could represent quantities, metric, or string that is relevant to the type of the data. Demand planning only supports numerical value as additional forecast input.

source\_event\_id

string

No

ID of the event created in the source system.

source\_update\_dttm

timestamp

No

Date time stamp of the update made in the source
system.

1You must enter a value. When you ingest data from SAP
or EDI, the default value for _string_ is SCN\_RESERVED\_NO\_VALUE\_PROVIDED.

2Foreign key

**Foreign key (FK)**

The table below lists the columns with the associated foreign key.

ColumnCategoryFK/Data entityFK/Columnproduct\_idProductproductidsite\_idNetworksiteidcustomer\_tpartner\_idOrganizationtrading\_partneridship\_to\_site\_idOutbound fulfilmentoutbound\_order\_lineship\_to\_site\_idgeo\_idOrganizationgeographyidship\_from\_site\_idOutbound fulfilmentoutbound\_order\_lineship\_from\_site\_id

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Forecast

forecast

All content copied from https://docs.aws.amazon.com/.
