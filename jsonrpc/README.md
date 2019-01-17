#JSON-RPC
For usage of the json-rpc api, the user should make a new client 
usingthe NewClient() method  
To run the tests included in this packages, the environment variable
 TESTURL should be set to the endpoint you wish to use. The tests are meant
  to be ran an ark testnet
  
TODO
Add a script which runs the tests and properly sets up a testnet.
go test wont work for this since wallet creation has to happen before
transactions etc.