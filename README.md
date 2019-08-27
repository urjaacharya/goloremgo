# goloremgo

`goloremgo` is a commandline tool for generating markdown files with random content. `goloremgo` is written in `go`.

## Installation

### Build from source

To build `goloremgo` yourself, go through the following steps. This assumes that `go` is already installed.

```
go get ...
go build
```

### Download Binaries

You can download the binaries for your operating system from here.

## Usage

```
goloremgo -p <path> -s [seed] -f -r
```

where,

- -p (REQUIRED): directory that contains the templates to be processed
- -f (OPTIONAL): overwrites files if they already exist
- -r (OPTIONAL): searches for templates recursively inside all sub-directories
- -s (OPTIONAL): seed for randomly creating content. Default is current time

Other available flags are:

- -h  print this usage information
- -v  version information

## Template Files

### Template File Naming

The template files have the format `LFS_<basename>_LFE_<number_of_files>.glg`

- `LFS_` denotes the start of the template files
- `<basename>` denotes the base name of the files that will be generated
- `_LFE_` follows the `<basename>` and denotes the end of the base name
- `<number_of_files>` denotes the number of the markdown files to be generated by `goloremgo`

For example, if the template name is `LFS_blog_LFE_2.glg`, `goloremgo` will generate 2 markdown files named `blog_1.md` and `blog_2.md`

### Template File Structure

Basically, template files have the structure of markdown files. Examples of template files and markdown files generated from them are available [here](https://github.com/urjaacharya/goloremgo/tree/master/example).

### Available Functions

#### `{{ words n sep }}`

`words` generates random words. It needs two arguments, `n` and `sep`. `n` is an integer representing the number of words to be generated. `sep` is the character which separates the words.

For example, `{{ words 5 " " }}` will generated five words separated by space.

#### `{{ sents n }}`

`sents` generates random sentences. It needs one argument, `n`. `n` is the number of sentences to be generated.

For example, `{{ sents 5 }}` will generate five sentences.

#### `{{ paras numSents numParas }}`

`paras` generates random paragraphs. It needs two arguments, `numSents` and `numParas`. `numSents` is an integer representing the number of sentences in each paragraph to be generated. `numParas` is an integer representing the total number of paragraphs to be generated. 

For example, `{{ paras 3 5 }}` will generate five paragraphs each with three sentences.

#### `{{ date startDate numDays format }}`

`date` generates a random date. It needs three arguments `startDate`, `numDays`, and `format`. `startDate`should have the format `yyyy-mm-dd`. A random date between the `startDate` and `numDays` after it (`startDate` + `numDays`) is returned. `format` represents the format of the date to be returned.

For example, `{{ date "2014-11-12" 250 "Jan 2, 06" }}` will generate a random date between 2014-11-12 and 250 days after it. The returned date will have a `format` similar to Jan 2, 06.

#### `{{ capFirst input }}`

`capFirst` will capitalize the first letter of the first word of the `input` string. 

For example, `{{ capFirst "apple sauce" }}` will return "Apple sauce".

#### `{{ capAll input }}`

`capAll` will capitalize all letters of the `input` string. 

For example, `{{ capAll "apple sauce" }}` will return "APPLE SAUCE".

#### `{{ capEach input }}`

`capEach` will capitalize first letter of each word in the `input` string. 

For example, `{{ capEach "apple sauce" }}` will return "Apple Sauce".
