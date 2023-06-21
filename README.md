# poller-app
This application polls a currency API at a defined interval and saves the response in a json file

## Running the application
In the root directory of the project, there is a makefile, ensure you have make command installed on your machine,
After that following options can be used:
- make tidy // updates dependencies
- make test // Runs all unit tests
- make build // builds a binary for the app in ./cmd directory.
- make clean // deletes the binary created by the build command.
- make run // will run the application.
