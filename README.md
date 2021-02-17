# 4mat

## -- Work in progress --
A simple text formatting tool to convert between well known text formats.  
Streams one or more source files into a single format.

Supports the following text formats:

- yaml
- json
- xml
- csv
- pdf

Converts to or from any of the supported formats.

### Usage

`4mat json myyamlfile.yaml`  
Will convert the single `myyamlfile.yaml` into json.

`4mat json myyamlfile.yaml myxml.xml myexcel.csv`  
Will convert the three files, from their respective formats,  
into json.

`4mat yaml myinfodoc.pdf`  
Converts the pdf file into yaml.


  
