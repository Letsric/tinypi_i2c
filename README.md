# tinypi_i2c

A go library to use raspberry pi's I2C with tinygo's I2C based drivers

The library basically works as a bridge between these librarys:

- github.com/ftl/i2c
- tinygo.org/x/drivers

## Example

In this example, I am using the tinygo I2C library to write text to a LCD-screen

```go
package main

import (
  "github.com/Letsric/tinypi_i2c"
  "tinygo.org/x/drivers/hd44780i2c"
)

func main() {
  bus := tinypii2c.New(1)
  screen := hd44780i2c.New(bus, 0x27)

  config := &hd44780i2c.Config{
    Width:  20,
    Height: 4,
  }
  screen.Configure(*config)

  screen.Print([]byte("Hello World!"))
}
```
