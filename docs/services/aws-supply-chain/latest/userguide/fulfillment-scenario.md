# Data mapping example for fulfillment

Below is an example to map brick and mortar or online sales to outbound order line dataset and optimize the historical demand setup. Use this example to structure your data for accurate forecasting.
Review the configurations in this example to make sure your forecasting models capture the different fulfillment scenarios.

###### Note

If the data fields _ship\_from\_site\_id_, _ship\_to\_site\_id_, and _channel\_id_ are selected for forecast granularity, make sure they have
values or enter _NULL_ as the value. The forecast will fail if the fields are blank.

Data fieldDescriptionScenario 1 – Store sales (POS)Scenario 2 – E-commerce demand fulfilled by storeScenario 3 – E-commerce demand fulfilled by online fulfillment center (direct to customer)ship\_from\_site\_idSite at which inventory is managedStore IDStore IDFulfillment Center IDship\_to\_site\_idSite that received the orderEnter _NULL_ to avoid forecast failureCountry, Region, State, or Zip – as applicableExternal retailer sore ID, or Country, Region, State, or Zip – as applicablechannel\_idMap how an item is soldBrick and mortarE-commerceE-commerce

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Prequisites before uploading your dataset

Data entities supported in AWS Supply Chain

All content copied from https://docs.aws.amazon.com/.
