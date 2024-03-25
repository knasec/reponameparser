# Repository and image names parser

This library is designed to parse repository names for more convenient use when you need to integrate with some external API.
For example to pass only project name, or group/repository pair

# Supported paths

## Git

### SSH
*State:* To do
```
ssh://[user@]example.com[:port]/subgroup/repository.git
```
### Git
*State:* To do
```
git://example.com[:port]/subgroup/repository.git
```
### http[s]
*State:* To do
```
http[s]://example.com[:port]/subgroup/repository.git
```
### ftp[s]
*State:* To do
```
ftp[s]://example.com[:port]/psubgroup/repository.git
```
### Scp
*State:* Parsing done. Need checker for pattern
```
[user@]example.com:group/subgroup/repository.git
```
## Docker Images

###
*State:* To do
```
example.com/imagename:version
```