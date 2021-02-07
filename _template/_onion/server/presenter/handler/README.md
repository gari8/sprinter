### Presenter layer

### OverView

- This layer receives information from the application layer.

### How to use

#### Create Presenter Handler

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
