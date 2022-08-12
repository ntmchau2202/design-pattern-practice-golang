Project structures
- entities: domain models; a set of data structure and functions that has wide enterprise business rules
- drivers: frameworks and drivers to connect to different tools like API, database, web framework, etc.
- uc (usecase): application business rules using domain models (entities); having
    - input ports: handling data from outer layer to uc
    - output ports: handling data from uc to outer layer
- interfaces: handling communication between inner and outer layer, concern only about tech logical, not business logic, eg. conversion, formating, etc.

Layer rules: drivers -> interfaces -> uc -> entities