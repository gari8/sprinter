### Application layer

### OverView

- this layer is like UseCase . This layer receives information from the repository layer.

### How to use

#### Create Application Handler

- At first: Create a function that belongs to a structure
```
ex...

func (e *exampleApplication) GetExample() (*model.Example, error) {
    // abridgement
}
```

- And then: Fill the interface
```
ex...

type ExampleApplication interface {
    GetExample() (*model.Example, error) // additional codes
}
```
