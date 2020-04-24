# Muto

Simple experiment to do live reloading of plugins.

## Usage
cd ./plugin1
go build -buildmode=plugin
cd ../plugin2
go build -buildmode=plugin
cd ..
go build
./muto

