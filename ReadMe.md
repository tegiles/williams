# Williams Race Viewer

## Backend

The backend of the race viewer is a Golang Application, running a basic HTTP server, providing the required F1 data. It runs on port 8080 of the local machine. 

### Running

1. Navigate to the backend folder, where you will find the main.go file. 
2. Run 'go run main.go' from the command line
3. To stop, simply end execution of the command line process

### Requirements

Golang Installation: https://go.dev/doc/install

## Frontend

The frontend of the race viewer is an Angular frontend component. It offers two tabs to view either circuit or driver information. You can click on entries in each table to get more information on the entry. 

### Running

1. Navigate to frontend/race-viewer folder
2. Run 'npm install' to ensure all packages are installed
3. Run 'ng serve --open' from the command line
4. The application will be available on the browser at localhost:4200

### Requirements

Angular Installation: https://v17.angular.io/guide/setup-local
