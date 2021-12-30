# DisneyLandPath
A project built to find the minimum path to reach everyride along the DisneyLand park (Anaheim). The project uses around a years worth of data of ride wait times along with a custom made distance dataset relating the distances of each ride to each other. 

**Run Instructions**
Inorder to run `main.go` run the command 
```
go run main.go
``` 
inside the main directory

**Build Instructions**
Inorder to compile the `main.go` file down to an exectuable `main` run the command 
```
go build main.go
```
**File Locations**
`CSV` contains all relevant CSV file sheets for ride times, walk times, and ride ids. `graph` contains the main files pertaining to the graph structure including the traversal functions. Inside the `graph` package includes the graph struct, the node struct, and the edge. Inside `utils` is files relating to parsing and modifying data from the CSVs.

This project is original based off a CS225 Final Project built by Eli Feinberg (elihf2), Sreyas Dhulipala (sreyasd2), and Jack Wang (jackw5).

