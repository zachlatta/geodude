# Geodude

<img src="http://i.imgur.com/IlE7oAi.png" alt="Geodude Icon" align="right" />
Geodude is a tiny command-line tool for geocoding addresses.

#### Features

* Geocodes and reverse geocodes addresses
* Minimalist interface
* Fast

### Installation

    $ go get github.com/zachlatta/geodude

### Usage

    $ geodude [address]

#### Examples

```
$ geodude 1600 amphitheatre parkway

Address: 1600 Amphitheatre Parkway, Mountain View, CA 94043, USA
Coordinates: 37.4219998, -122.0839596
```

```
$ geodude 1 infinite loop

Address: 1 Infinite Loop, Cupertino, CA 95014, USA
Coordinates: 37.331741, -122.0303329
```

## License

The MIT License (MIT)

Copyright (c) 2014 Zach Latta

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
