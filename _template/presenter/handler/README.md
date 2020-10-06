### Application layer

### OverView

- this layer is like UseCase . This layer receives information from the repository layer.

### How to use

#### Create Handler Function

- At first: Create a function that belongs to a structure
```
ex...

func (e *exampleHandler) ExampleIndex() (*model.Example, error) {
    // abridgement
}
```

- And then: Fill the interface
```
ex...

type ExampleHandler interface {
    ExampleIndex() (*model.Example, error) // additional codes
}
```