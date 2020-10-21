### Application layer

### OverView

- this layer is like UseCase . This layer receives information from the repository layer.

### How to use

#### Create Repository Handler

- At first: Create a function that belongs to a structure
```
ex...

func (e *exampleRepository) Fetch() (*model.Example, error) {
    // abridgement
}
```

- And then: Fill the interface
```
ex...

type ExampleRepository interface {
    Fetch() (*model.Example, error) // additional codes
}
```