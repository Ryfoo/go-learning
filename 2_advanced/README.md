
### Methods

Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v. 

### Methods and Pointer Indirection
for the recievers, things are a bit different than function parameters.
the compiler will handle values passing, in different ways
meaning that if the reciever type should be a pointer, and a value was passed, the compiler will 
automatically attach the address for it, for instance func (v *mytype) my_func(){}
and v was passed instead of &v, the compiler will automaticall pass the address of v instead
the same thing goes with dereferncing (*)

